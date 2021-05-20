<template>
  <section class="content-con">
    <GridCon>
      <div class="grid-tile tile-is-3">sags</div>
      <div class="grid-tile tile-is-2">dsfgds</div>
      <div class="grid-tile tile-is-5">dsfgds</div>
      <div class="grid-tile">dsfgfds</div>
      <div class="grid-tile tile-is-4">dsfgfds</div>
    </GridCon>
    <h3 class="title is-1 has-text-centered">Podejmij działanie</h3>
    <VerticalCon>
      <div class="vertical-flex">
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
      </div>
      <div class="vertical-flex">
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
      </div>
      <div class="vertical-flex">
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
        <div class="grid-tile">hi</div>
      </div>
    </VerticalCon>
  </section>
</template>

<script lang="ts">
import GridCon from "@/components/content/GridCon.vue";
import VerticalCon from "@/components/content/VerticalCon.vue";

export default {
  name: "Core",
  layout: "dashboard",
  components: { GridCon, VerticalCon },
  // async asyncData({ route, $axios }) {
  //   const response = await $axios.$post("/graphql", {
  //     query: `query get{enterprise(index: ${route.params.enterprise}){ id name address {street}}}`,
  //   });
  //
  //   console.log(response);
  // },
  mounted() {
    console.log(this.$auth.enterprise);
  },
};
</script>
