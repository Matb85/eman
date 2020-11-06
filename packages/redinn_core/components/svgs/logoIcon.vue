<template functional>
    <svg id="logoicon" xmlns="http://www.w3.org/2000/svg" width="300" height="250" viewBox="0 0 79.4 66.1">
        <path :fill="props.fill" d="M0 66.1C23.1 26.7 55-1.5 66.1.1c14 0-27.7 42.4 13.3 66H0"/>
    </svg>
</template>

<script>
export default {
    name: "logoIcon",
    functional: true,
    props: {
        fill: {
            required: false,
            type: String,
            default: "#FF3030"
        }
    }
}
</script>