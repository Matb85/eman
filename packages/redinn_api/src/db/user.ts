import { model, Document, Schema } from "mongoose";

export interface UserI {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  enterprises: number[];
}

export type UserDoc = UserI & Document<string>;

const UserSchema = new Schema<UserDoc>({
  email: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  firstName: { type: String, required: true },
  lastName: { type: String, required: true },
  enterprises: { type: [Schema.Types.ObjectId], ref: "Enterprise", default: [] },
});

export default model<UserDoc>("User", UserSchema);
