import app from "../src/app";
import supertest from "supertest";
const request = supertest(app);

it("Testing to see if Jest works", () => {
  expect(1).toBe(1);
});

it("gets the test endpoint", async done => {
  const response = await request.post("/graphql");

  expect(response.status).toBe(200);
  expect(response.body.message).toBe("pass!");
  done();
});
