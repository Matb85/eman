import { gql } from "apollo-server-express";

export default gql`
  type Enterprise {
    name: String!
    logo: String!
    address: Adress
    employees: [Int]
  }
  type Adress {
    country: String!
    zipcode: String!
    city: String!
    street: String!
  }

  input EnterpriseI {
    name: String!
    logo: String!
    address: AdressI
  }
  input AdressI {
    country: String!
    zipcode: String!
    city: String!
    street: String!
  }
`;
