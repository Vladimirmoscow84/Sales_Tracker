# Sales Tracker

ТЗ на проект находится в файле TASK.md

## Sales Tracker — сервис учёта доходов и расходов с аналитикой, фильтрами, сортировкой и визуализацией данных.

Приложение позволяет создавать финансовые записи, редактировать, удалять их, просматривать статистику, строить графики и анализировать финансовую динамику по дням, неделям и месяцам.

## Возможности
### Работа с транзакциями
```
Создание записи (POST /items)
Обновление записи (PUT /items/:id)
Удаление записи (DELETE /items/:id)
Получение всех записей (GET /items)
Получение записи по ID (GET /items/:id)
Получение записей за период (GET /items/period)
Получение отсортированных записей (GET /items/sorted)
```

### Аналитика
```
Общая статистика:
  сумма
  среднее значение
  количество
  медиана
  90-й перцентиль
Группировка:
  по дням (GET /analytics/day)
  по неделям (GET /analytics/week)
  по месяцам (GET /analytics/month)
  по категориям (GET /analytics/category)

Финансовая динамика:
  доходы
  расходы
  ```


## Фронтенд (HTML + JS + Chart.js)
![Скриншот интерфейса](assets/вэб.png)
```
Фильтрация по дате
Выбор периода анализа (день, неделя, месяц)
распределение по категориям
динамика финансов
сравнение доходов и расходов
Красивый и лёгкий интерфейс без фреймворков
```
## Технологии
```
Go (Gin, sqlx)
PostgreSQL (хранение транзакций)
HTML + Vanilla JS
Chart.js (визуализация данных)
CSS (кастомная цветовая схема)
```

## Структура проекта
```
/cmd/server              - точка входа сервера
/internal/app            - инициализация сущностей и запуск сервера
/internal/handlers       - HTTP-эндпоинты и хэндлеры
/internal/handlers/dto            - DTO для запросов/ответов
/internal/model          - модели данных
/internal/service        - бизнес-логика (CRUD + аналитика)
/internal/storage        - PostgreSQL слой (CRUD + аналитика)
/internal/middleware     - middleware Gin (логирование, CORS, recovery)
internal/cfg             - загрузка конфигурации (env, defaults)

/db/dumps                - SQL-миграции и дампы БД
/web                     - фронтенд (HTML, JS, CSS)
assets                   - изображения, скриншоты и статические файлы

TASK.md                  - Техническое задание
TEST_POSTMAN.md          - полный набор тестов Postman для проверки API
README.md                - документация проекта
```

## Быстрый старт
1. Создайте таблицу с помощью файла миграции db/dumps/ ....up
или с помощью SQL команды в теминале  SQL
```
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    amount BIGINT NOT NULL,          -- в копейках
    type TEXT NOT NULL,              -- income / expense
    category TEXT,
    event_date TIMESTAMP NOT NULL
);
```
2. Настроить подключение к БД
Через переменные окружения или .env
Пример:
```
DATABASE_URI="host=localhost port=5440 user=vladimir password=password dbname=sales_tracker sslmode=disable"
SERVER_ADDRESS=":7575"
```
3. Запустите сервер
go run cmd/server/main.go

4. Откройте интерфейс
http://localhost: (ваш порт)

5. Начинайте работу

Кейс тестов APIнаходится в файле TEST_POSTMAN.md

# Sales Tracker — это демонстрация чистой архитектуры на Go:
```
раздельные слои
строгие DTO
SQL-запросы через sqlx
фронтенд без фреймворков
визуализация данных в реальном времени
Приложение подходит как учебный пример и как основа для реального внутреннего финансового инструмента.
```
Автор

Разработчик: Vladimirmoscow84
Контакт: ccr1@yandex.ru

GitHub: github.com/Vladimirmoscow84