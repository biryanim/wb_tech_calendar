# Calendar
Простое HTTP-приложение для календаря на Go с использованием Gin. 
Реализует CRUD-операции для событий, поиск по дню, неделе, месяцу, 
а также кастомный мидлвар для логирования. 
Все события хранятся в памяти, архитектура проекта разделена по слоям.

## Структура проекта
```
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── calendar
│   │   │   ├── dto
│   │   │   │   └── dto.go
│   │   │   └── service.go
│   │   └── middleware
│   │       └── logger.go
│   ├── config
│   │   ├── config.go
│   │   └── http_config.go
│   ├── converter
│   │   └── converter.go
│   ├── model
│   │   └── event.go
│   └── service
│       ├── calendar
│       │   ├── service.go
│       │   └── service_test.go
│       └── service.go
└── Makefile
```

## API эндпоинты
| Метод | Путь              | Описание          |
| ----- | ----------------- | ----------------- |
| POST  | /create_event     | Создать событие   |
| POST  | /update_event     | Обновить событие  |
| POST  | /delete_event     | Удалить событие   |
| GET   | /events_for_day   | События на день   |
| GET   | /events_for_week  | События на неделю |
| GET   | /events_for_month | События на месяц  |

## Сборка 
`make build`