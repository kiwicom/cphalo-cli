package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/kiwicom/cphalo-go"
)

var (
	client  *cphalo.Client
	cmdName = "cphalo"

	cpAppKey    string
	cpAppSecret string

	rootCmd *cobra.Command
)

func main() {
	if cpAppKey = os.Getenv("CPHALO_APPLICATION_KEY"); cpAppKey == "" {
		fmt.Println("Environment variable CPHALO_APPLICATION_KEY must be set.")
		return
	}

	if cpAppSecret = os.Getenv("CPHALO_APPLICATION_SECRET"); cpAppSecret == "" {
		fmt.Println("Environment variable CPHALO_APPLICATION_SECRET must be set.")
		return
	}

	client = cphalo.NewClient(cpAppKey, cpAppSecret, nil)

	cobra.OnInitialize(setupLogging)

	rootCmd = &cobra.Command{Use: cmdName}
	rootCmd.PersistentFlags().Bool("verbose", false, "Output verbose information")
	rootCmd.AddCommand(
		cmdServerGroups,
		cmdServers,
		cmdCspAccounts,
		cmdAlertProfiles,
		cmdFirewallPolicies,
		cmdFirewallRules,
		cmdFirewallInterfaces,
		cmdFirewallServices,
		cmdFirewallZones,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute root command: %v", err)
	}
}

func setupLogging() {
	verbose, err := rootCmd.Flags().GetBool("verbose")
	if err != nil {
		fmt.Printf("failed parsing verbosity flag: %v\n", err)
		verbose = false
	}

	if !verbose {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
}
