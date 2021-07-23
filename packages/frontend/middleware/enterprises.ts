import { Middleware } from "@nuxt/types";

const enterprises: Middleware = ({ $auth, redirect, route }) => {
  if (!$auth.user) return;
  if ($auth.user.enterprises.length === 0 && !route.name?.includes("addenterprise")) redirect("/addenterprise");
};

export default enterprises;
