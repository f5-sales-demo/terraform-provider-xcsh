---
page_title: "xcsh_site_registrations_by_state reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state reference."
---

# xcsh_site_registrations_by_state reference

<a id="canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4cfb55aace60111a033a80daf77d9c9b3707fd1b956d620b3fbcb0cf5a67af0"></a>

## Property reference — Property reference / 872e9837982c / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- Property reference

<a id="canonical-a8c65eafccf74a99b4cee3cb28f2be5ffce49771032a73254048e9047cfbbb96"></a>

## Direct properties — Property reference / 872e9837982c / 3

- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-db84e5ed548d2312394a53b8856bf8998548c9d7b05e61eff5d11d998c7a9c32): complete subsection reference.

- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c): complete subsection reference.

<a id="canonical-750495cf0721392d1e9cb9e63d7b7b68e548727a0b3b88b8ada35b09fea3352f"></a>

<a id="canonical-e65d4f1c72ec9693e9d66c4dd75dd2b17ed5e2e27562ca735d0249abc6fd0bdc"></a>

## namespace property — Property reference / 872e9837982c / 4

Type: `"string"`. Optional, Computed.

Namespace. Registration namespace, only 'system' namespaces is accepted.

<a id="canonical-ddac6fd41c18336558c312ac40b1b715b84013562dde2837db7db02700f1c7f4"></a>

<a id="canonical-f8fbab27b3f4f1531462ccbdf967eb75c3d5a28c9b5149574f3419718b459c80"></a>

## state property — Property reference / 872e9837982c / 5

Type: `"string"`. Required.

\[Enum:
NOTSET|NEW|APPROVED|ADMITTED|RETIRED|FAILED|DONE|PENDING|ONLINE|UPGRADING|MAINTENANCE|FAILED\_INACTIVE\]
Defines states for registration object State isn't set Object was created (registration request was
received and object created) Registration was approved and waiting for configuration This state can
be set by user only if current state is NEW Registration is approved and prepared for to connect..
Possible values are \`NOTSET\`, \`NEW\`, \`APPROVED\`, \`ADMITTED\`, \`RETIRED\`, \`FAILED\`,
\`DONE\`, \`PENDING\`, \`ONLINE\`, \`UPGRADING\`, \`MAINTENANCE\`, \`FAILED\_INACTIVE\`. Defaults to
\`NOTSET\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOTSET",
    "NEW",
    "APPROVED",
    "ADMITTED",
    "RETIRED",
    "FAILED",
    "DONE",
    "PENDING",
    "ONLINE",
    "UPGRADING",
    "MAINTENANCE",
    "FAILED_INACTIVE"),
}
```

<a id="canonical-94b7bef909c0595e906e00d955faabcd2764c82a55f9f2e98d289d28d3e8adbc"></a>

## All schema paths — Property reference / 872e9837982c / 6

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `errors` | [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2f4ce87b1f910e3b1bdc4af0aa25520b1351013f3b73d718c8c1979cbe8295d2) |
| `errors.code` | [errors.code](data-sources--site_registrations_by_state--reference--group-001.md#canonical-25cffc968edf0e5c6f0b53519e31fbb326d1c27b3aa9e937a5d0f7194278c23a) |
| `errors.error_obj` | [errors.error_obj](data-sources--site_registrations_by_state--reference--group-001.md#canonical-8ad1e349cd34b26917cbc303d0ae7200a1d8055a09e084faced72212f9c70628) |
| `errors.error_obj.type_url` | [errors.error_obj.type_url](data-sources--site_registrations_by_state--reference--group-001.md#canonical-c440f0f45f1499eabd9314f07fb9c6ef6021ecbfa979fafc02a587334ed34fcd) |
| `errors.error_obj.value` | [errors.error_obj.value](data-sources--site_registrations_by_state--reference--group-001.md#canonical-402810b003af9d69f8d8945d8d2aed114c93840b0faa48cfedd689fe68284f5f) |
| `errors.message` | [errors.message](data-sources--site_registrations_by_state--reference--group-001.md#canonical-12b8e0447c6d3e4608b302c4b5f1cd71e987d26611efdde749d79f02a63302a4) |
| `items` | [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-bb2003288b68aef55834c33c75355b55d2922d3755d6069fd39fe2760f789433) |
| `items.annotations` | [items.annotations](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3e446def30581c52f75bf8f2d975c02bae8a79469e4549841ab93f58a8788150) |
| `items.description_spec` | [items.description_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-ef283be3ca3f9d677b227709099a53c8bd10af33e2bae343f82a9520da2ac40c) |
| `items.disabled` | [items.disabled](data-sources--site_registrations_by_state--reference--group-001.md#canonical-13ebd799ef106f8b5d55ac012555df9f855be724d0a428ba8a93b15ed0e8f128) |
| `items.get_spec` | [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-a61da75adf7ac99a816ad36c9393b778ec33f841c1b4c18c99e5399f020978fc) |
| `items.get_spec.infra` | [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-42a751645d5b1765384861c067aede853806a1fceb50f96e7a5ae70862afff39) |
| `items.get_spec.infra.availability_zone` | [items.get_spec.infra.availability_zone](data-sources--site_registrations_by_state--reference--group-001.md#canonical-9405634bc199506c9fd5fd800b7973fbcd2218e11b851bd33c02e73e20956c2b) |
| `items.get_spec.infra.bond_config` | [items.get_spec.infra.bond_config](data-sources--site_registrations_by_state--reference--group-001.md#canonical-f30725962dbd23246d5b09a4a032416b7d6988dd8a9079f548bfe240bfec9d76) |
| `items.get_spec.infra.bond_config.interfaces` | [items.get_spec.infra.bond_config.interfaces](data-sources--site_registrations_by_state--reference--group-001.md#canonical-e8f2143956811cc438de9a1ef3ea157800194de5fcb0bb717279b8bb8511c9c2) |
| `items.get_spec.infra.bond_config.mode` | [items.get_spec.infra.bond_config.mode](data-sources--site_registrations_by_state--reference--group-001.md#canonical-75557d04091c39f89e77c19bb33e9a7d2d7435622d0eed41c3a0a42b9223ddfb) |
| `items.get_spec.infra.bond_config.name` | [items.get_spec.infra.bond_config.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5a6f78e72c3afed598f290908b1131049c7e5f58e4d1ff720c91763bd1d07589) |
| `items.get_spec.infra.certified_hw` | [items.get_spec.infra.certified_hw](data-sources--site_registrations_by_state--reference--group-001.md#canonical-15f40b2271d1d7c70f7b40b00e7c2357c12e0f32004a95aad92268053aa9e3b3) |
| `items.get_spec.infra.domain` | [items.get_spec.infra.domain](data-sources--site_registrations_by_state--reference--group-001.md#canonical-a575c7f4d29b9e6eadd244c4b4d543a744e4dabd7b2dfa55a0fcc980c02faa2e) |
| `items.get_spec.infra.hostname` | [items.get_spec.infra.hostname](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b8ebd44166694cb36ef234c4be574f5a80a445ae23b392e27f5d95e1ffec6916) |
| `items.get_spec.infra.hugepages` | [items.get_spec.infra.hugepages](data-sources--site_registrations_by_state--reference--group-001.md#canonical-ce6098f47d18fb32ec33e1c30b482eabd01f3a37941a6faec45748647e64b8c5) |
| `items.get_spec.infra.hugepages.free` | [items.get_spec.infra.hugepages.free](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b138783d1787366d910a8109af32bc99f4415614b0092940dbc02c165816c392) |
| `items.get_spec.infra.hugepages.page_size` | [items.get_spec.infra.hugepages.page_size](data-sources--site_registrations_by_state--reference--group-001.md#canonical-d4d05b80e7c05e0142997acd98387d5330e1af4df80d78158230a4fc202e6a9c) |
| `items.get_spec.infra.hugepages.total` | [items.get_spec.infra.hugepages.total](data-sources--site_registrations_by_state--reference--group-001.md#canonical-94d43b30e3d4b084b30d797cc69ddffc29391f0f9e31f03313b4f2656561eef6) |
| `items.get_spec.infra.hw_info` | [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3fd098b622a08d9f02f9e3d59e287a50e2eb7bc56cb3db4bf9497ae88227a79d) |
| `items.get_spec.infra.hw_info.bios` | [items.get_spec.infra.hw_info.bios](data-sources--site_registrations_by_state--reference--group-001.md#canonical-13734ea8ce5c2cbb0230b603c8bc70060c04c17df19073495587ba6a2247c237) |
| `items.get_spec.infra.hw_info.bios.date` | [items.get_spec.infra.hw_info.bios.date](data-sources--site_registrations_by_state--reference--group-001.md#canonical-d35f0f8b9a9ff031c0fe582a7a9bef0a0d8d48a795ea9f59b305726487cc8538) |
| `items.get_spec.infra.hw_info.bios.vendor` | [items.get_spec.infra.hw_info.bios.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-623ac9db8b8bcaa1d0d532b3025a5c6e8256a0d48bc254ae339ffe1a81f79643) |
| `items.get_spec.infra.hw_info.bios.version` | [items.get_spec.infra.hw_info.bios.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-7d858c9a05e14595c74b797d84895a1640d56d5b64cda01a26e144c7880a47fd) |
| `items.get_spec.infra.hw_info.board` | [items.get_spec.infra.hw_info.board](data-sources--site_registrations_by_state--reference--group-001.md#canonical-27117624ce637c4df9a1628e66ef47671d3c21dc9c48775c004de214cb2ca48a) |
| `items.get_spec.infra.hw_info.board.asset_tag` | [items.get_spec.infra.hw_info.board.asset_tag](data-sources--site_registrations_by_state--reference--group-001.md#canonical-957a973c5694473be9a5f0d10d816e6670a594724fccbaeb2df06f45a7151979) |
| `items.get_spec.infra.hw_info.board.name` | [items.get_spec.infra.hw_info.board.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0ec53a24327708fe432c648bd25cb45df5b0af22021e7578e865348886966f5d) |
| `items.get_spec.infra.hw_info.board.serial` | [items.get_spec.infra.hw_info.board.serial](data-sources--site_registrations_by_state--reference--group-001.md#canonical-661d631c465969b9d2e39ba3b807bcce177e743a7fa8cd6e0c08bc6636b99e3d) |
| `items.get_spec.infra.hw_info.board.vendor` | [items.get_spec.infra.hw_info.board.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-17649f0014c4247c6774322c37f61cde3dedc0042683acfd3996c7b8ba61d48c) |
| `items.get_spec.infra.hw_info.board.version` | [items.get_spec.infra.hw_info.board.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-68eed800d841eb3c9f0d436607ea47ff4b48cc170261fc0bad40ed60097556e8) |
| `items.get_spec.infra.hw_info.chassis` | [items.get_spec.infra.hw_info.chassis](data-sources--site_registrations_by_state--reference--group-001.md#canonical-fd42e849b0045a38641c92ebc8d912df4e9ecaa0cda20209afb63594e753c9e7) |
| `items.get_spec.infra.hw_info.chassis.asset_tag` | [items.get_spec.infra.hw_info.chassis.asset_tag](data-sources--site_registrations_by_state--reference--group-001.md#canonical-a4c079208d1d38ddf68e892d2ff8c0a1b8d06ca5995f9e9b250ef2bef2ac98be) |
| `items.get_spec.infra.hw_info.chassis.serial` | [items.get_spec.infra.hw_info.chassis.serial](data-sources--site_registrations_by_state--reference--group-001.md#canonical-9e69383a7ac2c41a2bfb586a1fa5aa5c710a9775b7d37cc62793b6b70e20ee79) |
| `items.get_spec.infra.hw_info.chassis.type` | [items.get_spec.infra.hw_info.chassis.type](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b9c1648834cadaf7d5ce8b42874a2f8f5dbafb1c525702c63401e90c709078da) |
| `items.get_spec.infra.hw_info.chassis.vendor` | [items.get_spec.infra.hw_info.chassis.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-6963c6848971c3ddfb2fef3a42de4c7d2aace387214ef2a2fafdb73a479359c3) |
| `items.get_spec.infra.hw_info.chassis.version` | [items.get_spec.infra.hw_info.chassis.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-fc7c050ba8066ebdf70f43f66d01de24156dc84e256d21664d000a3dc72504f0) |
| `items.get_spec.infra.hw_info.cpu` | [items.get_spec.infra.hw_info.cpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-34e658547379e5e865eb1993352d2c0b319e5883dd7532dd67f352e0e16f5655) |
| `items.get_spec.infra.hw_info.cpu.cache` | [items.get_spec.infra.hw_info.cpu.cache](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2c2abad45c51def69818fd7739f14d3b5eccd8b50655f51fbe86326d92f5afe9) |
| `items.get_spec.infra.hw_info.cpu.cores` | [items.get_spec.infra.hw_info.cpu.cores](data-sources--site_registrations_by_state--reference--group-001.md#canonical-a78e7ce6460cc8271d96f3a30171b2531bcfdea9731a8fd421f5c7b373f1c53f) |
| `items.get_spec.infra.hw_info.cpu.cpus` | [items.get_spec.infra.hw_info.cpu.cpus](data-sources--site_registrations_by_state--reference--group-001.md#canonical-8d393c29024ce3eaec8819d20e0a63c456711dee8d8c043b54ecf6fa6b225d51) |
| `items.get_spec.infra.hw_info.cpu.model` | [items.get_spec.infra.hw_info.cpu.model](data-sources--site_registrations_by_state--reference--group-001.md#canonical-efeea466dd2fb8b52248c4fb740fc1ffec8fe0203135caf24177a7282f338885) |
| `items.get_spec.infra.hw_info.cpu.speed` | [items.get_spec.infra.hw_info.cpu.speed](data-sources--site_registrations_by_state--reference--group-001.md#canonical-51f31c03954bab760e927ea99c447b9bcb563a6cb1dd1de429ae2963e7675c8e) |
| `items.get_spec.infra.hw_info.cpu.threads` | [items.get_spec.infra.hw_info.cpu.threads](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b2af4d2923aca499e636d8e672516cfa6908e2048c4e23b6b309e06dc69a67cf) |
| `items.get_spec.infra.hw_info.cpu.vendor` | [items.get_spec.infra.hw_info.cpu.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-4ce8263e927569b55ca788b0145f3464dfa7d6898ec30b17dc6c7cfeeac00bbc) |
| `items.get_spec.infra.hw_info.gpu` | [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-4095f0e56f14dc36d885814341619b4579a0bb3388abb3e8107f9e196fbf6e32) |
| `items.get_spec.infra.hw_info.gpu.cuda_version` | [items.get_spec.infra.hw_info.gpu.cuda_version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5062039477ec16c3ffc6e14a166f654bf342c15cf170b84aced92c612954e1fe) |
| `items.get_spec.infra.hw_info.gpu.driver_version` | [items.get_spec.infra.hw_info.gpu.driver_version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5452371ed3a84568faa3b9229906b292d10d3b4ef3696cbcf55f62dfeb3e3fc8) |
| `items.get_spec.infra.hw_info.gpu.gpu_device` | [items.get_spec.infra.hw_info.gpu.gpu_device](data-sources--site_registrations_by_state--reference--group-001.md#canonical-dc91d62c6aa9d613ee5d448dc697beb6afb7ae5aa5fc6fdea141c622bad86ed1) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.id` | [items.get_spec.infra.hw_info.gpu.gpu_device.id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-c669513f3d7a8c377a4fa4c5b03a691d449a954d024b9be90a1d8c3058a908a9) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.processes` | [items.get_spec.infra.hw_info.gpu.gpu_device.processes](data-sources--site_registrations_by_state--reference--group-001.md#canonical-6ec7f61d5fd25b082caf72f4b7096fa3d844a97d312eb7652bee8e24a39847e8) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.product_name` | [items.get_spec.infra.hw_info.gpu.gpu_device.product_name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-12948f64048fe94385220a0c2f0c6b9b785cd153938641ae3c5398d813268ca9) |
| `items.get_spec.infra.hw_info.kernel` | [items.get_spec.infra.hw_info.kernel](data-sources--site_registrations_by_state--reference--group-001.md#canonical-ca92e3dfcab064e0d7f76a0f5d3bd3157c6fc41c6d31ab08609ffa790c82f622) |
| `items.get_spec.infra.hw_info.kernel.architecture` | [items.get_spec.infra.hw_info.kernel.architecture](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b1f97d1e99ddf98b117a8481af9e8e9bd48ef41cdd7041b57eda06c04970ba8f) |
| `items.get_spec.infra.hw_info.kernel.release` | [items.get_spec.infra.hw_info.kernel.release](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6a9087b1e470b56171ed489ad0dd6902fc1028cec9edb6ba7c2ca50d5360b720) |
| `items.get_spec.infra.hw_info.kernel.version` | [items.get_spec.infra.hw_info.kernel.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-abf966d2b2ed54df93f70c483bcc1d0a73103779b4019815624bffe8135071be) |
| `items.get_spec.infra.hw_info.memory` | [items.get_spec.infra.hw_info.memory](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1c58a8477d5ce7a103f0e4e3d948ab97f9f6bb6737f654636c996f562a8119eb) |
| `items.get_spec.infra.hw_info.memory.size_mb` | [items.get_spec.infra.hw_info.memory.size_mb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4f74549da77f2011f6071a0d1f5d431826cc54f90028ce8601bfe4628ed41415) |
| `items.get_spec.infra.hw_info.memory.speed` | [items.get_spec.infra.hw_info.memory.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9c526d39e8c345d5bd2ccb81ef4f8e95b5eb191491958cdbe78e885c218d5ea3) |
| `items.get_spec.infra.hw_info.memory.type` | [items.get_spec.infra.hw_info.memory.type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b642a5a6ecfcfdfce00a8af15882967f52191eff1a0e43c3647d66535466d1bb) |
| `items.get_spec.infra.hw_info.network` | [items.get_spec.infra.hw_info.network](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d2b4954f53630491a609afab87a6647695abb2e07b15fc71780de72e27f51a6b) |
| `items.get_spec.infra.hw_info.network.driver` | [items.get_spec.infra.hw_info.network.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7a24a07b8b96faaf97bab327fab4c37a6580511a4df8a0406980f1bb7953b29b) |
| `items.get_spec.infra.hw_info.network.ip_address` | [items.get_spec.infra.hw_info.network.ip_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-80042fc14f0bccffafa12683b1970255f515311217eeee7ab80f2e3a555a2cad) |
| `items.get_spec.infra.hw_info.network.link_quality` | [items.get_spec.infra.hw_info.network.link_quality](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b2f4b22c130b543cbd8c00d1981d3c8770b4e358cb162e4de81e9f7ba35d2852) |
| `items.get_spec.infra.hw_info.network.link_type` | [items.get_spec.infra.hw_info.network.link_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b1fb1be21cdbb7c3c1248d8325be62ca5d5133486be6d524c4d977810e04ad79) |
| `items.get_spec.infra.hw_info.network.mac_address` | [items.get_spec.infra.hw_info.network.mac_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dc38e8febc23742eb611c979fa74e7cb895c5e8748df33dfe032ce935c8c77d6) |
| `items.get_spec.infra.hw_info.network.name` | [items.get_spec.infra.hw_info.network.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-5447fa600b877821b2c06344303215724532e2b2685bbb68b07c2bbb147ac714) |
| `items.get_spec.infra.hw_info.network.port` | [items.get_spec.infra.hw_info.network.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9413a4355af6e0bd3afa61192a34302f0bbcce986f76d432a1482c9517577467) |
| `items.get_spec.infra.hw_info.network.speed` | [items.get_spec.infra.hw_info.network.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b5555cfd640ddbf7ae390b49ced4e1625b185425f6c88b2fc16d3a9e74232ee3) |
| `items.get_spec.infra.hw_info.numa_nodes` | [items.get_spec.infra.hw_info.numa_nodes](data-sources--site_registrations_by_state--reference--group-001.md#canonical-93a846b791c2f70fc85ceb809a4d7a609a7eb9b895061954745e623cca8d6c5b) |
| `items.get_spec.infra.hw_info.os` | [items.get_spec.infra.hw_info.os](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dc300487be75745220f6df55d6718951f8ce2a36c2e2d19a87d472aeeca41ced) |
| `items.get_spec.infra.hw_info.os.architecture` | [items.get_spec.infra.hw_info.os.architecture](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a5b295b2c40237e2d98be0af33dd1ec03f48a451a3ccc31b1834450a8a015d36) |
| `items.get_spec.infra.hw_info.os.name` | [items.get_spec.infra.hw_info.os.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-54b43b2ca6dc2ce27718c7cb38d7f7bfc44a56d95e12a1b2205421c3c324af2d) |
| `items.get_spec.infra.hw_info.os.release` | [items.get_spec.infra.hw_info.os.release](data-sources--site_registrations_by_state--reference--group-002.md#canonical-424d71db9347a344fc952e2e20c60a988c97e0ef04db229b0a9e970ed4ecf792) |
| `items.get_spec.infra.hw_info.os.vendor` | [items.get_spec.infra.hw_info.os.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2bd272da7078cca90023fbc06042a94afa0c692a40fdaf13efbe4a0b420bc47d) |
| `items.get_spec.infra.hw_info.os.version` | [items.get_spec.infra.hw_info.os.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-bada75cd8f4e48737f93fe0483679ed29569d60f19e2f6aba1018f045a1f831c) |
| `items.get_spec.infra.hw_info.product` | [items.get_spec.infra.hw_info.product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6bddb8d7f8546e84b21c056ce52a7f9ebf6cf6da86229d374927bba269509760) |
| `items.get_spec.infra.hw_info.product.name` | [items.get_spec.infra.hw_info.product.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-365d1b3948c11015d4faea8c96d6ee0573998f12512dabfae1b6258b70fe74a7) |
| `items.get_spec.infra.hw_info.product.serial` | [items.get_spec.infra.hw_info.product.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dcf1ba73e5203187467be3377a8effe13a4cb715c1a98f017897eda060c99988) |
| `items.get_spec.infra.hw_info.product.vendor` | [items.get_spec.infra.hw_info.product.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9978156a6a6494f5ef20a4bebfb7d427265ae33a6a07cf3c6315ad52c3e3d24e) |
| `items.get_spec.infra.hw_info.product.version` | [items.get_spec.infra.hw_info.product.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-489ea85bfaed952b1a1328e1f7775c50db13ddbc22117770afd76db6040488b4) |
| `items.get_spec.infra.hw_info.storage` | [items.get_spec.infra.hw_info.storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-180d50a76e1279f10adc37d612341d8f5f2302d4a484c150028d8001a2165032) |
| `items.get_spec.infra.hw_info.storage.driver` | [items.get_spec.infra.hw_info.storage.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d3b236cbd423fa927e1c4b7d2738842eab91f877af3020dffec468f34c8234b4) |
| `items.get_spec.infra.hw_info.storage.model` | [items.get_spec.infra.hw_info.storage.model](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4ebd22ff41298f7c5d0943aab4ff6017f510c484d64260de109cd186cf78990b) |
| `items.get_spec.infra.hw_info.storage.name` | [items.get_spec.infra.hw_info.storage.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1b64a9c1da0bcc05af94d640b125f75b5f1fe4c14805960f133db7a0eb519b51) |
| `items.get_spec.infra.hw_info.storage.serial` | [items.get_spec.infra.hw_info.storage.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-96e5706819a1487ade0995880a74c04bd24ad08066fb20dc3ba470b3becf0142) |
| `items.get_spec.infra.hw_info.storage.size_gb` | [items.get_spec.infra.hw_info.storage.size_gb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-12b2255f7adc3bc95fea43b4a2b9795b5b67d87be6e7f2b15b6a5937f8ac8179) |
| `items.get_spec.infra.hw_info.storage.vendor` | [items.get_spec.infra.hw_info.storage.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7dcf481505399224842c98dd0cab0c992155673f3b5a98918c637ef385a71968) |
| `items.get_spec.infra.hw_info.usb` | [items.get_spec.infra.hw_info.usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-803065e9c344a46ce355e7f692c1d98b3529498aa2c4f220c2a9fa5988c78b28) |
| `items.get_spec.infra.hw_info.usb.address` | [items.get_spec.infra.hw_info.usb.address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-eddd7475b3951fa2ad2caa1a21ea3d31a6959a17e063a2f1e42e2a23a0a4467a) |
| `items.get_spec.infra.hw_info.usb.b_device_class` | [items.get_spec.infra.hw_info.usb.b_device_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-002112cc67c750e7160626d83bb1f7d45974b25df9094cdfad09dc27a8365dfd) |
| `items.get_spec.infra.hw_info.usb.b_device_protocol` | [items.get_spec.infra.hw_info.usb.b_device_protocol](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c79c1085dccd213903f67df3dbb48c7e69bb260cfd19c7660342032c7071e178) |
| `items.get_spec.infra.hw_info.usb.b_device_sub_class` | [items.get_spec.infra.hw_info.usb.b_device_sub_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-71b8d97b0fe63fcef0fb40539afea4675adbdc598d5ca22bbeedf8a17164b691) |
| `items.get_spec.infra.hw_info.usb.b_max_packet_size` | [items.get_spec.infra.hw_info.usb.b_max_packet_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-bf8b2e1c06f7a14fc67626ff510b2604d468bcac0544118701b3e23d47b51c9d) |
| `items.get_spec.infra.hw_info.usb.bcd_device` | [items.get_spec.infra.hw_info.usb.bcd_device](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6f85c1c138d2fc7d2dd6ee334b61fdbc196e182662bf2134310004f463a271db) |
| `items.get_spec.infra.hw_info.usb.bcd_usb` | [items.get_spec.infra.hw_info.usb.bcd_usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-723e017ba2b3251a0fcf83066d37ded5b05da06fa5e83dc2ac4e4e24134c8053) |
| `items.get_spec.infra.hw_info.usb.bus` | [items.get_spec.infra.hw_info.usb.bus](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fe9fc08cfc2e546aaf8e9b05a7b1459467f0ca5ee7fa2361f4fab075320b6d7a) |
| `items.get_spec.infra.hw_info.usb.description_spec` | [items.get_spec.infra.hw_info.usb.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a70631daa6cc7f2702cfb0090bb71cfbfecd137b2374a5c61ac01e938eb2ed34) |
| `items.get_spec.infra.hw_info.usb.i_manufacturer` | [items.get_spec.infra.hw_info.usb.i_manufacturer](data-sources--site_registrations_by_state--reference--group-002.md#canonical-71f1f479ce18a6b31a71b3c3ad01bb7ad55d21162fe26f4736210b62a4573773) |
| `items.get_spec.infra.hw_info.usb.i_product` | [items.get_spec.infra.hw_info.usb.i_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f3e73e80ae4bab49629ab65e3f59455a34f4d945b52661f3775d08e7f50f60db) |
| `items.get_spec.infra.hw_info.usb.i_serial` | [items.get_spec.infra.hw_info.usb.i_serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f4cc353898e82c45f5a1a234d45c939d8792e8c00275f78b619a64484b8272e7) |
| `items.get_spec.infra.hw_info.usb.id_product` | [items.get_spec.infra.hw_info.usb.id_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cd761b6bfbd677fc50d48da20c75bac6fea06d31edb1750226bdc2bd7fad099a) |
| `items.get_spec.infra.hw_info.usb.id_vendor` | [items.get_spec.infra.hw_info.usb.id_vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f0612527ea8776898ecdd3fe84b35f79ce342bc82282ee0a37291893c5aeb6cb) |
| `items.get_spec.infra.hw_info.usb.port` | [items.get_spec.infra.hw_info.usb.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-64d03dcf5cf9308eec2beb793557e93f64a921f742129939764ee3f139fde87a) |
| `items.get_spec.infra.hw_info.usb.product_name` | [items.get_spec.infra.hw_info.usb.product_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9dc37cd2bb9a6ce28fe936bef4ff97b9b0393524e623e10abc970253f20978e3) |
| `items.get_spec.infra.hw_info.usb.speed` | [items.get_spec.infra.hw_info.usb.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b73496fe946fa82adbbac122a505245afaed8810bc8e4d3e4eba05fd0947d2bb) |
| `items.get_spec.infra.hw_info.usb.usb_type` | [items.get_spec.infra.hw_info.usb.usb_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2a555ea1bcd363159c91d384b0fdd98df7d27c6d9f80c62e1691b7a67d2772de) |
| `items.get_spec.infra.hw_info.usb.vendor_name` | [items.get_spec.infra.hw_info.usb.vendor_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-8d5863ac31f73c4219f7824d00d9fe4946fdbcb74b55135641730f6f2d442a5e) |
| `items.get_spec.infra.instance_id` | [items.get_spec.infra.instance_id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-501a411ec702f7b16858c2092bae5d7c3d3fdd935332cda6ef8c6e158cacae57) |
| `items.get_spec.infra.interfaces` | [items.get_spec.infra.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-89e19b760d093e82a645acad70f3a347a33b756c042d66728731dec65a13dccd) |
| `items.get_spec.infra.internet_proxy` | [items.get_spec.infra.internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2c84c10b714447b965468642e97ff460eea8ad6c7f5e2ade28f10f37755f1071) |
| `items.get_spec.infra.internet_proxy.http_proxy` | [items.get_spec.infra.internet_proxy.http_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a0b6de1d474c1ba61016ea87ef7db4abff34e1eab24489ab18f0d65406e822e4) |
| `items.get_spec.infra.internet_proxy.https_proxy` | [items.get_spec.infra.internet_proxy.https_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c211c4ff7cf5bc77776daf8fdceeeb990aed01c35fa8908e26634b1a3bf73c52) |
| `items.get_spec.infra.internet_proxy.no_proxy` | [items.get_spec.infra.internet_proxy.no_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f4a3aad6d68eb3cdd656b49e9228ae4e66d774e2b1ee105cef94f17845994278) |
| `items.get_spec.infra.internet_proxy.proxy_cacert_url` | [items.get_spec.infra.internet_proxy.proxy_cacert_url](data-sources--site_registrations_by_state--reference--group-002.md#canonical-adc5c8459dfb3c6f64016be838324b56dd4142112dd6a76b2eee18843282bfbc) |
| `items.get_spec.infra.is_slo_static` | [items.get_spec.infra.is_slo_static](data-sources--site_registrations_by_state--reference--group-001.md#canonical-29d9fc8f36ebc016a6adfe29fdb751815dd170578abb7b477557eb616bf6b711) |
| `items.get_spec.infra.machine_id` | [items.get_spec.infra.machine_id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b6445b17c6141b8ea28fa855f590e0d9fe0083d536567229439927ba7bbc70c3) |
| `items.get_spec.infra.provider_ref` | [items.get_spec.infra.provider_ref](data-sources--site_registrations_by_state--reference--group-001.md#canonical-252b57387e7011ea07310cf26acd99bfcaefb914612027fb014f2964eea5f49b) |
| `items.get_spec.infra.sw_info` | [items.get_spec.infra.sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7ec3b58e8f84ebdc0b494c470709b0a45095ef78e766dc727974c6fad5ec144e) |
| `items.get_spec.infra.sw_info.sw_version` | [items.get_spec.infra.sw_info.sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-593118d7940d313e90f1048f20aa5cc54b64c8c627f4ae05d1dab62db79bbbed) |
| `items.get_spec.infra.timestamp` | [items.get_spec.infra.timestamp](data-sources--site_registrations_by_state--reference--group-001.md#canonical-ef5227baacac553c88d9159d730322df315b141840756055c181edab824448a9) |
| `items.get_spec.infra.zone` | [items.get_spec.infra.zone](data-sources--site_registrations_by_state--reference--group-001.md#canonical-da623860865c7e6f4a07ca2ce5b1279d24ae97cd45bbb83739f0c5b0bcc39ccf) |
| `items.get_spec.passport` | [items.get_spec.passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c5924c211cef50fd6469340329a5f6179d8f7c4f2aaf8a026b0c4889160a772b) |
| `items.get_spec.passport.cluster_name` | [items.get_spec.passport.cluster_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f0c03a9ca6710de6f6e3f2c0283c1ef15da760cff9ebccca9f5fbc1a9bd1ad42) |
| `items.get_spec.passport.cluster_size` | [items.get_spec.passport.cluster_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fbf72cb9a28802034cb0a82298c3011b7ffdb5c53680023d4e7f7c1ec1e074b0) |
| `items.get_spec.passport.cluster_type` | [items.get_spec.passport.cluster_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-8af861b76f34a4e238bc3ed0f296f3d7bf62e93e68256b850a486a2d8aa6229f) |
| `items.get_spec.passport.default_os_version` | [items.get_spec.passport.default_os_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-82f077d86c1663185c9e415da2c3d943c957a1d405710cf07f97adb99065b72f) |
| `items.get_spec.passport.default_sw_version` | [items.get_spec.passport.default_sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a15a6d97f4fabcde54b25898686844c0b336c12a10f31f2cb8d487292de777cc) |
| `items.get_spec.passport.latitude` | [items.get_spec.passport.latitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-aeca2997257d6710eee8a46c4a4d4398d724da127df9e2fbd4cdccd4180ffe78) |
| `items.get_spec.passport.longitude` | [items.get_spec.passport.longitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fca69386a0c04b85b3ba8e83bd8374fcf5e7b27cbc6749e6eaefe150bae7b36d) |
| `items.get_spec.passport.operating_system_version` | [items.get_spec.passport.operating_system_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3c0a4c4306d95556dd3bbe56936a0b777dc6e00dcad29cf873425a97d5b17e12) |
| `items.get_spec.passport.private_network_name` | [items.get_spec.passport.private_network_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-55b0e68c6c795908479707aeaffe482123fb1eb4aa4f5106831122e6222f508e) |
| `items.get_spec.passport.volterra_software_version` | [items.get_spec.passport.volterra_software_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-40844c905be56f6d1e7d16d97b7451d5d3dcef9a2ff4b1ac41c35f725297413d) |
| `items.get_spec.token` | [items.get_spec.token](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5eb4049c30e6b27338d6f25fbe573152300f6edde7e97963d4241e8c61f80c50) |
| `items.labels` | [items.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9a64c33635a30ba148397aaafc742347abc47ef4fcccd0c9803e90a319a2346f) |
| `items.metadata` | [items.metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b95decf1feb18727bcec0ceb56c3698584bfac5387fa3a9b9e55cfd40c314eb) |
| `items.metadata.annotations` | [items.metadata.annotations](data-sources--site_registrations_by_state--reference--group-002.md#canonical-00d61b93eefbfb0fb116a9e56b653cfc30a7036269b330ae023d42897897f890) |
| `items.metadata.description_spec` | [items.metadata.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7e2c7137a873de4df57c0575fbdea6e1562ff4107a4a22b93524f1071458ae27) |
| `items.metadata.disable_spec` | [items.metadata.disable_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c2d6d2d565a62b1fd3bc41f5b081917ce8188ed73c22b0f6d59521a15d1a3dd9) |
| `items.metadata.labels` | [items.metadata.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0bfd0479441d301c13e440baefff79ae53761d143d1e1a2ae9cc8d529302029a) |
| `items.metadata.name` | [items.metadata.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f0afd7059da9061f957695564d3ad9ef4a855af3743fe03007960ec2d6ecbed1) |
| `items.metadata.namespace` | [items.metadata.namespace](data-sources--site_registrations_by_state--reference--group-002.md#canonical-164a284480b4ca449d5041d7d32765da78811a3892b0083029c11171cf7d3d91) |
| `items.name` | [items.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-490cb78e455ed2160dfde4b44ee2185d671292f1365c73a824de4018a3e3cd80) |
| `items.namespace` | [items.namespace](data-sources--site_registrations_by_state--reference--group-001.md#canonical-e900a43a6367722347c452a25bfb9f5c2ce8bcf522862da062b708ff270df72f) |
| `items.object` | [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e2301c31b8562cf7b7282c046393da837910dcf65479cae21ef07c65057a1ec6) |
| `items.object.metadata` | [items.object.metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-972a0ea2dc580a226c4128e98d44379de510aa817de3ba36010b4426ae5a1805) |
| `items.object.metadata.annotations` | [items.object.metadata.annotations](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6d1ec2c05f9cf25c136d7769ac4b7307692744335f22464ddba0a7cd56b45987) |
| `items.object.metadata.description_spec` | [items.object.metadata.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e0cfe8e5d7b7dbd6be4ad64349e29ec3508edba031d842c64b94c1f541990c0b) |
| `items.object.metadata.disable_spec` | [items.object.metadata.disable_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b35f889f490eb03d470d2248d29418e65ce159805349f65219d42f6485ddfbc4) |
| `items.object.metadata.labels` | [items.object.metadata.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ab0b59af197470d53a955da93f970ec5cd1523f705e413861660db79fb078617) |
| `items.object.metadata.name` | [items.object.metadata.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a14b001fec32c17c83153704fc0ae9b2e379353305a0f1676dcfcac07a079300) |
| `items.object.metadata.namespace` | [items.object.metadata.namespace](data-sources--site_registrations_by_state--reference--group-002.md#canonical-56f2f6a4beec69ec5ff5de9643a5e8712389d2e93c47e234091998e6f83d803c) |
| `items.object.metadata.uid` | [items.object.metadata.uid](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dc4a5f7c1fa1baba4378668c57bb4cdf24ec134fba6e9fd56b28f6c9b49131a8) |
| `items.object.spec` | [items.object.spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-be9d3a1c00a523010c92505ae75971e88c44c649a448efe56db87b1ca72352a0) |
| `items.object.spec.gc_spec` | [items.object.spec.gc_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c7114150662455325350e23911d86379d5654ed08c078e0b83e19ee8eed81283) |
| `items.object.spec.gc_spec.connected_regions` | [items.object.spec.gc_spec.connected_regions](data-sources--site_registrations_by_state--reference--group-002.md#canonical-5420a38271f93f74913a94c55bce7faa04bc703490692024c560fc844fcd7609) |
| `items.object.spec.gc_spec.infra` | [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3a1b6886e38fd9a4b649ae47136d0b00d0b582ea4d7cce6351f46a911a58362b) |
| `items.object.spec.gc_spec.infra.availability_zone` | [items.object.spec.gc_spec.infra.availability_zone](data-sources--site_registrations_by_state--reference--group-002.md#canonical-56a4d0b0fc7c37dcac07f4ea70af823d52c367959ea0b7a576d30ec4a65d83b2) |
| `items.object.spec.gc_spec.infra.bond_config` | [items.object.spec.gc_spec.infra.bond_config](data-sources--site_registrations_by_state--reference--group-002.md#canonical-94ee0406a804db9a9bf237f9a74b81a34035cffa1d513947c2392e645e5d0477) |
| `items.object.spec.gc_spec.infra.bond_config.interfaces` | [items.object.spec.gc_spec.infra.bond_config.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-09ae3fc7359a216fbd3ffdb353983cbb3f0a7efd6b42d57e4352fc987afe4304) |
| `items.object.spec.gc_spec.infra.bond_config.mode` | [items.object.spec.gc_spec.infra.bond_config.mode](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3b9a3525b8cd3e3bc857c4a225572cab3d7acebbc63bbfc909807150e7e50411) |
| `items.object.spec.gc_spec.infra.bond_config.name` | [items.object.spec.gc_spec.infra.bond_config.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a5fad6ea42e639b86dab66785aeaeafe3d224e0c028d793c22843699582de6ba) |
| `items.object.spec.gc_spec.infra.certified_hw` | [items.object.spec.gc_spec.infra.certified_hw](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c2b137812401f4f22cd11872f30b10ae82de8bceda2a1c0c0cbda50726d9f2b7) |
| `items.object.spec.gc_spec.infra.domain` | [items.object.spec.gc_spec.infra.domain](data-sources--site_registrations_by_state--reference--group-002.md#canonical-72332c48639090736a8d7cf4bf7b59543395b99de4859aa64b437872f54e5a0d) |
| `items.object.spec.gc_spec.infra.hostname` | [items.object.spec.gc_spec.infra.hostname](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d04e1944aff955919144e3c5446c2a3bda6c12d1dfd7abea4eb28299389ec572) |
| `items.object.spec.gc_spec.infra.hugepages` | [items.object.spec.gc_spec.infra.hugepages](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0fc59e5551e1a2bd862cbd8dc6faddb800dbafe5c357e6099492ec2cb1819d8e) |
| `items.object.spec.gc_spec.infra.hugepages.free` | [items.object.spec.gc_spec.infra.hugepages.free](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f0e6024186a7aae867118dc0a3c54e5d5e98ca6e5f5a1d58d38027c27210fed5) |
| `items.object.spec.gc_spec.infra.hugepages.page_size` | [items.object.spec.gc_spec.infra.hugepages.page_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7b4e91580166514f95747c0a34574178aba1bf379055b32867c9d164b8b821e8) |
| `items.object.spec.gc_spec.infra.hugepages.total` | [items.object.spec.gc_spec.infra.hugepages.total](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2b1095e542d166ce26aaab3f437a6b2923b22233f1a2453a5c58461cfedf16ac) |
| `items.object.spec.gc_spec.infra.hw_info` | [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b8022b9b769025fdec47994ea391c5b3fd0bac5043c43f0402aa8d560884f802) |
| `items.object.spec.gc_spec.infra.hw_info.bios` | [items.object.spec.gc_spec.infra.hw_info.bios](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7dddd5381cfba7bdd34b1abb82b9bc5203891591a0a3b9d6d06c3a8d97e9127a) |
| `items.object.spec.gc_spec.infra.hw_info.bios.date` | [items.object.spec.gc_spec.infra.hw_info.bios.date](data-sources--site_registrations_by_state--reference--group-002.md#canonical-5b40c7f0dd66adebaded344af38f4a2ab28953fb197913b12f03708b2aa1d493) |
| `items.object.spec.gc_spec.infra.hw_info.bios.vendor` | [items.object.spec.gc_spec.infra.hw_info.bios.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-775f30259b1b9ee627b819b83d8f083e8b2b4e2358d46ed18a630b40eb9b93a1) |
| `items.object.spec.gc_spec.infra.hw_info.bios.version` | [items.object.spec.gc_spec.infra.hw_info.bios.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0eeba65338d0551f1dc1948be06d7cc83ed24a3034b4cea6e8ff46b6cdf1b539) |
| `items.object.spec.gc_spec.infra.hw_info.board` | [items.object.spec.gc_spec.infra.hw_info.board](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1917d3b7bb2c3157ab4b2be35e1ef8a6023051a6f1d5b545027013fef3b920a4) |
| `items.object.spec.gc_spec.infra.hw_info.board.asset_tag` | [items.object.spec.gc_spec.infra.hw_info.board.asset_tag](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7fba29474234783b0d95904bd06878fc80b6404e01b9bb8c74d829c095007e2f) |
| `items.object.spec.gc_spec.infra.hw_info.board.name` | [items.object.spec.gc_spec.infra.hw_info.board.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fe99aeacd84be403458e801cf75b15670217bbca04c5bd977607ee54a97001a7) |
| `items.object.spec.gc_spec.infra.hw_info.board.serial` | [items.object.spec.gc_spec.infra.hw_info.board.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-332882ca81474256dc6df5f39a5df1b942dd55157f833819a1dac4df072f8fdc) |
| `items.object.spec.gc_spec.infra.hw_info.board.vendor` | [items.object.spec.gc_spec.infra.hw_info.board.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b3d31c9fa87e467b91b7e6e324167abcc136a86ef3349e02747945a8665152d1) |
| `items.object.spec.gc_spec.infra.hw_info.board.version` | [items.object.spec.gc_spec.infra.hw_info.board.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d856ba3f1de6849ad9300c88837a9d8ea09b223d38139c674494e621ec54c948) |
| `items.object.spec.gc_spec.infra.hw_info.chassis` | [items.object.spec.gc_spec.infra.hw_info.chassis](data-sources--site_registrations_by_state--reference--group-002.md#canonical-55b0ed6d9825302489bd5e264ceb957ca8f2922e81826731601fdc5b4e5863e8) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.asset_tag` | [items.object.spec.gc_spec.infra.hw_info.chassis.asset_tag](data-sources--site_registrations_by_state--reference--group-002.md#canonical-08f2aa66fabeab9fbb6b7675e3a2d66555a87886018e00943de68bb629ad78ad) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.serial` | [items.object.spec.gc_spec.infra.hw_info.chassis.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b23e94ba1c5f897dc794d072451cec3f655d0f50366f6f1b7090ad7a4962a970) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.type` | [items.object.spec.gc_spec.infra.hw_info.chassis.type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cf3412b1181dbe716f7f3526b19ef5c3fbd7c8c82f8510c3e77061d7d4325e74) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.vendor` | [items.object.spec.gc_spec.infra.hw_info.chassis.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fa9462d958a11ed79109726f42ec40a7e3f04b2fd054e834dbc9b8c96f8deab9) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.version` | [items.object.spec.gc_spec.infra.hw_info.chassis.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ad614427ada14124b718419a990104d031929837e877d3cee9bf9ae2f0fcdaef) |
| `items.object.spec.gc_spec.infra.hw_info.cpu` | [items.object.spec.gc_spec.infra.hw_info.cpu](data-sources--site_registrations_by_state--reference--group-002.md#canonical-705a93da8bc69d3bc4d31bbd47121f0c8b73449bebc4d698f3e24b44f7f6da14) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cache` | [items.object.spec.gc_spec.infra.hw_info.cpu.cache](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2dd31781aba3b9b00af0f6d692769eef16d54310c3a49a3f468e36dd30ff5a6a) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cores` | [items.object.spec.gc_spec.infra.hw_info.cpu.cores](data-sources--site_registrations_by_state--reference--group-002.md#canonical-03703ca8d12a7f2e3041b6abe6cb1a9b02c610bf4a476e5a24b80a80a09e456e) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cpus` | [items.object.spec.gc_spec.infra.hw_info.cpu.cpus](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dc1ba96dfc63c8ae2154812f7db33772a984a99ec2d8ca97618f4491a838b88b) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.model` | [items.object.spec.gc_spec.infra.hw_info.cpu.model](data-sources--site_registrations_by_state--reference--group-002.md#canonical-06b04082be0c4b925c91a4b72d6dc7f360e3e56f49485b891799857f762abf72) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.speed` | [items.object.spec.gc_spec.infra.hw_info.cpu.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6d94df9d9b0c1956a7fe8bb6854f82a1c307b215424bc3875dffe092ac325d27) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.threads` | [items.object.spec.gc_spec.infra.hw_info.cpu.threads](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fac0d5b647a190f4a2032d7dbad57a254b6fdb8f76226956d919e13d5d2c75b5) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.vendor` | [items.object.spec.gc_spec.infra.hw_info.cpu.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-23527fceb4f44dda70dd4a73c792a38d84e435922fc318a936301b3140b0cbd9) |
| `items.object.spec.gc_spec.infra.hw_info.gpu` | [items.object.spec.gc_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b0e454e1953faddfe3ca688ce5335dbca01f81298608aded1368ddad9aaab12d) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.cuda_version` | [items.object.spec.gc_spec.infra.hw_info.gpu.cuda_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0dc554573a13c96015dff0d16f639cd2752b2f9f9947653cf92cb8c33025ddc5) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.driver_version` | [items.object.spec.gc_spec.infra.hw_info.gpu.driver_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-91ef5dca8f0372b472afcc8adcb93c51b2346752d6a7c8711eb00b781f2c959d) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device](data-sources--site_registrations_by_state--reference--group-002.md#canonical-95145099fe49cbb34de8b729e86def70bc0fd65ce0df81ef6da5726c6f603fbe) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.id` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-56ca665ce2911fb8756ad0f92b0dd6acfa5566e2fc65a62401888f8fa3860e32) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.processes` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.processes](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b44bf50ecb4856acc5c74e6e0b3dcd8eceb2da5a091f7efb6e8cf512af8cfadd) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.product_name` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.product_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ad718f896a3b9c5947618f3641145cccac703aa329f55ae319f179868d30fbdd) |
| `items.object.spec.gc_spec.infra.hw_info.kernel` | [items.object.spec.gc_spec.infra.hw_info.kernel](data-sources--site_registrations_by_state--reference--group-002.md#canonical-de7cdf8cce51e3e50b6383946b86fe39215b98cd61af43e5aba55740b6711a49) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.architecture` | [items.object.spec.gc_spec.infra.hw_info.kernel.architecture](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d979238813bbc7159eed181e82e67dfa61645e95a8e6b25053f87ca7523b023b) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.release` | [items.object.spec.gc_spec.infra.hw_info.kernel.release](data-sources--site_registrations_by_state--reference--group-002.md#canonical-26c2ff4aa2063eac416de10a2d1dd1c6bc1e42097206b981fe07f8c14dd64d21) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.version` | [items.object.spec.gc_spec.infra.hw_info.kernel.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e3528b808bc4fe38b3eb2721aba13042694bca22f15dbaded5695317705ff529) |
| `items.object.spec.gc_spec.infra.hw_info.memory` | [items.object.spec.gc_spec.infra.hw_info.memory](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f0b7d73a5ab68e77ff07469a7aa49f72a16271b6573cad9c5a5b64302040c629) |
| `items.object.spec.gc_spec.infra.hw_info.memory.size_mb` | [items.object.spec.gc_spec.infra.hw_info.memory.size_mb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6ebd0b00763896ad5536f79e23368c043cc6deeba6e2fcfb40851fc130e8ef8c) |
| `items.object.spec.gc_spec.infra.hw_info.memory.speed` | [items.object.spec.gc_spec.infra.hw_info.memory.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-afe50c5c3dac3416bcf6b71f816ea417b565091e9656750cdb1adb8aebd00414) |
| `items.object.spec.gc_spec.infra.hw_info.memory.type` | [items.object.spec.gc_spec.infra.hw_info.memory.type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a02a0734eaa3ec1d2a1be672ca480679a7e2b2b20b1a88fcbf090e8feb68a153) |
| `items.object.spec.gc_spec.infra.hw_info.network` | [items.object.spec.gc_spec.infra.hw_info.network](data-sources--site_registrations_by_state--reference--group-002.md#canonical-97317d2baa152870d3b334d3ffccea4b404938a2f188f2aaded921ac2c95779c) |
| `items.object.spec.gc_spec.infra.hw_info.network.driver` | [items.object.spec.gc_spec.infra.hw_info.network.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fe99d7b2560e365ee3762cc65a14ea376bf8c60fbe55a5cc7fc7e2f47f460542) |
| `items.object.spec.gc_spec.infra.hw_info.network.ip_address` | [items.object.spec.gc_spec.infra.hw_info.network.ip_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d4aa0de5809c8b532df232bb0e5f4bdbbc1f60e0dabd024182fdea2c3d61f730) |
| `items.object.spec.gc_spec.infra.hw_info.network.link_quality` | [items.object.spec.gc_spec.infra.hw_info.network.link_quality](data-sources--site_registrations_by_state--reference--group-002.md#canonical-828b33f1bec6e95895c0c12d2f35eff0cf149df7b293e36a1aec8c3dd5bdad8a) |
| `items.object.spec.gc_spec.infra.hw_info.network.link_type` | [items.object.spec.gc_spec.infra.hw_info.network.link_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b7d97160f81e844c8da1e60bb0563a68aea0bc2e67eb8e7fd2ed5c4152f1a609) |
| `items.object.spec.gc_spec.infra.hw_info.network.mac_address` | [items.object.spec.gc_spec.infra.hw_info.network.mac_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-44cd9651c0e06ae30faf8db15ba3c479f23718caa25af0268959ae88b6403f2a) |
| `items.object.spec.gc_spec.infra.hw_info.network.name` | [items.object.spec.gc_spec.infra.hw_info.network.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-65b486c4c8ebf65c3ab24f3af6f07d23ab3817fa121a0c60379a859c094f33e4) |
| `items.object.spec.gc_spec.infra.hw_info.network.port` | [items.object.spec.gc_spec.infra.hw_info.network.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-23c643470910a0b444eedd43cbdf60ada4271e706142727ca16a0ee05acdfa0e) |
| `items.object.spec.gc_spec.infra.hw_info.network.speed` | [items.object.spec.gc_spec.infra.hw_info.network.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9baa24ce09a564b952fe9055dfaceffdfa265d819cdd9c90f63404499d386004) |
| `items.object.spec.gc_spec.infra.hw_info.numa_nodes` | [items.object.spec.gc_spec.infra.hw_info.numa_nodes](data-sources--site_registrations_by_state--reference--group-002.md#canonical-328d24adea38cb2390a3b42b8bcf58ee9f18864a54fd03aea9cf922642a5d57b) |
| `items.object.spec.gc_spec.infra.hw_info.os` | [items.object.spec.gc_spec.infra.hw_info.os](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1d70ca70676db73de6faa2b28a7e2472d661ad1c2ccc16461e869638dbb3f6f5) |
| `items.object.spec.gc_spec.infra.hw_info.os.architecture` | [items.object.spec.gc_spec.infra.hw_info.os.architecture](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2bfcf3ff888a354dc14f965ac31d66d5fc72402f3893b7b3c7f536d7ad7e74ea) |
| `items.object.spec.gc_spec.infra.hw_info.os.name` | [items.object.spec.gc_spec.infra.hw_info.os.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-99a4fa75b5531bc51191494f5900de18be252f3f0de409c0dc9aba9ca6327c72) |
| `items.object.spec.gc_spec.infra.hw_info.os.release` | [items.object.spec.gc_spec.infra.hw_info.os.release](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3dc7d1b1d86c0169702ab65394a924cff683580a8442cd190c34f723a1f5316a) |
| `items.object.spec.gc_spec.infra.hw_info.os.vendor` | [items.object.spec.gc_spec.infra.hw_info.os.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a847b21f107cc5c71eb8f6a0fa3341249513449932f9ccf8454c79f78f0b4854) |
| `items.object.spec.gc_spec.infra.hw_info.os.version` | [items.object.spec.gc_spec.infra.hw_info.os.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0c222c8387ab6869a3c11982914c6396f6dd0c0d54c7f5d764445b6406f05ff4) |
| `items.object.spec.gc_spec.infra.hw_info.product` | [items.object.spec.gc_spec.infra.hw_info.product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cffd3752ff5df7b365d232fb2704d24b50549bf7902af11a19d79c4ee56e4f45) |
| `items.object.spec.gc_spec.infra.hw_info.product.name` | [items.object.spec.gc_spec.infra.hw_info.product.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ca6994472a003a9bd09aa6027600907c995b506db56afacc7818b8a3fb8de54f) |
| `items.object.spec.gc_spec.infra.hw_info.product.serial` | [items.object.spec.gc_spec.infra.hw_info.product.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6ef0c55aa31612dbd56457e8e753ec2dfae4b41a1fba1ae57ec546d9690a08e0) |
| `items.object.spec.gc_spec.infra.hw_info.product.vendor` | [items.object.spec.gc_spec.infra.hw_info.product.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-8c4777b610d6e7372a3d54eb2b534d1519bfe96d6b1ce085f71bc07d246bb2ca) |
| `items.object.spec.gc_spec.infra.hw_info.product.version` | [items.object.spec.gc_spec.infra.hw_info.product.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-188d1c6d86d1a848bbb5748769fa7a44a1dfe46833b2d786885109f865de996b) |
| `items.object.spec.gc_spec.infra.hw_info.storage` | [items.object.spec.gc_spec.infra.hw_info.storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c1e4c7cac0b522bf0cd8e48fa40f25c4107fcc409ab52bed74b468f6487aa18a) |
| `items.object.spec.gc_spec.infra.hw_info.storage.driver` | [items.object.spec.gc_spec.infra.hw_info.storage.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-43d0c7e4b9fdf4e5e68ac5a09a8d9068a83a478e3557e90895ff7441419d5c98) |
| `items.object.spec.gc_spec.infra.hw_info.storage.model` | [items.object.spec.gc_spec.infra.hw_info.storage.model](data-sources--site_registrations_by_state--reference--group-002.md#canonical-62c73b391286fe15401c333beee569c96e7c2058ffc4d7a08d4dc1b4caa05981) |
| `items.object.spec.gc_spec.infra.hw_info.storage.name` | [items.object.spec.gc_spec.infra.hw_info.storage.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-38fe241d763f4f283728d16a6f03a44962274522929a9493fe2386c2df5b2aa9) |
| `items.object.spec.gc_spec.infra.hw_info.storage.serial` | [items.object.spec.gc_spec.infra.hw_info.storage.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1221addcfaab36807a647e5019242ea6b84bfce51af8586035b7f9476692e7b3) |
| `items.object.spec.gc_spec.infra.hw_info.storage.size_gb` | [items.object.spec.gc_spec.infra.hw_info.storage.size_gb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fd96f0e07c1a7547a96782f76f580a776ae0fd0c44c415164651592c7967a14a) |
| `items.object.spec.gc_spec.infra.hw_info.storage.vendor` | [items.object.spec.gc_spec.infra.hw_info.storage.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-664269373c5ee79d8d05167f1f680729d597a4cba78227ffe6a60693cdcd46cf) |
| `items.object.spec.gc_spec.infra.hw_info.usb` | [items.object.spec.gc_spec.infra.hw_info.usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-5a03f75b70340019fdc1a99bc359d485de43a29a5ed9a65fcc90e30b8fd5167b) |
| `items.object.spec.gc_spec.infra.hw_info.usb.address` | [items.object.spec.gc_spec.infra.hw_info.usb.address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3c8a7318c019d267cbc632658b8f1b3ab59d964927c63ede1f5474d2b0fc6093) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_class` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cabbfbc822d8bab82348010413f57668605c1c7f5a99735ba620d0e36cc772f2) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_protocol` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_protocol](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a61336a26099d7ba0b20441db5f346242f88c11d1b1f66f87794695b658b469d) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_sub_class` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_sub_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cfaaef3b456c304e05d651a651c6dc9f5f5cd2e80b2098f2fd329e99f72d0534) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_max_packet_size` | [items.object.spec.gc_spec.infra.hw_info.usb.b_max_packet_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-cd7338598ae0114e49ddaf1172883640c0207697d5f91095a988d3ed725b988e) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bcd_device` | [items.object.spec.gc_spec.infra.hw_info.usb.bcd_device](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6387cb4f1488b87d8a4f375404f9149717d74bb05629c29c19d3de354aabd163) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bcd_usb` | [items.object.spec.gc_spec.infra.hw_info.usb.bcd_usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7a3561022325e1ef3f20b3620951fd6e20bc4d043bf3d1534c5ab09def22db4c) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bus` | [items.object.spec.gc_spec.infra.hw_info.usb.bus](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e486fb8a3d7f7b2c50e978dcbdf9fc6d52bb5e2fe91c83d6a96c2ff597950cbd) |
| `items.object.spec.gc_spec.infra.hw_info.usb.description_spec` | [items.object.spec.gc_spec.infra.hw_info.usb.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-33133cdb8a96a487f15be098dd0c47a201e864716fdb6233d820d5ce61d1b1c6) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_manufacturer` | [items.object.spec.gc_spec.infra.hw_info.usb.i_manufacturer](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1ab6df23c2a7102674f8663bbed847a9ecb84e20c45aa2912ffac05729dd18ac) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_product` | [items.object.spec.gc_spec.infra.hw_info.usb.i_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-efa0adbcd0430ee9b3b4e8411678195d51f064e18474d9507dcdb9e6cb10b6b7) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_serial` | [items.object.spec.gc_spec.infra.hw_info.usb.i_serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-757134822062fd4506567417b6f0d1ad0570f8e613eb6df4d7cbd07abe51b55c) |
| `items.object.spec.gc_spec.infra.hw_info.usb.id_product` | [items.object.spec.gc_spec.infra.hw_info.usb.id_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-70fca0761beb15d10b4ed47bdc3cd346af6dc37fc9d5251f72fe60089d676947) |
| `items.object.spec.gc_spec.infra.hw_info.usb.id_vendor` | [items.object.spec.gc_spec.infra.hw_info.usb.id_vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-af3f4d4e16aa9ea492a7fdcbdecaff7440f85835616ef587529dfe0f2ca48a99) |
| `items.object.spec.gc_spec.infra.hw_info.usb.port` | [items.object.spec.gc_spec.infra.hw_info.usb.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-bb8abe53db2d15791794ec767a0c940b7076b62defd3e8a7e7f05cfb1763c4c0) |
| `items.object.spec.gc_spec.infra.hw_info.usb.product_name` | [items.object.spec.gc_spec.infra.hw_info.usb.product_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-72eb10e984f2f22456b18785cb44de9ac12237f20a6eb04603b90f04c81ca2ee) |
| `items.object.spec.gc_spec.infra.hw_info.usb.speed` | [items.object.spec.gc_spec.infra.hw_info.usb.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9d195a97e3322564ac6db08a4bd7534b534bfb5451f0e88864d87609f438c3b1) |
| `items.object.spec.gc_spec.infra.hw_info.usb.usb_type` | [items.object.spec.gc_spec.infra.hw_info.usb.usb_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d1df6bf5c1a565fe96cd32e1a8f695250516fb87f54212d17d9989caa13b787d) |
| `items.object.spec.gc_spec.infra.hw_info.usb.vendor_name` | [items.object.spec.gc_spec.infra.hw_info.usb.vendor_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-5eb5fb46119f3f2bab304e9c28cd30525672950fe1e94cd6d9f612ac5e09cc31) |
| `items.object.spec.gc_spec.infra.instance_id` | [items.object.spec.gc_spec.infra.instance_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4a4c8e6d0538bbc0dd1932910aeb99b0ba831e550ef8a7bf03dcb777a787d29f) |
| `items.object.spec.gc_spec.infra.interfaces` | [items.object.spec.gc_spec.infra.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3ecc8957218788c8f9b6730141b97e803025776cf6ab221e0229cb4374da230f) |
| `items.object.spec.gc_spec.infra.internet_proxy` | [items.object.spec.gc_spec.infra.internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-68fbb1cd40325b737843573d1c58d96238c9ba18fe9a577c2fb34dd85c7c73a4) |
| `items.object.spec.gc_spec.infra.internet_proxy.http_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.http_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-da9c185923c1db580327ab3546cc10b4b55fd687b31c8811caa5330382a2483a) |
| `items.object.spec.gc_spec.infra.internet_proxy.https_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.https_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ec488ef851907fdf590d8753d45fc24991bd58d4f6efd3f2c5bed31034958daa) |
| `items.object.spec.gc_spec.infra.internet_proxy.no_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.no_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3d25bc277dec8122f959a0313c7122ca21a651d10f22887de0fe83d896d03eec) |
| `items.object.spec.gc_spec.infra.internet_proxy.proxy_cacert_url` | [items.object.spec.gc_spec.infra.internet_proxy.proxy_cacert_url](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e9ae8f06884a0602d60ba7b22002dad8e04faefb6377ff5d1743d163e9ea1841) |
| `items.object.spec.gc_spec.infra.is_slo_static` | [items.object.spec.gc_spec.infra.is_slo_static](data-sources--site_registrations_by_state--reference--group-002.md#canonical-170fffeade3fd41493e70dc45f807e5556bf563b7ab86d1a56303a6a7792582d) |
| `items.object.spec.gc_spec.infra.machine_id` | [items.object.spec.gc_spec.infra.machine_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-713fcaf1c1443a252af2e0b378fd4e5cced3c8b538347e4ac363d22ebcb2b294) |
| `items.object.spec.gc_spec.infra.provider_ref` | [items.object.spec.gc_spec.infra.provider_ref](data-sources--site_registrations_by_state--reference--group-002.md#canonical-fc8b5bd3d765048e0cf0d98b3966e337e80cfabf0972832d17e809edbf9a9ca6) |
| `items.object.spec.gc_spec.infra.sw_info` | [items.object.spec.gc_spec.infra.sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c931bf4a06252caa11114a17821da1219b9bdbde6e0e8e4a92ec92c2b4935787) |
| `items.object.spec.gc_spec.infra.sw_info.sw_version` | [items.object.spec.gc_spec.infra.sw_info.sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f1b35ec83a3a8366bcaa6579a513863cc0d681dffd32b089e81df0e1b5ae9ef1) |
| `items.object.spec.gc_spec.infra.timestamp` | [items.object.spec.gc_spec.infra.timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f3555867554a616179db7183af38f53087a3e7d9f14b7fc5f7332587c8dbdfea) |
| `items.object.spec.gc_spec.infra.zone` | [items.object.spec.gc_spec.infra.zone](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4bfab0c687022401bc840146d2fd4191cabb95b35ace40f5e45711558c1a22ed) |
| `items.object.spec.gc_spec.passport` | [items.object.spec.gc_spec.passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d3744cec5955334673610363b1d804e65a3a650a0c51eb5b0256a7edd41ee9cd) |
| `items.object.spec.gc_spec.passport.cluster_name` | [items.object.spec.gc_spec.passport.cluster_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-914f7c21fe2d8c21552b749854b8946d01cbfaa59eaa5167617fb519b2fcb69a) |
| `items.object.spec.gc_spec.passport.cluster_size` | [items.object.spec.gc_spec.passport.cluster_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-9489ec9c6698cbdc448778075f4078f86aeec40bba9c7d29af74f0d2230c07b0) |
| `items.object.spec.gc_spec.passport.cluster_type` | [items.object.spec.gc_spec.passport.cluster_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a917cfe19a101d2de42ba7809c8ff6b59b56f1f5c8b5f32703c8022b922ab03d) |
| `items.object.spec.gc_spec.passport.default_os_version` | [items.object.spec.gc_spec.passport.default_os_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ebafb89468cd6ecc9d2ff71caf117e4734eb7caa8345daaf5a2d3e699a3dea75) |
| `items.object.spec.gc_spec.passport.default_sw_version` | [items.object.spec.gc_spec.passport.default_sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a79f637b4de194f9f02073b3a42e2cea6b4009b35d0389d506135d7df58b0514) |
| `items.object.spec.gc_spec.passport.latitude` | [items.object.spec.gc_spec.passport.latitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-08aa37afe05d95e3e3adaa866dad7122a2e8a96aca255084bcf1a2eba55d16b6) |
| `items.object.spec.gc_spec.passport.longitude` | [items.object.spec.gc_spec.passport.longitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-279d53c735553e33198ea8859dfdee4435a69f8c73d49f8f5a95d307569a79bc) |
| `items.object.spec.gc_spec.passport.operating_system_version` | [items.object.spec.gc_spec.passport.operating_system_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4f6fdfeca322040d71030b54a09d3d0f6062338a51c24b7aab7527f531e1411f) |
| `items.object.spec.gc_spec.passport.private_network_name` | [items.object.spec.gc_spec.passport.private_network_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-028c802c4b9e9a9f0b5fe88139031adca3cbc03b71bcf278ecd60d4bf13740fe) |
| `items.object.spec.gc_spec.passport.volterra_software_version` | [items.object.spec.gc_spec.passport.volterra_software_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-090d5424bc3879031658ec0017e806e08940d2e1b6342c44a1b2a12b63c08247) |
| `items.object.spec.gc_spec.role` | [items.object.spec.gc_spec.role](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4929927a3b373fdfc819498721fc86f5e379cbfe347178b9318b1b7b8c4f19df) |
| `items.object.spec.gc_spec.site` | [items.object.spec.gc_spec.site](data-sources--site_registrations_by_state--reference--group-002.md#canonical-95c44ca39900f6eb8fcf74b5a3339068bf861c622bf5eed268399934fdff0ee2) |
| `items.object.spec.gc_spec.site.kind` | [items.object.spec.gc_spec.site.kind](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a1270937a15b47aba99af32f1b3167b4e1389a3da553b02864016e6f9ab924d4) |
| `items.object.spec.gc_spec.site.name` | [items.object.spec.gc_spec.site.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-931f334a9d68a841d0ce530ba664a12578b02d16848ce7ad50d811f0c54eb984) |
| `items.object.spec.gc_spec.site.namespace` | [items.object.spec.gc_spec.site.namespace](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a68ea06409460eeb8461bf0d7756bd32acccc0e210af864665fde2356a11b7d3) |
| `items.object.spec.gc_spec.site.tenant` | [items.object.spec.gc_spec.site.tenant](data-sources--site_registrations_by_state--reference--group-002.md#canonical-676a93820c379dfa978fc205a3ccc3b90e70613d7728a2f053f92de6452b060b) |
| `items.object.spec.gc_spec.site.uid` | [items.object.spec.gc_spec.site.uid](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1a7517e05b121b4b7095b9ec1210c037ea216242ae8cb337b505753d162c6410) |
| `items.object.spec.gc_spec.token` | [items.object.spec.gc_spec.token](data-sources--site_registrations_by_state--reference--group-002.md#canonical-59b75529f65bf3b113045e2b008aaff2dec09381dc17822f8ceaf66757416af8) |
| `items.object.spec.gc_spec.tunnel_type` | [items.object.spec.gc_spec.tunnel_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-66561d6cf683381b54c498dd47cce506e3d0b2f56beb1e569ec514b39927f67c) |
| `items.object.status` | [items.object.status](data-sources--site_registrations_by_state--reference--group-002.md#canonical-60af0f9a068c583e14a1e32c8064e119a3927cb2f68755f586c371b658327ca7) |
| `items.object.status.current_state` | [items.object.status.current_state](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d2717312f3c3e05ccaa3333581ca066087e7bc1c78b5bbdd3076f4419130f295) |
| `items.object.status.object_status` | [items.object.status.object_status](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4141af61a19ac899a4bd41d3bd5888eb6f1a7d4f284b8b17614daf28be40a113) |
| `items.object.status.object_status.code` | [items.object.status.object_status.code](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ba2443b9b5aed5fdf316cf91669171af8aa8fac01c31c58865eb1148e33cd3b0) |
| `items.object.status.object_status.reason` | [items.object.status.object_status.reason](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a6fac05f60f8b61f88ec46ed5b765c39694a144937dd35b737a4076f0a8afd55) |
| `items.object.status.object_status.status` | [items.object.status.object_status.status](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1e3e1bd75b6eaee42265ab03b93d3406fd39e76e02fad9cf1dc72dd353fe45b2) |
| `items.object.status.parent_current_state` | [items.object.status.parent_current_state](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0cabded1363bfd0508253b0c68253f2b26d0f99806383a8d057a8b65c3880b47) |
| `items.object.status.state_update_timestamp` | [items.object.status.state_update_timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-4fa5f5ee763a93aae8ce0e2dd2fbe0412f42cdbf5205386bb77644bbb32befe5) |
| `items.object.system_metadata` | [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3fddd31983a0525e7f0ec6e3785f407b7e2f1c51b18ec1353e035e0fe7b8c531) |
| `items.object.system_metadata.creation_timestamp` | [items.object.system_metadata.creation_timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-f11ba3dee7a1adeaa6de986eb7ec24b23711f9ce525ee73fc421bc943dfb7f9e) |
| `items.object.system_metadata.creator_class` | [items.object.system_metadata.creator_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-b3d8a380a575822059bbc4632c6f09de13f4c883376fced76efcb2f28fa09981) |
| `items.object.system_metadata.creator_cookie` | [items.object.system_metadata.creator_cookie](data-sources--site_registrations_by_state--reference--group-002.md#canonical-dc3ae8718c2ddff621dac2b01ef9741f455f098ba1f36a2a8e189746ff2029d4) |
| `items.object.system_metadata.creator_id` | [items.object.system_metadata.creator_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-990bfc1d8ef8470c32681f23e7b583db393c75591aeeab8ad44f85d654c14245) |
| `items.object.system_metadata.deletion_timestamp` | [items.object.system_metadata.deletion_timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e0f5eb327d553a93fec759890ce2ba12f47fb809c7b86398d7e2d01842424667) |
| `items.object.system_metadata.direct_ref_hash` | [items.object.system_metadata.direct_ref_hash](data-sources--site_registrations_by_state--reference--group-002.md#canonical-810bf65e38758d1aa09cfcd74172559fa667aac4a5f94770f012978188b22e8e) |
| `items.object.system_metadata.finalizers` | [items.object.system_metadata.finalizers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a6eb8abcac57f434e21e35e68c1e8a6d1e43b1ab8278e1a466bcc8baf21e5f6e) |
| `items.object.system_metadata.initializers` | [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-d120877eebf84e237abced7708e5648d20671f4f5cc2b1e4d8a3658092b3c076) |
| `items.object.system_metadata.initializers.pending` | [items.object.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-e7269a7c93daf6350ddf4108408227058332a74e017a5ce164bb65ef2eba8605) |
| `items.object.system_metadata.initializers.pending.name` | [items.object.system_metadata.initializers.pending.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-27fcb5f779a437fd6d83398fb9a9634e16e97c27054e07cd76ad8497828c8cbe) |
| `items.object.system_metadata.initializers.result` | [items.object.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-40b2d22aa5545bc7edb167d58ffd9f1773455f58318e3f9e825aa26499239f4d) |
| `items.object.system_metadata.initializers.result.code` | [items.object.system_metadata.initializers.result.code](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f561fad416bb3f7e0bbf3b7415e0b821712b1a265534a58069944bf07194c2c1) |
| `items.object.system_metadata.initializers.result.reason` | [items.object.system_metadata.initializers.result.reason](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7fe0b39ab84b5f41362a4c9013d1af4113ed15d5c46ea9c48c116946afead738) |
| `items.object.system_metadata.initializers.result.status` | [items.object.system_metadata.initializers.result.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-6d10e060e21ab13bd50d2956f3aa2fb8d5b084e30025cc3829a8340d01dc0937) |
| `items.object.system_metadata.labels` | [items.object.system_metadata.labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0bcb9d25c56f96b7de6a207a0a5817c61006debe6c7adc4e51c780e54c3f57c2) |
| `items.object.system_metadata.modification_timestamp` | [items.object.system_metadata.modification_timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a67cd973a0899102bc81bb25a372edd3e7f29f37c1964d2e36ff94c7c25637e1) |
| `items.object.system_metadata.namespace` | [items.object.system_metadata.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-06aaa240001d9b8379715572ea969b79d8ef416be24077e1b789890b664889c6) |
| `items.object.system_metadata.namespace.kind` | [items.object.system_metadata.namespace.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-4e6493cc1f999f9c20a7436290e195742172a0aac40dc163743bf4283b55c099) |
| `items.object.system_metadata.namespace.name` | [items.object.system_metadata.namespace.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-c90918569e14a85a2afdce69d9982ae3859ff5d6e4a579c7ac47758d7830b365) |
| `items.object.system_metadata.namespace.namespace` | [items.object.system_metadata.namespace.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f1c9a4a07d90696b390fa15b034ea52d29c935845118b25f107ebb1d38c42050) |
| `items.object.system_metadata.namespace.tenant` | [items.object.system_metadata.namespace.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-38a2226f176f16aab130ce69c0a8686cbddf8bb01411898db197434d2e04b9ab) |
| `items.object.system_metadata.namespace.uid` | [items.object.system_metadata.namespace.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-19348bbc2becbb6442f33372ba7b2ad95c4e8158eaeb2c81e35b634fe01177ea) |
| `items.object.system_metadata.object_index` | [items.object.system_metadata.object_index](data-sources--site_registrations_by_state--reference--group-002.md#canonical-efe46e75973e67491462d62bb0fa9c16000b68a67f607c5002eea2a76dec5de3) |
| `items.object.system_metadata.owner_view` | [items.object.system_metadata.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-a0e7b40d37dd980b0e7c1fd21b254a6d4a15044fd1f819e396264861b0ccaa91) |
| `items.object.system_metadata.owner_view.kind` | [items.object.system_metadata.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-8589f28c8836558e2b2bc3302ae1bdd12099492a4af5ba121a3cf418702c45c3) |
| `items.object.system_metadata.owner_view.name` | [items.object.system_metadata.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-8b400ac879b5509b90059b01b44ca1b55fcbb83942de9421b73734bfbf0bf5eb) |
| `items.object.system_metadata.owner_view.namespace` | [items.object.system_metadata.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1120ea5704a8d5f08944dce88ea999f2018d7005a2d8a14373bac735a154fba7) |
| `items.object.system_metadata.owner_view.uid` | [items.object.system_metadata.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-44db5d9b0f873fb458d30b1529d7969f9e512b6f793b333d40b7b50ce0cd4b72) |
| `items.object.system_metadata.revision` | [items.object.system_metadata.revision](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e14e0864f442cfa34998318d6860df420f47bfdd1b6e556bd64d273e2c0c546d) |
| `items.object.system_metadata.sre_disable` | [items.object.system_metadata.sre_disable](data-sources--site_registrations_by_state--reference--group-002.md#canonical-e8b31b656e3aa6e7f6ca01e241bb658a04c6e389d68bdc082aa9bf4824dcd415) |
| `items.object.system_metadata.tenant` | [items.object.system_metadata.tenant](data-sources--site_registrations_by_state--reference--group-002.md#canonical-31e3b9604ada4fb5fd1db26c2ee5a3a25181220e1a3950ea4cab1522d58d3050) |
| `items.object.system_metadata.trace_info` | [items.object.system_metadata.trace_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3a56780497bedfcf9a69136280f8048fd9ae754b0c549a1389d8f16de3290591) |
| `items.object.system_metadata.uid` | [items.object.system_metadata.uid](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a7b7e98e3555298d5d26e8776fa5f58ff83b0eed86b7d36b25e941a3c7c84079) |
| `items.object.system_metadata.vtrp_id` | [items.object.system_metadata.vtrp_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-25cc256a3cb0072728c4dea4d0ff44ba91cd664524070352c2899abcf0f3df5b) |
| `items.object.system_metadata.vtrp_stale` | [items.object.system_metadata.vtrp_stale](data-sources--site_registrations_by_state--reference--group-002.md#canonical-7fde7ab998b4dd873bd3c9c414a729caf88b8dda6cd593231e820be3bd45cceb) |
| `items.owner_view` | [items.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-dcadc76bdff1c3b3510abe16b29d68cd606e60675015d4e6429ac5d7170afd6e) |
| `items.owner_view.kind` | [items.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-640b9955805dd2e29f2759bd4d4a9472f448ddf8da38ed20c837ff8d62b43438) |
| `items.owner_view.name` | [items.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-015b1f4361b04d55fd582da935c3095b111d75dd5e60b35af6dad073a4d2cdf8) |
| `items.owner_view.namespace` | [items.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-6dbd08b5c1512c8d7b8ebabfd6f2779ab3cd054d7a7d0548230c17f40c5dc9d6) |
| `items.owner_view.uid` | [items.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-bf6b7fe8a20fb7b50e198980d8475e7de5a7729cea6b2775932633c9c901ebe0) |
| `items.system_metadata` | [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-fbf9d31e1277cfb4777d53ddcf1be885d2271d7c1d6873cfe06b9ace35387688) |
| `items.system_metadata.creation_timestamp` | [items.system_metadata.creation_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-d156ad066a0335215b33bf80d1039c2fd3ba416ddd62e8a04b37e7326ca96e19) |
| `items.system_metadata.creator_class` | [items.system_metadata.creator_class](data-sources--site_registrations_by_state--reference--group-003.md#canonical-dbe7e65fcbc5b82da466da72e29132f434e8308f803ce985eb1c1f12751a075c) |
| `items.system_metadata.creator_id` | [items.system_metadata.creator_id](data-sources--site_registrations_by_state--reference--group-003.md#canonical-31eecb18d371edf103fa8bd561feafe6eb3968cd278bc36f814173111f4bf64a) |
| `items.system_metadata.deletion_timestamp` | [items.system_metadata.deletion_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f9689973d7ab11ea22e93b94c2fd69fd4cdccaca9091e8328a4b20bc49b413f3) |
| `items.system_metadata.finalizers` | [items.system_metadata.finalizers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-eb733c20f3e48fb3d741c6392f74dd8c447f841909fdeb37a73e1437bd9278ec) |
| `items.system_metadata.initializers` | [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-e0fff57ec3f45ae35f2ff8895289bffc542ae4e9bd38192d04fde852833387ee) |
| `items.system_metadata.initializers.pending` | [items.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2626063209967fe3bda32b6525355cac5cb7e8436871d792cc53a7edae3b19e7) |
| `items.system_metadata.initializers.pending.name` | [items.system_metadata.initializers.pending.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-55b8fab150b784817cf4d166129a44147b9e5990879c160d1e4410ec648ce8c4) |
| `items.system_metadata.initializers.result` | [items.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-e764da52e50848efbb1ab52820befcd88ebccb09416e8928f76853957692735e) |
| `items.system_metadata.initializers.result.code` | [items.system_metadata.initializers.result.code](data-sources--site_registrations_by_state--reference--group-003.md#canonical-4f1dec2668141d20df129d4d9363dc6ee685e87cc3363557bb322ba6b9820364) |
| `items.system_metadata.initializers.result.reason` | [items.system_metadata.initializers.result.reason](data-sources--site_registrations_by_state--reference--group-003.md#canonical-5005caf20fa9de05ce64fc2daff8d73c763b099b8f0a75d98af0b44fd55eb0d7) |
| `items.system_metadata.initializers.result.status` | [items.system_metadata.initializers.result.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-e44f62feaddb4435569642afbf717cf04e6078e03602809e51e507532f566f89) |
| `items.system_metadata.labels` | [items.system_metadata.labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-981a2260cf39b33bf4529ea1f3d45d1b012c27f195877640d12befcf586f0c2e) |
| `items.system_metadata.modification_timestamp` | [items.system_metadata.modification_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-068be60826a873b7fbd5c56770c232e2f6f4dcc31c4459fbb75dc5026144d680) |
| `items.system_metadata.object_index` | [items.system_metadata.object_index](data-sources--site_registrations_by_state--reference--group-003.md#canonical-9a1ae02dc798c9a59e0ef36c1e21b24bdf098dcd2856c7b82982b5fde28df68c) |
| `items.system_metadata.owner_view` | [items.system_metadata.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2e76e92ce2c67503c6ee58b6c4cbeabe2ad804ad4b3f8cca7f37958f021d6d14) |
| `items.system_metadata.owner_view.kind` | [items.system_metadata.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-dbb710c2e186beed0d35fe1eb8cf21ad44ee884736e4a6a4171139922063c757) |
| `items.system_metadata.owner_view.name` | [items.system_metadata.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-5ea5c115c0dcec04c31213bbc82ba79d6ddc006e182127897e77770f2c09626c) |
| `items.system_metadata.owner_view.namespace` | [items.system_metadata.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f3291aa7e2891010d68e1b5fa64d5c96de9297c623f90f14d8552ae3760a97d5) |
| `items.system_metadata.owner_view.uid` | [items.system_metadata.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-240766e3f3c4d11552326242b526c1d11956e483ecd23eda4cc01736c9bc7dd9) |
| `items.system_metadata.tenant` | [items.system_metadata.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-56370735295d422a9ec381396cd277a276d4f02bce4771a18462850c9a72362f) |
| `items.system_metadata.uid` | [items.system_metadata.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3e661322f5688ff710c516107d8b4a38244c7ec7de0ed8c5e89de8072f2ac9ef) |
| `items.tenant` | [items.tenant](data-sources--site_registrations_by_state--reference--group-001.md#canonical-fc799c70b6c9070cdf778c121091811493916eebe5de13db2ee362cd1a46f621) |
| `items.uid` | [items.uid](data-sources--site_registrations_by_state--reference--group-001.md#canonical-be6af56fa423b26fe17701ee5aae96329b80176706c5f9f2af6c277ee2104653) |
| `namespace` | [namespace](data-sources--site_registrations_by_state--reference--group-001.md#canonical-750495cf0721392d1e9cb9e63d7b7b68e548727a0b3b88b8ada35b09fea3352f) |
| `state` | [state](data-sources--site_registrations_by_state--reference--group-001.md#canonical-ddac6fd41c18336558c312ac40b1b715b84013562dde2837db7db02700f1c7f4) |

<a id="canonical-292d90378616739ac6868b14399ca8e82284733d328053cb267c688602a7bd27"></a>

## Next pages — Property reference / 872e9837982c / 7

- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-db84e5ed548d2312394a53b8856bf8998548c9d7b05e61eff5d11d998c7a9c32)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-db84e5ed548d2312394a53b8856bf8998548c9d7b05e61eff5d11d998c7a9c32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85b611ffbdb77271eeb04a360581602e57fa468319784a9c02fa9e82e73cf976"></a>

## errors — errors / b5c38e2e37df / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- errors

<a id="canonical-2f4ce87b1f910e3b1bdc4af0aa25520b1351013f3b73d718c8c1979cbe8295d2"></a>

Type: `"list"`. Computed.

Errors(if any) while listing items from collection.

<a id="canonical-2662e3405e63df00460161be0777435f8b9fb34f413ff7af48a5ce997cdcd582"></a>

## Direct properties — errors / b5c38e2e37df / 3

<a id="canonical-25cffc968edf0e5c6f0b53519e31fbb326d1c27b3aa9e937a5d0f7194278c23a"></a>

<a id="canonical-fb3a9486ffda9107a5b996e3222af80b45b399e07892a9bddb47219a4d33e7bf"></a>

## code property — errors / b5c38e2e37df / 4

Type: `"string"`. Computed.

\[Enum: EOK|EPERMS|EBADINPUT|ENOTFOUND|EEXISTS|EUNKNOWN|ESERIALIZE|EINTERNAL|EPARTIAL\] Union of all
possible error-codes from system - EOK: No error - EPERMS: Permissions error - EBADINPUT: Input is
not correct - ENOTFOUND: Not found - EEXISTS: Already exists - EUNKNOWN: Unknown/catchall error -
ESERIALIZE: Error in serializing/de-serializing - EINTERNAL: Server error - EPARTIAL.. Possible
values are \`EOK\`, \`EPERMS\`, \`EBADINPUT\`, \`ENOTFOUND\`, \`EEXISTS\`, \`EUNKNOWN\`,
\`ESERIALIZE\`, \`EINTERNAL\`, \`EPARTIAL\`. Defaults to \`EOK\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EOK",
    "EPERMS",
    "EBADINPUT",
    "ENOTFOUND",
    "EEXISTS",
    "EUNKNOWN",
    "ESERIALIZE",
    "EINTERNAL",
    "EPARTIAL"),
}
```

- [error_obj](data-sources--site_registrations_by_state--reference--group-001.md#canonical-449475b15da44f2ddfa3132d329b06b2e074021a149d1bd0d816ceb237e12fcf): complete subsection reference.

<a id="canonical-12b8e0447c6d3e4608b302c4b5f1cd71e987d26611efdde749d79f02a63302a4"></a>

<a id="canonical-930c141cf68030d33acfeb2ad9836f727ce7785594d538e404bd6a3eb7362083"></a>

## message property — errors / b5c38e2e37df / 5

Type: `"string"`. Computed.

Message. A human readable string of the error.

<a id="canonical-5fa0909054b2ef2ecfca9ef7de39082c56224da805202b569ec2af9a9561fc0c"></a>

## Next pages — errors / b5c38e2e37df / 6

- [errors.error_obj](data-sources--site_registrations_by_state--reference--group-001.md#canonical-449475b15da44f2ddfa3132d329b06b2e074021a149d1bd0d816ceb237e12fcf)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-449475b15da44f2ddfa3132d329b06b2e074021a149d1bd0d816ceb237e12fcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bae6676c12b3bede9f23a4150d38262cb4bde914f83a388757a2f2f95529efd7"></a>

## errors.error_obj — errors.error_obj / 5db57d2480b6 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-db84e5ed548d2312394a53b8856bf8998548c9d7b05e61eff5d11d998c7a9c32)
- errors.error_obj

<a id="canonical-8ad1e349cd34b26917cbc303d0ae7200a1d8055a09e084faced72212f9c70628"></a>

Type: `"single"`. Computed.

Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of
the serialized message. Protobuf library provides support to pack/unpack Any values in the form of
utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..

<a id="canonical-94aeee645e5af13a5113dcf17b37bc822562ff98313604401684a3f609dde41a"></a>

## Direct properties — errors.error_obj / 5db57d2480b6 / 3

<a id="canonical-c440f0f45f1499eabd9314f07fb9c6ef6021ecbfa979fafc02a587334ed34fcd"></a>

<a id="canonical-9a9ddc28cf89966b8d91930f907f8b18ae244dc69c6a28e5ce51ef966ac4f318"></a>

## type_url property — errors.error_obj / 5db57d2480b6 / 4

Type: `"string"`. Computed.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="canonical-402810b003af9d69f8d8945d8d2aed114c93840b0faa48cfedd689fe68284f5f"></a>

<a id="canonical-77495732aa4090bf5e67448611d3c9bb31c81aeac941fe8db4f79b4afe49ca18"></a>

## value property — errors.error_obj / 5db57d2480b6 / 5

Type: `"string"`. Computed.

Must be a valid serialized protocol buffer of the above specified type.

<a id="canonical-f2e1cfedeb9b19a208aaacb5bd1dbab5b96478988d4676ee33541c56e6fda801"></a>

## Next pages — errors.error_obj / 5db57d2480b6 / 6

- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-db84e5ed548d2312394a53b8856bf8998548c9d7b05e61eff5d11d998c7a9c32)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41a5adb6fcb40129f34d7d59224f1859dc0ec4aee14f2053796500079c1016e8"></a>

## items — items / 9d8b0ec7b3f7 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- items

<a id="canonical-bb2003288b68aef55834c33c75355b55d2922d3755d6069fd39fe2760f789433"></a>

Type: `"list"`. Computed.

Items represents the collection in response.

<a id="canonical-ca48bbbb184bd8ba952522c9247f41c611237dcd76ec8bb4fe300f06aa148438"></a>

## Direct properties — items / 9d8b0ec7b3f7 / 3

- [annotations](data-sources--site_registrations_by_state--reference--group-001.md#canonical-555bd14e91bec80da8a2920bf9561bd22b44bdf96f1f1b18424ee8865105124c): complete subsection reference.

<a id="canonical-ef283be3ca3f9d677b227709099a53c8bd10af33e2bae343f82a9520da2ac40c"></a>

<a id="canonical-9e92eda380c551310322121962ae23164e7996d9a685450619a3ec1cd716d780"></a>

## description_spec property — items / 9d8b0ec7b3f7 / 4

Type: `"string"`. Computed.

The description set for this registration.

<a id="canonical-13ebd799ef106f8b5d55ac012555df9f855be724d0a428ba8a93b15ed0e8f128"></a>

<a id="canonical-4de5746dc2152ae1fded64cfefcd7d917a73014721606d3a603818112f78eb14"></a>

## disabled property — items / 9d8b0ec7b3f7 / 5

Type: `"bool"`. Computed.

Value of true indicates registration is administratively disabled.

- [get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a): complete subsection reference.

- [labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-17fd8dce3f241d20b2849aea4776416e7f4eba3bd4e5878eb686e2c1810c8a92): complete subsection reference.

- [metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-023b24f5298e72edda243d9c7d4450d9bba825a0c21bd53d79f0d2db8770db13): complete subsection reference.

<a id="canonical-490cb78e455ed2160dfde4b44ee2185d671292f1365c73a824de4018a3e3cd80"></a>

<a id="canonical-c061b282e18f8efd3b81638a480b47db19df29843906594a4c23b14cbfb278ff"></a>

## name property — items / 9d8b0ec7b3f7 / 6

Type: `"string"`. Computed.

Name. The name of this registration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-e900a43a6367722347c452a25bfb9f5c2ce8bcf522862da062b708ff270df72f"></a>

<a id="canonical-824e7d152613357553ad590f46fa94fedcbef657afc15715667b39ce35dc3091"></a>

## namespace property — items / 9d8b0ec7b3f7 / 7

Type: `"string"`. Computed.

Namespace. The namespace this item belongs to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

- [object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383): complete subsection reference.

- [owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-9f1c2c7bdb366f4d4e91f49c44a8532f1c3abe65a9944ef6d3941b5c1acd7aba): complete subsection reference.

- [system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba): complete subsection reference.

<a id="canonical-fc799c70b6c9070cdf778c121091811493916eebe5de13db2ee362cd1a46f621"></a>

<a id="canonical-fa5d8b2e9868d1cf4b1190d13b44327d7f68a9d912348698c0730c39a93d334a"></a>

## tenant property — items / 9d8b0ec7b3f7 / 8

Type: `"string"`. Computed.

Tenant. The tenant this item belongs to.

<a id="canonical-be6af56fa423b26fe17701ee5aae96329b80176706c5f9f2af6c277ee2104653"></a>

<a id="canonical-11ed5c730c9329c88ef5b5c438196e7f9d029be39f388528889617665b4ae0bc"></a>

## uid property — items / 9d8b0ec7b3f7 / 9

Type: `"string"`. Computed.

UID. The unique uid of this registration.

<a id="canonical-1f2b3ebe79e6977e80b90fb2b3931efe4fd19aa2331808049079c98e953bccc4"></a>

## Next pages — items / 9d8b0ec7b3f7 / 10

- [items.annotations](data-sources--site_registrations_by_state--reference--group-001.md#canonical-555bd14e91bec80da8a2920bf9561bd22b44bdf96f1f1b18424ee8865105124c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-17fd8dce3f241d20b2849aea4776416e7f4eba3bd4e5878eb686e2c1810c8a92)
- [items.metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-023b24f5298e72edda243d9c7d4450d9bba825a0c21bd53d79f0d2db8770db13)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-9f1c2c7bdb366f4d4e91f49c44a8532f1c3abe65a9944ef6d3941b5c1acd7aba)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-555bd14e91bec80da8a2920bf9561bd22b44bdf96f1f1b18424ee8865105124c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8b9219d8d2257ad513548295efb550da35f5da1acec8bd34def054b6c9b1e81"></a>

## items.annotations — items.annotations / 1139b9e99ec6 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- items.annotations

<a id="canonical-3e446def30581c52f75bf8f2d975c02bae8a79469e4549841ab93f58a8788150"></a>

Type: `"single"`. Computed.

The set of annotations present on this registration.

<a id="canonical-ff0c3fe73f5428302fc90ee4b604e6f39666506eb98b3560abf5f771affa9866"></a>

## Direct properties — items.annotations / 1139b9e99ec6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-feb01887789ddfa051af5e43248907229dadde221a8f4caf98887a7aed8a1bd6"></a>

## Next pages — items.annotations / 1139b9e99ec6 / 4

- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a124541caa94d9eec65535fc2d4f6d46ca1e985273b7147ed7c861499dec9f8"></a>

## items.get_spec — items.get_spec / 0871e8396b75 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- items.get_spec

<a id="canonical-a61da75adf7ac99a816ad36c9393b778ec33f841c1b4c18c99e5399f020978fc"></a>

Type: `"single"`. Computed.

GET Registration. GET registration specification.

<a id="canonical-2fc730193a1ea10c4a5a62e4a01a2038447382f5ef994cb0dc7361699d79dc56"></a>

## Direct properties — items.get_spec / 0871e8396b75 / 3

- [infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639): complete subsection reference.

- [passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a37b09b669b83cc97c0b71234ee1fddabb461cd59c3142e159060d295540bea4): complete subsection reference.

<a id="canonical-5eb4049c30e6b27338d6f25fbe573152300f6edde7e97963d4241e8c61f80c50"></a>

<a id="canonical-cf04e8ed9c0fedd50187aa95b21afe4aed619a228c7628c11872eef82951b8d7"></a>

## token property — items.get_spec / 0871e8396b75 / 4

Type: `"string"`. Computed.

Token is used for machine and tenant identification.

<a id="canonical-c053f10ebb7c546a5baa8a15117581ec7461c1880ac71598f4f3f4e704ae6344"></a>

## Next pages — items.get_spec / 0871e8396b75 / 5

- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-a37b09b669b83cc97c0b71234ee1fddabb461cd59c3142e159060d295540bea4)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7a52a9abc459f21dd1f24af267504ce2f1d9017f8410fb8e7cf9342af0b2b45"></a>

## items.get_spec.infra — items.get_spec.infra / f331fa8282b0 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- items.get_spec.infra

<a id="canonical-42a751645d5b1765384861c067aede853806a1fceb50f96e7a5ae70862afff39"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

<a id="canonical-feafff2f025259a4e618dde445bb8f517252b66d4ac81086e56b685dda98bd2d"></a>

## Direct properties — items.get_spec.infra / f331fa8282b0 / 3

<a id="canonical-9405634bc199506c9fd5fd800b7973fbcd2218e11b851bd33c02e73e20956c2b"></a>

<a id="canonical-930cf1c191f2ae9a16cb1ddcf5a99ddcac943b3f09a5ac289a41f7851916ad16"></a>

## availability_zone property — items.get_spec.infra / f331fa8282b0 / 4

Type: `"string"`. Computed.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

- [bond_config](data-sources--site_registrations_by_state--reference--group-001.md#canonical-4435f82f8e3d2797356df5c83c95a8af6726da069528e4f1087673c087b58d58): complete subsection reference.

<a id="canonical-15f40b2271d1d7c70f7b40b00e7c2357c12e0f32004a95aad92268053aa9e3b3"></a>

<a id="canonical-069dd11446ca7b991390e07fbd394b6020d9624066f3414b74e4d7727d1b7fad"></a>

## certified_hw property — items.get_spec.infra / f331fa8282b0 / 5

Type: `"string"`. Computed.

Certified HW name used to map with F5XC certified\_hardware definition.

<a id="canonical-a575c7f4d29b9e6eadd244c4b4d543a744e4dabd7b2dfa55a0fcc980c02faa2e"></a>

<a id="canonical-6158ec90a81e27c1fb3cc1e3763a64ec91db026a1d17e3fc801ab144e467d41f"></a>

## domain property — items.get_spec.infra / f331fa8282b0 / 6

Type: `"string"`. Computed.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

<a id="canonical-b8ebd44166694cb36ef234c4be574f5a80a445ae23b392e27f5d95e1ffec6916"></a>

<a id="canonical-b52662204f7c262cfb71f933a894a676f5c4d6d94ecf7b077819b368e422d89b"></a>

## hostname property — items.get_spec.infra / f331fa8282b0 / 7

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

- [hugepages](data-sources--site_registrations_by_state--reference--group-001.md#canonical-d51691d91b7ca9b4b7293926cc3acc1e1642be2a6cdd13cb31cf7ba91bd5b46b): complete subsection reference.

- [hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee): complete subsection reference.

<a id="canonical-501a411ec702f7b16858c2092bae5d7c3d3fdd935332cda6ef8c6e158cacae57"></a>

<a id="canonical-bf2a6369299955a3e4dd51f7f7e8dffc32e870faa9c2d2398c859e2e662cc083"></a>

## instance_id property — items.get_spec.infra / f331fa8282b0 / 8

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ecf1682a19b648b1b0dd3a5a1a409280736f9d15800698213820573eac450000): complete subsection reference.

- [internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-db96aa04c48fb8b434e5e56fb276ee65f4e4bf7b78412459cc398e815ac3c199): complete subsection reference.

<a id="canonical-29d9fc8f36ebc016a6adfe29fdb751815dd170578abb7b477557eb616bf6b711"></a>

<a id="canonical-b625327c8bfbbe0486bcdf86ff885e3f048d1f016f953f9a7e99e0607dfd9cb6"></a>

## is_slo_static property — items.get_spec.infra / f331fa8282b0 / 9

Type: `"bool"`. Computed.

Is SLO Static. Indicates whether the SLO is static.

<a id="canonical-b6445b17c6141b8ea28fa855f590e0d9fe0083d536567229439927ba7bbc70c3"></a>

<a id="canonical-1a941c99a1705fd094e7e69d4496607f619be329f232187d18bef85c90985677"></a>

## machine_id property — items.get_spec.infra / f331fa8282b0 / 10

Type: `"string"`. Computed.

Machine ID - generated by operating system.

<a id="canonical-252b57387e7011ea07310cf26acd99bfcaefb914612027fb014f2964eea5f49b"></a>

<a id="canonical-d549dd58c5110b75fe68876ca98376ef819298f38aa7386e923a9b1706b79fa6"></a>

## provider_ref property — items.get_spec.infra / f331fa8282b0 / 11

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3985e1b4b290870f84cf589f9c1bdec24b6823f3a7f632e9886ffacf0d5fa859): complete subsection reference.

<a id="canonical-ef5227baacac553c88d9159d730322df315b141840756055c181edab824448a9"></a>

<a id="canonical-742e82330ba1106d60326c2e78862e0377634f3adcd7bea42d50f4a4f396d5f6"></a>

## timestamp property — items.get_spec.infra / f331fa8282b0 / 12

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="canonical-da623860865c7e6f4a07ca2ce5b1279d24ae97cd45bbb83739f0c5b0bcc39ccf"></a>

<a id="canonical-7d650e8e39b0322078487e4838f0559354f856efc2c0ede2fbf46312c338abdd"></a>

## zone property — items.get_spec.infra / f331fa8282b0 / 13

Type: `"string"`. Computed.

Instance zone (or region), depends on provider.

<a id="canonical-1365b3d60e7ff47ab3e1ef834d3e56015c69351ae507e25ed101bdb61498ef09"></a>

## Next pages — items.get_spec.infra / f331fa8282b0 / 14

- [items.get_spec.infra.bond_config](data-sources--site_registrations_by_state--reference--group-001.md#canonical-4435f82f8e3d2797356df5c83c95a8af6726da069528e4f1087673c087b58d58)
- [items.get_spec.infra.hugepages](data-sources--site_registrations_by_state--reference--group-001.md#canonical-d51691d91b7ca9b4b7293926cc3acc1e1642be2a6cdd13cb31cf7ba91bd5b46b)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [items.get_spec.infra.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-ecf1682a19b648b1b0dd3a5a1a409280736f9d15800698213820573eac450000)
- [items.get_spec.infra.internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-db96aa04c48fb8b434e5e56fb276ee65f4e4bf7b78412459cc398e815ac3c199)
- [items.get_spec.infra.sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3985e1b4b290870f84cf589f9c1bdec24b6823f3a7f632e9886ffacf0d5fa859)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-4435f82f8e3d2797356df5c83c95a8af6726da069528e4f1087673c087b58d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86d7f193ff240efa94966111927f0db0f14e453c8b5ad5a75565ee5029c63e8d"></a>

## items.get_spec.infra.bond_config — items.get_spec.infra.bond_config / 9ad45cfdd81c / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- items.get_spec.infra.bond_config

<a id="canonical-f30725962dbd23246d5b09a4a032416b7d6988dd8a9079f548bfe240bfec9d76"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

<a id="canonical-d454490febb37c4df44dae9805c6d31f59278d2111f5b931880ead155849fb01"></a>

## Direct properties — items.get_spec.infra.bond_config / 9ad45cfdd81c / 3

<a id="canonical-e8f2143956811cc438de9a1ef3ea157800194de5fcb0bb717279b8bb8511c9c2"></a>

<a id="canonical-d6752217d277021ec52833f462ff06ff20d77ed8699cb895198350660db8b8aa"></a>

## interfaces property — items.get_spec.infra.bond_config / 9ad45cfdd81c / 4

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

<a id="canonical-75557d04091c39f89e77c19bb33e9a7d2d7435622d0eed41c3a0a42b9223ddfb"></a>

<a id="canonical-6c1af79f1c6de7c3748babc21297b5ae0c6084dcf520db998ac84650b18e1cd5"></a>

## mode property — items.get_spec.infra.bond_config / 9ad45cfdd81c / 5

Type: `"string"`. Computed.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

<a id="canonical-5a6f78e72c3afed598f290908b1131049c7e5f58e4d1ff720c91763bd1d07589"></a>

<a id="canonical-cd21effaf96bf261623be12f00f94d74face15f444728f2eefcc7ad0a948923e"></a>

## name property — items.get_spec.infra.bond_config / 9ad45cfdd81c / 6

Type: `"string"`. Computed.

Bond Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

<a id="canonical-671ef0e8e53ca90fb4a55626218b6ef79963f9d9b5b2917c819f172450a08024"></a>

## Next pages — items.get_spec.infra.bond_config / 9ad45cfdd81c / 7

- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-d51691d91b7ca9b4b7293926cc3acc1e1642be2a6cdd13cb31cf7ba91bd5b46b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cd801ee76992865d4f48774f1dab5aa526c3d7582f71d589aa04572c84da1f8"></a>

## items.get_spec.infra.hugepages — items.get_spec.infra.hugepages / 38f9b25d17d8 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- items.get_spec.infra.hugepages

<a id="canonical-ce6098f47d18fb32ec33e1c30b482eabd01f3a37941a6faec45748647e64b8c5"></a>

Type: `"list"`. Computed.

Hugepage settings for CE on K8s SMV2 site.

<a id="canonical-1afd94747ca5dd8416bf2915428ef58fddb559c2a846530dba44db95bb8036ec"></a>

## Direct properties — items.get_spec.infra.hugepages / 38f9b25d17d8 / 3

<a id="canonical-b138783d1787366d910a8109af32bc99f4415614b0092940dbc02c165816c392"></a>

<a id="canonical-5e760cadae9d66da243d5f2f081beda58c84e58f073da76c437b868b378eb554"></a>

## free property — items.get_spec.infra.hugepages / 38f9b25d17d8 / 4

Type: `"number"`. Computed.

Free Hugepages. Total number of free hugepages present.

<a id="canonical-d4d05b80e7c05e0142997acd98387d5330e1af4df80d78158230a4fc202e6a9c"></a>

<a id="canonical-0cfce6dfd7bb980358cb18790b993814535df5b802f34e758076d03eee344cbe"></a>

## page_size property — items.get_spec.infra.hugepages / 38f9b25d17d8 / 5

Type: `"number"`. Computed.

Hugepage Size. Size of each hugepage.

<a id="canonical-94d43b30e3d4b084b30d797cc69ddffc29391f0f9e31f03313b4f2656561eef6"></a>

<a id="canonical-0236b5b49f938d86a783c35e2c4d88674097e81f70840c274a855fed1269fc1d"></a>

## total property — items.get_spec.infra.hugepages / 38f9b25d17d8 / 6

Type: `"number"`. Computed.

Total Hugepages. Total number of hugepages present.

<a id="canonical-27a1f0211632c4020b3c7b27cc827cc3b54c237771aa7e36626e080b77ee7355"></a>

## Next pages — items.get_spec.infra.hugepages / 38f9b25d17d8 / 7

- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b714c901495159064bbe84354dab4eba0c14a9a17cf74db8a0f642145ad2db33"></a>

## items.get_spec.infra.hw_info — items.get_spec.infra.hw_info / b58a038d8ea0 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- items.get_spec.infra.hw_info

<a id="canonical-3fd098b622a08d9f02f9e3d59e287a50e2eb7bc56cb3db4bf9497ae88227a79d"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

<a id="canonical-29e766f9cd66e1ec433297d9c05389f566fbbc99074650c156eecb8d4f3c6d7d"></a>

## Direct properties — items.get_spec.infra.hw_info / b58a038d8ea0 / 3

- [bios](data-sources--site_registrations_by_state--reference--group-001.md#canonical-8f98297c295e39853da3c3242651f54bed53236d39165e91fdffa6f483103578): complete subsection reference.

- [board](data-sources--site_registrations_by_state--reference--group-001.md#canonical-32b4a91956bb518846bc20d2465fa33f83f8fd47a6829f1409dac81dd2e0aeab): complete subsection reference.

- [chassis](data-sources--site_registrations_by_state--reference--group-001.md#canonical-84e891063acf08ee20a26a976ee6e114a4ef4ae2b7a02c0eb31da97d56e8a69f): complete subsection reference.

- [cpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-16c2ce93e2cd95cb6322bb881afe53a4143980ff4c10d8c3d5d268560729e74c): complete subsection reference.

- [gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-830380e9fcd4eb5688736c4241e29631698aef4f70bfb26302e02c6796e6c5a5): complete subsection reference.

- [kernel](data-sources--site_registrations_by_state--reference--group-001.md#canonical-395ba6354d6d65f89479970d31ce181d744f62da69c8129ccfc936963d373bc2): complete subsection reference.

- [memory](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3826b2faa5a603de2b56eb79d0435918daa6f944cb43704efe1a73777254e770): complete subsection reference.

- [network](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c6c77dae6a67f37cab63e651e9575a95817d8e26203eed13ca45498c0428bad5): complete subsection reference.

<a id="canonical-93a846b791c2f70fc85ceb809a4d7a609a7eb9b895061954745e623cca8d6c5b"></a>

<a id="canonical-ecb515b1fc5017773db8bb314cd1f28161b5518c98321668790b930b42bcea61"></a>

## numa_nodes property — items.get_spec.infra.hw_info / b58a038d8ea0 / 4

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

- [os](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2913d07325b4db4dd52176468d7dc9cd2684d1e4378e715fef65089631f97cb2): complete subsection reference.

- [product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-964f5f5f0f09c64c6f66fd23d0ef61f7b99c18a977678707f00d3ecb8613a5df): complete subsection reference.

- [storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-998242dfec8a9f723c4ac75056b703a5f71d28cccb45b200c75c3002d9001757): complete subsection reference.

- [usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-46e212c80d9bf02a58eeb14bfe57b7e13531c7de86e5078c1dd7d1269c5df6e8): complete subsection reference.

<a id="canonical-43643e6b97a372de67b718eb8ac64ec9713a1f393ebecb2c84c7b373f8a46ab7"></a>

## Next pages — items.get_spec.infra.hw_info / b58a038d8ea0 / 5

- [items.get_spec.infra.hw_info.bios](data-sources--site_registrations_by_state--reference--group-001.md#canonical-8f98297c295e39853da3c3242651f54bed53236d39165e91fdffa6f483103578)
- [items.get_spec.infra.hw_info.board](data-sources--site_registrations_by_state--reference--group-001.md#canonical-32b4a91956bb518846bc20d2465fa33f83f8fd47a6829f1409dac81dd2e0aeab)
- [items.get_spec.infra.hw_info.chassis](data-sources--site_registrations_by_state--reference--group-001.md#canonical-84e891063acf08ee20a26a976ee6e114a4ef4ae2b7a02c0eb31da97d56e8a69f)
- [items.get_spec.infra.hw_info.cpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-16c2ce93e2cd95cb6322bb881afe53a4143980ff4c10d8c3d5d268560729e74c)
- [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-830380e9fcd4eb5688736c4241e29631698aef4f70bfb26302e02c6796e6c5a5)
- [items.get_spec.infra.hw_info.kernel](data-sources--site_registrations_by_state--reference--group-001.md#canonical-395ba6354d6d65f89479970d31ce181d744f62da69c8129ccfc936963d373bc2)
- [items.get_spec.infra.hw_info.memory](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3826b2faa5a603de2b56eb79d0435918daa6f944cb43704efe1a73777254e770)
- [items.get_spec.infra.hw_info.network](data-sources--site_registrations_by_state--reference--group-002.md#canonical-c6c77dae6a67f37cab63e651e9575a95817d8e26203eed13ca45498c0428bad5)
- [items.get_spec.infra.hw_info.os](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2913d07325b4db4dd52176468d7dc9cd2684d1e4378e715fef65089631f97cb2)
- [items.get_spec.infra.hw_info.product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-964f5f5f0f09c64c6f66fd23d0ef61f7b99c18a977678707f00d3ecb8613a5df)
- [items.get_spec.infra.hw_info.storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-998242dfec8a9f723c4ac75056b703a5f71d28cccb45b200c75c3002d9001757)
- [items.get_spec.infra.hw_info.usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-46e212c80d9bf02a58eeb14bfe57b7e13531c7de86e5078c1dd7d1269c5df6e8)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-8f98297c295e39853da3c3242651f54bed53236d39165e91fdffa6f483103578"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbf9b023702fcabc412c74a7d7d7cb310f560a34eb39480fe78047fdd930eb38"></a>

## items.get_spec.infra.hw_info.bios — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.bios

<a id="canonical-13734ea8ce5c2cbb0230b603c8bc70060c04c17df19073495587ba6a2247c237"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

<a id="canonical-8cabadfc1e290cb4e03991edc47ea474e2ceaf1291fa36cd7890fc2836aea519"></a>

## Direct properties — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 3

<a id="canonical-d35f0f8b9a9ff031c0fe582a7a9bef0a0d8d48a795ea9f59b305726487cc8538"></a>

<a id="canonical-91d9a19cc65c3a7bf36c3a869ec4126b15d8d2ecc3d76c02be95ce1837d5d42c"></a>

## date property — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 4

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

<a id="canonical-623ac9db8b8bcaa1d0d532b3025a5c6e8256a0d48bc254ae339ffe1a81f79643"></a>

<a id="canonical-b88fe9da7320f8f15f8835d684670b95247d2bc0e7de582cfaa7939e365e1173"></a>

## vendor property — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 5

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="canonical-7d858c9a05e14595c74b797d84895a1640d56d5b64cda01a26e144c7880a47fd"></a>

<a id="canonical-8d5c2591ea8d1400056bca5773f2dde2f1b24021b76e97427185b3d4b6f6bb97"></a>

## version property — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 6

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.

<a id="canonical-3ac2ba1c43bc51462c5c97cba721d783cdc78789eb66ad335fc2d59cf3402b3c"></a>

## Next pages — items.get_spec.infra.hw_info.bios / 3d7b3270ec87 / 7

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-32b4a91956bb518846bc20d2465fa33f83f8fd47a6829f1409dac81dd2e0aeab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-278bc2d7089bdbba8e1a05b9f56d0ea99dd24e00e8acd0467ae8e34c4fe7c053"></a>

## items.get_spec.infra.hw_info.board — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.board

<a id="canonical-27117624ce637c4df9a1628e66ef47671d3c21dc9c48775c004de214cb2ca48a"></a>

Type: `"single"`. Computed.

Board Details. Board information.

<a id="canonical-906993ca5e87f3a78749aeeb8b68cb2ca3d54e9a10798ee412cca94f3d092836"></a>

## Direct properties — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 3

<a id="canonical-957a973c5694473be9a5f0d10d816e6670a594724fccbaeb2df06f45a7151979"></a>

<a id="canonical-1fba3b1c9c187f716a314c95b2bb94bf4fab8af1db405cee242cb1f0f1c2ea51"></a>

## asset_tag property — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 4

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_asset\_tag.

<a id="canonical-0ec53a24327708fe432c648bd25cb45df5b0af22021e7578e865348886966f5d"></a>

<a id="canonical-7ab78be56a304456e9295065c48c8b7a14edc8e6314c396aa79df7611df1bdb5"></a>

## name property — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 5

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-661d631c465969b9d2e39ba3b807bcce177e743a7fa8cd6e0c08bc6636b99e3d"></a>

<a id="canonical-18f9db68b104955f194f1c2805ab5b9a3c00be4776001a30d3323c4ff8c9ac87"></a>

## serial property — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 6

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_serial.

<a id="canonical-17649f0014c4247c6774322c37f61cde3dedc0042683acfd3996c7b8ba61d48c"></a>

<a id="canonical-d1d4dbae6389189ca6e3d67c1d5b4105448380624358800f626c31cfd115e125"></a>

## vendor property — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 7

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_vendor.

<a id="canonical-68eed800d841eb3c9f0d436607ea47ff4b48cc170261fc0bad40ed60097556e8"></a>

<a id="canonical-7b27dd1519fab29ee4ee44f12cd01c2230d718234400186802362ffcb0cf727a"></a>

## version property — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 8

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_version.

<a id="canonical-bf53600191b425fe9e92c713a79884f6344ee2f8ef05fd56a51c4452d483c154"></a>

## Next pages — items.get_spec.infra.hw_info.board / 80f1e90eec99 / 9

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-84e891063acf08ee20a26a976ee6e114a4ef4ae2b7a02c0eb31da97d56e8a69f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70e8a7f98cf5aae70c9df6826863af4353b460d37252f92f0d421f57ae49d441"></a>

## items.get_spec.infra.hw_info.chassis — items.get_spec.infra.hw_info.chassis / a70869ac855b / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.chassis

<a id="canonical-fd42e849b0045a38641c92ebc8d912df4e9ecaa0cda20209afb63594e753c9e7"></a>

Type: `"single"`. Computed.

Chassis Details. Chassis information.

<a id="canonical-28d1430e03862cc414481de264bb0aa39018fc215d53c923542cbeb61a86f6fa"></a>

## Direct properties — items.get_spec.infra.hw_info.chassis / a70869ac855b / 3

<a id="canonical-a4c079208d1d38ddf68e892d2ff8c0a1b8d06ca5995f9e9b250ef2bef2ac98be"></a>

<a id="canonical-9e071ebfeeaaa3049a21230179b3b7c154a75c3768146f7db9ecb23dcf8590fe"></a>

## asset_tag property — items.get_spec.infra.hw_info.chassis / a70869ac855b / 4

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

<a id="canonical-9e69383a7ac2c41a2bfb586a1fa5aa5c710a9775b7d37cc62793b6b70e20ee79"></a>

<a id="canonical-42a8c1261d85083c7920dd25769de692e4649b5c6fabd2063b69b5466e398adb"></a>

## serial property — items.get_spec.infra.hw_info.chassis / a70869ac855b / 5

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_serial.

<a id="canonical-b9c1648834cadaf7d5ce8b42874a2f8f5dbafb1c525702c63401e90c709078da"></a>

<a id="canonical-1b7e6a44f907f49d667e1a974501a0fd2a4104c8e7bf7c3a13e7c0c6dc3ac497"></a>

## type property — items.get_spec.infra.hw_info.chassis / a70869ac855b / 6

Type: `"number"`. Computed.

Information from /sys/class/dmi/ID/chassis\_type.

<a id="canonical-6963c6848971c3ddfb2fef3a42de4c7d2aace387214ef2a2fafdb73a479359c3"></a>

<a id="canonical-93a04d9a6eaf874915652b64a2720095c8f0ddcbf45b1d8decc549ddf8082b79"></a>

## vendor property — items.get_spec.infra.hw_info.chassis / a70869ac855b / 7

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_vendor.

<a id="canonical-fc7c050ba8066ebdf70f43f66d01de24156dc84e256d21664d000a3dc72504f0"></a>

<a id="canonical-1badea94d963bcdcbcc26b10576b123a4242b1503131e303257329408ed199ee"></a>

## version property — items.get_spec.infra.hw_info.chassis / a70869ac855b / 8

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_version.

<a id="canonical-d1d06556c33fb7bcef075f1fa5afcfe538cdad2c8d253d17dabc5f33b3721451"></a>

## Next pages — items.get_spec.infra.hw_info.chassis / a70869ac855b / 9

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-16c2ce93e2cd95cb6322bb881afe53a4143980ff4c10d8c3d5d268560729e74c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aaf286685d44edcac9d6307f037db9d5dbb913f3fe393d15490d7c24da8ef24"></a>

## items.get_spec.infra.hw_info.cpu — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.cpu

<a id="canonical-34e658547379e5e865eb1993352d2c0b319e5883dd7532dd67f352e0e16f5655"></a>

Type: `"single"`. Computed.

CPU Information. CPU information.

<a id="canonical-71bd053d06160b464be4c4f050df5feedb1cb5ca2a095d5380dc9707c5397f5a"></a>

## Direct properties — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 3

<a id="canonical-2c2abad45c51def69818fd7739f14d3b5eccd8b50655f51fbe86326d92f5afe9"></a>

<a id="canonical-07bff02c7909c020c6edb88e916c39180cba4259ba86ddeca77863aefc955768"></a>

## cache property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 4

Type: `"number"`. Computed.

Cache. CPU cache size in KB.

<a id="canonical-a78e7ce6460cc8271d96f3a30171b2531bcfdea9731a8fd421f5c7b373f1c53f"></a>

<a id="canonical-9657698ef0981bd36cceb64dd860f7d80e2fe20ff29a820b5badc1d0410b7041"></a>

## cores property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 5

Type: `"number"`. Computed.

Cores. Number of physical CPU cores.

<a id="canonical-8d393c29024ce3eaec8819d20e0a63c456711dee8d8c043b54ecf6fa6b225d51"></a>

<a id="canonical-d8bf8c7b8dbe8a2e9456d4b2499bab2ed0f85bde51da3e2a97726bcd05ec8a4d"></a>

## cpus property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 6

Type: `"number"`. Computed.

CPUs. Number of physical CPUs.

<a id="canonical-efeea466dd2fb8b52248c4fb740fc1ffec8fe0203135caf24177a7282f338885"></a>

<a id="canonical-da7757c65e30b8821a4abbbdbd7e36620ada40f392d2cc6e9e6f87536433a819"></a>

## model property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 7

Type: `"string"`. Computed.

Model. CPU model

<a id="canonical-51f31c03954bab760e927ea99c447b9bcb563a6cb1dd1de429ae2963e7675c8e"></a>

<a id="canonical-d43cd65720bbc3c13e1241c5d618decee979977d9ac9340ae38d6c5aa880a07c"></a>

## speed property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 8

Type: `"number"`. Computed.

Speed. CPU clock rate in MHz.

<a id="canonical-b2af4d2923aca499e636d8e672516cfa6908e2048c4e23b6b309e06dc69a67cf"></a>

<a id="canonical-24bbd39c4b1ad2933cc424331611d915f047e27d136fe5f5520ca4af99f048b1"></a>

## threads property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 9

Type: `"number"`. Computed.

Threads. Number of logical (HT) CPU cores.

<a id="canonical-4ce8263e927569b55ca788b0145f3464dfa7d6898ec30b17dc6c7cfeeac00bbc"></a>

<a id="canonical-e08d0d940bde8b92beea1cc7484841f89ad028ad992a57f0e6d972ea41f4b40e"></a>

## vendor property — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 10

Type: `"string"`. Computed.

Vendor. CPU vendor.

<a id="canonical-320ca1d00e27f1f5b27222df29599224c535e16afb0bd975eb41b769b042737b"></a>

## Next pages — items.get_spec.infra.hw_info.cpu / e1a4bc2ccfe1 / 11

- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-830380e9fcd4eb5688736c4241e29631698aef4f70bfb26302e02c6796e6c5a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3b3d8c9a19c82fa030f432a67f4c6ba0001077efa926355bcf355069ba28d4a"></a>

## items.get_spec.infra.hw_info.gpu — items.get_spec.infra.hw_info.gpu / 236bddaf73f2 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.gpu

<a id="canonical-4095f0e56f14dc36d885814341619b4579a0bb3388abb3e8107f9e196fbf6e32"></a>

Type: `"single"`. Computed.

GPU. GPU information on server.

<a id="canonical-58fc37a42363b4cfa6682c7842e6a8fb69ce8846db8b9ef6f9a5c2e25a5f8775"></a>

## Direct properties — items.get_spec.infra.hw_info.gpu / 236bddaf73f2 / 3

<a id="canonical-5062039477ec16c3ffc6e14a166f654bf342c15cf170b84aced92c612954e1fe"></a>

<a id="canonical-56916d324200c3d266b50e3cfbb2a1e2405a114c97a92cce20cd4b52995949ce"></a>

## cuda_version property — items.get_spec.infra.hw_info.gpu / 236bddaf73f2 / 4

Type: `"string"`. Computed.

Cuda Version. GPU Cuda Version.

<a id="canonical-5452371ed3a84568faa3b9229906b292d10d3b4ef3696cbcf55f62dfeb3e3fc8"></a>

<a id="canonical-c37ef3f60829d1c046ae1daaa6a849cc6fd0138aa36358cd3494b579940e81b7"></a>

## driver_version property — items.get_spec.infra.hw_info.gpu / 236bddaf73f2 / 5

Type: `"string"`. Computed.

Driver Version. GPU Driver Version.

- [gpu_device](data-sources--site_registrations_by_state--reference--group-001.md#canonical-6fdab37ff57867dc3c9f99588080863949efb9e269db8e0e0c3935b97539605f): complete subsection reference.

<a id="canonical-72c1312cf92f02c8c2f4810febdf40d69b6239506f418dc020149c5f5e941b0a"></a>

## Next pages — items.get_spec.infra.hw_info.gpu / 236bddaf73f2 / 6

- [items.get_spec.infra.hw_info.gpu.gpu_device](data-sources--site_registrations_by_state--reference--group-001.md#canonical-6fdab37ff57867dc3c9f99588080863949efb9e269db8e0e0c3935b97539605f)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-6fdab37ff57867dc3c9f99588080863949efb9e269db8e0e0c3935b97539605f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98e635916fa4cddf7597a520069671111752afc8bd555ea2f6955c7013e53077"></a>

## items.get_spec.infra.hw_info.gpu.gpu_device — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-830380e9fcd4eb5688736c4241e29631698aef4f70bfb26302e02c6796e6c5a5)
- items.get_spec.infra.hw_info.gpu.gpu_device

<a id="canonical-dc91d62c6aa9d613ee5d448dc697beb6afb7ae5aa5fc6fdea141c622bad86ed1"></a>

Type: `"list"`. Computed.

GPU devices. List of GPU devices in server.

<a id="canonical-ecd9eed44244bc62aa1e8f227c1b947f71ca0eca3307301fab49bafd53da2b2f"></a>

## Direct properties — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 3

<a id="canonical-c669513f3d7a8c377a4fa4c5b03a691d449a954d024b9be90a1d8c3058a908a9"></a>

<a id="canonical-30af8c078845e11475882ffa55d1c253df0ee2e2ee4a3e9b6ca0c49335dcd0e4"></a>

## id property — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 4

Type: `"string"`. Computed.

GPU ID. GPU ID

<a id="canonical-6ec7f61d5fd25b082caf72f4b7096fa3d844a97d312eb7652bee8e24a39847e8"></a>

<a id="canonical-581ba0d5374e57962a06a80310618068c1b64c278ed96830a7ee59b35bb38dc6"></a>

## processes property — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 5

Type: `"string"`. Computed.

Processes. GPU Processes.

<a id="canonical-12948f64048fe94385220a0c2f0c6b9b785cd153938641ae3c5398d813268ca9"></a>

<a id="canonical-374071013b6dee294aa5c39ef8473d0c8228b0ec9b7e219ff65f0df4b7fea3a2"></a>

## product_name property — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 6

Type: `"string"`. Computed.

Product Name. GPU Product Name.

<a id="canonical-6f427a6b1d46cca0a874c0ea5230eba59e9617983e78d4b4935e15b955b149b2"></a>

## Next pages — items.get_spec.infra.hw_info.gpu.gpu_device / 75c3524f0872 / 7

- [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-830380e9fcd4eb5688736c4241e29631698aef4f70bfb26302e02c6796e6c5a5)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-395ba6354d6d65f89479970d31ce181d744f62da69c8129ccfc936963d373bc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e87508406d862ced104690737315ffae16fbd064dd54c895b2cb64e4b0a785c"></a>

## items.get_spec.infra.hw_info.kernel — items.get_spec.infra.hw_info.kernel / 56fbf8b62587 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-b956fafb537777957e17f6d9aca6d5c5f6f06238257ab65292fa01ac07c30d7a)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2fd0b74c0e8ebcc9a99d16882715d8d2c2874ebea34d6563b71de4372fe2a639)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2903eac203ab1878382bf5bd6a29830083c946146d47792d041065d5722feeee)
- items.get_spec.infra.hw_info.kernel

<a id="canonical-ca92e3dfcab064e0d7f76a0f5d3bd3157c6fc41c6d31ab08609ffa790c82f622"></a>

Type: `"single"`. Computed.

Kernel. Kernel information.
