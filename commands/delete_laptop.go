package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var DeleteLaptopCmd = &cobra.Command{
	Use:   "delete-laptop",
	Short: "Delete a laptop from the system",
	Run: func(cmd *cobra.Command, args []string) {
		var laptopID int

		fmt.Print("Laptop ID to delete: ")
		fmt.Scanln(&laptopID)

		deleteLaptop(laptopID)
	},
}

func deleteLaptop(laptopID int) {
	database := db.InitDB()
	defer database.Close()

	deleteSQL := `DELETE FROM laptops WHERE id = ?`
	statement, err := database.Prepare(deleteSQL)
	if err != nil {
		fmt.Println("Error preparing delete statement:", err)
		return
	}

	_, err = statement.Exec(laptopID)
	if err != nil {
		fmt.Println("Error executing delete statement:", err)
		return
	}

	fmt.Println("Laptop deleted successfully!")
}
