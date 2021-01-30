<template>
  <div class="photo-editor-con is-flex is-flex-direction-column">
    <div class="pt-3 px-3 is-flex is-mobile is-justify-content-space-between">
      <b-button type="is-primary" icon-left="close" @click="reset">Resetuj zmiany</b-button>
      <div class="is-flex is-align-items-center">
        <RoundCard dim="2.5em"><img src="@/assets/imgs/property.jpg" alt="property image" /></RoundCard>
        <div class="ml-1">
          <h3>Zygmuntówka</h3>
          <p class="is-paragraph">@villazygmuntowka</p>
        </div>
      </div>
      <b-button type="is-primary" icon-right="chevron-double-up" @click="$emit('switchView', 'Upload')"
        >Zapisz i wróć</b-button
      >
    </div>
    <client-only>
      <PhotoEditor />
      <div slot="placeholder" class="notification is-white has-text-centered">
        <p>We're getting things ready...</p>
        <span class="icon mt-2"><i class="mdi mdi-24px mdi-loading mdi-spin"></i></span>
      </div>
    </client-only>
  </div>
</template>

<script>
import RoundCard from "@/components/RoundCard.vue";
import "@matb85/photoeditor/dist/photoEditor.css";
export default {
  name: "Editpost",
  components: {
    PhotoEditor: () => {
      if (process.browser) return import("@matb85/photoeditor/dist/photoEditor.umd.js");
    },
    RoundCard,
  },
  methods: {
    reset() {
      this.$store.commit("photoEditor/reset");
      this.$root.$emit("photoEditor/alterphoto");
    },
  },
};
</script>
