# webhookd

Сервис гарантированной доставки вебхуков.

## Запуск

```bash
go run ./cmd/webhookd
```

## Пример проверки

```bash
curl -i localhost:8080/healthz
```

Ожидаемый результат:

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 10 Oct 2026 11:55:34 GMT
Content-Length: 15

{"status":"ok"}


## Структура

- `cmd/webhookd` — точка входа: собирает зависимости и запускает сервис.
- `internal/api` — HTTP: приём и проверка событий, ответы клиентам.
- `internal/config` — загрузка настроек.
- `internal/storage` — хранение событий и очереди доставки в Postgres.
- `internal/delivery` — доставка событий получателям и повторные попытки.

Зависимости: `main` → `api`, `config`, `storage`, `delivery`; `api` → `storage`; `delivery` → `storage`. Пакеты `storage` и `config` ничего из `internal` не импортируют.