# Департамент финансов города Якутска

Современный веб-портал департамента финансов с мониторингом и аналитикой.

## Технологический стек

### Backend
- **Go** - REST API сервер
- **Prometheus** - сбор метрик
- **PostgreSQL** - база данных

### Frontend
- **Svelte** - современный UI фреймворк
- **Vite** - сборщик

### Мониторинг
- **Prometheus** - сбор и хранение метрик
- **Grafana** - визуализация метрик

## Функционал

- 📄 Информационные страницы о департаменте
- 📰 Новости и объявления
- 💰 Бюджетные данные и отчеты
- 📋 Документы и нормативные акты
- 📊 Мониторинг производительности

## Метрики

- HTTP метрики (запросы, статусы, время ответа)
- Системные метрики (CPU, память, диск)
- Бизнес-метрики (посещения, популярные страницы)
- Database метрики

## Быстрый старт

```bash
# Запустить все сервисы
make up

# Или через docker-compose
docker-compose up -d
```

После запуска доступны:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3001 (admin/admin)

## Разработка

### Backend
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Только мониторинг
```bash
make dev
```

## Инициализация базы данных

```bash
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

## Полезные команды

```bash
make help    # Показать все команды
make build   # Собрать Docker образы
make down    # Остановить сервисы
make logs    # Просмотр логов
make clean   # Очистить volumes
```
