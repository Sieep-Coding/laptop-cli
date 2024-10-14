// commands/list_laptop.go
package commands

import (
	"fmt"
	"laptop-donation-cli/db"
	"log"

	"github.com/spf13/cobra"
)

var ListLaptopsCmd = &cobra.Command{
	Use:   "list-laptop",
	Short: "List all donated laptops",
	Run: func(cmd *cobra.Command, args []string) {
		listLaptops()
	},
}

func listLaptops() {
	database := db.InitDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, donor_name, specs, status, donation_date FROM laptops")
	if err != nil {
		fmt.Println("Error querying laptops:", err)
		return
	}
	defer rows.Close()

	log.Println("Printing...")
	fmt.Println("ID | Donor | Specs | Status | Donation Date")
	for rows.Next() {
		var id int
		var donorName, specs, status, donationDate string
		rows.Scan(&id, &donorName, &specs, &status, &donationDate)
		fmt.Printf("%d | %s | %s | %s | %s\n", id, donorName, specs, status, donationDate)
	}
}
