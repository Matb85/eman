import { comparePasswords } from "./encryption";
import User, { UserDoc } from "@/db/user";
import passport from "passport";
import { Strategy as LocalStrategy } from "passport-local";
import { Strategy as JwtStrategy, ExtractJwt } from "passport-jwt";

passport.use(
  "local",
  new LocalStrategy(
    { usernameField: "email", passwordField: "password", session: false },
    async (email: string, password: string, done) => {
      User.findOne({ email })
        .then(async (user: UserDoc) => {
          if (!user) return done(null, false, { message: "Authentication failed" });

          const validation = await comparePasswords(password, user.password);
          if (validation) return done(null, user, { message: "Authentication succeeded" });

          return done(null, false, { message: "Authentication failed" });
        })
        .catch((err: Error) => {
          return done(err);
        });
    }
  )
);

passport.use(
  "jwt",
  new JwtStrategy(
    { secretOrKey: "TOP_SECRET", jwtFromRequest: ExtractJwt.fromAuthHeaderWithScheme("bearer") },
    async (token, done) => {
      try {
        return done(null, token.id);
      } catch (error) {
        done(error);
      }
    }
  )
);
