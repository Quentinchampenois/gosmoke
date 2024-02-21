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

	// Fetch all instances in database
	var instances []instances.Instance
	db.Preload("Requirements").Preload("Requirements.Contains").Find(&instances)

	for _, instance := range instances {
		fmt.Println(instance.Name)
		for _, requirement := range instance.Requirements {
			fmt.Println("  -", requirement.Name)
			fmt.Println("    -", requirement.Contains.Report())
		}
	}
}
