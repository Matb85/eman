<template>
  <div class="background">
    <logo></logo>
    <div id="arrow-con">
      <div v-for="arrow in arrows" :key="arrow.id">
        <doublearrow/>
      </div>
    </div>
    <logoIcon fill="#fff"/>
    <nuxt/>
  </div>
</template>

<script lang="ts">
import background from "@/mixins/background.js";
export default {
  mixins: [background],
  name: "background",
  
  
}
</script>


<style scoped lang="scss">
 @import "@/assets/components/_background.scss";
</style>
