<template>
  <section class="back-con">
    <Nuxt />
    <div class="background">
      <Logo />
      <div id="arrow-con">
        <div v-for="arrow in arrows" :key="arrow.id">
          <Doublearrow />
        </div>
      </div>
      <LogoIcon fill="#fff" />
      <div class="white-cover"></div>
    </div>
  </section>
</template>

<script lang="ts">
import Doublearrow from "@/components/svgs/Doublearrow.vue";
import LogoIcon from "@/components/logo/LogoIcon.vue";
import Logo from "@/components/logo/Logo.vue";
class ArrowInstance {
  id: number;
  constructor(id: number) {
    this.id = id;
  }
}
export default {
  name: "backgroundFullW",
  components: {
    Doublearrow,
    LogoIcon,
    Logo,
  },
  data() {
    const arrows = [];
    for (let i = 0; i < 21; i++) {
      arrows.push(new ArrowInstance(i));
    }
    return {
      bclass: "",
      arrows,
    };
  },
};
</script>

<style scoped lang="scss">
.background {
  position: absolute;
  background-color: $primary;
  overflow: hidden;
  width: 100%;
  height: 100%;
  > ::v-deep #logoicon {
    position: absolute;
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    display: block;
    height: auto;
    width: auto;
    min-width: 80%;
    min-height: 102%;
  }
  #arrow-con {
    transform: rotate(-10deg);
    position: absolute;
    top: 8%;
    left: -20%;
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-columns: repeat(5, 25vw);
    grid-template-rows: repeat(5, 25vh);
    div {
      position: relative;
      width: 100%;
      height: 100%;
      ::v-deep .doublearrow {
        position: absolute;
        width: 698px;
        height: auto;
      }
      &:nth-of-type(5n + 1) .doublearrow,
      &:nth-of-type(5n + 3) .doublearrow,
      &:nth-of-type(5n + 5) .doublearrow {
        top: 50%;
      }
    }
  }
  @media (max-height: $desktopHBP), (max-width: 1400px) {
    #arrow-con {
      top: 12%;
      div .doublearrow {
        width: calc(698px * 0.75);
      }
    }
  }
  @media (max-width: $desktopWBP) {
    #arrow-con {
      grid-template-columns: repeat(5, calc(1050px / 4));
    }
  }
  @media (min-width: #{$mobileWBP + 1px}) {
    #arrow-con div:nth-last-of-type(1) {
      top: -50%;
      grid-column: 4;
      grid-row: 5;
    }
  }
  @media (max-width: $mobileWBP) {
    #arrow-con {
      top: 16%;
      left: -40%;
      grid-template-columns: repeat(4, 300px);
      grid-template-rows: repeat(5, 120px);
      div:nth-of-type(5n + 1) .doublearrow,
      div:nth-of-type(5n + 3) .doublearrow,
      div:nth-of-type(5n + 5) .doublearrow {
        top: 0%;
      }
      div:nth-of-type(2n + 1) .doublearrow,
      div:nth-of-type(2n + 3) .doublearrow {
        top: 50%;
      }
    }
  }
  @media (max-aspect-ratio: 1/1) {
    > ::v-deep #logoicon {
      top: 50%;
      width: 240%;
      max-width: 1000px;
      transform: translate(-52%, -50%) rotate(20deg);
    }
  }
}

.back-con {
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: #fff;
  > *:not(.background) {
    position: absolute;
    z-index: 1;
  }
}

.back-con > ::v-deep .background-wide + .background {
  z-index: 0;
  height: 100vh;
  top: 0;
  left: 0;
  background: linear-gradient(180deg, rgba(255, 48, 48, 1) 99%, rgba(255, 255, 255, 1) 100%);
  > #logoicon {
    min-width: 100%;
  }
  .white-cover {
    position: absolute;
    bottom: 0%;
    display: block;
    width: 100%;
    height: 1px;
    background-color: #fff;
  }
  @media (max-aspect-ratio: 1/1) {
    > #logoicon {
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
