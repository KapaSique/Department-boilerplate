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
          <div class="logo-icon">💼</div>
          <div class="logo-text">
            <h1>Департамент финансов</h1>
            <p>города Якутска</p>
          </div>
        </Link>
      </div>

      <button class="mobile-toggle" on:click={toggleMenu} aria-label="Toggle menu">
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
    border-bottom: 1px solid var(--border);
    position: sticky;
    top: 0;
    z-index: 100;
    box-shadow: var(--shadow-sm);
    backdrop-filter: blur(12px);
    background: rgba(255, 255, 255, 0.95);
  }

  .header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1.25rem 0;
  }

  .logo :global(a) {
    display: flex;
    align-items: center;
    gap: 1rem;
    text-decoration: none;
  }

  .logo-icon {
    font-size: 2.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 56px;
    height: 56px;
    background: var(--gradient-primary);
    border-radius: 14px;
    box-shadow: 0 4px 12px rgba(8, 145, 178, 0.2);
  }

  .logo-text h1 {
    font-size: 1.375rem;
    font-weight: 700;
    color: var(--text);
    letter-spacing: -0.02em;
    line-height: 1.2;
  }

  .logo-text p {
    font-size: 0.875rem;
    color: var(--text-secondary);
    font-weight: 500;
    margin-top: 0.125rem;
  }

  .logo :global(a:focus-visible) {
    outline: 3px solid var(--focus-ring);
    outline-offset: 4px;
    border-radius: 8px;
  }

  .mobile-toggle {
    display: none;
    flex-direction: column;
    gap: 5px;
    padding: 0.75rem;
    border-radius: 8px;
    background: transparent;
    transition: background 0.2s;
  }

  .mobile-toggle:hover {
    background: var(--border-light);
  }

  .mobile-toggle span {
    width: 24px;
    height: 2px;
    background: var(--text);
    transition: all 0.3s;
    border-radius: 2px;
  }

  nav {
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  nav :global(a) {
    color: var(--text-secondary);
    font-weight: 600;
    font-size: 0.9375rem;
    transition: all 0.2s ease;
    padding: 0.625rem 1rem;
    border-radius: 10px;
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
      padding: 1.5rem;
      border-bottom: 1px solid var(--border);
      display: none;
      box-shadow: var(--shadow-lg);
      gap: 0.5rem;
    }

    nav.open {
      display: flex;
    }

    nav :global(a) {
      width: 100%;
      text-align: left;
      padding: 0.875rem 1rem;
    }
  }
</style>
