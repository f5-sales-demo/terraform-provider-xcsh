---
page_title: "xcsh_site_registration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registration reference."
---

# xcsh_site_registration reference

<a id="canonical-b894edd5f71c8e0dd62fda7420a5fb0ac3186df32f033ed0b50b08a89478d898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a5269ea376bba0ba67c92fe985a382ee162528047818e224e96491fccac25d2"></a>

## Property reference — Property reference / e3455772718d / 2

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)
- Property reference

<a id="canonical-79694e20c7fdd85f92ff09eb9da7a837dfcfa79544a21db78c82092fa9e2cbe3"></a>

## Direct properties — Property reference / e3455772718d / 3

<a id="canonical-aec9ad90c0e4fdc0e1cbc2ecea72c408007b996f26cd55222f19e9330c0fa14f"></a>

<a id="canonical-a45f92890b90387327442c6a6383aaa2a32b535053dfe7a18e6b0aa12afb2809"></a>

## cluster_name property — Property reference / e3455772718d / 4

Type: `"string"`. Computed.

Cluster name the CE registered with, as reported in its passport. Equals \`site\_name\` for a
correctly configured site.

<a id="canonical-350ba6ea57b8bfcdcf1e48f037a87aad25d25d9994df8b84aab1978821b580ea"></a>

<a id="canonical-5f70a32713ccc8e275572e2f80a25ec8f993d2d774c9195f1d660984d08541c8"></a>

## cluster_size property — Property reference / e3455772718d / 5

Type: `"number"`. Computed.

Number of nodes the CE reported for its cluster (1 for a single-node site, 3 for a three-node site).

<a id="canonical-fa579a82631ec69205ec49094aad8469c02f98a5ab7f304ecb07789408fb9dc3"></a>

<a id="canonical-116e1a40e862d1d4edc68796feaf165c562dd2e51a5f611733a0d90dd7a09d91"></a>

## found property — Property reference / e3455772718d / 6

Type: `"bool"`. Computed.

Whether a registration was resolved. \`false\` (with no error) while the CE has not registered yet —
gate an approval's \`count\` on this.

<a id="canonical-443b29fc9489a0cdc3e207e85a0bb3f5e82c031e20d8c6ebb34286c7e9104767"></a>

<a id="canonical-d91b60e69b352ab4914a817491b8749f4600b2029f16430935529b4ffa9bac2e"></a>

## hostname property — Property reference / e3455772718d / 7

Type: `"string"`. Optional, Computed.

Node hostname used to pick one registration when a multi-node site has several. Optional for a
single-node site; when omitted, the resolved node's hostname is returned here. Hostnames are only
unique within a site.

<a id="canonical-48f42d52eb3d001879f9d93ac7da9a42b8a07241932bd070dfc9ca1d60a80b88"></a>

<a id="canonical-8ecc802ed18868e54e56b84e8e05d3652d159467f4cc1d97f20d80b5b614b177"></a>

## id property — Property reference / e3455772718d / 8

Type: `"string"`. Computed.

Identifier of this lookup: the registration name when one is found, otherwise null.

<a id="canonical-193026a36589f07744daf09dd5df9dfba417e0e61fd5239fba98f41bbd2f479d"></a>

<a id="canonical-bbf53a4beaa9c70aedd0db442ba0a5e046fc56ecfa69599043e0b8cb93bcabb1"></a>

## instance_id property — Property reference / e3455772718d / 9

Type: `"string"`. Computed.

Infrastructure instance identifier reported by the CE registration
(\`get\_spec.infra.instance\_id\`). This distinguishes rebuilt nodes that reuse the same site and
hostname.

<a id="canonical-59f589bcf85990144ef8aa61c42601925237721557f4eb6b0ca660ea17c4a305"></a>

<a id="canonical-6cc1b20467ff9273c0c94faacf10003324b4b00b378a5979e315e1bb12092bf1"></a>

## name property — Property reference / e3455772718d / 10

Type: `"string"`. Computed.

Registration name (\`r-&lt;uuid&gt;\`) to pass to \`xcsh\_registration\_approval\`. Null when
\`found\` is \`false\`.

<a id="canonical-d069f563f997fcb0fcc0c404752d234fa45f3bfae4e6408a682d4e61cb0978e8"></a>

<a id="canonical-4aba3ee446a3501d84dfce6451a3fc6ac5136871133f4a46f6244a7318578709"></a>

## namespace property — Property reference / e3455772718d / 11

Type: `"string"`. Optional, Computed.

Namespace holding the registrations. Defaults to \`system\`, where site registrations live.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtLeast(1),
}
```

<a id="canonical-a2abb2396946b59d43c6167d695cdeb256278bae031e07166ffb49d75aaf0bc0"></a>

<a id="canonical-4ec098f57d628889b455fe6704082267fb48904e7cc1c0746839aad6dca82832"></a>

## provider_type property — Property reference / e3455772718d / 12

Type: `"string"`. Computed.

Infrastructure provider the CE reported, e.g. \`AZURE\`, \`AWS\`, \`GCP\`, \`VMWARE\`.

<a id="canonical-bd96ec30d28ccf43382431b60a2e0daed347dcafbe798b7b30f4c36485c3ae7f"></a>

<a id="canonical-273e9e1a785349ce7fc30416657d7032dfa3d85b9916cc5e35a9cb57acbbe775"></a>

## site_name property — Property reference / e3455772718d / 13

Type: `"string"`. Required.

Name of the F5 XC site whose CE registration should be resolved. Matched against each registration's
\`get\_spec.passport.cluster\_name\`.

<a id="canonical-4f81c64a81f3f1cf529547b26527331c0bd87c535e7ca98278d55c088eb77fec"></a>

<a id="canonical-6dfbdba4fb0f136b59b8f832226e132e7cfae0b67b85caa56e16abd0cc22bc0a"></a>

## state property — Property reference / e3455772718d / 14

Type: `"string"`. Computed.

Current registration state, e.g. \`PENDING\` (awaiting approval) or \`ONLINE\` (node admitted and
healthy).

<a id="canonical-f2a2fc680ef54b4da46b5be2ed0391652e963e3839c5095664ccf785979112e4"></a>

<a id="canonical-b5703dd14e702956959cc39de5de9aba6b774afe5d421cd7079e4dbbd8917da0"></a>

## uid property — Property reference / e3455772718d / 15

Type: `"string"`. Computed.

Unique identifier of the registration (the \`&lt;uuid&gt;\` part of the name).

<a id="canonical-aa68de8bbf25c0d91fb5c341ea8d5fc8d47b8b394c85767f8784d83a1e0c8907"></a>

## All schema paths — Property reference / e3455772718d / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cluster_name` | [cluster_name](data-sources--site_registration--reference--group-001.md#canonical-aec9ad90c0e4fdc0e1cbc2ecea72c408007b996f26cd55222f19e9330c0fa14f) |
| `cluster_size` | [cluster_size](data-sources--site_registration--reference--group-001.md#canonical-350ba6ea57b8bfcdcf1e48f037a87aad25d25d9994df8b84aab1978821b580ea) |
| `found` | [found](data-sources--site_registration--reference--group-001.md#canonical-fa579a82631ec69205ec49094aad8469c02f98a5ab7f304ecb07789408fb9dc3) |
| `hostname` | [hostname](data-sources--site_registration--reference--group-001.md#canonical-443b29fc9489a0cdc3e207e85a0bb3f5e82c031e20d8c6ebb34286c7e9104767) |
| `id` | [id](data-sources--site_registration--reference--group-001.md#canonical-48f42d52eb3d001879f9d93ac7da9a42b8a07241932bd070dfc9ca1d60a80b88) |
| `instance_id` | [instance_id](data-sources--site_registration--reference--group-001.md#canonical-193026a36589f07744daf09dd5df9dfba417e0e61fd5239fba98f41bbd2f479d) |
| `name` | [name](data-sources--site_registration--reference--group-001.md#canonical-59f589bcf85990144ef8aa61c42601925237721557f4eb6b0ca660ea17c4a305) |
| `namespace` | [namespace](data-sources--site_registration--reference--group-001.md#canonical-d069f563f997fcb0fcc0c404752d234fa45f3bfae4e6408a682d4e61cb0978e8) |
| `provider_type` | [provider_type](data-sources--site_registration--reference--group-001.md#canonical-a2abb2396946b59d43c6167d695cdeb256278bae031e07166ffb49d75aaf0bc0) |
| `site_name` | [site_name](data-sources--site_registration--reference--group-001.md#canonical-bd96ec30d28ccf43382431b60a2e0daed347dcafbe798b7b30f4c36485c3ae7f) |
| `state` | [state](data-sources--site_registration--reference--group-001.md#canonical-4f81c64a81f3f1cf529547b26527331c0bd87c535e7ca98278d55c088eb77fec) |
| `uid` | [uid](data-sources--site_registration--reference--group-001.md#canonical-f2a2fc680ef54b4da46b5be2ed0391652e963e3839c5095664ccf785979112e4) |

<a id="canonical-721b093f77674105560534bc869ff3145b8d681718ee28c7b4b4f8234adc4b9b"></a>

## Next pages — Property reference / e3455772718d / 17

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)
