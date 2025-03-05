package main


import (
	"os"
	//"time"
        "strings"

	"github.com/nats-io/nats.go"
        "github.com/nats-io/nats.go/jetstream"
	//"context"
	//"fmt"
	"sync"
	"log"
)


func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
          println("problem connecting to NATS")
	}
	defer nc.Drain()
     

        wg := sync.WaitGroup{}
        wg.Add(1)

        if _, err := nc.QueueSubscribe("erdis.kvo.>", "queue_erdis_kvo", func(msg *nats.Msg)  {
		msgHandler(msg)
		//wg.Done()
        }); err != nil {
            log.Fatal(err)
	    println(err.Error())
        }

        wg.Wait()

}

func msgHandler(msg *nats.Msg){
     // make error
     cmd_args := strings.Split(msg.Subject, ".")
     cmd := cmd_args[2]  + "." + cmd_args[3]

     if cmd == "list.append" {
          list_append(msg)
     }
     if cmd == "list.remove_first" {
          list_remove_first(msg)
     }
     if cmd == "list.remove_all" {
          list_remove_all(msg)
     }
     if cmd == "list.valueExists" {
	  ocurrences := list_valueExists(msg)
	  str_ocurrences := intArrayToString(ocurrences,",")
          msg.Respond([]byte(str_ocurrences))
     }
     if cmd == "list.pop" {
         list_pop(msg)
     }
     if cmd == "list.push" {
	 list_push(msg)
     }

     if cmd == "list.kv_delete" {
         kv_delete(msg)
     }
}


//  google ai // Here's how you can convert a jetstream.Message to a nats.Msg
func jetstreamMsgToNatsMsg(jsMsg jetstream.Msg) *nats.Msg {

    return &nats.Msg{

        Subject: jsMsg.Subject(),

        Data:   jsMsg.Data(),

        Reply:  jsMsg.Reply(),

        // Add any other relevant metadata from jsMsg if needed

    }

}
