// main.go
package main

import (
	"fmt"
	"laptop-donation-cli/commands"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "laptop-cli",
		Short: "Laptop Donation Tracker CLI",
	}

	// Registering Commands
	rootCmd.AddCommand(commands.AddLaptopCmd)
	rootCmd.AddCommand(commands.ListLaptopsCmd)
	rootCmd.AddCommand(commands.LogRepairCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
