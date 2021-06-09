package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	flag.IntVar(&nodeId, "n", 0, "node ID")
	flag.StringVar(&target, "t", "", "set node into sender mode to specified target address")
	flag.IntVar(&requests, "r", 50, "ammount of reqeusts to be done")
	flag.StringVar(&delay, "d", "1s", "delay between requests")
	flag.Parse()

	if target != "" {
		startSender()
	} else {
		startReveiver()
	}
}

func startSender() {
	var err error

	fmt.Println("Sleeping for 2 seconds ...")
	time.Sleep(2 * time.Second)

	d, err := time.ParseDuration(delay)
	if err != nil {
		panic(err)
	}
	t := time.NewTicker(d)
	i := 0
	rttC := make(chan time.Duration)
	rtts := make([]time.Duration, 0, requests)

	go func() {
		for rtt := range rttC {
			rtts = append(rtts, rtt)
		}
	}()

	for {
		if i >= requests {
			time.Sleep(1 * time.Second)
			close(rttC)
			break
		}
		i++
		<-t.C
		go func() {
			if err = send(rttC); err != nil {
				log.Println(err)
			}
		}()
	}

	var avg, min, max time.Duration
	var der float64
	for _, rtt := range rtts {
		avg += rtt
		if min == 0 || rtt < min {
			min = rtt
		}
		if rtt > max {
			max = rtt
		}
	}
	avg = avg / time.Duration(len(rtts))

	for _, rtt := range rtts {
		der += math.Pow(float64(rtt-avg), 2)
	}
	der = der / float64(len(rtts))
	derd := time.Duration(math.Sqrt(der))

	fmt.Printf("rtt stats over %d requests (1 every %s):\n", len(rtts), delay)
	fmt.Printf("Average: %s (%dns)\n", avg, avg)
	fmt.Printf("Min:     %s (%dns)\n", min, min)
	fmt.Printf("Max:     %s (%dns)\n", max, max)
	fmt.Printf("SD:      %s (%dns)\n", derd, derd)
}

func send(rttC chan time.Duration) (err error) {
	startT := time.Now()

	w := NewTimeWrapper(RandomBook())

	buff := bytes.NewBuffer([]byte{})
	if err = json.NewEncoder(buff).Encode(w); err != nil {
		return
	}

	res, err := http.Post(target, "application/json", buff)
	if err != nil {
		return
	}

	tr := time.Now().UnixNano()
	json.NewDecoder(res.Body).Decode(w)
	w.AddStageNow("decoded")
	w.AddStage("received", tr)

	rtt := time.Since(startT)
	rttC <- rtt

	return
}

func startReveiver() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Post("/", func(ctx *fiber.Ctx) (err error) {
		w := new(TimeWrapper)
		t := time.Now().UnixNano()

		// Deserialize payload into object
		if err = ctx.BodyParser(w); err != nil {
			return
		}

		w.AddStageNow("parsed")
		w.AddStage("received", t)

		// Serialize object back to JSON and
		// send response
		return ctx.JSON(w)
	})

	panic(app.Listen(":80"))
}
