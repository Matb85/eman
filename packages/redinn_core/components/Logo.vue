<template>
  <section class="logo-con">
    <logoIcon fill="#fff"/>
    <a href="#" class="title is-1 redinn-name ml-1 is-unselectable is-clickable">Redinn</a>
  </section>
</template>

<script lang="ts">
import logoIcon from "@/components/svgs/logoIcon.vue"
export default {
  name: "logo",
  components: {
    logoIcon,
  },
}
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
  .redinn-name {align-self: center;}
}
@media (max-width: $desktopWBP) {
  .logo-con {  
    top: 10px;
    left: 10px;
  }
}
</style>
