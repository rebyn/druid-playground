package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type View struct {
	Timestamp  time.Time `json:"timestamp"`
	CreatorId  uint      `json:"creator_id"`
	UserId     uint      `json:"user_id"`
	PostId     uint      `json:"post_id"`
}

func main() {
	dayEventsChan := make(chan *View, 1000)
	timeZone, _ := time.LoadLocation("Asia/Bangkok")
	startDate := time.Date(2020, 9, 01, 0, 0, 0, 0, timeZone)
	numberOfDays := 7
	numberOfCreators := 5
	maxViewsPerClientPerDay := 500

	// receive Views and put all of them to file
	writeDoneChan := make(chan struct{}, 0)
	go func(dayChan <-chan *View) {
		fp, _ := os.Create("./generated/events.json")
		defer fp.Close()

		i := 0
		for tr := range dayChan {
			if i%10000 == 0 {
				fmt.Println("Syncing file to disk, written:", i)
				fp.Sync()
			}

			b, err := json.Marshal(tr)
			if err != nil {
				panic(err)
			}

			if _, err := fp.Write(b); err != nil {
				panic(err)
			}
			if _, err := fp.Write([]byte("\n")); err != nil {
				panic(err)
			}
			i++
		}

		writeDoneChan <- struct{}{}
	}(dayEventsChan)

	// generate all Views
	for d := 0; d < numberOfDays; d++ {
		for c := 0; c < numberOfCreators; c++ {
			if c%10 == 0 {
				fmt.Println("creators done:", c, "for day:", d)
			}

			numToday := rand.Intn(maxViewsPerClientPerDay)
			startHour := rand.Intn(23)

			for i := 0; i < numToday; i++ {
				trDate := startDate.
					AddDate(0, 0, d).
					Add(time.Hour * time.Duration(startHour)).
					Add(time.Minute * time.Duration(rand.Intn(59)))

				dayEventsChan <- &View{
					Timestamp:  trDate,
					UserId:     uint(rand.Intn(numberOfCreators)),
					CreatorId:  uint(c + 1),
					PostId:     uint(rand.Intn(100)),
				}
				startHour++
			}
		}
	}

	close(dayEventsChan)
	fmt.Println("finished View event generation")

	<-writeDoneChan
	fmt.Println("finished writing")
}
