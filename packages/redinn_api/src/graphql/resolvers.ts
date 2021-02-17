import User, { UserI } from "../db/user";
import { Error } from "mongoose";
type SafeUser = Omit<UserI, "password" | "id">;

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
    user(_: unknown, __: unknown, context: { user: string }): Promise<SafeUser> {
      return User.findById(context.user)
        .then(
          (user: UserI): SafeUser => {
            return {
              email: user.email,
              firstName: user.firstName,
              lastName: user.lastName,
              enterprises: user.enterprises,
            };
          }
        )
        .catch((err: Error) => err);
    },
  },
};
