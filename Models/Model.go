package Models

import (
	"fmt"
	"todo-app/Config"

	_ "github.com/go-sql-driver/mysql"
)

//fetch all todos
func GetAllTodos(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//insert a todo
func CreateATodo(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//Get a todo
func GetATodo(todo *Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//Update a todo
func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

//Delete a todo
func DeleteATodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
