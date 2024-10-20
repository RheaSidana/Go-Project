package user

import (
	"encoding/json"

	models "go-project/models"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	repository Repository
}

func (h *Handler) PatchUserHandler(w http.ResponseWriter, r *http.Request) {
	var userUpdates models.User
	err := json.NewDecoder(r.Body).Decode(&userUpdates)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.repository.PatchUser(&userUpdates)
	if err != nil {
		http.Error(w, "Error patching user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		User: *updatedUser,
		Message: "User patched successfully",
	})
}
func (h *Handler) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.repository.UpdateUser(&user)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		User: *updatedUser,
		Message: "User updated successfully",
	})
}
func (h *Handler) DeleteAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.repository.DeleteAll()
	if err != nil {
		http.Error(w, "Error deleting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseList{
		Users: users,
		Message: "All users deleted successfully!",
	})
}

func (h *Handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.repository.GetAll()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseList{
		Users: users,
		Message: "All users returned successfully!",
	})
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL format. Expected /users/{id}", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.repository.Delete(id)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		User: *user,
		Message: "Delete successfully",
	})
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.repository.GetByEmail(newUser.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = isEncryptedPassword(user.Password, newUser.Password)
	if err != nil {
		http.Error(w, "invalid password provided", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		User: *user,
		Message: "user logged in successfully",
	})
}

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	encryptPassword, err := hashPassword(newUser.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	newUser.Password = encryptPassword

	createdUser, err := h.repository.Create(&newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL format. Expected /users/{id}", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.repository.Get(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := h.repository.Create(&newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}
