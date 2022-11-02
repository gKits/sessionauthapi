package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/gKits/sessionauthapi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	dbUser := os.Getenv("databaseUser")
	dbPasswd := os.Getenv("databasePassword")
	dbName := os.Getenv("databaseName")
	dbHost := os.Getenv("databaseHost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPasswd)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	fmt.Println("successfully connected to database!")

	return db, nil
}

func CreateUser(u *models.User) (*models.User, error) {
	db, err := ConnectDB()
	if err != nil {
		return &models.User{}, err
	}
	defer db.Close()

	err = db.Create(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	return u, nil
}

func GetUserByEmail(email string) (models.User, error) {
	u := models.User{}

	db, err := ConnectDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	err = db.Model(&models.User{}).Where("Email = ?", email).Take(&u).Error
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func GetUserById(uid uint) (models.User, error) {
	u := models.User{}

	db, err := ConnectDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	err = db.First(&models.User{}).Where("ID = ?", uid).Take(&u).Error
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	u.Password = ""
	return u, nil
}
