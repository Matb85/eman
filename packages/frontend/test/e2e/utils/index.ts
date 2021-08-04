import axios from "axios";

function randomString(length: number) {
  var result = "";
  var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  var charactersLength = characters.length;
  for (var i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

interface UserI {
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

export const axiosApi = axios.create({ baseURL: "http://localhost:3001", timeout: 5000 });
