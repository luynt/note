package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(CatchTimeOut())
}

func CatchTimeOut() error {
	err := make(chan error, 1)
	fmt.Println(time.Now().Second())
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Over 2s ")
			fmt.Println(time.Now().Second())
			err <- errors.New(t.String())
		}
	}()
	go func() {
		for c:=range time.Tick(1000){
			fmt.Printf("\r %v" ,c.String())
		}
	}()
	time.Sleep(3 * time.Second)
	defer ticker.Stop()
	fmt.Println("Ticker stopped")
	return <-err
}
