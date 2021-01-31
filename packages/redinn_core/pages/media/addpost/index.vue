<template>
  <div class="addpost">
    <Navbar :curPage="{ name: 'Media', to: '/media' }" style="position: relative" />

    <section class="p-3 is-flex is-mobile is-justify-content-space-between">
      <b-button
        type="is-primary"
        :icon-left="nav[curView.name].left.icon"
        @click="switchView(nav[curView.name].left.to)"
        >{{ nav[curView.name].left.text }}</b-button
      >
      <div class="is-flex is-align-items-center">
        <RoundCard dim="2.5em"><img src="@/assets/imgs/property.jpg" alt="property image" /></RoundCard>
        <div class="ml-1">
          <h3>Zygmuntówka</h3>
          <p class="is-paragraph">@villazygmuntowka</p>
        </div>
      </div>
      <b-button
        type="is-primary"
        :icon-right="nav[curView.name].right.icon"
        @click="switchView(nav[curView.name].right.to)"
        >{{ nav[curView.name].right.text }}</b-button
      >
    </section>
    <keep-alive>
      <component :is="curView" class="view" @switchView="switchView($event)" />
    </keep-alive>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import Navbar from "@/components/nav/Navbar.vue";
import Upload from "@/components/addcontent/Upload.vue";
import PhotoEditor from "@/components/addcontent/PhotoEditor.vue";
import Share from "@/components/addcontent/Share.vue";
import RoundCard from "@/components/RoundCard.vue";

class Navbtn {
  text: string;
  icon: string;
  to: string;
  constructor(text: string, icon: string, to: string = "") {
    this.text = text;
    this.icon = icon;
    this.to = to;
  }
}
@Component({ components: { Navbar, PhotoEditor, RoundCard } })
export default class Addpost extends Vue {
  curView: typeof PhotoEditor | typeof Upload | typeof Share = Upload;
  nav = {
    // Editpost = PhotoEditor
    Editpost: {
      left: new Navbtn("Resetuj zmiany", "close", "Share"),
      right: new Navbtn("Zapisz i wróć", "chevron-double-up", "Share"),
    },
    Upload: {
      left: new Navbtn("Anuluj", "close", "/media"),
      right: new Navbtn("Kolejny krok", "chevron-double-right", "Share"),
    },
    Share: {
      left: new Navbtn("Wróć", "chevron-double-left", "Upload"),
      right: {},
    },
  };
  switchView(event: string) {
    if (event === "PhotoEditor") this.curView = PhotoEditor;
    if (event === "Upload") this.curView = Upload;
    if (event === "Share") this.curView = Share;
    if (event[0] === "/") this.$router.push({ path: event });
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
