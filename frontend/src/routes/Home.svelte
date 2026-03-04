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
    <h1>Департамент финансов города Якутска</h1>
    <p>Управление бюджетом и финансами муниципального образования</p>
  </div>
</div>

<section class="section">
  <div class="container">
    <div class="grid grid-3">
      <div class="card feature">
        <div class="icon">📰</div>
        <h3>Новости</h3>
        <p>Актуальные новости и объявления департамента</p>
        <Link to="/news" class="btn btn-primary">Перейти</Link>
      </div>

      <div class="card feature">
        <div class="icon">💰</div>
        <h3>Бюджет</h3>
        <p>Бюджетные данные и финансовые отчеты</p>
        <Link to="/budget" class="btn btn-primary">Перейти</Link>
      </div>

      <div class="card feature">
        <div class="icon">📋</div>
        <h3>Документы</h3>
        <p>Нормативные акты и официальные документы</p>
        <Link to="/documents" class="btn btn-primary">Перейти</Link>
      </div>
    </div>
  </div>
</section>

<section class="section">
  <div class="container">
    <h2 class="section-title">Последние новости</h2>

    {#if loading}
      <p>Загрузка...</p>
    {:else if news.length > 0}
      <div class="grid grid-2">
        {#each news as item}
          <Link to="/news/{item.id}">
            <div class="card news-card">
              <h3>{item.title}</h3>
              <p class="date">{new Date(item.published_at).toLocaleDateString('ru-RU')}</p>
              <p class="excerpt">{item.content.substring(0, 150)}...</p>
            </div>
          </Link>
        {/each}
      </div>
      <div style="text-align: center; margin-top: 2rem;">
        <Link to="/news" class="btn btn-primary">Все новости</Link>
      </div>
    {:else}
      <p>Новостей пока нет</p>
    {/if}
  </div>
</section>

<style>
  .hero {
    background: linear-gradient(135deg, var(--primary) 0%, var(--primary-dark) 100%);
    color: white;
    padding: 5rem 0;
    text-align: center;
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
    background: radial-gradient(circle at 30% 50%, rgba(34, 211, 238, 0.2) 0%, transparent 50%);
    pointer-events: none;
  }

  .hero h1 {
    font-size: 2.75rem;
    margin-bottom: 1.25rem;
    font-weight: 700;
    letter-spacing: -0.02em;
    position: relative;
  }

  .hero p {
    font-size: 1.375rem;
    opacity: 0.95;
    max-width: 700px;
    margin: 0 auto;
    line-height: 1.6;
    position: relative;
  }

  .feature {
    text-align: center;
    transition: transform 0.25s ease;
  }

  .feature:hover {
    transform: translateY(-4px);
  }

  .icon {
    font-size: 3.5rem;
    margin-bottom: 1.25rem;
    filter: drop-shadow(0 2px 8px rgba(8, 145, 178, 0.15));
  }

  .feature h3 {
    font-size: 1.5rem;
    margin-bottom: 0.75rem;
    color: var(--text);
    font-weight: 600;
  }

  .feature p {
    color: var(--text-secondary);
    margin-bottom: 1.5rem;
    line-height: 1.6;
  }

  .news-card {
    transition: all 0.25s ease;
  }

  .news-card:hover {
    border-color: var(--secondary);
  }

  .news-card h3 {
    font-size: 1.25rem;
    margin-bottom: 0.75rem;
    color: var(--text);
    font-weight: 600;
    line-height: 1.4;
  }

  .date {
    font-size: 0.875rem;
    color: var(--text-secondary);
    margin-bottom: 0.75rem;
    font-weight: 500;
  }

  .excerpt {
    color: var(--text-secondary);
    line-height: 1.7;
  }

  @media (max-width: 768px) {
    .hero h1 {
      font-size: 1.75rem;
    }

    .hero p {
      font-size: 1rem;
    }
  }
</style>
