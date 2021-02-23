import { gql } from "apollo-server-express";

export default gql`
  type Enterprise {
    name: String!
    logo: String!
    address: Address!
    employees: [Employee!]
  }
  type Address {
    country: String!
    zipcode: String!
    city: String!
    street: String!
  }
  type Employee {
    ref: String
    role: String!
  }

  input EnterpriseI {
    name: String!
    address: AddressI!
  }
  input AddressI {
    country: String!
    zipcode: String!
    city: String!
    street: String!
  }
`;
