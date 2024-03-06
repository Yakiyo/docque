import { pb } from "./pocketbase";

export async function modifyApp(
  app: Record<string, string>
): Promise<Record<string, string>> {
  app.docId = app.doctor;
  app.doctor = await pb
    .collection("doctors")
    .getOne(app.doctor)
    .then((v) => v.name);
    
  const date = new Date(Date.parse(app.start));
  const datestr = `${date.getHours()}:${date.getMinutes()} - ${date.getMonth()}/${date.getDate()}/${date.getFullYear()}`;
  app.start = datestr;
  let duration = 30;
  if (app.end) {
    // @ts-ignore
    const diff = Math.abs(date - new Date(Date.parse(app.end)));
    duration = diff / 60_000;
  }
  app.duration = `${duration} minutes`;
  return app;
}
