package commands

import (
	"fmt"
	"laptop-donation-cli/db"
	"log"

	"github.com/spf13/cobra"
)

var ListDesktopsCmd = &cobra.Command{
	Use:   "list-desktop",
	Short: "List available desktops in inventory",
	Run: func(cmd *cobra.Command, args []string) {
		listDesktops()
	},
}

func listDesktops() {
	database := db.InitDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, donor_name, specs, status, donation_date FROM desktops")
	if err != nil {
		fmt.Println("Error querying desktop table:", err)
		return
	}
	defer rows.Close()

	log.Println("Printing...")
	fmt.Println("ID | Donor | Specs | Status | Donation Date")

	for rows.Next() {
		var id int
		var donorName, specs, status, donation_date string
		rows.Scan(&id, &donorName, &specs, &status, &donation_date)
		fmt.Printf("%d | %s | %s | %s | %s\n", id, donorName, specs, status, donation_date)
	}
}
