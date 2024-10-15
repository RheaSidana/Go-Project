package main

import (
	"fmt"
	"go-project/initializer"
	"go-project/migrations/functionality"
	"log"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	defer initializer.CloseDb()
 
	for _, migrationFile := range functionality.MigrationFiles {
		sqlFile, err := functionality.OpenFile(migrationFile)
		if err != nil {
			fmt.Print(err)
		}
		defer sqlFile.Close()

		var sqlString string = functionality.ReadSqlFile(sqlFile)

		err = functionality.ExecuteCommand(sqlString)
		if err != nil {
			fmt.Print(err)
		}

		log.Println("Migration ran successfully for ", migrationFile, "!")
	}

}
