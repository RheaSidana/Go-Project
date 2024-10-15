package user

import "go-project/initializer"

func initRepository() Repository{
	return NewRepository(initializer.Db)
}

func initHandler(repository Repository) Handler{
	return Handler{
		repository: repository,
	}
}