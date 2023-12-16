import express from 'express';

const app = express();
const port = Number(process.env.PORT ?? '3000');

app.get('/', (_req, res) => {
	if (process.env.NODE_ENV !== 'production') {
		return console.log('Hello World');
	}
	res.redirect('https://github.com/Yakiyo/docque');
});

app.listen(port, () => console.info(`Listening to port ${port}`));
