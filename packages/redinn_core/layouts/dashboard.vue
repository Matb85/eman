<template>
  <section>
    <Navbar :curPage="curPage" :links="links" />

    <Sidebar :links="links">
      <div v-for="branch in $store.state[$route.name].subtreelinks" :key="branch.name">
        <hr class="my-2" />
        <h3>{{ branch.name }}</h3>
        <div class="routes">
          <ColorfulBtn
            v-for="sublink in branch.sublinks"
            :key="sublink.name"
            tag="nuxt-link"
            :to="sublink.to"
            :type="sublink.type"
            expanded
            rounded
            outlined
            class="my-2"
            :addedclass="sublink.class"
            :icon="sublink.icon"
            :customIcon="sublink.customIcon"
            >{{ sublink.name }}</ColorfulBtn
          >
        </div>
      </div>
    </Sidebar>
    <Nuxt />
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { PrimeRouteI } from "@/store";
import Navbar from "@/components/nav/Navbar.vue";
import Sidebar from "@/components/nav/Sidebar.vue";
import ColorfulBtn from "@/components/ColorfulBtn.vue";
@Component({
  components: {
    Navbar,
    Sidebar,
    ColorfulBtn,
  },
})
export default class Dashboard extends Vue {
  get links() {
    return this.$store.state.dashboards.filter((link: PrimeRouteI) => link.name !== this.curPage.name);
  }
  get curPage() {
    const nm = this.$route.name as string;
    return {
      name: nm[0].toUpperCase() + nm.substring(1),
      to: this.$route.path,
    };
  }
}
</script>

<style lang="scss" scoped>
section > ::v-deep .content-con {
  width: calc(100% - var(--sidew));
  height: calc(100% - var(--navh));
  position: absolute;
  bottom: 0;
  right: 0;
  padding: 1em;
  overflow-y: auto;
  > section,
  > div {
    margin-bottom: 2rem;
  }
}

@media (max-width: $sidebarWBP) {
  section > ::v-deep .content-con {
    left: 0;
    width: 100%;
  }
}
@media (max-width: 370px) {
  section > ::v-deep .content-con {
    --top: 62px;
  }
}
</style>
