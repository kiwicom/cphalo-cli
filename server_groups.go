package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	cmdServerGroups *cobra.Command
)

func init() {
	cmdServerGroups = &cobra.Command{
		Use:   "server-groups",
		Short: "Manage server groups",
	}

	cmdServerGroups.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all server groups",
		Run: func(cmd *cobra.Command, args []string) {
			listServerGroups()
		},
	})

	cmdServerGroups.AddCommand(&cobra.Command{
		Use:   "delete [group-id]",
		Short: "Delete a server group",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteServerGroup(args[0])
		},
	})
}

func listServerGroups() {
	format := "%s\t%s\t%s\t\n"

	resp, err := client.ListServerGroups()

	if err != nil {
		log.Fatalf("failed to get server groups: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no server groups found")
		return
	}

	log.Printf("found %d server groups", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Name", "ParentID")
	for _, sg := range resp.Groups {
		fmt.Fprintf(w, format, sg.ID, sg.Name, sg.ParentID)
	}
	w.Flush()
}

func deleteServerGroup(groupID string) {
	if err := client.DeleteFirewallZone(groupID); err != nil {
		fmt.Printf("Aborting. Could not delete server group: %v\n", err)
		return
	}

	fmt.Printf("Server group %s deleted.\n", groupID)
}
