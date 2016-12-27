# go-mssqldb-crud

How to make friends GO and MSSQL

## Install

    go get github.com/andriylesch/go-mssqldb-crud


In the project will be described subjects like :

- 1. Go with SQL Server driver is unable to connect successfully
- 2. Description about CRUD logic with pure SQL

## Go with SQL Server driver is unable to connect successfully

There are a lot of information about how your GO app can do requests to MySQL/ NoSql / SqlLite / ... databases.
In this project it will be shared information what should be done before start writting your GO app.

`NOTE : during installation MSSQL server you need to choose Authentication Mode as "Mixed Mode (Sql Server authentication and Windows)".
this small remark will save your time in your next step in configuration.`

On stackoverflow site you can find page http://stackoverflow.com/questions/32010749/go-with-sql-server-driver-is-unable-to-connect-successfully-login-fail (really thanks "Richard Chambers" with his answer). On this page you can do step by step your changeds in "Sql Server Configuration Manager".

## Description about CRUD logic with pure SQL

in solution there is sql file with schemas and data (file name "sqlLocalDB.sql")

Test application is using connection string :

```go
driver={SQL Server};server=.\\GO;database=LocalDb;user id=#####;password=####;
```

Methods :

- GetUsers
- GetUsersSP (`using stored procedure`) 
- GetUserById
- GetUserByNick
- UpdateUser
- InsertUser
- DeleteUser


For all examples was created custom `User` type

```go

type User struct {
	Id     int    `db:"Id"`
	Nick   string `db:"Nick"`
	Email  string `db:"EMail"`
	Nom    string `db:"Nom"`
	Prenom string `db:"Prenom"`
	Adr1   string `db:"Adr1"`
	Ville  string `db:"Ville"`
}

```

Code of method `GetUsers`:

```go

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
		err1 := rows.StructScan(&row)

		if err1 != nil {
			fmt.Println("Scan: %v", err)
		}

		result = append(result, row)
	}

	return result
}

``` 

And execute method :

```go

func main() {
	fmt.Println("------ GetUsers ------------")
	resultDb := managers.GetUsers()
  	fmt.Println(resultDb[0].ToString())
}

```
