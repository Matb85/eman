import { Context } from "../index";
import Enterprise, { EnterpriseI, EnterpriseDoc, Employee } from "@/db/enterprise";
import User, { UserDoc } from "@/db/user";
import { comparePasswords } from "@/auth/encryption";
import { UserInputError, ApolloError } from "apollo-server-express";

interface Message {
  message: string;
}
export default {
  Query: {
    // enterprise(index: Int) {...}
    async enterprise(_: unknown, { index }: { index: number }, context: Context): Promise<EnterpriseI> {
      const user = (await User.findById(context.user).populate("enterprises")) as UserDoc;
      return user.enterprises[index] as unknown as EnterpriseI;
    },
  },
  Mutation: {
    // addEnterprise(data: $enterprise) { ... }
    addEnterprise(_: unknown, { data }: { data: EnterpriseI }, context: Context): Promise<EnterpriseI> {
      return Enterprise.create({
        name: data.name,
        address: data.address,
        employees: [{ ref: context.user, permissions: "111111" }],
      })
        .then(enterprise => {
          User.findByIdAndUpdate(context.user, { $addToSet: { enterprises: enterprise._id } }).catch(
            (err: Error) => err
          );
          return enterprise;
        })
        .catch(err => err);
    },

    async deleteEnterprise(_: unknown, args: Record<string, string>, context: Context): Promise<Message> {
      const enterprise = (await Enterprise.findById(args.enterpriseID)) as EnterpriseDoc;
      if (!enterprise) throw new UserInputError("incorrect enterprise ID");
      console.log("employees: " + enterprise.employees[0]);

      const requesterEntry = enterprise.employees.find(u => u.ref == context.user) as Employee;
      console.log("requesterEntry: " + requesterEntry);
      const requester = (await User.findById(requesterEntry.ref)) as UserDoc;

      if (!(await comparePasswords(args.password, requester.password))) throw new UserInputError("incorrect password");
      if (!parseInt(requesterEntry.permissions[0]))
        throw new ApolloError("you don't have permissions to delete this enterprise");

      await User.updateMany(
        { _id: { $in: enterprise.employees.map(e => e.ref) } },
        { $pull: { enterprises: enterprise._id } }
      ).catch((err: Error) => err);

      await enterprise.remove().catch(err => err);

      return { message: "enterprise deleted" };
    },
  },
};
