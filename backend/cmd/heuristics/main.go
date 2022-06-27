package main

import (
	"log"
	"openreplay/backend/internal/config/heuristics"
	"openreplay/backend/pkg/handlers"
	web2 "openreplay/backend/pkg/handlers/web"
	"openreplay/backend/pkg/intervals"
	logger "openreplay/backend/pkg/log"
	"openreplay/backend/pkg/messages"
	"openreplay/backend/pkg/queue"
	"openreplay/backend/pkg/queue/types"
	"openreplay/backend/pkg/sessions"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Llongfile)

	// Load service configuration
	cfg := heuristics.New()

	// HandlersFabric returns the list of message handlers we want to be applied to each incoming message.
	handlersFabric := func() []handlers.MessageProcessor {
		return []handlers.MessageProcessor{
			// web handlers
			&web2.ClickRageDetector{},
			&web2.CpuIssueDetector{},
			&web2.DeadClickDetector{},
			&web2.MemoryIssueDetector{},
			&web2.NetworkIssueDetector{},
			&web2.PerformanceAggregator{},
			// iOS's handlers
			//&ios2.AppNotResponding{},
			//&ios2.ClickRageDetector{},
			//&ios2.PerformanceAggregator{},
			// Other handlers (you can add your custom handlers here)
			//&custom.CustomHandler{},
		}
	}

	// Create handler's aggregator
	builderMap := sessions.NewBuilderMap(handlersFabric)

	// Init logger
	statsLogger := logger.NewQueueStats(cfg.LoggerTimeout)

	// Init producer and consumer for data bus
	producer := queue.NewProducer()
	consumer := queue.NewMessageConsumer(
		cfg.GroupHeuristics,
		[]string{
			cfg.TopicRawWeb,
		},
		func(sessionID uint64, msg messages.Message, meta *types.Meta) {
			statsLogger.Collect(sessionID, meta)
			builderMap.HandleMessage(sessionID, msg, msg.Meta().Index)
		},
		false,
	)

	log.Printf("Heuristics service started\n")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	tick := time.Tick(intervals.EVENTS_COMMIT_INTERVAL * time.Millisecond)
	for {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			producer.Close(cfg.ProducerTimeout)
			consumer.Commit()
			consumer.Close()
			os.Exit(0)
		case <-tick:
			builderMap.IterateReadyMessages(func(sessionID uint64, readyMsg messages.Message) {
				producer.Produce(cfg.TopicAnalytics, sessionID, messages.Encode(readyMsg))
			})
			producer.Flush(cfg.ProducerTimeout)
			consumer.Commit()
		default:
			if err := consumer.ConsumeNext(); err != nil {
				log.Fatalf("Error on consuming: %v", err)
			}
		}
	}
}
