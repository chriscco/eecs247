package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

const (
	filename = "text/daredevil.txt"
)

// Sender call SendMessage() when there is data in channel 
//	@param msgCh 
//	@param producer 
//	@param wg 
func Sender(msgCh <- chan string, producer sarama.SyncProducer, wg *sync.WaitGroup) {
	defer wg.Done() 
	for msg := range msgCh {
		SendMessage(msg, producer)
	}
}

// Reader reader from file, send lines to channel 
// remember to close channel before exiting 
//	@param filename filename
//	@param producer 
//	@param msgCh Channel to transmit data by lines 
func Reader(filename string, producer sarama.SyncProducer, msgCh chan <- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close() 
	scanner := bufio.NewScanner(file) 

	start := time.Now() 
	for scanner.Scan() {
		msgCh <- scanner.Text() 
	}
	// send this to Spark to indicate all data has been sent to Kakfa
	// doesn't seem to work properly if also using goroutine 
	msgCh <- "__END__"
	fmt.Printf("Time Taken (Go Producer): %dms\n", time.Since(start).Milliseconds())

	if err := scanner.Err(); err != nil {
		log.Fatal("Scanner error: ", err)
	}
	close(msgCh)
}

// main 
func main() {

}   