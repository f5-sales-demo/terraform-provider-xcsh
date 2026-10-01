---
page_title: "xcsh_site_cloud_init reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_cloud_init reference."
---

# xcsh_site_cloud_init reference

<a id="canonical-449146e9b12ccd06dd9f1b3726adeaaf418081765a16b4774522c5040c4e90da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-288153598c164d03c6ab0d27aa881f8dbd2a3cee4864fef186ffbd770958690d"></a>

## Property reference — Property reference / fe16a54890ab / 2

Breadcrumbs:

- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)
- Property reference

<a id="canonical-4b9f5107a7bec4ac57e601ed23d1a2c681627999217c44c5bcc9afe14a8f08fb"></a>

## Direct properties — Property reference / fe16a54890ab / 3

<a id="canonical-73ec927e6aaea20b8af4f8a52074f534e534036751a6a547d35266f51b3b5d69"></a>

<a id="canonical-4885bb8df31b045b870b941e41fd534655d99b45d6c79b970ce9c6cc5b4628af"></a>

## cloud_init_config property — Property reference / fe16a54890ab / 4

Type: `"string"`. Computed, Sensitive.

Cloud-init template with an unresolved token placeholder; substitute a separately issued site-bound
JWT before deployment. This sensitive value is stored in Terraform state; protect state access
accordingly.

<a id="canonical-efba95b75c9db8d1ae8ce5fa07e6042bfb7989bfaefa0a884ed13423dc627623"></a>

<a id="canonical-d3ef0d3d04a241682ea0e31b7fd220b41feb4ceafc48b2e95a1ca0a570979c43"></a>

## enable_management_network property — Property reference / fe16a54890ab / 5

Type: `"bool"`. Optional.

Management network choice for this cloud-init config.

<a id="canonical-07449237127fba92ea4a1e1b10feba08a45b4347704ac72078169bc181432f10"></a>

<a id="canonical-9da41f17af5e2192710e25c32f1f3ffe85080911fa796fda0d58fd2f4683b606"></a>

## provider_ref property — Property reference / fe16a54890ab / 6

Type: `"string"`. Required.

Provider for that cloud-init config.

<a id="canonical-9bef5f3a7c54735870327474a0c934a69c4c43e8aeee54c6ac49790da18081b2"></a>

<a id="canonical-0e595a201ff7bd3eda9be367224008762d107d86bcc9f146888d168f6d2ef6b9"></a>

## site_name property — Property reference / fe16a54890ab / 7

Type: `"string"`. Required.

Site name for this cloud-init config.

<a id="canonical-01a0d5a2024de70c090d4a953582806f9b99c67d1934935ab0b0c3f335e952b2"></a>

## All schema paths — Property reference / fe16a54890ab / 8

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cloud_init_config` | [cloud_init_config](data-sources--site_cloud_init--reference--group-001.md#canonical-73ec927e6aaea20b8af4f8a52074f534e534036751a6a547d35266f51b3b5d69) |
| `enable_management_network` | [enable_management_network](data-sources--site_cloud_init--reference--group-001.md#canonical-efba95b75c9db8d1ae8ce5fa07e6042bfb7989bfaefa0a884ed13423dc627623) |
| `provider_ref` | [provider_ref](data-sources--site_cloud_init--reference--group-001.md#canonical-07449237127fba92ea4a1e1b10feba08a45b4347704ac72078169bc181432f10) |
| `site_name` | [site_name](data-sources--site_cloud_init--reference--group-001.md#canonical-9bef5f3a7c54735870327474a0c934a69c4c43e8aeee54c6ac49790da18081b2) |

<a id="canonical-1f47685cf3d7857a26301b700d3b206bc38c7353c0a09a549c4d77b59c45e3b1"></a>

## Next pages — Property reference / fe16a54890ab / 9

- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md#canonical-bb83253dcef707bb0c00385589b2ba967da3ad61756414aa9015e45e3d75e738)
