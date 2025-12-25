# Тестовое задание для Samokat.
# Описание задачи:
Создайте HTTP-сервер на Go, который будет обрабатывать следующие эндпоинты:


POST /todos — создать новую задачу


GET /todos — получить список всех задач


GET /todos/{id} — получить задачу по идентификатору


PUT /todos/{id} — обновить задачу по идентификатору


DELETE /todos/{id} — удалить задачу по идентификатору


# Инструкция по запуску
**P.S Все секретки были выложены в открытый доступ для более удобного тестирования задания, иначе бы попали в .gitignore**


**1. Скачать [ZIP-архив](https://github.com/northerf/EcomTest/archive/refs/heads/main.zip) или клонировать [репозиторий](https://github.com/northerf/EcomTest.git)**

**2. Убедиться в наличие у вас установленного [Docker](https://www.docker.com/), Golang**


**3. Ввести в консоль "docker-compose up"**

**ВАЖНОЕ замечание: проверьте порты перед запуском и освободите 8080**


**4. Можете начинать пользоваться!**

## Как проверить, все ли успешно?

### Если написано:Starting on port 8080... - тогда все отлично!

# **Как дергать ручки в моём сервисе?**

### Чтобы дергать ручки вам понадобиться Postman или его аналог(я использовал Postman)

### 1. POST /todos — создать новую задачу

**POST http://localhost:8080/todos**

Headers: Content-Type: application/json 

Body:
```json
{
    "title": "test",
    "description": "описание",
    "completed": false
}
```

Ответ: 201 Created
```json
{
    "id": 1,
    "title": "test",
    "description": "описание",
    "completed": false
}
```

**В случае неверного тела запроса, ответ будет таким:**
Body:
```json
{
    "id": "1",
    "title": "test",
    "description": "описание",
    "completed": false
}
```

Ответ: 400 Bad Request
```json
{
  "error": "Invalid JSON"
}
```

**Или если, тело все же правильное (значение ID числовое), но оно в моем задании не требуется, т.к. ID автоматически ставится системой дабы избежать дубликатов:**

Body:
```json
{
    "id": 1,
    "title": "test",
    "description": "описание",
    "completed": false
}
```

Ответ: 400 Bad Request
```json
{
  "error": "ID is not required"
}
```

### 2. GET /todos — получить список всех задач

**GET http://localhost:8080/todos**

Ответ: 200 OK
```json
[
    {
        "id": 1,
        "title": "test",
        "description": "Пробелы должны игнорироваться",
        "completed": false
    }
]
```

### 3. GET /todos/{id} — получить задачу по идентификатору

**GET http://localhost:8080/todos/{id}**

Ответ: 200 OK

```json
{
    "id": 1,
    "title": "test",
    "description": "Пробелы должны игнорироваться",
    "completed": false
}
```
**В случае неверного ID, ответ будет таким:**

Ответ: 404 Not Found
```json
{
    "error": "Todo not found"
}
```

### 4. PUT /todos/{id} — обновить задачу по идентификатору

**PUT http://localhost:8080/todos/{id}**

Body:
```json
{
    "id": 1,
    "title": "test123",
    "description": "123",
    "completed": true
}
```
Ответ: 200 OK

```json
{
    "id": 1,
    "title": "test123",
    "description": "123",
    "completed": true
}
```
**Опять же будет обрабатываться ошибка, если не будет найден ID**

### 5. DELETE /todos/{id} — удалить задачу по идентификатору

**DELETE http://localhost:8080/todos/{id}**

Ответ: 204 No Content

## Тесты

### Чтобы запустить тесты необходимо в терминале ввести команду "go test ./... "
Тесты:

[Тесты для хендлеров](https://github.com/northerf/EcomTest/blob/main/internal/handler/handler_test.go)


[Тесты для репозитория](https://github.com/northerf/EcomTest/blob/main/internal/repository/in-memory_test.go)
