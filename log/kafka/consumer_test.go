package kafka

import "testing"

func TestConsumer(t *testing.T) {

	addrList := []string{"172.20.10.3:19092", "172.20.10.3:19093", "172.20.10.3:19094"}

	ConsumeLog(addrList)

}
