<script lang="ts">
	import Queue from '../components/Queue.svelte';
	import type { AppointmentResponse } from '$lib';

	export let data;
	// sort appointments based on doctor
	const { appointments } = data;
	const docQueues: Record<string, AppointmentResponse[]> = {};
	appointments.forEach((x) => {
		if (!(x.doctor.name in docQueues)) {
			docQueues[x.doctor.name] = [];
		}
		docQueues[x.doctor.name].push(x);
	});
</script>

<svelte:head>
	<title>Docque - Doctor's Appointment Queue</title>
</svelte:head>

<div class="flex w-full flex-col items-center justify-between">
	{#each Object.entries(docQueues) as [doctor, appointments]}
		<div class="py-3">
			<Queue {doctor} {appointments} />
		</div>
	{/each}
</div>
