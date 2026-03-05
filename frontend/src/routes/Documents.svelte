<script>
  import { onMount } from 'svelte'
  import { api } from '../lib/api.js'

  let documents = []
  let loading = true
  let selectedCategory = ''

  const categories = [
    { value: '', label: 'Все категории' },
    { value: 'budget', label: 'Бюджет' },
    { value: 'reports', label: 'Отчеты' },
    { value: 'regulations', label: 'Нормативные акты' },
    { value: 'orders', label: 'Приказы' }
  ]

  onMount(() => {
    loadDocuments()
  })

  async function loadDocuments() {
    loading = true
    try {
      documents = await api.getDocuments(selectedCategory)
    } catch (error) {
      console.error('Failed to load documents:', error)
    } finally {
      loading = false
    }
  }

  function handleCategoryChange(event) {
    selectedCategory = event.target.value
    loadDocuments()
  }
</script>

<section class="section">
  <div class="container">
    <div class="page-header">
      <h1 class="page-title">Документы и нормативные акты</h1>
      <p class="page-description">Официальные документы, нормативные акты и приказы департамента финансов</p>
    </div>

    <div class="filters">
      <div class="filter-group">
        <label for="category">Категория:</label>
        <select id="category" bind:value={selectedCategory} on:change={handleCategoryChange}>
          {#each categories as cat}
            <option value={cat.value}>{cat.label}</option>
          {/each}
        </select>
      </div>
    </div>

    {#if loading}
      <div class="loading">
        <div class="spinner"></div>
        <p>Загрузка документов...</p>
      </div>
    {:else if documents.length > 0}
      <div class="documents-list">
        {#each documents as doc}
          <div class="card document-item">
            <div class="document-icon">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z" />
              </svg>
            </div>
            <div class="document-content">
              <h3>{doc.title}</h3>
              {#if doc.description}
                <p class="description">{doc.description}</p>
              {/if}
              <div class="document-meta">
                <span class="badge badge-primary">{doc.category}</span>
                <time>{new Date(doc.published_at).toLocaleDateString('ru-RU', {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric'
                })}</time>
              </div>
            </div>
            <a href={doc.file_url} class="btn btn-primary" download>
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3" />
              </svg>
              Скачать
            </a>
          </div>
        {/each}
      </div>
    {:else}
      <div class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z" />
        </svg>
        <p>Документов не найдено</p>
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

  .filters {
    margin-bottom: 3rem;
    display: flex;
    gap: 1.5rem;
    flex-wrap: wrap;
  }

  .filter-group {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .filter-group label {
    font-weight: 600;
    color: var(--text);
    font-size: 0.9375rem;
  }

  select {
    padding: 0.75rem 1.25rem;
    border: 2px solid var(--border);
    border-radius: 12px;
    background: var(--surface);
    color: var(--text);
    font-size: 0.9375rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    min-width: 200px;
  }

  select:hover {
    border-color: var(--primary);
  }

  select:focus {
    outline: none;
    border-color: var(--primary);
    box-shadow: 0 0 0 3px rgba(8, 145, 178, 0.1);
  }

  .documents-list {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .document-item {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  .document-icon {
    width: 56px;
    height: 56px;
    border-radius: 14px;
    background: rgba(8, 145, 178, 0.1);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .document-icon svg {
    width: 28px;
    height: 28px;
    color: var(--primary);
  }

  .document-content {
    flex: 1;
  }

  .document-content h3 {
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
    color: var(--text);
    font-weight: 700;
    line-height: 1.4;
  }

  .description {
    color: var(--text-secondary);
    font-size: 0.9375rem;
    margin-bottom: 0.75rem;
    line-height: 1.6;
  }

  .document-meta {
    display: flex;
    gap: 1rem;
    align-items: center;
    flex-wrap: wrap;
  }

  .document-meta time {
    color: var(--text-muted);
    font-size: 0.875rem;
    font-weight: 500;
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

    .document-item {
      flex-direction: column;
      align-items: flex-start;
    }

    .document-item .btn {
      width: 100%;
    }
  }
</style>
