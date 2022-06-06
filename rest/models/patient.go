package models

type Patient struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  FullName  string `json:"fullName"`
  CPF string `json:"cpf"`
}