const dirname = "/media"
import { defaultlinks } from "~/store/globals/globals"
export const state = () => ({
    subtreelinks: [
        defaultlinks(dirname),
        {
            id: 1, name: "Media",
            sublinks: [
                {
                    id: 2,
                    name: "Google Business",
                    to: dirname + "/statistics",
                    customIcon: "gbusiness",
                    type: "is-gbusiness",
                    class: "backlit"
                },
                {
                    id: 1,
                    name: "Facebook",
                    to: dirname + "/calendar",
                    customIcon: "facebook",
                    type: "is-facebook",
                },
                {
                    id: 0,
                    name: "Instagram",
                    to: dirname,
                    customIcon: "instagram",
                    type: "is-instagram",
                    class: "instagram whiteicon"
                },
                {
                    id: 3,
                    name: "Pinterest",
                    to: dirname + "/statistics",
                    customIcon: "pinterest",
                    type: "is-pinterest",
                },
                {
                    id: 4,
                    name: "Twitter",
                    to: dirname + "/statistics",
                    customIcon: "twitter",
                    type: "is-twitter",
                },
            ]
        }
    ],
})

export const mutations = {
}

export const actions = {
}