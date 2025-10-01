package main

import (
	"fmt"
	"net"
)

// resolveTargets returns a slice of IPv4 addresses for the given target.
// If the target is already an IP it returns it directly (as a 1-element slice).
// If it's a hostname it resolves via DNS and returns IPv4 addresses.
// Returns an error if lookup fails or no IPv4 addresses are found.
func resolveTargets(target string) ([]string, error) {
	// If it's already a valid IP, return it
	if ip := net.ParseIP(target); ip != nil {
		// prefer IPv4 only for now
		if ip.To4() != nil {
			return []string{target}, nil
		}
		return nil, fmt.Errorf("only IPv4 supported (got IPv6): %s", target)
	}

	// Otherwise resolve hostname
	ips, err := net.LookupIP(target)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, ip := range ips {
		if ip4 := ip.To4(); ip4 != nil {
			out = append(out, ip4.String())
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no IPv4 addresses found for %s", target)
	}
	return out, nil
}
