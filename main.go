package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c := make(chan int)
	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			c <- 1
		}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c <- 2
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c <- 3
	}()
	for i := range c {
		log.Println("i", i)
		time.Sleep(2 * time.Second)
	}
}
