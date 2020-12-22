<template>
  <b-button
    :tag="tag"
    :to="to"
    class="b-btn"
    expanded
    outlined
    rounded
    :class="addedclass"
    :type="type"
    :icon-left="icon"
  >
    <img v-if="customIcon" :src="require('@/assets/imgs/logos/' + customIcon + '.svg')" /><slot />
  </b-button>
</template>

<script>
export default {
  name: "colorfulbtn",
  props: {
    icon: {
      required: false,
      type: String,
    },
    customIcon: {
      required: false,
      type: String,
    },
    addedclass: {
      required: false,
      type: String,
    },
    type: {
      default: "is-primary",
      required: false,
      type: String,
    },
    tag: {
      required: false,
      type: String,
    },
    to: {
      required: false,
      type: String,
    },
  },
};
</script>

<style lang="scss">
.b-btn {
  img {
    position: absolute;
    height: 1em;
    top: calc(50% - 0.5em);
    left: 1em;
    z-index: 9;
  }
  .icon {
    position: absolute;
    left: 1.3em;
  }
}
.backlit::after {
  content: "";
  position: absolute;
  display: block;
  height: 1.4em;
  width: 1.4em;
  top: calc(50% - 0.7em);
  left: 0.8em;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.5);
  z-index: 8;
}
.whiteicon:hover img {
  filter: grayscale(1) brightness(10000%);
}
.instagram {
  padding: 0 !important;
  border: none !important;
  background: linear-gradient(225deg, #d800b9 0%, #ff0100 50%, #ffd600 100%);
  span:not(.icon) {
    width: 100%;
    height: 100%;
    padding: calc(0.5em) calc(1.25em);
    background-color: white;
    color: black;
    border-radius: $radius-rounded;
    mix-blend-mode: screen;
    border: 1px solid rgb(0, 0, 0);
  }
  &:hover span:not(.icon) {
    background-color: rgba(255, 255, 255, 0);
    color: white;
    border: none;
    mix-blend-mode: initial;
  }
}
</style>
