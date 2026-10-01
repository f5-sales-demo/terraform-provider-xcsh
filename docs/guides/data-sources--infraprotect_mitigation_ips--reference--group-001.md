---
page_title: "xcsh_infraprotect_mitigation_ips reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_mitigation_ips reference."
---

# xcsh_infraprotect_mitigation_ips reference

<a id="canonical-e50d1fb885e742097142394b7c46c24d524be2fe0d2364e28add4376fdca80f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-481bab23bf4b4d5084bb2b0596d12995d959c98596bbe5a91e97667d79bb9b83"></a>

## Property reference — Property reference / 6503e5ab69ae / 2

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
- Property reference

<a id="canonical-eec8374b5aa681f7b9212d17f6083a8d9062ff9bb815febbd86994f1fc8097e1"></a>

## Direct properties — Property reference / 6503e5ab69ae / 3

- [ips](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-7f5b6990129ceba4e470330536851761a24a57b61659962add31db3a171cd654): complete subsection reference.

<a id="canonical-c7f793d38e8df87bd4885abd2ebab48c95b91dfd57428e9228372952a613cc30"></a>

<a id="canonical-4f88c3859e8654bcb4300e313832bf3216bdaf51626aa703f833ac3fe31f88b2"></a>

## mitigation_id property — Property reference / 6503e5ab69ae / 4

Type: `"string"`. Required.

Mitigation ID ID of the mitigation we want to GET the IPs for.

<a id="canonical-f85eedd119218430b7d48446f7a7fb65fc216b04c49a13c6cef2266dd602c014"></a>

<a id="canonical-48045040825b469dbaab1cd9f45eea1a955876503d4122d5f7ff090f382a99aa"></a>

## namespace property — Property reference / 6503e5ab69ae / 5

Type: `"string"`. Required.

Namespace This request is supported only in system namespace.

<a id="canonical-86e2a81d61297394fdffeed26f6af753323a59d05bdec5271c40c6c6256e9ebc"></a>

## All schema paths — Property reference / 6503e5ab69ae / 6

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ips` | [ips](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-b0e2f89e7c2bd3e09bc6a2fa18e82b2947fb38559b48e039bdc00b124e8ed7c5) |
| `ips.ip` | [ips.ip](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-d32bd1fb2221f74db46accf85cb2af55d263c9c7587ad671ce101ecb7e6092a6) |
| `ips.log_count` | [ips.log_count](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-64cddd0e6d8599a5fcfd72a62fc89be48a750d8074aa2f5db827e523d052168c) |
| `mitigation_id` | [mitigation_id](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-c7f793d38e8df87bd4885abd2ebab48c95b91dfd57428e9228372952a613cc30) |
| `namespace` | [namespace](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-f85eedd119218430b7d48446f7a7fb65fc216b04c49a13c6cef2266dd602c014) |

<a id="canonical-88389dc3614196cea4397f9e2797c6bea62c70f25085f500dae8bdb96269bd40"></a>

## Next pages — Property reference / 6503e5ab69ae / 7

- [ips](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-7f5b6990129ceba4e470330536851761a24a57b61659962add31db3a171cd654)
- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)

<a id="canonical-7f5b6990129ceba4e470330536851761a24a57b61659962add31db3a171cd654"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c2c8df27925d38f851b1c63e07f78df7623e1d249b8742e1681ebef295aff83"></a>

## ips — ips / a74e8b9c0602 / 2

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
- [Property reference](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-e50d1fb885e742097142394b7c46c24d524be2fe0d2364e28add4376fdca80f4)
- ips

<a id="canonical-b0e2f89e7c2bd3e09bc6a2fa18e82b2947fb38559b48e039bdc00b124e8ed7c5"></a>

Type: `"list"`. Computed.

List of IP addresses along with log count.

<a id="canonical-bcbe2fdc7faa2427f7909d1d00934e456e4c68f19c3ddb1da2fe151319aefc5c"></a>

## Direct properties — ips / a74e8b9c0602 / 3

<a id="canonical-d32bd1fb2221f74db46accf85cb2af55d263c9c7587ad671ce101ecb7e6092a6"></a>

<a id="canonical-7e4cc22514419fa726247f01707b5e7178bfb6a8ff7ac8033e52b7abd140da78"></a>

## ip property — ips / a74e8b9c0602 / 4

Type: `"string"`. Computed.

IP. Mitigation source IP.

<a id="canonical-64cddd0e6d8599a5fcfd72a62fc89be48a750d8074aa2f5db827e523d052168c"></a>

<a id="canonical-24dafa7d6765980fc41d831fce703c699641b0de899f8272d6b3d2afc39711b1"></a>

## log_count property — ips / a74e8b9c0602 / 5

Type: `"string"`. Computed.

Number of times the IP appears in the log.

<a id="canonical-989bd5da0237b7bc7dbbdfc92b80d729464fba95d9198424dda7a597b9db18b0"></a>

## Next pages — ips / a74e8b9c0602 / 6

- [Property reference](data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-e50d1fb885e742097142394b7c46c24d524be2fe0d2364e28add4376fdca80f4)
- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
