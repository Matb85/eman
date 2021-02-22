import { ApolloServer, gql } from "apollo-server-express";
import { merge } from "lodash";
import UserTypeDefs from "./user/types";
import UserResolvers from "./user/resolvers";
import EnterpriseTypeDefs from "./enterprise/types";
import EnterpriseResolvers from "./enterprise/resolvers";

const Query = gql`
  type Query {
    user: User
    enterprise: Enterprise
  }
  type Mutation {
    addEnterprise(data: EnterpriseI): Enterprise
  }
`;

export default new ApolloServer({
  typeDefs: [Query, UserTypeDefs, EnterpriseTypeDefs],
  resolvers: merge({}, UserResolvers, EnterpriseResolvers),
  context: context => ({ user: context.req.user }),
});

export interface Context {
  user: string;
}
