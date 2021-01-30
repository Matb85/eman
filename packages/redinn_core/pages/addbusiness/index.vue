<template>
  <section>
    <h1 class="title is-2">Witaj w Redinn, Maria!</h1>
    <h2 class="title is-2">Oto plan działania:</h2>
    <div class="steps">
      <RoundCard class="mb-2" dim="2em" fontSize="1.4em">1</RoundCard>
      <h3 class="title is-3 mb-1">Dodaj Firmę</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">
        Dodaj swoją działalność ich więcej oraz zaznacz usługi dla działalności. Jeśli masz więcej firm będziesz je mógł
        dodawać po kolei.
      </p>

      <RoundCard class="mb-2" dim="2em" fontSize="1.4em">2</RoundCard>
      <h3 class="title is-3 mb-1">Skonfiguruj usługi</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">
        Połącz usługi z zewnętrznymi kontami i dostosuj je do własnych potrzeb.
      </p>

      <RoundCard class="mb-2" dim="2em" fontSize="1.4em">3</RoundCard>
      <h3 class="title is-3 mb-1">Rusz podbijać rynek!</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">
        Tak to już! Zacznij odkrywać opcje i potencjał naszego ekosystemu. W razie poroblemów polecamy sprawdzić nasze
        samouczki lub centrum pomocy.
      </p>
    </div>
    <div class="next-btn">
      <b-button expanded class="mt-4" tag="nuxt-link" to="/addbusiness/add" type="is-primary">Zaczynajmy!</b-button>
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
.add-business .steps {
  width: 40%;
  min-width: 300px;
  margin: 0 auto;
}
@media (max-width: 310px) {
  .add-business .steps {
    min-width: 250px;
  }
}
</style>
