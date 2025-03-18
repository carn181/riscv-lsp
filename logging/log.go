package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init(){
	f, _ := os.OpenFile("/tmp/lsp.log", os.O_RDWR | os.O_CREATE, 0644)

	Logger = log.New(f, "", 0)
}

func ReceivedMessage(str string){
	Logger.Printf("Received Message:\n%s\n", str)
}

func SentMessage(str string){
	Logger.Printf("Sent Message:\n%s\n", str)
}
