package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var (
	cmdFirewallZones *cobra.Command
)

func init() {
	cmdFirewallZones = &cobra.Command{
		Use:   "firewall-zones",
		Short: "Manage firewall zones",
	}

	cmdFirewallZones.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all firewall zones",
		Run: func(cmd *cobra.Command, args []string) {
			listFirewallZones()
		},
	})

	cmdFirewallZones.AddCommand(&cobra.Command{
		Use:   "delete [zone-id]",
		Short: "Delete a firewall zone",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteFirewallZone(args[0])
		},
	})
}

func listFirewallZones() {
	resp, err := client.ListFirewallZones()

	if err != nil {
		log.Fatalf("failed to get firewall zones: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall zones found")
		return
	}

	log.Printf("found %d firewall zones", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", "ID", "Name", "IP", "Description", "Permanent")
	for _, p := range resp.Zones {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\t\n", p.ID, p.Name, p.IPAddress, p.Description, p.System)
	}
	w.Flush()
}

func deleteFirewallZone(zoneID string) {
	if err := client.DeleteFirewallZone(zoneID); err != nil {
		fmt.Printf("Aborting. Could not delete firewall zone: %v\n", err)
		return
	}

	fmt.Printf("Firewall zone %s deleted.\n", zoneID)
}
