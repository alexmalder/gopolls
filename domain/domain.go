package domain

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func FindByCn(cn string) (*ldap.Entry, error) {
	conn, err := ldap.DialURL("ldap://127.0.0.1:1389")
	if err != nil {
		log.Fatalf("Failed to connect: %s\n", err)
	}
	defer conn.Close()

	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=%v,dc=example,dc=ru", cn), // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"cn"},                          // A list attributes to retrieve
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(sr.Entries) != 1 {
		return nil, fmt.Errorf("Couldn't not find or too many entries")
	}
	return sr.Entries[0], nil
}

func ExampleConn_WhoAmI() {
	conn, err := ldap.DialURL("ldap://127.0.0.1:1389")
	if err != nil {
		log.Fatalf("Failed to connect: %s\n", err)
	}
	defer conn.Close()

	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=ru", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"cn", "givenName", "sn", "uidNumber", "description", "objectClass"}, // A list attributes to retrieve
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		for _, attr := range entry.Attributes {
			fmt.Printf("%s | %v | %v\n", entry.DN, attr.Name, attr.Values)
		}
	}
}
