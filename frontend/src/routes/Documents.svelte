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
    <h1 class="section-title">Документы и нормативные акты</h1>

    <div class="filters">
      <select bind:value={selectedCategory} on:change={handleCategoryChange}>
        {#each categories as cat}
          <option value={cat.value}>{cat.label}</option>
        {/each}
      </select>
    </div>

    {#if loading}
      <p>Загрузка документов...</p>
    {:else if documents.length > 0}
      <div class="documents-list">
        {#each documents as doc}
          <div class="card document-item">
            <div class="doc-icon">📄</div>
            <div class="doc-content">
              <h3>{doc.title}</h3>
              {#if doc.description}
                <p class="description">{doc.description}</p>
              {/if}
              <div class="doc-meta">
                <span class="category">{doc.category}</span>
                <time>{new Date(doc.published_at).toLocaleDateString('ru-RU')}</time>
              </div>
            </div>
            <a href={doc.file_url} class="btn btn-primary" download>Скачать</a>
          </div>
        {/each}
      </div>
    {:else}
      <div class="card">
        <p>Документов не найдено</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .filters {
    margin-bottom: 2rem;
  }

  select {
    padding: 0.75rem 1rem;
    border: 1px solid var(--border);
    border-radius: 6px;
    background: var(--surface);
    color: var(--text);
    font-size: 1rem;
    cursor: pointer;
  }

  .documents-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .document-item {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  .doc-icon {
    font-size: 2rem;
  }

  .doc-content {
    flex: 1;
  }

  .doc-content h3 {
    font-size: 1.125rem;
    margin-bottom: 0.5rem;
    color: var(--text);
  }

  .description {
    color: var(--text-secondary);
    font-size: 0.875rem;
    margin-bottom: 0.5rem;
  }

  .doc-meta {
    display: flex;
    gap: 1rem;
    font-size: 0.875rem;
  }

  .category {
    color: var(--primary);
    font-weight: 500;
  }

  time {
    color: var(--text-secondary);
  }

  @media (max-width: 768px) {
    .document-item {
      flex-direction: column;
      align-items: flex-start;
    }
  }
</style>
