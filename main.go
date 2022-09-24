package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type MyClient struct {
	WAClient       *whatsmeow.Client
	eventHandlerID uint32
}

func (mycli *MyClient) register() {
	mycli.eventHandlerID = mycli.WAClient.AddEventHandler(mycli.myEventHandler)
}

func (mycli *MyClient) myEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
		/*sender := v.Info.MessageSource.Sender
				cli.SendMessage(context.Background(), targetJID, "", &waProto.Message{
		    		Conversation: proto.String("Hello, World!"),
				})
		*/
		mycli.WAClient.SendMessage(context.Background(), v.Info.MessageSource.Sender, "", &waProto.Message{
			Conversation: proto.String("Hello, World!"),
		})
	case *events.Receipt:
		fmt.Println("Received a receipt!", v.Chat.User)
	}
}

func main() {

	cli, err := WAConnect()
	if err != nil {
		fmt.Println(err)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	cli.WAClient.Disconnect()
}
