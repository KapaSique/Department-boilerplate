<script>
  import { onMount } from 'svelte'
  import { api } from '../lib/api.js'

  let contacts = []
  let loading = true

  onMount(async () => {
    try {
      contacts = await api.getContacts()
    } catch (error) {
      console.error('Failed to load contacts:', error)
    } finally {
      loading = false
    }
  })
</script>

<section class="section">
  <div class="container">
    <h1 class="section-title">Контакты</h1>

    {#if loading}
      <p>Загрузка...</p>
    {:else if contacts.length > 0}
      <div class="contacts-grid">
        {#each contacts as contact}
          <div class="card contact-card">
            <h3>{contact.department}</h3>
            <div class="contact-info">
              <div class="contact-item">
                <span class="icon">👤</span>
                <div>
                  <strong>Руководитель</strong>
                  <p>{contact.head}</p>
                </div>
              </div>
              <div class="contact-item">
                <span class="icon">📞</span>
                <div>
                  <strong>Телефон</strong>
                  <p><a href="tel:{contact.phone}">{contact.phone}</a></p>
                </div>
              </div>
              <div class="contact-item">
                <span class="icon">✉️</span>
                <div>
                  <strong>Email</strong>
                  <p><a href="mailto:{contact.email}">{contact.email}</a></p>
                </div>
              </div>
            </div>
          </div>
        {/each}
      </div>
    {:else}
      <div class="card">
        <p>Контактная информация недоступна</p>
      </div>
    {/if}

    <div class="card map-placeholder">
      <h3>Как нас найти</h3>
      <p>677000, Республика Саха (Якутия), г. Якутск, ул. Орджоникидзе, д. 24</p>
      <p class="note">Режим работы: Пн-Пт 9:00-18:00, обед 13:00-14:00</p>
    </div>
  </div>
</section>

<style>
  .contacts-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
  }

  .contact-card h3 {
    font-size: 1.25rem;
    margin-bottom: 1.5rem;
    color: var(--text);
  }

  .contact-info {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .contact-item {
    display: flex;
    gap: 1rem;
    align-items: flex-start;
  }

  .icon {
    font-size: 1.5rem;
  }

  .contact-item strong {
    display: block;
    font-size: 0.875rem;
    color: var(--text);
    margin-bottom: 0.25rem;
  }

  .contact-item p {
    color: var(--text-secondary);
    font-size: 0.875rem;
  }

  .contact-item a {
    color: var(--primary);
  }

  .map-placeholder {
    margin-top: 2rem;
  }

  .map-placeholder h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
    color: var(--text);
  }

  .map-placeholder p {
    color: var(--text-secondary);
    line-height: 1.6;
  }

  .note {
    margin-top: 0.5rem;
    font-size: 0.875rem;
    font-style: italic;
  }
</style>
