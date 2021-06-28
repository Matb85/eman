export interface PrimeRouteI {
  name: string;
  to: string;
}
class PrimeRoute {
  name: string;
  to: string;
  constructor(name: string) {
    this.name = name;
    this.to = "/" + name[0].toLowerCase() + name.substring(1);
  }
}

export const state = () => ({
  dashboards: [new PrimeRoute("Core"), new PrimeRoute("Media"), new PrimeRoute("Pages"), new PrimeRoute("Teams")],
});

export const mutations = {};

export const actions = {};
