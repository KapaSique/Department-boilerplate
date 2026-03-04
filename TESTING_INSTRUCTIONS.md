# 🧪 Инструкции по тестированию проекта

## Перед началом

1. **Запустите Docker Desktop**
   - Откройте приложение Docker Desktop
   - Дождитесь, пока Docker полностью запустится (иконка в трее станет зеленой)

2. **Освободите порты**
   ```bash
   # Проверьте, что порты свободны
   lsof -i :3000,8080,9090,3001,5432
   
   # Если порты заняты, остановите процессы
   ```

## Автоматическое тестирование (рекомендуется)

```bash
cd /Users/artemcike/Documents/Imi-labs/Department-boilerplate
./test.sh
```

Скрипт выполнит:
- ✅ Проверку Docker
- 🧹 Очистку старых контейнеров
- 🔨 Сборку всех образов
- ▶️ Запуск сервисов
- 💾 Инициализацию БД
- ✅ Проверку всех endpoints
- 📊 Вывод итогового отчета

## Ручное тестирование

### Шаг 1: Сборка

```bash
cd /Users/artemcike/Documents/Imi-labs/Department-boilerplate

# Очистка
docker-compose down -v

# Сборка
docker-compose build --no-cache
```

### Шаг 2: Запуск

```bash
# Запустить все сервисы
docker-compose up -d

# Проверить статус
docker-compose ps
```

### Шаг 3: Инициализация БД

```bash
# Подождать 10-15 секунд, затем:
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

### Шаг 4: Проверка сервисов

#### Backend API
```bash
# Health check
curl http://localhost:8080/api/v1/health

# Получить новости
curl http://localhost:8080/api/v1/news | jq

# Получить информацию о департаменте
curl http://localhost:8080/api/v1/department/info | jq

# Метрики Prometheus
curl http://localhost:8080/metrics | grep http_requests_total
```

#### Frontend
```bash
# Открыть в браузере
open http://localhost:3000

# Или проверить через curl
curl -I http://localhost:3000
```

#### Prometheus
```bash
# Открыть UI
open http://localhost:9090

# Проверить targets
open http://localhost:9090/targets

# Проверить метрики
curl http://localhost:9090/api/v1/query?query=http_requests_total
```

#### Grafana
```bash
# Открыть UI (admin/admin)
open http://localhost:3001

# Проверить health
curl http://localhost:3001/api/health
```

## Тестирование функционала

### 1. Тест новостей

```bash
# Получить список новостей
curl http://localhost:8080/api/v1/news

# Получить конкретную новость
curl http://localhost:8080/api/v1/news/1
```

### 2. Тест документов

```bash
# Все документы
curl http://localhost:8080/api/v1/documents

# Документы по категории
curl "http://localhost:8080/api/v1/documents?category=budget"
```

### 3. Тест бюджета

```bash
# Все бюджетные данные
curl http://localhost:8080/api/v1/budget

# За конкретный год
curl "http://localhost:8080/api/v1/budget?year=2025"
```

### 4. Тест метрик

```bash
# Сделать несколько запросов
for i in {1..10}; do
  curl -s http://localhost:8080/api/v1/news > /dev/null
done

# Проверить метрики
curl http://localhost:8080/metrics | grep http_requests_total
```

## Проверка логов

```bash
# Все логи
docker-compose logs

# Конкретный сервис
docker-compose logs backend
docker-compose logs frontend
docker-compose logs postgres

# Следить за логами в реальном времени
docker-compose logs -f backend
```

## Проверка производительности

```bash
# Использование ресурсов
docker stats

# Нагрузочное тестирование (если установлен ab)
ab -n 1000 -c 10 http://localhost:8080/api/v1/news
```

## Проверка доступности (a11y)

1. Откройте http://localhost:3000 в браузере
2. Откройте DevTools (F12)
3. Перейдите в Lighthouse
4. Запустите аудит Accessibility
5. Проверьте, что score >= 90

## Проверка responsive дизайна

1. Откройте http://localhost:3000
2. Откройте DevTools (F12)
3. Включите Device Toolbar (Cmd+Shift+M)
4. Проверьте на разных устройствах:
   - iPhone SE (375px)
   - iPad (768px)
   - Desktop (1440px)

## Проверка клавиатурной навигации

1. Откройте http://localhost:3000
2. Используйте только клавиатуру:
   - Tab - переход между элементами
   - Enter - активация ссылок/кнопок
   - Escape - закрытие модальных окон
3. Проверьте, что focus states видны

## Остановка и очистка

```bash
# Остановить сервисы
docker-compose down

# Остановить и удалить volumes
docker-compose down -v

# Полная очистка
docker system prune -a --volumes
```

## Troubleshooting

### Проблема: "Cannot connect to Docker daemon"
**Решение**: Запустите Docker Desktop

### Проблема: "Port already in use"
**Решение**: 
```bash
# Найти процесс
lsof -i :8080

# Убить процесс
kill -9 <PID>
```

### Проблема: "Backend не подключается к БД"
**Решение**:
```bash
# Проверить логи PostgreSQL
docker-compose logs postgres

# Пересоздать контейнер БД
docker-compose down -v
docker-compose up -d postgres
sleep 10
docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql
```

### Проблема: "Frontend показывает ошибку 502"
**Решение**:
```bash
# Проверить, что backend запущен
docker-compose ps backend

# Перезапустить frontend
docker-compose restart frontend
```

## CI/CD тестирование

```bash
# Установить act (если еще не установлен)
brew install act

# Запустить GitHub Actions локально
act -j backend-test
act -j frontend-build
act -j docker-build
```

## Итоговый чеклист

- [ ] Docker Desktop запущен
- [ ] Все порты свободны
- [ ] Образы собраны успешно
- [ ] Все контейнеры запущены
- [ ] База данных инициализирована
- [ ] Backend отвечает на /health
- [ ] Frontend открывается в браузере
- [ ] Prometheus собирает метрики
- [ ] Grafana показывает дашборды
- [ ] API endpoints возвращают данные
- [ ] Логи не содержат ошибок

## Успешный результат

При успешном тестировании вы должны увидеть:

```
🎉 Все тесты пройдены успешно!

Сервисы доступны по адресам:
  • Frontend:   http://localhost:3000
  • Backend:    http://localhost:8080
  • Prometheus: http://localhost:9090
  • Grafana:    http://localhost:3001 (admin/admin)
```

Проект готов к использованию! 🚀
