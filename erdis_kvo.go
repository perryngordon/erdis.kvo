package main


import (
//	"fmt"
	"os"
	"time"
        "strings"

	"github.com/nats-io/nats.go"
	//"github.com/nats-io/nats.go/jetstream"

)


func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
          println("problem here")
	}
	defer nc.Drain()


	sub, _ := nc.SubscribeSync("erdis.kvo.>")
	msg, _ := sub.NextMsg(10 * time.Millisecond)
	//println(msg)

         for {
	         msg, _ = sub.NextMsg(10 * time.Millisecond)
		 if msg != nil {
			 msgHandler(msg)
			 //println(string(msg.Data))
                         //msg.Respond([]byte("hello, " + name))
		 }

	 }

	defer sub.Unsubscribe()

}

func msgHandler(msg *nats.Msg){
     // make error
     cmd_args := strings.Split(msg.Subject, ".")
     cmd := cmd_args[2]  + "." + cmd_args[3]

     if cmd == "list.append" {
          list_append(msg)
     }
}

