import { ApolloServer } from "apollo-server-express";
import typeDefs from "./types";
import resolvers from "./resolvers";

export default new ApolloServer({
  typeDefs,
  resolvers,
  context: context => ({ user: context.req.user }),
});
