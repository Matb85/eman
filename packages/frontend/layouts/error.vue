<template>
  <div class="apl-form">
    <svg xmlns="http://www.w3.org/2000/svg" width="90" height="90" fill="#FF3030" viewBox="0 0 48 48">
      <path
        d="M22 30h4v4h-4zm0-16h4v12h-4zm1.99-10C12.94 4 4 12.95 4 24s8.94 20 19.99 20S44 35.05 44 24 35.04 4 23.99 4zM24 40c-8.84 0-16-7.16-16-16S15.16 8 24 8s16 7.16 16 16-7.16 16-16 16z"
      />
    </svg>

    <div class="title">{{ message }}</div>
    <div class="title">{{ error.statusCode }}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import { NuxtError } from "@nuxt/types/app/index";

@Component({ layout: "background" })
export default class Error extends Vue {
  @Prop() error: NuxtError;

  get statusCode() {
    return (this.error && this.error.statusCode) || 500;
  }
  get message() {
    return this.error.message || "unknown error occured";
  }
  head() {
    return {
      title: this.message,
    };
  }
}
</script>
<style scoped lang="scss">
@import "@/assets/scss/partials/_loginregister.scss";
</style>
