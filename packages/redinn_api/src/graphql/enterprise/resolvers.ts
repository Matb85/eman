import { Context } from "../index";
import Enterprise, { EnterpriseI } from "@/db/enterprise";
import User, { UserDoc } from "@/db/user";

export default {
  Query: {
    // enterprise(id: $id) {...}
    enterprise(_: unknown, { id }: { id: number }, context: Context): Promise<EnterpriseI> {
      return User.findById(context.user)
        .populate("enterprises")
        .then((user: UserDoc) => user.enterprises[id]);
    },
  },
  Mutation: {
    // addEnterprise(data: $enterprise) { ... }
    addEnterprise(_: unknown, { data }: { data: EnterpriseI }, context: Context): Promise<EnterpriseI> {
      return Enterprise.create({
        name: data.name,
        address: data.address,
        employees: [{ ref: context.user, role: "admin" }],
      })
        .then(enterprise => {
          User.findByIdAndUpdate(context.user, { $addToSet: { enterprises: enterprise._id } }).catch(
            (err: Error) => err
          );
          console.log("data:", enterprise);
          return enterprise;
        })
        .catch(err => err);
    },
  },
};
