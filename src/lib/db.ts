import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

(async () => {
    try {
        await prisma.$connect();
    } catch (error) {
        console.error('Failed to initialize connection to database', error);
        await prisma.$disconnect().catch(() => {});
        process.exit(1);
    }
})();

export { prisma };
