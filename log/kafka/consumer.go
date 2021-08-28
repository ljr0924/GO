package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/mongo"
)

type Detail struct {
	Name string `bson:"name" json:"name"`
	Ip   string `bson:"ip" json:"ip"`
}

type Log struct {
	Key    string `json:"key" bson:"key"`
	Detail *Detail
}

func ProcessMsg(coll *mongo.Collection, msg *sarama.ConsumerMessage) error {

	var detail Detail
	err := json.Unmarshal(msg.Value, &detail)
	if err != nil {
		return err
	}

	log := &Log{
		Key:    string(msg.Key),
		Detail: &detail,
	}

	_, err = coll.InsertOne(context.TODO(), log)
	if err != nil {
		return err
	}

	return nil
}

func ConsumeLog(addr []string) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	c, err := sarama.NewConsumer(addr, config)
	if err != nil {
		fmt.Printf("new consumer err: %s\n", err.Error())
		return
	}
	topics, err := c.Topics()
	if err != nil {
		fmt.Printf("get topics err: %s\n", err.Error())
	}

	fmt.Printf("topics: %v\n", topics)

	topic := "logs"

	partitions, err := c.Partitions(topic)
	if err != nil {
		fmt.Printf("get topics err: %s", err.Error())
		return
	}

	var pcList []*sarama.PartitionConsumer
	for _, p := range partitions {
		pc, err := c.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("get consumer partition err: %s", err.Error())
			continue
		}
		fmt.Printf("new consumer partition topic: %s partition: %d\n", topic, p)

		pcList = append(pcList, &pc)

	}

	var wg sync.WaitGroup

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	wg.Add(len(pcList))

	coll := GetCollLog()

	for _, c := range pcList {
		defer (*c).AsyncClose()
		go func(cp sarama.PartitionConsumer) {
			fmt.Println("start")
			for {
				select {
				case msg := <-cp.Messages():
					fmt.Printf("get message from topic: %s partition: %d ---> key: %s, value: %s\n", msg.Topic, msg.Partition, msg.Key, msg.Value)
					err := ProcessMsg(coll, msg)
					if err != nil {
						fmt.Printf("process msg err: %s", err.Error())
					}
				case err := <-cp.Errors():
					fmt.Printf("get message from topic: %s partition: %d ---> %s\n", err.Topic, err.Partition, err.Err.Error())
					wg.Done()
					return
				case <-signals:
					fmt.Println("stop")
					wg.Done()
					return
				}
			}
		}(*c)
	}

	wg.Wait()

}
