<template>
  <section class="add-business">
    <roundCard class="mb-2" dim="2em" fontSize="1.4em">3</roundCard>
    <h1 class="title is-2">Rusz podbijać rynek!</h1>
        <img class="steps" src="~/assets/imgs/rocket.svg" alt="rocket">
    <div class="next-btn">
      <b-button rounded expanded class="mb-6 mt-4" tag="nuxt-link" to="/core" type="is-primary">Przejdź do Redinn</b-button>
    </div>
    <div class="steps">
      <roundCard class="mb-2" dim="2em" fontSize="1.4em">?</roundCard>
      <h3 class="title is-3 mb-1">Masz więcej działalności?</h3>
    </div>
    <div class="next-btn">
      <b-button rounded outlined expanded class="mb-6 mt-4" tag="nuxt-link" to="/addbusiness/step1" type="is-primary">Dodaj je!</b-button>
    </div>
  </section>
</template>

<script lang="ts">
import roundCard from "@/components/roundCard.vue"
export default {
    name: "addobject",
    layout: "backgroundFullW",
    components: {
      roundCard
    }
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_addbusiness.scss";
</style>