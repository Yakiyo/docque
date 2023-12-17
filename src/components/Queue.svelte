<script lang="ts">
	import type { AppointmentResponse } from '$lib';
	export let doctor: string;
	export let appointments: AppointmentResponse[];
	// console.log(appointments.find(a => !a.start), doctor)

	// appointments.sort((a, b) => a.start.getTime() - b.start.getTime());
</script>

<table border="25" class="table-auto border-spacing-5">
	<caption class="text-lg font-bold"><a href={`/${doctor}`}>Appointments for Dr. {doctor}</a></caption>
	<tr>
		<th>Serial</th>
		<th>Patient</th>
		<th>Estimated Time</th>
		<th>Estimated Duration</th>
	</tr>
	{#if appointments.length < 1}
		<tr>
			<td colspan="4">Dr. {doctor} currently has no appointments</td>
		</tr>
	{/if}
	{#each appointments as appointment, i}
		<tr>
			<td>{i + 1}</td>
			<td>{appointment.patient}</td>
			<td>{appointment.start}</td>
			<td>{appointment.duration}</td>
		</tr>
	{/each}
</table>

<style>
	td {
		text-align: center;
		padding: 15px;
	}
	table,
	th,
	td {
		border: 1px solid;
	}
</style>
