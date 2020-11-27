package user

import (
	"fiber-go-crud/database"
	"fiber-go-crud/models"
	"log"
)

// Get all users
func GetUsers() (models.Users, error) {
	result := models.Users{}
	db := database.ConnectPgInit()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err.Error())
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			log.Fatal(err.Error())
			return result, err
		}
		result = append(result, user)
	}

	return result, nil
}
