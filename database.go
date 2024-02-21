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

	db.Create(&instances.Instance{Name: "Decidim App Develop", URL: "https://decidim-app-develop.osp.dev/"})
	db.Create(&instances.Instance{Name: "Decidim App k8s", URL: "https://develop.decidim-app.k8s.osp.cat/"})

	db.Create(&instances.Request{InstanceID: 1, Name: "HTTP Requirements", StatusCode: 200, ResponseTime: 1000})
	db.Create(&instances.Request{InstanceID: 2, Name: "HTTP Requirements", StatusCode: 200, ResponseTime: 1000})
	db.Create(&instances.Contains{InstanceID: 1, Name: "Decidim App Develop", Expected: "decidim-app"})
	db.Create(&instances.Contains{InstanceID: 2, Name: "Decidim App k8s", Expected: "decidim-app"})
	db.Create(&instances.Contains{InstanceID: 2, Name: "Decidim App k8s", Expected: "matomo"})
	db.Create(&instances.Requirement{InstanceID: 1, ContainsID: 3, Name: "Decidim App Develop"})
	db.Create(&instances.Requirement{InstanceID: 1, ContainsID: 0, Name: "Has valid HTTP status", RequestID: 1})
	db.Create(&instances.Requirement{InstanceID: 2, ContainsID: 0, Name: "Has valid HTTP status", RequestID: 2})
}
