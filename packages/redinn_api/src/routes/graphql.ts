import { Router } from "express";
const router = Router();

router.post("/", async (req, res) => {
  res.send("Hello World");
  console.log(req.body);
});

export default router;
