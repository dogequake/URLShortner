# URL Shortener

Минималистичный сервис сокращения ссылок на Go (чистая архитектура, in-memory).

## Возможности
- Быстрое сокращение длинных ссылок до коротких.
- Редирект по короткой ссылке на оригинальный URL.
- Чистая архитектура: легко заменить in-memory на Redis/PostgreSQL.
- Только стандартная библиотека Go.

## Запуск

1. Перейдите в папку проекта:
   ```
   cd /home/nyen/projects/URLShortner
   ```
2. Запустите сервер:
   ```
   go run main.go
   ```
   Сервер стартует на `:8080` и будет доступен по адресу http://147.45.187.208:8080

## Как пользоваться

### 1. Сократить ссылку

#### PowerShell (Windows 10/11)
```powershell
curl -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"url":"https://www.dns-shop.ru/profile/order/all/?page=1&orderKey=%D0%91-00683871"}' http://147.45.187.208:8080/shorten
```

#### Windows CMD (одна строка)
```
curl -X POST -H "Content-Type: application/json" -d "{\"url\":\"https://www.dns-shop.ru/profile/order/all/?page=1&orderKey=%D0%91-00683871\"}" http://147.45.187.208:8080/shorten
```

#### Linux/Mac
```
curl -X POST -H "Content-Type: application/json" \
  -d '{"url": "https://www.dns-shop.ru/profile/order/all/?page=1&orderKey=%D0%91-00683871"}' \
  http://147.45.187.208:8080/shorten
```

**Ответ:**
```
{"short_url":"http://147.45.187.208:8080/AbCdEf"}
```

### 2. Перейти по короткой ссылке
- Откройте короткую ссылку из ответа в браузере.
- Или выполните:
  ```
  curl -v http://147.45.187.208:8080/AbCdEf
  ```
  (где AbCdEf — ваш shortID из ответа)

## Пример

1. Сокращаем ссылку:
   ```
   curl -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"url":"https://www.dns-shop.ru/profile/order/all/?page=1&orderKey=%D0%91-00683871"}' http://147.45.187.208:8080/shorten
   ```
   Ответ:
   ```
   {"short_url":"http://147.45.187.208:8080/FvZr9Y"}
   ```
2. Открываем короткую ссылку:
   - В браузере: http://147.45.187.208:8080/FvZr9Y
   - Или через curl:
     ```
     curl -v http://147.45.187.208:8080/FvZr9Y
     ```

## Структура проекта
- `main.go` — запуск сервера
- `handler/` — HTTP-обработчики
- `service/` — бизнес-логика
- `repository/` — in-memory map
- `model/` — структуры данных

## Заметки
- Для хранения используется map (in-memory).
- Легко заменить хранилище на Redis/PostgreSQL.
- Используется только стандартная библиотека.
