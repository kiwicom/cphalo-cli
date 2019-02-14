package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/kiwicom/cphalo-go"
	"io/ioutil"
	"log"
	"os"
)

var (
	client  *cphalo.Client
	cmdName = "cphalocli"

	cpAppKey    string
	cpAppSecret string

	rootCmd *cobra.Command
)

func main() {
	if cpAppKey = os.Getenv("CP_APPLICATION_KEY"); cpAppKey == "" {
		fmt.Println("Environment variable CP_APPLICATION_KEY must be set.")
		return
	}

	if cpAppSecret = os.Getenv("CP_APPLICATION_SECRET"); cpAppSecret == "" {
		fmt.Println("Environment variable CP_APPLICATION_SECRET must be set.")
		return
	}

	client = cphalo.NewClient(cpAppKey, cpAppSecret)

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

	rootCmd.Execute()
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
