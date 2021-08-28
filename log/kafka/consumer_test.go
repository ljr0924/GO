package kafka

import "testing"

func TestConsumer(t *testing.T) {

	addrList := []string{"192.168.2.167:19092", "192.168.2.167:19093", "192.168.2.167:19094"}

	ConsumeLog(addrList)

}
