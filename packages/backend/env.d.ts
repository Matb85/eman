declare namespace NodeJS {
  export interface ProcessEnv {
    PORT: string;
    DB: string;
    NODE_ENV: string;
  }
}
