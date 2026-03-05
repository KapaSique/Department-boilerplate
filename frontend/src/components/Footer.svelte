<script>
  const currentYear = new Date().getFullYear()
</script>

<footer>
  <div class="container">
    <div class="footer-content">
      <div class="footer-section">
        <div class="footer-logo">
          <div class="logo-icon">💼</div>
          <div>
            <h3>Департамент финансов</h3>
            <p>Окружной администрации города Якутска</p>
          </div>
        </div>
      </div>

      <div class="footer-section">
        <h4>Контакты</h4>
        <ul>
          <li>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 01-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 00-1.091-.852H4.5A2.25 2.25 0 002.25 4.5v2.25z" />
            </svg>
            +7 (4112) 42-26-26
          </li>
          <li>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
            </svg>
            depfin@yakutsk.org
          </li>
        </ul>
      </div>

      <div class="footer-section">
        <h4>Адрес</h4>
        <ul>
          <li>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
            </svg>
            677000, Республика Саха (Якутия)
          </li>
          <li>г. Якутск, ул. Орджоникидзе, д. 24</li>
          <li class="work-hours">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            Пн-Пт: 9:00-18:00, обед 13:00-14:00
          </li>
        </ul>
      </div>
    </div>

    <div class="footer-bottom">
      <p>&copy; {currentYear} Департамент финансов города Якутска. Все права защищены.</p>
      <div class="footer-links">
        <a href="/about">О департаменте</a>
        <a href="/contacts">Контакты</a>
      </div>
    </div>
  </div>
</footer>

<style>
  footer {
    background: var(--text);
    color: rgba(255, 255, 255, 0.9);
    margin-top: 6rem;
    padding: 4rem 0 2rem;
  }

  .footer-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 3rem;
    margin-bottom: 3rem;
  }

  .footer-logo {
    display: flex;
    align-items: flex-start;
    gap: 1rem;
  }

  .logo-icon {
    font-size: 2rem;
    width: 48px;
    height: 48px;
    background: var(--gradient-primary);
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .footer-section h3 {
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
    color: white;
    font-weight: 700;
  }

  .footer-section h4 {
    font-size: 1rem;
    margin-bottom: 1.25rem;
    color: white;
    font-weight: 700;
    letter-spacing: -0.01em;
  }

  .footer-section p {
    color: rgba(255, 255, 255, 0.7);
    font-size: 0.9375rem;
    line-height: 1.6;
  }

  .footer-section ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .footer-section li {
    color: rgba(255, 255, 255, 0.7);
    font-size: 0.9375rem;
    line-height: 1.8;
    display: flex;
    align-items: flex-start;
    gap: 0.75rem;
  }

  .footer-section li svg {
    width: 20px;
    height: 20px;
    flex-shrink: 0;
    margin-top: 0.125rem;
    color: var(--secondary);
  }

  .work-hours {
    margin-top: 0.75rem;
    padding-top: 0.75rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
  }

  .footer-bottom {
    padding-top: 2rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    gap: 1.5rem;
  }

  .footer-bottom p {
    color: rgba(255, 255, 255, 0.6);
    font-size: 0.875rem;
    margin: 0;
  }

  .footer-links {
    display: flex;
    gap: 2rem;
  }

  .footer-links a {
    color: rgba(255, 255, 255, 0.7);
    font-size: 0.875rem;
    font-weight: 500;
    transition: color 0.2s;
  }

  .footer-links a:hover {
    color: white;
  }

  @media (max-width: 768px) {
    footer {
      padding: 3rem 0 1.5rem;
    }

    .footer-content {
      gap: 2.5rem;
    }

    .footer-bottom {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;
    }

    .footer-links {
      gap: 1.5rem;
    }
  }
</style>
