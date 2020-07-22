package Models

import (
	"first-api-go/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Handle Request

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//Create User
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//Fetch User from ID
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//Update User
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//Delete User
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
