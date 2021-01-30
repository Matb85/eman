<template>
  <section class="vertical-con"><slot /></section>
</template>

<script>
export default {
  name: "VerticalCon",
};
</script>

<style lang="scss" scoped>
.vertical-con {
  position: relative;
  width: 100%;
  height: auto;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: auto;
  grid-gap: 1em;
  .vertical-flex {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: auto;
    > * {
      margin-bottom: 1em;
    }
  }
}
</style>
