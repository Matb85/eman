import { comparePasswords } from "./encryption";
import auth from "./userManagemet";
import passport from "passport";
import { Strategy as LocalStrategy } from "passport-local";
import { Strategy as JwtStrategy, ExtractJwt } from "passport-jwt";
import { UserDoc } from "../models/user";

passport.use(
  "local",
  new LocalStrategy(
    { usernameField: "email", passwordField: "password", session: false },
    async (email: string, password: string, done) => {
      auth
        .GetUser(email)
        .then(async (user: UserDoc) => {
          console.log("granted user", user._id);
          if (!user) return done(null, false, { message: "Authentication failed" });

          const validation = await comparePasswords(password, user.password);
          if (validation) return done(null, user, { message: "Authentication succeeded" });

          return done(null, false, { message: "Authentication failed" });
        })
        .catch(err => {
          return done(err);
        });
    }
  )
);

passport.use(
  "jwt",
  new JwtStrategy(
    { secretOrKey: "TOP_SECRET", jwtFromRequest: ExtractJwt.fromBodyField("token") },
    async (token, done) => {
      try {
        return done(null, token.id);
      } catch (error) {
        done(error);
      }
    }
  )
);
