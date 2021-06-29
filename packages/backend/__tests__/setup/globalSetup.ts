require("dotenv").config(); // eslint-disable-line
import { MongoMemoryServer } from "mongodb-memory-server-core";

const port = 8000;
const dbName = "test";
process.env.DATABASE_HOST = "127.0.0.1";
process.env.DATABASE_PORT = "" + port;
process.env.DATABASE_NAME = dbName;
process.env.BROWSER_BASE_URL = "http://localhost:" + (port + 1);

const mongod = new MongoMemoryServer({
  instance: { port, dbName },
});

export default async function (): Promise<void> {
  await mongod.start();

  global["mongod"] = mongod;
}
