---
page_title: "xcsh_smsv2_kvm_runtime_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime_interface reference."
---

# xcsh_smsv2_kvm_runtime_interface reference

<a id="canonical-c66a286cbeb4e0f17e9c67b0ec0c11e7d79620ab55229d34ee040b173be79841"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a04945a4fa0cd55f4c9573180abb247099bad786f895a1ddbdd79ed9718f00df"></a>

## Property reference — Property reference / 6dd453a8774e / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)
- Property reference

<a id="canonical-73133fb8791659dcb50cf646741e0a293e351055231f5f65f1013cde408adc21"></a>

## Direct properties — Property reference / 6dd453a8774e / 3

<a id="canonical-ffc7a187bc2e7e0933a5045de8db4a585d92619c218879d63f094ac21eb0f8b9"></a>

<a id="canonical-742c2ea8fca12668645259e5c5b53b65794cbfd9b784bb0d1647fb4504b1567f"></a>

## configured property — Property reference / 6dd453a8774e / 4

Type: `"bool"`. Computed.

Whether the exact owned SLI currently uses static IPv4 configuration.

<a id="canonical-eff093e0a88c29e69752361f686cb4e3af2be3f22625e068f7c19c4809abea2e"></a>

<a id="canonical-e36f146b8b64988ce7adee9bde502cc5d7a35819bc3254895c5ae4d8cac3d874"></a>

## device property — Property reference / 6dd453a8774e / 5

Type: `"string"`. Computed.

Exact live registration device resolved from the expected MAC.

<a id="canonical-56e85536fdd1ce446a23cf73c1c01578e32b72c79f701e53bfb7e991e3116128"></a>

<a id="canonical-841fe45464c4a4ef56ca130c2aeb51c88adc3e0ff5975fec700c0c8d4fd6c270"></a>

## expected_mac property — Property reference / 6dd453a8774e / 6

Type: `"string"`. Required.

Terraform-owned SLI MAC used for live registration correlation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-df0ef2bc4018c3c1a291c9cab15fbd519153e85b03c1cdad94bada7e005d63d1"></a>

<a id="canonical-c06dac92265ca61757aa29d2ccd19c6b88f1ae7d9de72e1f5e9def15ea15683d"></a>

## hostname property — Property reference / 6dd453a8774e / 7

Type: `"string"`. Computed.

Exact live registration hostname resolved from the expected MAC.

<a id="canonical-cc1235f82563db6679b459b2f068224bcf2924c7d47cb54ef3df395cb2d083f5"></a>

<a id="canonical-1e91b621830e4805efc6538840be92baab00c65776b584cbd5253a61b1c1603e"></a>

## id property — Property reference / 6dd453a8774e / 8

Type: `"string"`. Computed.

Stable namespace/interface identity.

<a id="canonical-ac8ecd9bfc5878a41f1270148a8945910cb3cece78542d6548ae0ffcc6f16e23"></a>

<a id="canonical-816b122feba2e07a7cf780ea05ff47e11e9bb58309fc4c34ccb7110f2838446e"></a>

## interface_name property — Property reference / 6dd453a8774e / 9

Type: `"string"`. Computed.

Exact platform-generated SLI child name resolved from live ownership. Platform names may exceed 64
characters.

<a id="canonical-1a11b2d131d404271f86445c8bf6ff64a790c74f3718a35e9268bcb86385d739"></a>

<a id="canonical-ee8877019b119c60c9a9ebb2d9fbabf7f3b9f106724c968f606c40543bf3d4dc"></a>

## ipv4_cidr property — Property reference / 6dd453a8774e / 10

Type: `"string"`. Required.

Static IPv4 host address and prefix to configure on the SLI.

<a id="canonical-e867a939519f1a9b91fb69d065df6ddbd120ba65cb23ccc02c5a04ca85da021f"></a>

<a id="canonical-0bef29d458d47d578b1e8e23db8a0dbc68285782ac0d06078c035b372652d5b1"></a>

## namespace property — Property reference / 6dd453a8774e / 11

Type: `"string"`. Optional, Computed.

Namespace containing the site and runtime child. KVM SMSv2 supports only \`system\`.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{stringvalidator.OneOf("system")}
```

<a id="canonical-013c435595f8785eab95e3cf6ea1f53ab2d8be59dec9678d971cd11c001bfdd5"></a>

<a id="canonical-206fb38e511a557bb0876992c4c521753ba8e142971d170aa79e942a6aaf00df"></a>

## owner_uid property — Property reference / 6dd453a8774e / 12

Type: `"string"`. Computed.

Secure Mesh Site v2 UID that owns the runtime child.

<a id="canonical-658259f6187f4c24f9fe7d807ea50f40e250c844a092172bafbd4b5e22db41df"></a>

<a id="canonical-d196138cffc0448ea976b7bbc41c431a64aee927a36464bea20c0baa2805eac0"></a>

## resource_version property — Property reference / 6dd453a8774e / 13

Type: `"string"`. Computed.

Latest XC concurrency version observed after reconciliation.

<a id="canonical-416b7938fec8ff786ab6aff8f2c7312eb012b77281ebf772dea05146a055b95e"></a>

<a id="canonical-d3601aaf754f6add264f1d418a882cfb790f9ab290fc1d46bf4ca227f3b6913a"></a>

## site property — Property reference / 6dd453a8774e / 14

Type: `"string"`. Required.

Secure Mesh Site v2 configuration name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-3baa1efe6fb31c2ae323784e1ecbf436ec45ea42d032b1f24a53d8606064a61e"></a>

## All schema paths — Property reference / 6dd453a8774e / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `configured` | [configured](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-ffc7a187bc2e7e0933a5045de8db4a585d92619c218879d63f094ac21eb0f8b9) |
| `device` | [device](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-eff093e0a88c29e69752361f686cb4e3af2be3f22625e068f7c19c4809abea2e) |
| `expected_mac` | [expected_mac](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-56e85536fdd1ce446a23cf73c1c01578e32b72c79f701e53bfb7e991e3116128) |
| `hostname` | [hostname](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-df0ef2bc4018c3c1a291c9cab15fbd519153e85b03c1cdad94bada7e005d63d1) |
| `id` | [id](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-cc1235f82563db6679b459b2f068224bcf2924c7d47cb54ef3df395cb2d083f5) |
| `interface_name` | [interface_name](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-ac8ecd9bfc5878a41f1270148a8945910cb3cece78542d6548ae0ffcc6f16e23) |
| `ipv4_cidr` | [ipv4_cidr](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-1a11b2d131d404271f86445c8bf6ff64a790c74f3718a35e9268bcb86385d739) |
| `namespace` | [namespace](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-e867a939519f1a9b91fb69d065df6ddbd120ba65cb23ccc02c5a04ca85da021f) |
| `owner_uid` | [owner_uid](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-013c435595f8785eab95e3cf6ea1f53ab2d8be59dec9678d971cd11c001bfdd5) |
| `resource_version` | [resource_version](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-658259f6187f4c24f9fe7d807ea50f40e250c844a092172bafbd4b5e22db41df) |
| `site` | [site](resources--smsv2_kvm_runtime_interface--reference--group-001.md#canonical-416b7938fec8ff786ab6aff8f2c7312eb012b77281ebf772dea05146a055b95e) |

<a id="canonical-0418e39d6573854e1a81ebd31391cf1580a681a121ff932161a173f81452e054"></a>

## Next pages — Property reference / 6dd453a8774e / 16

- [xcsh_smsv2_kvm_runtime_interface](../resources/smsv2_kvm_runtime_interface.md#canonical-2f0c16a5489e111182e91669fdd8243ae8c574bdb1fc5fc1e4ab1808527039b3)
