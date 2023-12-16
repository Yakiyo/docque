import { fastify } from "fastify";

const app = fastify({
  logger: true,
  clientErrorHandler: (error) => console.error(error),
});

app.listen({ port: Number(process.env.PORT ?? 3000) }, (err, port) => {
  if (err) {
    console.error("Unexpected error when starting server", err);
    process.exit(1);
  }
  console.info(`Listening to ${port}`);
});
