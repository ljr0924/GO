package kafka

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

func TestMockLog(t *testing.T) {

	type Log struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	logDirs := []string{
		"/Users/banana/git_project/GO/log/log_dir1/log.log",
		"/Users/banana/git_project/GO/log/log_dir2/log.log",
		"/Users/banana/git_project/GO/log/log_dir3/log.log",
		"/Users/banana/git_project/GO/log/log_dir4/log.log",
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(4)

	for _, d := range logDirs {
		go func(addr string) {
			file, err := os.OpenFile(addr, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				t.Logf("open file err: %s", err.Error())
				wg.Done()
				return
			}
			defer file.Close()
			t.Logf("stat to mock: %s", addr)
			w := bufio.NewWriter(file)
			var cnt int
			for {
				select {
				case <- signals:
					wg.Done()
					return
				default:
					data, _ := json.Marshal(&Log{Name: fmt.Sprintf("name%d", cnt), Age: cnt})
					log := append(data, []byte("\n")...)
					t.Logf("write log: %s", log)
					_, _ = w.WriteString(string(log))
					_ = w.Flush()
					cnt++
					time.Sleep(3*time.Second)
				}
			}
		}(d)
	}

	wg.Wait()
}
