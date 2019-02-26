package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	cmdFirewallPolicies *cobra.Command
)

func init() {
	cmdFirewallPolicies = &cobra.Command{
		Use:   "firewall-policies",
		Short: "Manage firewall policies",
	}

	cmdFirewallPolicies.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all firewall policies",
		Run: func(cmd *cobra.Command, args []string) {
			listFirewallPolicies()
		},
	})

	cmdFirewallPolicies.AddCommand(&cobra.Command{
		Use:   "delete [policy-id]",
		Short: "Delete a firewall policy",
		Long:  "Delete a firewall policy, including all attached firewall rules",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteFirewallPolicy(args[0])
		},
	})
}

func listFirewallPolicies() {
	format := "%s\t%s\t%s\t\n"
	resp, err := client.ListFirewallPolicies()

	if err != nil {
		log.Fatalf("failed to get firewall policies: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall policies found")
		return
	}

	log.Printf("found %d firewall policies", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Name", "Platform")
	for _, p := range resp.Policies {
		fmt.Fprintf(w, format, p.ID, p.Name, p.Platform)
	}
	w.Flush()
}

func deleteFirewallPolicy(policyID string) {
	resp, err := client.ListFirewallRules(policyID)

	if err != nil {
		log.Fatalf("failed to get firewall rules: %v", err)
	}

	// take care of attached firewall rules
	if resp.Count > 0 {
		fmt.Print("Firewall policy has attached rules, please type YES to confirm deletion: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if strings.ToLower(scanner.Text()) != "yes" {
			fmt.Println("Deletion cancelled.")
			return
		}

		for _, rule := range resp.Rules {
			if err = client.DeleteFirewallRule(policyID, rule.ID); err != nil {
				fmt.Printf("Aborting. Could not delete firewall rule: %v\n", err)
				return
			}

			fmt.Printf("Rule %s deleted.\n", rule.ID)
		}
	}

	if err = client.DeleteFirewallPolicy(policyID); err != nil {
		fmt.Printf("Could not delete firewall policy: %v\n", err)
	} else {
		fmt.Printf("Policy %s deleted.\n", policyID)
	}
}
