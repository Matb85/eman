<template>
  <section class="logo-con">
    <LogoIcon :fill="$attrs.fill" />
    <a href="#" class="title is-1 eman-name ml-1 is-unselectable is-clickable">Eman</a>
  </section>
</template>

<script lang="ts">
import LogoIcon from "./LogoIcon.vue";
export default {
  name: "Logo",
  components: {
    LogoIcon,
  },
};
</script>

<style scoped lang="scss">
.logo-con {
  position: absolute;
  display: flex;
  flex-wrap: no-wrap;
  flex-direction: row;
  top: 25px;
  left: 50px;
  svg {
    height: 4em;
    width: auto;
    object-fit: cover;
  }
  .eman-name {
    align-self: center;
  }
}
@media (max-width: $desktopWBP) {
  .logo-con {
    top: 10px;
    left: 10px;
  }
}
</style>
