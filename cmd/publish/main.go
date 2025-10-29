package main

import (
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publish-client", stan.NatsURL("nats://127.0.0.1:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	data1 := []byte(`{
	  "order_uid": "b563feb7b2b84b6test",
	  "track_number": "WBILMTESTTRACK",
	  "entry": "WBIL",
	  "delivery": {
	    "name": "Test Testov",
	    "phone": "+9720000000",
	    "zip": "2639809",
	    "city": "Kiryat Mozkin",
	    "address": "Ploshad Mira 15",
	    "region": "Kraiot",
	    "email": "test@gmail.com"
	  },
	  "payment": {
	    "transaction": "b563feb7b2b84b6test",
	    "request_id": "",
	    "currency": "RUB",
	    "provider": "wbpay",
	    "amount": 181700,
	    "payment_dt": 1637907727,
	    "bank": "alpha",
	    "delivery_cost": 150000,
	    "goods_total": 31700,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 9934930,
	      "track_number": "WBILMTESTTRACK",
	      "price": 45300,
	      "rid": "ab4219087a764ae0btest",
	      "name": "Mascaras",
	      "sale": 30,
	      "size": "0",
	      "total_price": 31700,
	      "nm_id": 2389212,
	      "brand": "Vivienne Sabo",
	      "status": 202
	    }
	  ],
	  "locale": "en",
	  "internal_signature": "",
	  "customer_id": "test",
	  "delivery_service": "meest",
	  "shardkey": "9",
	  "sm_id": 99,
	  "date_created": "2021-11-26T06:22:19Z",
	  "oof_shard": "1"
	}`)
	err = sc.Publish("orders", data1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 1 with UID: %s\n", "b563feb7b2b84b6test")

	data2 := []byte(`{
	  "order_uid": "U2RuKhsnAs52455rtest",
	  "track_number": "WBMSKCLOTH001",
	  "entry": "WBMSK",
	  "delivery": {
	    "name": "Ivan Ivanov",
	    "phone": "+79001234567",
	    "zip": "101000",
	    "city": "Moscow",
	    "address": "Tverskaya ulitsa 1",
	    "region": "Moscow Oblast",
	    "email": "ivan@example.ru"
	  },
	  "payment": {
	    "transaction": "U2RuKhsnAs52455rtest",
	    "request_id": "req-456",
	    "currency": "RUB",
	    "provider": "sberpay",
	    "amount": 8500,
	    "payment_dt": 1729862400,
	    "bank": "sberbank",
	    "delivery_cost": 500,
	    "goods_total": 8000,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 1234567,
	      "track_number": "WBMSKCLOTH001",
	      "price": 3000,
	      "rid": "cloth-rid-001",
	      "name": "Cotton T-Shirt",
	      "sale": 10,
	      "size": "M",
	      "total_price": 2700,
	      "nm_id": 5678901,
	      "brand": "Nike",
	      "status": 202
	    },
	    {
	      "chrt_id": 7654321,
	      "track_number": "WBMSKCLOTH001",
	      "price": 5000,
	      "rid": "cloth-rid-002",
	      "name": "Slim Fit Jeans",
	      "sale": 0,
	      "size": "32",
	      "total_price": 5000,
	      "nm_id": 8901234,
	      "brand": "Levi's",
	      "status": 202
	    }
	  ],
	  "locale": "ru",
	  "internal_signature": "sig-123",
	  "customer_id": "cust-ivan",
	  "delivery_service": "boxberry",
	  "shardkey": "1",
	  "sm_id": 10,
	  "date_created": "2025-10-24T10:00:00Z",
	  "oof_shard": "2"
	}`)
	err = sc.Publish("orders", data2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 2 with UID: %s\n", "U2RuKhsnAs52455rtest")

	data3 := []byte(`{
	  "order_uid": "1xF5TS7yiwGkmMwHtest",
	  "track_number": "WBSPBELEC002",
	  "entry": "WBSPB",
	  "delivery": {
	    "name": "Anna Petrova",
	    "phone": "+78121234567",
	    "zip": "191186",
	    "city": "Saint Petersburg",
	    "address": "Nevsky Prospekt 22",
	    "region": "Leningrad Oblast",
	    "email": "anna@example.ru"
	  },
	  "payment": {
	    "transaction": "1xF5TS7yiwGkmMwHtest",
	    "request_id": "req-789",
	    "currency": "RUB",
	    "provider": "tinkoffpay",
	    "amount": 50000,
	    "payment_dt": 1729866000,
	    "bank": "tinkoff",
	    "delivery_cost": 2000,
	    "goods_total": 48000,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 9876543,
	      "track_number": "WBSPBELEC002",
	      "price": 50000,
	      "rid": "elec-rid-001",
	      "name": "Smartphone Xiaomi Redmi Note 13",
	      "sale": 5,
	      "size": "0",
	      "total_price": 47500,
	      "nm_id": 1357924,
	      "brand": "Xiaomi",
	      "status": 202
	    }
	  ],
	  "locale": "ru",
	  "internal_signature": "sig-456",
	  "customer_id": "cust-anna",
	  "delivery_service": "cdek",
	  "shardkey": "3",
	  "sm_id": 20,
	  "date_created": "2025-10-24T11:00:00Z",
	  "oof_shard": "3"
	}`)
	err = sc.Publish("orders", data3)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 3 with UID: %s\n", "1xF5TS7yiwGkmMwHtest")

	data4 := []byte(`{
	  "order_uid": "oNiDeZMp3W7QWJzxtest",
	  "track_number": "WBEKBAPPL003",
	  "entry": "WBEKB",
	  "delivery": {
	    "name": "Sergey Sidorov",
	    "phone": "+73431234567",
	    "zip": "620000",
	    "city": "Yekaterinburg",
	    "address": "Lenina Prospekt 50",
	    "region": "Sverdlovsk Oblast",
	    "email": "sergey@example.ru"
	  },
	  "payment": {
	    "transaction": "oNiDeZMp3W7QWJzxtest",
	    "request_id": "req-012",
	    "currency": "RUB",
	    "provider": "yandexpay",
	    "amount": 15000,
	    "payment_dt": 1729869600,
	    "bank": "yandex",
	    "delivery_cost": 1500,
	    "goods_total": 13500,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 2468135,
	      "track_number": "WBEKBAPPL003",
	      "price": 5000,
	      "rid": "appl-rid-001",
	      "name": "Electric Kettle Bosch",
	      "sale": 20,
	      "size": "0",
	      "total_price": 4000,
	      "nm_id": 9753102,
	      "brand": "Bosch",
	      "status": 202
	    },
	    {
	      "chrt_id": 5310864,
	      "track_number": "WBEKBAPPL003",
	      "price": 10000,
	      "rid": "appl-rid-002",
	      "name": "Toaster Philips",
	      "sale": 5,
	      "size": "0",
	      "total_price": 9500,
	      "nm_id": 8642091,
	      "brand": "Philips",
	      "status": 202
	    }
	  ],
	  "locale": "ru",
	  "internal_signature": "sig-789",
	  "customer_id": "cust-sergey",
	  "delivery_service": "dpd",
	  "shardkey": "4",
	  "sm_id": 30,
	  "date_created": "2025-10-24T12:00:00Z",
	  "oof_shard": "4"
	}`)
	err = sc.Publish("orders", data4)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 4 with UID: %s\n", "oNiDeZMp3W7QWJzxtest")

	data5 := []byte(`{
	  "order_uid": "SMjNghpbzYPuRRb4test",
	  "track_number": "WBALMBKS004",
	  "entry": "WBALM",
	  "delivery": {
	    "name": "Aigerim Zhumagaliyeva",
	    "phone": "+77123456789",
	    "zip": "050000",
	    "city": "Almaty",
	    "address": "Abay Avenue 44",
	    "region": "Almaty Region",
	    "email": "aigerim@example.kz"
	  },
	  "payment": {
	    "transaction": "SMjNghpbzYPuRRb4test",
	    "request_id": "req-345",
	    "currency": "RUB",
	    "provider": "kaspi-pay",
	    "amount": 4000,
	    "payment_dt": 1729873200,
	    "bank": "kaspi",
	    "delivery_cost": 1000,
	    "goods_total": 3000,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 1357924,
	      "track_number": "WBALMBKS004",
	      "price": 1500,
	      "rid": "books-rid-001",
	      "name": "Harry Potter and the Philosopher's Stone",
	      "sale": 0,
	      "size": "0",
	      "total_price": 1500,
	      "nm_id": 2468135,
	      "brand": "J.K. Rowling",
	      "status": 202
	    },
	    {
	      "chrt_id": 8642091,
	      "track_number": "WBALMBKS004",
	      "price": 1500,
	      "rid": "books-rid-002",
	      "name": "1984 by George Orwell",
	      "sale": 0,
	      "size": "0",
	      "total_price": 1500,
	      "nm_id": 9753102,
	      "brand": "George Orwell",
	      "status": 202
	    }
	  ],
	  "locale": "kz",
	  "internal_signature": "sig-012",
	  "customer_id": "cust-aigerim",
	  "delivery_service": "kazpost",
	  "shardkey": "5",
	  "sm_id": 40,
	  "date_created": "2025-10-24T13:00:00Z",
	  "oof_shard": "5"
	}`)
	err = sc.Publish("orders", data5)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 5 with UID: %s\n", "SMjNghpbzYPuRRb4test")

	data6 := []byte(`{
	  "order_uid": "cHHKGKA1o5M4SWfRtest",
	  "track_number": "WBNSKTOYS005",
	  "entry": "WBNSK",
	  "delivery": {
	    "name": "Dmitry Kuznetsov",
	    "phone": "+73831234567",
	    "zip": "630000",
	    "city": "Novosibirsk",
	    "address": "Krasny Prospekt 17",
	    "region": "Novosibirsk Oblast",
	    "email": "dmitry@example.ru"
	  },
	  "payment": {
	    "transaction": "cHHKGKA1o5M4SWfRtest",
	    "request_id": "req-678",
	    "currency": "RUB",
	    "provider": "ozonpay",
	    "amount": 6000,
	    "payment_dt": 1729876800,
	    "bank": "ozonbank",
	    "delivery_cost": 800,
	    "goods_total": 5200,
	    "custom_fee": 0
	  },
	  "items": [
	    {
	      "chrt_id": 1122334,
	      "track_number": "WBNSKTOYS005",
	      "price": 6000,
	      "rid": "toys-rid-001",
	      "name": "LEGO Classic Creative Bricks",
	      "sale": 15,
	      "size": "0",
	      "total_price": 5100,
	      "nm_id": 4455667,
	      "brand": "LEGO",
	      "status": 202
	    }
	  ],
	  "locale": "ru",
	  "internal_signature": "sig-345",
	  "customer_id": "cust-dmitry",
	  "delivery_service": "sdek",
	  "shardkey": "6",
	  "sm_id": 50,
	  "date_created": "2025-10-24T14:00:00Z",
	  "oof_shard": "6"
	}`)
	err = sc.Publish("orders", data6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published order 6 with UID: %s\n", "cHHKGKA1o5M4SWfRtest")

	log.Println("All orders published to NATS")
}
