package database

import (
	"fmt"
	"os"

	"github.com/gKits/sessionauthapi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func ConnectDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPasswd)
    fmt.Println(dbURI)
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

func GetUserBy(criteria, searchkey string) (models.User, error) {
    user := models.User{}

    db, err := ConnectDB()
    if err != nil {
        return models.User{}, err
    }
    defer db.Close()

    err = db.Model(&models.User{}).Where(fmt.Sprintf("%s = ?", criteria), searchkey).Take(&user).Error
    if err != nil {
        return models.User{}, err
    }

    user.Password = ""
    return user, nil
}

func ValidateUser(user models.User) error {
    validation := models.User{}

    db, err := ConnectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    err = db.Model(&models.User{}).Where("Email = ?", user.Email).Take(&validation).Error
    if err != nil {
        return err
    }

	err = bcrypt.CompareHashAndPassword([]byte(validation.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

    return nil
}
