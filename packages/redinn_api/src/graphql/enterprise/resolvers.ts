import { Context } from "../index";
import Enterprise, { EnterpriseI } from "@/db/enterprise";
import User from "@/db/user";

export default {
  Query: {
    enterprise(_: unknown, id: number): Promise<EnterpriseI> {
      return Enterprise.findById(id)
        .then((enterprise: EnterpriseI) => {
          return enterprise;
        })
        .catch((err: Error) => err);
    },
  },
  Mutation: {
    addEnterprise(_: unknown, data: EnterpriseI, context: Context): Promise<EnterpriseI> {
      return Enterprise.create({
        name: data.name,
        adress: data.address,
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
