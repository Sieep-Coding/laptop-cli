package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var DeleteRepairCmd = &cobra.Command{
	Use:   "delete-repair",
	Short: "Delete a repair record",
	Run: func(cmd *cobra.Command, args []string) {
		var repairID int

		fmt.Print("Repair ID to delete: ")
		fmt.Scanln(&repairID)

		deleteRepair(repairID)
	},
}

func deleteRepair(repairID int) {
	database := db.InitDB()
	defer database.Close()

	deleteSQL := `DELETE FROM repairs WHERE id = ?`
	statement, err := database.Prepare(deleteSQL)
	if err != nil {
		fmt.Println("Error preparing delete statement:", err)
		return
	}

	_, err = statement.Exec(repairID)
	if err != nil {
		fmt.Println("Error executing delete statement:", err)
		return
	}

	fmt.Println("Repair deleted successfully!")
}
