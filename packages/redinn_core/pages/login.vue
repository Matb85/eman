<template>
  <main class="apl-form">
    <h1 class="is-size-2 is-unselectable has-text-centered">Dobrze Cię widzieć!</h1>
    <form class="mt-5" @submit.prevent="login">
      <b-input v-model="credentials.email" type="text" placeholder="Email" rounded></b-input>
      <b-input
        v-model="credentials.password"
        type="password"
        placeholder="Hasło"
        class="mt-3"
        rounded
        password-reveal
      ></b-input>
      <div class="mt-3 submit-con">
        <b-button tag="input" type="is-primary" native-type="submit" value="Zaloguj" class="is-size-4 py-0" expanded />
      </div>
    </form>
    <nuxt-link to="/register" class="new-account">Załóż konto</nuxt-link>
  </main>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator";

@Component({ layout: "background" })
export default class Login extends Vue {
  credentials = {
    email: "",
    password: "",
  };

  login(): void {
    this.$auth
      .loginWith("local", { data: this.credentials })
      .then(async () => {
        const user = await this.$axios.$post("/graphql", {
          query: "query { user {email, firstName, lastName, enterprises }}",
        });
        this.$auth.setUser(user.data.user);
      })
      .catch((err) => {
        console.log(err);
      });
  }
}
</script>
