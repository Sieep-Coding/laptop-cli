package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var ListRepairsCmd = &cobra.Command{
	Use:   "list-repairs",
	Short: "List all repairs",
	Run: func(cmd *cobra.Command, args []string) {
		listRepairs()
	},
}

// List all repairs that are currently active in the system.
// We use a little ASCII in order to display in the terminal
// We could also CRUD this into a web app or something if we ever wanted
func listRepairs() {
	database := db.InitDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, laptop_id, issue, technician, status, reasoningForRepair FROM repairs")
	if err != nil {
		fmt.Println("Error querying repairs:", err)
		return
	}
	defer rows.Close()

	fmt.Println("ID | Laptop ID | Issue | Technician | Status | Reason For Repair")
	for rows.Next() {
		var id, laptopID int
		var issue, technician, status, reasoningForRepair string
		err := rows.Scan(&id, &laptopID, &issue, &technician, &status, &reasoningForRepair)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("%d | %d | %s | %s | %s | %s\n", id, laptopID, issue, technician, status, reasoningForRepair)
	}
}
