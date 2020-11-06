<template>
  <section class="add-business">
    <roundCard class="mb-2" dim="2em" fontSize="1.4em">2</roundCard>
    <h1 class="title is-2">Skonfiguruj usługi firmy</h1>
    <h3 class="subtitle is-4 mt-5 mb-1">Redinn Media</h3>
    <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">Zaloguj się do google calendar i wybranych serwisów</p>
    <div class="flex-form columns is-centered mb-0">
        <colorfulbtn class="column" icon="gcalendar" addedclass="backlit">Google Business</colorfulbtn>
    </div>
    <div class="flex-form columns is-multiline is-centered">
        <colorfulbtn class="column" icon="gbusiness" addedclass="backlit">Google Business</colorfulbtn>
        <colorfulbtn class="column" icon="facebook">Facebook</colorfulbtn>
        <colorfulbtn class="column" icon="instagram" addedclass="whiteicon instagram">Instagram</colorfulbtn>
        <colorfulbtn class="column" icon="pinterest">Pinterest</colorfulbtn>
        <colorfulbtn class="column" icon="twitter" addedclass="whiteicon">Twitter</colorfulbtn>
    </div>

    <h3 class="subtitle is-4 mt-5">Redinn Pages</h3>
    <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">Kliknij odpowiednią opcję</p>

    <div class="services input-form">
        <tileCard>
            <h4 class="title is-3">Przenieś stronę do Redinn</h4>
            <img src="~/assets/imgs/move_page.svg" alt="move your page to Redinn">
            <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
        </tileCard>
        <tileCard>
            <h4 class="title is-3">Utwórz własną stronę Redinn</h4>
            <img src="~/assets/imgs/create_page.svg" alt="crete your own Redinn page">
            <p class="is-size-6 is-family-secondary has-text-weight-bold">Zarządzaj social mediami, publikuj na wielu playformach jednocześnie, używaj botów do like’owania, oraz wiele więcej!</p>
        </tileCard>
    </div>

    <div class="next-btn">
      <b-button rounded expanded class="mb-6 mt-6" tag="nuxt-link" to="/addbusiness/step3" type="is-primary">Idź dalej</b-button>
    </div>
  </section>
</template>

<script lang="ts">
import roundCard from "@/components/roundCard.vue"
import tileCard from "@/components/tileCard.vue"
import colorfulbtn from "@/components/btnCornerIcon.vue"
export default {
    name: "addobject",
    layout: "backgroundFullW",
    components: {
      roundCard,
      tileCard,
      colorfulbtn
    }
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_addbusiness.scss";
.flex-form {
    width: calc(50% + 2em);
    min-width: 600px;
    height: auto;
    margin: 0 auto;
    >*{
        flex: 0 0 calc(33% - 1em);
        margin: 0.5em;
    }
}
@media (max-width: $tablet){
    .flex-form {min-width: 200px;}
}
</style>