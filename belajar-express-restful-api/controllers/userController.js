const db = require("../config/db");

exports.createUser = async (req, res) => {
  const { name, email, password } = req.body;
  try {
    const result = await db.query(
      "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *",
      [name, email, password]
    );
    res.status(201).json({
      status: "success",
      data: {
        user: result.rows[0],
      },
    });
  } catch (error) {
    res.status(500).json({
      status: "error",
      message: error.message,
    });
  }
};
