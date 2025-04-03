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

	//laptop
	rootCmd.AddCommand(commands.ListLaptopsCmd)
	rootCmd.AddCommand(commands.DeleteLaptopCmd)
	rootCmd.AddCommand(commands.UpdateLaptopStatusCmd)

	//repairs
	rootCmd.AddCommand(commands.DeleteRepairCmd)
	rootCmd.AddCommand(commands.LogRepairCmd)
	rootCmd.AddCommand(commands.ListRepairsCmd)

	//recipients
	rootCmd.AddCommand(commands.AddRecipientCmd)
	rootCmd.AddCommand(commands.ListRecipientsCmd)

	//desktop
	rootCmd.AddCommand(commands.AddDesktopCmd)
	rootCmd.AddCommand(commands.ListDesktopsCmd)
	rootCmd.AddCommand(commands.UpdateDesktopStatusCmd)

	// rootCmd.AddCommand(commands.removeRepairCmd)
	// rootCmd.AddCommand(commands.AddPhoneCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
