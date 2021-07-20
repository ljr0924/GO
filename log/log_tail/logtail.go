package log_tail

import (
	"bufio"
	"fmt"
	"github.com/hpcloud/tail"
	"os"
	"path"
	"runtime"
	"time"
)

func main() {

	logRelative := "../log_dir/log.log"
	_, filePath, _, _ := runtime.Caller(0)
	fmt.Println(filePath)

	// 拼接日志文件路径
	dataPath := path.Join(path.Dir(filePath), logRelative)
	fmt.Println(dataPath)

	go writeLog(dataPath)



	tailFile, err := tail.TailFile(dataPath, tail.Config{
		ReOpen:    true, // 文件被移除或被打包，需要重新打开
		Follow:    true, // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取
		MustExist: false, // 支持文件不存在，文件不存在也不会报错
		Poll:      true,
	})
	if err != nil {
		fmt.Printf("tail file err: %s\n", err.Error())
		return
	}

	for {
		msg, ok := <- tailFile.Lines
		if !ok {
			fmt.Printf("file reopen, file name: %s\n", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
		}

		fmt.Println("msg: ", msg.Text)
	}


}

func writeLog(dataPath string) {
	logFile, err := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file err: %s\n", err.Error())
		return
	}

	w := bufio.NewWriter(logFile)
	for i := 0; i < 100; i++ {
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintln(w, "current time is "+timeStr)
		time.Sleep(time.Millisecond * 200)
		w.Flush()
	}

	logBak := time.Now().Format("20060102150405") + ".log"
	logBak = path.Join(path.Dir(dataPath), logBak)

	logFile.Close()
	err = os.Rename(dataPath, logBak)
	if err != nil {
		fmt.Println("rename error ", err.Error())
		return
	}

}

