import { prisma } from '$lib';
import { json } from '@sveltejs/kit';

/**
 * Fetch all appointments of a doctor
 */
export async function GET() {
	const data = await prisma.appointment
		.findMany({
			select: {
				id: true,
				start: true,
				end: true,
				doctor: true
			},
			orderBy: {
				doctor: {
					name: 'asc'
				}
			}
		})
		.then((a) =>
			a.map((a) => ({
				duration: Math.abs(a.end.getTime() - a.start.getTime()),
				...a
			}))
		)
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
