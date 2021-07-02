<template>
  <section class="upload-con">
    <section class="preview px-3">
      <div ref="imgs" class="columns mt-0 is-mobile">
        <div v-for="img in uploadedImgs" :key="img.undefined" class="column img-act-con" :data-id="img.id">
          <div class="columns is-vcentered is-centered is-multiline is-align-content-center">
            <div class="m-2 is-clickable" @click="$emit('switchView', 'PhotoEditor')">
              <b-icon icon="image-edit-outline" type="is-light" size="is-large" />
              <p class="is-paragraph has-text-centered has-text-white">Edit</p>
            </div>
            <div class="m-2 is-clickable">
              <b-icon icon="trash-can-outline" type="is-light" size="is-large" />
              <p class="is-paragraph has-text-centered has-text-white">{{ img.order }}</p>
            </div>
          </div>
          <img :src="img.src" :alt="img.name" />
        </div>
        <div v-if="uploadedImgs.length <= 10" class="column py-5 is-flex upload-form">
          <b-upload accept=".jpg,.jpeg,.png" drag-drop multiple class="is-flex-grow-0" @input="addImg($event)">
            <section class="section">
              <div class="content has-text-centered">
                <b-icon icon="upload" size="is-large"></b-icon>
                <p>Upuść zdjęcia lub kliknij aby je dodać</p>
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
  readonly name: string;
  readonly id: number;
  order: number;
}
@Component
export default class Upload extends Vue {
  $refs!: { imgs: HTMLElement };

  uploadedImgs: Array<imgI> = [];
  addImg(imgs: FileList) {
    const status: Promise<void>[] = [];
    for (let i = inputCounter; i < imgs.length; i++) {
      status.push(this.readImg(imgs[i]));
    }
    inputCounter = imgs.length;
    // wait for all the imgs to be read and pushed into this.uploadedImgs
    Promise.allSettled(status).then(() => {
      const a = this.uploadedImgs.map((i) => {
        return { id: i.id, name: i.name };
      });
      // save a map of images in case of a page reload
      sessionStorage.setItem("uploadedimgsmap", JSON.stringify(a));
    });
  }
  readImg(img: File) {
    return new Promise<void>((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => {
        sessionStorage.setItem("uploadedimg" + this.uploadedImgs.length, reader.result as string);
        this.uploadedImgs.push({
          src: reader.result,
          name: img.name,
          id: this.uploadedImgs.length,
          order: this.uploadedImgs.length,
        });
        resolve();
      };
      reader.onerror = reject;
      reader.readAsDataURL(img);
    });
  }
  mounted() {
    console.log(this.$route);
    // it has to be a distinct function because of scoping interference
    const sortImgs = (order: number[]) => {
      console.log(order);
      console.log(this.uploadedImgs);
      let j = 0;
      for (let i = 0; i < order.length - 1; i++) {
        this.uploadedImgs[order[i]].order = j;
        j++;
      }
      console.log(this.uploadedImgs.map((i) => i.order));

      // update the map of images in case of a page reload -- see addImg()
      sessionStorage.setItem(
        "uploadedimgsmap",
        JSON.stringify(
          [...this.uploadedImgs].sort(function (a: imgI, b: imgI) {
            if (a.order > b.order) return 1;
            if (b.order > a.order) return -1;
            return 0;
          })
        )
      );
    };
    Sortable.create(this.$refs.imgs, {
      filter: ".upload-form",
      direction: "horizontal",
      animation: 150,
      dataIdAttr: "data-id",
      store: {
        get(sortable) {
          console.log(sortable);
          return [""];
        },
        set(sortable) {
          sortImgs(sortable.toArray().map((x) => parseInt(x)));
        },
      },
    });

    if (!sessionStorage.getItem("uploadedimgsmap")) return;
    const uploadedimgsmap = JSON.parse(sessionStorage.getItem("uploadedimgsmap") as string);
    this.uploadedImgs = uploadedimgsmap.map((x: imgI): imgI => {
      console.log("uploadedimg" + x.id);
      return {
        src: sessionStorage.getItem("uploadedimg" + x.id),
        id: x.id,
        name: x.name,
        order: x.order,
      };
    });
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
