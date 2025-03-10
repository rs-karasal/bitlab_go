# Lesson 1 – Базовый Mood Tracker

### Что реализовано:
- Подключение к PostgreSQL
- Модель `Mood` с полями `Value` и `Note`
- CRUD эндпоинты:
  - `GET /moods` – получить список всех настроений
  - `POST /moods` – создать настроение
  - `DELETE /moods/:id` – удалить настроение

### Как запустить:
1. Настрой `.env` или просто подставь данные в `db.Connect()`
2. Выполни `go run main.go`
3. Протестируй через Postman / Insomnia