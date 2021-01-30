<template>
  <div
    id="round-card"
    class="has-text-white has-background-primary"
    :style="{ width: dim, height: dim, 'line-height': dim, 'font-size': fontSize }"
  >
    <slot />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";

@Component
export default class RoundCard extends Vue {
  @Prop({ required: true }) dim: string;
  @Prop() fontSize: string;
}
</script>

<style scoped lang="scss">
#round-card {
  position: relative;
  margin: 0 auto;
  border-radius: 50%;
  text-align: center;
  overflow: hidden;
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
</style>
