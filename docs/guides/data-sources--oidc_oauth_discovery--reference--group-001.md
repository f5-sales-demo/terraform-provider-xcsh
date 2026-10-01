---
page_title: "xcsh_oidc_oauth_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery reference."
---

# xcsh_oidc_oauth_discovery reference

<a id="canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffe004706d3f4e4e5688f5356d4a3be4d5994c003f3b86ef63d5d49f07439bb0"></a>

## Property reference — Property reference / 34359263992c / 2

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
- Property reference

<a id="canonical-490c491c49e035deb271e09b8551528c31a8aefd6b12d3dfa0ca374edcd9a1ca"></a>

## Direct properties — Property reference / 34359263992c / 3

<a id="canonical-120dbfc0bbaed6d2e2a7c6f87dd16ca535546997490669440d4f52be87a1af73"></a>

<a id="canonical-6e601b4aae30964cf84a36c8ed09b3de78e1bf95003c8d0f0d8b42f372a58e74"></a>

## authentication_uri property — Property reference / 34359263992c / 4

Type: `"string"`. Computed.

Authentication URI. OAuth 2.0 authorization endpoint URI.

- [auto_jwt_config_name](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-07999a827444bdc5e60edfebf8d1b6a68afbebb4302f8b5f4cd09c29cea9ab7e): complete subsection reference.

<a id="canonical-4a1b952d9c80424641ccbf5b8c2bd5c7117c5c5d05b7afcb0a1681dee822ab51"></a>

<a id="canonical-a0d316b66264cbd0079fbb7f3e28ccc1678a8c07ea0f07feba53958a1031a91e"></a>

## introspect_uri property — Property reference / 34359263992c / 5

Type: `"string"`. Computed.

OAuth 2.0 token introspection endpoint URI.

<a id="canonical-a2946676d1cc3ba35e6723c8361fc6b5b273e6702971ba9de739157828ab3a6a"></a>

<a id="canonical-afb521e507d608bd47e3bdab8f0a040d3c2a17bdd5ab0ba67cc0ad747736579c"></a>

## namespace property — Property reference / 34359263992c / 6

Type: `"string"`. Required.

Namespace Namespace to scope the request.

<a id="canonical-1bd4aec2f9d251e61550ad107c1f33dde0c3c050ffc0c2dd1a810b7d61673fcb"></a>

<a id="canonical-2906a31715088c9847cc1dcc12ee9eb6e9d1173a7cf8b725837d2e4c0addc772"></a>

## openid_cfg_uri property — Property reference / 34359263992c / 7

Type: `"string"`. Optional.

OpenID Configuration URI. OpenID provider metadata URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

<a id="canonical-a5c5682f4f12ca1591adde67be93132932686be0bc454f0a7466fb50097bdbb8"></a>

<a id="canonical-5236cc0f625c64ebc79bcac2846d87f6db2cfbf6e06098bc0014775a569a67dc"></a>

## token_uri property — Property reference / 34359263992c / 8

Type: `"string"`. Computed.

Token URI. OAuth 2.0 token endpoint URI.

<a id="canonical-811a03a4f7f18600d43d168d8c9b762d36199d820c2e4fc122abed78dbda87e8"></a>

<a id="canonical-40606342d2ad47b8c4df0e700b938567867d29e18451698daef7fa72d8160063"></a>

## token_validation_scope_uri property — Property reference / 34359263992c / 9

Type: `"string"`. Computed.

Scope used when validating tokens against the provider.

- [trusted_ca](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-7268e3a8cae2c61ee1069c59780b9e4e8cff938574fbd22d2ff3e18173c79117): complete subsection reference.

<a id="canonical-e8c4f443fd5a5ec887a3a1f24de6fbbe7ebf29fd61574e68049cfaf844d45357"></a>

<a id="canonical-6b93ac2c33e8e072db4b944ca8e0395a1ec422be915131886a12e070289d601a"></a>

## userinfo_request_uri property — Property reference / 34359263992c / 10

Type: `"string"`. Computed.

UserInfo Request URI. OpenID Connect UserInfo endpoint URI.

<a id="canonical-c296202ff3546f807d7d9f8b469b8239d8aec204134a2767af080799f65f74e6"></a>

## All schema paths — Property reference / 34359263992c / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `authentication_uri` | [authentication_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-120dbfc0bbaed6d2e2a7c6f87dd16ca535546997490669440d4f52be87a1af73) |
| `auto_jwt_config_name` | [auto_jwt_config_name](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-7b922931a8c698fb4a1de58ce5262f7d544bbce657054748e28f694061f5d930) |
| `auto_jwt_config_name.name` | [auto_jwt_config_name.name](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-bf616d1576e8e9ab46b3f4cb6af9febe684e48f7fc436822911535eb8785c1d8) |
| `auto_jwt_config_name.namespace` | [auto_jwt_config_name.namespace](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-19cf4b15cb8135e6b68d7e17325d9d003a0f8c9044d4983a5ef7e35241f60715) |
| `auto_jwt_config_name.tenant` | [auto_jwt_config_name.tenant](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-a36e1603d7cc44cafebfc8223b30e61a8a6d8ea4d0c820b5671776f05db4208f) |
| `introspect_uri` | [introspect_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-4a1b952d9c80424641ccbf5b8c2bd5c7117c5c5d05b7afcb0a1681dee822ab51) |
| `namespace` | [namespace](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-a2946676d1cc3ba35e6723c8361fc6b5b273e6702971ba9de739157828ab3a6a) |
| `openid_cfg_uri` | [openid_cfg_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-1bd4aec2f9d251e61550ad107c1f33dde0c3c050ffc0c2dd1a810b7d61673fcb) |
| `token_uri` | [token_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-a5c5682f4f12ca1591adde67be93132932686be0bc454f0a7466fb50097bdbb8) |
| `token_validation_scope_uri` | [token_validation_scope_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-811a03a4f7f18600d43d168d8c9b762d36199d820c2e4fc122abed78dbda87e8) |
| `trusted_ca` | [trusted_ca](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-52897a8a086950abc23584d82bafac11257994b4f086f223113420773ece9785) |
| `trusted_ca.name` | [trusted_ca.name](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-d09583797dfe517b33b2007abd086aef2d28735af136c1b23273e8445bd90091) |
| `trusted_ca.namespace` | [trusted_ca.namespace](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-a4c98fef88a3cd03c684b6f4470a9760d4c33291fefc59ab1636ccc25dc7867e) |
| `trusted_ca.tenant` | [trusted_ca.tenant](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-845d073fe2160b9a9cc32ca560cc7e94538f1420132b50c52c9cedd63a82fdfe) |
| `userinfo_request_uri` | [userinfo_request_uri](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-e8c4f443fd5a5ec887a3a1f24de6fbbe7ebf29fd61574e68049cfaf844d45357) |

<a id="canonical-f96a54c768f189ffad28d2576557c00e3383be656b35f4657af359bca753544c"></a>

## Next pages — Property reference / 34359263992c / 12

- [auto_jwt_config_name](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-07999a827444bdc5e60edfebf8d1b6a68afbebb4302f8b5f4cd09c29cea9ab7e)
- [trusted_ca](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-7268e3a8cae2c61ee1069c59780b9e4e8cff938574fbd22d2ff3e18173c79117)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)

<a id="canonical-07999a827444bdc5e60edfebf8d1b6a68afbebb4302f8b5f4cd09c29cea9ab7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1066c5390f76cd102086df44bfd0749f5b86833723d488cdeb718384169bd684"></a>

## auto_jwt_config_name — auto_jwt_config_name / e1a0dda32a29 / 2

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
- [Property reference](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a)
- auto_jwt_config_name

<a id="canonical-7b922931a8c698fb4a1de58ce5262f7d544bbce657054748e28f694061f5d930"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

<a id="canonical-770bba5c162e20f8b97de989563cb137d30edaf08fcd0573d2a205ea45434ba3"></a>

## Direct properties — auto_jwt_config_name / e1a0dda32a29 / 3

<a id="canonical-bf616d1576e8e9ab46b3f4cb6af9febe684e48f7fc436822911535eb8785c1d8"></a>

<a id="canonical-ec26672e6da5f51929377c3c4118b6df39f47594060181fd965d2359171bf98b"></a>

## name property — auto_jwt_config_name / e1a0dda32a29 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

<a id="canonical-19cf4b15cb8135e6b68d7e17325d9d003a0f8c9044d4983a5ef7e35241f60715"></a>

<a id="canonical-69e89daa56cb577909eea1185cc008f19f6bcf46e4a12b5e33d6f13d1aa8eb1a"></a>

## namespace property — auto_jwt_config_name / e1a0dda32a29 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

<a id="canonical-a36e1603d7cc44cafebfc8223b30e61a8a6d8ea4d0c820b5671776f05db4208f"></a>

<a id="canonical-42239e2febe64707da4350646223630a75321aebc4c79a62e1a94112a46fea2d"></a>

## tenant property — auto_jwt_config_name / e1a0dda32a29 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-684a1a6eb65d2875feb51ea9544eb808621afc1a9d982c1995a1e1490b84d554"></a>

## Next pages — auto_jwt_config_name / e1a0dda32a29 / 7

- [Property reference](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)

<a id="canonical-7268e3a8cae2c61ee1069c59780b9e4e8cff938574fbd22d2ff3e18173c79117"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8093fb077cf32397a2f9a1b5a2ee572d7fd0121af53831ef7adaafebce41c15"></a>

## trusted_ca — trusted_ca / d1cc7e0d5f5e / 2

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
- [Property reference](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a)
- trusted_ca

<a id="canonical-52897a8a086950abc23584d82bafac11257994b4f086f223113420773ece9785"></a>

Type: `"single"`. Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

<a id="canonical-029b8886a133b1cc547ecf8ed0880b3f3bd2510e9cb892fe2df41beb181a1131"></a>

## Direct properties — trusted_ca / d1cc7e0d5f5e / 3

<a id="canonical-d09583797dfe517b33b2007abd086aef2d28735af136c1b23273e8445bd90091"></a>

<a id="canonical-b6cdabfd2bdd04002a3af6c87f553a8fa0346815be3ebb2dbee583bb2ed96741"></a>

## name property — trusted_ca / d1cc7e0d5f5e / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

<a id="canonical-a4c98fef88a3cd03c684b6f4470a9760d4c33291fefc59ab1636ccc25dc7867e"></a>

<a id="canonical-c5b9b093c4e196e5cffa3cd9b2b82ebf49e7f2db238e9a29afe631c3ddb9a133"></a>

## namespace property — trusted_ca / d1cc7e0d5f5e / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

<a id="canonical-845d073fe2160b9a9cc32ca560cc7e94538f1420132b50c52c9cedd63a82fdfe"></a>

<a id="canonical-87810dcef9218680031a15c46e697007355baec9d78465eaff29971b60961545"></a>

## tenant property — trusted_ca / d1cc7e0d5f5e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-bcdd2d2d6611849e9b5b7c359b3aa80eb44048d0cf14c385c829ff27043839b7"></a>

## Next pages — trusted_ca / d1cc7e0d5f5e / 7

- [Property reference](data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-8080b48cb0feeea357a5116a2e8d2f87213d62c9c935f7da9a09cb40d2eabf8a)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-6699f2a062d3230c5b4e6d55d20704f4124103bcf43f441267242ac4de16ca5a)
