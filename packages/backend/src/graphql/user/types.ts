import { gql } from "apollo-server-express";

export default gql`
  type User {
    id: ID!
    email: String!
    firstName: String!
    lastName: String!
    enterprises: [ID!]
  }
`;
