package main

import (
	"fmt"
	"net/http"
	"rest/platform"

	"github.com/gin-gonic/gin"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

type Item struct {
	FullName        string `json:"fullName"`
	Cpf             string `json:"cpf"`
	Address         string `json:"address"`
	City            string `json:"city"`
	State           string `json:"state"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Hospital        string `json:"hospital"`
	CardNo          string `json:"cardNo"`
	AppointmentDate string `json:"appointmentDate"`
}

func ExibePost(c *gin.Context) {
	b := Item{}
	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db, _ := sql.Open("sqlite3", "./patient.db")

	feed := patient.FromSQLite(db)

	feed.SetPatient(patient.Item{
		FullName: b.FullName,
		Cpf: b.Cpf,
		Address: b.Address,
		City: b.City,
		State: b.State,
		Phone: b.Phone,
		Email: b.Email,
		Hospital: b.Hospital,
		CardNo: b.CardNo,
		AppointmentDate: b.AppointmentDate,
	})
	
	fmt.Println("success")
	c.JSON(http.StatusAccepted, &b)
}

func main() {
	r := gin.Default()

	r.POST("/patient", ExibePost)

	r.Run(":5000")
}
