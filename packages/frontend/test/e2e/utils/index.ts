import axios from "axios";

export const axiosApi = axios.create({ baseURL: "http://localhost:3001/api", timeout: 5000 });

function randomString(length: number) {
  var result = "";
  var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  var charactersLength = characters.length;
  for (var i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

// User ralated utils
export interface UserI {
  id: string;
  email: string;
  password: string;
  firstname: string;
  lastname: string;
  enterprises: string[];
}

class User implements UserI {
  id: string = "";
  email: string;
  password: string;
  firstname: string;
  lastname: string;
  enterprises: string[] = [];
  constructor(email: string, password: string, firstname: string, lastname: string) {
    this.email = email;
    this.password = password;
    this.firstname = firstname;
    this.lastname = lastname;
  }
}

export function mockUser(): UserI {
  return new User("user@example.com" + randomString(10), "example-password" + randomString(10), "example", "user");
}

// Enterprise related utils
export interface EnterpiseI {
  id: string;
  name: string;
  logo: string;
  address: AddressI;
  enterprises: EmployeeI[];
}
interface EmployeeI {
  ref: string;
  permissions: string;
}
interface AddressI {
  country: string;
  street: string;
  zipcode: string;
  city: string;
}
class Enterpise implements EnterpiseI {
  id: string = "";
  name: string;
  logo: string = "";
  address: AddressI;
  enterprises: EmployeeI[] = [];
  constructor(name: string, address: AddressI) {
    this.name = name;
    this.address = address;
  }
}

export function mockEnterpise(): EnterpiseI {
  const address: AddressI = {
    country: "country" + randomString(10),
    street: "street" + randomString(10),
    zipcode: "zipcode" + randomString(10),
    city: "city" + randomString(10),
  };
  return new Enterpise("user@example.com" + randomString(10), address);
}
