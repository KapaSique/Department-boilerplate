-- Initialize database schema
CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS documents (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    file_url VARCHAR(500) NOT NULL,
    category VARCHAR(100),
    published_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS budget_reports (
    id SERIAL PRIMARY KEY,
    year INTEGER NOT NULL,
    quarter INTEGER,
    title VARCHAR(255) NOT NULL,
    revenue DECIMAL(15,2),
    expenses DECIMAL(15,2),
    balance DECIMAL(15,2),
    file_url VARCHAR(500),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_news_published ON news(published_at DESC);
CREATE INDEX IF NOT EXISTS idx_documents_category ON documents(category);
CREATE INDEX IF NOT EXISTS idx_budget_year ON budget_reports(year DESC, quarter DESC);

-- Insert sample data
INSERT INTO news (title, content, published_at) VALUES
('Утверждён бюджет города на 2026 год', 'Городская Дума утвердила бюджет города Якутска на 2026 год. Общий объём доходов составит 15,2 млрд рублей, расходов - 15,5 млрд рублей.', '2026-01-15 10:00:00'),
('Отчёт об исполнении бюджета за 2025 год', 'Департамент финансов представил отчёт об исполнении бюджета города за 2025 год. Доходы исполнены на 98,5%, расходы - на 97,2%.', '2026-02-10 14:30:00'),
('Новые меры поддержки малого бизнеса', 'С 1 марта вступают в силу новые меры финансовой поддержки субъектов малого и среднего предпринимательства.', '2026-02-28 09:00:00');

INSERT INTO documents (title, description, file_url, category, published_at) VALUES
('Решение о бюджете на 2026 год', 'Решение Якутской городской Думы об утверждении бюджета города на 2026 год', '/files/budget-2026.pdf', 'budget', '2026-01-15'),
('Положение о бюджетном процессе', 'Положение о бюджетном процессе в городском округе "город Якутск"', '/files/budget-process.pdf', 'regulations', '2025-12-01'),
('Отчёт об исполнении бюджета 2025', 'Годовой отчёт об исполнении бюджета города за 2025 год', '/files/report-2025.pdf', 'reports', '2026-02-10');

INSERT INTO budget_reports (year, quarter, title, revenue, expenses, balance) VALUES
(2025, 4, 'Бюджет 4 квартал 2025', 3800000000, 3750000000, 50000000),
(2025, 3, 'Бюджет 3 квартал 2025', 3600000000, 3650000000, -50000000),
(2026, 1, 'Бюджет 1 квартал 2026', 3900000000, 3850000000, 50000000);
