### Тестирование  API через Postman:
Запустите сервис: go run cmd/server/main.go и убедитесь по логам в терминале, что сервис запустился 
```
[postgres] connect to DB successfully
[app] Connected to Postgres successfully
[app]storage initialized successfully
[app] Service initialized successfully
[app] starting server on :7575
```

Вбейте в адресной строке Postman  адрес в соответствии с Вашим cfg (например http://localhost:7575) (при методе GET нажмите SEND - в поле body будет отображен html код стартовой страницы, это подтверждает удачное подключение к сервису)

1. POST /items — создание транзакции
```
Body (JSON, raw):
{
  "name": "Покупка ноутбука",
  "description": "Для работы",
  "amount": 75000.50,
  "type": "expense",
  "category": "electronics",
  "event_date": "2025-11-26"
}
```
Ожидаемый результат:
201 Created
{"id": <new_id>}

2. PUT /items/:id — обновление транзакции
```
Body (JSON, raw):
{
  "name": "Покупка ноутбука",
  "description": "Для работы и учебы",
  "amount": 76000,
  "type": "expense",
  "category": "electronics",
  "event_date": "2025-11-27"
}
```
Ожидаемый результат:
204 No Content

3. GET /analytics — общая аналитика
Возвращает суммарные значения по типам транзакций (income / expense) или другие показатели, если они предусмотрены в сервисе.
Пример запроса:
GET http://localhost:7575/analytics
```
Ожидаемый результат: Суммы уже в рублях.
{
  "income": 50000,
  "expense": 76000
}
```
4. GET /analytics/day — группировка по дням
```
Query params:
from — необязательный, начало периода, формат YYYY-MM-DD
to — необязательный, конец периода, формат YYYY-MM-DD
Пример запроса:
GET http://localhost:7575/analytics/day?from=2025-11-01&to=2025-11-30
Ожидаемый результат:
{
  "2025-11-26": 76000,
  "2025-11-25": 50000
}
```

5. GET /analytics/week — группировка по неделям
```
GET http://localhost:7575/analytics/week?from=2025-11-01&to=2025-11-30

Пример результата:
{
  "2025-W48": 126000
}
```
6. GET /analytics/month — группировка по месяцам
```
GET http://localhost:7575/analytics/month?from=2025-01-01&to=2025-12-31

Пример результата:
{
  "2025-11": 126000
}
```
7. GET /analytics/category — группировка по категориям
```
GET http://localhost:7575/analytics/category?from=2025-11-01&to=2025-11-30

Пример результата:
{
  "electronics": 76000,
  "food": 50000
}
```