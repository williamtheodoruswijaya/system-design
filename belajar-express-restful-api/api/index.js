const express = require("express");
const serverless = require("serverless-http");
const userRoutes = require("../routes/userRoutes");

const app = express();
app.use(express.json());

// Debug log
app.use((req, res, next) => {
  console.log(`${req.method} ${req.path}`);
  next();
});

app.use("/users", userRoutes);

// Catch unhandled errors
app.use((err, req, res, next) => {
  console.error("Unhandled error:", err);
  res.status(500).json({ status: "error", message: "Internal Server Error" });
});

module.exports.handler = serverless(app);
