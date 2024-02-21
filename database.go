package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gosmoke/instances"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gosmoke.sqlite3"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automigrate your model
	err = db.AutoMigrate(&instances.Instance{})
	err = db.AutoMigrate(&instances.Requirement{})
	err = db.AutoMigrate(&instances.Contains{})
	err = db.AutoMigrate(&instances.Request{})
	if err != nil {
		return nil, err
	} // Assuming Instance is your GORM model struct

	seeding(db)
	return db, nil
}

func seeding(db *gorm.DB) {
	if db.Find(&instances.Instance{}).RowsAffected > 0 {
		return
	}

	db.Create(&instances.Instance{Name: "Google", URL: "https://google.fr/"})
	db.Create(&instances.Request{InstanceID: 1, Name: "HTTP Requirements", StatusCode: 200})
	db.Create(&instances.Contains{InstanceID: 1, Name: "Has a search button", Expected: "Recherche Google"})
	db.Create(&instances.Requirement{InstanceID: 1, ContainsID: 1, Name: ""})
	db.Create(&instances.Requirement{InstanceID: 1, RequestID: 1, Name: ""})

	db.Create(&instances.Instance{Name: "Mozilla FR", URL: "https://www.mozilla.org/fr/"})
	db.Create(&instances.Request{InstanceID: 2, Name: "HTTP Requirements", StatusCode: 200})
	db.Create(&instances.Contains{InstanceID: 2, Name: "Has a CEO's message", Expected: "La bonne santé d’Internet et de la vie en ligne est notre raison d’exister."})
	db.Create(&instances.Requirement{InstanceID: 2, ContainsID: 2, Name: ""})
	db.Create(&instances.Requirement{InstanceID: 2, RequestID: 2, Name: ""})

}
