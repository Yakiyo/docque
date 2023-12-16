import { prisma } from '$lib/db';
import { addDuration } from '$lib/index';
import { json } from '@sveltejs/kit';

/**
 * Return all appointments of a doctor
 */
export async function GET({ params: { doc } }) {
	const doctor = await prisma.doctor.findUnique({
		where: {
			name: doc
		},
		include: {
			appointments: true
		}
	});
	if (!doctor) {
		return json({
			ok: false,
			error: `Invalid doc name: ${doc}. No entry found`,
			data: []
		});
	}
	return json({
		ok: true,
		data: doctor.appointments.map(addDuration),
		error: null
	});
}
