# Changelog

## 0.24.0 (2024-03-26)

Full Changelog: [v0.23.1...v0.24.0](https://github.com/orbcorp/orb-go/compare/v0.23.1...v0.24.0)

### Features

* add IsKnown method to enums ([#111](https://github.com/orbcorp/orb-go/issues/111)) ([50d1623](https://github.com/orbcorp/orb-go/commit/50d1623f43a1f87b7b84b0100d1da5bd0db6cb52))
* **api:** introduce credits status ([658ff1e](https://github.com/orbcorp/orb-go/commit/658ff1ea9605f7d96ff3dc15cb20d6f5afb362b8))
* **api:** remove `scaling_factor` ([#114](https://github.com/orbcorp/orb-go/issues/114)) ([658ff1e](https://github.com/orbcorp/orb-go/commit/658ff1ea9605f7d96ff3dc15cb20d6f5afb362b8))


### Chores

* **internal:** update generated pragma comment ([#110](https://github.com/orbcorp/orb-go/issues/110)) ([713043c](https://github.com/orbcorp/orb-go/commit/713043c8db17cccf3fc68ff8584ac4d9c0159464))


### Documentation

* fix typo in CONTRIBUTING.md ([#108](https://github.com/orbcorp/orb-go/issues/108)) ([7a107b3](https://github.com/orbcorp/orb-go/commit/7a107b30d9bce7ee47077afcb22025151bc2fd84))
* **readme:** consistent use of sentence case in headings ([#113](https://github.com/orbcorp/orb-go/issues/113)) ([bef493a](https://github.com/orbcorp/orb-go/commit/bef493a87cdac0083fc213a814284e7eade7ea52))
* **readme:** document file uploads ([#115](https://github.com/orbcorp/orb-go/issues/115)) ([c5efd2e](https://github.com/orbcorp/orb-go/commit/c5efd2e33eb0eb3d7955699d389610e9efc2482a))
* updated invoice description ([#112](https://github.com/orbcorp/orb-go/issues/112)) ([7d55ac1](https://github.com/orbcorp/orb-go/commit/7d55ac1377a18e7d1513108573c084d3796999ef))

## 0.23.1 (2024-03-18)

Full Changelog: [v0.23.0...v0.23.1](https://github.com/orbcorp/orb-go/compare/v0.23.0...v0.23.1)

### Bug Fixes

* **api:** improve union handling for request bodies ([#106](https://github.com/orbcorp/orb-go/issues/106)) ([93a5ed2](https://github.com/orbcorp/orb-go/commit/93a5ed25222c0f49a5d716f6b69f62c4ce210008))

## 0.23.0 (2024-03-15)

Full Changelog: [v0.22.0...v0.23.0](https://github.com/orbcorp/orb-go/compare/v0.22.0...v0.23.0)

### Features

* **api:** create invoice metadata param ([#104](https://github.com/orbcorp/orb-go/issues/104)) ([c585633](https://github.com/orbcorp/orb-go/commit/c58563397fc91a2e03b75dbb93b11e15393a9007))

## 0.22.0 (2024-03-14)

Full Changelog: [v0.21.2...v0.22.0](https://github.com/orbcorp/orb-go/compare/v0.21.2...v0.22.0)

### Features

* **api:** add matrix with allocation price ([#103](https://github.com/orbcorp/orb-go/issues/103)) ([e819aad](https://github.com/orbcorp/orb-go/commit/e819aad807bd126774e7b3011ffbd5b87de31283))
* set user-agent header by default when making requests ([#101](https://github.com/orbcorp/orb-go/issues/101)) ([cd92451](https://github.com/orbcorp/orb-go/commit/cd92451f97743eb75297e6f246f6d4222842fbfc))

## 0.21.2 (2024-03-12)

Full Changelog: [v0.21.1...v0.21.2](https://github.com/orbcorp/orb-go/compare/v0.21.1...v0.21.2)

### Bug Fixes

* **client:** don't include ? in path unless necessary ([#99](https://github.com/orbcorp/orb-go/issues/99)) ([1449fb4](https://github.com/orbcorp/orb-go/commit/1449fb4be288821ebba2c4e41efb774f312041a8))

## 0.21.1 (2024-03-12)

Full Changelog: [v0.21.0...v0.21.1](https://github.com/orbcorp/orb-go/compare/v0.21.0...v0.21.1)

### Bug Fixes

* fix String() behavior of param.Field ([#97](https://github.com/orbcorp/orb-go/issues/97)) ([3f3e715](https://github.com/orbcorp/orb-go/commit/3f3e71588321bec0707f1a33ab68d93a7db57f92))

## 0.21.0 (2024-03-11)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/orbcorp/orb-go/compare/v0.20.0...v0.21.0)

### Features

* **api:** updates ([#95](https://github.com/orbcorp/orb-go/issues/95)) ([d7d4253](https://github.com/orbcorp/orb-go/commit/d7d42534a7eab74012d07121c556da659a49d18c))

## 0.20.0 (2024-03-08)

Full Changelog: [v0.19.1...v0.20.0](https://github.com/orbcorp/orb-go/compare/v0.19.1...v0.20.0)

### Features

* **api:** updates ([#94](https://github.com/orbcorp/orb-go/issues/94)) ([81eafc8](https://github.com/orbcorp/orb-go/commit/81eafc8004ceb7ce84ee7c25a983cff5c82b6f29))


### Documentation

* **contributing:** add a CONTRIBUTING.md ([#92](https://github.com/orbcorp/orb-go/issues/92)) ([7240de0](https://github.com/orbcorp/orb-go/commit/7240de0827f83cd5960dc6bf2cf2bfe50d74a309))

## 0.19.1 (2024-03-07)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/orbcorp/orb-go/compare/v0.19.0...v0.19.1)

### Bug Fixes

* fix Price and Ledger union deserialization ([#90](https://github.com/orbcorp/orb-go/issues/90)) ([384fd44](https://github.com/orbcorp/orb-go/commit/384fd44a8138a776092918e17bee2462d592d2b8))

## 0.19.0 (2024-03-07)

Full Changelog: [v0.18.1...v0.19.0](https://github.com/orbcorp/orb-go/compare/v0.18.1...v0.19.0)

### Features

* implement public RawJSON of response structs ([#88](https://github.com/orbcorp/orb-go/issues/88)) ([05bfb4b](https://github.com/orbcorp/orb-go/commit/05bfb4b860dd4fc79578f7f65ba6e1118fd4ad45))


### Bug Fixes

* fix union deserialization for multiple objects ([#89](https://github.com/orbcorp/orb-go/issues/89)) ([896ac1e](https://github.com/orbcorp/orb-go/commit/896ac1e70ef476c9750eb727cb0b83e3bb2528ad))


### Chores

* **ci:** uses Stainless GitHub App for releases ([#83](https://github.com/orbcorp/orb-go/issues/83)) ([c0fb3d0](https://github.com/orbcorp/orb-go/commit/c0fb3d0a21294b287e76c81b7a591d597f443e57))
* **internal:** bump timeout threshold in tests ([#81](https://github.com/orbcorp/orb-go/issues/81)) ([69709ef](https://github.com/orbcorp/orb-go/commit/69709efb1df345372e301f7da269322d64493053))
* **internal:** refactor release environment script ([#84](https://github.com/orbcorp/orb-go/issues/84)) ([edbfe98](https://github.com/orbcorp/orb-go/commit/edbfe98064b97f8038b63ea56654b5fa496619ac))
* **internal:** update deps ([#86](https://github.com/orbcorp/orb-go/issues/86)) ([c746c73](https://github.com/orbcorp/orb-go/commit/c746c73574700de853ecfb8a0257c0ee073d1aa3))


### Documentation

* **readme:** improve wording ([#87](https://github.com/orbcorp/orb-go/issues/87)) ([83f9d09](https://github.com/orbcorp/orb-go/commit/83f9d09e1cc0d693aa1b10377ff3b3fb5eb460be))

## 0.18.1 (2024-02-07)

Full Changelog: [v0.18.0...v0.18.1](https://github.com/orbcorp/orb-go/compare/v0.18.0...v0.18.1)

### Bug Fixes

* change status serialization behavior ([#80](https://github.com/orbcorp/orb-go/issues/80)) ([3d52b74](https://github.com/orbcorp/orb-go/commit/3d52b741f0b97bdac1f438954d5d1aa1ec6499f2))


### Chores

* **interal:** make link to api.md relative ([#77](https://github.com/orbcorp/orb-go/issues/77)) ([14c8dad](https://github.com/orbcorp/orb-go/commit/14c8dad60679a5dabc7b3f63f1da9d06c127d57f))
* **internal:** adjust formatting ([#79](https://github.com/orbcorp/orb-go/issues/79)) ([1abb53e](https://github.com/orbcorp/orb-go/commit/1abb53e52f6c480b4161cc775f65842077130ef9))

## 0.18.0 (2024-02-01)

Full Changelog: [v0.17.1...v0.18.0](https://github.com/orbcorp/orb-go/compare/v0.17.1...v0.18.0)

### Features

* **api:** add `version` to plan ([#75](https://github.com/orbcorp/orb-go/issues/75)) ([d3ff4de](https://github.com/orbcorp/orb-go/commit/d3ff4de20191cb973137fbffffa85b872f652644))

## 0.17.1 (2024-01-31)

Full Changelog: [v0.17.0...v0.17.1](https://github.com/orbcorp/orb-go/compare/v0.17.0...v0.17.1)

### Bug Fixes

* properly register discriminated unions ([#74](https://github.com/orbcorp/orb-go/issues/74)) ([adaf56d](https://github.com/orbcorp/orb-go/commit/adaf56dcc7dc9a986ad2b7c79a7e034b5902ee9d))


### Chores

* **internal:** support pre-release versioning ([#72](https://github.com/orbcorp/orb-go/issues/72)) ([7d2a140](https://github.com/orbcorp/orb-go/commit/7d2a140b0bc8ee0cb1bd71e255a0b3653e391975))

## 0.17.0 (2024-01-30)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/orbcorp/orb-go/compare/v0.16.0...v0.17.0)

### Features

* **api:** price schema updates ([#70](https://github.com/orbcorp/orb-go/issues/70)) ([faaf9bc](https://github.com/orbcorp/orb-go/commit/faaf9bc254f2a12103b15a7632f885d48af31d5b))

## 0.16.0 (2024-01-30)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/orbcorp/orb-go/compare/v0.15.0...v0.16.0)

### Features

* **api:** add `external_customer_id` ([#69](https://github.com/orbcorp/orb-go/issues/69)) ([f220ce4](https://github.com/orbcorp/orb-go/commit/f220ce44b0ecab7da0bd286a053fe05e11ef0173))


### Chores

* **internal:** parse date-time strings more leniently ([#67](https://github.com/orbcorp/orb-go/issues/67)) ([41e0d63](https://github.com/orbcorp/orb-go/commit/41e0d63fe5cda972c2db39b2c6ccd91d32fc1e66))

## 0.15.0 (2024-01-22)

Full Changelog: [v0.14.2...v0.15.0](https://github.com/orbcorp/orb-go/compare/v0.14.2...v0.15.0)

### Features

* **api:** introduce per-price cost v2, credit top-ups ([#65](https://github.com/orbcorp/orb-go/issues/65)) ([718e39d](https://github.com/orbcorp/orb-go/commit/718e39d321a574257a4ed660a25e469d5759ce04))

## 0.14.2 (2024-01-18)

Full Changelog: [v0.14.1...v0.14.2](https://github.com/orbcorp/orb-go/compare/v0.14.1...v0.14.2)

### Bug Fixes

* **ci:** ignore stainless-app edits to release PR title ([#63](https://github.com/orbcorp/orb-go/issues/63)) ([a0c58ce](https://github.com/orbcorp/orb-go/commit/a0c58ceee640711d431b9a82c649417fc384d1af))

## 0.14.1 (2024-01-17)

Full Changelog: [v0.14.0...v0.14.1](https://github.com/orbcorp/orb-go/compare/v0.14.0...v0.14.1)

### Bug Fixes

* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#62](https://github.com/orbcorp/orb-go/issues/62)) ([191baab](https://github.com/orbcorp/orb-go/commit/191baab152ab720eebd5778a2ac8efa0e4f8cfbf))


### Chores

* **internal:** speculative retry-after-ms support ([#60](https://github.com/orbcorp/orb-go/issues/60)) ([c1d6958](https://github.com/orbcorp/orb-go/commit/c1d6958ed48e1cd1bbd1c766361693cd518a4855))

## 0.14.0 (2024-01-17)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/orbcorp/orb-go/compare/v0.13.0...v0.14.0)

### Features

* **api:** updates ([#58](https://github.com/orbcorp/orb-go/issues/58)) ([65f323b](https://github.com/orbcorp/orb-go/commit/65f323b78966225dd18334fd466419cb7b34e96c))

## 0.13.0 (2024-01-15)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/orbcorp/orb-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** updates ([#57](https://github.com/orbcorp/orb-go/issues/57)) ([f86171d](https://github.com/orbcorp/orb-go/commit/f86171d1ed077c619168194d4569933074584080))


### Chores

* formatting ([#55](https://github.com/orbcorp/orb-go/issues/55)) ([f46706f](https://github.com/orbcorp/orb-go/commit/f46706fb3f0f24949a2af6cf13a7dce5abdef5f9))

## 0.12.0 (2024-01-12)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/orbcorp/orb-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** add beta evaluate price endpoint ([#54](https://github.com/orbcorp/orb-go/issues/54)) ([725a916](https://github.com/orbcorp/orb-go/commit/725a9160f6ae76bb1ef3ed5912f969dd47aa1ee7))


### Chores

* add .keep files for examples and custom code directories ([#52](https://github.com/orbcorp/orb-go/issues/52)) ([cead4f2](https://github.com/orbcorp/orb-go/commit/cead4f2d8656de487ab6cae31434ca114ebf1547))
* **internal:** minor updates to pagination ([#50](https://github.com/orbcorp/orb-go/issues/50)) ([9df26bf](https://github.com/orbcorp/orb-go/commit/9df26bf86604a530c74de6c4363ff01bf33fdf55))


### Documentation

* **readme:** improve api reference ([#53](https://github.com/orbcorp/orb-go/issues/53)) ([8f49f03](https://github.com/orbcorp/orb-go/commit/8f49f03d846a0b724baea013f8d415fac1affc54))

## 0.11.0 (2024-01-01)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/orbcorp/orb-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** add currency fields ([#48](https://github.com/orbcorp/orb-go/issues/48)) ([f621f46](https://github.com/orbcorp/orb-go/commit/f621f461983cd31158044bde23c5328412afec9b))

## 0.10.0 (2023-12-26)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/orbcorp/orb-go/compare/v0.9.0...v0.10.0)

### Features

* **internal:** fallback to json serialization if no serialization methods are defined ([#41](https://github.com/orbcorp/orb-go/issues/41)) ([2d282c0](https://github.com/orbcorp/orb-go/commit/2d282c0e4550c0b51f575b99c852a9634df5176b))


### Bug Fixes

* use brackets instead of commas for array query params ([#47](https://github.com/orbcorp/orb-go/issues/47)) ([9aafd58](https://github.com/orbcorp/orb-go/commit/9aafd58ae040bf36fae103fc8b3f7efe4178e491))


### Chores

* **ci:** run release workflow once per day ([#44](https://github.com/orbcorp/orb-go/issues/44)) ([61c34a3](https://github.com/orbcorp/orb-go/commit/61c34a3ae53d1123ea6230c23505f28b610df046))


### Documentation

* **api:** updates ([#46](https://github.com/orbcorp/orb-go/issues/46)) ([2ef33b3](https://github.com/orbcorp/orb-go/commit/2ef33b31c98515320387680b13b48dc45e6de0df))
* avoid normalizing trailing space ([#43](https://github.com/orbcorp/orb-go/issues/43)) ([cd88196](https://github.com/orbcorp/orb-go/commit/cd8819666864839993edbcf58265effe587e4aba))
* **options:** fix link to readme ([#45](https://github.com/orbcorp/orb-go/issues/45)) ([af3e4d5](https://github.com/orbcorp/orb-go/commit/af3e4d53252338a35e30cbb84365186ae146c5cd))

## 0.9.0 (2023-12-11)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/orbcorp/orb-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** updates ([#39](https://github.com/orbcorp/orb-go/issues/39)) ([6d7952e](https://github.com/orbcorp/orb-go/commit/6d7952ef944dced3a3b6e7cfdb01629e538d0af0))

## 0.8.0 (2023-12-08)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/orbcorp/orb-go/compare/v0.7.0...v0.8.0)

### Features

* **api:** remove unsupported field ([#38](https://github.com/orbcorp/orb-go/issues/38)) ([f724da3](https://github.com/orbcorp/orb-go/commit/f724da3326ff0f51342f85449920a252e1afc10c))

## 0.7.0 (2023-11-28)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/orbcorp/orb-go/compare/v0.6.1...v0.7.0)

### Features

* **api:** updates ([#34](https://github.com/orbcorp/orb-go/issues/34)) ([1a6c9ce](https://github.com/orbcorp/orb-go/commit/1a6c9ce28f28e9270852136896de687d6b1b263c))


### Documentation

* **api:** update metadata docstrings ([#32](https://github.com/orbcorp/orb-go/issues/32)) ([56bfc55](https://github.com/orbcorp/orb-go/commit/56bfc5548b3144f4ed7bd97b6f5af7797e8d2bb5))

## 0.6.1 (2023-11-17)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/orbcorp/orb-go/compare/v0.6.0...v0.6.1)

### Bug Fixes

* stop sending default idempotency headers with GET requests ([#31](https://github.com/orbcorp/orb-go/issues/31)) ([ef9e67d](https://github.com/orbcorp/orb-go/commit/ef9e67dcfb2ae27fd7f5c8f2fe21ccf445c3614c))


### Documentation

* **readme:** fix link to docs website ([#29](https://github.com/orbcorp/orb-go/issues/29)) ([b9b2a06](https://github.com/orbcorp/orb-go/commit/b9b2a06209f119de878316acce52bd49a5db6a7a))

## 0.6.0 (2023-11-16)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/orbcorp/orb-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** updates ([#28](https://github.com/orbcorp/orb-go/issues/28)) ([1dce8d6](https://github.com/orbcorp/orb-go/commit/1dce8d6b3d7830bf33339844f78edff8f3c84ef4))


### Refactors

* do not include `JSON` fields when serialising back to json ([#26](https://github.com/orbcorp/orb-go/issues/26)) ([b6f991c](https://github.com/orbcorp/orb-go/commit/b6f991c4c492e60e1c8132cca7f92b2d456ac24d))

## 0.5.0 (2023-11-09)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/orbcorp/orb-go/compare/v0.4.1...v0.5.0)

### Features

* **api:** updates ([#24](https://github.com/orbcorp/orb-go/issues/24)) ([7bded2d](https://github.com/orbcorp/orb-go/commit/7bded2ddd30b100a44032f936a59f70e57cc2a60))

## 0.4.1 (2023-11-08)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/orbcorp/orb-go/compare/v0.4.0...v0.4.1)

### Bug Fixes

* make options.WithHeader utils case-insensitive ([#22](https://github.com/orbcorp/orb-go/issues/22)) ([cf9e4db](https://github.com/orbcorp/orb-go/commit/cf9e4db374fc4f8555a541a24e69598d8f69c1f5))

## 0.4.0 (2023-11-06)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/orbcorp/orb-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** remove unsupported params ([#21](https://github.com/orbcorp/orb-go/issues/21)) ([ae0cb79](https://github.com/orbcorp/orb-go/commit/ae0cb792d7f9c7484714d22d977ac858f5303e1e))
* **client:** allow binary returns ([#19](https://github.com/orbcorp/orb-go/issues/19)) ([c70fbf8](https://github.com/orbcorp/orb-go/commit/c70fbf8ffef8cdc3d2a53113710c8cee7091be6e))
* **github:** include a devcontainer setup ([#18](https://github.com/orbcorp/orb-go/issues/18)) ([c2efbd4](https://github.com/orbcorp/orb-go/commit/c2efbd43c0a1a5e67c404375bdfd4f00eabe47ff))
* type alias enum values from shared in package root ([#16](https://github.com/orbcorp/orb-go/issues/16)) ([8f4a8ae](https://github.com/orbcorp/orb-go/commit/8f4a8aeb3ffc9b342c6797475ea6af2d27b5db29))


### Documentation

* **readme:** improve example snippets ([#20](https://github.com/orbcorp/orb-go/issues/20)) ([33a83b3](https://github.com/orbcorp/orb-go/commit/33a83b310946d5863777967ae01f3adea37bd263))

## 0.3.0 (2023-10-27)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/orbcorp/orb-go/compare/v0.2.1...v0.3.0)

### Features

* **api:** updates ([#9](https://github.com/orbcorp/orb-go/issues/9)) ([3f9d760](https://github.com/orbcorp/orb-go/commit/3f9d760586553746b0cde168579c0d744cf7207f))

## 0.2.1 (2023-10-26)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/orbcorp/orb-go/compare/v0.2.0...v0.2.1)

### Bug Fixes

* rename customer.credits.ledger.create_entry_by_exteral_id and RequestValidationErrors ([#7](https://github.com/orbcorp/orb-go/issues/7)) ([778f476](https://github.com/orbcorp/orb-go/commit/778f4766e6a730ead8840915e71c16a30a8f347d))

## 0.2.0 (2023-10-26)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/orbcorp/orb-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** updates ([#6](https://github.com/orbcorp/orb-go/issues/6)) ([371e13a](https://github.com/orbcorp/orb-go/commit/371e13a303d1735e4160a6f72a50d4d826c3855f))
* **client:** adjust retry behavior ([#3](https://github.com/orbcorp/orb-go/issues/3)) ([bf7bb50](https://github.com/orbcorp/orb-go/commit/bf7bb5031fcc57f7b0137849ff9d428c0e254094))

## 0.1.0 (2023-10-23)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/orbcorp/orb-go/compare/v0.0.1...v0.1.0)

### Features

* **init:** initial commit ([a6b37c9](https://github.com/orbcorp/orb-go/commit/a6b37c951e4d9607aad5f4b1bab7a3711dcc7805))


### Documentation

* improve code examples ([cbdd9fe](https://github.com/orbcorp/orb-go/commit/cbdd9fe679634618845ae23aa1bc1e52481b8226))
