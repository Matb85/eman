<template>
  <section id="btm" class="apl-form">
    <h1 class="is-size-2 is-unselectable has-text-centered">Witaj na pokładzie!</h1>
    <b-input type="text" placeholder="Imę i nazwisko" class="mt-5" rounded></b-input>
    <b-input type="text" placeholder="email" class="mt-4" rounded></b-input>
    <b-input type="password" placeholder="Hasło" class="mt-4" rounded></b-input>
    <b-input type="password" placeholder="Powtórz hasło" class="mt-4" rounded></b-input>
    <div class="mt-5 form-con">
      <b-button tag="nuxt-link" to="/core" class="is-primary is-size-4 py-0" rounded expanded>Zarejestruj się</b-button>
    </div>
    <nuxt-link to="/login" class="new-account subtitle is-size-6 has-text-centered is-unselectable is-clickable"
      >Masz konto? Zaloguj się</nuxt-link
    >
  </section>
</template>

<script lang="ts">
export default {
  name: "Register",
  layout: "background",
};
</script>
<style scoped lang="scss">
@import "@/assets/scss/partials/_loginregister.scss";
#btm {
  top: initial;
  bottom: 1em;
}
@media (max-width: $mobileWBP) {
  #btm {
    bottom: 8%;
  }
}
</style>
