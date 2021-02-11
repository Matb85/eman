import User, { UserDoc, UserI } from "../models/user";

async function CreateUser(credentials: UserI): Promise<UserDoc> {
  return await User.create(credentials)
    .then(data => {
      return data;
    })
    .catch(err => {
      throw err;
    });
}

async function GetUser(email: string): Promise<UserDoc> {
  return await User.findOne({ email })
    .then((data: UserDoc) => {
      return data;
    })
    .catch((err: Error) => {
      throw err;
    });
}

export default {
  CreateUser,
  GetUser,
};
