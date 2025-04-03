package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var AddDesktopCmd = &cobra.Command{
	Use:   "add-desktop",
	Short: "Add a new desktop donation",
	Run: func(cmd *cobra.Command, args []string) {
		var donorName, specs string
		fmt.Print("Donor Name: ")
		fmt.Scanln(&donorName)

		fmt.Print("Specs: ")
		fmt.Scanln(&specs)

		saveDesktop(donorName, specs)
	},
}

func saveDesktop(donorName, specs string) {
	database := db.InitDB()
	defer database.Close()

	insertDesktopSQL := `INSERT INTO desktop (donor_name, specs, status, donation_date)
                        VALUES (?, ?, 'donated', DATE('now'));`
	stmt, err := database.Prepare(insertDesktopSQL)
	if err != nil {
		fmt.Println("Error creating a new desktop during db prep:", err)
		return
	}

	_, err = stmt.Exec(donorName, specs)
	if err != nil {
		fmt.Println("Error creating a new desktop during execution:", err)
		return
	}
	fmt.Println("Desktop added successfully!")
}
