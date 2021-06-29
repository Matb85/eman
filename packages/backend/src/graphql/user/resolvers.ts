import User, { UserDoc, UserI } from "@/db/user";
import { Context } from "../index";

type SafeUser = Omit<UserI, "password">;

export default {
  Query: {
    async user(_: unknown, __: unknown, context: Context): Promise<SafeUser> {
      const user = (await User.findById(context.user)) as UserDoc;
      return {
        email: user.email,
        firstName: user.firstName,
        lastName: user.lastName,
        enterprises: user.enterprises,
      };
    },
  },
};
