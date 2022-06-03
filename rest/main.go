package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

type Body struct {
	// json tag to de-serialize json body
	 Name string `json:"nome"`
  }

type Patient struct {
     // json tag to de-serialize json body
	 FullName string `json:"fullName"`
	 Cpf string `json:"cpf"`
	 Address string `json:"address"`
	 City string `json:"city"`
	 State string `json:"state"`
	 Phone string `json:"phone"`
	 Email string `json:"email"`
	 Hospital string `json:"hospital"`
	 CardNo string `json:"cardNo"`
	 AppointmentDate string `json:"appointmentDate"`

  }

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{
		"id": "1",
		"nome": "Maicon",
	})
}

func ExibePost(c *gin.Context){
	b := Patient{}
	if err:=c.BindJSON(&b);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	 }
	 fmt.Println(b)
	 c.JSON(http.StatusAccepted,&b) 
}

func main(){
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.POST("/teste", ExibePost)
	r.POST("/patient", ExibePost)
	models.ConnectDatabase()
	r.Run(":5000")
}