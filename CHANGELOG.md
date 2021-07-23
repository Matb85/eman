# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

# 0.1.0 (2021-07-23)


### Bug Fixes

* **api:** :rotating_light: ts issues ([964c970](https://github.com/Redinn-Labs/redinn-core/commit/964c9702784fccd3ee039cdc1c24908dc9e32cd9))
* **auth:** :bug: toke field ([6ccb7ee](https://github.com/Redinn-Labs/redinn-core/commit/6ccb7ee19c34848b7a5fed9a37f220355cbd3d9f))
* **graphql:** :bug: addenterprise Mutation ([688ecaf](https://github.com/Redinn-Labs/redinn-core/commit/688ecaf6ed44cfd417fb1a22aa42fafec713471e))
* **graphql:** :bug: deleteEnterprise errors ([52aaf66](https://github.com/Redinn-Labs/redinn-core/commit/52aaf666bcaa56113c629e3df5793967775d3fd0))
* **graphql:** :goal_net: error status codes ([afe486e](https://github.com/Redinn-Labs/redinn-core/commit/afe486efb5a2e7d06343b61c5855000fd43a399f))
* **media:** simplify routing in the photo uploader ([9495d3e](https://github.com/Redinn-Labs/redinn-core/commit/9495d3ec7152b4a63dad0e3f6375b7745f870db6))
* **website:** :bug: addEnterprise flow ([7b73917](https://github.com/Redinn-Labs/redinn-core/commit/7b73917f48496c41ed72965fbd36fc18587d3c1f))
* **website:** :bug: conditions for fetching enterprises ([55332d6](https://github.com/Redinn-Labs/redinn-core/commit/55332d61bd5f4301d3607226acc058204bbfd35c))
* **website:** :rotating_light: linters' warnings ([e8e2889](https://github.com/Redinn-Labs/redinn-core/commit/e8e288959aac9fc2f308403f5c0ca650bebdf9eb))
* **website:** :truck: addbusiness -> addenterprise ([bb807ec](https://github.com/Redinn-Labs/redinn-core/commit/bb807ecc212b0a098504d078b21faade5bc7d40a))
* :art: enterprise middleware ([e6e5daa](https://github.com/Redinn-Labs/redinn-core/commit/e6e5daad0e573fcf103e3a4a665de43903f3d003))
* impove performace - remove unnecessary css ([e7d7133](https://github.com/Redinn-Labs/redinn-core/commit/e7d7133a9923aef0be92b228e707a7211758e251))


### Features

* **api:** :building_construction: db package ([86606d3](https://github.com/Redinn-Labs/redinn-core/commit/86606d36f3ab757357fb1f767f6056e8b485d686))
* **api:** :goal_net: better error messages ([3437805](https://github.com/Redinn-Labs/redinn-core/commit/343780580e2e385d5e7b776d7622bc9e03af23ee))
* **api:** :lock: route for registering ([3962fb7](https://github.com/Redinn-Labs/redinn-core/commit/3962fb7431e8743c803558b922cf29b0df9ddf8f))
* **api:** :sparkles: graphql playground ([55a7715](https://github.com/Redinn-Labs/redinn-core/commit/55a7715fc792b28fd4bf42d3109a836b44ef76c3))
* **api:** :sparkles: types for users ([3a7b3fb](https://github.com/Redinn-Labs/redinn-core/commit/3a7b3fbf01ccb2d8deb1b96fb20c2a12aeb4d271))
* **auth:** :lock: /login route ([a9da29a](https://github.com/Redinn-Labs/redinn-core/commit/a9da29a94e0970694c20366b2a3f66d1150b263a))
* **auth:** :lock: bcrypt for passwords ([2a1cdb4](https://github.com/Redinn-Labs/redinn-core/commit/2a1cdb4fea48b78baf7ce367944fb07f3ddcbb3c))
* **auth:** :sparkles: add basic jwt + login authorization and user registration ([7682d31](https://github.com/Redinn-Labs/redinn-core/commit/7682d313c09d5d419c9404392b64c0e21e2105c4))
* **core:** :art: simplify the experience of adding an enterprise ([66f5b04](https://github.com/Redinn-Labs/redinn-core/commit/66f5b04630257e235cded30c5ac3dd31f806873f))
* **core:** :lock: add middleware that redirects a user to /addbusiness if they don't have any enterprises ([5694310](https://github.com/Redinn-Labs/redinn-core/commit/569431007ba4363bcc0b4ee05106605098fd30e1))
* **core:** :sparkles: enable users to add enterprises ([4d158bf](https://github.com/Redinn-Labs/redinn-core/commit/4d158bf06f25c0830608237a00e135010eb0822b))
* **graphql:** :sparkles: add auth to graphql and allow user's data fetching ([88bd05f](https://github.com/Redinn-Labs/redinn-core/commit/88bd05fd82f7e10d26966e84b1ff37c9e9d2e61d))
* **graphql:** :sparkles: add mutation - delete enterprise ([9dead37](https://github.com/Redinn-Labs/redinn-core/commit/9dead371d619b3fe308393991c649166badd4190))
* **graphql:** :sparkles: add types and model for the enterprise data type ([e8cd5b6](https://github.com/Redinn-Labs/redinn-core/commit/e8cd5b6d12238dde63d2dbfa471d3ddbe9d41daf))
* **graphql:** :sparkles: begin writing types and resolvers ([535ec03](https://github.com/Redinn-Labs/redinn-core/commit/535ec03f5e5d71904b3d564ec729dcfedde5cd70))
* **graphql:** :sparkles: enterprise resolvers ([5ec8f34](https://github.com/Redinn-Labs/redinn-core/commit/5ec8f34fefff12ea966aa48697ae57b54b53e297))
* **graphql:** :sparkles: fetch users from the db ([5471bf6](https://github.com/Redinn-Labs/redinn-core/commit/5471bf635e2394517cbcd8a04a728fc1400ce2fe))
* **graphql:** :zap: utils ([bf3d484](https://github.com/Redinn-Labs/redinn-core/commit/bf3d484c5689cdaabb1d3f2d3240a725690ac82b))
* :heavy_plus_sign: load .env variables thanks to dotenv ([55d3d16](https://github.com/Redinn-Labs/redinn-core/commit/55d3d164c948ae76365daef15374393b8fde4de5))
* :loud_sound: add morgan for better logs ([a562174](https://github.com/Redinn-Labs/redinn-core/commit/a562174129eb725f74d164cc730b920dbb7b6be2))
* :sparkles: add a db model for enterprises ([e55e3c9](https://github.com/Redinn-Labs/redinn-core/commit/e55e3c90350dabb3258ede1d8e13aae951b2633a))
* :sparkles: add authentication, authorization and registration ([d3cfdd4](https://github.com/Redinn-Labs/redinn-core/commit/d3cfdd4079120acd4b07895458f9cbac117bb485))
* :sparkles: add login functionality thanks to Nuxt Auth ([7ba46d6](https://github.com/Redinn-Labs/redinn-core/commit/7ba46d6b481730cabedcfee035073f5a2dfbed05))
* :wrench: add types for process.env variables ([c3cb99c](https://github.com/Redinn-Labs/redinn-core/commit/c3cb99c07f0adf0ec42c7b6a991d2f8fb62eabb9))
* add types for process.env variables ([1b4e660](https://github.com/Redinn-Labs/redinn-core/commit/1b4e660d3922f2c9b95da9b8fd6e416f7dd5429c))
* setup eslint and typescript ([1c65755](https://github.com/Redinn-Labs/redinn-core/commit/1c657559e89f22d9f0b18e618ed8a60b40ce1005))
* **media:** /addpost/upload: save images to sessionStorage in case of a reload right after uploading ([77dbf1d](https://github.com/Redinn-Labs/redinn-core/commit/77dbf1dadf2c4bf5971cb64965e7e22ef6bfd5bb))
* **ui:** :lipstick: redo the background layout ([a461933](https://github.com/Redinn-Labs/redinn-core/commit/a4619333dff18616dda5c43d39219d635c663266))


### Performance Improvements

* **api:** :zap: call less funcs ([e50c5e1](https://github.com/Redinn-Labs/redinn-core/commit/e50c5e1707866d3b80a27f891bda99087ff843e2))
* **api:** :zap: centralize access to db ([58822c9](https://github.com/Redinn-Labs/redinn-core/commit/58822c9407dc0538456fa6c5de4a2bbc0d84a14b))
* **api:** :zap: don't constantly convert users' ids ([36c36aa](https://github.com/Redinn-Labs/redinn-core/commit/36c36aa6522d0f3b528141bf1e1a5ef8c9de7560))
* **api:** :zap: setup timeouts ([7302c39](https://github.com/Redinn-Labs/redinn-core/commit/7302c3988a4dc5b039242933e27a31507bf7fe42))
* **graphql:** :zap: use _id instead of uuid ([fed6e1c](https://github.com/Redinn-Labs/redinn-core/commit/fed6e1ccc527a9763b79c5f44400dc6cedcde1e5))
* **website:** :art: use async/await ([46709d6](https://github.com/Redinn-Labs/redinn-core/commit/46709d60d648aed724a1d619237e9308ab00991e))
