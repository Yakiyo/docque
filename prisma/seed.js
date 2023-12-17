// @ts-check

import { PrismaClient } from '@prisma/client';
import { faker } from '@faker-js/faker';

const prisma = new PrismaClient();

async function generateDoctor() {
	const doctor = await prisma.doctor.create({
		data: {
			name: faker.person.fullName()
		}
	});
	// generate n number of appointment instances
	const n = faker.number.int({ min: 5, max: 20 });
	console.info(`Generating ${n} appointments for Dr. ${doctor.name}`);
	for (let i = 1; i <= n; i++) {
		await generateAppointment(doctor);
	}
}

/**
 *
 * @param {import('@prisma/client').Doctor} doctor
 */
async function generateAppointment(doctor) {
	const start = faker.date.soon({ days: 1 });
	// random duration of the appointment **IN MINUTES**
	const span = faker.number.int({ min: 10, max: 60 });
	return await prisma.appointment.create({
		data: {
			patient: faker.person.fullName(),
			start,
			end: new Date(start.getTime() + span * 60000),
			doctorId: doctor.id
		}
	});
}

async function main() {
	await prisma.$connect();
	const n = faker.number.int({ min: 4, max: 9 });
	console.info(`Generating ${n} doctors`);
	for (let i = 1; i <= n; i++) {
		await generateDoctor();
	}
}

main()
	.then(() => {})
	.catch(async (e) => {
		console.error(e);
		await prisma.$disconnect();
		process.exit(1);
	});
