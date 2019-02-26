package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdServers *cobra.Command
)

func init() {
	cmdServers = &cobra.Command{
		Use:   "servers",
		Short: "Manage servers",
	}

	cmdServers.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all servers",
		Run: func(cmd *cobra.Command, args []string) {
			listServers()
		},
	})

	cmdServers.AddCommand(&cobra.Command{
		Use:   "delete [server-id]",
		Short: "Delete a server",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteServer(args[0])
		},
	})
}

func listServers() {
	format := "%s\t%s\t%s\t\n"
	resp, err := client.ListServers()

	if err != nil {
		log.Fatalf("failed to get servers: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no servers found")
		return
	}

	log.Printf("found %d servers", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Group Name", "Platform")
	for _, s := range resp.Servers {
		fmt.Fprintf(w, format, s.ID, s.GroupName, s.Platform)
	}
	w.Flush()
}

func deleteServer(serverID string) {
	if err := client.DeleteServer(serverID); err != nil {
		log.Errorf("Aborting. Could not delete server: %v\n", err)
		return
	}

	fmt.Printf("Server %s deleted.\n", serverID)
}
