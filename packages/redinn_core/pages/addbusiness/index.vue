<template>
  <section class="add-business background-wide">
    <h1 class="title is-2">Witaj w Redinn, Maria!</h1>
    <h2 class="title is-2">Oto plan działania:</h2>
    <div class="steps">
      <roundCard class="mb-2" dim="2em" fontSize="1.4em">1</roundCard>
      <h3 class="title is-3 mb-1">Dodaj Firmę</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">Dodaj swoją działalność ich więcej oraz  zaznacz usługi dla działalności. Jeśli masz więcej firm będziesz je mógł dodawać po kolei.</p>
      
      <roundCard class="mb-2" dim="2em" fontSize="1.4em">2</roundCard>
      <h3 class="title is-3 mb-1">Skonfiguruj usługi</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">Połącz usługi z zewnętrznymi kontami i dostosuj je do własnych potrzeb.</p>
      
      <roundCard class="mb-2" dim="2em" fontSize="1.4em">3</roundCard>
      <h3 class="title is-3 mb-1">Rusz podbijać rynek!</h3>
      <p class="is-size-6 is-family-secondary has-text-weight-bold mb-4">Tak to już!  Zacznij odkrywać opcje i potencjał naszego ekosystemu. W razie poroblemów polecamy sprawdzić nasze samouczki lub centrum pomocy.</p>
    </div>
    <div class="next-btn">
      <b-button rounded expanded class="mb-6 mt-4" tag="nuxt-link" to="/addbusiness/step1" type="is-primary">Zaczynajmy!</b-button>
    </div>
  </section>
</template>

<script lang="ts">
import roundCard from "@/components/roundCard.vue"
export default {
    name: "addobject",
    layout: "background",
    components: {
      roundCard
    }
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_addbusiness.scss";
.add-business .steps {
  width: 40%;
  min-width: 300px;
  margin: 0 auto;
}
@media (max-width: 310px){
  .add-business .steps {min-width: 250px;}
}
</style>