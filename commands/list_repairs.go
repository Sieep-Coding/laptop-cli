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

func listRepairs() {
	database := db.InitDB()
	defer database.Close()

	rows, err := database.Query("SELECT id, laptop_id, issue, technician, status FROM repairs")
	if err != nil {
		fmt.Println("Error querying repairs:", err)
		return
	}
	defer rows.Close()

	fmt.Println("ID | Laptop ID | Issue | Technician | Status")
	for rows.Next() {
		var id, laptopID int
		var issue, technician, status string
		err := rows.Scan(&id, &laptopID, &issue, &technician, &status)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("%d | %d | %s | %s | %s\n", id, laptopID, issue, technician, status)
	}
}
