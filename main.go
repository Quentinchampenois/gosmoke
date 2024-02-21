package main

import (
	"fmt"
	"gosmoke/instances"
	"net/http"
)

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	var i []instances.Instance
	db.Preload("Requirements").Preload("Requirements.Contains").Preload("Requirements.Request").Find(&i)

	for _, instance := range i {
		req, err := http.Get(instance.URL)
		if err != nil {
			panic(err)
		}

		fmt.Println("Instance:", instance.Name)
		fmt.Println("  - URL:", instance.URL)

		for _, requirement := range instance.Requirements {
			requirement.Expectation(*req)
			fmt.Println("    -", requirement.Report())
		}
	}
}
