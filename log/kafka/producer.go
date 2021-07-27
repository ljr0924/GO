package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func NewKafProducer(addr []string) (*KafProducer, error) {
	p, err := CreateKafkaProducer(addr)
	if err != nil {
		fmt.Errorf("create kafka error: %s", err.Error())
		return nil, err
	}
	return &KafProducer{Producer: p}, nil
}

func CreateKafkaProducer(addr []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()

	// 等待服务器所有副本都保存成功才响应
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 随机的分区类型：返回一个分区器，返回一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 是否等待成功后响应
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("create kafka producer err: ", err.Error())
		return nil, err
	}
	fmt.Println("create kafka producer success")

	return producer, err
}

type KafProducer struct {
	Producer sarama.SyncProducer
}

func (k *KafProducer) PutIntoKafka(keyStr, valStr string) {

	// 构建发送的消息
	msg := &sarama.ProducerMessage{
		Topic: "logs",
		Key:   sarama.StringEncoder(keyStr),
		Value: sarama.StringEncoder(valStr),
	}
	partition, offset, err := k.Producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg fail")
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("partition = %d, offset = %d, value = %s", partition, offset, valStr)

}
