<template>
  <main class="apl-form">
    <h1 class="is-size-2 is-unselectable has-text-centered">Dobrze Cię widzieć!</h1>
    <form class="mt-5" @submit.prevent="login">
      <b-input data-testid="login-email" v-model="credentials.email" type="text" placeholder="Email" rounded></b-input>
      <b-input
        data-testid="login-password"
        v-model="credentials.password"
        type="password"
        placeholder="Hasło"
        class="mt-3"
        rounded
        password-reveal
      ></b-input>
      <div class="mt-3 submit-con">
        <b-button data-testid="login-submit" type="is-primary" native-type="submit" class="is-size-4 py-0" expanded
          >Zaloguj</b-button
        >
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

  async login(): Promise<void> {
    try {
      await this.$auth.loginWith("local", { data: this.credentials });
      const user = await this.$axios.$post("/api/graphql", {
        query: "query { user {email, firstName, lastName, enterprises }}",
      });
      this.$auth.setUser(user.data.user);
    } catch (err) {
      this.$buefy.snackbar.open({
        message: "Coś poszło nie tak... Spróbuj ponownie później",
        type: "is-danger",
        queue: false,
        position: "is-bottom-right",
      });
      console.log(err);
    }
  }
}
</script>
