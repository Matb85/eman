<template>
  <section class="px-6 pt-3 tile">
    <div class="tile is-6 is-vertical pr-2">
      <h3 class="title is-4 mb-2">Tekst postu:</h3>
      <b-input maxlength="200" type="textarea"></b-input>
      <h3 class="title is-4 mb-2">Opublikuj w:</h3>
      <div class="columns is-multiline m-0">
        <div v-for="sublink in $store.state.media.subtreelinks[1].sublinks" :key="sublink.name" class="p-1 column is-6">
          <ColorfulBtn
            tag="nuxt-link"
            :to="sublink.to"
            :type="sublink.type"
            expanded
            rounded
            outlined
            :addedclass="sublink.class"
            :icon="sublink.icon"
            :customIcon="sublink.customIcon"
            >{{ sublink.name }}</ColorfulBtn
          >
        </div>
      </div>
    </div>
    <div class="tile is-6">
      <img src="@/assets/imgs/property.jpg" alt="property image" style="object-fit: contain" />
    </div>
  </section>
</template>

<script>
import ColorfulBtn from "@/components/ColorfulBtn.vue";

export default {
  name: "Share",
  components: { ColorfulBtn },
};
</script>
