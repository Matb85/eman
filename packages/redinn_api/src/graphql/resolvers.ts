import User, { UserI } from "../models/user";
// import { GetUser } from "../auth/userManagement";
import { Error } from "mongoose";

const books = [
  {
    title: "Harry Potter and the Sorcerer's stone",
    author: "J.K. Rowling",
  },
  {
    title: "Jurassic Park",
    author: "Michael Crichton",
  },
];

export default {
  Query: {
    books: (): typeof books => books,
    user(_: string, args: { id: string }): Promise<Omit<UserI, "password" | "id">> {
      return User.findById(args.id)
        .then((user: UserI) => {
          console.log(user, args);
          return {
            email: user.email,
            firstName: user.firstName,
            lastName: user.lastName,
            enterprises: user.enterprises,
          };
        })
        .catch((err: Error) => err);
    },
  },
};
