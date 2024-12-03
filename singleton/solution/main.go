package main

import (
	"log"
	"sync"
	"time"
)

type Config struct {
	logAllowed bool
}

func (c Config) LogAllowed() bool { return c.logAllowed }

/*
Điểm yếu của singleton -> application biết nhiều về biến này
-> Nên sử dụng interface
*/
type application struct {
	once   sync.Once
	config *Config
}

type Application interface {
	GetConfig() *Config
}

var singletonApp = &application{once: sync.Once{}}

func GetApplication() Application {
	return singletonApp
}

func (app *application) GetConfig() *Config {
	/*
		Racing if using this
		if app.config == nil {
			log.Println("It should be run only once. But you'll it many times!!")
			app.loadConfig()
		}
		return app.config
	*/
	if app.config == nil {
		app.once.Do(func() {
			log.Println("Loading config once and forever.")
			app.loadConfig()
		})
	}
	return app.config
}

func (app *application) loadConfig() {
	time.Sleep(100)
	app.config = &Config{logAllowed: true}
}

func main() {
	// Demo 1000 requests to service at a same time (1000 RPS)
	// I made this code for simple demo, not a real practice!

	rps := 1000
	wg := sync.WaitGroup{}
	wg.Add(rps)

	for i := 1; i <= rps; i++ {
		go func(idx int) {
			//requestHandler(idx)
			if GetApplication().GetConfig().LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func requestHandler(requestIdx int) {
	if GetApplication().GetConfig().LogAllowed() {
		log.Printf("Handling request %d .... please wait.\n", requestIdx)
	}
}
