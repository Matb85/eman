<template>
  <section class="apl-form background-medium">
    <h1 class="is-size-2 is-unselectable has-text-centered">Witaj z powrotem!</h1>
    <b-input type="text" placeholder="Nazwa Użytkownika" class="mt-5" rounded></b-input>
    <b-input type="password" placeholder="Hasło" class="mt-4" password-reveal rounded></b-input>
    <div class="mt-5 form-con">
      <b-button tag="nuxt-link" to="/core" class="is-primary is-size-4 py-0" rounded expanded>Zaloguj</b-button>
    </div>
    <nuxt-link to="/register" class="subtitle has-text-centered new-account is-unselectable is-clickable"
      >Załóż konto</nuxt-link
    >
  </section>
</template>

<script lang="ts">
export default {
  name: "Login",
  layout: "background",
};
</script>

<style lang="scss">
@import "@/assets/scss/partials/_loginregister.scss";
</style>
