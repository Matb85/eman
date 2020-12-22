<template>
    <section class="content-con">
        <section class="grid-con">
            <div class="grid-tile tile-is-3">sags</div>
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

<script lang="ts">
export default {
    name: "Media",
    layout: "dashboard",
    components: {
    },
}
</script>

<style lang="scss" scoped>
@import "@/assets/components/_dashboards.scss";
</style>
