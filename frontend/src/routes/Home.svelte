<script>
  import { onMount } from 'svelte'
  import { Link } from 'svelte-routing'
  import { api } from '../lib/api.js'

  let news = []
  let loading = true

  onMount(async () => {
    try {
      news = await api.getNews(3)
    } catch (error) {
      console.error('Failed to load news:', error)
    } finally {
      loading = false
    }
  })
</script>

<div class="hero">
  <div class="container">
    <div class="hero-content">
      <div class="hero-badge">
        <span class="badge badge-primary">Официальный портал</span>
      </div>
      <h1>Департамент финансов<br/>города Якутска</h1>
      <p>Управление бюджетом и финансами муниципального образования.<br/>Прозрачность, эффективность, ответственность.</p>
      <div class="hero-actions">
        <Link to="/budget" class="btn btn-primary">Бюджет города</Link>
        <Link to="/documents" class="btn btn-secondary">Документы</Link>
      </div>
    </div>
  </div>
</div>

<section class="section">
  <div class="container">
    <div class="grid grid-3">
      <Link to="/news">
        <div class="card feature">
          <div class="feature-icon">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" />
            </svg>
          </div>
          <h3>Новости</h3>
          <p>Актуальные новости и объявления департамента финансов</p>
          <span class="feature-link">Перейти →</span>
        </div>
      </Link>

      <Link to="/budget">
        <div class="card feature">
          <div class="feature-icon accent">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z" />
            </svg>
          </div>
          <h3>Бюджет</h3>
          <p>Бюджетные данные и финансовые отчеты города</p>
          <span class="feature-link">Перейти →</span>
        </div>
      </Link>

      <Link to="/documents">
        <div class="card feature">
          <div class="feature-icon secondary">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z" />
            </svg>
          </div>
          <h3>Документы</h3>
          <p>Нормативные акты и официальные документы</p>
          <span class="feature-link">Перейти →</span>
        </div>
      </Link>
    </div>
  </div>
</section>

<section class="section section-alt">
  <div class="container">
    <div class="section-header">
      <h2 class="section-title">Последние новости</h2>
      <Link to="/news" class="btn btn-secondary">Все новости</Link>
    </div>

    {#if loading}
      <div class="loading">
        <div class="spinner"></div>
        <p>Загрузка новостей...</p>
      </div>
    {:else if news.length > 0}
      <div class="grid grid-2">
        {#each news as item}
          <Link to="/news/{item.id}">
            <div class="card news-card">
              <div class="news-meta">
                <span class="badge badge-primary">Новость</span>
                <time>{new Date(item.published_at).toLocaleDateString('ru-RU', {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric'
                })}</time>
              </div>
              <h3>{item.title}</h3>
              <p class="excerpt">{item.content.substring(0, 180)}...</p>
              <span class="read-more">Читать далее →</span>
            </div>
          </Link>
        {/each}
      </div>
    {:else}
      <div class="empty-state">
        <p>Новостей пока нет</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .hero {
    background: var(--gradient-primary);
    color: white;
    padding: 6rem 0;
    position: relative;
    overflow: hidden;
  }

  .hero::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background:
      radial-gradient(circle at 20% 50%, rgba(34, 211, 238, 0.3) 0%, transparent 50%),
      radial-gradient(circle at 80% 80%, rgba(34, 197, 94, 0.2) 0%, transparent 50%);
    pointer-events: none;
  }

  .hero-content {
    position: relative;
    max-width: 800px;
    text-align: center;
    margin: 0 auto;
  }

  .hero-badge {
    margin-bottom: 1.5rem;
  }

  .hero h1 {
    font-size: 3.5rem;
    margin-bottom: 1.5rem;
    font-weight: 800;
    letter-spacing: -0.03em;
    line-height: 1.1;
  }

  .hero p {
    font-size: 1.25rem;
    opacity: 0.95;
    line-height: 1.7;
    margin-bottom: 2.5rem;
    font-weight: 400;
  }

  .hero-actions {
    display: flex;
    gap: 1rem;
    justify-content: center;
    flex-wrap: wrap;
  }

  .feature {
    text-align: left;
    position: relative;
    overflow: hidden;
  }

  .feature::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: var(--gradient-primary);
    transform: scaleX(0);
    transform-origin: left;
    transition: transform 0.3s ease;
  }

  .feature:hover::before {
    transform: scaleX(1);
  }

  .feature-icon {
    width: 56px;
    height: 56px;
    border-radius: 14px;
    background: rgba(8, 145, 178, 0.1);
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 1.5rem;
    transition: all 0.3s ease;
  }

  .feature-icon svg {
    width: 28px;
    height: 28px;
    color: var(--primary);
  }

  .feature-icon.accent {
    background: rgba(34, 197, 94, 0.1);
  }

  .feature-icon.accent svg {
    color: var(--accent);
  }

  .feature-icon.secondary {
    background: rgba(34, 211, 238, 0.1);
  }

  .feature-icon.secondary svg {
    color: var(--secondary);
  }

  .feature:hover .feature-icon {
    transform: scale(1.1) rotate(5deg);
  }

  .feature h3 {
    font-size: 1.5rem;
    margin-bottom: 0.75rem;
    color: var(--text);
    font-weight: 700;
  }

  .feature p {
    color: var(--text-secondary);
    line-height: 1.7;
    margin-bottom: 1.5rem;
  }

  .feature-link {
    color: var(--primary);
    font-weight: 600;
    font-size: 0.9375rem;
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    transition: gap 0.2s ease;
  }

  .feature:hover .feature-link {
    gap: 0.75rem;
  }

  .section-alt {
    background: white;
  }

  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 3rem;
    flex-wrap: wrap;
    gap: 1.5rem;
  }

  .news-card {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .news-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.25rem;
    gap: 1rem;
  }

  .news-meta time {
    font-size: 0.875rem;
    color: var(--text-muted);
    font-weight: 500;
  }

  .news-card h3 {
    font-size: 1.375rem;
    margin-bottom: 1rem;
    color: var(--text);
    font-weight: 700;
    line-height: 1.4;
  }

  .excerpt {
    color: var(--text-secondary);
    line-height: 1.7;
    margin-bottom: 1.5rem;
    flex-grow: 1;
  }

  .read-more {
    color: var(--primary);
    font-weight: 600;
    font-size: 0.9375rem;
    display: inline-flex;
    align-items: center;
    transition: gap 0.2s ease;
    gap: 0.5rem;
  }

  .news-card:hover .read-more {
    gap: 0.75rem;
  }

  .loading {
    text-align: center;
    padding: 4rem 0;
  }

  .spinner {
    width: 48px;
    height: 48px;
    border: 4px solid var(--border);
    border-top-color: var(--primary);
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
    margin: 0 auto 1rem;
  }

  @keyframes spin {
    to { transform: rotate(360deg); }
  }

  .empty-state {
    text-align: center;
    padding: 4rem 0;
    color: var(--text-muted);
  }

  @media (max-width: 768px) {
    .hero {
      padding: 4rem 0;
    }

    .hero h1 {
      font-size: 2.25rem;
    }

    .hero p {
      font-size: 1.125rem;
    }

    .section-header {
      flex-direction: column;
      align-items: flex-start;
    }
  }
</style>
