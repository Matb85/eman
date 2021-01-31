<template>
  <section class="upload-con">
    <section class="preview px-3">
      <div ref="imgs" class="columns mt-0 is-mobile">
        <div v-for="img in uploadedImgs" :key="img.id" class="column img-act-con" :data-id="img.id">
          <div class="columns is-vcentered is-centered is-multiline is-align-content-center">
            <div class="m-2 is-clickable" @click="$emit('switchView', 'PhotoEditor')">
              <b-icon icon="image-edit-outline" type="is-light" size="is-large" />
              <p class="is-paragraph has-text-centered has-text-white">Edit</p>
            </div>
            <div class="m-2 is-clickable">
              <b-icon icon="trash-can-outline" type="is-light" size="is-large" />
              <p class="is-paragraph has-text-centered has-text-white">{{ img.id }}</p>
            </div>
          </div>
          <img src="@/assets/imgs/property.jpg" :alt="img.name" />
        </div>
        <div v-if="uploadedImgs.length <= 10" class="column py-6 is-flex upload-form">
          <b-upload accept=".jpg,.jpeg,.png" drag-drop multiple class="is-flex-grow-0" @input="addImg($event)">
            <section class="section">
              <div class="content has-text-centered">
                <b-icon icon="upload" size="is-large"></b-icon>
                <p>Drop your files here or click to upload</p>
              </div>
            </section>
          </b-upload>
        </div>
      </div>
    </section>
    <h2 class="has-text-centered title is-4 mt-1 mb-4">Filtry</h2>
    <section class="filters preview">
      <div class="columns is-mobile">
        <div v-for="filter in [1, 2, 3, 4, 5, 6, 7]" :key="filter" class="column p-2">
          <div class="is-clipped">
            <img src="@/assets/imgs/property.jpg" alt="property image" />
          </div>
          <p class="has-text-centered">Default</p>
        </div>
      </div>
    </section>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import Sortable from "sortablejs";
let inputCounter = 0;
interface imgI {
  src: string | ArrayBuffer | null;
  name: string;
  id: number;
}
@Component
export default class Upload extends Vue {
  $refs!: {
    imgs: HTMLElement;
  };
  uploadedImgs: Array<imgI> = [
    { id: 0, src: "", name: "0" },
    { id: 1, src: "", name: "1" },
    { id: 2, src: "", name: "2" },
  ];
  addImg(imgs: FileList) {
    for (let i = inputCounter; i < imgs.length; i++) {
      const reader = new FileReader();
      reader.onload = () => {
        this.uploadedImgs.push({
          src: reader.result,
          name: imgs[i].name,
          id: this.uploadedImgs.length,
        });
      };
      reader.readAsDataURL(imgs[i]);
    }
    inputCounter = imgs.length;
  }

  mounted() {
    const sortImgs = (order: number[]) => {
      let j = 0;
      for (let i = 0; i < order.length - 1; i++) {
        this.uploadedImgs[order[i]].id = j;
        j++;
      }
    };
    Sortable.create(this.$refs.imgs, {
      filter: ".upload-form",
      direction: "horizontal",
      animation: 150,
      dataIdAttr: "data-id",
      store: {
        get(sortable) {
          console.log(sortable);
        },
        set(sortable) {
          sortImgs(sortable.toArray().map((x) => parseInt(x)));
        },
      },
    });
    this.uploadedImgs[0].id = 1;
    this.uploadedImgs[1].id = 0;
  }
}
</script>

<style lang="scss" scoped>
.notification {
  width: 300px;
  margin: auto;
  position: relative;
  top: 50%;
  .mdi-spin::before {
    animation-duration: 750ms;
  }
}

.img-act-con {
  cursor: grab;
  position: relative;
  > div.columns {
    position: absolute;
    width: calc(100% - #{$column-gap * 2});
    height: calc(100% - #{$column-gap * 2});
    top: $column-gap * 2;
    left: $column-gap * 2;
    background-color: rgba(0, 0, 0, 0.4);
    opacity: 1;
    transition: 0.4s opacity;
  }
  &:hover div {
    opacity: 1;
  }
}

.preview {
  height: 50%;
  width: 100%;
  overflow-x: auto;
  overflow-y: hidden;
  &::-webkit-scrollbar {
    display: none;
  }
  .columns {
    width: 100%;
    height: 100%;
    .column {
      height: 100%;
      flex: 0 0 auto;
      img {
        height: 100%;
        flex: 0;
      }
    }
  }
}

.filters {
  height: 6.5rem;
  width: 60%;
  margin: auto;
  > .columns .column div {
    height: 100%;
    img {
      filter: brightness(100%) blur(10px);
      transform: scale(1.05);
      transition: filter 0.4s;
      &:hover {
        filter: brightness(150%) blur(10px);
      }
    }
  }
}
</style>
