<script>
  const currentYear = new Date().getFullYear()
</script>

<footer>
  <div class="container">
    <div class="footer-content">
      <div class="footer-section">
        <h3>Департамент финансов</h3>
        <p>Окружной администрации города Якутска</p>
      </div>

      <div class="footer-section">
        <h4>Контакты</h4>
        <p>Телефон: +7 (4112) 42-26-26</p>
        <p>Email: depfin@yakutsk.org</p>
      </div>

      <div class="footer-section">
        <h4>Адрес</h4>
        <p>677000, Республика Саха (Якутия)</p>
        <p>г. Якутск, ул. Орджоникидзе, д. 24</p>
      </div>
    </div>

    <div class="footer-bottom">
      <p>&copy; {currentYear} Департамент финансов города Якутска</p>
    </div>
  </div>
</footer>

<style>
  footer {
    background: var(--surface);
    border-top: 1px solid var(--border);
    margin-top: 4rem;
    padding: 2rem 0 1rem;
  }

  .footer-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    margin-bottom: 2rem;
  }

  .footer-section h3 {
    font-size: 1.125rem;
    margin-bottom: 0.5rem;
    color: var(--text);
  }

  .footer-section h4 {
    font-size: 1rem;
    margin-bottom: 0.5rem;
    color: var(--text);
  }

  .footer-section p {
    color: var(--text-secondary);
    font-size: 0.875rem;
    line-height: 1.6;
  }

  .footer-bottom {
    padding-top: 1rem;
    border-top: 1px solid var(--border);
    text-align: center;
  }

  .footer-bottom p {
    color: var(--text-secondary);
    font-size: 0.875rem;
  }
</style>
