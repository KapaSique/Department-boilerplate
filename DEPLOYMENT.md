# Руководство по развертыванию

## Требования

- Docker и Docker Compose
- Go 1.22+ (для локальной разработки)
- Node.js 20+ (для локальной разработки)

## Быстрый старт

### 1. Клонировать репозиторий

```bash
git clone https://github.com/KapaSique/Department-boilerplate.git
cd Department-boilerplate
```

### 2. Запустить все сервисы

```bash
make up
```

Или через docker-compose:

```bash
docker-compose up -d
```

### 3. Инициализировать базу данных

```bash
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

### 4. Открыть приложение

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3001 (admin/admin)

## Локальная разработка

### Backend

```bash
cd backend
go mod download
go run cmd/server/main.go
```

Backend будет доступен на http://localhost:8080

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend будет доступен на http://localhost:5173

### Только мониторинг

Для запуска только PostgreSQL, Prometheus и Grafana:

```bash
make dev
```

## Структура проекта

```
.
├── backend/
│   ├── cmd/server/          # Точка входа приложения
│   ├── internal/
│   │   ├── handlers/        # HTTP обработчики
│   │   ├── models/          # Модели данных
│   │   ├── database/        # Работа с БД
│   │   └── metrics/         # Prometheus метрики
│   ├── init.sql             # Инициализация БД
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/      # Компоненты (Header, Footer)
│   │   ├── routes/          # Страницы приложения
│   │   ├── lib/             # API клиент
│   │   └── styles/          # Глобальные стили
│   └── Dockerfile
├── monitoring/
│   ├── prometheus/          # Конфигурация Prometheus
│   └── grafana/             # Дашборды и datasources
├── docker-compose.yml       # Оркестрация сервисов
└── Makefile                 # Команды для разработки
```

## API Endpoints

### Новости
- `GET /api/v1/news` - Список новостей
- `GET /api/v1/news/:id` - Новость по ID

### Документы
- `GET /api/v1/documents` - Список документов
- `GET /api/v1/documents/:id` - Документ по ID

### Бюджет
- `GET /api/v1/budget` - Бюджетные данные
- `GET /api/v1/budget/reports` - Отчеты

### Информация
- `GET /api/v1/department/info` - О департаменте
- `GET /api/v1/department/contacts` - Контакты

### Служебные
- `GET /api/v1/health` - Проверка здоровья
- `GET /metrics` - Метрики Prometheus

## Метрики Prometheus

### HTTP метрики
- `http_requests_total` - Общее количество запросов
- `http_request_duration_seconds` - Длительность запросов

### Бизнес-метрики
- `page_views_total` - Просмотры страниц
- `news_views_total` - Просмотры новостей
- `document_downloads_total` - Скачивания документов
- `budget_report_views_total` - Просмотры бюджетных отчетов

### Database метрики
- `db_queries_total` - Количество запросов к БД
- `db_query_duration_seconds` - Длительность запросов к БД

## Grafana Dashboard

После запуска Grafana доступна по адресу http://localhost:3001

**Логин**: admin
**Пароль**: admin

Дашборд "Finance Department Metrics" включает:
- HTTP запросы в секунду
- Длительность HTTP запросов (p95)
- HTTP статус коды
- Просмотры страниц
- Просмотры новостей
- Скачивания документов
- Просмотры бюджетных отчетов
- Длительность запросов к БД
- Запросы к БД в секунду

## Остановка сервисов

```bash
make down
```

Или:

```bash
docker-compose down
```

## Очистка данных

Для полной очистки volumes и образов:

```bash
make clean
```

## Troubleshooting

### Backend не запускается

Проверьте, что PostgreSQL запущен и доступен:

```bash
docker-compose logs postgres
```

### Frontend не может подключиться к API

Убедитесь, что backend запущен на порту 8080:

```bash
curl http://localhost:8080/api/v1/health
```

### Prometheus не собирает метрики

Проверьте targets в Prometheus UI:
http://localhost:9090/targets

### Grafana не показывает данные

Убедитесь, что datasource настроен правильно:
http://localhost:3001/datasources
