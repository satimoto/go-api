package sync

import (
	"context"
	"log"
	"sync"
	"time"
)

func (r *SyncService) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Sync Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup

	go r.startLoop()
}

func (r *SyncService) Sync() {
	r.waitGroup.Add(1)
	r.mutex.Lock()
	log.Printf("Start Poi sync")

	r.PoiResolver.SyncronizePois()

	log.Printf("End Poi sync")
	r.waitGroup.Done()
	r.mutex.Unlock()
}

func (r *SyncService) startLoop() {
	timeNow := time.Now().UTC()
	lastUpdated := time.Date(
		timeNow.Year(),
		timeNow.Month(),
		timeNow.Day(),
		1,
		30,
		0,
		0,
		timeNow.Location())
	startTime := lastUpdated.Add(time.Hour * 24)

	if lastUpdated.After(timeNow) {
		startTime = lastUpdated
	}

	waitDuration := startTime.Sub(timeNow)

	for {
		select {
		case <-r.shutdownCtx.Done():
			log.Printf("Shutting down Sync Service")
			return
		case <-time.After(waitDuration):
		}

		waitDuration = time.Hour * 24
		r.Sync()
	}
}
