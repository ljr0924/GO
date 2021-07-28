package kafka_demo

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"sync"

	"testing"
)

func TestConsumer(t *testing.T) {

	topic := []string{"test1"}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 初始化消费者
	consumer, err := cluster.NewConsumer(Address, "group-1", topic, config)
	if err != nil {
		t.Fatalf("init consumer err: %s", err.Error())
	}
	defer consumer.Close()

	go func() {
		for err := range consumer.Errors() {
			t.Errorf("consumer err: %s", err.Error())
			return
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	Loop:
	for {
		t.Logf("start")
		select {
		case msg, ok := <-consumer.Messages():
			t.Logf("msg %v", msg)
			if ok {
				t.Logf("get message: %s", string(msg.Value))
				consumer.MarkOffset(msg, "")
			}
		case <- signals:
			break Loop
		}
	}

}

func TestSaramaConsumer(t *testing.T) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	c, err := sarama.NewConsumer(Address, config)
	if err != nil {
		t.Fatalf("new consumer err: %s", err.Error())
	}
	topics, err := c.Topics()
	if err != nil {
		t.Fatalf("get topics err: %s", err.Error())
	}

	t.Logf("topics: %v", topics)

	topic := "logs"

	partitions, err := c.Partitions(topic)
	if err != nil {
		t.Fatalf("get topics err: %s", err.Error())
	}

	var pcList []*sarama.PartitionConsumer
	for _, p := range partitions {
		pc, err := c.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			t.Errorf("get consumer partition err: %s", err.Error())
			continue
		}
		t.Logf("new consumer partition topic:%s partition: %d", topic, p)

		pcList = append(pcList, &pc)

	}

	var wg sync.WaitGroup

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	wg.Add(len(pcList))

	for _, c := range pcList {
		defer (*c).AsyncClose()
		go func(cp sarama.PartitionConsumer) {
			t.Log("start")
			for {
				select {
				case msg := <- cp.Messages():
					t.Logf("get message from topic: %s partition: %d ---> %s", msg.Topic, msg.Partition, msg.Value)
				case err := <- cp.Errors():
					t.Errorf("get message from topic: %s partition: %d ---> %s", err.Topic, err.Partition, err.Err.Error())
					wg.Done()
					return
				case <- signals:
					t.Log("stop")
					wg.Done()
					return
				}
			}
		}(*c)
	}

	wg.Wait()

}
