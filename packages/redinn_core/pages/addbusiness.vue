<template>
  <section class="add-business background-wide">
    <h1 class="title is-2">Dodaj działalność</h1>

    <h3 class="subtitle is-4 mt-5">Ogólne informacje</h3>
    <div class="input-form">
      <b-field label="Nazwa firmy" label-position="on-border">
        <b-input rounded placeholder="np. Redinn" />
      </b-field>
      <b-field label="Typ działalności" label-position="on-border">
        <b-select rounded expanded placeholder="np. hoterlarstwo">
          <option>Hotelarstwo</option>
          <option>Usugi</option>
          <option>Przemysł</option>
        </b-select>
      </b-field>
    </div>

    <h3 class="subtitle is-4 mt-5">Adress</h3>
    <div class="input-form">
      <b-field label="Państwo" label-position="on-border">
        <b-input rounded placeholder="np. Polska" />
      </b-field>
      <b-field label="Ulica" label-position="on-border">
        <b-input rounded placeholder="np. Strzeców 4" />
      </b-field>
      <b-field label="miasto" label-position="on-border">
        <b-input rounded placeholder="np. Kraków" />
      </b-field>
      <b-field label="kod pocztowy" label-position="on-border">
        <b-input rounded placeholder="np. 31-422" type="text" pattern="[0-9]*" />
      </b-field>
    </div>

    <h3 class="subtitle is-4 mt-5">Logo działalności</h3>

    <b-upload drag-drop accept="image/*" required type="is-dark">
      <section class="section content has-text-centered px-6">
        <b-icon icon="upload" size="is-large" />
        <p>Upuść lub wybierz plik</p>
      </section>
    </b-upload>

    <div class="next-btn">
      <b-button class="mt-6" tag="nuxt-link" to="/addbusiness/choose" type="is-primary">Dodaj</b-button>
    </div>
  </section>
</template>

<script lang="ts">
export default {
  name: "Addbusiness",
  layout: "background",
};
</script>

<style lang="scss">
@import "@/assets/scss/partials/_addbusiness.scss";
</style>
