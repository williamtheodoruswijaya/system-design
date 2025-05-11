const express = require("express");
const serverless = require("serverless-http");
const userRoutes = require("../routes/userRoutes");
const app = express();

app.use(express.json());
app.use("/users", userRoutes);

module.exports.handler = serverless(app);
