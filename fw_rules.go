package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var (
	cmdFirewallRules *cobra.Command
)

func init() {
	cmdFirewallRules = &cobra.Command{
		Use:   "firewall-rules",
		Short: "Manage firewall rules for a specific policy",
	}

	cmdFirewallRules.AddCommand(&cobra.Command{
		Use:   "list [policy-id]",
		Short: "List all firewall rules for a specific policy",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			listFirewallRules(args[0])
		},
	})

	cmdFirewallRules.AddCommand(&cobra.Command{
		Use:   "delete [policy-id] [rule-id]",
		Short: "Delete a firewall rule for a specific policy",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deleteFirewallRule(args[0], args[1])
		},
	})

	cmdFirewallRules.AddCommand(&cobra.Command{
		Use:   "delete-all [policy-id]",
		Short: "Delete all firewall rule for a specific policy",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteAllFirewallRules(args[0])
		},
	})
}

func listFirewallRules(policyID string) {
	format := "%s\t%s\t%s\t%t\t\n"
	resp, err := client.ListFirewallRules(policyID)

	if err != nil {
		log.Fatalf("failed to get firewall rules: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall rules found")
		return
	}

	log.Printf("found %d firewall rules", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", "policyID", "Chain", "Action", "Active")
	for _, r := range resp.Rules {
		fmt.Fprintf(w, format, r.ID, r.Chain, r.Action, r.Active)
	}
	w.Flush()
}

func deleteFirewallRule(policyID, ruleID string) error {
	if err := client.DeleteFirewallRule(policyID, ruleID); err != nil {
		return fmt.Errorf("could not delete firewall rule: %v\n", err)
	}

	fmt.Printf("Rule %s deleted.\n", ruleID)

	return nil
}

func deleteAllFirewallRules(policyID string) {
	resp, err := client.ListFirewallRules(policyID)

	if err != nil {
		log.Fatalf("failed to get firewall rules: %v", err)
	}

	for _, rule := range resp.Rules {
		if err = deleteFirewallRule(policyID, rule.ID); err != nil {
			fmt.Printf("Aborting: %v\n", err)
			return
		}

		fmt.Printf("Rule %s deleted.\n", rule.ID)
	}
}
