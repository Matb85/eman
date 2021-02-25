<template>
  <section class="background-layout">
    <Nuxt class="main" />
    <Logo />
    <section class="fins">
      <LogoIcon />
      <LogoIcon />
    </section>
  </section>
</template>

<script lang="ts">
import LogoIcon from "@/components/logo/LogoIcon.vue";
import Logo from "@/components/logo/Logo.vue";

export default {
  name: "BackgroundFullW",
  components: {
    LogoIcon,
    Logo,
  },
};
</script>

<style lang="scss">
.background-layout {
  width: 100%;
  height: 100%;
  > .main {
    position: relative;
    top: 8rem;
    left: 5rem;
    width: 100%;
    padding: 0 0.75rem;
    max-width: 24rem;
  }
  > .logo-con {
    position: absolute;
    top: 1.5rem;
    left: 1.5rem;
  }
  > .fins {
    position: absolute;
    z-index: -10;
    right: 0;
    top: 0;
    height: 100%;
    width: auto;
    > .logoicon {
      opacity: 0.7;
      height: 100%;
      width: auto;
      position: absolute;
      top: 0;
      right: 0;
      &:last-of-type {
        transform: scaleY(-1);
      }
    }
  }
}
@media (max-width: $tablet) {
  .background-layout > .main {
    left: initial;
    margin: 0 auto;
  }
  .background-layout > .fins > .logoicon {
    opacity: 0.7;
    &:last-of-type {
      display: none;
    }
  }
}
</style>
