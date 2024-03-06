<script>
  import { link } from "svelte-routing";
    import { pb } from "../pocketbase";

  const logged = !!pb.authStore.token;
</script>

<div class="navbar bg-base-100 shadow-lg rounded-lg">
  <div class="navbar-start">
    <div class="dropdown">
      <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          ><path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h7"
          /></svg
        >
      </div>
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li><a href="https://github.com/Yakiyo/docque">Github</a></li>
      </ul>
    </div>
  </div>
  <div class="flex-1">
    <a class="btn btn-ghost text-xl" href="/" use:link>Docque</a>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal px-1">
      {#if !logged}
        <li><a href="/login" use:link><img src="/images/user.svg" alt="user-logo">Login</a></li>
      {/if}
    </ul>
  </div>
</div>
