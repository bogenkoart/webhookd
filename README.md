# webhookd

Сервис гарантированной доставки вебхуков.

## Запуск

```bash
go run ./cmd/webhookd
```

## Структура

- `cmd/webhookd` — точка входа: собирает зависимости и запускает сервис.
- `internal/api` — HTTP: приём и проверка событий, ответы клиентам.
- `internal/config` — загрузка настроек.
- `internal/storage` — хранение событий и очереди доставки в Postgres.
- `internal/delivery` — доставка событий получателям и повторные попытки.

Зависимости: `main` → `api`, `config`, `storage`, `delivery`; `api` → `storage`; `delivery` → `storage`. Пакеты `storage` и `config` ничего из `internal` не импортируют.