# Order Service Demo

Сервис для управления заказами с кэшированием, интеграцией NATS Streaming и PostgreSQL.

## Описание

Это демонстрационный сервис, который:
- Получает данные о заказах из NATS Streaming
- Сохраняет их в PostgreSQL
- Кэширует в памяти для быстрого доступа
- Предоставляет веб-интерфейс для просмотра заказов по id
- Восстанавливает кэш из БД при перезапуске

## Технологии

- **Go 1.25.3 - основной язык
- **PostgreSQL** - хранение данных
- **NATS Streaming** - очередь сообщений
- **Gorilla Mux** - HTTP роутер


## Структура проекта

```
order-service-demo/
├── cmd/
│   ├── main/main.go           # Основной сервис
│   └── publish/main.go        # Публикация в NATS
├── internal/
│   ├── config/                # Конфигурация
│   ├── db/                    # Работа с БД
│   ├── handler/               # HTTP handlers
│   ├── models/                # Модели данных
│   └── service/               # NATS сервис
├── pkg/
│   └── cache/                 # In-memory кэш
└── scripts/                   # Скрипты тестирования
```


## Установка


### Проект

```bash
# Склонировать репозиторий
git clone <your-repo-url>
cd order-service-demo

# Установить зависимости
go mod download
```

## Запуск

### Шаг 1: Установить переменную окружения

```bash
export DB_CONN="postgres://postgres:postgres@localhost:5432/orders?sslmode=disable"
```

### Шаг 2: Запустить NATS Streaming (Терминал 1)

```bash
nats-streaming-server -cluster_id test-cluster
```

Должно появиться: `Starting nats-streaming-server...`

### Шаг 3: Запустить сервис (Терминал 2)

```bash
go run ./cmd/main/main.go
```


### Шаг 4: Опубликовать тестовые данные (Терминал 3)

```bash
go run ./cmd/publish/main.go
```


### Шаг 5: Проверить работу

Открыть в браузере: **http://localhost:8080/ui**



### Стресс-тест (опционально)

```bash

# Запустить тест
bash ./scripts/wrk_test.sh   

```

---

**Видео демонстрация:** [Ссылка на видео]