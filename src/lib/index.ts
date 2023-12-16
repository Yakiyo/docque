import type { Appointment } from '@prisma/client';

/**
 * Calculates and adds duration of the appointment and returns it
 */
export function addDuration<T extends Pick<Appointment, 'start' | 'end'>>(
	a: T
): T & { duration: number } {
	return {
		duration: Math.abs(a.end.getTime() - a.start.getTime()),
		...a
	};
}

export * from './db';
