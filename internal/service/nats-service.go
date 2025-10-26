package service

import (
	"encoding/json"
	"log"
	"time"

	"order-service-demo/internal/db"
	"order-service-demo/internal/models"
	"order-service-demo/pkg/cache"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/stan.go"
)

func InitNATSAndSubscribe(natsURL string, dbPool *pgxpool.Pool) {
	var sc stan.Conn
	var err error

	for attempts := 0; attempts < 5; attempts++ {
		sc, err = stan.Connect("test-cluster", "order-client", stan.NatsURL(natsURL), stan.ConnectWait(10*time.Second))
		if err == nil {
			break
		}
		log.Printf("NATS connect attempt %d failed: %v", attempts+1, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to NATS after retries:", err)
	}

	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Printf("Invalid JSON: %v", err)
			return
		}
		if order.OrderUID == "" {
			log.Println("Invalid order: no UID")
			return
		}

		if err := db.SaveOrder(dbPool, order, m.Data); err != nil {
			log.Printf("DB error: %v", err)
			return
		}

		cache.CacheMu.Lock()
		cache.Cache[order.OrderUID] = string(m.Data)
		cache.CacheMu.Unlock()

		m.Ack()
	}, stan.DurableName("order-durable"), stan.SetManualAckMode())
	if err != nil {
		log.Fatal("Subscribe error:", err)
	}

	log.Println("Subscribed to NATS channel 'orders'")

}
