# Lesson 2 – Mood Tracker с авторизацией (JWT)

## Что реализовано:
- Регистрация и авторизация пользователей (JWT)
- Привязка настроений к пользователю
- Защита маршрутов через middleware
- CRUD для настроений:
  - `GET /moods` – получить список своих настроений
  - `GET /moods/:id` – получить конкретное настроение
  - `POST /moods` – создать настроение
  - `DELETE /moods/:id` – удалить настроение

---

## Эндпоинты регистрации и входа:
- `POST /register` – регистрация нового пользователя
- `POST /login` – вход и получение JWT токена

📌 Все маршруты `/moods` требуют заголовок: Authorization: Bearer <токен>

---

## Как запустить:

1. Убедись, что в PostgreSQL существует схема `mood`
2. Настрой подключение к БД в `db/db.go` (DSN строка)
3. Перейди в папку `Lesson2`
4. Установи зависимости:
    go get -u github.com/gofiber/fiber/v2 github.com/golang-jwt/jwt/v5 golang.org/x/crypto gorm.io/driver/postgres gorm.io/gorm
5. Запусти сервер:
    go run main.go