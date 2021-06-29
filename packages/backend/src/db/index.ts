import mongoose from "mongoose";

const DB_URL =
  "mongodb://" + process.env.DATABASE_HOST + ":" + process.env.DATABASE_PORT + "/" + process.env.DATABASE_NAME;

mongoose.set("useNewUrlParser", true);
mongoose.set("useFindAndModify", false);
mongoose.set("useCreateIndex", true);
mongoose.set("useUnifiedTopology", true);
mongoose.connect(DB_URL);

const db = mongoose.connection;
db.on("error", console.error.bind(console, "connection error:"));
db.once("open", function () {
  console.log("connected to the db!");
});
