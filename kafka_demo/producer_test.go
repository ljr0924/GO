package kafka_demo

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestSyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	p, err := sarama.NewSyncProducer(Address, config)
	if err != nil {
		t.Fatalf("connect err %s", err.Error())
		return
	}
	defer p.Close()

	topic := "test1"
	srcValue := "sync message index=%d"
	_, _, _ = p.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("-----------------start------------------"),
	})
	if err != nil {
		return
	}

	for i := 0; i < 100; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf(srcValue, i)),
		}
		partition, offset, err := p.SendMessage(msg)
		if err != nil {
			t.Logf("send message err: %s", err.Error())
		} else {
			t.Logf("send message success, partition: %d, offset: %d", partition, offset)
		}
		time.Sleep(time.Second * 1)
	}
}

func TestAsyncProducer(t *testing.T) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	p, err := sarama.NewAsyncProducer(Address, config)
	if err != nil {
		t.Errorf("new async producer err: %s", err.Error())
		return
	}

	defer p.AsyncClose()

	topic := "test1"
	srcValue := "sync message index=%d"

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	var success int
	var errNum int
	var cnt int
	Loop:
	for {
		cnt ++
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf(srcValue, cnt)),
		}

		select {
		case p.Input() <- msg:
		case <-p.Successes():
			success++
		case <-p.Errors():
			errNum++
		case <- sig:
			t.Logf("success num: %d", success)
			t.Logf("error num: %d", errNum)
			break Loop
		}
		time.Sleep(time.Second*3)
	}

}
