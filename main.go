package main

import (
	"fmt"

	"github.com/jsrdriguez/go-hands-on/config"
)

func main() {
	database := config.InitDB()

	defer database.Close()

	fmt.Println(database)

}
