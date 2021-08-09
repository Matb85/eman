import { mockUser, axiosApi, UserI } from ".";

export async function RegisterUser(): Promise<UserI> {
  const user = mockUser();
  // register the user
  const response = await axiosApi.post("/auth/register", user).catch((err) => err.response);
  expect(response.data.message).toBe("registration successful");
  expect(response.status).toBe(200);
  return user;
}

export async function Login(user: UserI): Promise<void> {
  // login the user
  const response = await axiosApi.post("/auth/login", user).catch((err) => err.response);
  expect(response.data.message).toBe("Welcome to Redinn!");
  expect(response.status).toBe(200);
  browser.setCookies({ name: "auth._token.local", value: response.data.access_token });
}

export function Logout() {
  return browser.deleteCookies(["auth._token.local", "auth._token_expiration.local"]);
}
