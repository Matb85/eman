<template>
    <section class="back-con">
        <div class="background">
            <logo></logo>
            <div id="arrow-con">
            <div v-for="arrow in arrows" :key="arrow.id">
                <doublearrow/>
            </div>
            </div>
            <logoIcon fill="#fff"/>
        <div class="white-cover"></div>
        </div>
        <nuxt/>
    </section>
</template>

<script lang="ts">
import background from "@/mixins/background.js";
export default {
  mixins: [background],
  name: "backgroundFullW",
}
</script>

<style scoped lang="scss">
@import "@/assets/components/_background.scss";
.back-con {
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: #fff;
    .background {
        height: 100vh;
        top: 0;
        left: 0; 
            background: rgb(255,48,48);
            background: linear-gradient(180deg, rgba(255,48,48,1) 99%, rgba(255,255,255,1) 100%);
        > ::v-deep #logoicon  {
            min-width: 100%;
        }
    }
    .white-cover{
        position: absolute;
        bottom: 0%;
        display: block;
        width: 100%;
        height: 1px;
        background-color: #fff;
    }
    > *:not(.background){position: absolute;}
}
@media (max-aspect-ratio: 1/1){
    .back-con {
        .background > ::v-deep #logoicon {
            max-width: 120vh;
        }
        .white-cover {
            height: 40%;
            background-color: #fff; 
            &::after {
                content: "";
                display: block;
                position: absolute;
                top: 50%;
                left: 85%;
                width: 100%;
                height: 500px;
                max-width: 1000px;
                background-color: #fff;
                transform: translate(-52%, -50%) rotate(-50deg);
            }
        }
    }
}
</style>
