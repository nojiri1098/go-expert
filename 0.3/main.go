package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type group struct {
	singleflight.Group
	sync.WaitGroup
}

func (g *group) do(key string) {
	g.Add(1)

	go func() {
		defer g.Done()

		v, err, shared := g.Group.Do("key", func() (interface{}, error) {
			time.Sleep(1 * time.Second)
			log.Printf("cached %s\n", key)
			return key, nil
		})
		log.Println(v, err, shared)
	}()
}

func (g *group) Forgot() {
	g.Group.Forget("key")
}

func main() {
	var g group
	g.do("foo")
	g.do("bar")
	g.Wait()
}
