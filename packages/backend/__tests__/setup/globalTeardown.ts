import { MongoMemoryServer } from "mongodb-memory-server-core";

export default async function (): Promise<void> {
  const mongod: MongoMemoryServer = global["mongod"];
  await mongod.stop();
}
