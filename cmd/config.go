package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmdConfig *cobra.Command
)

const configTemplate = "application_key: %s\napplication_secret: %s\n"

func init() {
	cmdConfig = &cobra.Command{
		Use:   "config",
		Short: "Prepare configuration file",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// overwrite parent func
		},
		Run: func(cmd *cobra.Command, args []string) {
			configPath := cfgFile

			// if config path is not set, set it in $HOME
			if configPath == "" {
				home, err := homedir.Dir()
				if err != nil {
					log.Fatal(err)
				}
				configPath = fmt.Sprintf("%s/%s.yaml", home, configName)
			}

			if _, err := os.Stat(configPath); os.IsExist(err) {
				fmt.Printf("Config already exists at %s\n", configPath)
				fmt.Println("If you would like to update it, delete existing file first.")

				return
			}

			key, secret, err := readCredentials(os.Stdin)

			if err != nil {
				log.Fatal(err)
			}

			content := fmt.Sprintf(configTemplate, key, secret)

			if err = ioutil.WriteFile(configPath, []byte(content), 0600); err != nil {
				log.Fatalf("failed to write to %s: %v", configPath, err)
			}

			fmt.Printf("config file successfully written to %s\n", configPath)
		},
	}
}

func readCredentials(r io.Reader) (string, string, error) {
	scanner := bufio.NewScanner(r)

	key := viper.GetString("application_key")
	secret := viper.GetString("application_secret")

	if key == "" {
		fmt.Print("application key: ")
		scanner.Scan()

		key = scanner.Text()
	}

	if key == "" {
		return "", "", fmt.Errorf("key not given")
	}

	if secret == "" {
		fmt.Print("application secret: ")
		scanner.Scan()

		secret = scanner.Text()
	}

	if secret == "" {
		return "", "", fmt.Errorf("secret not given")
	}

	return key, secret, nil
}
