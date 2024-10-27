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
     
	/**
        use_jetstream := os.Getenv("NATS_ERDIS_KVO_USE_JS")
        if use_jetstream == "true" {
             println("jetstream disabled")
             js, _ := jetstream.New(nc)
             ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
             defer cancel()
             // create stream (idempotent)
	     cfg := jetstream.StreamConfig{
		Name:      "ERDIS_KVO",
		Retention: jetstream.WorkQueuePolicy,
		Subjects:  []string{"erdis.kvo.>"},
	     }
             stream, err := js.CreateStream(ctx, cfg)
	     if stream == nil{
		     println("srteam not created")
	     }
	     if err != nil{
		     println(err.Error())
	     }

	     hostname, err := os.Hostname()
	     println(hostname)
             cons, err := js.CreateOrUpdateConsumer(ctx, "ERDIS_KVO", jetstream.ConsumerConfig{
             //cons, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
                 Name: "erdis_kvo_"+hostname,
             })
             if err != nil {
		      println("error creating consumer")
                      println(err.Error())
             }
             iter, _ := cons.Messages()
             for {
                msg, err := iter.Next()
                // Next can return error, e.g. when iterator is closed or no heartbeats were received
                if err != nil {
                    //handle error
                }
                //fmt.Printf("Received a JetStream message: %s\n", string(msg.Data()))
                fmt.Printf("message: %s\n", msg.Subject())
		msg2 := jetstreamMsgToNatsMsg(msg)
                msgHandler(msg2)
                //msgHandler(msg.Parent)
		//msg2 :=  nats.Msg(msg) 
		//msgHandler(msg2)
                msg.Ack()
             }
             //iter.Stop
        
        }else{
        println("jetstream disabled")
	**/

	////sub, _ := nc.SubscribeSync("erdis.kvo.>", "erdis_kvo")
	////msg, _ := sub.NextMsg(10 * time.Millisecond)

        //// for {
	////         msg, _ = sub.NextMsg(10 * time.Millisecond)
	////	 if msg != nil {
	////		 msgHandler(msg)
	////	 }

	//// }

	////defer sub.Unsubscribe()
        //}

        wg := sync.WaitGroup{}
        wg.Add(1)

        //sub, err := nc.QueueSubscribe("erdis.kvo.>", "queue_erdis_kvo", msgHandler(msg *nats.Msg)) // {
	//if sub == nil{}
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
