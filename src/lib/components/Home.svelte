<script lang="ts">
  import { onMount } from "svelte";
  import { pb } from "../pocketbase";

  let apps: Record<string, string>[] = [];
  let loaded = false;
  onMount(async () => {
	apps = await pb.collection('appointments').getFullList({
		sort: '-start',
	})
	for (const app of apps) {
		app.doctor = await pb.collection('doctors').getOne(app.doctor).then(v => v.name)
		const date = new Date(Date.parse(apps[0].start))
		const datestr = `${date.getMonth()}/${date.getDate()}/${date.getFullYear()} - ${date.getHours()}`
		console.log('str', datestr);
	}
	console.log(Date.parse(apps[0].start))
	loaded = true
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
          <th>Time</th>
        </tr>
      </thead>
	  <tbody>
		{#each apps as app, i}
			<tr class="hover">
				<td>{i + 1}</td>
				<td>{app.patient}</td>
				<td>{app.doctor}</td>
				<td>{app.start}</td>
			</tr>
		{/each}
	  </tbody>
    </table>
  </div>
{/if}
