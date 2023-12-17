import type { Appointment, Doctor } from '@prisma/client';

/**
 * Calculates and adds duration of the appointment and returns it
 */
export function addDuration<T extends Pick<Appointment, 'start' | 'end'>>(
	a: T
): T & { duration: number } {
	return {
		duration: Math.abs((a.end.getTime() - a.start.getTime()) / 60000),
		...a
	};
}

export type AppointmentResponse = Appointment & {
	doctor: Doctor;
	duration: number;
};

export type ApiResponse<T> = {
	ok: boolean;
	error: string | null;
	data: T;
};

export * from './db';
