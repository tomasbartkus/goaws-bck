package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"

	sns "github.com/tomasbartkus/GoAws/gosns"
	sqs "github.com/tomasbartkus/GoAws/gosqs"
	"github.com/tomasbartkus/GoAws/conf"
)

func BadRequest(w http.ResponseWriter, req *http.Request) {
	resp := "Bad Request"
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, resp)
}

// hello world, the web server
func IndexServer(w http.ResponseWriter, req *http.Request) {
	switch req.FormValue("Action") {
	/*** SQS Actions ***/
	case "ListQueues":
		sqs.ListQueues(w, req)
	case "CreateQueue":
		sqs.CreateQueue(w, req)
	case "GetQueueAttributes":
		sqs.GetQueueAttributes(w, req)
	case "SendMessage":
		sqs.SendMessage(w, req)
	case "ReceiveMessage":
		sqs.ReceiveMessage(w, req)
	case "DeleteMessage":
		sqs.DeleteMessage(w, req)
	case "GetQueueUrl":
		sqs.GetQueueUrl(w, req)
	case "PurgeQueue":
		sqs.PurgeQueue(w, req)
	case "DeleteQueue":
		sqs.DeleteQueue(w, req)

	/*** SNS Actions ***/
	case "ListTopics":
		sns.ListTopics(w, req)
	/*** SNS Actions ***/
	case "CreateTopic":
		sns.CreateTopic(w, req)
	case "DeleteTopic":
		sns.DeleteTopic(w, req)
	case "Subscribe":
		sns.Subscribe(w, req)
	case "SetSubscriptionAttributes":
		sns.SetSubscriptionAttributes(w, req)
	case "ListSubscriptionsByTopic":
		sns.ListSubscriptionsByTopic(w, req)
	case "ListSubscriptions":
		sns.ListSubscriptions(w, req)
	case "Unsubscribe":
		sns.Unsubscribe(w, req)
	case "Publish":
		sns.Publish(w, req)

	/*** Bad Request ***/
	default:
		log.Println("Action:", req.FormValue("Action"))
		BadRequest(w, req)
	}
}

func main() {
	env := "Local"
	if len(os.Args) == 2 {
		env = os.Args[1]
	}
	var portNumber string
	var filename string
	flag.StringVar(&portNumber, "port", "", "Port number to listen on")
	flag.StringVar(&filename, "config", "", "config file location + name")
	flag.Parse()

	portNumber = conf.LoadYamlConfig(filename, env, portNumber)

	r := mux.NewRouter()
	r.HandleFunc("/", IndexServer).Methods("GET", "POST")
	r.HandleFunc("/queue/{queueName}", IndexServer).Methods("GET", "POST")

	go func() {
		log.Printf("GoAws listening on: 0.0.0.0:9324\n")
		err := http.ListenAndServe(":9324", r)
		log.Fatal(err)
	}()

	go func() {
		log.Printf("GoAws listening on: 0.0.0.0:9292\n")
		err := http.ListenAndServe(":9292", r)
		log.Fatal(err)
	}()

	log.Printf("GoAws listening on: 0.0.0.0:%s\n", portNumber)
	err := http.ListenAndServe(":"+portNumber, r)
	log.Fatal(err)
}
