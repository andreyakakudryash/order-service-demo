package db

import (
	"context"
	"encoding/json"
	"log"

	"order-service-demo/internal/models"
	"order-service-demo/pkg/cache"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(connString string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}

func SaveOrder(dbPool *pgxpool.Pool, order models.Order, data []byte) error {
	ctx := context.Background()

	if !isValidJSON(data) {
		log.Println("Invalid JSON data")
	}

	_, err := dbPool.Exec(ctx, "INSERT INTO orders (order_number, order_data) VALUES ($1, $2) ON CONFLICT (order_number) DO UPDATE SET order_data = EXCLUDED.order_data",
		order.OrderUID, data)
	return err
}

func LoadCacheFromDB(db *pgxpool.Pool) int {
	ctx := context.Background()
	rows, err := db.Query(ctx, "SELECT order_number, order_data FROM orders")
	if err != nil {
		log.Fatalf("Failed to load orders from database: %v", err)
	}
	defer rows.Close()

	count := 0
	cache.CacheMu.Lock()
	defer cache.CacheMu.Unlock()

	for rows.Next() {
		var uid string
		var data json.RawMessage
		if err := rows.Scan(&uid, &data); err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		cache.Cache[uid] = string(data)
		count++
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
	}

	return count
}

func isValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}
