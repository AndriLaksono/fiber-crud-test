package user

import (
	"fiber-go-crud/database"
	"log"
)

func StoreUser(username string, password string) error {
	db := database.ConnectPgInit()
	_, err := db.Exec("INSERT INTO users (username, password) VALUES($1,$2)", username, password)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
