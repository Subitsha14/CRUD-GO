package main

import (
	"CRUD-GO/initializers"
	//"CRUD-GO/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

}
