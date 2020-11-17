package blog

import (
	"fmt"
)

func Debug(message string) {
	fmt.Println("[Level Debug] message : " + message)
}

func Warn(message string) {
	fmt.Println("[Level Warn] message : " + message)
}

func Error(message string) {
	fmt.Println("[Level Error] message : " + message)
}

func Info(message string) {
	fmt.Println("[Level Info] message : " + message)
}
