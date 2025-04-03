package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var UpdateDesktopStatusCmd = &cobra.Command{
	Use:   "update-desktop-status",
	Short: "Update the status of a desktop (e.g., 'under repair', 'distributed')",
	Run: func(cmd *cobra.Command, args []string) {
		var desktopID int
		var newStatus string

		fmt.Print("Desktop ID: ")
		fmt.Scanln(&desktopID)

		fmt.Print("New Status: ")
		fmt.Scanln(&newStatus)

		updateDesktopStatus(desktopID, newStatus)
	},
}

func updateDesktopStatus(desktopID int, newStatus string) {
	database := db.InitDB()
	defer database.Close()

	updateSQL := `UPDATE desktops SET status = ? WHERE id = ?`
	statement, err := database.Prepare(updateSQL)
	if err != nil {
		fmt.Println("Error preparing update statement:", err)
		return
	}
	_, err = statement.Exec(newStatus, desktopID)
	if err != nil {
		fmt.Println("Error executing update statement:", err)
		return
	}
	fmt.Println("Desktop status updated successfully!")
}
