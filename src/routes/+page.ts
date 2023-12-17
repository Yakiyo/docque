import type { ApiResponse, AppointmentResponse } from '$lib';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const response = (await fetch('/api/appointments').then((r) => r.json())) as ApiResponse<
		AppointmentResponse[]
	>;
	if (!response.ok) {
		error(400, response.error!);
	}
	return {
		appointments: response.data
	};
};
