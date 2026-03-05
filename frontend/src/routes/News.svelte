<script>
  import { onMount } from 'svelte'
  import { Link } from 'svelte-routing'
  import { api } from '../lib/api.js'

  let news = []
  let loading = true

  onMount(async () => {
    try {
      news = await api.getNews(20)
    } catch (error) {
      console.error('Failed to load news:', error)
    } finally {
      loading = false
    }
  })
</script>

<section class="section">
  <div class="container">
    <div class="page-header">
      <h1 class="page-title">Новости и объявления</h1>
      <p class="page-description">Актуальные новости и важные объявления департамента финансов города Якутска</p>
    </div>

    {#if loading}
      <div class="loading">
        <div class="spinner"></div>
        <p>Загрузка новостей...</p>
      </div>
    {:else if news.length > 0}
      <div class="news-list">
        {#each news as item}
          <Link to="/news/{item.id}">
            <article class="card news-item">
              <div class="news-header">
                <div class="news-meta">
                  <span class="badge badge-primary">Новость</span>
                  <time>{new Date(item.published_at).toLocaleDateString('ru-RU', {
                    day: 'numeric',
                    month: 'long',
                    year: 'numeric'
                  })}</time>
                </div>
                <h2>{item.title}</h2>
              </div>
              <p class="news-excerpt">{item.content.substring(0, 220)}...</p>
              <span class="read-more">
                Читать полностью
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
                </svg>
              </span>
            </article>
          </Link>
        {/each}
      </div>
    {:else}
      <div class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" />
        </svg>
        <p>Новостей пока нет</p>
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
    background: var(--gradient-primary);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .page-description {
    font-size: 1.125rem;
    color: var(--text-secondary);
    line-height: 1.7;
  }

  .news-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .news-item {
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
    overflow: hidden;
  }

  .news-item::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 4px;
    height: 100%;
    background: var(--gradient-primary);
    transform: scaleY(0);
    transform-origin: top;
    transition: transform 0.3s ease;
  }

  .news-item:hover::before {
    transform: scaleY(1);
  }

  .news-header {
    margin-bottom: 1.25rem;
  }

  .news-meta {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1rem;
    flex-wrap: wrap;
  }

  .news-meta time {
    font-size: 0.875rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  .news-item h2 {
    font-size: 1.5rem;
    color: var(--text);
    font-weight: 700;
    line-height: 1.4;
    letter-spacing: -0.02em;
  }

  .news-excerpt {
    color: var(--text-secondary);
    line-height: 1.7;
    margin-bottom: 1.5rem;
  }

  .read-more {
    color: var(--primary);
    font-weight: 600;
    font-size: 0.9375rem;
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    transition: gap 0.2s ease;
  }

  .read-more svg {
    width: 18px;
    height: 18px;
  }

  .news-item:hover .read-more {
    gap: 0.75rem;
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
    font-size: 1rem;
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

    .page-description {
      font-size: 1rem;
    }

    .news-item h2 {
      font-size: 1.25rem;
    }
  }
</style>
