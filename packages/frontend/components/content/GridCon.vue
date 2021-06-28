<template>
  <section class="grid-con">
    <slot />
  </section>
</template>

<script>
export default {
  name: "GridCon",
};
</script>

<style lang="scss" scoped>
.grid-con {
  position: relative;
  width: 100%;
  height: auto;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: min(100px);
  grid-gap: 1em;
}
.tile-is-2 {
  grid-column: span 2;
}
.tile-is-3 {
  grid-column: span 3;
}
.tile-is-4 {
  grid-column: span 4;
}
.tile-is-5 {
  grid-column: span 5;
}
</style>
