import { Context } from "@nuxt/types";
import type { Auth, RecursivePartial, ModuleOptions } from "@nuxtjs/auth-next/dist/index";
import { EnterpriseI, UserI } from "@/helpers/authTypes";

export default async function ({ $auth, $axios, route }: Context) {
  if (!$auth.loggedIn || isNaN(parseInt(route.params.enterprise))) return;
  const response = await $axios.$post("/api/graphql", {
    query: `query get{enterprise(index: ${route.params.enterprise}){ name address {country zipcode city street }}}`,
  });
  $auth.enterprise = response.data.enterprise;
}

// types for nuxt-auth

interface ExtendedAuth extends Omit<Auth, "user"> {
  enterprise: EnterpriseI;
  user: UserI;
}

declare module "vue/types/vue" {
  interface Vue {
    $auth: ExtendedAuth;
  }
}
declare module "@nuxt/types" {
  interface Context {
    $auth: ExtendedAuth;
  }
  interface NuxtAppOptions {
    $auth: ExtendedAuth;
  }

  interface Configuration {
    auth?: RecursivePartial<ModuleOptions>;
  }
}

declare module "vue/types/options" {
  // eslint-disable-next-line no-unused-vars,@typescript-eslint/no-unused-vars
  interface ComponentOptions<V> {
    auth?: true | false | "guest";
  }
}

declare module "vuex/types/index" {
  // eslint-disable-next-line no-unused-vars,@typescript-eslint/no-unused-vars
  interface Store<S> {
    $auth: Auth;
  }
}
