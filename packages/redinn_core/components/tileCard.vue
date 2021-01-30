<template>
  <div class="service-tile">
    <slot />
  </div>
</template>

<script>
export default {
  name: "TileCard",
};
</script>

<style scoped lang="scss">
.service-tile {
  width: 100%;
  height: auto;
  position: relative;
  box-shadow: 1px 3px 10px 1px rgba(0, 0, 0, 0.25);
  border-radius: 2em;
  padding: 1.5em;
  img {
    margin-bottom: 1.5em;
    width: 100%;
    height: auto;
    object-fit: cover;
  }
}
</style>
