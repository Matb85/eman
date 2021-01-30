<template>
  <div class="addpost">
    <Navbar :curPage="{ name: 'Media', to: '/media' }" style="position: relative" />
    <keep-alive>
      <component :is="currentView" class="view" @switchView="switchView($event)" />
    </keep-alive>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import Navbar from "@/components/nav/Navbar.vue";
import Upload from "@/components/addcontent/Upload.vue";
import PhotoEditor from "@/components/addcontent/PhotoEditor.vue";
import Share from "@/components/addcontent/Share.vue";

@Component({ components: { Navbar, PhotoEditor } })
export default class Addpost extends Vue {
  currentView: typeof PhotoEditor | typeof Upload | typeof Share = Upload;

  switchView(event: string) {
    switch (event) {
      case "PhotoEditor":
        this.currentView = PhotoEditor;
        break;
      case "Upload":
        this.currentView = Upload;
        break;
      case "Share":
        this.currentView = Share;
        break;
    }
  }
}
</script>
<style scoped>
.addpost {
  width: 100%;
  height: 100%;
}
.view {
  width: 100%;
  height: calc(100% - var(--navh));
}
</style>
