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
Body (JSON, raw):
```
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
Body (JSON, raw):
```
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
