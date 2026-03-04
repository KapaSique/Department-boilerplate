<script>
  import { Link } from 'svelte-routing'

  let mobileMenuOpen = false

  function toggleMenu() {
    mobileMenuOpen = !mobileMenuOpen
  }
</script>

<header>
  <div class="container">
    <div class="header-content">
      <div class="logo">
        <Link to="/">
          <h1>Департамент финансов</h1>
          <p>города Якутска</p>
        </Link>
      </div>

      <button class="mobile-toggle" on:click={toggleMenu}>
        <span></span>
        <span></span>
        <span></span>
      </button>

      <nav class:open={mobileMenuOpen}>
        <Link to="/" on:click={() => mobileMenuOpen = false}>Главная</Link>
        <Link to="/news" on:click={() => mobileMenuOpen = false}>Новости</Link>
        <Link to="/documents" on:click={() => mobileMenuOpen = false}>Документы</Link>
        <Link to="/budget" on:click={() => mobileMenuOpen = false}>Бюджет</Link>
        <Link to="/about" on:click={() => mobileMenuOpen = false}>О департаменте</Link>
        <Link to="/contacts" on:click={() => mobileMenuOpen = false}>Контакты</Link>
      </nav>
    </div>
  </div>
</header>

<style>
  header {
    background: var(--surface);
    border-bottom: 2px solid var(--border);
    position: sticky;
    top: 0;
    z-index: 100;
    box-shadow: 0 2px 12px rgba(8, 145, 178, 0.08);
    backdrop-filter: blur(8px);
  }

  .header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 0;
  }

  .logo h1 {
    font-size: 1.375rem;
    font-weight: 700;
    color: var(--text);
    letter-spacing: -0.01em;
  }

  .logo p {
    font-size: 0.875rem;
    color: var(--text-secondary);
    font-weight: 500;
  }

  .logo :global(a:focus-visible) {
    outline: 3px solid var(--focus-ring);
    outline-offset: 4px;
    border-radius: 4px;
  }

  .mobile-toggle {
    display: none;
    flex-direction: column;
    gap: 4px;
    padding: 0.5rem;
  }

  .mobile-toggle span {
    width: 24px;
    height: 2px;
    background: var(--text);
    transition: all 0.3s;
  }

  nav {
    display: flex;
    gap: 2rem;
  }

  nav :global(a) {
    color: var(--text-secondary);
    font-weight: 600;
    font-family: 'Lexend', sans-serif;
    transition: all 0.2s ease;
    padding: 0.5rem 0.75rem;
    border-radius: 6px;
    position: relative;
  }

  nav :global(a:hover) {
    color: var(--primary);
    background: rgba(8, 145, 178, 0.08);
  }

  nav :global(a.active) {
    color: var(--primary);
    background: rgba(8, 145, 178, 0.12);
  }

  nav :global(a:focus-visible) {
    outline: 3px solid var(--focus-ring);
    outline-offset: 2px;
  }

  @media (max-width: 768px) {
    .mobile-toggle {
      display: flex;
    }

    nav {
      position: absolute;
      top: 100%;
      left: 0;
      right: 0;
      background: var(--surface);
      flex-direction: column;
      padding: 1rem;
      border-bottom: 1px solid var(--border);
      display: none;
    }

    nav.open {
      display: flex;
    }
  }
</style>
