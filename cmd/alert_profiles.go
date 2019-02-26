package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdAlertProfiles *cobra.Command
)

func init() {
	cmdAlertProfiles = &cobra.Command{
		Use:   "alert-profiles",
		Short: "Manage alert profiles",
	}

	cmdAlertProfiles.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all alert profiles",
		Run: func(cmd *cobra.Command, args []string) {
			listAlertProfiles()
		},
	})
}

func listAlertProfiles() {
	format := "%s\t%s\t\n"
	resp, err := client.ListAlertProfiles()

	if err != nil {
		log.Fatalf("failed to get alert profiles: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no alert profiles found")
		return
	}

	log.Printf("found %d alert profiles", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Name")
	for _, p := range resp.AlertProfiles {
		fmt.Fprintf(w, format, p.ID, p.Name)
	}
	w.Flush()
}
