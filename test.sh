#!/bin/bash

set -e

echo "🚀 Полный цикл тестирования проекта Department Boilerplate"
echo "============================================================"
echo ""

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Функция для проверки
check_service() {
    local url=$1
    local name=$2
    local max_attempts=30
    local attempt=1

    echo -n "Проверка $name... "

    while [ $attempt -le $max_attempts ]; do
        if curl -f -s "$url" > /dev/null 2>&1; then
            echo -e "${GREEN}✅ OK${NC}"
            return 0
        fi
        sleep 1
        attempt=$((attempt + 1))
    done

    echo -e "${RED}❌ FAIL${NC}"
    return 1
}

# 1. Проверка Docker
echo "1️⃣  Проверка Docker..."
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker не запущен!${NC}"
    echo "Запустите Docker Desktop и попробуйте снова."
    exit 1
fi
echo -e "${GREEN}✅ Docker запущен${NC}"
echo ""

# 2. Очистка
echo "2️⃣  Очистка старых контейнеров..."
docker-compose down -v > /dev/null 2>&1 || true
echo -e "${GREEN}✅ Очистка завершена${NC}"
echo ""

# 3. Сборка
echo "3️⃣  Сборка Docker образов..."
echo "   Это может занять несколько минут..."
if docker-compose build --no-cache; then
    echo -e "${GREEN}✅ Сборка успешна${NC}"
else
    echo -e "${RED}❌ Ошибка сборки${NC}"
    exit 1
fi
echo ""

# 4. Запуск
echo "4️⃣  Запуск сервисов..."
if docker-compose up -d; then
    echo -e "${GREEN}✅ Сервисы запущены${NC}"
else
    echo -e "${RED}❌ Ошибка запуска${NC}"
    exit 1
fi
echo ""

# 5. Ожидание
echo "5️⃣  Ожидание готовности сервисов..."
sleep 15
echo -e "${GREEN}✅ Ожидание завершено${NC}"
echo ""

# 6. Инициализация БД
echo "6️⃣  Инициализация базы данных..."
if docker exec -i finance-dept-db psql -U postgres -d finance_dept < backend/init.sql > /dev/null 2>&1; then
    echo -e "${GREEN}✅ База данных инициализирована${NC}"
else
    echo -e "${YELLOW}⚠️  База данных уже инициализирована или ошибка${NC}"
fi
echo ""

# 7. Проверка сервисов
echo "7️⃣  Проверка работоспособности сервисов..."
echo ""

FAILED=0

check_service "http://localhost:8080/api/v1/health" "Backend API" || FAILED=$((FAILED + 1))
check_service "http://localhost:3000" "Frontend" || FAILED=$((FAILED + 1))
check_service "http://localhost:9090/-/healthy" "Prometheus" || FAILED=$((FAILED + 1))
check_service "http://localhost:3001/api/health" "Grafana" || FAILED=$((FAILED + 1))

echo ""

# 8. Проверка API endpoints
echo "8️⃣  Проверка API endpoints..."
echo ""

echo -n "GET /api/v1/news... "
if curl -f -s "http://localhost:8080/api/v1/news" | grep -q "id"; then
    echo -e "${GREEN}✅ OK${NC}"
else
    echo -e "${RED}❌ FAIL${NC}"
    FAILED=$((FAILED + 1))
fi

echo -n "GET /api/v1/department/info... "
if curl -f -s "http://localhost:8080/api/v1/department/info" | grep -q "name"; then
    echo -e "${GREEN}✅ OK${NC}"
else
    echo -e "${RED}❌ FAIL${NC}"
    FAILED=$((FAILED + 1))
fi

echo -n "GET /metrics... "
if curl -f -s "http://localhost:8080/metrics" | grep -q "http_requests_total"; then
    echo -e "${GREEN}✅ OK${NC}"
else
    echo -e "${RED}❌ FAIL${NC}"
    FAILED=$((FAILED + 1))
fi

echo ""

# 9. Итоги
echo "============================================================"
if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}🎉 Все тесты пройдены успешно!${NC}"
    echo ""
    echo "Сервисы доступны по адресам:"
    echo "  • Frontend:   http://localhost:3000"
    echo "  • Backend:    http://localhost:8080"
    echo "  • Prometheus: http://localhost:9090"
    echo "  • Grafana:    http://localhost:3001 (admin/admin)"
    echo ""
    echo "Для просмотра логов: docker-compose logs -f"
    echo "Для остановки: docker-compose down"
    exit 0
else
    echo -e "${RED}❌ Тесты провалены: $FAILED ошибок${NC}"
    echo ""
    echo "Для просмотра логов: docker-compose logs"
    echo "Для отладки: см. DOCKER_DEBUG.md"
    exit 1
fi
