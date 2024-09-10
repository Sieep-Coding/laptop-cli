package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var AddRecipientCmd = &cobra.Command{
	Use:   "add-recipient",
	Short: "Add a recipient for a donated laptop",
	Run: func(cmd *cobra.Command, args []string) {
		var recipientName, contactInfo string
		var laptopID int

		fmt.Print("Recipient Name: ")
		fmt.Scanln(&recipientName)

		fmt.Print("Contact Info: ")
		fmt.Scanln(&contactInfo)

		fmt.Print("Laptop ID: ")
		fmt.Scanln(&laptopID)

		addRecipient(recipientName, contactInfo, laptopID)
	},
}

func addRecipient(name, contactInfo string, laptopID int) {
	database := db.InitDB()
	defer database.Close()

	insertRecipientSQL := `INSERT INTO recipients (name, contact_info, laptop_id, received_at)
                           VALUES (?, ?, ?, DATE('now'))`
	statement, err := database.Prepare(insertRecipientSQL)
	if err != nil {
		fmt.Println("Error preparing insert statement:", err)
		return
	}

	_, err = statement.Exec(name, contactInfo, laptopID)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return
	}

	fmt.Println("Recipient added successfully!")
}
