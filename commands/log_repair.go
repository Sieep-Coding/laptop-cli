// commands/log_repair.go
package commands

import (
	"fmt"
	"laptop-donation-cli/db"

	"github.com/spf13/cobra"
)

var LogRepairCmd = &cobra.Command{
	Use:   "log-repair",
	Short: "Log a repair for a laptop",
	Run: func(cmd *cobra.Command, args []string) {
		var laptopID int
		var issue, technician string

		fmt.Print("Laptop ID: ")
		fmt.Scanln(&laptopID)

		fmt.Print("Issue: ")
		fmt.Scanln(&issue)

		fmt.Print("Technician: ")
		fmt.Scanln(&technician)

		logRepair(laptopID, issue, technician)
	},
}

func logRepair(laptopID int, issue, technician string) {
	database := db.InitDB()
	defer database.Close()

	insertRepairSQL := `INSERT INTO repairs (laptop_id, issue, technician, status)
                        VALUES (?, ?, ?, 'pending');`
	statement, err := database.Prepare(insertRepairSQL)
	if err != nil {
		fmt.Println("Error preparing insert statement:", err)
		return
	}

	_, err = statement.Exec(laptopID, issue, technician)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return
	}

	fmt.Println("Repair logged successfully!")
}
