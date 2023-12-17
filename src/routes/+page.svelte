<script lang="ts">
	export let data;
    // sort appointments based on doctor
	const appointments = data.appointments.sort((x, y) =>
		x.doctor.name! > y.doctor.name! ? 1 : x.doctor.name! < y.doctor.name! ? -1 : 0
	);
    
</script>

<svelte:head>
	<title>Docque - Doctor's Appointment Queue</title>
</svelte:head>
