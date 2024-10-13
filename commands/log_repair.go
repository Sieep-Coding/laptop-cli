// commands/log_repair.go
package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var LogRepairCmd = &cobra.Command{
	Use:   "log-repair",
	Short: "Log a repair for a laptop",
	Run: func(cmd *cobra.Command, args []string) {
		var laptopID int
		var issue, technician string

		fmt.Print("Laptop ID: ")
		fmt.Scanln(&laptopID)

		fmt.Print("Issue: ")
		fmt.Scanln(&issue)

		fmt.Print("Technician: ")
		fmt.Scanln(&technician)

		logRepair(laptopID, issue, technician)
	},
}

// Our logging function for the laptop repairs
// we can add a repair on it's own individdual table.
// LaptopID is used as the foreign key reference back to the main laptop table.
func logRepair(laptopID int, issue, technician string) {
	database := db.InitDB()
	defer database.Close()

	insertRepairSQL := `INSERT INTO repairs (laptop_id, issue, technician, status)
                        VALUES (?, ?, ?, 'pending');`
	statement, err := database.Prepare(insertRepairSQL)
	if err != nil {
		fmt.Println("Error preparing insert statement:", err)
		return
	}
	_, err = statement.Exec(laptopID, issue, technician)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return
	}
	fmt.Println("Repair logged successfully!")
}

// A simple function to remove any repair when given the proper laptopID
// This introduces a new variable called reason for removal
func removeRepair(laptopID int, issue, reasoningForRepair, technician string) {
	database := db.InitDB()
	defer database.Close()

	fmt.Scanf()

	removeRepairSQL := `DELETE INTO repairs (laptop_id, issue, technician, status)
											WHERE laptop_id = ?`
	statement, err := database.Prepare(removeRepairSQL)
	if err != nil {
		fmt.Println("Error preparing for laptop ID (and row) removal", err)
		return
	}
	_, err = statement.Exec(laptopID, issue, technician)
	if err != nil {
		fmt.Println("Error executing laptop removal query", err)
	}
	fmt.Println("Removed repair successfully!")
}
