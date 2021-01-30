<template>
  <section>
    <RoundCard class="mb-2" dim="2em" fontSize="1.4em">3</RoundCard>
    <h1 class="title is-2">Rusz podbijać rynek!</h1>
    <img class="steps" src="~/assets/imgs/rocket.svg" alt="rocket" />
    <div class="next-btn">
      <b-button expanded class="mb-6 mt-4" tag="nuxt-link" to="/core" type="is-primary">Przejdź do Redinn</b-button>
    </div>
    <div class="steps">
      <RoundCard class="mb-2" dim="2em" fontSize="1.4em">?</RoundCard>
      <h3 class="title is-3 mb-1">Masz więcej działalności?</h3>
    </div>
    <div class="next-btn">
      <b-button outlined expanded class="mt-4" tag="nuxt-link" to="/addbusiness/add" type="is-primary"
        >Dodaj je!</b-button
      >
    </div>
  </section>
</template>

<script lang="ts">
import RoundCard from "@/components/RoundCard.vue";
export default {
  name: "Addobject",
  components: {
    RoundCard,
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/scss/partials/_addbusiness.scss";
</style>
