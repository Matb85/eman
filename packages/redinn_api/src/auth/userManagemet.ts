import User, { UserI } from "../models/user";

async function CreateUser(email: string, password: string): Promise<UserI> {
  return await User.create({ email, password })
    .then(data => {
      return data;
    })
    .catch(err => {
      throw err;
    });
}

async function GetUser(email: string): Promise<UserI> {
  return await User.findOne({ email })
    .then(data => {
      return data;
    })
    .catch(err => {
      throw err;
    });
}

export default {
  CreateUser,
  GetUser,
};
