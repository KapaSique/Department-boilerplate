<script>
  import { onMount } from 'svelte'
  import { api } from '../lib/api.js'

  let info = null
  let loading = true

  onMount(async () => {
    try {
      info = await api.getDepartmentInfo()
    } catch (error) {
      console.error('Failed to load department info:', error)
    } finally {
      loading = false
    }
  })
</script>

<section class="section">
  <div class="container">
    <h1 class="section-title">О департаменте</h1>

    {#if loading}
      <p>Загрузка...</p>
    {:else if info}
      <div class="about-content">
        <div class="card main-info">
          <h2>{info.name}</h2>
          <p class="description">{info.description}</p>

          <div class="info-grid">
            <div class="info-item">
              <strong>Руководитель:</strong>
              <span>{info.head}</span>
            </div>
            <div class="info-item">
              <strong>Адрес:</strong>
              <span>{info.address}</span>
            </div>
            <div class="info-item">
              <strong>Телефон:</strong>
              <span>{info.phone}</span>
            </div>
            <div class="info-item">
              <strong>Email:</strong>
              <span>{info.email}</span>
            </div>
            <div class="info-item">
              <strong>Режим работы:</strong>
              <span>{info.work_hours}</span>
            </div>
          </div>
        </div>

        {#if info.departments && info.departments.length > 0}
          <div class="card">
            <h3>Структура департамента</h3>
            <ul class="departments-list">
              {#each info.departments as dept}
                <li>{dept}</li>
              {/each}
            </ul>
          </div>
        {/if}
      </div>
    {:else}
      <div class="card">
        <p>Информация недоступна</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .about-content {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .main-info h2 {
    font-size: 1.75rem;
    margin-bottom: 1rem;
    color: var(--text);
  }

  .description {
    color: var(--text-secondary);
    line-height: 1.8;
    margin-bottom: 2rem;
  }

  .info-grid {
    display: grid;
    gap: 1rem;
  }

  .info-item {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .info-item strong {
    color: var(--text);
    font-size: 0.875rem;
  }

  .info-item span {
    color: var(--text-secondary);
  }

  h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
    color: var(--text);
  }

  .departments-list {
    list-style: none;
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .departments-list li {
    padding: 0.75rem;
    background: var(--background);
    border-radius: 6px;
    border: 1px solid var(--border);
  }
</style>
