package log_tail

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func WatchLogFile(key string, dataPath string, ctx context.Context, keyChan chan<- string) {

	defer func() {
		if errCover := recover(); errCover != nil {
			fmt.Printf("goroutine watch %s panic\n", dataPath)
			fmt.Println(errCover)
			keyChan <- key
		}
	}()

	fmt.Println("begin goroutine watch log file ", dataPath)

	tailFile, err := tail.TailFile(dataPath, tail.Config{
		ReOpen:    true, // 文件被移除或被打包，需要重新打开
		Follow:    true, // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取
		MustExist: false, // 支持文件不存在，文件不存在也不会报错
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err: ", err)
		return
	}

	for {
		select {
		case msg, ok := <-tailFile.Lines:
			if !ok {
				fmt.Printf("file reopen, file name: %s\n", tailFile.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			fmt.Println("msg: ", msg.Text)
		case <- ctx.Done():
			fmt.Println("receive main goroutine exit msg")
			fmt.Println("watch log file ", dataPath, " goroutine exited" )
			return
		}
	}

}


