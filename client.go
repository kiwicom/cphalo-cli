package main

import (
	"fmt"
	"gitlab.skypicker.com/terraform-provider-cphalo/api"
	"log"
	"os"
	"text/tabwriter"
)

var (
	client *api.Client
)

func main() {
	client = api.NewClient(os.Getenv("CP_APPLICATION_KEY"), os.Getenv("CP_APPLICATION_SECRET"))

	if len(os.Args) < 2 {
		help()
		return
	}

	endpoint := os.Args[1]

	// disable log output
	//log.SetFlags(0)
	//log.SetOutput(ioutil.Discard)

	switch endpoint {
	case "server_groups":
		listServerGroups()
	case "servers":
		listServers()
	case "firewall_policies":
		listFirewallPolicies()
	case "csp_accounts":
		listCSPAccounts()
	case "alert_profiles":
		listAlertProfiles()
	case "firewall_rules":
		if len(os.Args) < 3 {
			help("please provide ID of firewall policy")
			return
		}
		listFirewallRules(os.Args[2])
	default:
		help()
	}
}

func help(extra ...string) {
	fmt.Printf("Usage:\n\t%s endpoint\n", os.Args[0])

	if len(extra) > 0 {
		fmt.Println()
	}

	for _, e := range extra {
		fmt.Println(e)
	}
}

func listCSPAccounts() {
	format := "%s\t%s\t%s\t%s\t\n"

	resp, err := client.ListCSPAccounts()

	if err != nil {
		log.Fatalf("failed to get CSP accounts: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no CSP accounts found")
		return
	}

	log.Printf("found %d CSP accounts", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Type", "Name", "Created at")
	for _, a := range resp.CSPAccounts {
		fmt.Fprintf(w, format, a.ID, a.CSPAccountType, a.AccountDisplayName, a.CreatedAt)
	}
	w.Flush()
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

func listFirewallPolicies() {
	format := "%s\t%s\t%s\t\n"
	resp, err := client.ListFirewallPolicies()

	if err != nil {
		log.Fatalf("failed to get firewall policies: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall policies found")
		return
	}

	log.Printf("found %d firewall policies", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, format, "ID", "Name", "Platform")
	for _, p := range resp.Policies {
		fmt.Fprintf(w, format, p.ID, p.Name, p.Platform)
	}
	w.Flush()
}

func listFirewallRules(ID string) {
	format := "%s\t%s\t%s\t%t\t\n"
	resp, err := client.ListFirewallRules(ID)

	if err != nil {
		log.Fatalf("failed to get firewall rules: %v", err)
	}

	if resp.Count == 0 {
		fmt.Println("no firewall rules found")
		return
	}

	log.Printf("found %d firewall rules", resp.Count)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", "ID", "Chain", "Action", "Active")
	for _, r := range resp.Rules {
		fmt.Fprintf(w, format, r.ID, r.Chain, r.Action, r.Active)
	}
	w.Flush()
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
