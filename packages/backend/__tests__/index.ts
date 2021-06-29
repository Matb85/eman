import app from "../src/app";
import supertest from "supertest";
const request = supertest(app);

let token: string;
let enterpriseID: string;
const email = "doe@gmail.com";
const password = "1qsxvfe3";

describe("auth", () => {
  it("create John Doe's account", async () => {
    const response = await request.post("/auth/register").send({
      firstName: "John",
      lastName: "Doe",
      email,
      password,
    });
    expect(response.status).toBe(201);
  });

  it("login to the John Doe's account", async () => {
    const response = await request.post("/auth/login").send({
      email,
      password,
    });

    expect(response.status).toBe(200);
    expect(response.body.message).toBe("login successful");

    token = response.body.token;
  });
});

const graphqlHeaders = (token: string) => ({ Accept: "application/json", Authorization: "bearer " + token });

describe("graphql", () => {
  it("get user data", async () => {
    const response = await request.post("/graphql").set(graphqlHeaders(token)).send({
      query: "query {user {firstName}}",
    });

    expect(response.status).toBe(200);
    expect(typeof response.body.data.user.firstName).toBe("string");
  });

  it("create an enterprise", async () => {
    const response = await request
      .post("/graphql")
      .set(graphqlHeaders(token))
      .send({
        query: `mutation add($enterprise: EnterpriseI){
          addEnterprise(data: $enterprise) { 
            name address { country zipcode} employees {ref permissions}
          }}`,
        variables: {
          enterprise: {
            name: "Redinn",
            address: { country: "Poland", zipcode: "43-542", city: "Kraków", street: "Strzelców 4a" },
          },
        },
      });
    console.log("response ", response.body);
    expect(response.status).toBe(200);
  });

  it("get enterprise", async () => {
    const response = await request.post("/graphql").set(graphqlHeaders(token)).send({
      query: "query get{enterprise(index: 0){ id name}}",
    });
    enterpriseID = response.body.data.enterprise.id;

    expect(response.status).toBe(200);
    console.log("enterprise: ", response.body);
  });

  /*
  it("update enterprise", async () => {
    const response = await request
      .post("/graphql")
      .set(graphqlHeaders(token))
      .send({
        query: `mutation update($id: ID){
          updateEnterprise(enterpriseID: $id){
            message
          }
        }`,
        variables: {
          id: enterpriseID,
        },
      });

    console.log("response:: ", response.body);
    expect(response.status).toBe(200);
  });
*/
  it("delete enterprise", async () => {
    const response = await request
      .post("/graphql")
      .set(graphqlHeaders(token))
      .send({
        query: `mutation delete($password: String, $id: ID){
          deleteEnterprise(password: $password, enterpriseID: $id){
            message
          }
        }`,
        variables: {
          password,
          id: enterpriseID,
        },
      });

    console.log("response:: ", response.body);
    expect(response.status).toBe(200);
  });
});
