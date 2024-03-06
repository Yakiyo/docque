<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import Home from "./lib/components/Home.svelte";
</script>

<Router url="">
  <div>
    <Route path="/" component={Home}></Route>
  </div>
</Router>
