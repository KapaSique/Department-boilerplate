<script>
  import { Router, Route } from 'svelte-routing'
  import Header from './components/Header.svelte'
  import Footer from './components/Footer.svelte'
  import Home from './routes/Home.svelte'
  import News from './routes/News.svelte'
  import NewsDetail from './routes/NewsDetail.svelte'
  import Documents from './routes/Documents.svelte'
  import Budget from './routes/Budget.svelte'
  import About from './routes/About.svelte'
  import Contacts from './routes/Contacts.svelte'
</script>

<Router>
  <Header />

  <main>
    <Route path="/" component={Home} />
    <Route path="/news" component={News} />
    <Route path="/news/:id" component={NewsDetail} />
    <Route path="/documents" component={Documents} />
    <Route path="/budget" component={Budget} />
    <Route path="/about" component={About} />
    <Route path="/contacts" component={Contacts} />
  </main>

  <Footer />
</Router>

<style>
  main {
    min-height: calc(100vh - 200px);
  }
</style>
