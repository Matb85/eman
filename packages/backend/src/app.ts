import "./db";
import express from "express";
import bodyParser from "body-parser";
import morgan from "morgan";
import passport from "passport";
import authRoutes from "./auth";
import apollo from "./graphql";
import "./auth/strategies";

const app = express();

app.use(morgan(":date[iso] :method :url status :status :res[content-length] bytes - :response-time ms"));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true })); // for parsing application/x-www-form-urlencoded

app.use(passport.initialize());

app.use("/auth", authRoutes);

app.post("/graphql", passport.authenticate("jwt", { session: false }), apollo.getMiddleware({ path: "/" }));

export default app;
