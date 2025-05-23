<template>
  <nav id="navbar">
    <div id="start">
      <div id="brand" class="mx-2">
        <LogoIcon />
        <nuxt-link :to="curPage.to" class="eman-name title is-3 ml-1 is-unselectable"
          >Eman {{ curPage.name }}</nuxt-link
        >
      </div>
      <div id="links">
        <nuxt-link v-for="link in links" :key="link.name" class="has-text-dark is-size-5 mx-2" :to="link.to">
          Eman {{ link.name }}
        </nuxt-link>
      </div>
    </div>
    <div id="end" class="is-unselectable">
      <div id="dialog" class="mx-2">
        <h2 class="title is-4">Witaj {{ $auth.user.firstName }}</h2>
        <p>Co słychać?</p>
      </div>
      <RoundCard dim="2.2em" fontSize="1.4em" @click.native="profileMenu = !profileMenu">{{
        $auth.user.firstName ? $auth.user.firstName[0] : ""
      }}</RoundCard>
      <MenuBurger @click.native="toggleSidebar" />
      <div v-show="profileMenu" id="profile-menu">
        <div class="current-profile mb-3">
          <RoundCard dim="2.2em" fontSize="1.4em">M</RoundCard>
          <div class="profile-basic-info ml-2">
            <p class="title is-5 mb-0">{{ $auth.user.firstName + " " + $auth.user.lastName }}</p>
            <p class="has-text-weight-bold is-family-secondary">{{ $auth.user.email }}</p>
          </div>
        </div>
        <b-menu-list>
          <b-menu-item icon="cog" label="Ustawienia"></b-menu-item>
          <b-menu-item icon="account" label="Profil"></b-menu-item>
          <b-menu-item icon="information-outline" label="Pomoc"></b-menu-item>
          <b-menu-item icon="logout" label="Wyloguj" @click="$auth.logout()"></b-menu-item>
        </b-menu-list>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import { PrimeRouteI } from "@/store";

import MenuBurger from "@/components/svgs/menuBurger.vue";
import LogoIcon from "@/components/logo/LogoIcon.vue";
import RoundCard from "@/components/roundCard.vue";

@Component({
  components: {
    LogoIcon,
    RoundCard,
    MenuBurger,
  },
})
export default class Navbar extends Vue {
  @Prop() curPage: PrimeRouteI;
  @Prop() links: Array<PrimeRouteI>;

  profileMenu = false;
  toggleSidebar() {
    this.$root.$emit("toggle sidebar");
  }
}
</script>

<style scoped lang="scss">
#navbar {
  font-size: 15px;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 2;
  background-color: #fff;
  width: 100%;
  height: 4.8em;
  display: flex;
  justify-content: space-between;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.25);
  * {
    align-self: center;
  }
  #start {
    display: flex;
    justify-content: center;
    align-content: flex-start;
    #brand {
      display: flex;
      justify-content: flex-start;
      svg {
        height: 4em;
        width: auto;
      }
      .eman-name {
        line-height: 2.5em;
        position: relative;
        &::after {
          content: "";
          display: block;
          width: 100%;
          height: 1px;
          background-color: $primary;
          position: absolute;
          bottom: 0.6em;
        }
      }
    }
    #links {
      display: flex;
      justify-content: center;
      align-content: flex-start;
      > a {
        margin-top: 0.3em;
        &:hover {
          color: $primary !important;
        }
      }
    }
  }
  #end {
    display: flex;
    justify-content: flex-end;
    margin-right: 1.5em;
    #dialog {
      display: flex;
      flex-direction: column;
      * {
        text-align: right;
        align-self: flex-end;
        line-height: 0;
      }
    }
    #round-card,
    .menu-burger {
      cursor: pointer;
    }
    .menu-burger {
      display: none;
      margin-left: 0.5em;
    }
    #profile-menu {
      z-index: -1;
      position: absolute;
      top: 5.5em;
      right: 1.5em;
      padding: 1.5em 3em;
      height: auto;
      width: auto;
      background-color: white;
      border-radius: 0.5em;
      box-shadow: 0 0 6px rgba(0, 0, 0, 0.25);
      .current-profile {
        display: flex;
        width: 100%;
        * {
          flex: 0 0 auto;
        }
      }
    }
  }
}
@media (max-width: $sidebarWBP) {
  #navbar {
    #start #links {
      display: none;
    }
    #end .menu-burger {
      display: block;
    }
  }
}
@media (max-width: 550px) {
  #navbar #end {
    margin-right: 0;
    #dialog {
      display: none;
    }
    .menu-burger {
      margin-left: 0;
    }
  }
}
@media (max-width: 370px) {
  #navbar {
    font-size: 13px;
  }
}
</style>
