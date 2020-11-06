export function defaultlinks(dirname) {
    return {
        id: 0, name: "Narzędzia",
        sublinks: [
            {
                id: 0,
                name: "Główna",
                to: dirname,
                icon: "home"
            },
            {
                id: 1,
                name: "Kalendarz",
                to: dirname + "/calendar",
                icon: "chart-line-variant"
            },
            {
                id: 2,
                name: "Statystyki",
                to: dirname + "/statistics",
                icon: "calendar-range"
            },
        ]
    }
}