# avito-shop


## Структура проекта

### Основные директории

- `cmd/` - точки входа в приложение
    - `app/` - инициализация и настройка HTTP-сервера
    - `config/` - конфигурация приложения
    - `initDB/` - инициализация базы данных и запуск миграций

- `internal/` - внутренняя логика приложения
    - `auth/` - сервисы для аутентификации
    - `controller/` - хендлеры
    - `mapper/` - преобразование в нужный тип inventory и transactions
    - `mocks/` - моки репозиториев для тестов
    - `model/` - бизнес-модели
    - `repository/` - репозитории и реализации с постгресом
    - `usecase/` - юзкейсы для транзакций, покупок, пользователя и авторизации 

- `tests/` - тесты
    - `unit_tests/` - модульные тесты
    - `e2e_tests/` - интеграционные тесты

- `migrations/` - SQL миграции
- `seed/` - начальные данные для таблицы с мерчом

### Docker файлы

- `docker-compose.yaml` - основной файл для запуска сервиса
- `docker-compose.test.yaml` - конфигурация для тестового окружения


## Запуск проекта

### Основное окружение

```bash
docker-compose up --build
```

### Запуск unit тестов
```bash
go test avito-shop/tests/unit_tests
```

### Запуск e2e тестов
```bash
docker-compose -f docker-compose.test.yaml up --build
```

### Запуск линтера
```bash
golangci-lint run ./...
```
