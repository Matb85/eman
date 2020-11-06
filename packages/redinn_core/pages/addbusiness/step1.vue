<template>
  <section class="add-business">
    <roundCard class="mb-2" dim="2em" fontSize="1.4em">1</roundCard>
    <h1 class="title is-2">Dodaj Firmę</h1>

    <h3 class="subtitle is-4 mt-5">Ogólne informacje</h3>
    <div class="input-form">
         <b-field label="Nazwa firmy" label-position="on-border">
            <b-input rounded placeholder="np. Redinn"></b-input>
        </b-field>
         <b-field label="Typ działalności" label-position="on-border">
            <b-select rounded expanded placeholder="Select a name">
                <option>Hotelarstwo</option>
                <option>Usugi</option>
                <option>Przemysł</option>
            </b-select>
        </b-field>
    </div>

    <h3 class="subtitle is-4 mt-5">Adress Firmy</h3>
    <div class="input-form">
         <b-field label="Ulica" label-position="on-border">
            <b-input rounded placeholder="np. Strzeców 4"></b-input>
        </b-field>
         <b-field label="miasto" label-position="on-border">
            <b-input rounded placeholder="Kraków"></b-input>
        </b-field>
         <b-field label="kod pocztowy" label-position="on-border">
            <b-input rounded placeholder="31-422"></b-input>
        </b-field>
    </div>
    
    <h3 class="subtitle is-4 mt-5">Logo lub zdjęcie firmy</h3>
    <b-field><b-upload drag-drop accept="image/*" required type="is-dark">
            <section class="section content has-text-centered">
                    <b-icon
                        icon="upload"
                        size="is-large">
                    </b-icon>
                    <p>Upuść lub wybierz plik</p>
            </section>
    </b-upload></b-field>

    <h3 class="title is-2 mt-6">Wybież usługi dla Firmy</h3>
    <div class="services input-form">
            <tileCard>
                <h4 class="title is-3">Redinn Media</h4>
                <img src="~/assets/imgs/socials.svg" alt="Redinn Media">
                <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
            </tileCard>
            <tileCard>
                <h4 class="title is-3">Redinn Pages</h4>
                <img src="~/assets/imgs/wall_post.svg" alt="Redinn Media">
                <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
            </tileCard>
            <tileCard>
                <h4 class="title is-3">Redinn Teams</h4>
                <img src="~/assets/imgs/work_chat.svg" alt="Redinn Media">
                <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
            </tileCard>
            <tileCard>
                <h4 class="title is-3">Redinn Book</h4>
                <img src="~/assets/imgs/travel_booking.svg" alt="Redinn Media">
                <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
            </tileCard>
        </div>

    <div class="next-btn">
      <b-button rounded expanded class="mb-6 mt-6" tag="nuxt-link" to="/addbusiness/step2" type="is-primary">Idź dalej</b-button>
    </div>
  </section>
</template>

<script lang="ts">
import roundCard from "@/components/roundCard.vue"
import tileCard from "@/components/tileCard.vue"

export default {
    name: "addbusiness",
    layout: "backgroundFullW",
    components: {
      roundCard,
      tileCard
    }
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_addbusiness.scss";
</style>