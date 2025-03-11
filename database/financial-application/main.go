package main

import (
	"errors"
	"fmt"
	"go-modules/database/financial-application/pkg"

	"log"
)

func main() {
	db, err := pkg.DbCon()
	if err != nil {
		log.Fatal("Failed to connect to db ", err)
	}
	defer db.Close()
	fmt.Println("Connected to db")
	user_id := 1
	var newBalance float64
	newBalance = 1000
	err = pkg.UpdateBalance(db, user_id, newBalance)
	if err != nil {
		if errors.Is(err, pkg.ErrVersionConflict) {
			log.Println("Version conflict")
		} else {
			log.Fatalf("Failed to update balance %v", err)
		}
	} else {
		fmt.Println("Balance updated successfully")
	}

	var userCount int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User count: ", userCount)
}
