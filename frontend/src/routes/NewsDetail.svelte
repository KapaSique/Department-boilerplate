<script>
  import { onMount } from 'svelte'
  import { api } from '../lib/api.js'

  export let id

  let news = null
  let loading = true
  let error = null

  onMount(async () => {
    try {
      news = await api.getNewsById(id)
    } catch (err) {
      error = 'Не удалось загрузить новость'
      console.error('Failed to load news:', err)
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
    {:else if error}
      <div class="error-state">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z" />
        </svg>
        <p>{error}</p>
        <a href="/news" class="btn btn-primary">Вернуться к новостям</a>
      </div>
    {:else if news}
      <article class="news-detail">
        <div class="news-detail-header">
          <div class="breadcrumb">
            <a href="/">Главная</a>
            <span>/</span>
            <a href="/news">Новости</a>
            <span>/</span>
            <span>{news.title}</span>
          </div>

          <div class="news-detail-meta">
            <span class="badge badge-primary">Новость</span>
            <time>{new Date(news.published_at).toLocaleDateString('ru-RU', {
              day: 'numeric',
              month: 'long',
              year: 'numeric'
            })}</time>
          </div>

          <h1>{news.title}</h1>
        </div>

        <div class="news-detail-content">
          {news.content}
        </div>

        <div class="news-detail-footer">
          <a href="/news" class="btn btn-secondary">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
            </svg>
            Вернуться к новостям
          </a>
        </div>
      </article>
    {/if}
  </div>
</section>

<style>
  .news-detail {
    max-width: 800px;
    margin: 0 auto;
  }

  .news-detail-header {
    margin-bottom: 3rem;
  }

  .breadcrumb {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    margin-bottom: 2rem;
    font-size: 0.875rem;
    flex-wrap: wrap;
  }

  .breadcrumb a {
    color: var(--text-secondary);
    transition: color 0.2s;
  }

  .breadcrumb a:hover {
    color: var(--primary);
  }

  .breadcrumb span:last-child {
    color: var(--text-muted);
  }

  .breadcrumb span:not(:last-child) {
    color: var(--border);
  }

  .news-detail-meta {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.5rem;
    flex-wrap: wrap;
  }

  .news-detail-meta time {
    font-size: 0.9375rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  .news-detail h1 {
    font-size: 2.5rem;
    margin-bottom: 0;
    color: var(--text);
    font-weight: 800;
    line-height: 1.2;
    letter-spacing: -0.03em;
  }

  .news-detail-content {
    line-height: 1.8;
    color: var(--text);
    white-space: pre-wrap;
    font-size: 1.0625rem;
    padding: 2.5rem;
    background: var(--surface);
    border-radius: 16px;
    border: 1px solid var(--border-light);
    box-shadow: var(--shadow-sm);
  }

  .news-detail-footer {
    margin-top: 3rem;
    padding-top: 2rem;
    border-top: 1px solid var(--border);
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

  .loading p {
    color: var(--text-muted);
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  .error-state {
    text-align: center;
    padding: 6rem 0;
  }

  .error-state svg {
    width: 64px;
    height: 64px;
    color: #ef4444;
    margin: 0 auto 1.5rem;
  }

  .error-state p {
    color: var(--text-secondary);
    font-size: 1.125rem;
    margin-bottom: 2rem;
  }

  @media (max-width: 768px) {
    .news-detail h1 {
      font-size: 2rem;
    }

    .news-detail-content {
      padding: 1.5rem;
      font-size: 1rem;
    }
  }
</style>
