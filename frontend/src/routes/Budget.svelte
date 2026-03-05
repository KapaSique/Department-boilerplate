<script>
  import { onMount } from 'svelte'
  import { api } from '../lib/api.js'

  let budgetData = []
  let loading = true
  let selectedYear = ''

  const currentYear = new Date().getFullYear()
  const years = Array.from({ length: 5 }, (_, i) => currentYear - i)

  onMount(() => {
    loadBudgetData()
  })

  async function loadBudgetData() {
    loading = true
    try {
      budgetData = await api.getBudgetData(selectedYear)
    } catch (error) {
      console.error('Failed to load budget data:', error)
    } finally {
      loading = false
    }
  }

  function handleYearChange(event) {
    selectedYear = event.target.value
    loadBudgetData()
  }

  function formatCurrency(value) {
    return new Intl.NumberFormat('ru-RU', {
      style: 'currency',
      currency: 'RUB',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0
    }).format(value)
  }
</script>

<section class="section">
  <div class="container">
    <div class="page-header">
      <h1 class="page-title">Бюджетные данные и отчеты</h1>
      <p class="page-description">Информация о доходах, расходах и исполнении бюджета города Якутска</p>
    </div>

    <div class="filters">
      <div class="filter-group">
        <label for="year">Год:</label>
        <select id="year" bind:value={selectedYear} on:change={handleYearChange}>
          <option value="">Все годы</option>
          {#each years as year}
            <option value={year}>{year}</option>
          {/each}
        </select>
      </div>
    </div>

    {#if loading}
      <div class="loading">
        <div class="spinner"></div>
        <p>Загрузка данных...</p>
      </div>
    {:else if budgetData.length > 0}
      <div class="budget-list">
        {#each budgetData as report}
          <div class="card budget-item">
            <div class="budget-header">
              <div>
                <h3>{report.title}</h3>
                <div class="budget-period">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5" />
                  </svg>
                  {report.year} год, {report.quarter} квартал
                </div>
              </div>
            </div>

            <div class="budget-stats">
              <div class="stat revenue">
                <div class="stat-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 18L9 11.25l4.306 4.307a11.95 11.95 0 015.814-5.519l2.74-1.22m0 0l-5.94-2.28m5.94 2.28l-2.28 5.941" />
                  </svg>
                </div>
                <div class="stat-content">
                  <span class="label">Доходы</span>
                  <span class="value">{formatCurrency(report.revenue)}</span>
                </div>
              </div>

              <div class="stat expenses">
                <div class="stat-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6L9 12.75l4.286-4.286a11.948 11.948 0 014.306 6.43l.776 2.898m0 0l3.182-5.511m-3.182 5.511l-5.511-3.181" />
                  </svg>
                </div>
                <div class="stat-content">
                  <span class="label">Расходы</span>
                  <span class="value">{formatCurrency(report.expenses)}</span>
                </div>
              </div>

              <div class="stat balance" class:positive={report.balance >= 0} class:negative={report.balance < 0}>
                <div class="stat-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 3v17.25m0 0c-1.472 0-2.882.265-4.185.75M12 20.25c1.472 0 2.882.265 4.185.75M18.75 4.97A48.416 48.416 0 0012 4.5c-2.291 0-4.545.16-6.75.47m13.5 0c1.01.143 2.01.317 3 .52m-3-.52l2.62 10.726c.122.499-.106 1.028-.589 1.202a5.988 5.988 0 01-2.031.352 5.988 5.988 0 01-2.031-.352c-.483-.174-.711-.703-.59-1.202L18.75 4.971zm-16.5.52c.99-.203 1.99-.377 3-.52m0 0l2.62 10.726c.122.499-.106 1.028-.589 1.202a5.989 5.989 0 01-2.031.352 5.989 5.989 0 01-2.031-.352c-.483-.174-.711-.703-.59-1.202L5.25 4.971z" />
                  </svg>
                </div>
                <div class="stat-content">
                  <span class="label">Баланс</span>
                  <span class="value">{formatCurrency(report.balance)}</span>
                </div>
              </div>
            </div>

            {#if report.file_url}
              <a href={report.file_url} class="btn btn-secondary" download>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3" />
                </svg>
                Скачать отчет
              </a>
            {/if}
          </div>
        {/each}
      </div>
    {:else}
      <div class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z" />
        </svg>
        <p>Бюджетных данных не найдено</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .page-header {
    text-align: center;
    max-width: 700px;
    margin: 0 auto 4rem;
  }

  .page-title {
    font-size: 3rem;
    font-weight: 800;
    margin-bottom: 1rem;
    letter-spacing: -0.03em;
    background: var(--gradient-accent);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .page-description {
    font-size: 1.125rem;
    color: var(--text-secondary);
    line-height: 1.7;
  }

  .filters {
    margin-bottom: 3rem;
  }

  .filter-group {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .filter-group label {
    font-weight: 600;
    color: var(--text);
    font-size: 0.9375rem;
  }

  select {
    padding: 0.75rem 1.25rem;
    border: 2px solid var(--border);
    border-radius: 12px;
    background: var(--surface);
    color: var(--text);
    font-size: 0.9375rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    min-width: 200px;
  }

  select:hover {
    border-color: var(--primary);
  }

  select:focus {
    outline: none;
    border-color: var(--primary);
    box-shadow: 0 0 0 3px rgba(8, 145, 178, 0.1);
  }

  .budget-list {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .budget-item {
    position: relative;
    overflow: hidden;
  }

  .budget-item::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: var(--gradient-accent);
  }

  .budget-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 2rem;
    gap: 1rem;
  }

  .budget-header h3 {
    font-size: 1.5rem;
    color: var(--text);
    font-weight: 700;
    margin-bottom: 0.75rem;
  }

  .budget-period {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: var(--text-muted);
    font-size: 0.9375rem;
    font-weight: 500;
  }

  .budget-period svg {
    width: 18px;
    height: 18px;
  }

  .budget-stats {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
  }

  .stat {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 1.25rem;
    background: var(--background);
    border-radius: 12px;
    border: 1px solid var(--border-light);
  }

  .stat-icon {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .stat-icon svg {
    width: 24px;
    height: 24px;
  }

  .stat.revenue .stat-icon {
    background: rgba(34, 197, 94, 0.1);
    color: #10b981;
  }

  .stat.expenses .stat-icon {
    background: rgba(245, 158, 11, 0.1);
    color: #f59e0b;
  }

  .stat.balance .stat-icon {
    background: rgba(8, 145, 178, 0.1);
    color: var(--primary);
  }

  .stat.balance.negative .stat-icon {
    background: rgba(239, 68, 68, 0.1);
    color: #ef4444;
  }

  .stat-content {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .label {
    font-size: 0.875rem;
    color: var(--text-muted);
    font-weight: 600;
  }

  .value {
    font-size: 1.375rem;
    font-weight: 700;
    color: var(--text);
  }

  .stat.revenue .value {
    color: #10b981;
  }

  .stat.expenses .value {
    color: #f59e0b;
  }

  .stat.balance.positive .value {
    color: #10b981;
  }

  .stat.balance.negative .value {
    color: #ef4444;
  }

  .loading {
    text-align: center;
    padding: 6rem 0;
  }

  .spinner {
    width: 48px;
    height: 48px;
    border: 4px solid var(--border);
    border-top-color: var(--accent);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto 1.5rem;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  .empty-state {
    text-align: center;
    padding: 6rem 0;
  }

  .empty-state svg {
    width: 64px;
    height: 64px;
    color: var(--text-muted);
    margin: 0 auto 1.5rem;
    opacity: 0.5;
  }

  .empty-state p {
    color: var(--text-muted);
    font-size: 1.125rem;
  }

  @media (max-width: 768px) {
    .page-title {
      font-size: 2.25rem;
    }

    .budget-stats {
      grid-template-columns: 1fr;
    }
  }
</style>
