---
page_title: "xcsh_nginx_server reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_server reference."
---

# xcsh_nginx_server reference

<a id="canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-064f5d21da8e858dfe1df82f39cfc3a1e8fe012debb221afabf129c78325fec5"></a>

## Property reference — Property reference / b6ff9f6f9d5c / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- Property reference

<a id="canonical-b3d34856c4edfe6dd274022549731d16753268ff97260ce081249693f15dfe30"></a>

## Direct properties — Property reference / b6ff9f6f9d5c / 3

<a id="canonical-fa408ceeb2cc57418954519dc2b7766f049cc7131d1853c7c53612fea7c12100"></a>

<a id="canonical-f846f5a24ccd52aa11254a643aa6be9d8847fe99069ee0601150898851762b52"></a>

## annotations property — Property reference / b6ff9f6f9d5c / 4

Type: `["map", "string"]`. Computed.

Annotations.

- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed): complete subsection reference.

<a id="canonical-bf71598a7b826ec78ac99b53c6c128a110055583096f4a99dfd706cf14315d2a"></a>

<a id="canonical-cf3073b17f472237a423124f5d18772b71b1f1e5261440f9d74e027f3607d0e0"></a>

## description property — Property reference / b6ff9f6f9d5c / 5

Type: `"string"`. Computed.

Description.

<a id="canonical-4f2938089fbac9fc76bb70012e5211eef9271d540fe87e4183cdb92666316240"></a>

<a id="canonical-719b503fd267a042a2fb168e9dcdf0ad8ca82872b3b47fccc9b63413d6237e5d"></a>

## id property — Property reference / b6ff9f6f9d5c / 6

Type: `"string"`. Computed.

Unique identifier.

<a id="canonical-a6640aca4726cf498b7315ab9b51def6aafc4b53a3d36cead17a761ef2a0b43d"></a>

<a id="canonical-23b83646b9084ad15c3d5fcacb02d32e97c4a0abee1393ac03f200d069930a2e"></a>

## labels property — Property reference / b6ff9f6f9d5c / 7

Type: `["map", "string"]`. Computed.

Labels.

<a id="canonical-f9349e4daf099371e3c193cbc1e72daa7da9bfcdf1a3b981a94994080b2708e9"></a>

<a id="canonical-f6510ec9478043bcb065f2f2976e4c709fa34da96d4a31357768fe69d9c4e231"></a>

## name property — Property reference / b6ff9f6f9d5c / 8

Type: `"string"`. Required.

Name of the NginxServer to look up.

<a id="canonical-f73e7590060bc36820f8d831b7441dcc1c229e52f1133af6b2e94c57f54f189b"></a>

<a id="canonical-0feddae92e09b20b2ef453f0d80429374d359c4f329c47cbd5281b8ea2af19d9"></a>

## namespace property — Property reference / b6ff9f6f9d5c / 9

Type: `"string"`. Required.

Namespace of the NginxServer.

- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d): complete subsection reference.

<a id="canonical-cad2bb602ee4aa9a3dcd1f87f8ac927a74bdc523801703a802890d65ce08527a"></a>

## All schema paths — Property reference / b6ff9f6f9d5c / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_server--reference--group-001.md#canonical-fa408ceeb2cc57418954519dc2b7766f049cc7131d1853c7c53612fea7c12100) |
| `dataplane_ref` | [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-59fd5c9357a918df7bb93c6909bef8ffb15834de4b33f91ad0c650538ec614c9) |
| `dataplane_ref.nginx_csg` | [dataplane_ref.nginx_csg](data-sources--nginx_server--reference--group-001.md#canonical-6f104a15e1dd52ec5eff013bad60bf6443353dab118004eeaf6447b711e706e9) |
| `dataplane_ref.nginx_csg.name` | [dataplane_ref.nginx_csg.name](data-sources--nginx_server--reference--group-001.md#canonical-837b052cf749efb45b7ccd0d240a3245bfa9a554ec5fdbd5cd00cb87c0d7d873) |
| `dataplane_ref.nginx_csg.namespace` | [dataplane_ref.nginx_csg.namespace](data-sources--nginx_server--reference--group-001.md#canonical-7cb1da10a40c10ef40c464947eff87a63836bbd97365d2052dee981d83cdce30) |
| `dataplane_ref.nginx_csg.tenant` | [dataplane_ref.nginx_csg.tenant](data-sources--nginx_server--reference--group-001.md#canonical-bf4605bcb15ff8288ad93c6291551272bf9ee00a099c77c7797fcf19f419075f) |
| `dataplane_ref.nginx_instance` | [dataplane_ref.nginx_instance](data-sources--nginx_server--reference--group-001.md#canonical-be403bb879ec65960f5441ca5081dcfcfe77db3bd0cf68c7907ca8f1138b4143) |
| `dataplane_ref.nginx_instance.name` | [dataplane_ref.nginx_instance.name](data-sources--nginx_server--reference--group-001.md#canonical-1a870ebb6a343572f50618d685448e6ecc2434812375ed92276d0a16ee2367ad) |
| `dataplane_ref.nginx_instance.namespace` | [dataplane_ref.nginx_instance.namespace](data-sources--nginx_server--reference--group-001.md#canonical-d62af18b9db0cef12b76edb3c468f3cf9e6a5ab14a69082b142fded8dda2b9a3) |
| `dataplane_ref.nginx_instance.tenant` | [dataplane_ref.nginx_instance.tenant](data-sources--nginx_server--reference--group-001.md#canonical-0566541b6d1856e3b477d3aa8c3d16e2c30d9b53178b14a10abab8228ffdb4c0) |
| `description` | [description](data-sources--nginx_server--reference--group-001.md#canonical-bf71598a7b826ec78ac99b53c6c128a110055583096f4a99dfd706cf14315d2a) |
| `id` | [id](data-sources--nginx_server--reference--group-001.md#canonical-4f2938089fbac9fc76bb70012e5211eef9271d540fe87e4183cdb92666316240) |
| `labels` | [labels](data-sources--nginx_server--reference--group-001.md#canonical-a6640aca4726cf498b7315ab9b51def6aafc4b53a3d36cead17a761ef2a0b43d) |
| `name` | [name](data-sources--nginx_server--reference--group-001.md#canonical-f9349e4daf099371e3c193cbc1e72daa7da9bfcdf1a3b981a94994080b2708e9) |
| `namespace` | [namespace](data-sources--nginx_server--reference--group-001.md#canonical-f73e7590060bc36820f8d831b7441dcc1c229e52f1133af6b2e94c57f54f189b) |
| `server_spec` | [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-652bca6236c3e6111a75ae4c72eace0bc834992fa32e51661fb9fb9bc225cef6) |
| `server_spec.api_discovery_spec` | [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-5bbc463c344595ec971aec203bb05bbd9f5313f05cc78a17450b3c263bdebd8c) |
| `server_spec.api_discovery_spec.disabled` | [server_spec.api_discovery_spec.disabled](data-sources--nginx_server--reference--group-001.md#canonical-7667d5a9296b270a332c7489860dfeb248213ca787b43384dff421924d64a628) |
| `server_spec.api_discovery_spec.enabled` | [server_spec.api_discovery_spec.enabled](data-sources--nginx_server--reference--group-001.md#canonical-aac571d1a4c208d7c956485186a34cfa0c81da3187da80ea1186ac726ce46b3c) |
| `server_spec.domains` | [server_spec.domains](data-sources--nginx_server--reference--group-001.md#canonical-83d4a09d77ff9c0d9728ba7ec89f636f247dc1d77472897fdb552a8d87ab2285) |
| `server_spec.locations` | [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a0f3ea1105ba1add42285cc1c517e79531cb2d12f4d246b15888a732e22e6b91) |
| `server_spec.locations.api_discovery_spec` | [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d9941d0b8e4b1765d9c647e0cb20e40f26fdf8d1e07f84a874655ddb55abd9e3) |
| `server_spec.locations.api_discovery_spec.disabled` | [server_spec.locations.api_discovery_spec.disabled](data-sources--nginx_server--reference--group-001.md#canonical-44ff3c4bdb34817a8a5b36b5366fe7c25adf743c727337f3706b102d5ed11b09) |
| `server_spec.locations.api_discovery_spec.enabled` | [server_spec.locations.api_discovery_spec.enabled](data-sources--nginx_server--reference--group-001.md#canonical-38432fcfd8878e15110fdb53067f766737af9e7d96855c6846cb846803dc1585) |
| `server_spec.locations.definition` | [server_spec.locations.definition](data-sources--nginx_server--reference--group-001.md#canonical-afbf823d98ecc27eeaa5b4340103e9b38e6ed32062c4245aae6e76a9633fe77a) |
| `server_spec.locations.name` | [server_spec.locations.name](data-sources--nginx_server--reference--group-001.md#canonical-ea64337daf7be396a78c8f1de4beadf002934dcda62bd6d987b60ec5053e838a) |
| `server_spec.locations.waf_spec` | [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-70d6ae10e3ef2d93bbd5a7520b4ffeb1754b9a122ed07b102ece0347ee83dafb) |
| `server_spec.locations.waf_spec.blocking_waf_mode` | [server_spec.locations.waf_spec.blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-06fe7a573eebd3f04a4c9777365670c2897e739b4a50835ad6b060b2e4e5861e) |
| `server_spec.locations.waf_spec.distributed_cloud_policy_management` | [server_spec.locations.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-431f50eef6ad14fe2400ce12b3941f4c88f6c5c252d67f2f36ef151c0c89f353) |
| `server_spec.locations.waf_spec.monitoring_waf_mode` | [server_spec.locations.waf_spec.monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-51820a84a88c52271591c80376d5ab82caae1b60b10905171efca2b059156940) |
| `server_spec.locations.waf_spec.nginx_policy_management` | [server_spec.locations.waf_spec.nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-e82b778ebce2bc4f0317274434627eb82b5b9f97b34959b2c9bd49a52da29a0d) |
| `server_spec.locations.waf_spec.none_waf_mode` | [server_spec.locations.waf_spec.none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-38379bb9be093ba27cbab3c6f872c97b7e1f271131d76252b8e807c31a629c6f) |
| `server_spec.locations.waf_spec.policy_file_name` | [server_spec.locations.waf_spec.policy_file_name](data-sources--nginx_server--reference--group-001.md#canonical-043b0138fb2e67b9e21b5049e63e294c3e65d9e24e05902005edfc45038f06a0) |
| `server_spec.locations.waf_spec.policy_name` | [server_spec.locations.waf_spec.policy_name](data-sources--nginx_server--reference--group-001.md#canonical-ae1b0218f260c34a0680e19625c5981dcf0464b3f542893325eaae54025fa768) |
| `server_spec.locations.waf_spec.security_log_enabled` | [server_spec.locations.waf_spec.security_log_enabled](data-sources--nginx_server--reference--group-001.md#canonical-e0a52aa1db5ee3a971a6bf2df2e10b6fc9bf9d553cac94111cd44cb603c68592) |
| `server_spec.locations.waf_spec.security_log_file_names` | [server_spec.locations.waf_spec.security_log_file_names](data-sources--nginx_server--reference--group-001.md#canonical-4a6213297c2b79bdff9964caeb458fe7ea2d09d0b43c6d769588e5a61f2343e3) |
| `server_spec.nginx_one_object_id` | [server_spec.nginx_one_object_id](data-sources--nginx_server--reference--group-001.md#canonical-0bbb0a4f5707e02b55e0bea70a060640efd68bd205f43a0d957ad59b4879b203) |
| `server_spec.nginx_one_object_name` | [server_spec.nginx_one_object_name](data-sources--nginx_server--reference--group-001.md#canonical-6c031e0be051c03822e7fc04f73798c397360d64dba1f15029c13d1b7798bcfc) |
| `server_spec.port` | [server_spec.port](data-sources--nginx_server--reference--group-001.md#canonical-ed67278fb87d270d55df3bccbd7fac85b90ef9747140217051f67ab8de05a08e) |
| `server_spec.server_name` | [server_spec.server_name](data-sources--nginx_server--reference--group-001.md#canonical-ae8aadbe812c15401a50620482e2f31b154d5ac8a0673ddafe6435f932fd658c) |
| `server_spec.total_routes` | [server_spec.total_routes](data-sources--nginx_server--reference--group-001.md#canonical-f0c53299950268c58e522895cc8603a50bd6e1c1048f40049dd6d6fcda0d8e36) |
| `server_spec.waf_spec` | [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-3c19225c51d0b6daccb2e377c4676fefa0f44828f84c5dea951465776b64e529) |
| `server_spec.waf_spec.blocking_waf_mode` | [server_spec.waf_spec.blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-94b2189b032bdd9ab3615f8e5affa3f7df0c5a0483d1176a9c8b184b6b1b5aa8) |
| `server_spec.waf_spec.distributed_cloud_policy_management` | [server_spec.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-ae6885d05917e503ca34899083e8b88132d6e74590fd2bbe58d8362c3cc21505) |
| `server_spec.waf_spec.monitoring_waf_mode` | [server_spec.waf_spec.monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-65ca80fb697f9ae4370154d1b2db63bcc2544a92e664f39a717fbad105716e42) |
| `server_spec.waf_spec.nginx_policy_management` | [server_spec.waf_spec.nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-f9d99c3a120d9c671b692df6ee12e9e8327c826593638e5a8561208b9e9bda51) |
| `server_spec.waf_spec.none_waf_mode` | [server_spec.waf_spec.none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-41b71fdcaa4951d525b1a8328f988c3c45b9a5f324cc8721cfb808d6ed9cd1ff) |
| `server_spec.waf_spec.policy_file_name` | [server_spec.waf_spec.policy_file_name](data-sources--nginx_server--reference--group-001.md#canonical-ada9cf03ce47f71d52ff16a0c286207c21db9a98c36f69402b7d8cffb537e4f1) |
| `server_spec.waf_spec.policy_name` | [server_spec.waf_spec.policy_name](data-sources--nginx_server--reference--group-001.md#canonical-43646c1a88d2cec12f473fcccd439380cfb9d7b3b92ba9d5f68cf5ba7e958b16) |
| `server_spec.waf_spec.security_log_enabled` | [server_spec.waf_spec.security_log_enabled](data-sources--nginx_server--reference--group-001.md#canonical-ae0143e666d7d17ed3e32bf10467c514fb99d47c9d575ff2e9dbeba8566b99c6) |
| `server_spec.waf_spec.security_log_file_names` | [server_spec.waf_spec.security_log_file_names](data-sources--nginx_server--reference--group-001.md#canonical-2c0762836e69a962e3eb066891b003d5ec003fc372cd0133f33d341d0511d48e) |

<a id="canonical-f7e3474b74875ac9822603614496c91380327fe50c5cc9860cf44f68bebbbf0a"></a>

## Next pages — Property reference / b6ff9f6f9d5c / 11

- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59309565fb2c81966ff92c5c5a12b9fbcf94d31c35f5b8e0300baa30a2556d7e"></a>

## dataplane_ref — dataplane_ref / 22933fad279a / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- dataplane_ref

<a id="canonical-59fd5c9357a918df7bb93c6909bef8ffb15834de4b33f91ad0c650538ec614c9"></a>

Type: `"single"`. Computed.

DataplaneReference.

<a id="canonical-aa9f17cecc0cc513ba7fda41df3337ad33b1414caa44da3a9de65d0a7b646c9d"></a>

## Direct properties — dataplane_ref / 22933fad279a / 3

- [nginx_csg](data-sources--nginx_server--reference--group-001.md#canonical-7ad685117718496c5d8f5d431940e2312aa5ae9809177a898ba4d5fa7d65b566): complete subsection reference.

- [nginx_instance](data-sources--nginx_server--reference--group-001.md#canonical-26d3376c10655c1d740402e4d5c940da80e238bcce99c16f81c04fefbafe0517): complete subsection reference.

<a id="canonical-748c12fecd1a6407521f23e5f7f04db2a1ee12abf4663d0162bdb530bb5a8c4b"></a>

## Next pages — dataplane_ref / 22933fad279a / 4

- [dataplane_ref.nginx_csg](data-sources--nginx_server--reference--group-001.md#canonical-7ad685117718496c5d8f5d431940e2312aa5ae9809177a898ba4d5fa7d65b566)
- [dataplane_ref.nginx_instance](data-sources--nginx_server--reference--group-001.md#canonical-26d3376c10655c1d740402e4d5c940da80e238bcce99c16f81c04fefbafe0517)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-7ad685117718496c5d8f5d431940e2312aa5ae9809177a898ba4d5fa7d65b566"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9e328d41559e1226816aba0fdfe09444c1d2bc7447ddedfaf33b9fb6222c92f"></a>

## dataplane_ref.nginx_csg — dataplane_ref.nginx_csg / e6e7cfae45d1 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed)
- dataplane_ref.nginx_csg

<a id="canonical-6f104a15e1dd52ec5eff013bad60bf6443353dab118004eeaf6447b711e706e9"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

<a id="canonical-d94adb2a8775edec70c56cac09f66c801153b16b923ce3e82c7b35171e9c101c"></a>

## Direct properties — dataplane_ref.nginx_csg / e6e7cfae45d1 / 3

<a id="canonical-837b052cf749efb45b7ccd0d240a3245bfa9a554ec5fdbd5cd00cb87c0d7d873"></a>

<a id="canonical-c71f21f65a66cb044f66bd64c41d425432e99c17d16f814bbc230842b390f55d"></a>

## name property — dataplane_ref.nginx_csg / e6e7cfae45d1 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-7cb1da10a40c10ef40c464947eff87a63836bbd97365d2052dee981d83cdce30"></a>

<a id="canonical-2ef2a773813e2fd5f28c553bb20e44165400646f4d750cdec0c9a4ab9703b7f3"></a>

## namespace property — dataplane_ref.nginx_csg / e6e7cfae45d1 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="canonical-bf4605bcb15ff8288ad93c6291551272bf9ee00a099c77c7797fcf19f419075f"></a>

<a id="canonical-b2af8080c535b6f296c0a56c71c07764e39e39627a2c628df2c1be9dd50db1c5"></a>

## tenant property — dataplane_ref.nginx_csg / e6e7cfae45d1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-98fef430ce07f63e3996892a6f61048be7e39fe817e1130c1b231e1d45893cc3"></a>

## Next pages — dataplane_ref.nginx_csg / e6e7cfae45d1 / 7

- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-26d3376c10655c1d740402e4d5c940da80e238bcce99c16f81c04fefbafe0517"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a0b91505b129e9aa34128371f036cc0cd5841acfed390c09acd28ccca1e2d25"></a>

## dataplane_ref.nginx_instance — dataplane_ref.nginx_instance / 503fc219dc2f / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed)
- dataplane_ref.nginx_instance

<a id="canonical-be403bb879ec65960f5441ca5081dcfcfe77db3bd0cf68c7907ca8f1138b4143"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

<a id="canonical-10a0778f91acf180b034c184149df1ddcd21f92482d47d6ed48cd0aaedd19254"></a>

## Direct properties — dataplane_ref.nginx_instance / 503fc219dc2f / 3

<a id="canonical-1a870ebb6a343572f50618d685448e6ecc2434812375ed92276d0a16ee2367ad"></a>

<a id="canonical-95cd38a559a7ed4c309b0618a44f8b409f64372eb38034b3a7481f1e2dc24b25"></a>

## name property — dataplane_ref.nginx_instance / 503fc219dc2f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-d62af18b9db0cef12b76edb3c468f3cf9e6a5ab14a69082b142fded8dda2b9a3"></a>

<a id="canonical-51a0e69c78a95944d893268102cb06f55148be0f229164e485a9d27a6c2c9afa"></a>

## namespace property — dataplane_ref.nginx_instance / 503fc219dc2f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="canonical-0566541b6d1856e3b477d3aa8c3d16e2c30d9b53178b14a10abab8228ffdb4c0"></a>

<a id="canonical-e220d6790900e2417a810533fb4653d60b6ef6b67c74f00cb0b5e7d26c738a17"></a>

## tenant property — dataplane_ref.nginx_instance / 503fc219dc2f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-758db491745f4b1dd53a47dcb122e9a43e2cdaf1784da330810f52a2fb10d56d"></a>

## Next pages — dataplane_ref.nginx_instance / 503fc219dc2f / 7

- [dataplane_ref](data-sources--nginx_server--reference--group-001.md#canonical-28c982205935f107691e31cef2e9ad66f30ccc98e799f7436fe49c16ffc73bed)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c354ffff3f100cfc010d5fa4dd74d3b6e77cc57a76a64f8c4b34e83c85b2f388"></a>

## server_spec — server_spec / b2d226bffa1a / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- server_spec

<a id="canonical-652bca6236c3e6111a75ae4c72eace0bc834992fa32e51661fb9fb9bc225cef6"></a>

Type: `"single"`. Computed.

Configuration for server\_spec.

<a id="canonical-bab721679b38c8c35b55231498c49b15a4aafec269b5e5b7bf071c66b8209b5c"></a>

## Direct properties — server_spec / b2d226bffa1a / 3

- [api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb): complete subsection reference.

<a id="canonical-83d4a09d77ff9c0d9728ba7ec89f636f247dc1d77472897fdb552a8d87ab2285"></a>

<a id="canonical-4f0c6ab695dce686198230c2be0073d7540b35d99b082bb53bba8a1ee3af7121"></a>

## domains property — server_spec / b2d226bffa1a / 4

Type: `["list", "string"]`. Computed.

Server name list specified as $\{server\_name\} in NGINX config. If no value is specified
corresponding to this variable, 'default' is used Reference:
https&#58;//nginx.org/en/docs/HTTP/ngx\_http\_core\_module.html\#server.

- [locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863): complete subsection reference.

<a id="canonical-0bbb0a4f5707e02b55e0bea70a060640efd68bd205f43a0d957ad59b4879b203"></a>

<a id="canonical-e0cf81f990ada9bfa51d034439c9277ed38d32bd8808880879a072ecd7ee8934"></a>

## nginx_one_object_id property — server_spec / b2d226bffa1a / 5

Type: `"string"`. Computed.

Signifies the uniqueness identifier for NGINX One representation of this NGINX server.

<a id="canonical-6c031e0be051c03822e7fc04f73798c397360d64dba1f15029c13d1b7798bcfc"></a>

<a id="canonical-6178f405ffae1df02683c2cdf645b6d525cf4817b6b79dd490d0ce33156c5174"></a>

## nginx_one_object_name property — server_spec / b2d226bffa1a / 6

Type: `"string"`. Computed.

Hostname value set for Instance or Name for a Config Sync Group in NGINX One.

<a id="canonical-ed67278fb87d270d55df3bccbd7fac85b90ef9747140217051f67ab8de05a08e"></a>

<a id="canonical-ef550660d2b1d0f6a54025485c456d0206fe1c769937c729462210e10554ef02"></a>

## port property — server_spec / b2d226bffa1a / 7

Type: `"number"`. Computed.

Signifies the port configured for the NGINX server.

<a id="canonical-ae8aadbe812c15401a50620482e2f31b154d5ac8a0673ddafe6435f932fd658c"></a>

<a id="canonical-ff37498f10e69341a7738077179fb60f8fbc676664282b6dcdb25d971d039a6f"></a>

## server_name property — server_spec / b2d226bffa1a / 8

Type: `"string"`. Computed.

Signifies the combination of first element in domains array and the port configured for the NGINX
server.

<a id="canonical-f0c53299950268c58e522895cc8603a50bd6e1c1048f40049dd6d6fcda0d8e36"></a>

<a id="canonical-b3828349aaadab4ac4f6669eb633dc241233a869009a2c6ef2bbdc8ba9d5608d"></a>

## total_routes property — server_spec / b2d226bffa1a / 9

Type: `"number"`. Computed.

Total locations configured in the NGINX Server.

- [waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee): complete subsection reference.

<a id="canonical-670cdb0aafe8d8bab2370ecaf8a2784de20f513ad8ee5539ab6613d7d11ad3e0"></a>

## Next pages — server_spec / b2d226bffa1a / 10

- [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c03e35bee86427693d7cb1ca04a019bd23b818ea8820d2a78d187579a095693"></a>

## server_spec.api_discovery_spec — server_spec.api_discovery_spec / cb928c48c2be / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- server_spec.api_discovery_spec

<a id="canonical-5bbc463c344595ec971aec203bb05bbd9f5313f05cc78a17450b3c263bdebd8c"></a>

Type: `"single"`. Computed.

Configuration for api\_discovery\_spec.

<a id="canonical-18dd0c43e42a6891bbae5650947adc4f0ae578ec6dbac6a4eecbdd9f567696f4"></a>

## Direct properties — server_spec.api_discovery_spec / cb928c48c2be / 3

- [disabled](data-sources--nginx_server--reference--group-001.md#canonical-da17a06343102eef8a93e982437abb4297a78df1c790e916494e08658aa8800e): complete subsection reference.

- [enabled](data-sources--nginx_server--reference--group-001.md#canonical-e39013bff8106c1ca76cda8c96ee63f1baa3361f82090cd5860e5abd0c7e8ed6): complete subsection reference.

<a id="canonical-67f3fe7a5e6627646cc17e590860677b01f8efd72986fd4726511146aef10426"></a>

## Next pages — server_spec.api_discovery_spec / cb928c48c2be / 4

- [server_spec.api_discovery_spec.disabled](data-sources--nginx_server--reference--group-001.md#canonical-da17a06343102eef8a93e982437abb4297a78df1c790e916494e08658aa8800e)
- [server_spec.api_discovery_spec.enabled](data-sources--nginx_server--reference--group-001.md#canonical-e39013bff8106c1ca76cda8c96ee63f1baa3361f82090cd5860e5abd0c7e8ed6)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-da17a06343102eef8a93e982437abb4297a78df1c790e916494e08658aa8800e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e30cfba1b414a28cc93149b59330b38e2623460cc4b6db7172a6ff458f24d4e9"></a>

## server_spec.api_discovery_spec.disabled — server_spec.api_discovery_spec.disabled / 91ba4e271a26 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb)
- server_spec.api_discovery_spec.disabled

<a id="canonical-7667d5a9296b270a332c7489860dfeb248213ca787b43384dff421924d64a628"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-4905a1cebb7fcf6476ca94e15ba5d25a5993869807e3e025394a7f1a66965b95"></a>

## Direct properties — server_spec.api_discovery_spec.disabled / 91ba4e271a26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4fea4e3b10208832a17b3ffb48beb8f5ec204ede9aa3602942353e13d3d0e7bd"></a>

## Next pages — server_spec.api_discovery_spec.disabled / 91ba4e271a26 / 4

- [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-e39013bff8106c1ca76cda8c96ee63f1baa3361f82090cd5860e5abd0c7e8ed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c1ca5e2e522adc93a0380ba730ceae6a05c6916ae0b02e838fff29645543c94"></a>

## server_spec.api_discovery_spec.enabled — server_spec.api_discovery_spec.enabled / 21c18d600bc6 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb)
- server_spec.api_discovery_spec.enabled

<a id="canonical-aac571d1a4c208d7c956485186a34cfa0c81da3187da80ea1186ac726ce46b3c"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-2497964e7f002175c685f1308871233e48516a5b66a09edcb2b272023fdfdf76"></a>

## Direct properties — server_spec.api_discovery_spec.enabled / 21c18d600bc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03430fdec9439ed2c77b16b64f7da37df3e3bf9d0a52876dec7fb1fa7e839108"></a>

## Next pages — server_spec.api_discovery_spec.enabled / 21c18d600bc6 / 4

- [server_spec.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-d2514125884ff776c605f4533037937287c00a1bdf2afba64820b1a0bbbbc8eb)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23ac484ad1e6464013ffd0877d40095b1a4dd80940a092879152a89580611daf"></a>

## server_spec.locations — server_spec.locations / 9e12e034dbc2 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- server_spec.locations

<a id="canonical-a0f3ea1105ba1add42285cc1c517e79531cb2d12f4d246b15888a732e22e6b91"></a>

Type: `"list"`. Computed.

Configuration of the set of locations corresponding to this server.

<a id="canonical-73eca2ad6516660031462708f4de6024efc7ec8e527236818f26e5c433c836d9"></a>

## Direct properties — server_spec.locations / 9e12e034dbc2 / 3

- [api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28): complete subsection reference.

<a id="canonical-afbf823d98ecc27eeaa5b4340103e9b38e6ed32062c4245aae6e76a9633fe77a"></a>

<a id="canonical-17e96cb00c9cc4c307ae06b6ce77e9a053a4bf4e57f9ce0264b2048e13245691"></a>

## definition property — server_spec.locations / 9e12e034dbc2 / 4

Type: `"string"`. Computed.

Location definition specified as the attributes of $\{location\} block in NGINX config. This
includes both the optional\_modifier and the location\_match combined. A location can either be
defined by a prefix string, or by a regular expression.

<a id="canonical-ea64337daf7be396a78c8f1de4beadf002934dcda62bd6d987b60ec5053e838a"></a>

<a id="canonical-d0bd91e8a95c6dfef77e179c17e99e81edf9469b9b914de2761f1cd89a35ea84"></a>

## name property — server_spec.locations / 9e12e034dbc2 / 5

Type: `"string"`. Computed.

Uniqueness identifier for a location definition.

- [waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160): complete subsection reference.

<a id="canonical-bc8a70ec610521dff7c89e6ca8feca0dc68c1c8804b02aa43577fa5a8bcd61ac"></a>

## Next pages — server_spec.locations / 9e12e034dbc2 / 6

- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8bf846b6989c0b1fd9f0c99f1a46a51f48ed337bd8a2578e8c0ff0497cf1c6f"></a>

## server_spec.locations.api_discovery_spec — server_spec.locations.api_discovery_spec / 5fd3761f83bc / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- server_spec.locations.api_discovery_spec

<a id="canonical-d9941d0b8e4b1765d9c647e0cb20e40f26fdf8d1e07f84a874655ddb55abd9e3"></a>

Type: `"single"`. Computed.

Configuration for api\_discovery\_spec.

<a id="canonical-d830041247767cf1e62c39238078e7cdc0871c4d45d0f0a2de1334e81c1ee478"></a>

## Direct properties — server_spec.locations.api_discovery_spec / 5fd3761f83bc / 3

- [disabled](data-sources--nginx_server--reference--group-001.md#canonical-0f4d14567fc784c218ea8a46e4b859f59700847358fb8f0bb5f414548bdb7cf3): complete subsection reference.

- [enabled](data-sources--nginx_server--reference--group-001.md#canonical-e08df24bafdc8d9455b0c8b00da0f9a59e96085220426f8e150d843fc78d40cf): complete subsection reference.

<a id="canonical-9cc10842ffad928d0949c003bb6349a88e4634cca00bc4d47bfea2d311cec122"></a>

## Next pages — server_spec.locations.api_discovery_spec / 5fd3761f83bc / 4

- [server_spec.locations.api_discovery_spec.disabled](data-sources--nginx_server--reference--group-001.md#canonical-0f4d14567fc784c218ea8a46e4b859f59700847358fb8f0bb5f414548bdb7cf3)
- [server_spec.locations.api_discovery_spec.enabled](data-sources--nginx_server--reference--group-001.md#canonical-e08df24bafdc8d9455b0c8b00da0f9a59e96085220426f8e150d843fc78d40cf)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-0f4d14567fc784c218ea8a46e4b859f59700847358fb8f0bb5f414548bdb7cf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a13491813293b452ee78a06bc62d62775e87fa1d561c69511da60cdb0b70a3cd"></a>

## server_spec.locations.api_discovery_spec.disabled — server_spec.locations.api_discovery_spec.disabled / f2945d052cc5 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28)
- server_spec.locations.api_discovery_spec.disabled

<a id="canonical-44ff3c4bdb34817a8a5b36b5366fe7c25adf743c727337f3706b102d5ed11b09"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-fde4597c5f486ac07d7ca2ff3a317385bffd40ee493e927bcb1624bb15c5b952"></a>

## Direct properties — server_spec.locations.api_discovery_spec.disabled / f2945d052cc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eabab92796fd90864c556a52afac5a6ed1d5198371c2ed6821ede0f163c5b032"></a>

## Next pages — server_spec.locations.api_discovery_spec.disabled / f2945d052cc5 / 4

- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-e08df24bafdc8d9455b0c8b00da0f9a59e96085220426f8e150d843fc78d40cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5054bb6b014273e32f1cab8c19a0202e809389aa6917b5f346d60e7df76abdc"></a>

## server_spec.locations.api_discovery_spec.enabled — server_spec.locations.api_discovery_spec.enabled / 9e68f213e5b4 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28)
- server_spec.locations.api_discovery_spec.enabled

<a id="canonical-38432fcfd8878e15110fdb53067f766737af9e7d96855c6846cb846803dc1585"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-bf398b12a25d6e9d31e1b3833aed6b8c05f7ffaf46f1f37745d351952cbe3bd4"></a>

## Direct properties — server_spec.locations.api_discovery_spec.enabled / 9e68f213e5b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed4ef25dcbc069c8aa9d13493fcab8c545be4bfc5626c2e784f3f9bcaf0d6cf2"></a>

## Next pages — server_spec.locations.api_discovery_spec.enabled / 9e68f213e5b4 / 4

- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--reference--group-001.md#canonical-4fa002ffe3708b4301e1298bccb263d925b53a7b23f6752ad131f06c76ba4e28)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e51a61780b32ecbd5fa864bb365ffed982c379ad4a0be40486b591d470bf6a2"></a>

## server_spec.locations.waf_spec — server_spec.locations.waf_spec / fa2f3369cb99 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- server_spec.locations.waf_spec

<a id="canonical-70d6ae10e3ef2d93bbd5a7520b4ffeb1754b9a122ed07b102ece0347ee83dafb"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

<a id="canonical-814998fd68ac8781c086e4b231380d2e0fc7475a676263278e0e746e450810d9"></a>

## Direct properties — server_spec.locations.waf_spec / fa2f3369cb99 / 3

- [blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-b6262430946203b9c68f080890c607df232664c2ca173e57b5142cb034bec738): complete subsection reference.

- [distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-affb442d8b80cb67670cf07cbfe92e5a60cd5a6f0dd889a8657123744d06b8d4): complete subsection reference.

- [monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-b6a8ebd0e6ddd6344f58b1816c0f028dc74f6fc523c7a8ec012016f9ccebf6f9): complete subsection reference.

- [nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-bb5910a6ba93cfab29f5508c144616c8d1c7293cbec84753c744f8732058fdd6): complete subsection reference.

- [none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-1e669e8b32a4b0361d4976977ac69402da98348015e5bf03b86dbf3fdd757b3b): complete subsection reference.

<a id="canonical-043b0138fb2e67b9e21b5049e63e294c3e65d9e24e05902005edfc45038f06a0"></a>

<a id="canonical-4a220c004cbc617b3a79de456ec82ab99d01f36fe1d0372a25932b4eb0bb2627"></a>

## policy_file_name property — server_spec.locations.waf_spec / fa2f3369cb99 / 4

Type: `"string"`. Computed.

WAF Policy File Name. Policy file name for WAF.

<a id="canonical-ae1b0218f260c34a0680e19625c5981dcf0464b3f542893325eaae54025fa768"></a>

<a id="canonical-592231014f4c35301c5441a138c540d420a71c9eb510feb44f2d868a2eed01e8"></a>

## policy_name property — server_spec.locations.waf_spec / fa2f3369cb99 / 5

Type: `"string"`. Computed.

WAF Policy Name. Policy name configured for WAF.

<a id="canonical-e0a52aa1db5ee3a971a6bf2df2e10b6fc9bf9d553cac94111cd44cb603c68592"></a>

<a id="canonical-993218f14ebe7f304d9115c9b318ebc74219782d47c2c1e93304024b42722533"></a>

## security_log_enabled property — server_spec.locations.waf_spec / fa2f3369cb99 / 6

Type: `"bool"`. Computed.

Specifies if security logging is enabled.

<a id="canonical-4a6213297c2b79bdff9964caeb458fe7ea2d09d0b43c6d769588e5a61f2343e3"></a>

<a id="canonical-62b64e50a450c809e5de7bc1cdbaffb4192cadacaa83b3603b052c208731e46e"></a>

## security_log_file_names property — server_spec.locations.waf_spec / fa2f3369cb99 / 7

Type: `["list", "string"]`. Computed.

Specifies the list of security log files specification.

<a id="canonical-76d6dcb9d7d2b639097529e4352bf48f07b4587a5de0b23df4912d887f783ae5"></a>

## Next pages — server_spec.locations.waf_spec / fa2f3369cb99 / 8

- [server_spec.locations.waf_spec.blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-b6262430946203b9c68f080890c607df232664c2ca173e57b5142cb034bec738)
- [server_spec.locations.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-affb442d8b80cb67670cf07cbfe92e5a60cd5a6f0dd889a8657123744d06b8d4)
- [server_spec.locations.waf_spec.monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-b6a8ebd0e6ddd6344f58b1816c0f028dc74f6fc523c7a8ec012016f9ccebf6f9)
- [server_spec.locations.waf_spec.nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-bb5910a6ba93cfab29f5508c144616c8d1c7293cbec84753c744f8732058fdd6)
- [server_spec.locations.waf_spec.none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-1e669e8b32a4b0361d4976977ac69402da98348015e5bf03b86dbf3fdd757b3b)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-b6262430946203b9c68f080890c607df232664c2ca173e57b5142cb034bec738"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d391298eed03a97f66377eef77afc897ab6e4d19f21a7e65552b990c222b04b1"></a>

## server_spec.locations.waf_spec.blocking_waf_mode — server_spec.locations.waf_spec.blocking_waf_mode / 2ad333c6afe2 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- server_spec.locations.waf_spec.blocking_waf_mode

<a id="canonical-06fe7a573eebd3f04a4c9777365670c2897e739b4a50835ad6b060b2e4e5861e"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-e705e0e62f422d44d19d218498bd9689977464c314626d9841638220c6ea3c84"></a>

## Direct properties — server_spec.locations.waf_spec.blocking_waf_mode / 2ad333c6afe2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a020b6985ca45f04f66c7353c4bb39479b537a3f45d0e743b81712414eefdf5"></a>

## Next pages — server_spec.locations.waf_spec.blocking_waf_mode / 2ad333c6afe2 / 4

- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-affb442d8b80cb67670cf07cbfe92e5a60cd5a6f0dd889a8657123744d06b8d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-793e432e1ee68ac24b8b0beb88188504044e67acf667aa21abec262cf1587065"></a>

## server_spec.locations.waf_spec.distributed_cloud_policy_management — server_spec.locations.waf_spec.distributed_cloud_policy_management / e893d737b965 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- server_spec.locations.waf_spec.distributed_cloud_policy_management

<a id="canonical-431f50eef6ad14fe2400ce12b3941f4c88f6c5c252d67f2f36ef151c0c89f353"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for distributed cloud policy management.

<a id="canonical-1c91addbc5d1c790dbe7837e9299bf3ef3f85e839ca089d72d479076be35a962"></a>

## Direct properties — server_spec.locations.waf_spec.distributed_cloud_policy_management / e893d737b965 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3190874fd4d31bb607a92c45c9419a148dd8bcd36d2b1e5b79898e4a26331101"></a>

## Next pages — server_spec.locations.waf_spec.distributed_cloud_policy_management / e893d737b965 / 4

- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-b6a8ebd0e6ddd6344f58b1816c0f028dc74f6fc523c7a8ec012016f9ccebf6f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb9519bea087748f0b4391f12d67bf6cbcb97a890ed64361a88dec642343a25e"></a>

## server_spec.locations.waf_spec.monitoring_waf_mode — server_spec.locations.waf_spec.monitoring_waf_mode / 3d443692e3fd / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- server_spec.locations.waf_spec.monitoring_waf_mode

<a id="canonical-51820a84a88c52271591c80376d5ab82caae1b60b10905171efca2b059156940"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for monitoring waf mode.

<a id="canonical-7fa6b62b44a61db8b1a10fc45b910852286aab47d64db76fe709613b960b1a48"></a>

## Direct properties — server_spec.locations.waf_spec.monitoring_waf_mode / 3d443692e3fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b462816a4ec0f90b0093bf7b9ec4e9f14621bf69c92b25a9c0e46cf6af0120de"></a>

## Next pages — server_spec.locations.waf_spec.monitoring_waf_mode / 3d443692e3fd / 4

- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-bb5910a6ba93cfab29f5508c144616c8d1c7293cbec84753c744f8732058fdd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-525a383e5e049e1d62a5787ef1e1a0f96d7a3547695fb7be057b89d3cc1e0e68"></a>

## server_spec.locations.waf_spec.nginx_policy_management — server_spec.locations.waf_spec.nginx_policy_management / 4fc4fd2ad352 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- server_spec.locations.waf_spec.nginx_policy_management

<a id="canonical-e82b778ebce2bc4f0317274434627eb82b5b9f97b34959b2c9bd49a52da29a0d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nginx policy management.

<a id="canonical-524f1b540ff7ad5932cbd5bdf4fc9abf8a7abc2a6511960e98c98a6d0ec28bf3"></a>

## Direct properties — server_spec.locations.waf_spec.nginx_policy_management / 4fc4fd2ad352 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3caf027ea0b4b9f2da05182345e60937b58109e0415a4566e5f45232f287386"></a>

## Next pages — server_spec.locations.waf_spec.nginx_policy_management / 4fc4fd2ad352 / 4

- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-1e669e8b32a4b0361d4976977ac69402da98348015e5bf03b86dbf3fdd757b3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-612e7114ebfd5fd9805a98b5f2ad619a2f28d8f89dd618f4dac74f3ec12917a2"></a>

## server_spec.locations.waf_spec.none_waf_mode — server_spec.locations.waf_spec.none_waf_mode / 624c78887a87 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.locations](data-sources--nginx_server--reference--group-001.md#canonical-a3075c28b19fe347eb6616776efadc26f2853d35beb3a759a400c0e4a88eb863)
- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- server_spec.locations.waf_spec.none_waf_mode

<a id="canonical-38379bb9be093ba27cbab3c6f872c97b7e1f271131d76252b8e807c31a629c6f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for none waf mode.

<a id="canonical-4aa08f87a580ec02ac90f02a72cea6f1499fd588f1777827db063f558812e833"></a>

## Direct properties — server_spec.locations.waf_spec.none_waf_mode / 624c78887a87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bac28d08f12bc88a343974105397e590cbc3c1d7ab65e2e47cf856c18d4acb31"></a>

## Next pages — server_spec.locations.waf_spec.none_waf_mode / 624c78887a87 / 4

- [server_spec.locations.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-d0edde0c350edb5d8a6f57a5c2df4ad9052be13c19b1f666aa631b246c1f6160)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-994722a040cbab979f458a41da7f5461205a3621121bb15962ec1ca109c65e0c"></a>

## server_spec.waf_spec — server_spec.waf_spec / e83a96ba5bd3 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- server_spec.waf_spec

<a id="canonical-3c19225c51d0b6daccb2e377c4676fefa0f44828f84c5dea951465776b64e529"></a>

Type: `"single"`. Computed.

Configuration for waf\_spec.

<a id="canonical-8585f51d27f69113dc60a93f459ad7cfdc4f07f6762fd673c569d1dd3179b451"></a>

## Direct properties — server_spec.waf_spec / e83a96ba5bd3 / 3

- [blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-254976d0c00c977ee869e54e329da08eadbddf15d01aecf26e7fd2c10473c5e1): complete subsection reference.

- [distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-34698be7e157b9b9e82b92c03bf71f965f34a0f647d4f0f329e683fdfa1edf95): complete subsection reference.

- [monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-7b9d0f41a0d43eedd6b7cdd1e6e52061337297d95dea152d61c6444567ca3cb0): complete subsection reference.

- [nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-83dd907df3b2c4e3254113db4a0a77436eb068c64e5af28787e6e43f17fb724a): complete subsection reference.

- [none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-54e015aec5cc07be1d2264d695ebf32903083bf0eecdf3cdd60541f3d1fde574): complete subsection reference.

<a id="canonical-ada9cf03ce47f71d52ff16a0c286207c21db9a98c36f69402b7d8cffb537e4f1"></a>

<a id="canonical-74775f0fc9fc3be9c6eeb82e78f348f725edab0033f3062dfce0fb25fbb4057e"></a>

## policy_file_name property — server_spec.waf_spec / e83a96ba5bd3 / 4

Type: `"string"`. Computed.

WAF Policy File Name. Policy file name for WAF.

<a id="canonical-43646c1a88d2cec12f473fcccd439380cfb9d7b3b92ba9d5f68cf5ba7e958b16"></a>

<a id="canonical-a857f6872c715d6eaf968c142de9fd4eee2a8a022fbbeb7fe545e169d9a8baa2"></a>

## policy_name property — server_spec.waf_spec / e83a96ba5bd3 / 5

Type: `"string"`. Computed.

WAF Policy Name. Policy name configured for WAF.

<a id="canonical-ae0143e666d7d17ed3e32bf10467c514fb99d47c9d575ff2e9dbeba8566b99c6"></a>

<a id="canonical-a495cd0238584fe12ccda6c46e3477e9f5e0c0849e9c87c217732e7eae22a721"></a>

## security_log_enabled property — server_spec.waf_spec / e83a96ba5bd3 / 6

Type: `"bool"`. Computed.

Specifies if security logging is enabled.

<a id="canonical-2c0762836e69a962e3eb066891b003d5ec003fc372cd0133f33d341d0511d48e"></a>

<a id="canonical-26bf7f4121b9338c342be943e4bf54f83efe68ba99bdbc6b0321986fa5496885"></a>

## security_log_file_names property — server_spec.waf_spec / e83a96ba5bd3 / 7

Type: `["list", "string"]`. Computed.

Specifies the list of security log files specification.

<a id="canonical-a521f838f1761c93178a710bf6ab1e33b2712eb4125fc9e26ace0670903098ad"></a>

## Next pages — server_spec.waf_spec / e83a96ba5bd3 / 8

- [server_spec.waf_spec.blocking_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-254976d0c00c977ee869e54e329da08eadbddf15d01aecf26e7fd2c10473c5e1)
- [server_spec.waf_spec.distributed_cloud_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-34698be7e157b9b9e82b92c03bf71f965f34a0f647d4f0f329e683fdfa1edf95)
- [server_spec.waf_spec.monitoring_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-7b9d0f41a0d43eedd6b7cdd1e6e52061337297d95dea152d61c6444567ca3cb0)
- [server_spec.waf_spec.nginx_policy_management](data-sources--nginx_server--reference--group-001.md#canonical-83dd907df3b2c4e3254113db4a0a77436eb068c64e5af28787e6e43f17fb724a)
- [server_spec.waf_spec.none_waf_mode](data-sources--nginx_server--reference--group-001.md#canonical-54e015aec5cc07be1d2264d695ebf32903083bf0eecdf3cdd60541f3d1fde574)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-254976d0c00c977ee869e54e329da08eadbddf15d01aecf26e7fd2c10473c5e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b7bf2dfaa92230c8a2641b2f0a58eff51ae4dbb43e08edadb842c2eb218763d"></a>

## server_spec.waf_spec.blocking_waf_mode — server_spec.waf_spec.blocking_waf_mode / b03af2a0cc07 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- server_spec.waf_spec.blocking_waf_mode

<a id="canonical-94b2189b032bdd9ab3615f8e5affa3f7df0c5a0483d1176a9c8b184b6b1b5aa8"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-42f8987fe47f297998b9443ac1966364c407085b0aa3a500c95ef4f7e269c913"></a>

## Direct properties — server_spec.waf_spec.blocking_waf_mode / b03af2a0cc07 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-802cbac9516ead3a2540aa9ee7f50bd626fc2b6e7cb22e0b56fdfbb02f492767"></a>

## Next pages — server_spec.waf_spec.blocking_waf_mode / b03af2a0cc07 / 4

- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-34698be7e157b9b9e82b92c03bf71f965f34a0f647d4f0f329e683fdfa1edf95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d83d40e2a5a0becf2ded128335fdbf969a8db911dcbcc29040e8a00ed3549315"></a>

## server_spec.waf_spec.distributed_cloud_policy_management — server_spec.waf_spec.distributed_cloud_policy_management / 33352b34c835 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- server_spec.waf_spec.distributed_cloud_policy_management

<a id="canonical-ae6885d05917e503ca34899083e8b88132d6e74590fd2bbe58d8362c3cc21505"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for distributed cloud policy management.

<a id="canonical-e6e7ce39ec3be4179bf530b1a2c181359ac1839cc56910b964d8a29cb8134e03"></a>

## Direct properties — server_spec.waf_spec.distributed_cloud_policy_management / 33352b34c835 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6e5ba47848cd7bbe219d4a076df3c364e7aeb648a91a3b4e96f8d2ca61aa956"></a>

## Next pages — server_spec.waf_spec.distributed_cloud_policy_management / 33352b34c835 / 4

- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-7b9d0f41a0d43eedd6b7cdd1e6e52061337297d95dea152d61c6444567ca3cb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f6c1fa55ce90f298b2bf1750c1e8c492e7b6d02bd5f5a9f081714410f1c0563"></a>

## server_spec.waf_spec.monitoring_waf_mode — server_spec.waf_spec.monitoring_waf_mode / 5743bd463348 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- server_spec.waf_spec.monitoring_waf_mode

<a id="canonical-65ca80fb697f9ae4370154d1b2db63bcc2544a92e664f39a717fbad105716e42"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for monitoring waf mode.

<a id="canonical-ac9bfe420c93144fe05fe8ca78d0e097f00aad30babbc1101405720b16beac93"></a>

## Direct properties — server_spec.waf_spec.monitoring_waf_mode / 5743bd463348 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1edc875eb36da1c8313285479283bdf612638922fccea5a87d40f65f1a7092a2"></a>

## Next pages — server_spec.waf_spec.monitoring_waf_mode / 5743bd463348 / 4

- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-83dd907df3b2c4e3254113db4a0a77436eb068c64e5af28787e6e43f17fb724a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c1bd554c622b5225c6f0b50af67ad98d6d71d24557a5547d3cb0358133ff6f3"></a>

## server_spec.waf_spec.nginx_policy_management — server_spec.waf_spec.nginx_policy_management / aa6c8176979b / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- server_spec.waf_spec.nginx_policy_management

<a id="canonical-f9d99c3a120d9c671b692df6ee12e9e8327c826593638e5a8561208b9e9bda51"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nginx policy management.

<a id="canonical-c961d0eaa718f48ebce7e764ac8dd00baab504a68540e310cd161c48d8a2bc46"></a>

## Direct properties — server_spec.waf_spec.nginx_policy_management / aa6c8176979b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c73539a8b08759907dbcc4a27b10cda93821b19ad1f58d4bf2261627c2386a18"></a>

## Next pages — server_spec.waf_spec.nginx_policy_management / aa6c8176979b / 4

- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-54e015aec5cc07be1d2264d695ebf32903083bf0eecdf3cdd60541f3d1fde574"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8946faba8757182401fcbc128cbfd3afc7d94c2ea124c651be714e53c9f12c56"></a>

## server_spec.waf_spec.none_waf_mode — server_spec.waf_spec.none_waf_mode / d95efd06bc24 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Property reference](data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [server_spec](data-sources--nginx_server--reference--group-001.md#canonical-1acb410979a4cd5f50d26e5cba9a63aee3ae9293b2af2eda8907309a7b896c7d)
- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- server_spec.waf_spec.none_waf_mode

<a id="canonical-41b71fdcaa4951d525b1a8328f988c3c45b9a5f324cc8721cfb808d6ed9cd1ff"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for none waf mode.

<a id="canonical-f4a4c96819b7038a68a61326b5a4f19e77b5b6e4bda25908a7cc1ca979ef2fe0"></a>

## Direct properties — server_spec.waf_spec.none_waf_mode / d95efd06bc24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e321101891499ac45c2913a04ae4e1bc1b0ed745a40ee8d1abf95ad266c30c19"></a>

## Next pages — server_spec.waf_spec.none_waf_mode / d95efd06bc24 / 4

- [server_spec.waf_spec](data-sources--nginx_server--reference--group-001.md#canonical-9fbef0bdb99940bdecb01dceadd3fee3b29c1a905ac9508f01827cb3aafc12ee)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
