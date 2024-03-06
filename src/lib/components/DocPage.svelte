<script lang="ts">
  import { onMount } from "svelte";
  import { pb } from "../pocketbase";
  import { modifyApp } from "../utils";

  export let id: string;
  let loaded = false;
  let apps: Record<string, string>[] = [];
  onMount(async () => {
    const apps = (await pb.collection("appointments").getFullList({
      filter: `doctor.id='${id}'`,
    })) as Record<string, string>[];
    console.log(apps.length);
    for (let app of apps) {
      app = await modifyApp(app);
    }
    console.log(apps, apps.length);

    loaded = true;
  });
  console.log(apps.length)
</script>

{#if !loaded}
  <div class="flex justify-center h-max items-center">
    <!-- <img src="/loader.svg" alt="loader-gif"> -->
    <span class="loading loading-ring loading-lg"></span>
  </div>
{:else}
  <div class="overflow-x-auto p-4 m-7">
    <table class="table border-2">
      <thead>
        <tr>
          <th></th>
          <th>Patient</th>
          <th>Duration</th>
          <th>Time</th>
        </tr>
      </thead>
      <tbody>
        {apps.length}
        {#each apps as app, i}
          <tr class="hover">
            <td>{i + 1}</td>
            <td>{app.patient}</td>
            <td>{app.duration}</td>
            <td>{app.start}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
{/if}
