<template>
  <section ref="sidebar" class="sidebar">
    <div id="property" class="mt-4">
      <b-button type="is-text" size="is-small" class="has-background-white px-0 py-0" active>Zmień obiekt</b-button>
      <RoundCard dim="120px"><img src="@/assets/imgs/property.jpg" alt="property image" /></RoundCard>
      <h3 class="is-title is-4">Zygmuntówka</h3>
      <p class="has-text-weight-bold is-family-secondary">Ul. Franciszka Koterby Kuriera AK 37</p>
    </div>
    <div id="routes-con" class="px-3">
      <span class="nav-mobile">
        <hr class="my-2" />
        <h3>Nawigacja</h3>
        <b-menu-list>
          <b-menu-item
            v-for="link in links"
            :key="link.name"
            :to="link.to"
            :label="'Redinn ' + link.name"
            tag="nuxt-link"
          ></b-menu-item>
        </b-menu-list>
      </span>

      <slot />
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";

import RoundCard from "@/components/RoundCard.vue";
@Component({
  components: {
    RoundCard,
  },
})
export default class Sidebar extends Vue {
  $refs!: {
    sidebar: HTMLElement;
  };
  @Prop() links: Array<object>;
  created() {
    this.$root.$on("toggle sidebar", () => {
      this.$refs.sidebar.classList.toggle("sidebar-opened");
    });
  }
}
</script>

<style scoped lang="scss">
.sidebar {
  position: absolute;
  bottom: 0;
  left: 0;
  z-index: 1;
  width: var(--sidew);
  height: calc(100% - var(--navh));
  background-color: #fff;
  box-shadow: 1px 0 4px rgba(0, 0, 0, 0.25);
  transition: transform 400ms;
  overflow-y: auto;
}
#property {
  text-align: center;
  position: relative;
  width: 90%;
  margin: 0 auto;
  height: auto;
  button {
    position: absolute;
    top: -0.8em;
    left: -0.6em;
    opacity: 0.6;
    &:focus {
      box-shadow: none;
    }
    &:hover {
      opacity: 1;
    }
  }
}
.nav-mobile {
  display: none;
}
#routes-con {
  width: 100%;
  height: auto;
}

@media (max-width: $sidebarWBP) {
  .sidebar {
    transform: translateX(-250px);
  }
  .sidebar.sidebar-opened {
    transform: translateX(0);
  }
  .nav-mobile {
    display: block;
  }
}
</style>
