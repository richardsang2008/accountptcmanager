package utility

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/richardsang2008/accountptcmanager/model"


)

type Log struct {
}


func (l *Log) New(logfile string, loglevel model.LogLevel)  {

	log.SetOutput(&lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     3,    //days
		Compress:   true, // disabled by default
	})
	switch loglevel{
	case model.DEBUG:
		log.SetLevel(log.DebugLevel)
	case model.INFO:
		log.SetLevel(log.InfoLevel)
	case model.ERROR:
		log.SetLevel(log.ErrorLevel)
	case model.PANIC:
		log.SetLevel(log.PanicLevel)
	case model.WARNING:
		log.SetLevel(log.WarnLevel)
	}
	/*f, err := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	return f,nil*/

}
/*func (l *Log) Close(f *os.File) {
	f.Close()
}*/
func (l *Log) Debug(args ...interface{}) {
	log.Debug(args)
}
func (l *Log) Panic(args ...interface{}) {
	log.Panic(args)
}
func (l *Log) Info(args ...interface{}) {
	log.Info(args)
}
func (l *Log) Error(args ...interface{}) {
	log.Error(args)
}
func (l *Log) Warning(args ...interface{}) {
	log.Warning(args)
}
