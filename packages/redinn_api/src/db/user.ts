import { model, Document, Schema } from "mongoose";

export interface UserI {
  id: string;
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  enterprises: number[];
}

export interface UserDoc extends UserI, Document<string> {
  id: string;
}

const UserSchema = new Schema<UserDoc>({
  email: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  firstName: { type: String, required: true },
  lastName: { type: String, required: true },
  enterprises: { type: [Number], default: [] },
});

export default model<UserDoc>("User", UserSchema);
