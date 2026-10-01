---
page_title: "xcsh_addon_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service reference."
---

# xcsh_addon_service reference

<a id="canonical-11d603208fbd7d6adbc72e5aae21645bfb00d74593d5c996a45e0a5dfc81a055"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e2de6f73f6613aabf67aecd4970e93c8469b23b7cd955df70f04140382aeff2"></a>

## Property reference — Property reference / 78887c8213e5 / 2

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)
- Property reference

<a id="canonical-764a5d683b3009c36cd017f7fdbe44801d21258ac9587a8d166868cd8f7be0bc"></a>

## Direct properties — Property reference / 78887c8213e5 / 3

<a id="canonical-f23b9167346a63dc6785d5d9685e6bc33bbc402c06bfd3ead7ac8b00bd2a2266"></a>

<a id="canonical-36a190ec76eb5a65238b44139b9cf472fef072223ed7d790d4fc186e96025133"></a>

## activation_type property — Property reference / 78887c8213e5 / 4

Type: `"string"`. Computed.

How the addon service is activated. Possible values: \`self\` (user can activate directly),
\`partial\` (requires partial SRE management), \`managed\` (requires full manual intervention).

<a id="canonical-0d02ec89f2303be85d9c79e38915c8319bc6e1d54ec1ed0e74ad5901af50092b"></a>

<a id="canonical-a0e58df3e1206dac94844b6b4cc1df61d0e0c5de8629b9be961ca0b939f8ed4c"></a>

## addon_service_group_display_name property — Property reference / 78887c8213e5 / 5

Type: `"string"`. Computed.

Display name of the addon service group.

<a id="canonical-146a739f95d279c985d1e20bf08647e24f623696f28d6011f90b923bc9d54e5b"></a>

<a id="canonical-78e4af080e6e9fd23a1d8eebe6ccb9f0fe9b62876eeeca1233a46d995cfa20d0"></a>

## addon_service_group_name property — Property reference / 78887c8213e5 / 6

Type: `"string"`. Computed.

Name of the addon service group this service belongs to.

<a id="canonical-af36415a70bbd0d9c6479bc5b885a8f96db48013a191a202169a9010b59a2cb3"></a>

<a id="canonical-0a321fc09432963016c0dd7d28bae7ecd6f64ec0aeb831a20e3c4aec3106c106"></a>

## display_name property — Property reference / 78887c8213e5 / 7

Type: `"string"`. Computed.

Human-readable display name of the addon service.

<a id="canonical-3e16e03a729ca5186a318cbeb00e10e67dc1412659b350f6a569d56f77fefe67"></a>

<a id="canonical-92004e7640fd6cf478ab7859bf522c101bb09452903c04fae69a31e516fe08e0"></a>

## id property — Property reference / 78887c8213e5 / 8

Type: `"string"`. Computed.

Unique identifier for the data source.

<a id="canonical-6d7e7e04f546959a9f141c323d749dbcd2698bccf996a7b0df7484aa31b291a3"></a>

<a id="canonical-76946f2972d6f4c127c15af6a3d075d5633d19de5f2195820da8c2e482e4dbb6"></a>

## name property — Property reference / 78887c8213e5 / 9

Type: `"string"`. Required.

Name of the addon service (e.g., \`bot\_defense\`, \`client\_side\_defense\`).

<a id="canonical-5d5a49ec1559eaa20193ab8251f00241c24147e23a47927d3467b6e5be419600"></a>

<a id="canonical-5a7c58f9147d6ff7775394f8f45d073f04737d851c9ffdbd0d2a89b378a0c8b8"></a>

## namespace property — Property reference / 78887c8213e5 / 10

Type: `"string"`. Optional, Computed.

Namespace where the addon service is defined. Usually \`shared\`.

<a id="canonical-7b0955e529e52caaee2500635497b2947b79fa091ce907463fd05756b48fd35a"></a>

<a id="canonical-3aaa90c1a68382cc28bdd5381b54166723644ff493fd99a81d01b5487125410c"></a>

## tier property — Property reference / 78887c8213e5 / 11

Type: `"string"`. Computed.

Subscription tier required for this addon service. Possible values: \`NO\_TIER\`, \`BASIC\`,
\`STANDARD\`, \`ADVANCED\`, \`PREMIUM\`.

<a id="canonical-0f16ca4d5281148cec86373e31bd9a0e8ae3858c2a0933cbb2c492e8dfbbb57b"></a>

## All schema paths — Property reference / 78887c8213e5 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `activation_type` | [activation_type](data-sources--addon_service--reference--group-001.md#canonical-f23b9167346a63dc6785d5d9685e6bc33bbc402c06bfd3ead7ac8b00bd2a2266) |
| `addon_service_group_display_name` | [addon_service_group_display_name](data-sources--addon_service--reference--group-001.md#canonical-0d02ec89f2303be85d9c79e38915c8319bc6e1d54ec1ed0e74ad5901af50092b) |
| `addon_service_group_name` | [addon_service_group_name](data-sources--addon_service--reference--group-001.md#canonical-146a739f95d279c985d1e20bf08647e24f623696f28d6011f90b923bc9d54e5b) |
| `display_name` | [display_name](data-sources--addon_service--reference--group-001.md#canonical-af36415a70bbd0d9c6479bc5b885a8f96db48013a191a202169a9010b59a2cb3) |
| `id` | [id](data-sources--addon_service--reference--group-001.md#canonical-3e16e03a729ca5186a318cbeb00e10e67dc1412659b350f6a569d56f77fefe67) |
| `name` | [name](data-sources--addon_service--reference--group-001.md#canonical-6d7e7e04f546959a9f141c323d749dbcd2698bccf996a7b0df7484aa31b291a3) |
| `namespace` | [namespace](data-sources--addon_service--reference--group-001.md#canonical-5d5a49ec1559eaa20193ab8251f00241c24147e23a47927d3467b6e5be419600) |
| `tier` | [tier](data-sources--addon_service--reference--group-001.md#canonical-7b0955e529e52caaee2500635497b2947b79fa091ce907463fd05756b48fd35a) |

<a id="canonical-34c0f712d48920f7bc442b8b33a65a04c85f3dc70a493a797cc5bbe1bdff90ff"></a>

## Next pages — Property reference / 78887c8213e5 / 13

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)
