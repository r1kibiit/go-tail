package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ch := make(chan string, 10)
	f, err := os.Open("test.log")
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	// 末尾10行以外破棄
	for sc.Scan() {
		if len(ch) >= 10 {
			<-ch
		}
		ch <- sc.Text()
	}

	go func() {
		for {
			t, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(t)
		}
	}()

	go func() {
		for {
			sc := bufio.NewScanner(f)
			for sc.Scan() {
				ch <- sc.Text()
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(time.Minute)
}
