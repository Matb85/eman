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
.content-con .br-act-btns{
    position: fixed;
    bottom: 1rem;
    right: 2rem;
    display: flex;
    flex-direction: column;
    width: 6em;
    transition: width 0.4s;
    .act-btn {
        i.mdi.mdi-plus {
            font-size: 1.3em;
        }
        &:nth-of-type(1){
            position: relative;
            top: 0;
            right: 0;
            z-index: -1;
            opacity: 0;
            transform: translateY(2.5em);
            transition: opacity 0.4s, transform 0.4s;
        }
        &:nth-of-type(2){
            overflow: hidden;
            > ::v-deep span:not(.icon) {
                margin-right: -2em;
                transition: margin-right 0.4s;
                span{
                    display: inline-block;
                    transform: translateY(2.5rem);
                    transition: transform 0.4s;
                }
            }
        }
    }
    &:hover {
        width: 10em;
        .act-btn{
        &:nth-of-type(1){
            transform: translateY(-0.5rem);
            opacity: 1;
        }
        &:nth-of-type(2)> ::v-deep span:not(.icon){margin-right:0;
            span{
            transform: translateY(0);
        }
        }
    }
    }
}
</style>
