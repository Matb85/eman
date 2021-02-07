import express from "express";
import bodyParser from "body-parser";
import passport from "passport";
import authRoutes from "./routes/auth";
import graphql from "./routes/graphql";
import "./auth/passport";

const port = process.env.PORT || 3000;
const app = express();
app.use(bodyParser.json());

app.use(passport.initialize());
app.use(passport.session());

app.use("/auth", authRoutes);
app.use("/graphql", graphql);

app.listen(port, () => console.log("App listening on port " + port));
