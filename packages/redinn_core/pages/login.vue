<template>
  <section class="apl-form background-medium">
    <h1 class="is-size-2 is-unselectable has-text-centered">Witaj z powrotem!</h1>
    <form class="mt-5" @submit.prevent="login">
      <b-input v-model="credentials.email" type="text" placeholder="Email" rounded></b-input>
      <b-input
        v-model="credentials.password"
        type="password"
        placeholder="Hasło"
        class="mt-4"
        rounded
        password-reveal
      ></b-input>
      <div class="mt-5 form-con">
        <b-button tag="input" native-type="submit" value="Zaloguj" class="is-primary is-size-4 py-0" expanded />
      </div>
    </form>
    <nuxt-link to="/register" class="subtitle has-text-centered new-account is-unselectable is-clickable"
      >Załóż konto</nuxt-link
    >
  </section>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator";

interface CredentialsI {
  email: string;
  password: string;
}

@Component({ layout: "background" })
export default class Login extends Vue {
  credentials: CredentialsI = {
    email: "",
    password: "",
  };

  async login(): Promise<void> {
    try {
      const response = await this.$auth.loginWith("local", { data: this.credentials });
      console.log(response);
    } catch (err) {
      console.log(err);
    }
  }
}
</script>

<style lang="scss">
@import "@/assets/scss/partials/_loginregister.scss";
</style>
