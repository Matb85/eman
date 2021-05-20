<template>
  <section class="dashbord">
    <Navbar :curPage="curPage" :links="links" />
    <Sidebar :links="links">
      <div v-for="branch in $store.state[curPage.lowerCaseName].subtreelinks" :key="branch.name">
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
import Navbar from "@/components/nav/Navbar.vue";
import Sidebar from "@/components/nav/Sidebar.vue";
import ColorfulBtn from "@/components/ColorfulBtn.vue";
import { PrimeRouteI } from "@/store";

@Component({ components: { Navbar, Sidebar, ColorfulBtn } })
export default class Dashboard extends Vue {
  get links() {
    return this.$store.state.dashboards.filter((link: PrimeRouteI) => link.name !== this.curPage.name);
  }
  get curPage() {
    const rm = this.$route.name as string;
    const nm = rm.slice(rm.indexOf("enterprise") + 11).substring(0, rm.includes("-") ? rm.indexOf("-") : rm.length);
    return {
      lowerCaseName: nm,
      name: nm[0].toUpperCase() + nm.substring(1),
      to: this.$route.path,
    };
  }
}
</script>

<style lang="scss">
.dashbord > .content-con {
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
  .dashbord > .content-con {
    left: 0;
    width: 100%;
  }
}
@media (max-width: 370px) {
  .dashbord > .content-con {
    --top: 62px;
  }
}

//buttons in the tight bottom corner
.br-act-btns {
  position: fixed;
  bottom: 1rem;
  right: 2rem;
  display: flex;
  flex-direction: column;
  width: 6rem;
  transition: width 0.4s;
  > .act-btn {
    i.mdi.mdi-plus {
      font-size: 1.3rem;
    }
    &:nth-of-type(1) {
      position: relative;
      top: 0;
      right: 0;
      z-index: -1;
      opacity: 0;
      transform: translateY(2.5rem);
      transition: opacity 0.4s, transform 0.4s;
    }
    &:nth-of-type(2) {
      overflow: hidden;
      > span:not(.icon) {
        margin-right: -2rem;
        transition: margin-right 0.4s;
        span {
          display: inline-block;
          transform: translateY(2.5rem);
          transition: transform 0.4s;
        }
      }
    }
  }
  &:hover {
    width: 10em;
    > .act-btn {
      &:nth-of-type(1) {
        transform: translateY(-0.5rem);
        opacity: 1;
      }
      &:nth-of-type(2) > span:not(.icon) {
        margin-right: 0;
        span {
          transform: translateY(0);
        }
      }
    }
  }
}
</style>
