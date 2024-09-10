# Laptop Donation Tracker CLI

Laptop Donation Tracker is a Command Line Interface (CLI) application built with Go and SQLite. It helps nonprofits manage laptop donations, track repairs, and distribute devices to recipients.

## Features

- **Add a laptop donation**: Track donated laptops with donor name and specs.
- **List donated laptops**: View a list of all donated laptops, including their status.
- **Log repairs**: Log repairs performed on a laptop and track the status of repairs.
- **Update laptop status**: Update the status of a laptop (e.g., "under repair," "distributed").
- **List repairs**: View all repairs logged for laptops.
- **Add a recipient**: Assign a recipient to a donated laptop.
- **List recipients**: Display a list of all recipients and the laptops they received.
- **Delete a laptop**: Remove a laptop from the system.
- **Delete a repair**: Remove a repair record.

## Requirements

- Go 1.16 or later
- SQLite 3


# HELP

```
Usage:
  laptop-cli [command]

Available Commands:
  add-laptop           Add a new laptop donation
  add-recipient        Add a recipient for a donated laptop
  completion           Generate the autocompletion script for the specified shell
  delete-laptop        Delete a laptop from the system
  delete-repair        Delete a repair record
  help                 Help about any command
  list-laptop          List all donated laptops
  list-recipients      List all recipients and the laptops they received
  list-repairs         List all repairs
  log-repair           Log a repair for a laptop
  update-laptop-status Update the status of a laptop (e.g., 'under repair', 'distributed')

Flags:
  -h, --help   help for laptop-cli

Use "laptop-cli [command] --help" for more information about a command.

```
