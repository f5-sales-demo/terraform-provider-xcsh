---
page_title: "xcsh_access_active_session reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session reference."
---

# xcsh_access_active_session reference

<a id="canonical-f9145d41000cc5377b88dd55e2dfd831d018e430e7e6108f72457e70d726428d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3fae0ab4c0c0778013778fdf26219f2d25cec9cea0d78f576d137d3aea8a2da"></a>

## Property reference — Property reference / 101036385201 / 2

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
- Property reference

<a id="canonical-8e92be9a92f5ed8cbeecdac312d97e35506c2c33cf078bc331408d58a7583137"></a>

## Direct properties — Property reference / 101036385201 / 3

<a id="canonical-6854a1910303e8920b02c2769add2b8a5d905c7ab7275faaac0e92fda123ced8"></a>

<a id="canonical-88b68e43d1255ff6e901f1d25023dc0337986433d2fe13066e29c3e2f4aa268a"></a>

## client_ip property — Property reference / 101036385201 / 4

Type: `"string"`. Computed.

Client IP of the user that connected via the session.

<a id="canonical-f3d1e9ca4bb27dd69abbae464c0e9076a6be1726f37197931bd37e0f9a835342"></a>

<a id="canonical-eeb13c10641b40288652d24c45f9ebf328881b8124ecc31a4ebbfb7a531c30b4"></a>

## expiration_time property — Property reference / 101036385201 / 5

Type: `"string"`. Computed.

Expiration time of the session in RFC 3339 format.

<a id="canonical-c0ce9693ca644f253163196e11c9ebd64f262be01e9181baa3a79c2d789529d6"></a>

<a id="canonical-216efcbd9b93adb40530068d6cdfde573da4662befa35f965ce47f00db10c1b3"></a>

## id property — Property reference / 101036385201 / 6

Type: `"string"`. Required.

ID ID of the session.

<a id="canonical-70a55109e86cf2765fde3b07614985770865341d49813c5a9353237b277021c4"></a>

<a id="canonical-2489e4828083de7faac41ba6e6023c9495b316dc25f288cecd3fb8d8c8032552"></a>

## last_activity_time property — Property reference / 101036385201 / 7

Type: `"string"`. Computed.

Last activity time of the session in RFC 3339 format.

<a id="canonical-eda95bf6959e7d32c4d43f58d2b16926a1c7df56e5ccce44b21726dcb9846cce"></a>

<a id="canonical-ed620d0c2862ddfd6c453c96a095f7582e457de3eec7e278b60898d85df82b59"></a>

## namespace property — Property reference / 101036385201 / 8

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request.

<a id="canonical-9d8dfc3600fb47793def86f139b0328879ef415040cf51d10344260e8d82b6c4"></a>

<a id="canonical-4a8c3216e4f6b3b9df5be75d24d1e339cf4a762505799b68d1bb8f792b068483"></a>

## policy property — Property reference / 101036385201 / 9

Type: `"string"`. Computed.

Policy. Policy associated with the session.

<a id="canonical-611fdd92522efc09d9716debce7d9db29ac80e53e12c904f7ff96d1171a47649"></a>

<a id="canonical-c71232b0e7f551d26aaa030651121051f548d37889ffcaeb16d07a0440112df0"></a>

## site property — Property reference / 101036385201 / 10

Type: `"string"`. Computed.

Site. Site where the session is created.

<a id="canonical-ba825e81a407be6727fb354ce3297d6ee4ff96931af46cae7dbbe638840500e1"></a>

<a id="canonical-747dd72460d83aa0113b4943c4d1d0252518469a7f4d10e3fea7b8ed4ce5e8ee"></a>

## start_time property — Property reference / 101036385201 / 11

Type: `"string"`. Computed.

Start time of the session in RFC 3339 format.

<a id="canonical-0e8dd58a7ea34eb993da158f4281be138064ae947d2787fc3d48d5e432c86dc5"></a>

<a id="canonical-19b96b01de896e6b3315b6ae2438c578ca3152431f682243dc3edfdf45edc7b9"></a>

## status property — Property reference / 101036385201 / 12

Type: `"string"`. Computed.

\[Enum: PENDING|ESTABLISHED\] The actual status of an active session Session status is pending
Session status is established. Possible values are \`PENDING\`, \`ESTABLISHED\`. Defaults to
\`PENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PENDING",
    "ESTABLISHED"),
}
```

<a id="canonical-819e01569eec30e5fcd7f10a9a759261a80b431a9774aaa3dbc9b2c8678f139f"></a>

<a id="canonical-fa714c1f54b2de788506f50a79b9c2ed32b237565794b17d615c9298f9b9ed4f"></a>

## username property — Property reference / 101036385201 / 13

Type: `"string"`. Computed.

Username used in authentication of this session.

- [variables](data-sources--access_active_session--reference--group-001.md#canonical-fe9deefd6f8b4b60488c573dfc53749b5da80da187273355455bd00691eacb17): complete subsection reference.

<a id="canonical-be9ac6cb9faead13f8af8aac0703b089a2ab7790ef77a202fd2815a80b5c8c95"></a>

<a id="canonical-4b3351a42d8d3aab8fe454b8442944db323cb98d6b137725d2abcb4b66edd73b"></a>

## virtual_server property — Property reference / 101036385201 / 14

Type: `"string"`. Computed.

Virtual server associated with the session.

<a id="canonical-5a4436261edaecf5bf40923c30a07c6f75cbc5fea3636e35e529072626d0241a"></a>

## All schema paths — Property reference / 101036385201 / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `client_ip` | [client_ip](data-sources--access_active_session--reference--group-001.md#canonical-6854a1910303e8920b02c2769add2b8a5d905c7ab7275faaac0e92fda123ced8) |
| `expiration_time` | [expiration_time](data-sources--access_active_session--reference--group-001.md#canonical-f3d1e9ca4bb27dd69abbae464c0e9076a6be1726f37197931bd37e0f9a835342) |
| `id` | [id](data-sources--access_active_session--reference--group-001.md#canonical-c0ce9693ca644f253163196e11c9ebd64f262be01e9181baa3a79c2d789529d6) |
| `last_activity_time` | [last_activity_time](data-sources--access_active_session--reference--group-001.md#canonical-70a55109e86cf2765fde3b07614985770865341d49813c5a9353237b277021c4) |
| `namespace` | [namespace](data-sources--access_active_session--reference--group-001.md#canonical-eda95bf6959e7d32c4d43f58d2b16926a1c7df56e5ccce44b21726dcb9846cce) |
| `policy` | [policy](data-sources--access_active_session--reference--group-001.md#canonical-9d8dfc3600fb47793def86f139b0328879ef415040cf51d10344260e8d82b6c4) |
| `site` | [site](data-sources--access_active_session--reference--group-001.md#canonical-611fdd92522efc09d9716debce7d9db29ac80e53e12c904f7ff96d1171a47649) |
| `start_time` | [start_time](data-sources--access_active_session--reference--group-001.md#canonical-ba825e81a407be6727fb354ce3297d6ee4ff96931af46cae7dbbe638840500e1) |
| `status` | [status](data-sources--access_active_session--reference--group-001.md#canonical-0e8dd58a7ea34eb993da158f4281be138064ae947d2787fc3d48d5e432c86dc5) |
| `username` | [username](data-sources--access_active_session--reference--group-001.md#canonical-819e01569eec30e5fcd7f10a9a759261a80b431a9774aaa3dbc9b2c8678f139f) |
| `variables` | [variables](data-sources--access_active_session--reference--group-001.md#canonical-fd3751921795b67e3b0aa02e7a454c57197d9eb7c9aa4c4e2304ca3a30afe50f) |
| `variables.value` | [variables.value](data-sources--access_active_session--reference--group-001.md#canonical-91d6a05a457c88a56fce6ac0cbae2134e79a512057bdf8806aee6e73bf5a79c6) |
| `variables.variable` | [variables.variable](data-sources--access_active_session--reference--group-001.md#canonical-f07a11c876873f177b9eb0d6db61ab15a46aef2c24699c509766c429db14d6c0) |
| `virtual_server` | [virtual_server](data-sources--access_active_session--reference--group-001.md#canonical-be9ac6cb9faead13f8af8aac0703b089a2ab7790ef77a202fd2815a80b5c8c95) |

<a id="canonical-10326d02c28550f180bf13f79fa3935cd321f688883ac2363be0e8e14144afd4"></a>

## Next pages — Property reference / 101036385201 / 16

- [variables](data-sources--access_active_session--reference--group-001.md#canonical-fe9deefd6f8b4b60488c573dfc53749b5da80da187273355455bd00691eacb17)
- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)

<a id="canonical-fe9deefd6f8b4b60488c573dfc53749b5da80da187273355455bd00691eacb17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-150b18e5791ee3d8fc93563b6b7b0acb69262b2b6bb4ae61690980ad6e7539af"></a>

## variables — variables / cefeea73a726 / 2

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
- [Property reference](data-sources--access_active_session--reference--group-001.md#canonical-f9145d41000cc5377b88dd55e2dfd831d018e430e7e6108f72457e70d726428d)
- variables

<a id="canonical-fd3751921795b67e3b0aa02e7a454c57197d9eb7c9aa4c4e2304ca3a30afe50f"></a>

Type: `"list"`. Computed.

Variables. Session variables as key-value pairs.

<a id="canonical-4937de6a9f77fadf2a0541c03b44cbf87245d0cc30f15f6361872d81c2a18fce"></a>

## Direct properties — variables / cefeea73a726 / 3

<a id="canonical-91d6a05a457c88a56fce6ac0cbae2134e79a512057bdf8806aee6e73bf5a79c6"></a>

<a id="canonical-8ff29f4387bd0baf68069370946d51de6fcde5a13bd00e3c21bf6a702c2718a4"></a>

## value property — variables / cefeea73a726 / 4

Type: `"string"`. Computed.

Variable Value. The value of the session variable.

<a id="canonical-f07a11c876873f177b9eb0d6db61ab15a46aef2c24699c509766c429db14d6c0"></a>

<a id="canonical-fe9caf4135875f0d2fd0dd56c78737a93c06df08819845cfd0c8830a14f23b59"></a>

## variable property — variables / cefeea73a726 / 5

Type: `"string"`. Computed.

Variable Name. The name of the session variable.

<a id="canonical-f7c4032215b2d3d78f6c8715f3230b0c3b9e83fed93b4591369bcf0cfbe4b8ef"></a>

## Next pages — variables / cefeea73a726 / 6

- [Property reference](data-sources--access_active_session--reference--group-001.md#canonical-f9145d41000cc5377b88dd55e2dfd831d018e430e7e6108f72457e70d726428d)
- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
