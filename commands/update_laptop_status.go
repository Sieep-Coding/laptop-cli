package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var UpdateLaptopStatusCmd = &cobra.Command{
	Use:   "update-laptop-status",
	Short: "Update the status of a laptop (e.g., 'under repair', 'distributed')",
	Run: func(cmd *cobra.Command, args []string) {
		var laptopID int
		var newStatus string

		fmt.Print("Laptop ID: ")
		fmt.Scanln(&laptopID)

		fmt.Print("New Status: ")
		fmt.Scanln(&newStatus)

		updateLaptopStatus(laptopID, newStatus)
	},
}

func updateLaptopStatus(laptopID int, newStatus string) {
	database := db.InitDB()
	defer database.Close()

	updateSQL := `UPDATE laptops SET status = ? WHERE id = ?`
	statement, err := database.Prepare(updateSQL)
	if err != nil {
		fmt.Println("Error preparing update statement:", err)
		return
	}

	_, err = statement.Exec(newStatus, laptopID)
	if err != nil {
		fmt.Println("Error executing update statement:", err)
		return
	}

	fmt.Println("Laptop status updated successfully!")
}
