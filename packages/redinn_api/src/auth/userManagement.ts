import User, { UserDoc, UserI } from "../models/user";

export async function CreateUser(credentials: UserI): Promise<UserDoc> {
  return await User.create(credentials)
    .then(data => {
      return data;
    })
    .catch(err => {
      throw err;
    });
}

type UserProp<P, T> = Record<Extract<keyof UserI, P>, T>;

export async function GetUser(data: UserProp<"id", string> | UserProp<"email", string>): Promise<UserDoc> {
  return await User.findOne(data)
    .then((data: UserDoc) => {
      return data;
    })
    .catch((err: Error) => {
      throw err;
    });
}
