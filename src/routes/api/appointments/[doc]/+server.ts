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

/**
 * Adds a new appointment for a doctor
 */
export async function POST({ params: { doc }, request }) {
	// NOTE: duration is and MUST be in minutes
	const { start, duration, patient } = (await request.json()) as Record<string, string>;
	if ([start, duration, patient].some((x) => !x)) {
		return json({
			ok: false,
			data: null,
			error: 'Missing required property. Must provide start, duration and patient'
		});
	}

	const doctor = await prisma.doctor.findUnique({
		where: {
			name: doc
		},
		select: {
			id: true
		}
	});
	if (!doctor) {
		return json({
			ok: false,
			data: null,
			error: `Invalid doctor name. No entry found for ${doc}`
		});
	}
	// TODO: make sure theres no existing in the same time and duration
	const err = await prisma.appointment
		.create({
			data: {
				start: new Date(start),
				end: new Date(new Date(start).getTime() + Number(duration) * 60000),
				patient,
				doctorId: doctor.id
			}
		})
		.then(() => null)
		.catch((e) => `${e}`);

	return json({
		ok: !err ? true : false,
		data: null,
		error: err
	});
}
