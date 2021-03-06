package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/dedis/student_19_libp2p_tlc/model"
	"github.com/dedis/student_19_libp2p_tlc/protobuf/messagepb"
	"github.com/dedis/student_19_libp2p_tlc/transport/test_utils"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"os"
	"testing"
	"time"
	"unicode/utf8"
)

var usernames = []string{"a@localhost.localdomain", "b@localhost.localdomain", "c@localhost.localdomain", "d@localhost.localdomain", "e@localhost.localdomain"}
var passwords = []string{"apassword", "bpassword", "cpassword", "dpassword", "epassword"}

func TestSendMail(t *testing.T) {
	//usernames := []string{"a@localhost.localdomain", "b@localhost.localdomain", "c@localhost.localdomain", "d@localhost.localdomain", "e@localhost.localdomain"}
	SendMail(usernames[0], []string{usernames[1]}, "444", []byte("Testing from mail package"), "apassword")
}

func TestGetMail(t *testing.T) {
	GetMail(usernames[4], passwords[4], 5)
}

// setupHosts is responsible for creating tlc nodes
func setup(n int) []*model.Node {
	// nodes used in tlc model
	nodes := make([]*model.Node, n)

	for i := range nodes {
		// First delete existing emails inside inbox
		deleteMail(usernames[i], passwords[i])

		//var comm model.CommunicationInterface
		var comm *mail
		comm = new(mail)

		comm.username = usernames[i]
		comm.password = passwords[i]
		comm.recentIndex = 1
		comm.addressBook = usernames
		comm.imapSession = StartImapSession(usernames[i], passwords[i])

		nodes[i] = &model.Node{
			Id:           i,
			TimeStep:     0,
			ThresholdWit: n,
			ThresholdAck: n,
			Acks:         0,
			ConvertMsg:   &messagepb.Convert{},
			Comm:         comm,
			History:      make([]model.Message, 0)}
	}
	return nodes
}

func deleteMail(username string, password string) {
	// Connect to server
	c, err := client.DialTLS(mailReceiveServer+":993", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	if err := <-done; err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	if mbox.Messages == 0 {
		fmt.Printf("Error : %v\n", "No message in inbox")
		return
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(1, mbox.Messages)

	// First mark the message as deleted
	item := imap.FormatFlagsOp(imap.AddFlags, true)
	flags := []interface{}{imap.DeletedFlag}
	if err := c.Store(seqset, item, flags, nil); err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	// Then delete it
	if err := c.Expunge(nil); err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	fmt.Println("Successfully deleted emails")

}

func TestMail(t *testing.T) {
	logFile, _ := os.OpenFile("../../logs/NoFailure_Mail.log", os.O_RDWR|os.O_CREATE, 0666)
	model.Logger1 = log.New(logFile, "", log.Ltime|log.Lmicroseconds)
	nodes := setup(5)
	test_utils.StartTest(nodes, 10, 0)
	test_utils.LogOutput(t, nodes)
}

func TestMailProto(t *testing.T) {
	msg := model.Message{Source: 2, Step: 4, MsgType: model.Raw, History: []model.Message{}}
	msgBytes := *messagepb.Convert{}.MessageToBytes(msg)
	fmt.Println("BYTES", msgBytes)
	fmt.Println(utf8.Valid(msgBytes))
	SendMail(usernames[0], usernames, "", msgBytes, passwords[0])
	time.Sleep(time.Second)
	//data := GetMail(usernames[2], passwords[2], 1)
	data := GetMailSubject(StartImapSession(usernames[2], passwords[2]), 1)
	fmt.Println("BYTES", data)
	fmt.Println(messagepb.Convert{}.BytesToModelMessage(data))
}
