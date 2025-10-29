package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service-demo/pkg/cache"

	"github.com/gorilla/mux"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cache.CacheMu.RLock()
	data, ok := cache.Cache[id]
	cache.CacheMu.RUnlock()

	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, data)
}

func UIHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Order Viewer</title>
		<style>
			* { margin: 0; padding: 0; box-sizing: border-box; }
			
			body {
				font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
				background: linear-gradient(135deg, #8a9be6ff 0%, #4b86a2ff 100%);
				min-height: 100vh;
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 20px;
			}ы
			
			.container {
				width: 100%;
				max-width: 500px;
			}
			
			.search-card {
				background: white;
				border-radius: 20px;
				padding: 40px;
				box-shadow: 0 20px 60px rgba(0,0,0,0.3);
				animation: fadeIn 0.5s ease;
			}
			
			@keyframes fadeIn {
				from { opacity: 0; transform: translateY(20px); }
				to { opacity: 1; transform: translateY(0); }
			}
			
			h1 {
				font-size: 28px;
				font-weight: 700;
				color: #1a202c;
				margin-bottom: 10px;
				text-align: center;
			}
			
			.subtitle {
				color: #718096;
				text-align: center;
				margin-bottom: 30px;
				font-size: 14px;
			}
			
			.input-group {
				margin-bottom: 20px;
			}
			
			label {
				display: block;
				font-size: 14px;
				font-weight: 600;
				color: #4a5568;
				margin-bottom: 8px;
			}
			
			input[type="text"] {
				width: 100%;
				padding: 14px 18px;
				font-size: 15px;
				border: 2px solid #e2e8f0;
				border-radius: 12px;
				transition: all 0.3s ease;
				font-family: 'Courier New', monospace;
			}
			
			input[type="text"]:focus {
				outline: none;
				border-color: #667eea;
				box-shadow: 0 0 0 3px rgba(102,126,234,0.1);
			}
			
			button {
				width: 100%;
				padding: 14px;
				background: linear-gradient(135deg, #8a9be6ff 0%, #4b86a2ff 100%);
				color: white;
				border: none;
				border-radius: 12px;
				font-size: 16px;
				font-weight: 600;
				cursor: pointer;
				transition: all 0.3s ease;
				box-shadow: 0 4px 15px rgba(102,126,234,0.4);
			}
			
			button:hover {
				transform: translateY(-2px);
				box-shadow: 0 6px 20px rgba(102,126,234,0.6);
			}
			
			button:active {
				transform: translateY(0);
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="search-card">
				<h1>Order Viewer</h1>
				<p class="subtitle">Введите ID заказа для просмотра деталей</p>
				<form method="POST" action="/ui">
					<div class="input-group">
						<label for="id">Order ID</label>
						<input type="text" id="id" name="id" placeholder="b563feb7b2b84b6test" required>
					</div>
					<button type="submit">Найти заказ</button>
				</form>
			</div>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func GetOrderUI(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	cache.CacheMu.RLock()
	data, ok := cache.Cache[id]
	cache.CacheMu.RUnlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Order Not Found</title>
			<style>
				* { margin: 0; padding: 0; box-sizing: border-box; }
				body {
					font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
					background: linear-gradient(135deg, #8a9be6ff 0%, #4b86a2ff 100%);
					min-height: 100vh;
					display: flex;
					align-items: center;
					justify-content: center;
					padding: 20px;
				}
				.error-card {
					background: white;
					border-radius: 20px;
					padding: 40px;
					text-align: center;
					box-shadow: 0 20px 60px rgba(0,0,0,0.3);
					max-width: 400px;
				}
				.error-icon {
					font-size: 60px;
					margin-bottom: 20px;
				}
				h2 {
					color: #1a202c;
					margin-bottom: 10px;
				}
				p {
					color: #718096;
					margin-bottom: 30px;
				}
				a {
					display: inline-block;
					padding: 12px 30px;
					background: linear-gradient(135deg, #8a9be6ff 0%, #4b86a2ff 100%);
					color: white;
					text-decoration: none;
					border-radius: 12px;
					font-weight: 600;
					transition: all 0.3s ease;
				}
				a:hover {
					transform: translateY(-2px);
					box-shadow: 0 6px 20px rgba(102,126,234,0.6);
				}
			</style>
		</head>
		<body>
			<div class="error-card">
				<div class="error-icon">❌</div>
				<h2>Заказ не найден</h2>
				<p>ID: ` + id + `</p>
				<a href="/ui">← Вернуться к поиску</a>
			</div>
		</body>
		</html>
		`
		fmt.Fprint(w, html)
		return
	}

	var order map[string]interface{}
	json.Unmarshal([]byte(data), &order)

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Order Details</title>
		<style>
			* { margin: 0; padding: 0; box-sizing: border-box; }
			
			body {
				font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
				background: linear-gradient(135deg, #8a9be6ff 0%, #4b86a2ff 100%);
				min-height: 100vh;
				padding: 40px 20px;
			}
			
			.container {
				max-width: 900px;
				margin: 0 auto;
			}
			
			.header {
				background: white;
				border-radius: 20px;
				padding: 30px;
				margin-bottom: 20px;
				box-shadow: 0 10px 40px rgba(0,0,0,0.2);
			}
			
			.header h1 {
				font-size: 24px;
				color: #1a202c;
				margin-bottom: 5px;
			}
			
			.order-id {
				font-family: 'Courier New', monospace;
				color: #667eea;
				font-size: 14px;
			}
			
			.back-btn {
				display: inline-block;
				margin-top: 15px;
				padding: 10px 20px;
				background: #f7fafc;
				color: #4a5568;
				text-decoration: none;
				border-radius: 10px;
				font-size: 14px;
				font-weight: 600;
				transition: all 0.3s ease;
			}
			
			.back-btn:hover {
				background: #edf2f7;
				transform: translateX(-3px);
			}
			
			.section {
				background: white;
				border-radius: 20px;
				padding: 30px;
				margin-bottom: 20px;
				box-shadow: 0 10px 40px rgba(0,0,0,0.2);
			}
			
			.section-title {
				font-size: 18px;
				font-weight: 700;
				color: #1a202c;
				margin-bottom: 20px;
				padding-bottom: 10px;
				border-bottom: 2px solid #f7fafc;
			}
			
			.info-grid {
				display: grid;
				grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
				gap: 15px;
			}
			
			.info-item {
				padding: 15px;
				background: #f7fafc;
				border-radius: 10px;
				transition: all 0.3s ease;
			}
			
			.info-item:hover {
				background: #edf2f7;
				transform: translateY(-2px);
			}
			
			.info-label {
				font-size: 12px;
				text-transform: uppercase;
				color: #718096;
				font-weight: 600;
				margin-bottom: 5px;
				letter-spacing: 0.5px;
			}
			
			.info-value {
				font-size: 15px;
				color: #1a202c;
				font-weight: 500;
			}
			
			.items-table {
				width: 100%;
				border-collapse: collapse;
			}
			
			.items-table thead {
				background: #f7fafc;
			}
			
			.items-table th {
				padding: 12px;
				text-align: left;
				font-size: 12px;
				text-transform: uppercase;
				color: #718096;
				font-weight: 600;
				letter-spacing: 0.5px;
			}
			
			.items-table td {
				padding: 15px 12px;
				border-top: 1px solid #f7fafc;
				color: #1a202c;
			}
			
			.items-table tr:hover {
				background: #f7fafc;
			}
			
			.price {
				color: #667eea;
				font-weight: 600;
			}
			
			.status {
				display: inline-block;
				padding: 4px 12px;
				background: #c6f6d5;
				color: #22543d;
				border-radius: 20px;
				font-size: 12px;
				font-weight: 600;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>📦 Детали заказа</h1>
				<div class="order-id">` + id + `</div>
				<a href="/ui" class="back-btn">← Назад к поиску</a>
			</div>
			
			<div class="section">
				<div class="section-title">Общая информация</div>
				<div class="info-grid">
					<div class="info-item">
						<div class="info-label">Track Number</div>
						<div class="info-value">` + fmt.Sprint(order["track_number"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Entry</div>
						<div class="info-value">` + fmt.Sprint(order["entry"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Locale</div>
						<div class="info-value">` + fmt.Sprint(order["locale"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Customer ID</div>
						<div class="info-value">` + fmt.Sprint(order["customer_id"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Delivery Service</div>
						<div class="info-value">` + fmt.Sprint(order["delivery_service"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Date Created</div>
						<div class="info-value">` + fmt.Sprint(order["date_created"]) + `</div>
					</div>
				</div>
			</div>`

	delivery := order["delivery"].(map[string]interface{})
	html += `
			<div class="section">
				<div class="section-title">Доставка</div>
				<div class="info-grid">
					<div class="info-item">
						<div class="info-label">Имя</div>
						<div class="info-value">` + fmt.Sprint(delivery["name"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Телефон</div>
						<div class="info-value">` + fmt.Sprint(delivery["phone"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Email</div>
						<div class="info-value">` + fmt.Sprint(delivery["email"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Город</div>
						<div class="info-value">` + fmt.Sprint(delivery["city"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Адрес</div>
						<div class="info-value">` + fmt.Sprint(delivery["address"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Регион</div>
						<div class="info-value">` + fmt.Sprint(delivery["region"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">ZIP</div>
						<div class="info-value">` + fmt.Sprint(delivery["zip"]) + `</div>
					</div>
				</div>
			</div>`

	payment := order["payment"].(map[string]interface{})
	html += `
			<div class="section">
				<div class="section-title">Оплата</div>
				<div class="info-grid">
					<div class="info-item">
						<div class="info-label">Transaction</div>
						<div class="info-value">` + fmt.Sprint(payment["transaction"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Currency</div>
						<div class="info-value">` + fmt.Sprint(payment["currency"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Provider</div>
						<div class="info-value">` + fmt.Sprint(payment["provider"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Bank</div>
						<div class="info-value">` + fmt.Sprint(payment["bank"]) + `</div>
					</div>
					<div class="info-item">
						<div class="info-label">Сумма</div>
						<div class="info-value price">` + fmt.Sprint(payment["amount"]) + ` ₽</div>
					</div>
					<div class="info-item">
						<div class="info-label">Доставка</div>
						<div class="info-value price">` + fmt.Sprint(payment["delivery_cost"]) + ` ₽</div>
					</div>
					<div class="info-item">
						<div class="info-label">Товары</div>
						<div class="info-value price">` + fmt.Sprint(payment["goods_total"]) + ` ₽</div>
					</div>
				</div>
			</div>
			
			<div class="section">
				<div class="section-title">Товары</div>
				<table class="items-table">
					<thead>
						<tr>
							<th>Название</th>
							<th>Бренд</th>
							<th>Цена</th>
							<th>Скидка</th>
							<th>Итого</th>
							<th>Статус</th>
						</tr>
					</thead>
					<tbody>`

	items := order["items"].([]interface{})
	for _, item := range items {
		i := item.(map[string]interface{})
		html += `
						<tr>
							<td>` + fmt.Sprint(i["name"]) + `</td>
							<td>` + fmt.Sprint(i["brand"]) + `</td>
							<td class="price">` + fmt.Sprint(i["price"]) + ` ₽</td>
							<td>` + fmt.Sprint(i["sale"]) + `%</td>
							<td class="price">` + fmt.Sprint(i["total_price"]) + ` ₽</td>
							<td><span class="status">` + fmt.Sprint(i["status"]) + `</span></td>
						</tr>`
	}

	html += `
					</tbody>
				</table>
			</div>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}
