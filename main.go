// main.go
package main

import (
	"fmt"
	"laptop-donation-cli/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "laptop-cli",
		Short: "Laptop Donation Tracker CLI",
	}

	rootCmd.AddCommand(commands.ListLaptopsCmd)
	rootCmd.AddCommand(commands.LogRepairCmd)
	rootCmd.AddCommand(commands.UpdateLaptopStatusCmd)
	rootCmd.AddCommand(commands.ListRepairsCmd)
	rootCmd.AddCommand(commands.AddRecipientCmd)
	rootCmd.AddCommand(commands.ListRecipientsCmd)
	rootCmd.AddCommand(commands.DeleteLaptopCmd)
	rootCmd.AddCommand(commands.DeleteRepairCmd)
	// rootCmd.AddCommand(commands.removeRepairCmd)
	// rootCmd.AddCommand(commands.AddPhoneCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
