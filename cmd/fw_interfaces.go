package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdFirewallInterfaces *cobra.Command
)

func init() {
	cmdFirewallInterfaces = &cobra.Command{
		Use:   "firewall-interfaces",
		Short: "Manage firewall interfaces",
	}

	cmdFirewallInterfaces.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all firewall interfaces",
		Run: func(cmd *cobra.Command, args []string) {
			listFirewallInterfaces()
		},
	})

	cmdFirewallInterfaces.AddCommand(&cobra.Command{
		Use:   "delete [interface-id]",
		Short: "Delete a firewall interface",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteFirewallInterface(args[0])
		},
	})
}

func listFirewallInterfaces() {
	resp, err := client.ListFirewallInterfaces()

	if err != nil {
		log.Fatalf("failed to get firewall interfaces: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall interfaces found")
		return
	}

	log.Printf("found %d firewall interfaces", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t\n", "ID", "Name", "Permanent")
	for _, p := range resp.Interfaces {
		fmt.Fprintf(w, "%s\t%s\t%t\t\n", p.ID, p.Name, p.System)
	}
	w.Flush()
}

func deleteFirewallInterface(interfaceID string) {
	if err := client.DeleteFirewallInterface(interfaceID); err != nil {
		log.Errorf("Aborting. Could not delete firewall interface: %v\n", err)
		return
	}

	fmt.Printf("Firewall interface %s deleted.\n", interfaceID)
}
