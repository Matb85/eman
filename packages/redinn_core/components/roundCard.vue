<template>
  <div id="round-card" class="has-text-white has-background-primary" :style="{width: dim, height: dim, 'line-height': dim, 'font-size': fontSize}"><slot/></div>
</template>

<script>
export default {
    name: "roundCard",
    props: {
        dim: {
            required: true,
            type: String,
        },
        fontSize: {
            required: false,
            type: String,
        },
    },
}
</script>

<style scoped lang="scss">
    #round-card {
        position: relative;
        margin: 0 auto;
        border-radius: 50%;
        text-align: center;
        overflow: hidden;
        img {
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
    }
</style>