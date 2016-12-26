package managers

import (
	"fmt"
	"go-mssqldb-crud/models"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const connectionString string = "driver={SQL Server};server=.\\GO;database=LocalDb;user id=testUser;password=123;"

func getSqlDB() *sqlx.DB {

	db, err := sqlx.Connect("mssql", connectionString)
	if err != nil {
		fmt.Println("Fatal error : ", err)
		return nil
	}

	return db
}

func UpdateUser(model models.User) bool {

	if (models.User{}) == model || model.Id <= 0 {
		return false
	}

	db := getSqlDB()
	if db == nil {
		return false
	}

	defer db.Close()

	_, err := db.Exec(
		"UPDATE [dbo].[Users] SET [Nick] = ? ,[EMail] = ? ,[Nom] = ? ,[Prenom] = ? ,[Adr1] = ? ,[Ville] = ? WHERE [Id] = ?",
		model.Nick,
		model.Email,
		model.Nom,
		model.Prenom,
		model.Adr1,
		model.Ville,
		model.Id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func InsertUser(model models.User) bool {

	if (models.User{}) == model {
		return false
	}

	db := getSqlDB()
	if db == nil {
		return false
	}

	defer db.Close()

	_, err := db.Exec(
		"INSERT INTO [dbo].[Users] ([Nick],[EMail],[Nom],[Prenom],[Adr1],[Ville]) VALUES (? ,? ,? ,? ,? ,?)	",
		model.Nick,
		model.Email,
		model.Nom,
		model.Prenom,
		model.Adr1,
		model.Ville)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func getUser(userId int, nick string) models.User {

	var result models.User

	db := getSqlDB()
	if db == nil {
		return result
	}

	defer db.Close()
	var rows *sqlx.Rows
	var err error

	if userId <= 0 {
		rows, err = db.Queryx("SELECT * FROM [dbo].[Users] WITH(NOLOCK) WHERE [Nick] = ?", nick)
	} else {
		rows, err = db.Queryx("SELECT * FROM [dbo].[Users] WITH(NOLOCK) WHERE [Id] = ?", userId)
	}

	if err != nil {
		fmt.Println(err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		err1 := rows.StructScan(&result)
		if err1 != nil {
			fmt.Println("Scan: %v", err)
		}

		break
	}

	return result
}

func GetUserByNick(nick string) models.User {

	var result models.User
	if len(nick) == 0 {
		return result
	}

	return getUser(0, nick)

}

func GetUserById(userId int) models.User {

	var result models.User
	if userId <= 0 {
		return result
	}

	return getUser(userId, "")
}

func GetUsers() []models.User {

	db := getSqlDB()
	if db == nil {
		return nil
	}

	defer db.Close()

	rows, err := db.Queryx("SELECT [Id], [Nick], [EMail], [Nom], [Prenom], [Adr1], [Ville] FROM [dbo].[Users] WITH(NOLOCK)")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	var result []models.User

	for rows.Next() {
		var row models.User
		//err1 := rows.Scan(&r.Id, &r.Nick, &r.Email, &r.Nom, &r.Prenom, &r.Adr1, &r.Ville)

		err1 := rows.StructScan(&row)

		if err1 != nil {
			fmt.Println("Scan: %v", err)
		}

		result = append(result, row)
	}

	return result
}

func GetUsersSP() []models.User {

	db := getSqlDB()
	if db == nil {
		return nil
	}

	defer db.Close()

	rows, err := db.Queryx("EXEC [dbo].[GetUsers]")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	var result []models.User

	for rows.Next() {
		var row models.User
		err1 := rows.StructScan(&row)

		if err1 != nil {
			fmt.Println("Scan: %v", err)
		}

		result = append(result, row)
	}

	return result
}

func DeleteUser(userId int) bool {

	if userId <= 0 {
		return false
	}

	db := getSqlDB()
	if db == nil {
		return false
	}

	defer db.Close()

	_, err := db.Exec("DELETE FROM [dbo].[Users] WHERE [Id] = ?", userId)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
