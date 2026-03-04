# Docker Build & CI/CD Testing Guide

## Проблема и решение

### Исправленные проблемы:
1. ✅ **ETXTBSY error** в frontend - исправлено через `npm ci --ignore-scripts`
2. ✅ **Obsolete version field** - удалено из docker-compose.yml
3. ✅ **Оптимизация сборки** - добавлены .dockerignore файлы

## Запуск Docker

### 1. Запустить Docker Desktop

```bash
# Проверить, что Docker запущен
docker info

# Если не запущен, запустите Docker Desktop вручную
open -a Docker
```

### 2. Очистить старые образы и контейнеры

```bash
cd /Users/artemcike/Documents/Imi-labs/Department-boilerplate

# Остановить и удалить все контейнеры
docker-compose down -v

# Удалить старые образы (опционально)
docker system prune -a
```

### 3. Собрать образы заново

```bash
# Собрать все образы
docker-compose build --no-cache

# Или собрать по отдельности для отладки
docker-compose build backend
docker-compose build frontend
```

### 4. Запустить сервисы

```bash
# Запустить все сервисы
docker-compose up -d

# Или запустить с логами для отладки
docker-compose up
```

### 5. Проверить статус

```bash
# Проверить запущенные контейнеры
docker-compose ps

# Проверить логи
docker-compose logs -f backend
docker-compose logs -f frontend
```

## Инициализация базы данных

```bash
# Дождаться запуска PostgreSQL (5-10 секунд)
sleep 10

# Инициализировать БД
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

## Проверка работоспособности

### Backend
```bash
# Health check
curl http://localhost:8080/api/v1/health

# Получить новости
curl http://localhost:8080/api/v1/news

# Метрики Prometheus
curl http://localhost:8080/metrics
```

### Frontend
```bash
# Открыть в браузере
open http://localhost:3000
```

### Мониторинг
```bash
# Prometheus
open http://localhost:9090

# Grafana (admin/admin)
open http://localhost:3001
```

## CI/CD тестирование

### Локальное тестирование GitHub Actions

```bash
# Установить act (если еще не установлен)
brew install act

# Запустить CI pipeline локально
act -j backend-test
act -j frontend-build
act -j docker-build
```

### Проверка линтеров

#### Backend (Go)
```bash
cd backend

# Установить golangci-lint
brew install golangci-lint

# Запустить линтер
golangci-lint run

# Форматирование
go fmt ./...
go vet ./...
```

#### Frontend (JavaScript)
```bash
cd frontend

# Установить зависимости
npm install

# Проверить сборку
npm run build

# Запустить dev сервер
npm run dev
```

## Отладка проблем

### Проблема: Docker daemon не запущен
```bash
# Решение: Запустить Docker Desktop
open -a Docker

# Подождать 10-20 секунд и проверить
docker info
```

### Проблема: ETXTBSY при npm install
```bash
# Решение: Уже исправлено в Dockerfile
# Используется: npm ci --ignore-scripts && npm rebuild esbuild
```

### Проблема: Backend не подключается к БД
```bash
# Проверить, что PostgreSQL запущен
docker-compose ps postgres

# Проверить логи PostgreSQL
docker-compose logs postgres

# Проверить переменные окружения
docker-compose exec backend env | grep DB_
```

### Проблема: Frontend не может подключиться к Backend
```bash
# Проверить, что backend запущен
curl http://localhost:8080/api/v1/health

# Проверить nginx конфигурацию
docker-compose exec frontend cat /etc/nginx/conf.d/default.conf

# Проверить логи nginx
docker-compose logs frontend
```

## Полный цикл тестирования

```bash
#!/bin/bash

echo "🚀 Запуск полного цикла тестирования..."

# 1. Очистка
echo "🧹 Очистка старых контейнеров..."
docker-compose down -v

# 2. Сборка
echo "🔨 Сборка образов..."
docker-compose build --no-cache

# 3. Запуск
echo "▶️  Запуск сервисов..."
docker-compose up -d

# 4. Ожидание
echo "⏳ Ожидание запуска сервисов..."
sleep 15

# 5. Инициализация БД
echo "💾 Инициализация базы данных..."
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql

# 6. Проверка
echo "✅ Проверка сервисов..."
curl -f http://localhost:8080/api/v1/health && echo "✅ Backend OK" || echo "❌ Backend FAIL"
curl -f http://localhost:3000 && echo "✅ Frontend OK" || echo "❌ Frontend FAIL"
curl -f http://localhost:9090/-/healthy && echo "✅ Prometheus OK" || echo "❌ Prometheus FAIL"
curl -f http://localhost:3001/api/health && echo "✅ Grafana OK" || echo "❌ Grafana FAIL"

echo "🎉 Тестирование завершено!"
```

Сохраните этот скрипт как `test.sh` и запустите:

```bash
chmod +x test.sh
./test.sh
```

## Мониторинг в реальном времени

```bash
# Следить за всеми логами
docker-compose logs -f

# Следить за конкретным сервисом
docker-compose logs -f backend

# Проверить использование ресурсов
docker stats
```

## Остановка и очистка

```bash
# Остановить сервисы
docker-compose down

# Остановить и удалить volumes
docker-compose down -v

# Полная очистка (осторожно!)
docker system prune -a --volumes
```
