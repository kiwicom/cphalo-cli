package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var (
	cmdFirewallServices *cobra.Command
)

func init() {
	cmdFirewallServices = &cobra.Command{
		Use:   "firewall-services",
		Short: "Manage firewall services",
	}

	cmdFirewallServices.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all firewall services",
		Run: func(cmd *cobra.Command, args []string) {
			listFirewallServices()
		},
	})

	cmdFirewallServices.AddCommand(&cobra.Command{
		Use:   "delete [service-id]",
		Short: "Delete a firewall service",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteFirewallService(args[0])
		},
	})
}

func listFirewallServices() {
	resp, err := client.ListFirewallServices()

	if err != nil {
		log.Fatalf("failed to get firewall services: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall services found")
		return
	}

	log.Printf("found %d firewall services", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", "ID", "Name", "Port", "Protocol", "Permanent")
	for _, p := range resp.Services {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\t\n", p.ID, p.Name, p.Port, p.Protocol, p.System)
	}
	w.Flush()
}

func deleteFirewallService(serviceID string) {
	if err := client.DeleteFirewallService(serviceID); err != nil {
		fmt.Printf("Aborting. Could not delete firewall service: %v\n", err)
		return
	}

	fmt.Printf("Firewall service %s deleted.\n", serviceID)
}
