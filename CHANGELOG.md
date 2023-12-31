# Changelog

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
