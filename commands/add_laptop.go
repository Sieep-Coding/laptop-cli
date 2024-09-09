// commands/add_laptop.go
package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var AddLaptopCmd = &cobra.Command{
	Use:   "add-laptop",
	Short: "Add a new laptop donation",
	Run: func(cmd *cobra.Command, args []string) {
		var donorName, specs string
		fmt.Print("Donor Name: ")
		fmt.Scanln(&donorName)

		fmt.Print("Laptop Specs: ")
		fmt.Scanln(&specs)

		saveLaptop(donorName, specs)
	},
}

func saveLaptop(donorName, specs string) {
	database := db.InitDB()
	defer database.Close()

	insertLaptopSQL := `INSERT INTO laptops (donor_name, specs, status, donation_date) 
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

	fmt.Println("Laptop added successfully!")
}
