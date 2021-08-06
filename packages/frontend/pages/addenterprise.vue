<template>
  <form class="add-enterprise" @submit.prevent="addEnterprise">
    <h1 class="title is-2">Dodaj działalność</h1>
    <h3 class="is-size-4">Ogólne informacje</h3>

    <b-field label="Nazwa firmy" label-position="on-border" class="my-4">
      <b-input v-model="enterpriseData.name" required rounded placeholder="np. Redinn" />
    </b-field>

    <h3 class="is-size-4">Adress</h3>

    <div class="input-form">
      <b-field v-for="(f, nm) in addressFields" :key="f.label" :label="f.label" label-position="on-border">
        <b-input v-model="enterpriseData.address[nm]" required rounded :placeholder="f.placeholder" />
      </b-field>
    </div>

    <h3 class="is-size-4">Logo lub zdjęcie</h3>

    <b-upload v-model="logo" drag-drop accept="image/*" required type="is-dark" class="my-4" expanded>
      <section class="py-4 content has-text-centered">
        <b-icon icon="upload" size="is-large" />
        <p>Upuść lub wybierz plik</p>
      </section>
    </b-upload>

    <div class="submit-con">
      <b-button tag="input" native-type="submit" value="Dodaj" type="is-primary" expanded />
    </div>
    <nuxt-link to="/core" class="new-account">Anuluj</nuxt-link>
  </form>
</template>

<script lang="ts">
import { Vue, Component } from "nuxt-property-decorator";

class AddressInput {
  label: string;
  placeholder: string;
  constructor(label: string, placeholder: string) {
    this.label = label;
    this.placeholder = placeholder;
  }
}
@Component({ layout: "background" })
export default class AddEnterprise extends Vue {
  logo: File | null = null;
  enterpriseData = {
    name: "",
    address: { country: "", zipcode: "", city: "", street: "" },
  };

  addressFields = {
    country: new AddressInput("Państwo", "np. Polska"),
    zipcode: new AddressInput("kod pocztowy", "np. 43-234"),
    city: new AddressInput("Miasto", "np. Kraków"),
    street: new AddressInput("ulica", "np. Strzelców 4"),
  };

  async addEnterprise() {
    try {
      // add enterprise info
      const response = await this.$axios
        .$post("/secure/graphql", {
          query: `mutation add($enterprise: EnterpriseI!){
          addEnterprise(data: $enterprise) { name }
          }`,
          variables: { enterprise: this.enterpriseData },
        })
        .catch((err) => err.response);
      // update the user
      await this.$auth.fetchUser();
      // upload the logo
      await this.uploadLogo();
      // if OK inform the user & redirect to the home page
      this.$buefy.snackbar.open({
        message: "Dodano działalność o nazwie " + response.data.addEnterprise.name,
        type: "is-success",
        queue: false,
        position: "is-bottom-right",
      });
      this.$router.push({
        path: "/0/media",
      });
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
  async uploadLogo() {
    // get the upload token
    const response = await this.$axios
      .$post("/secure/graphql", {
        query: `mutation logo($filename: String!){
              enterpriseLogo(index: ${this.$auth.user.enterprises.length - 1}, filename: $filename) { uploadtoken }
            }`,
        variables: {
          filename: this.logo!.name,
        },
      })
      .catch((err) => err);
    console.log(response.data);
    console.log(this.logo);
    // post the logo file
    const formData = new FormData();
    formData.append("file", this.logo!);
    await this.$axios.$post("/secure/assets", formData, {
      headers: { "Content-Type": "multipart/form-data", Upload: response.data.enterpriseLogo.uploadtoken },
    });
  }
}
</script>

<style lang="scss">
.add-enterprise {
  text-align: center;
  padding: 0 0 2rem 0;
  .input-form {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-auto-rows: auto;
    grid-gap: 0.5rem;
    margin: 1rem 0;
  }
}
@media (max-width: 550px) {
  .add-enterprise {
    padding: 0 0.75rem 2rem 0.75rem;
    max-width: 300px !important;
  }
  .add-enterprise .input-form {
    grid-template-columns: repeat(1, 1fr);
    min-width: 250px;
  }
}
</style>
