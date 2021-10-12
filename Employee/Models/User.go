package Models

import (
	"API_Golang/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(employee *[]Employee) (err error) {
	if err = Config.DB.Find(employee).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(employee *Employee) (err error) {
	if err = Config.DB.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(employee *Employee, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(employee).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(employee *Employee, id string) (err error) {
	fmt.Println(employee)
	Config.DB.Save(employee)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(employee *Employee, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(employee)
	return nil
}
