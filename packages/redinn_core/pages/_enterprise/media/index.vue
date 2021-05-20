<template>
  <section class="content-con">
    <GridCon class="grid-con">
      <div class="grid-tile tile-is-3">sags</div>
      <div class="grid-tile tile-is-2">dsfgds</div>
      <div class="grid-tile tile-is-5">dsfgds</div>
      <div class="grid-tile">dsfgfds</div>
      <div class="grid-tile tile-is-4">dsfgfds</div>
    </GridCon>

    <h3 class="title is-1 has-text-centered">Podejmij działanie</h3>

    <VerticalCon class="vertical-con">
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
    <div class="br-act-btns">
      <b-button class="act-btn px-5" tag="nuxt-link" to="/media/addpost" type="is-primary" icon-left="plus"
        >Dodaj Relację</b-button
      >
      <b-button class="act-btn px-5" tag="nuxt-link" to="/media/addpost" type="is-primary" icon-left="plus"
        >Dodaj <span>Post</span></b-button
      >
    </div>
  </section>
</template>

<script lang="ts">
import GridCon from "@/components/content/GridCon.vue";
import VerticalCon from "@/components/content/VerticalCon.vue";
export default {
  name: "Media",
  layout: "dashboard",
  components: { GridCon, VerticalCon },
};
</script>
