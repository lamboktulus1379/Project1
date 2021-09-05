package Config

<<<<<<< HEAD
import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)

func Produce(ctx context.Context) {
	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})

		if err != nil {
			panic("Could not write message" + err.Error())
		}

		fmt.Println("writes: ", i)
		i++

		time.Sleep(time.Second)
	}
}

func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
		GroupID: "my-gra-tech",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}
=======
import "fmt"

func Send() {
	fmt.Print("Send kafka")
}
>>>>>>> 61f8f77 (Added Typing Controller)
