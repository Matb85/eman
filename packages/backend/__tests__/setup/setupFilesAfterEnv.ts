import mongoose from "mongoose";

afterAll(() => {
  mongoose.connection.close();
});
