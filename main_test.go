package main

import "os"

func Example_main() {
	os.Setenv("CPHALO_APPLICATION_KEY", "test")
	os.Setenv("CPHALO_APPLICATION_SECRET", "test")
	main()

	//Output: CLI tool for connection to CPHalo API
	//
	// Usage:
	//   cphalo [command]
	//
	//Available Commands:
	//   alert-profiles      Manage alert profiles
	//   config              Prepare configuration file
	//   csp-accounts        Manage CSP accounts
	//   firewall-interfaces Manage firewall interfaces
	//   firewall-policies   Manage firewall policies
	//   firewall-rules      Manage firewall rules for a specific policy
	//   firewall-services   Manage firewall services
	//   firewall-zones      Manage firewall zones
	//   help                Help about any command
	//   server-groups       Manage server groups
	//   servers             Manage servers
	//
	//Flags:
	//   -h, --help            help for cphalo
	//       --key string      Application key
	//       --path string     config file (default is $HOME/.cphalo.yaml)
	//       --secret string   Application secret
	//
	//Use "cphalo [command] --help" for more information about a command.
}
