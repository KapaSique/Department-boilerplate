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
      <p>Загрузка...</p>
    {:else if error}
      <div class="card">
        <p>{error}</p>
      </div>
    {:else if news}
      <article class="news-detail">
        <h1>{news.title}</h1>
        <time>{new Date(news.published_at).toLocaleDateString('ru-RU', {
          year: 'numeric',
          month: 'long',
          day: 'numeric'
        })}</time>
        <div class="content">
          {news.content}
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

  h1 {
    font-size: 2rem;
    margin-bottom: 1rem;
    color: var(--text);
  }

  time {
    display: block;
    font-size: 0.875rem;
    color: var(--text-secondary);
    margin-bottom: 2rem;
  }

  .content {
    line-height: 1.8;
    color: var(--text);
    white-space: pre-wrap;
  }

  @media (max-width: 768px) {
    h1 {
      font-size: 1.5rem;
    }
  }
</style>
