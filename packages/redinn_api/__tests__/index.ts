import app from "../src/app";
import supertest from "supertest";
const request = supertest(app);

let token: string;
const email = "doe@gmail.com";
const password = "1qsxvfe3";
let enterpriseID: string;

describe("auth", () => {
  it("login to the John Doe's account", async done => {
    const response = await request.post("/auth/login").send({
      email,
      password,
    });

    expect(response.status).toBe(200);
    expect(response.body.message).toBe("login successful");

    token = response.body.token;
    done();
  });
});

const graphqlHeaders = (token: string) => ({ Accept: "application/json", Authorization: "bearer " + token });

describe("graphql", () => {
  it("get user data", async done => {
    const response = await request
      .post("/graphql")
      .set(graphqlHeaders(token))
      .send({
        query: "query {user {firstName}}",
      });

    expect(response.status).toBe(200);
    expect(typeof response.body.data.user.firstName).toBe("string");
    done();
  });

  it("create an enterprise", async done => {
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

    done();
  });

  it("get enterprise", async done => {
    const response = await request
      .post("/graphql")
      .set(graphqlHeaders(token))
      .send({
        query: "query get{enterprise(index: 0){ id name}}",
      });
    enterpriseID = response.body.data.enterprise.id;
    console.log(enterpriseID);
    expect(response.status).toBe(200);
    console.log("enterprise: ", response.body);
    done();
  });

  it("delete enterprise", async done => {
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
          password: password,
          id: enterpriseID,
        },
      });

    console.log("response:: ", response.body);
    expect(response.status).toBe(200);
    done();
  });
});
