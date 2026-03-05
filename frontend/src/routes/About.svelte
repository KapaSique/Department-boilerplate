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
    {#if loading}
      <div class="loading">
        <div class="spinner"></div>
        <p>Загрузка...</p>
      </div>
    {:else if info}
      <div class="page-header">
        <h1 class="page-title">О департаменте</h1>
      </div>

      <div class="about-content">
        <div class="card main-info">
          <h2>{info.name}</h2>
          <p class="description">{info.description}</p>

          <div class="info-grid">
            <div class="info-item">
              <div class="info-icon">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
                </svg>
              </div>
              <div>
                <strong>Руководитель</strong>
                <span>{info.head}</span>
              </div>
            </div>

            <div class="info-item">
              <div class="info-icon">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                </svg>
              </div>
              <div>
                <strong>Адрес</strong>
                <span>{info.address}</span>
              </div>
            </div>

            <div class="info-item">
              <div class="info-icon">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z" />
                </svg>
              </div>
              <div>
                <strong>Телефон</strong>
                <span>{info.phone}</span>
              </div>
            </div>

            <div class="info-item">
              <div class="info-icon">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
                </svg>
              </div>
              <div>
                <strong>Email</strong>
                <span>{info.email}</span>
              </div>
            </div>

            <div class="info-item">
              <div class="info-icon">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <strong>Режим работы</strong>
                <span>{info.work_hours}</span>
              </div>
            </div>
          </div>
        </div>

        {#if info.departments && info.departments.length > 0}
          <div class="card">
            <h3>Структура департамента</h3>
            <div class="departments-grid">
              {#each info.departments as dept}
                <div class="department-item">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 21h19.5m-18-18v18m10.5-18v18m6-13.5V21M6.75 6.75h.75m-.75 3h.75m-.75 3h.75m3-6h.75m-.75 3h.75m-.75 3h.75M6.75 21v-3.375c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21M3 3h12m-.75 4.5H21m-3.75 3.75h.008v.008h-.008v-.008zm0 3h.008v.008h-.008v-.008zm0 3h.008v.008h-.008v-.008z" />
                  </svg>
                  {dept}
                </div>
              {/each}
            </div>
          </div>
        {/if}
      </div>
    {:else}
      <div class="empty-state">
        <p>Информация недоступна</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .page-header {
    text-align: center;
    margin-bottom: 4rem;
  }

  .page-title {
    font-size: 3rem;
    font-weight: 800;
    letter-spacing: -0.03em;
    background: var(--gradient-primary);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .about-content {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .main-info h2 {
    font-size: 1.75rem;
    margin-bottom: 1.5rem;
    color: var(--text);
    font-weight: 700;
  }

  .description {
    color: var(--text-secondary);
    line-height: 1.8;
    margin-bottom: 3rem;
    font-size: 1.0625rem;
  }

  .info-grid {
    display: grid;
    gap: 2rem;
  }

  .info-item {
    display: flex;
    align-items: flex-start;
    gap: 1rem;
  }

  .info-icon {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    background: rgba(8, 145, 178, 0.1);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .info-icon svg {
    width: 24px;
    height: 24px;
    color: var(--primary);
  }

  .info-item > div {
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }

  .info-item strong {
    color: var(--text);
    font-size: 0.9375rem;
    font-weight: 700;
  }

  .info-item span {
    color: var(--text-secondary);
    line-height: 1.6;
  }

  h3 {
    font-size: 1.5rem;
    margin-bottom: 2rem;
    color: var(--text);
    font-weight: 700;
  }

  .departments-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 1rem;
  }

  .department-item {
    display: flex;
    align-items: center;
    gap: 0.875rem;
    padding: 1.25rem;
    background: var(--background);
    border-radius: 12px;
    border: 1px solid var(--border-light);
    transition: all 0.2s;
    font-weight: 500;
    color: var(--text);
  }

  .department-item:hover {
    border-color: var(--primary);
    background: var(--surface);
  }

  .department-item svg {
    width: 20px;
    height: 20px;
    color: var(--primary);
    flex-shrink: 0;
  }

  .loading {
    text-align: center;
    padding: 6rem 0;
  }

  .spinner {
    width: 48px;
    height: 48px;
    border: 4px solid var(--border);
    border-top-color: var(--primary);
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
    color: var(--text-muted);
  }

  @media (max-width: 768px) {
    .page-title {
      font-size: 2.25rem;
    }

    .departments-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
