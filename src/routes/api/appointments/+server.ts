import { addDuration, prisma } from '$lib';
import { json } from '@sveltejs/kit';

/**
 * Fetch all appointments of a doctor
 */
export async function GET() {
	const data = await prisma.appointment
		.findMany({
			include: {
				doctor: true,
			},
			orderBy: {
				doctor: {
					name: 'asc'
				}
			}
		})
		.then((a) => a.map(addDuration))
		.catch((e) => `${e}`);

	if (typeof data === 'string') {
		return json({
			ok: false,
			error: data,
			data: []
		});
	}

	return json({
		ok: true,
		error: null,
		data
	});
}
