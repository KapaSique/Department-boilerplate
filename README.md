<div align="center">

# 🏛️ Департамент финансов города Якутска

### Современный веб-портал с мониторингом и аналитикой

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Svelte](https://img.shields.io/badge/Svelte-4.0+-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)](https://svelte.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Prometheus](https://img.shields.io/badge/Prometheus-Metrics-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)](https://prometheus.io/)
[![Grafana](https://img.shields.io/badge/Grafana-Dashboard-F46800?style=for-the-badge&logo=grafana&logoColor=white)](https://grafana.com/)

[🚀 Быстрый старт](#-быстрый-старт) • [📖 Документация](#-документация) • [🎨 Дизайн](#-дизайн-система) • [📊 Метрики](#-метрики)

</div>

---

## 📋 О проекте

Полнофункциональный веб-портал департамента финансов с интегрированной системой мониторинга производительности и бизнес-аналитики.

### ✨ Основные возможности

- 📰 **Новости и объявления** — актуальная информация для граждан
- 💰 **Бюджетные данные** — прозрачность финансовых отчетов
- 📋 **Документы** — нормативные акты и официальные документы
- 📊 **Мониторинг** — real-time метрики производительности
- 🎯 **Аналитика** — визуализация данных в Grafana
- ♿ **Доступность** — WCAG AAA совместимость

---

## 🛠️ Технологический стек

<div align="center">

### Backend
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Gorilla Mux](https://img.shields.io/badge/Gorilla_Mux-00ADD8?style=for-the-badge&logo=go&logoColor=white)

### Frontend
![Svelte](https://img.shields.io/badge/Svelte-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

### Мониторинг
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)
![Grafana](https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white)

### DevOps
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Docker Compose](https://img.shields.io/badge/Docker_Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Nginx](https://img.shields.io/badge/Nginx-009639?style=for-the-badge&logo=nginx&logoColor=white)

</div>

---

## 🚀 Быстрый старт

### Автоматический запуск (рекомендуется)

```bash
# Клонировать репозиторий
git clone https://github.com/KapaSique/Department-boilerplate.git
cd Department-boilerplate

# Запустить Docker Desktop, затем:
./test.sh
```

Скрипт автоматически соберет, запустит и протестирует все сервисы.

### Ручной запуск

```bash
# Клонировать
git clone https://github.com/KapaSique/Department-boilerplate.git
cd Department-boilerplate

# Запустить
make up

# Инициализировать БД
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

### 🌐 Доступ к сервисам

| Сервис | URL | Описание |
|--------|-----|----------|
| 🎨 **Frontend** | http://localhost:3000 | Веб-интерфейс портала |
| 🔧 **Backend API** | http://localhost:8080 | REST API сервер |
| 📊 **Prometheus** | http://localhost:9090 | Система мониторинга |
| 📈 **Grafana** | http://localhost:3001 | Дашборды (admin/admin) |

---

## 📖 Документация

### API Endpoints

#### 📰 Новости
```http
GET /api/v1/news              # Список новостей
GET /api/v1/news/:id          # Новость по ID
```

#### 📋 Документы
```http
GET /api/v1/documents         # Список документов
GET /api/v1/documents/:id     # Документ по ID
```

#### 💰 Бюджет
```http
GET /api/v1/budget            # Бюджетные данные
GET /api/v1/budget/reports    # Финансовые отчеты
```

#### ℹ️ Информация
```http
GET /api/v1/department/info     # О департаменте
GET /api/v1/department/contacts # Контакты
GET /api/v1/health              # Health check
```

#### 📊 Метрики
```http
GET /metrics                  # Prometheus метрики
```

---

## 📊 Метрики

### HTTP Метрики
- `http_requests_total` — общее количество HTTP запросов
- `http_request_duration_seconds` — длительность обработки запросов

### Бизнес-метрики
- `page_views_total` — просмотры страниц
- `news_views_total` — просмотры новостей
- `document_downloads_total` — скачивания документов
- `budget_report_views_total` — просмотры бюджетных отчетов

### Database метрики
- `db_queries_total` — количество запросов к БД
- `db_query_duration_seconds` — длительность запросов к БД

---

## 🎨 Дизайн-система

### Типографика
- **Заголовки**: Lexend (600-700)
- **Основной текст**: Source Sans 3 (400-500)
- **Минимальный размер**: 16px

### Цветовая палитра

```css
--primary:    #0891B2  /* Cyan 600 */
--secondary:  #22D3EE  /* Cyan 400 */
--cta:        #22C55E  /* Green 500 */
--background: #ECFEFF  /* Cyan 50 */
--text:       #164E63  /* Cyan 900 */
```

### Доступность
- ✅ WCAG AAA совместимость
- ✅ Контраст текста 4.5:1+
- ✅ Видимые focus states (3px)
- ✅ Touch targets 44x44px
- ✅ Поддержка `prefers-reduced-motion`
- ✅ Keyboard navigation

---

## 🏗️ Структура проекта

```
.
├── backend/
│   ├── cmd/server/          # Точка входа
│   ├── internal/
│   │   ├── handlers/        # HTTP обработчики
│   │   ├── models/          # Модели данных
│   │   ├── database/        # Работа с БД
│   │   └── metrics/         # Prometheus метрики
│   ├── init.sql             # Инициализация БД
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/      # Компоненты UI
│   │   ├── routes/          # Страницы
│   │   ├── lib/             # API клиент
│   │   └── styles/          # Глобальные стили
│   └── Dockerfile
├── monitoring/
│   ├── prometheus/          # Конфигурация Prometheus
│   └── grafana/             # Дашборды Grafana
├── docker-compose.yml       # Оркестрация
├── Makefile                 # Команды разработки
└── DEPLOYMENT.md            # Руководство по развертыванию
```

---

## 💻 Разработка

### Backend (Go)

```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend (Svelte)

```bash
cd frontend
npm install
npm run dev
```

### Полезные команды

```bash
make help    # Показать все команды
make dev     # Запустить dev окружение
make build   # Собрать Docker образы
make down    # Остановить сервисы
make logs    # Просмотр логов
make clean   # Очистить volumes
```

---

## 🐳 Docker

### Сервисы

- **postgres** — PostgreSQL 16
- **backend** — Go REST API
- **frontend** — Svelte + Nginx
- **prometheus** — Система мониторинга
- **grafana** — Визуализация метрик

### Volumes

- `postgres-data` — данные PostgreSQL
- `prometheus-data` — данные Prometheus
- `grafana-data` — конфигурация Grafana

---

## 📈 Grafana Dashboard

После запуска откройте http://localhost:3001

**Логин**: `admin`
**Пароль**: `admin`

### Доступные панели

- 📊 HTTP запросы в секунду
- ⏱️ Длительность HTTP запросов (p95)
- 🔢 HTTP статус коды
- 👁️ Просмотры страниц
- 📰 Просмотры новостей
- 📥 Скачивания документов
- 💰 Просмотры бюджетных отчетов
- 🗄️ Длительность запросов к БД
- 📊 Запросы к БД в секунду

---

## 🤝 Вклад в проект

Мы приветствуем вклад в развитие проекта!

1. Fork репозитория
2. Создайте feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit изменения (`git commit -m 'Add some AmazingFeature'`)
4. Push в branch (`git push origin feature/AmazingFeature`)
5. Откройте Pull Request

---

## 📝 Лицензия

Этот проект создан для департамента финансов города Якутска.

---

## 📞 Контакты

**Департамент финансов Окружной администрации города Якутска**

- 📍 677000, Республика Саха (Якутия), г. Якутск, ул. Орджоникидзе, д. 24
- 📞 +7 (4112) 42-26-26
- 📧 depfin@yakutsk.org
- 🌐 https://depfin.sakha.gov.ru

---

<div align="center">

### 🌟 Сделано с использованием современных технологий

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![Made with Svelte](https://img.shields.io/badge/Made%20with-Svelte-FF3E00?style=flat-square&logo=svelte)](https://svelte.dev/)
[![Powered by Docker](https://img.shields.io/badge/Powered%20by-Docker-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)

**[⬆ Вернуться к началу](#-департамент-финансов-города-якутска)**

</div>
