<template>
  <div class="photoeditor">
    <client-only placeholder="Loading...">
      <photoEditor />
    </client-only>
  </div>
</template>

<script>
import '@matb85/photoeditor/dist/photoEditor.css'
export default {
  name: 'upload',
  components: {
    photoEditor: () => {
      if (process.browser) return import('@matb85/photoeditor')
    },
  },
}
</script>
<style scoped>
.photoeditor {
  width: 100%;
  height: 100%;
}
</style>
