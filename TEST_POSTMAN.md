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


