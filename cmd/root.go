package cmd

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/kiwicom/cphalo-go"
)

var (
	cfgFile string
	client  *cphalo.Client

	configName = ".cphalo"
)

var rootCmd = &cobra.Command{
	Use:   "cphalo",
	Short: "CLI tool for connection to CPHalo API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		key := viper.GetString("application_key")
		secret := viper.GetString("application_secret")

		if key == "" {
			log.Fatal("application_key is not set")
		}
		if secret == "" {
			log.Fatal("application_secret is not set")
		}

		client = cphalo.NewClient(key, secret, nil)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		cmdConfig,
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
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cphalo.yaml)")
	rootCmd.PersistentFlags().String("key", "", "Application key")
	rootCmd.PersistentFlags().String("secret", "", "Application secret")

	if err := viper.BindPFlag("application_key", rootCmd.PersistentFlags().Lookup("key")); err != nil {
		log.Fatal(err)
	}
	if err := viper.BindPFlag("application_secret", rootCmd.PersistentFlags().Lookup("secret")); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		viper.SetConfigName(configName)
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}
	viper.SetEnvPrefix("CPHALO")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("using config file: %s", viper.ConfigFileUsed())
	}
}
