import mongoose, { Document, Schema } from "mongoose";

mongoose.set("useNewUrlParser", true);
mongoose.set("useFindAndModify", false);
mongoose.set("useCreateIndex", true);
mongoose.set("useUnifiedTopology", true);
mongoose.connect(process.env.DB || "mongodb://127.0.0.1:27017/", {});

const db = mongoose.connection;
db.on("error", console.error.bind(console, "connection error:"));
db.once("open", function() {
  console.log("connected to the db!");
});

export interface UserI {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  enterprises: number[];
}

export interface UserDoc extends UserI, Document<number> {}

const UserSchema = new Schema<UserDoc>({
  email: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  firstName: { type: String, required: true },
  lastName: { type: String, required: true },
  enterprises: { type: [Number], default: [] },
});

export default mongoose.model<UserDoc>("User", UserSchema);
