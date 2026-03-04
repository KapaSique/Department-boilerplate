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
    <h1 class="section-title">Новости и объявления</h1>

    {#if loading}
      <p>Загрузка новостей...</p>
    {:else if news.length > 0}
      <div class="news-list">
        {#each news as item}
          <Link to="/news/{item.id}">
            <article class="card news-item">
              <div class="news-header">
                <h2>{item.title}</h2>
                <time>{new Date(item.published_at).toLocaleDateString('ru-RU', {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric'
                })}</time>
              </div>
              <p class="news-excerpt">{item.content.substring(0, 200)}...</p>
              <span class="read-more">Читать далее →</span>
            </article>
          </Link>
        {/each}
      </div>
    {:else}
      <div class="card">
        <p>Новостей пока нет</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .news-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .news-item {
    transition: all 0.2s;
  }

  .news-item:hover {
    transform: translateY(-2px);
  }

  .news-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 1rem;
    margin-bottom: 1rem;
  }

  .news-header h2 {
    font-size: 1.5rem;
    color: var(--text);
    flex: 1;
  }

  time {
    font-size: 0.875rem;
    color: var(--text-secondary);
    white-space: nowrap;
  }

  .news-excerpt {
    color: var(--text-secondary);
    line-height: 1.6;
    margin-bottom: 1rem;
  }

  .read-more {
    color: var(--primary);
    font-weight: 500;
  }

  @media (max-width: 768px) {
    .news-header {
      flex-direction: column;
    }

    .news-header h2 {
      font-size: 1.25rem;
    }
  }
</style>
