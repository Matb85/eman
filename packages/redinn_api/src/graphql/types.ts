import { gql } from "apollo-server-express";

export default gql`
  type Query {
    books: [Book]
    user: User
  }
  type Book {
    title: String
    author: String
  }
  type User {
    id: ID!
    email: String!
    firstName: String!
    lastName: String!
    enterprises: [ID!]
  }
`;
