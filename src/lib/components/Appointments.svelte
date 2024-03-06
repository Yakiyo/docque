<!-- 
    @component
    A svelte component that fetches all appointment's and displays them in a table
 -->
<script lang="ts">
    import { onMount } from "svelte";
    import { pb } from "../pocketbase";
  import { modifyApp } from "../utils";
  
    let apps: Record<string, string>[] = [];
    let loaded = false;
    onMount(async () => {
      apps = await pb.collection("appointments").getFullList({
        sort: "-start",
      });
      for (let app of apps) {
        app = await modifyApp(app);
      }
      console.log(apps.map((x) => x.start + "\n"));
      loaded = true;
    });
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
            <th>Doctor</th>
            <th>Duration</th>
            <th>Time</th>
          </tr>
        </thead>
        <tbody>
          {#each apps as app, i}
            <tr class="hover">
              <td>{i + 1}</td>
              <td>{app.patient}</td>
              <td class="underline"><a href={`/doc/${app.docId}`}>Dr. {app.doctor}</a></td>
              <td>{app.duration}</td>
              <td>{app.start}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
  