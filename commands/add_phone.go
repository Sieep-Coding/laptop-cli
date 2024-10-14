// commands/add_laptop.go
package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var AddPhoneCmd = &cobra.Command{
	Use:   "add-phone",
	Short: "Add a new phone donation",
	Run: func(cmd *cobra.Command, args []string) {
		var donorName, specs string
		fmt.Print("Donor Name: ")
		fmt.Scanln(&donorName)

		fmt.Print("Phone Specs: ")
		fmt.Scanln(&specs)

		savePhone(donorName, specs)
	},
}

func savePhone(donorName, specs string) {
	database := db.InitDB()
	defer database.Close()

	insertLaptopSQL := `INSERT INTO phones (donor_name, specs, status, donation_date) 
                        VALUES (?, ?, 'donated', DATE('now'));`
	statement, err := database.Prepare(insertLaptopSQL)
	if err != nil {
		fmt.Println("Error preparing insert statement:", err)
		return
	}

	_, err = statement.Exec(donorName, specs)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return
	}

	fmt.Println("Phone added successfully!")
}
