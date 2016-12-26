package main

import (
	"fmt"
	"go-mssqldb-crud/managers"
	"go-mssqldb-crud/models"

	uuid "github.com/satori/go.uuid"
)

func main() {

	fmt.Println("\n ------ PURE SQL ------------ \n ")
	fmt.Println("------ GetUsers ------------")
	resultDb := managers.GetUsers()
	fmt.Printf("There are users %d in DB \n\n", len(resultDb))

	fmt.Println("------ GetUsers via SP ------------")
	resultDb = managers.GetUsersSP()
	fmt.Printf("There are users %d in DB \n\n", len(resultDb))

	fmt.Println("------ GetUsersById ------------")
	userDB := managers.GetUserById(resultDb[0].Id)
	fmt.Println(resultDb[0].ToString())

	fmt.Println("------ UpdateUser ------------")

	userDB.Nick = "testUser"
	userDB.Adr1 = "Adr_" + uuid.NewV4().String()
	userDB.Ville = "Kiev"

	result := managers.UpdateUser(userDB)
	fmt.Println("User was updated in DB :", result)

	updatedUser := managers.GetUserByNick(userDB.Nick)
	fmt.Println(updatedUser.ToString())

	fmt.Println("------ InsertUser ------------")

	userObj := models.User{
		Adr1:   "AIns-" + uuid.NewV4().String(),
		Email:  "test@gmail.com",
		Nick:   "testNick",
		Nom:    "testNom",
		Prenom: "testPrenom",
		Ville:  "testVille"}

	fmt.Println(userObj.ToString())

	isInserted := managers.InsertUser(userObj)
	fmt.Println("User was insert : ", isInserted)

	fmt.Println("------ GetUsersByNick ------------")
	insertedUser := managers.GetUserByNick(userObj.Nick)
	fmt.Println(insertedUser.ToString())

	fmt.Println("------ DeleteUser ------------")
	deleteUser := managers.GetUserByNick(insertedUser.Nick)
	managers.DeleteUser(deleteUser.Id)

	fmt.Println("------ Test if user was deleted ------------")
	deletedUser1 := managers.GetUserByNick(deleteUser.Nick)
	fmt.Println(deletedUser1.ToString())

}
