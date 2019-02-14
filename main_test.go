package main

import "os"

func ExampleHelp() {
	os.Setenv("CP_APPLICATION_KEY", "test")
	os.Setenv("CP_APPLICATION_SECRET", "test")
	main()

	//Output:Usage:
	//   cphalo [command]
	//
	//Available Commands:
	//   alert-profiles      Manage alert profiles
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
	//   -h, --help      help for cphalo
	//       --verbose   Output verbose information
	//
	//Use "cphalo [command] --help" for more information about a command.

}
