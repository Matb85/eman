import User, { UserI } from "@/db/user";
import { Context } from "../index";

type SafeUser = Omit<UserI, "password">;

export default {
  Query: {
    user(_: unknown, __: unknown, context: Context): Promise<SafeUser> {
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
