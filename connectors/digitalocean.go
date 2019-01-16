package connectors

import (
	"net"

	"gopkg.in/alecthomas/kingpin.v2"
)

// DigitalOceanDNSConnector implements the connector interface and provides methods
// to interact with the DigitalOcean DNS API
type DigitalOceanDNSConnector struct {
}

// Register this connector in the list
func init() {
	RegisterConnector("digitalocean", &DigitalOceanDNSConnector{})
}

// Assert that we meet the interface at compile time
var _ Connector = (*DigitalOceanDNSConnector)(nil)

// SetupCommand is used to setup the command line details
func (do *DigitalOceanDNSConnector) SetupCommand(app *kingpin.CmdClause) {

}

// AddNode creates a new record for that node on the DO DNS API
func (do *DigitalOceanDNSConnector) AddNode(address string, ip net.IP, ttl int) error {
	return nil
}

// UpdateNode updates a record for that node on the DO DNS API
func (do *DigitalOceanDNSConnector) UpdateNode(address string, ip net.IP, ttl int) error {
	return nil
}

// DeleteNode deletes the record for that node on the DO DNS API
func (do *DigitalOceanDNSConnector) DeleteNode(address string) error {
	return nil
}
