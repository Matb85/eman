<template>
    <section>
        <navbar :curPage="curPage" :links="links"/>
        
        <sidebar :links="links">
            <div v-for="branch in $store.state[curPage.toLowerCase()].subtreelinks" :key="branch.id">
                <hr class="my-2">
                <h3>{{branch.name}}</h3>
                <div class="routes">
                    <colorfulbtn v-for="sublink in branch.sublinks" :key="sublink.id" 
                    tag="nuxt-link" :to="sublink.to" :type="sublink.type"
                    expanded rounded outlined class="my-2" :addedclass="sublink.class"
                    :icon="sublink.icon" :customIcon="sublink.customIcon">{{sublink.name}}</colorfulbtn>
                </div>
            </div>
        </sidebar>
        <nuxt/>
    </section>
</template>

<script>
import navbar from "@/components/navbar.vue"
import sidebar from "@/components/sidebar.vue"
import colorfulbtn from "@/components/btnCornerIcon.vue"
export default {
    name: "dashboard",
    components: {
        navbar,
        sidebar,
        colorfulbtn
    },
    computed: {
        links(){
            return this.$store.state.dashboards.filter((link) => link.name != this.curPage)
        },
        curPage(){
            const path = this.$route.path
            const slash = (path.slice(2).includes("/"))? path.slice(2).indexOf("/")+2 : path.length
            return path.charAt(1).toUpperCase() + path.slice(2, slash)
        }
    },
}
</script>