package conn

import (
	"bufio"
	"net"
	"bytes"
)

/*
This package is present on both the client and server recieving it. This will effectively
allow communication in a bidirectional way by only reading or doing an action when the
sender has said to do so.

Both parties must have all od the constants, such that they both can understand and send
messages in the same format.

The structure of the information is:


First 6 bytes: Identifier
7th byte: Command ID
8th and onward: Data packaged with it

Data packaged by commands is:

JOIN_REQUEST: Owner in string
CLIENT_EXIT_REQUEST, PING_REQUEST, PING_RESPONSE: None
SERVER_EXIT_REQUEST: Who left the server in string
JOIN_NOTIF: Who joined the server in string
MESSAGE: Message in string

Useage of the commands:

JOIN_REQUEST: Sent by client to server to notify a user join
JOIN_NOTIF: Sent by server to notify clients that someone has joined the server
EXIT_NOTIF: Sent by server to notify clients that someone has left the server
PING: Sent by server to check if clients are still online
PING_RESPONSE: Sent by client to respond to a server ping

MESSAGE: Sent by clients to send a message to the server, which then
the server will broadcast back to every client and when the clients recieve
said request, they will display the message sent
*/

type commandID byte

var verify []byte = []byte("XCHAT")

const (
	DELIMITER byte = 0xFF
	MESSAGE_DELIMITER byte = 0xF9
)

const (
	CMD_JOIN_REQUEST commandID = iota
	CMD_JOIN_NOTIF
	CMD_EXIT_NOTIF
	CMD_PING
	CMD_PING_RESPONSE
	CMD_MESSAGE
)

type command struct {
	id   commandID
	data []byte
}

func SendCommand(conn net.Conn, id commandID, data []byte) error {
	writer := bufio.NewWriter(conn)
	var err error
	if len(data) != 0 {
		_, err = writer.Write(append(append(append(verify, byte(id)), data...), DELIMITER))
	} else {
		_, err = writer.Write(append(append(verify, byte(id)), DELIMITER))
	}
	return err
}

func RetrieveCommand(conn net.Conn) (command) {
	reader := bufio.NewReader(conn)
	comm, err := reader.ReadBytes(DELIMITER)

	if err != nil || !bytes.Equal(comm[:len(verify)], verify){
		panic("Error at retrieving command: " + err.Error())
	}

	return command{commandID(comm[len(verify)]), comm[len(verify)+1:len(comm)-1]}

}

func OwnerMessageToBytes(owner string, message string) []byte {
	return []byte(append(append([]byte(owner), MESSAGE_DELIMITER), []byte(message)...))
}

func BytesToOwnerMessage(bytes []byte) (string, string) {
	for i, b := range(bytes) {
		if b == MESSAGE_DELIMITER {
			return string(bytes[:i]), string(bytes[i+1:])
		}
	}
	return "MissingNo", "Error getting message body"
}
