package main

import (
	"fmt"
	"gosmoke/instances"
)

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	var i []instances.Instance
	db.Preload("Requirements").Preload("Requirements.Contains").Preload("Requirements.Request").Find(&i)

	for _, instance := range i {
		fmt.Println(instance.Name)
		for _, requirement := range instance.Requirements {
			fmt.Println("    -", requirement.Report())
		}
	}
}
