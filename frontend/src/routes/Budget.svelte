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
      minimumFractionDigits: 0
    }).format(value)
  }
</script>

<section class="section">
  <div class="container">
    <h1 class="section-title">Бюджетные данные и отчеты</h1>

    <div class="filters">
      <select bind:value={selectedYear} on:change={handleYearChange}>
        <option value="">Все годы</option>
        {#each years as year}
          <option value={year}>{year}</option>
        {/each}
      </select>
    </div>

    {#if loading}
      <p>Загрузка данных...</p>
    {:else if budgetData.length > 0}
      <div class="budget-list">
        {#each budgetData as report}
          <div class="card budget-item">
            <div class="budget-header">
              <h3>{report.title}</h3>
              <span class="period">{report.year} год, {report.quarter} квартал</span>
            </div>

            <div class="budget-stats">
              <div class="stat">
                <span class="label">Доходы</span>
                <span class="value revenue">{formatCurrency(report.revenue)}</span>
              </div>
              <div class="stat">
                <span class="label">Расходы</span>
                <span class="value expenses">{formatCurrency(report.expenses)}</span>
              </div>
              <div class="stat">
                <span class="label">Баланс</span>
                <span class="value balance" class:positive={report.balance >= 0} class:negative={report.balance < 0}>
                  {formatCurrency(report.balance)}
                </span>
              </div>
            </div>

            {#if report.file_url}
              <a href={report.file_url} class="btn btn-primary" download>Скачать отчет</a>
            {/if}
          </div>
        {/each}
      </div>
    {:else}
      <div class="card">
        <p>Бюджетных данных не найдено</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .filters {
    margin-bottom: 2rem;
  }

  select {
    padding: 0.75rem 1rem;
    border: 1px solid var(--border);
    border-radius: 6px;
    background: var(--surface);
    color: var(--text);
    font-size: 1rem;
    cursor: pointer;
  }

  .budget-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .budget-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
    gap: 1rem;
  }

  .budget-header h3 {
    font-size: 1.25rem;
    color: var(--text);
  }

  .period {
    font-size: 0.875rem;
    color: var(--text-secondary);
    white-space: nowrap;
  }

  .budget-stats {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1.5rem;
    margin-bottom: 1.5rem;
  }

  .stat {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .label {
    font-size: 0.875rem;
    color: var(--text-secondary);
    font-weight: 500;
  }

  .value {
    font-size: 1.5rem;
    font-weight: 700;
  }

  .revenue {
    color: #10b981;
  }

  .expenses {
    color: #f59e0b;
  }

  .balance.positive {
    color: #10b981;
  }

  .balance.negative {
    color: #ef4444;
  }

  @media (max-width: 768px) {
    .budget-header {
      flex-direction: column;
      align-items: flex-start;
    }

    .budget-stats {
      grid-template-columns: 1fr;
    }
  }
</style>
