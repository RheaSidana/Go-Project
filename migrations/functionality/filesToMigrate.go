package functionality

import (
	"bufio"
	"fmt"
	"go-project/initializer"
	"os"
	"strings"
)

var MigrationFiles = []string{
	"./migrations/user/001_create_users_table.sql",
}

func OpenFile(file string) (*os.File, error) {
	sqlFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error reading migration file: %v", err)
	}

	return sqlFile, nil
}

func ReadSqlFile(sqlFile *os.File) string {
	var sqlBytes string
	scanner := bufio.NewScanner(sqlFile)
	for scanner.Scan() {
		fmt.Println("File reading!")
		line := scanner.Text()
		sqlBytes += strings.Trim(line, " ") + " "
	}
	fmt.Println("File read!", sqlBytes)

	return sqlBytes
}

func ExecuteCommand(sqlBytes string) (error) {
	_, err := initializer.Db.Exec(sqlBytes)
	if err != nil {
		return fmt.Errorf("error executing migration: %v", err)
	}

	return err
}
