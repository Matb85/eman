<template>
  <section class="content-con">
    <GridCon>
      <div class="grid-tile tile-is-3">{{ medium }}</div>
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
    <div class="br-act-btns">
      <b-button class="act-btn px-5" rounded type="is-primary" icon-left="plus">Dodaj Relację</b-button>
      <b-button class="act-btn px-5" rounded type="is-primary" icon-left="plus">Dodaj <span> Post</span></b-button>
    </div>
  </section>
</template>

<script>
import GridCon from "@/components/content/GridCon.vue";
import VerticalCon from "@/components/content/VerticalCon.vue";

export default {
  name: "Medium",
  layout: "dashboard",
  components: { GridCon, VerticalCon },
  asyncData({ params, error, store }) {
    console.log(params);
    if (!store.state.media.media.includes(params.medium))
      error({ statusCode: 404, message: "This page could not be found" });
    const medium = params.medium;
    return { medium };
  },
};
</script>
