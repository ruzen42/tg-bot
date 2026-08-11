package logger	

import (
	"fmt"
	"time"
	"github.com/go-color-term/go-color-term/coloring"
)

func getTime() string {
	return time.Now().Format("15:04:05")
}

func log(type string, color func(string) string, msg string) {
	fmt.Printf("%s %s %s\n", getTime(), color(type), msg)
}

func Debug(msg string) { log("DEBUG", coloring.Grey, msg) }
func Info(msg string)  { log("INFO", coloring.Green, msg) }
func Warn(msg string)  { log("WARN", coloring.Yellow, msg) }
func Error(msg string) { log("ERROR", coloring.Red, msg) }
func Fatal(msg string) {
	log(" FATAL ", coloring.Red, msg)
	os.Exit(1)
}
