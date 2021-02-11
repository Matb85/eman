require("dotenv").config(); // eslint-disable-line
import express from "express";
import bodyParser from "body-parser";
import morgan from "morgan";
import passport from "passport";
import authRoutes from "./auth/routes";
import apollo from "./graphql/routes";
import "./auth/strategies";

const port = process.env.PORT || 3000;
const app = express();
app.use(morgan(":date[iso] :method :url status :status :res[content-length] bytes - :response-time ms"));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true })); // for parsing application/x-www-form-urlencoded

app.use(passport.initialize());
app.use(passport.session());

app.use("/auth", authRoutes);

app.use("/graphql", /*passport.authenticate("jwt", { session: false }),*/ apollo.getMiddleware({ path: "/" }));

app.listen(port, () => console.log("App listening on port " + port));
