export class SubRoute {
  name: string;
  icon: string;
  to: string;
  constructor(name: string, icon: string, to: string) {
    this.name = name;
    this.icon = icon;
    this.to = to;
  }
}

export function defaultlinks(dirname: string) {
  return {
    name: "Narzędzia",
    sublinks: [
      new SubRoute("Główna", "home", dirname),
      new SubRoute("Kalendarz", "calendar-range", dirname + "/calendar"),
      new SubRoute("Statystyki", "chart-line-variant", dirname + "/statistics"),
    ],
  };
}
