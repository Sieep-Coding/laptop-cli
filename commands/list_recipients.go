package commands

import (
	"fmt"
	"laptop-donation-cli/db"
	"log"

	"github.com/spf13/cobra"
)

var ListRecipientsCmd = &cobra.Command{
	Use:   "list-recipients",
	Short: "List all recipients and the laptops they received",
	Run: func(cmd *cobra.Command, args []string) {
		listRecipients()
	},
}

func listRecipients() {
	database := db.InitDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, name, contact_info, laptop_id, received_at FROM recipients")
	if err != nil {
		fmt.Println("Error querying recipients:", err)
		return
	}
	defer rows.Close()

	log.Printf("Printing...")
	fmt.Println("ID | Name | Contact Info | Laptop ID | Received At")
	for rows.Next() {
		var id, laptopID int
		var name, contactInfo, receivedAt string
		err := rows.Scan(&id, &name, &contactInfo, &laptopID, &receivedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("%d | %s | %s | %d | %s\n", id, name, contactInfo, laptopID, receivedAt)
	}
}
