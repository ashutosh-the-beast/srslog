package srslog

import (
	"fmt"
	"net"
)

// netConn has an internal net.Conn and adheres to the serverConn interface,
// allowing us to send syslog messages over the network.
type netConn struct {
	conn  net.Conn
	local bool
}

// writeString formats syslog messages using time.RFC3339 and includes the
// hostname, and sends the message to the connection.
func (n *netConn) writeString(framer Framer, formatter Formatter, p Priority, hostname, tag, msg string) error {
	fmt.Println("Framer 17 :>", framer)
	fmt.Println("Formatter 18:>", formatter)
	if framer == nil {
		framer = DefaultFramer
	}
	if formatter == nil {
		formatter = DefaultFormatter
	}
	fmt.Println("Framer 25 :>", framer)
	fmt.Println("Formatter26:>", formatter)
	formattedMessage := framer(formatter(p, hostname, tag, msg))

	fmt.Println("Formatted Message :>", formattedMessage)
	_, err := n.conn.Write([]byte(formattedMessage))
	return err
}

// close the network connection
func (n *netConn) close() error {
	return n.conn.Close()
}
