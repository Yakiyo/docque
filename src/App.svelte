<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import Home from "./lib/components/Home.svelte";
  import DocPage from "./lib/components/DocPage.svelte";
  import Navbar from "./lib/components/Navbar.svelte";
</script>

<Navbar></Navbar>

<Router url="">
  <div>
    <Route path="/" component={Home}></Route>
    <Route path="/doc/:id" let:params>
      <DocPage id={params.id}></DocPage>
    </Route>
  </div>
</Router>
