import { Middleware } from "@nuxt/types";

const enterprises: Middleware = ({ $auth, redirect, route }) => {
  console.log($auth.user);
  if (!$auth.user) return;
  if ($auth.user.enterprises.length === 0 && !route.name?.includes("addbusiness")) redirect("/addbusiness");
};

export default enterprises;
