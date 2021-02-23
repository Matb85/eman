import { Context } from "../index";
import Enterprise, { EnterpriseI } from "@/db/enterprise";
import User from "@/db/user";

export default {
  Query: {
    // enterprise(id: $id) {...}
    enterprise(_: unknown, id: number): Promise<EnterpriseI> {
      return Enterprise.findById(id)
        .then((enterprise: EnterpriseI) => {
          return enterprise;
        })
        .catch((err: Error) => err);
    },
  },
  Mutation: {
    // addEnterprise(data: $enterprise) { ... }
    addEnterprise(_: unknown, { data }: { data: EnterpriseI }, context: Context): Promise<EnterpriseI> {
      console.log(data);
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
