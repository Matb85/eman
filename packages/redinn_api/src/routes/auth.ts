import { Router } from "express";
import passport from "passport";
import userManagemet from "../auth/userManagemet";
import { GeneratePasswordHash } from "../auth/encryption";
import jwt from "jsonwebtoken";
import { UserI } from "../models/user";
const router = Router();

interface CredentialsI {
  email: string;
  password: string;
}

router.post("/register", async (req, res) => {
  const { password, email }: CredentialsI = req.body;
  if (!password || !email) return res.send({ error: "Provide an email and password." });
  if (await userManagemet.GetUser(email)) return res.send({ error: "this email is already taken." });

  return userManagemet
    .CreateUser(email, await GeneratePasswordHash(password))
    .then(() => {
      return res.status(201).send({ message: "An account has been created. Welcome to Redinn!" });
    })
    .catch(err => {
      throw err;
    });
});

router.post("/login", async (req, res, next) => {
  passport.authenticate("local", { session: false }, (err, user: UserI, { message } = "") => {
    if (err || !user) return res.status(400).send({ err, message });

    const token = jwt.sign({ id: user._id, email: user.email }, "TOP_SECRET", { expiresIn: "15m" });

    return res.send({ message: "login successful", token });
  })(req, res, next);
});

export default router;
