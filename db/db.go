package db

// shared `Client` instance to use
var Client *PrismaClient

func Init() error {
	Client = NewClient()
	return Client.Prisma.Connect()
}

func Close() error {
	return Client.Prisma.Disconnect()
}
