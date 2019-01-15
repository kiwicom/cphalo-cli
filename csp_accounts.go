package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var (
	cmdCspAccounts *cobra.Command
)

func init() {
	cmdCspAccounts = &cobra.Command{
		Use:   "csp-accounts",
		Short: "Manage CSP accounts",
	}

	cmdCspAccounts.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all CSP accounts",
		Run: func(cmd *cobra.Command, args []string) {
			listCSPAccounts()
		},
	})

	cmdCspAccounts.AddCommand(&cobra.Command{
		Use:   "delete [csp-account-id]",
		Short: "Delete a CSP account",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteCSPAccount(args[0])
		},
	})
}

func listCSPAccounts() {
	format := "%s\t%s\t%s\t%s\t\n"

	resp, err := client.ListCSPAccounts()

	if err != nil {
		log.Fatalf("failed to get CSP accounts: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no CSP accounts found")
		return
	}

	log.Printf("found %d CSP accounts", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Type", "Name", "Created at")
	for _, a := range resp.CSPAccounts {
		fmt.Fprintf(w, format, a.ID, a.CSPAccountType, a.AccountDisplayName, a.CreatedAt)
	}
	w.Flush()
}

func deleteCSPAccount(accountID string) {
	if err := client.DeleteCSPAccount(accountID); err != nil {
		fmt.Printf("Aborting. Could not delete CSP account: %v\n", err)
		return
	}

	fmt.Printf("CSP account %s deleted.\n", accountID)
}
