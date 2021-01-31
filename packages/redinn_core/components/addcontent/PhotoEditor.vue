<template>
  <div class="photo-editor-con is-flex is-flex-direction-column">
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
import "@matb85/photoeditor/dist/photoEditor.css";
export default {
  name: "Editpost",
  components: {
    PhotoEditor: () => {
      if (process.browser) return import("@matb85/photoeditor/dist/photoEditor.umd.js");
    },
  },
  methods: {
    reset() {
      this.$store.commit("photoEditor/reset");
      this.$root.$emit("photoEditor/alterphoto");
    },
  },
};
</script>
