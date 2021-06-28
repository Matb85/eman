<template>
  <div class="addpost">
    <Navbar :curPage="{ name: 'Media', to: '/media' }" style="position: relative" />

    <section class="p-3 is-flex is-mobile is-justify-content-space-between is-relative">
      <b-button type="is-primary" :icon-left="nav[curView].left.icon" @click="switchView(nav[curView].left.to)">{{
        nav[curView].left.text
      }}</b-button>
      <div class="is-flex is-justify-content-center absolute-card">
        <div class="is-flex is-align-items-center mt-3">
          <RoundCard dim="2.5em"><img src="@/assets/imgs/property.jpg" alt="property image" /></RoundCard>
          <div class="ml-1">
            <h3>Zygmuntówka</h3>
            <p class="is-paragraph">@villazygmuntowka</p>
          </div>
        </div>
      </div>
      <b-button type="is-primary" :icon-right="nav[curView].right.icon" @click="switchView(nav[curView].right.to)">{{
        nav[curView].right.text
      }}</b-button>
    </section>
    <NuxtChild class="view" @switchView="switchView($event)" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import Navbar from "@/components/nav/Navbar.vue";

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
@Component({ components: { Navbar, RoundCard } })
export default class Addpost extends Vue {
  get curView() {
    return this.$route.path.slice(this.$route.path.lastIndexOf("/") + 1);
  }
  nav = {
    edit: {
      left: new Navbtn("Resetuj zmiany", "close", "/media/addpost/share"),
      right: new Navbtn("Zapisz i wróć", "chevron-double-up", "/media/addpost/upload"),
    },
    upload: {
      left: new Navbtn("Anuluj", "close", "/media"),
      right: new Navbtn("Kolejny krok", "chevron-double-right", "/media/addpost/share"),
    },
    share: {
      left: new Navbtn("Wróć", "chevron-double-left", "/media/addpost/upload"),
      right: new Navbtn("Opublikuj", "fire", "/media"),
    },
  };
  switchView(event: string) {
    this.$router.push({ path: event });
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
  height: calc(100% - var(--navh) * 2 + 4px);
}
.absolute-card {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  z-index: -1;
}
</style>
