import { Router } from "express";
import passport from "passport";
import User from "@/db/user";
import { GeneratePasswordHash } from "./encryption";
import jwt from "jsonwebtoken";
import { UserDoc, UserI } from "../db/user";
const router = Router();

router.post("/register", async (req, res) => {
  const cred: UserI = req.body; // user credentials
  if (!cred.password || !cred.email || !cred.firstName || !cred.lastName)
    return res.send({ error: "Provide all necessary credentials" });
  if (await User.findOne({ email: cred.email })) return res.send({ error: "this email is already taken." });

  cred.password = await GeneratePasswordHash(cred.password);
  cred.enterprises = [];

  return User.create(cred)
    .then(() => {
      return res.status(201).send({ message: "An account has been created. Welcome to Redinn!" });
    })
    .catch(err => {
      throw err;
    });
});

router.post("/login", async (req, res, next) => {
  passport.authenticate("local", { session: false }, (err, user: UserDoc, { message } = "") => {
    if (err || !user) return res.status(400).send({ err, message });

    const token = jwt.sign({ id: user.id }, "TOP_SECRET", { expiresIn: "30m" });

    return res.send({ message: "login successful", token });
  })(req, res, next);
});

export default router;
