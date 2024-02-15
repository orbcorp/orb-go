# Changelog

## 0.0.1 (2024-02-15)

Full Changelog: [...abc-v0.0.1](https://github.com/orbcorp/orb-go/compare/...abc-v0.0.1)

### Features

* **api:** add `external_customer_id` ([#69](https://github.com/orbcorp/orb-go/issues/69)) ([c6291e7](https://github.com/orbcorp/orb-go/commit/c6291e72513f02fc0c96896efb769301f431b385))
* **api:** add `version` to plan ([#75](https://github.com/orbcorp/orb-go/issues/75)) ([c34137d](https://github.com/orbcorp/orb-go/commit/c34137df0ca31411d51ebd772c7deca8cfb8e6c0))
* **api:** add beta evaluate price endpoint ([#54](https://github.com/orbcorp/orb-go/issues/54)) ([b9dfffe](https://github.com/orbcorp/orb-go/commit/b9dfffe57eaba2d7e7614704c7307f2e8399504a))
* **api:** add currency fields ([#48](https://github.com/orbcorp/orb-go/issues/48)) ([e9b9331](https://github.com/orbcorp/orb-go/commit/e9b933140182f1a6959f8dbb829831723b40b2ad))
* **api:** introduce per-price cost v2, credit top-ups ([#65](https://github.com/orbcorp/orb-go/issues/65)) ([06c0941](https://github.com/orbcorp/orb-go/commit/06c09411f71b3ef7922bfe9c18018d99d2860c02))
* **api:** price schema updates ([#70](https://github.com/orbcorp/orb-go/issues/70)) ([de8e0b2](https://github.com/orbcorp/orb-go/commit/de8e0b25713a4299b202d39311b45aa28ca689c0))
* **api:** remove unsupported field ([#38](https://github.com/orbcorp/orb-go/issues/38)) ([30f2e9c](https://github.com/orbcorp/orb-go/commit/30f2e9c93d7bbbe315cf599f9abc0c54c84bc311))
* **api:** remove unsupported params ([#21](https://github.com/orbcorp/orb-go/issues/21)) ([871d9dc](https://github.com/orbcorp/orb-go/commit/871d9dc16fbdd8d810b652d6a483f689d7c33ba1))
* **api:** updates ([#24](https://github.com/orbcorp/orb-go/issues/24)) ([a51a4a1](https://github.com/orbcorp/orb-go/commit/a51a4a166b3382ef9f23c836d4da1a579b653303))
* **api:** updates ([#28](https://github.com/orbcorp/orb-go/issues/28)) ([91641a7](https://github.com/orbcorp/orb-go/commit/91641a7980728fa16b27be17b994e497d7b83da2))
* **api:** updates ([#34](https://github.com/orbcorp/orb-go/issues/34)) ([0827662](https://github.com/orbcorp/orb-go/commit/08276624382966ef154ee19ca8d4a7525b2ea9a7))
* **api:** updates ([#39](https://github.com/orbcorp/orb-go/issues/39)) ([65af26c](https://github.com/orbcorp/orb-go/commit/65af26c46122a4960fe1ea1ee5095f32a175af0a))
* **api:** updates ([#57](https://github.com/orbcorp/orb-go/issues/57)) ([26324dc](https://github.com/orbcorp/orb-go/commit/26324dcca3d2b191246eb6bc10ec4bf6a5dcbf27))
* **api:** updates ([#58](https://github.com/orbcorp/orb-go/issues/58)) ([5babc2d](https://github.com/orbcorp/orb-go/commit/5babc2da21302b416861aeb00e2edb04a1853ad1))
* **api:** updates ([#6](https://github.com/orbcorp/orb-go/issues/6)) ([018f7b6](https://github.com/orbcorp/orb-go/commit/018f7b67a08e74d64f37c4a17f6ab23efc85f35f))
* **api:** updates ([#9](https://github.com/orbcorp/orb-go/issues/9)) ([7909ee5](https://github.com/orbcorp/orb-go/commit/7909ee53faad9187c9aee6a61209f562fb84ba44))
* **client:** adjust retry behavior ([#3](https://github.com/orbcorp/orb-go/issues/3)) ([689fc35](https://github.com/orbcorp/orb-go/commit/689fc35ca3fb7daf35570aaf8f2cf22aa282969c))
* **client:** allow binary returns ([#19](https://github.com/orbcorp/orb-go/issues/19)) ([b65919b](https://github.com/orbcorp/orb-go/commit/b65919b7a3a1f27302fa8badd9c6e869a3820293))
* **github:** include a devcontainer setup ([#18](https://github.com/orbcorp/orb-go/issues/18)) ([87c093b](https://github.com/orbcorp/orb-go/commit/87c093bcce0fb62effbc1c23cc832b9000e76db5))
* **init:** initial commit ([a6b37c9](https://github.com/orbcorp/orb-go/commit/a6b37c951e4d9607aad5f4b1bab7a3711dcc7805))
* **internal:** fallback to json serialization if no serialization methods are defined ([#41](https://github.com/orbcorp/orb-go/issues/41)) ([64281ba](https://github.com/orbcorp/orb-go/commit/64281bafb6199f9031296240255d93846b476e66))
* type alias enum values from shared in package root ([#16](https://github.com/orbcorp/orb-go/issues/16)) ([6a4efe0](https://github.com/orbcorp/orb-go/commit/6a4efe07b8ae1359432e310c03d40faaefb9daf6))


### Bug Fixes

* change status serialization behavior ([#80](https://github.com/orbcorp/orb-go/issues/80)) ([80b4475](https://github.com/orbcorp/orb-go/commit/80b447548e3e58a7a11878f9f9e61b963d99b593))
* **ci:** ignore stainless-app edits to release PR title ([#63](https://github.com/orbcorp/orb-go/issues/63)) ([881023f](https://github.com/orbcorp/orb-go/commit/881023ff945df58f365e5b733cb2c6d3cb3d747b))
* make options.WithHeader utils case-insensitive ([#22](https://github.com/orbcorp/orb-go/issues/22)) ([d7b32ac](https://github.com/orbcorp/orb-go/commit/d7b32ac3e3b2ffe9e5cb22ab8f69071e8c2aaefd))
* properly register discriminated unions ([#74](https://github.com/orbcorp/orb-go/issues/74)) ([9a856b6](https://github.com/orbcorp/orb-go/commit/9a856b6969f493dd883a5ac2530e040d8d03c552))
* rename customer.credits.ledger.create_entry_by_exteral_id and RequestValidationErrors ([#7](https://github.com/orbcorp/orb-go/issues/7)) ([5b1636b](https://github.com/orbcorp/orb-go/commit/5b1636ba1e476aa887d37a22a617e0d710afaabc))
* stop sending default idempotency headers with GET requests ([#31](https://github.com/orbcorp/orb-go/issues/31)) ([9f902fc](https://github.com/orbcorp/orb-go/commit/9f902fc2d552e8c27a0f03410105c5a386d72a52))
* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#62](https://github.com/orbcorp/orb-go/issues/62)) ([bf5735b](https://github.com/orbcorp/orb-go/commit/bf5735bcebfa90371a8a72f13a1a1a53ecc58758))
* use brackets instead of commas for array query params ([#47](https://github.com/orbcorp/orb-go/issues/47)) ([dda3342](https://github.com/orbcorp/orb-go/commit/dda334237db8a99bc3bf05b98d488a57770e8a3d))

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
