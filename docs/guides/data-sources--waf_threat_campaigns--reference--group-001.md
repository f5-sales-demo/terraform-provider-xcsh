---
page_title: "xcsh_waf_threat_campaigns reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threat_campaigns reference."
---

# xcsh_waf_threat_campaigns reference

<a id="canonical-b6e27744b7fb85106a6f572c628de8e4a957ca299bbde4a98a1b3f663513454a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7b9cfa59566f777696b5c54ac1157c633a202a1fda9d62253c97f7d8204f60c"></a>

## Property reference — Property reference / f80190199e37 / 2

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
- Property reference

<a id="canonical-cc3d22a31a6ab01c0dbc2f14af620068e730b9548a88d92d2d932e600f0bffa2"></a>

## Direct properties — Property reference / f80190199e37 / 3

<a id="canonical-554594458fff75d14b8ec81ea2f2b4ebd5ce640d0e429d4738758ebc9e530b20"></a>

<a id="canonical-6651cd586abf34b333a33c656725868e5b7b04ffb1c5bcd35856e73dfcdc51be"></a>

## known_version property — Property reference / f80190199e37 / 4

Type: `"string"`. Optional.

Version of the threat campaigns list that the client currently has, can be used for caching.

<a id="canonical-186f0bee4d81137a79cb51f75aa3f7d4b4086d939567cc3dc8b091941be67597"></a>

<a id="canonical-617dbd4b946e28bddc13d26a0982f6a4e032efb86947e4bfdfb4969b5714b7b1"></a>

## not_modified property — Property reference / f80190199e37 / 5

Type: `"bool"`. Computed.

Indicates if the attack signatures list has not been modified since the last version.

- [threat_campaigns](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-56407a2a2d542f5d0a27f604bfcf64f65d02e0274dc6a1cc5f16c77dbaca1849): complete subsection reference.

<a id="canonical-34b59746b3c8a68d79192a82e68cb1b495020fe25a20b35f4feb13e822d0210f"></a>

<a id="canonical-2a026695a577aec9bcdff299650843d99e1ce2791a4eb8674f3480510cac3648"></a>

## version property — Property reference / f80190199e37 / 6

Type: `"string"`. Computed.

Version of the attack signatures list, can be used for caching.

<a id="canonical-baa11879b2e219380f63e8f85c1875c31e320a56e6a330fcd58fd16496b02e7b"></a>

## All schema paths — Property reference / f80190199e37 / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `known_version` | [known_version](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-554594458fff75d14b8ec81ea2f2b4ebd5ce640d0e429d4738758ebc9e530b20) |
| `not_modified` | [not_modified](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-186f0bee4d81137a79cb51f75aa3f7d4b4086d939567cc3dc8b091941be67597) |
| `threat_campaigns` | [threat_campaigns](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-a20a5000f40ab6e5a3b3237052ce13890f749681d71834d09e46f7ef81d01db6) |
| `threat_campaigns.attack_type` | [threat_campaigns.attack_type](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-9901a075d1d6b3a5055068b21ab8a2ac696423e8be863ac719c52217b2e590b4) |
| `threat_campaigns.description_spec` | [threat_campaigns.description_spec](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-a526cc9a497e87b79633ce970d75f28cedb4dc0a76547a67e05cad317182e04f) |
| `threat_campaigns.id` | [threat_campaigns.id](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-b8757e4f514b28f7afdf7dd323e8c8041d746041f8cc1adfb724af5b5525a9c2) |
| `threat_campaigns.intent` | [threat_campaigns.intent](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-a695498d6d5cadb766f2f2e504d4e21bac777ccb8a223a6220dde25595c79085) |
| `threat_campaigns.last_update` | [threat_campaigns.last_update](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-5334feffb205ee8dbeeb29376bac7751458e742b497d7dda7739dfd9cba802e3) |
| `threat_campaigns.malwares` | [threat_campaigns.malwares](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-2f128264812582d4ecbe2fd0edf6b87cc0b0bb85f6876a613ee1a12ba01d3496) |
| `threat_campaigns.name` | [threat_campaigns.name](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-8066bba5b206615a63751d2b350171eac7f62fdca4fbe76404546777f6431f58) |
| `threat_campaigns.references` | [threat_campaigns.references](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-7879a4a6c8c2eec2981954bfebe5bcda0a02a4fa1f4d45fd5637cd5f9622ec67) |
| `threat_campaigns.risk` | [threat_campaigns.risk](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-75c0122be630a75aa800065235c9d9de9837d2a914a0658d0c4af8d232fe3079) |
| `threat_campaigns.systems` | [threat_campaigns.systems](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-c5f00efdf2cc63821de09ac50b59e96ef725a3c6cd1dd4e90de2a8111672462d) |
| `version` | [version](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-34b59746b3c8a68d79192a82e68cb1b495020fe25a20b35f4feb13e822d0210f) |

<a id="canonical-3bce946050cc508e86957f7c00f9c8a5a3a45acd5aca8b8267085660c1fb149d"></a>

## Next pages — Property reference / f80190199e37 / 8

- [threat_campaigns](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-56407a2a2d542f5d0a27f604bfcf64f65d02e0274dc6a1cc5f16c77dbaca1849)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)

<a id="canonical-56407a2a2d542f5d0a27f604bfcf64f65d02e0274dc6a1cc5f16c77dbaca1849"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97b04827089a307f729f14c0b386c8ca7e7f8bb87c8bbf58ecf2f9bd0d07b09a"></a>

## threat_campaigns — threat_campaigns / 22b7db4a006a / 2

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
- [Property reference](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-b6e27744b7fb85106a6f572c628de8e4a957ca299bbde4a98a1b3f663513454a)
- threat_campaigns

<a id="canonical-a20a5000f40ab6e5a3b3237052ce13890f749681d71834d09e46f7ef81d01db6"></a>

Type: `"list"`. Computed.

Threat Campaigns. A list of all supported threat campaigns.

<a id="canonical-bfa99f23bf8b6944456f26a5a990b0bb7eb4642e126656b2cdf5839cfc793792"></a>

## Direct properties — threat_campaigns / 22b7db4a006a / 3

<a id="canonical-9901a075d1d6b3a5055068b21ab8a2ac696423e8be863ac719c52217b2e590b4"></a>

<a id="canonical-f6bcdc4176c6a1dd01c97e4cd5b26afc5937e19b1819d9641f76568ad2c46e32"></a>

## attack_type property — threat_campaigns / 22b7db4a006a / 4

Type: `"string"`. Computed.

Attack Type. The Threat Campaign Attack Type.

<a id="canonical-a526cc9a497e87b79633ce970d75f28cedb4dc0a76547a67e05cad317182e04f"></a>

<a id="canonical-80de5768469d8cda19135e403238c020fa9c4ea3443cfd80b26b368a908511a3"></a>

## description_spec property — threat_campaigns / 22b7db4a006a / 5

Type: `"string"`. Computed.

Description. The Threat Campaign Description.

<a id="canonical-b8757e4f514b28f7afdf7dd323e8c8041d746041f8cc1adfb724af5b5525a9c2"></a>

<a id="canonical-6e3089395f9213b7038be095f51aa0d0487868b0a045a122f0c1bb11b81c9f1d"></a>

## id property — threat_campaigns / 22b7db4a006a / 6

Type: `"string"`. Computed.

ID. The Threat Campaign ID.

<a id="canonical-a695498d6d5cadb766f2f2e504d4e21bac777ccb8a223a6220dde25595c79085"></a>

<a id="canonical-4394bbe2f800feb02cf5afb42ca932f27d7c1cc3ccfec801706be68f1fde2e72"></a>

## intent property — threat_campaigns / 22b7db4a006a / 7

Type: `"string"`. Computed.

Intent. The Threat Campaign Intent.

<a id="canonical-5334feffb205ee8dbeeb29376bac7751458e742b497d7dda7739dfd9cba802e3"></a>

<a id="canonical-8657308edaadd61a7057a46a1f6c1ee07b877c54cccbbef8e38b7f65193cb9ce"></a>

## last_update property — threat_campaigns / 22b7db4a006a / 8

Type: `"string"`. Computed.

Last Update. The Threat Campaign last update time.

<a id="canonical-2f128264812582d4ecbe2fd0edf6b87cc0b0bb85f6876a613ee1a12ba01d3496"></a>

<a id="canonical-95a57abe116d71f520c4c80ecd321ccafc97eadd6be8833e3c7749d46c14a065"></a>

## malwares property — threat_campaigns / 22b7db4a006a / 9

Type: `["list", "string"]`. Computed.

Malwares. The Threat Campaign Malwares.

<a id="canonical-8066bba5b206615a63751d2b350171eac7f62fdca4fbe76404546777f6431f58"></a>

<a id="canonical-ee9624fafc1cd019c1f86a4e83b5169b92abccd9426e8cab18a81c337e397467"></a>

## name property — threat_campaigns / 22b7db4a006a / 10

Type: `"string"`. Computed.

Name. The Threat Campaign Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-7879a4a6c8c2eec2981954bfebe5bcda0a02a4fa1f4d45fd5637cd5f9622ec67"></a>

<a id="canonical-b8602533ebf39759d6468618e7ec508bba7758458b468932c232bfd256563879"></a>

## references property — threat_campaigns / 22b7db4a006a / 11

Type: `["list", "string"]`. Computed.

References. The Threat Campaign References.

<a id="canonical-75c0122be630a75aa800065235c9d9de9837d2a914a0658d0c4af8d232fe3079"></a>

<a id="canonical-43efc51f724bcd3246cc37335eb05ca27425c4305c0d05f619b8c86170932f84"></a>

## risk property — threat_campaigns / 22b7db4a006a / 12

Type: `"string"`. Computed.

Risk. The Threat Campaign Risk.

<a id="canonical-c5f00efdf2cc63821de09ac50b59e96ef725a3c6cd1dd4e90de2a8111672462d"></a>

<a id="canonical-ae25700ae4bc740a3d00561d8f9e9f6f044220f850ff1357b0674bd73046f7e5"></a>

## systems property — threat_campaigns / 22b7db4a006a / 13

Type: `["list", "string"]`. Computed.

Systems. The Threat Campaign Systems.

<a id="canonical-6d5b43340cf89cc89c2e4d80d9d3fa0f741423c75a167934b102335e5e2edcf0"></a>

## Next pages — threat_campaigns / 22b7db4a006a / 14

- [Property reference](data-sources--waf_threat_campaigns--reference--group-001.md#canonical-b6e27744b7fb85106a6f572c628de8e4a957ca299bbde4a98a1b3f663513454a)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md#canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9)
