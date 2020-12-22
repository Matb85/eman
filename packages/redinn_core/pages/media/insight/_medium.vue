<template>
<section class="content-con">
        <section class="grid-con">
            <div class="grid-tile tile-is-3">{{this.medium}}</div>
            <div class="grid-tile tile-is-2">dsfgds</div>
            <div class="grid-tile tile-is-5">dsfgds</div>
            <div class="grid-tile">dsfgfds</div>
            <div class="grid-tile tile-is-4">dsfgfds</div>
        </section>

        <h3 class="title is-1 has-text-centered">Podejmij działanie</h3>
        
        <section class="vertical-con">
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
        </section>
        <div class="br-act-btns">
            <b-button class="act-btn px-5" rounded type="is-primary" icon-left="plus">Dodaj Relację</b-button>
            <b-button class="act-btn px-5" rounded type="is-primary" icon-left="plus">Dodaj <span> Post</span></b-button>
        </div>
    </section>
</template>

<script>
export default {
  name: "medium",
  layout: "dashboard",
    async asyncData({ params, error, store }) {
        if(! store.state.media.media.includes(params.medium)) error({statusCode: 404, message: "This page could not be found"})
      const medium = params.medium
      return { medium }
    }
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_dashboards.scss";
</style>
