# docque
Doctor's queue app. A prototype of a web application that can be used for viewing & managing appointments of a doctor within a hospital. Made for experimental & learning purposes.

A doctor's appointments are sorted via their date and showed. The home page shows all the doctor's queues. Further each doctor has their own page that only shows their appointment (todo!).

## Stack
This is a fullstack app made using [SvelteKit](https://kit.svelte.dev). SvelteKit's [Api Routes](https://learn.svelte.dev/tutorial/get-handlers) are used for getting and putting appointments to the backend. An sqlite database is used with [prisma](https://prisma.io) for storing the data.

## Development
### Required tools
- [nodejs](https://nodejs.org) - LTS or svelte supported versions
- [pnpm](https://pnpm.io) - for dependency managment, tho yarn/npm should work too
- [git](https://git-scm.com) - for cloning the repo

Clone the repository
```bash
$ git clone https://github.com/Yakiyo/docque && cd docque
```
Before running, the database needs to be synced
```bash
$ pnpm prisma db push
```
During development, dummy data can be generated via the [seed script](./prisma/seed.js) which uses [fakerjs](https://fakerjs.dev) for generating fake data. To seed the data, use:
```bash
$ pnpm seed
```
Then just run it
```bash
$ pnpm dev # for a development build

$ pnpm build && pnpm preview # for previewing a production build
```
[Vercel](https://vercel.com) & [Netlify](https://www.netlify.com) are some sites that provide out-of-the-box support for SvelteKit and can be used for deploying it.
