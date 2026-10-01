---
page_title: "xcsh_addon_service_activation_status reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status reference."
---

# xcsh_addon_service_activation_status reference

<a id="canonical-f41bda08979dbea479c6a4df981142036fa7233a29104511d138789f181b019b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbe7bdcf3a9cb16c783f706931fbb981ca2c160f75ac9fd2add64aa037100fd1"></a>

## Property reference — Property reference / 38d3f36a8eb8 / 2

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)
- Property reference

<a id="canonical-26aa5487627d8831304e5d3cb45327e4b73e738ad534d2a64fd55b563bcca449"></a>

## Direct properties — Property reference / 38d3f36a8eb8 / 3

<a id="canonical-f22e40a26b02d51d3fdd512253f30d6022deb223e7d8b9d39782823cd769db55"></a>

<a id="canonical-a251e979befcd36114e659f92353ca6604759ff8f3d2736bc6626861630d0bb1"></a>

## addon_service property — Property reference / 38d3f36a8eb8 / 4

Type: `"string"`. Required.

Name of the addon service to check (e.g., \`bot\_defense\`, \`client\_side\_defense\`).

<a id="canonical-8c325481f5f9bde8b9e6a24ddde56c0fd5a45f170971b40924e46b5cc5821664"></a>

<a id="canonical-83c67d6ad77ae8f2ea391be262fbfe2a8a7ec520d3b730287bf1794fec7aa569"></a>

## can_activate property — Property reference / 38d3f36a8eb8 / 5

Type: `"bool"`. Computed.

Whether the addon service can be activated. True if state is \`AS\_NONE\` (not yet subscribed) or
\`AS\_SUBSCRIBED\` (already active).

<a id="canonical-08045c262fabf011a91708387695807e0d715a24a2049aab9db4b2d2c6ce2e29"></a>

<a id="canonical-8c6b46bc8abdc91367cf27ba1b21bd1e1299243506046b398b54c2e389d82a99"></a>

## id property — Property reference / 38d3f36a8eb8 / 6

Type: `"string"`. Computed.

Unique identifier for the data source.

<a id="canonical-429ba3008a463268e73c643dff432bf36440e2fa0ae175b8c1f299e31cfa0579"></a>

<a id="canonical-64722e438b3ca5befd898e632e88910059238353568f88214546f2242cc1c11a"></a>

## message property — Property reference / 38d3f36a8eb8 / 7

Type: `"string"`. Computed.

Human-readable message describing the current activation status.

<a id="canonical-bb444d1e5610fe066d34311366646b7bfb86882c737c5dedb879526fbe8582f3"></a>

<a id="canonical-0d6d0a7736d39a04ab7e47e3acc661dd3b29f489ce5c21b117df817796dc27da"></a>

## state property — Property reference / 38d3f36a8eb8 / 8

Type: `"string"`. Computed.

Current state of the addon service subscription. Possible values: \`AS\_NONE\`, \`AS\_PENDING\`,
\`AS\_SUBSCRIBED\`, \`AS\_ERROR\`.

<a id="canonical-d079e3a9a0869750cf2806a8a78bfd26ad26cf487f24a636193f58838f4da013"></a>

## All schema paths — Property reference / 38d3f36a8eb8 / 9

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `addon_service` | [addon_service](data-sources--addon_service_activation_status--reference--group-001.md#canonical-f22e40a26b02d51d3fdd512253f30d6022deb223e7d8b9d39782823cd769db55) |
| `can_activate` | [can_activate](data-sources--addon_service_activation_status--reference--group-001.md#canonical-8c325481f5f9bde8b9e6a24ddde56c0fd5a45f170971b40924e46b5cc5821664) |
| `id` | [id](data-sources--addon_service_activation_status--reference--group-001.md#canonical-08045c262fabf011a91708387695807e0d715a24a2049aab9db4b2d2c6ce2e29) |
| `message` | [message](data-sources--addon_service_activation_status--reference--group-001.md#canonical-429ba3008a463268e73c643dff432bf36440e2fa0ae175b8c1f299e31cfa0579) |
| `state` | [state](data-sources--addon_service_activation_status--reference--group-001.md#canonical-bb444d1e5610fe066d34311366646b7bfb86882c737c5dedb879526fbe8582f3) |

<a id="canonical-29df5012ba482e87b421f24d1cd931c7a92b6a3df071ac958c0a7f37be39477d"></a>

## Next pages — Property reference / 38d3f36a8eb8 / 10

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)
