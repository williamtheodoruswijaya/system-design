const express = require("express");
const bodyParser = require("body-parser");
const userRoutes = require("./routes/userRoutes");
require("dotenv").config();

const app = express();
app.use(bodyParser.json());

app.use("/users", userRoutes);

// Optional: central error handler
// app.use(require('./middlewares/errorHandler'));

module.exports = app;
