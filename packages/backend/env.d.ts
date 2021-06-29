declare namespace NodeJS {
  interface ProcessEnv {
    DATABASE_HOST: string;
    DATABASE_PORT: string;
    DATABASE_NAME: string;
    BROWSER_BASE_URL: string;
    PORT: string;
    NODE_ENV: string;
  }
  interface Global {
    mongod: MongoMemoryServer;
  }
}
