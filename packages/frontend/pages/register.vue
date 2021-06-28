<template>
  <main class="apl-form">
    <h1 class="is-size-2 has-text-centered">Witaj na pokładzie!</h1>
    <form class="mt-5" @submit.prevent="register">
      <b-field
        v-for="(value, name) in feedback"
        :key="name"
        :type="feedback[name].message.length > 0 ? 'is-danger' : ''"
        :message="feedback[name].message"
      >
        <b-input
          v-model="credentials[name]"
          :type="feedback[name].type"
          :placeholder="feedback[name].placeholder"
          rounded
          @input="feedback[name].validate(credentials[name])"
        />
      </b-field>

      <div class="mt-3 submit-con">
        <b-button
          tag="input"
          native-type="submit"
          value="Utwórz konto"
          class="is-size-4 py-0"
          type="is-primary"
          expanded
        />
      </div>
    </form>
    <nuxt-link to="/login" class="new-account">Masz konto? Zaloguj się</nuxt-link>
  </main>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator";
interface RegisterCred {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

class InputValidation {
  type: string;
  placeholder: string;
  validate: (str: string) => void;
  message = "";
  constructor(type: string, placeholder: string, validate: (str: string) => void) {
    this.type = type;
    this.placeholder = placeholder;
    this.validate = validate;
  }
}

@Component({ layout: "background", auth: "guest" })
export default class Register extends Vue {
  credentials: RegisterCred = {
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  };
  feedback = {
    firstName: new InputValidation("text", "Imię", function (this: InputValidation) {}),
    lastName: new InputValidation("text", "Nazwisko", function (this: InputValidation) {}),
    email: new InputValidation("text", "Email", function (this: InputValidation, str: string) {
      if (/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/.test(str)) {
        this.message = "";
        return;
      }
      this.message = "wpisany adres jest niepoprawny";
    }),
    password: new InputValidation("password", "Hasło", function (this: InputValidation, str: string) {
      if (str.length >= 8) {
        this.message = "";
        return;
      }
      this.message = "hasło musi mieć co najmniej 8 znaków";
    }),
  };

  async register(): Promise<void> {
    try {
      const result = await this.$axios.$post("/auth/register", this.credentials);
      console.log(result);
      if (result.error) return;
      this.$auth.loginWith("local", {
        data: this.credentials,
      });
    } catch (err) {
      console.log(err);
    }
  }
}
</script>
