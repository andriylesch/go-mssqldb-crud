package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	Id     int    `db:"Id"`
	Nick   string `db:"Nick"`
	Email  string `db:"EMail"`
	Nom    string `db:"Nom"`
	Prenom string `db:"Prenom"`
	Adr1   string `db:"Adr1"`
	Ville  string `db:"Ville"`
}

func (user *User) ToString() string {

	var buffer bytes.Buffer
	buffer.WriteString("------- User --------")
	buffer.WriteString("\n Id : " + strconv.Itoa(user.Id))
	buffer.WriteString("\n Nick : " + user.Nick)
	buffer.WriteString("\n Email : " + user.Email)
	buffer.WriteString("\n Nom : " + user.Nom)
	buffer.WriteString("\n Prenom : " + user.Prenom)
	buffer.WriteString("\n Adr1 : " + user.Adr1)
	buffer.WriteString("\n Ville : " + user.Ville)
	buffer.WriteString("\n---------------------\n")

	return buffer.String()
}

func (user *User) ToJson() string {
	result, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return string(result)
}
