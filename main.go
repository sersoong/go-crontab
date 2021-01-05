package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * * * * *", hello)
	c.Start()
	select {}
}

func hello() {
	log.Println("hello")
}
