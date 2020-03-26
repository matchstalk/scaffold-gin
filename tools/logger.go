package tools

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func InitLogger() {
	switch RunMode(viper.GetString("app.runmode")) {
	case RunModeDev, RunModeTest:
		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
	case RunModeProd:
		f, err := os.OpenFile(viper.GetString("logger.dir")+"/shop-"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		if err != nil {
			log.Fatal("log  init failed")
		}

		var info os.FileInfo
		info, err = f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		fileWriter := logFileWriter{f, info.Size()}
		log.SetOutput(&fileWriter)
		log.SetLevel(log.ErrorLevel)
	}

	log.SetReportCaller(true)
}

type logFileWriter struct {
	file *os.File
	size int64
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)
	//每天一个文件
	if p.file.Name() != viper.GetString("logger.dir")+"/shop-"+time.Now().Format("2006-01-02")+".log" {
		p.file.Close()
		p.file, _ = os.OpenFile(viper.GetString("logger.dir")+"/shop-"+time.Now().Format("2006-01-02")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		p.size = 0
	}
	return n, e
}
