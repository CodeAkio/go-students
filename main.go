package main

import (
	"github.com/CodeAkio/go-students/models"
	"github.com/CodeAkio/go-students/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Pedro", CPF: "000.000.000-00", RG: "00.000.000-0"},
		{Name: "Ana", CPF: "111.111.111-11", RG: "11.111.111-1"},
		{Name: "Julia", CPF: "222.222.222-22", RG: "22.222.222-2"},
	}

	routes.HandleRequests()
}
