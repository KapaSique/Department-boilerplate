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
    padding: 4rem 0;
    text-align: center;
  }

  .hero h1 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
  }

  .hero p {
    font-size: 1.25rem;
    opacity: 0.9;
  }

  .feature {
    text-align: center;
  }

  .icon {
    font-size: 3rem;
    margin-bottom: 1rem;
  }

  .feature h3 {
    font-size: 1.5rem;
    margin-bottom: 0.5rem;
  }

  .feature p {
    color: var(--text-secondary);
    margin-bottom: 1rem;
  }

  .news-card h3 {
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
    color: var(--text);
  }

  .date {
    font-size: 0.875rem;
    color: var(--text-secondary);
    margin-bottom: 0.5rem;
  }

  .excerpt {
    color: var(--text-secondary);
    line-height: 1.6;
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
