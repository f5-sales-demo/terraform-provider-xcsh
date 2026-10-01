---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-381657b3480437c10fcf316c0afe71a07b20c353c9a6e7681ccb54fd7764d96e"></a>

## Property reference — Property reference / a744dbd1f731 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- Property reference

<a id="canonical-8578c465ac663ac909ec49a7def1afacc2a360d62ef296803c03093d877bc125"></a>

## Direct properties — Property reference / a744dbd1f731 / 3

- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac): complete subsection reference.

<a id="canonical-b70e0e787a77ad8e5b3645fe2cab3f3ad4ccffa412f0395ee7ad9afae9cabdf7"></a>

<a id="canonical-ca1753ca04a8855564f2699c302ace497e6a23a26cabee776683652fb3345d41"></a>

## annotations property — Property reference / a744dbd1f731 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b): complete subsection reference.

<a id="canonical-eb585dd6b063e58fbb8e1ed31413368394afb9f4cb7ad81be70fcdba2c39d4c3"></a>

<a id="canonical-66ef00c60bcfbf5b20bd73d1c9c7e83a557f41ae2e68bb781b05a780e9f70941"></a>

## description property — Property reference / a744dbd1f731 / 5

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-fef20dabcb2ad2e829c054a80cc77087ca35a1b43924717958ad9cbfe13ab91f"></a>

<a id="canonical-4eef2e91585ec20f1ab8d9402e565de376fddaed28ba755fecd4b6e336806fe2"></a>

## disable property — Property reference / a744dbd1f731 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-57513081a09b8abb9a60bc3abfe19d773b66067615be1a8de1c42c5f20a01415"></a>

<a id="canonical-fe2d21c0f70b1467d6519535e7e4ff5ea2cdf85e79617c0c02bb8d0c5b713bf0"></a>

## id property — Property reference / a744dbd1f731 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-58e1feceba1cc5dc0e29099842965bf040dcdf9ee383ce30502e44581d236fa2): complete subsection reference.

<a id="canonical-a94767af6a33ec7df337d8f8c8901bea4a37f1ed3f37a6fac191f73cd9dae4f6"></a>

<a id="canonical-2df1859b8ef14ca9e5beeb125b8939f87ae11093fc35935940dc3628c62eef11"></a>

## labels property — Property reference / a744dbd1f731 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-04bc9f74502d8ad52eaa52994e94b471c36cc5c7eca34734297c5bf85ef6556d): complete subsection reference.

<a id="canonical-b752b61bbc2d3c9004349d5ba9375e42697b4b70c65feb589a8ec7e25c4810ad"></a>

<a id="canonical-9787731fc664e1b4b7c014ac111bf2deb0166c636ef5256a9e9e5b2796a53244"></a>

## name property — Property reference / a744dbd1f731 / 9

Type: `"string"`. Required.

Name of the BIG-IP HTTP Proxy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-e84d49a199aed817223bae39bb08099aefcf9207d5a924bf6b8c1efb6946a8f1"></a>

<a id="canonical-af20a052e6196121e6425ae1c2158ddf5e3a71d52b2a386b31d157b8c8a2b48a"></a>

## namespace property — Property reference / a744dbd1f731 / 10

Type: `"string"`. Required.

Namespace where the BIG-IP HTTP Proxy is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2): complete subsection reference.

- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a): complete subsection reference.

- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf): complete subsection reference.

- [timeouts](resources--bigip_http_proxy--reference--group-004.md#canonical-6dfd33c696868d9aaa2f87009aa602ef5a1be81177bdfd35ec1953db11750b92): complete subsection reference.

<a id="canonical-98827456a4d98fa2a27684b4c416c68b9b5052c240eab12aafac30b15e5d1c7b"></a>

## All schema paths — Property reference / a744dbd1f731 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-a245461473faa06a1136f76879ec1ae234563a4a198520f8da9c2f73fbe9dbf8) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](resources--bigip_http_proxy--reference--group-001.md#canonical-ca985752719eacf36f287bdd6e54895f0a7d113043d942de92139634fc340fd9) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-21dce0387b5edbaeb5e2c5ed61ae1106bc6b28aa78ca965e5d81816090cd87ad) |
| `annotations` | [annotations](resources--bigip_http_proxy--reference--group-001.md#canonical-b70e0e787a77ad8e5b3645fe2cab3f3ad4ccffa412f0395ee7ad9afae9cabdf7) |
| `ddos_profile` | [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-d9c0552c6bc9f1722a230fd8dbe5c3ec5638749dc0027abdd3a3aa705d057867) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-02b0ef95a035fe65d28453b824caf35283ac9c2311bb9354186fe42a31533e92) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-9a3dfb04717900fcb908a8a0f8a473c8410e707d740cf4de4cc552170f213380) |
| `description` | [description](resources--bigip_http_proxy--reference--group-001.md#canonical-eb585dd6b063e58fbb8e1ed31413368394afb9f4cb7ad81be70fcdba2c39d4c3) |
| `disable` | [disable](resources--bigip_http_proxy--reference--group-001.md#canonical-fef20dabcb2ad2e829c054a80cc77087ca35a1b43924717958ad9cbfe13ab91f) |
| `id` | [id](resources--bigip_http_proxy--reference--group-001.md#canonical-57513081a09b8abb9a60bc3abfe19d773b66067615be1a8de1c42c5f20a01415) |
| `irules` | [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-3011b3314c93539cefe83e9ecfe627c8bb8085fb4286f6248a6b2b76bf12ed63) |
| `irules.irules` | [irules.irules](resources--bigip_http_proxy--reference--group-001.md#canonical-aa005ef36d45e63cce4634ada13cc58857e9ef184641eddc2bc79cfe7efdbc0e) |
| `irules.irules.name` | [irules.irules.name](resources--bigip_http_proxy--reference--group-001.md#canonical-8d254192484074b5db9678bfdf3f61e29958720f1b2344b7e549bc5c8c20757b) |
| `irules.irules.namespace` | [irules.irules.namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-b7649b35116c034ac8ecf89f043fc56bbe3e37c310b7aa2ecdaf57766fc72b8f) |
| `irules.irules.tenant` | [irules.irules.tenant](resources--bigip_http_proxy--reference--group-001.md#canonical-769435ceaf302afef8f574b0f897a740242f72ec763fc7501b936b3ccbe007cc) |
| `labels` | [labels](resources--bigip_http_proxy--reference--group-001.md#canonical-a94767af6a33ec7df337d8f8c8901bea4a37f1ed3f37a6fac191f73cd9dae4f6) |
| `lb_algorithm` | [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-639dccf092030b6ea115b031644580ffd2cf87e0e5f57883d719c422d4791da7) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--bigip_http_proxy--reference--group-001.md#canonical-48e56275efa603b505163573abcdc96ea9355528a4c19ea6603802abd939664d) |
| `name` | [name](resources--bigip_http_proxy--reference--group-001.md#canonical-b752b61bbc2d3c9004349d5ba9375e42697b4b70c65feb589a8ec7e25c4810ad) |
| `namespace` | [namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-e84d49a199aed817223bae39bb08099aefcf9207d5a924bf6b8c1efb6946a8f1) |
| `origin_pools` | [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-bff420c9dcac9003b1b1cf86d1f48c359ccc36a63f3ffb709f26633f290ed7b5) |
| `origin_pools.pools` | [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-fb42e8e2a33adca0af5e42c56a9bf6bbb0410ec6873cc6e1ed31bb4790cb8512) |
| `origin_pools.pools.name` | [origin_pools.pools.name](resources--bigip_http_proxy--reference--group-001.md#canonical-b58c37206966132e098899fa2b852838e54668a96401bcfbe21347c28ac6df67) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-e5eddd711f418f181d88576d74b8b61708baf248384b3b50eeb1af868c11c168) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](resources--bigip_http_proxy--reference--group-001.md#canonical-b8bca6fc5a390d96735ad200ef19553308437b8ddc939bd1f26b8b9ed52e74bf) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-7d4e468e6a58209e4488ac91128697fffed299f88ed07c12035b3648dc972dad) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-9faff68e85044bd7e44dcb09348b77bb07ed49854faffab26101109ed2b8ed24) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-9198610961bd03ba004dd63f8f9ae5853e65a84e703ffd8ca2c7d6bf2be97312) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-d12882321f19a1d130f70a701dfd85de832f4ac5d23a05036a5ba3b5dd93ffdd) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--bigip_http_proxy--reference--group-001.md#canonical-5a40146f78fe61d13862ffc4f960f24bd6f3af86caf0ebd825108addac4cf2e1) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--bigip_http_proxy--reference--group-001.md#canonical-c39bce1b15e838ce934ff946f0e430f395a98f4d00cc42ae7eb28511cd26e589) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](resources--bigip_http_proxy--reference--group-001.md#canonical-8cd6a338af752a00b65cdcc2ed2dfd6004022546ddb72e7f770f9328eb1f6d25) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](resources--bigip_http_proxy--reference--group-001.md#canonical-48b34c416e0713d385e1f8a366985834a08e520cf721a66c1362702c9c612952) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](resources--bigip_http_proxy--reference--group-001.md#canonical-8fb39008ceea2e5e150312389b1cdf44549b56d4ce6a8cd378824dac84a85611) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](resources--bigip_http_proxy--reference--group-001.md#canonical-319757f1a535b8cd2f17d9fe83829a558ad8014507c899d313bf25e379c154da) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](resources--bigip_http_proxy--reference--group-001.md#canonical-e2bebbf461d5873f564360e7923c2176220602f995efcb184a252041127013a8) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-05010ee3dab4c54633a80d65f9fdfe74991299190f45e31127ed79f757a1d6e0) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0300907df319b13dc6850d94b60d71cf2d7ca28baf7de43ce1eb9ca88f0adefc) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-829d9509735656256b1cb9588c26b8f8d14f2c9ddaa4bd60d5f62883758c1bc0) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-97e39fd170317efae818de0467542269121a32df3f193c9234381d6eece89205) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](resources--bigip_http_proxy--reference--group-001.md#canonical-755a88bc018d114a6e2ba6f2bc7962ed8990550d284d74240ae9030f221e7029) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](resources--bigip_http_proxy--reference--group-001.md#canonical-9584ba3795b5480ef8ff059b23d75c5cfc0b5a7baede61b17ab168ddc22bab14) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-ab165180bad9d4fb6a3ecb23bc399c6f7536265d36b89f2a2a1bd71efbcdd811) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](resources--bigip_http_proxy--reference--group-002.md#canonical-a9d6bb5892dc4beb3374928e29301286c7cb69da7209dd4066d1b33aa302224b) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-0b9ab319993f855548bb990cf286da4c05ec466c9059ca5c97cd1c34f0a32e83) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-115c8fe7abab8571863e76d23e85732b039124d53a725ca8b7d667fddf5b3964) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-f457abe20734855ac56accccb848bc26fda708514d5de6b27c51a25ccee68439) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-fddeca3f7cf95a4b8e942237a70f40ff097fdb5065641a24aa224bbd7c5bdfbb) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-2b0ae0e8a47845973d254214c28ffda0ffab5297aee18c33af6bfad368f7b210) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-daa92cc01f797073c0bbe205d6146642ae0fc7ccb99f2848796f908f95f0c376) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-09b6cb57709790ac1201489204d788c7d5a0228cb29b2c5606af95fed6c5b605) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-64346720c3aadf3f1f83329d6bad6497e49eb59a7d0bc1bbce567e4bca59e49d) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-9be5f87c5129410aa601c770dfd25c9d0c7f98b18ab2775e182e97acd2608035) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-e5c12022e34775fd3d0f4c22f30cd8a1831facc8afe318e55a1f726991a0c3db) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--reference--group-002.md#canonical-f2b4e68076cf3d87110ff70a8c17b04d8aca23b3bb6150393bfb295c6bf5cc09) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](resources--bigip_http_proxy--reference--group-002.md#canonical-257cd1a085a52e805633de70db18e2b6c5b5a9d48666633464b5398a93a74c8d) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-b9ac5037b66a974ad7fc7d962c0fe41998e9c2a83efb44391f15f340aa107adf) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-2f46a42df25a722dfa0edfdadc6165a5aced2898e732967c9d960747022d0111) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-4abed645f474da3a75ab754c019735bb512d5b52160e35355c762f248a8e91a0) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-92ad27a4a3c72d98a62b42c3b61fc652b66c6e89d3d9e7c065ad74d70c47fa88) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](resources--bigip_http_proxy--reference--group-002.md#canonical-bdda455da5206e5503684792d231c5f2f101d314497eba0d56468c8802270a11) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](resources--bigip_http_proxy--reference--group-002.md#canonical-182801bf521eb4f42f7cf6e4b9857fa706e17327292156ae16e3b1b59b3bbed7) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-db0eca5ff7f3c026ca902c95d1424f2a7b36789ce1a1231a4d27b6da73cff35c) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-c073d1e92c0fe2ffb567fc8c04c8682fd27b2dcb2acc2da11e0e39b90576ddbc) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-2fa618c80590a083fab2b84cdf4b3482b652141a44afdc1c5db3bf39469ee61c) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](resources--bigip_http_proxy--reference--group-002.md#canonical-6e66afcc6da1fc458207cd9ecc4437a9b4ff31ac3abde9b727aa43e6fad60b59) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-daf31fb7f8057b496ab5d2fc21daf0c3c46f8eae97ccf04a806b20d53935ff41) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-e6ccba1a7698c1ac89c41b89c89308410a97f447a85c7fe6d6e0444ddda3ccaf) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-e1f0014e870f11d422a647b8718fc8a9620ccb24baed50b583193cd35c39c351) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-7c65e5a6571b7045c585f2eb5c46d7c6207a38cabc134af10a05f5ac89921f5a) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-6cfa1dc8bc409e5f7dcf443cfc4ed85cfabec60e7c490ccb4eecd3c7779d62da) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-560b092e51cc06fbf452526a7ba134d21a68668fa7da6dc5659f6ff6a3bf12df) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-0324e4d62ad11c626d38811dc4091490ff9f3feab130579c6e0b27a55a3b9092) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-8e05525a92746aead25ca2a74f1ddef662dd274b6e69946f8f0a01dfbb8f7d82) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-934bf957d3afc37edb659a6c8cf20de0fef1608f2d55b555a812ee997644bfcf) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-4ef35c03f7e343139ba72f3f8dcd7e8d2ec047d599aef0c29e4df287c9f50a05) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--reference--group-002.md#canonical-5b54d399f8eefd2929d76c5d9ac9c31574677f57814c8dd8ba50c2c273f2d7a8) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-84acc5e32f941a19084cc6863a24cfc32aaee1694e1539008fff3cec4e7a812c) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-7a78c119ffbc8370c6319bdb971b5a231774c04c267c5c68b325b93a2e7ee966) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](resources--bigip_http_proxy--reference--group-002.md#canonical-f2c5a0876cdfafe0e2e9580568fe6ec1d1078b40d96520fa7346f612ba09b30f) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](resources--bigip_http_proxy--reference--group-002.md#canonical-b5f87f4d37b2c3cd7cca75bd4a6ced9d2cb14f760b209eb39272a1c7a74889f8) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](resources--bigip_http_proxy--reference--group-002.md#canonical-1be4b21479e9947e2bf2534b9a789c78899768094e242198fc02677069f58cff) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](resources--bigip_http_proxy--reference--group-001.md#canonical-a265d272f1fb14df5948f28f9b982ee56fe8028cd1c27588bea90d66e7dd0aa4) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](resources--bigip_http_proxy--reference--group-001.md#canonical-8fb27a6dfdd8d9e31450042c56ff5adb75025ed0ff8ce5bafdd59ef470f1a232) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](resources--bigip_http_proxy--reference--group-001.md#canonical-5589fb70691586957938c5d1ae461275c095bda1093a3817f40ffc5e17513713) |
| `proxy_advertisement` | [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-de0220efa04dfbaddb3b33e08a1cb7605d9751c475a3eb675c2d280db4a3f28c) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-755582f63681a94aa38c421dcc269820798375b29da54890002ae03df0131c05) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-9592b4b1df16d17e7db73b173ae486f7884f5145c480e5cb7bc4f0105800f3ed) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-182716908c40545b1d73f0eaee45aaf610773bc5c6012d9313051b445e407bc4) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-5d6eca053f9283a4e7e45a30d0cedfb69092c8484976e083df7581389e9b6909) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-97d0882e0e816e2cebf956f53f5a2331b06364435cd0367a1b4a4493731334b9) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-5b6595d513bddce700b6822cf86ee991e00e5f3ed7f5264ef471e27fd043fbde) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-06a4f2f060c8353300daf94c5b3fe7aefd80ee9881dc81af18a4178bfc5840e0) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-11c923eeea590c7be6948024b83ac6d61182cf95ac2b1338f54425a532381dff) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-fce3fe0e31050b884b6b8ab2b6b87e4a0e6a57522acfc4437f23290872951611) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-e5c66680a5cd48a01aeaa4ce3ffa0e8d38ad3bc3b719daa57899dcf77c53a28e) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-d54c1a11f451eed0ffef1c6de8046cdedf115a2cffc840d03f849615d0548f35) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d8eb69bf116aef9d7a9e3a466db64efe1cd1fd4cde36499ff86be5a0cf6db7) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-39f0be748941fc91cdb2279340e238968721da63c3267e6b3d5f26cfc198f541) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-aedc0702d3ab0c5ccd461b0ae45cda6836640d6111e16f81ab869b04fc858e42) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-4c69ca6a5bb663e799aaf73145095d3eb6326c731ba7cfdb9543c423db998d06) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-9ab86565a10ac424d7e020463991265fd6612670f4dc4564baec29a0f187bc4b) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-4de201bc27912c9b172e53d701f213be2a7365e3180f08afdd8677b913f76ada) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--bigip_http_proxy--reference--group-002.md#canonical-315969f14aa8272ade8ce195fb24743cfd0d8263fbb6b58f7e94586e18af758a) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--bigip_http_proxy--reference--group-002.md#canonical-7f65c3cb0f8f4ca7b68404819b0e1229c4bfbfe31618c6b65e026697a5065c7d) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-1c69909f5aa462fa6fdcd880c074cc2d3d08596a63dcd29bbc4e83aa94692c4c) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-7db859a85787b73335e8433014cedbfb0e7dbc7fb8a8bf738dd48d99dccc026d) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--bigip_http_proxy--reference--group-002.md#canonical-dac11af03fd7221b8ccbf4bc5cc93e8cabfa1d8ddbf1a498a431d9c9997925ca) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--bigip_http_proxy--reference--group-002.md#canonical-8738bd360de97dacfe818b46010911db1020378900683bfaa059ecde291adc63) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-2115b2293c3dd747abb667ca9e3b62e4f67bdd8f7f7cbe39947c3ed854015e57) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-09fe766955784b930329d3d882f176b1a9ca7204662ce2fb0df1bf7d32bb52b4) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-9c538802947c2587b8f2f31ddb98510ab2e31b5a87430dec3df9063d31605f98) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--bigip_http_proxy--reference--group-002.md#canonical-87a3cb118377001c4680bf117fc15155871afe6823cd03a9b9c59513c0e93eb4) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-80246c1f2b956d0fbd276631b7798159d7ffef977a200e0e20774867acfb03cf) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-cd4b92f103be0b6690a75828490935b59978f28f516fd5c4e2199506e7ec5f53) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-bf6f54fa3f0a357a84dcbe813d6ab4c06ac120660f839be3df74564e7cd7f7bf) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-ce241ed070877134d2377eb857388e3b4a5a690c3ab3f970af5bffdcc46cd54c) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-348c7c6e0a48e5dcf06d8b5c3824491b2d680a8a89bc9f3e5895bb8408549ac6) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-4a353091edc4ee2d679bde8f5a8d2eeb4168e33fd43c8416796a0fff222aa652) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--bigip_http_proxy--reference--group-002.md#canonical-380824253800a008eb5499764dbaea5d6dbae4489cdd441e5d33c14607230072) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-dbfb66b43d236c1324fac70758d1b05222c4283a43bdde5b1ba908555b330589) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-b962ea151483ea8f1b4ea48b1dd8defa0a9bdc4177fc2c22c29b69359ccfd413) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-dbf4770aef71706cad2881037199031396b77850bbb3c0d3abde75faa06e21fc) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--bigip_http_proxy--reference--group-002.md#canonical-de7f32319bdd3d7669f32d512496c84ae416e1676dc9671628f02d40891a32db) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-e63b8f6f148c1543675f59fd0c9c8907e1943c821ff18292e374ddb54d0da10a) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-bc23ac8f1f16d1389125dcad15c58a66f91af38dde7a98e80b7229633d7c3eac) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-88680d258b39fff6410521f73e464ce14481dbefc3eeedb88ed6cfa95216dac6) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-3bfaaf47181961ac99edbb9a6940af02b39632c8f97d822c24941d66a5d732fa) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-28c2dde4fed71e5ccce17e37f0d124cdc3d81f509b5f2a8b805fcc3d769cf459) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-cbbd452b1618d3ebdadfa020da4bc2debc59beb706240c802df23cc33467052b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--bigip_http_proxy--reference--group-002.md#canonical-85449426bfa455c0af996a0532d936ce6809d27583133c2f2d3cf82b254cded1) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-4206592ea7b27b2996e52d3463b2a535dfcf79fc80a9c82abcbcf7ae820b555b) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-fa8c99780e2f0d51918945b36bcaf8d74df7281cf5f33a3e725790003341b5b5) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-857b421813c57394af9a9ec74c2a0ea55061971590ca925def86cc7adbf98395) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-93c6d422bd06b7fad4df483fa6253dd112fe4acda4df4bc966a59a78d143fb4c) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-f10cd8e87c0cc3b92f2f267f51611f15c31703bacebec7c68af99e68087bc24c) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--reference--group-002.md#canonical-3418d5a3ab7b68e301977b1daab521119ea3863ffcf6e790e868b29240bfe174) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-fe67a041527f98c67e9cefcd9284b4bddb47925a250c424e4914b8387159c6b9) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-a2ab1b5194e4582e61ebb41ced16ce8f39eb84e6a884fc4771584d52fe59902e) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-2d799fba917c211ba6ff4df94d8d04df88c29164838aee9968975cde1d73f277) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-a5d8c5c3598c9a72dcd703a621be8e42c70dee85707ed16b308415c46658101c) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--bigip_http_proxy--reference--group-003.md#canonical-8bd4f4c16f8b5a7674a4faee079e7139b6e52458198524060d6645c532824b3a) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-8669ae18b658509fc3e58c1567b0b9f8e0c468ae3b2900611ab306a02e201d87) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-38982c279fb87666a347dd2f6327e5f5b8e50904a6fae3bd2e8aa4df7f9e33a6) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--bigip_http_proxy--reference--group-003.md#canonical-48c812065580a2ea07b05dd8c53411a5e8a56996f7519141fc161d47ec6fb1e4) |
| `proxy_config` | [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-393817bce832cba65ce788bb455b93d2361a3acf7d70c0c209d16ce7fe906499) |
| `proxy_config.domains` | [proxy_config.domains](resources--bigip_http_proxy--reference--group-003.md#canonical-3b37bc515b8a8d14272c48d565783cc5ff6c219707ca2c5f4f09ea8c71a80009) |
| `proxy_config.http` | [proxy_config.http](resources--bigip_http_proxy--reference--group-003.md#canonical-8fdb3823d250b342a9c3491ea57433fa1cfd2500dafdee5de3bbd686e65f0fb7) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](resources--bigip_http_proxy--reference--group-003.md#canonical-5c57f3767667a2a27383c68908b9fab5a911285d6efc2b104c9d64e053bab797) |
| `proxy_config.http.port` | [proxy_config.http.port](resources--bigip_http_proxy--reference--group-003.md#canonical-0ef545bfee9afd0dbe5e8efc58d11fc4e0f9ef0223eb5bc7064ba88ffbae689d) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](resources--bigip_http_proxy--reference--group-003.md#canonical-96028dd8a68d42d01baa5b1c7160ec78ade35f0f109aa51588b8e5784d0e01af) |
| `proxy_config.https` | [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-d7ba2d6b1c65e156882a61906e192ab13cac2784851176fbfe6470d06d3748b1) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](resources--bigip_http_proxy--reference--group-003.md#canonical-05bc32e787c935e6e0af7bbb56b42607056d87f02cfd0b2d5c751a90834d9008) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](resources--bigip_http_proxy--reference--group-003.md#canonical-3c03a18964c054e9682e6280faafd9fc270cb4fd5733ab6f186d75568031cc0b) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-5b613686c29e6e3e67cea41fb7f261da9c867d208b9de70de63ed947f3b85d1f) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-578508cb3a88168d0287502e5811cfd22662681e12c9ce24c39b74ba8d74886f) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-73aa8b2e7167f857e5a481420a7ea43aebf799380efb952b442f5b377d7ce051) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](resources--bigip_http_proxy--reference--group-003.md#canonical-ed67099a35b039de36d96b1d7597e070ac6250adb46254ae8d1f587a9109a00c) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-bf3f08497c044314c76f8983ad5ed14cc85b29a25b4f6b43631c6f2a78c76f14) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-358e63cdf62cb159386bd96955df9b390095f9579237eb2845c7fa226f53517a) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-ebb206f6f8d94deed75db42ebcc3754b2b04183787c284207cdb3a5d0fba4624) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-c2ad5915f33a2d7f830ec3f2300d94ef7900c8c44abc4140519529526be9082a) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-824bcf954b96ad918b0c9af7e2aa4add26fefe25847a4b0274203c9cd6c604f3) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-9c27a040f5225d7406699d05ac7f11128ab725b2a8c34d870d737abffc30d35f) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-00f5e5b19490e694c36965a009c8dbe63ab9aa9a2d116ad761306fda3b3145ca) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-c96a1ff50b5a3fd8703a608d205b987855f3b5fd867d2e81731609fcee3311ba) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-ff1c876852b8c36ed77296e3ab06229569462443d8ca686db332a24ed16657be) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-080bd63feeac5d47c6e62e47db22fbe113f2aba65c087135d6e9a1b8bd3ad922) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-7a5e1e3908558626883b41a15505a2240bd9f29749aab03d17744bde520fad26) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-31469aaaf9db9f1d864fa0b93a13e2ccef31142c48de49b76042d2c4a598e70d) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](resources--bigip_http_proxy--reference--group-003.md#canonical-45b3acd581d120085b3fed2d4ce43f254c9027c201499eb83f685b61869e7383) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-d91791e83ae838d417e3f509e6f03d1c23682b0f099288cc1953ff13f021839a) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-90cd7a4ee5dcdf12446444d1f275527772c139b2914c0d412bd3c090318590dd) |
| `proxy_config.https.port` | [proxy_config.https.port](resources--bigip_http_proxy--reference--group-003.md#canonical-77a526a48086f81d72a0b81a2040624b091efd2a1bb2986bddcc79b7df36b1b8) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](resources--bigip_http_proxy--reference--group-003.md#canonical-dc2102357c1f30f1df1ca5dea847ba193c5756a883087c1d10eb293c611f1de8) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](resources--bigip_http_proxy--reference--group-003.md#canonical-cdfa91fd445288e4601e1918915db6081b9357e3bab22d712b70883744a4eb9f) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-2ff80792778940a864793bb40570db1a909590583c99e208bd885c4414d38a3c) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-bd4999eed78144989cd4db9b0b74a387f4521b12d9f2e02a562e374cafd97805) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](resources--bigip_http_proxy--reference--group-003.md#canonical-c872796d1428e512135213bedc145496003aa8031b796c4a39df40d5910ac789) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-a9f92a396f4cb8c05e6b2c7a0ebf42117ff79b7258ac403eee0d8558677fdc50) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-74a247680179575538231b6ec156ddaf4338640e88b3f6e868c4f68a9d45bde0) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-f9dd2a1fb59cc16f0651e7e8e3eecf0047de27be6e01448d017986ee242a6313) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-9fb2c7ab0788fab94ea7d8c76732185c46695fd51a235fe60cdc109bc5954993) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-0dce79ec83f4d8288e8581aa705ec725f968025d5fa961bcbdc8415edab5a4f9) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-003.md#canonical-901aec354adb9f20385cd521bda246b07eaf1bcee8f76211c469d44a812e6210) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-003.md#canonical-9d52aecd74dbe4ed038eb3b4254d61832cb4ce7efcb83c36e0c8120e0ea8051c) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-003.md#canonical-98aeeac38c86dac31d61f285ff7729c53affbfcf38e63343758ad0ae66aba485) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-c7f0d15f0d46ef80f81f8e648a599ec4a0fdb5bd76c4e45f26183cf50dc1a13b) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-84676e3211481fe9133b502c29b3edb00f5537ee9b6349d982b7627be8b2c3d3) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--reference--group-003.md#canonical-baf50061acafe10e948a9aa4e8104b4e212d499bed258cd65e1a42626d0e201c) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-82cf13559eda168fe8119e873f36047a234e1f8fcc024866f88db78d68a66b01) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-003.md#canonical-c7ca7a8fcb8bbfa08375a628cd84a87f6dfafc1f74e21af6363413a773cff1df) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--reference--group-003.md#canonical-502b24de89ca1e1b1a0068a6e97a427a6284d44c7212ccae3b837a9ee016b6ef) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-003.md#canonical-df4943326cd7324c37dc0d6cc7c1c36900ae2a7f741b239bfe89f366cada7436) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-393a408547a811624866cd08b5f4c024fe9d3da658cf0fc6a7363312b6bbcd5c) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-b73b1d254c29abe7e88b4eb2701bf6fef9a94b20a6952a9b451c2201bf59d932) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-003.md#canonical-c81784c97c6f7c0c0621f7143e3ba2c13ec2229fe146c2732588c97b6da28729) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-003.md#canonical-80b9e158851e1e884d68dcbf7226316688a4edb27d02313740a8556113bfedae) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-003.md#canonical-438a70b70d09cad647641536e0d48ef4d22160f992296d0d9efb457921bcdcb9) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-7811c4df08ae8ccdedee402ccbbfd76d04aee57c9a507af4ab3b890bb2ac8118) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-9af38733057be6492361f13851c2261252f049bd863e76508e346e4bde740cf1) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-003.md#canonical-03a2db73dc3fecfefad3906ec81f92f423033822f2da0854a4cb13361acf2a2c) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-003.md#canonical-153f22e9f40f3cdd74f99dd54735075e1c10f55a57a98ee4e0c076fae3ef79dd) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-003.md#canonical-3e9acc1797065511707f9e5ee32f7819d1b60f5587645633b55db80f43abae58) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-003.md#canonical-a6bc1484321bac4a7d584c15120471adb1dc5d46c6840dec1bc6721781710f72) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-003.md#canonical-18498408501693d229a28f817fd49a988b6de53602355e564ab4a852dd2acfdf) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](resources--bigip_http_proxy--reference--group-003.md#canonical-f4303b0268f2e4a4f4ebaabad8a2d7538740a04d2ea2de06ea8195d2ad14904a) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-28c4ef3f7b1d0f50063efed6b8fb5bdbe6928a2a162502e4e1432dfe1a4b51f4) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](resources--bigip_http_proxy--reference--group-003.md#canonical-0e9d063d2d26348ccbea12fd56101ea7deb8481038aa0e9ee5e277fb07edcff1) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-bfb937f013fbe490efc3c2371d149f1b4b5a4f34f92db2e49e59b9a3f7a6f48c) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--bigip_http_proxy--reference--group-003.md#canonical-50a83b298ff5b34ce82487dbc6699367c2adadb6451ad4229d16b238bdcead6d) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](resources--bigip_http_proxy--reference--group-003.md#canonical-dfd3d17b361c5e5cdffbacd6ae84eb201e2312c6079de8e2aead80b4e9f1b9ce) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-003.md#canonical-ef2e62d7f08ed175aeba1ad3b5de280c77dd2f27d6f516783c3e4b303677a4a3) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-003.md#canonical-c77e33afd614778021b9f115f9c1c655255fb73f2b55292f79197111900511a9) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-c2b3d64a85a37e3ca9aef94de67e22e403fee32369dba423940b67b673b85718) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--bigip_http_proxy--reference--group-003.md#canonical-268eb8209778de1f9e0f2c658e731158e40dcaa33688384feb0aca0a2e80f5d2) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--bigip_http_proxy--reference--group-003.md#canonical-196cf5cf4a6bd640b04dc8d627d1a76dad203fa8c3b55e57ee5bee2d25e742b1) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--bigip_http_proxy--reference--group-003.md#canonical-250a35845114e5a7179e65bb07d4cff6af3f69a125d9a050b2e3ec0390beded4) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--reference--group-003.md#canonical-d1e5237a84a4a3cfd416f1125ed9ae80598f22ddf1726b7d37004789606771b9) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--bigip_http_proxy--reference--group-003.md#canonical-148e0b4a8b1c1a70a3ced2358e236a3263ba04028c46bc87c52437aba1fa53b6) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--bigip_http_proxy--reference--group-003.md#canonical-8b899a3ff4ac4aebfbe1c8278a4c5168ca7f1cee9ce1fca20d50ee2eecece695) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](resources--bigip_http_proxy--reference--group-003.md#canonical-f08d2fe1755a51720d01252eb692896f4295e54a6e8b22cfae790d1e151a124d) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-003.md#canonical-93ecc36fa57cb54f9e2d0e9ebde281af2bace84a712d0e725200797ef21838ab) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--reference--group-003.md#canonical-a40d1133e7c8c8606beebf1f5153afb67aca392c74be4734ce5b60017381cd7c) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-003.md#canonical-fb3d6bf292be22848ddabf3b3a5923f95c1cf3dd206bf0f6b434bcab51821573) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-003.md#canonical-af2ac28dca91fd30f505363e35731419202134b0a4160c91473389abf8821558) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-003.md#canonical-114553bf2c513d42647d6cbf803ed14021240a0233d4ab823c36421064f6c210) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--reference--group-003.md#canonical-4afd0f8fa2def1390f28512d0526535bc25801d43fe9311496644bdcd06de2f1) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--reference--group-003.md#canonical-2642c54c589756b71c250c311fbac0dbdc7e33b39523b54535b7f2618663d480) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-3b7eab2bdefc6de1345f229a642aee5dcd84a1b4fcc0226107c2e4fb7d50e762) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-a2cfd5f0d13d917425506225e2fda5b10bbb14631551a8342cc52613e0eba3a2) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-004.md#canonical-7b87cb89664363140822129008932f18748fcf82da58396f8391d172b01178bd) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1bbfa2daf4a6a1a0882cd09f5b87f921ba1f08e5f9bbdaff2bdeb85ebeca592a) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-004.md#canonical-066037aec065facd0a8a39dfbd7f80bf10a86ec768c94ad19b6dcce72831f552) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-dab11c48577b0c25f27ac111e7d84eeeb2c1fb6b3c675cf42e359eb1faf0fdf6) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-73bf7bb1b3c002381f10c4bfe93d64a3c7eed079d53c89c0143a11fb5b224a29) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-9116507121d70ef3a1547c42a935d32044b914437b0da3e330563d43c618cb5e) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-5f52e905a14c2482f07f04777009d14452dd608a840d76cabe3317b9b11ad7f2) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-004.md#canonical-29efaf3f1038de19a30b19a94f31de6d996158976ec440271a761d4c52d78e7d) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-d30e9e42775e50aaa8d63da6b202c879dbb9220d83c304eb59dd544d54ca53c6) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-0a8c32b25d086678fcaece890c55c582e4b4e20e1edd642bc879e71bee3a0fc4) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-004.md#canonical-79a3457143dd8399674d31f12bc0dae65ac0f3a6b9f10c505d2d1429b45a9220) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-842a8cfcc03e1d149117caf689a8f5a05df0be98dbf4b2380fac139fb2194128) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-868f10fe1085de369723b747f1d57e9e5d93ee636e321bfdee5f4fe4f2674990) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-004.md#canonical-2d37fd2098beb3c1d49cd7d43376c78010d48da9fb69e5848b880ba9e463a863) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-6e73f6a85aed1b37e9e1792428626239c9ddb53cd35bd009dba978caca799f2e) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](resources--bigip_http_proxy--reference--group-004.md#canonical-49cbc6e967036ce230d86df6fe6c5783a90686d7313a411ce0feaa31ad27b9f8) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](resources--bigip_http_proxy--reference--group-004.md#canonical-87db4c1524c42c7dab588cd78f2bf6431b6a9fb2d47f9fcdd070dde6208ca351) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-8b68ea9f38d8c3462f1f03cb69838e58dffb2fe65b39090c1cb6fe8c581a71cf) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-55c5364e26157de294c5fc705de2ccb04094b64de7e0012563da77f9aff516b5) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-c4da4f577e3775bf63626abfdd6e36cd2eeca53e25ce5c35d694a7751c5f80db) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](resources--bigip_http_proxy--reference--group-004.md#canonical-18af64db4e311d4300619caf4c2788e3fb8b16167048ada82171e5c3aa63400b) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-e77d40b407696c87dbeaa232bc62835d2379891dbf90b3a4c9de42af382ba351) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-46c67bf9c1a3f3a9282a360d292037022f969e16c985722ee7a163aeafb7981f) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-e81d8537f776d8beda7c715b636e863ca5fb6428fe6f316a16392db70ff8721a) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-688863d8768b7dfa9b3505cdc0dc9806a224933b38e7d0498935afae4da0a1e3) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-2ac09f80a108783afbe72e1b3ee56d7b09dacdb18dd8b3a943073eefba31b5b5) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-e15485a4cc03412614fe65912546121c655f56d561ef516748208d8437f42091) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-b63627d3566db26fc58addd032afc8304c05a9fa0a840063256548e99863e534) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-3497eddb36495381b3aca9797b6f5a5669a3a45ca54437185b9b1d9e8004d3db) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-46bb6e5123ab3e50c6896eda15286455cd33a02c85555e6fbff75b16ec090370) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-8dfff61c77fd6dccd478839f98a907a8fc97b7aaa8cf87d47d113f7bd0b3c184) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-1dde5b3889ba9e797e56f805cb221f365cbec507f6e276dd0efdab84e5432a1d) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-e620f4ba679bb3982780ae873178e9239b29c8f66fcd51450d7be23923998a31) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](resources--bigip_http_proxy--reference--group-004.md#canonical-6ae6383fcd154436155071e01152e05d452b77a1911af305cc6c89ea0a1c6e28) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-7af6a275121c6ecfc5cda3ba03ba4e6fa38f1f2bc7d8457b2e55749a75389b04) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-a80f9ce26fae60cd6a30b868716f56772e37eeb5b7f8f05f0c237acd14290caa) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-457dea171029fcff29b5f14f4209b7a5cbe31b554d7cd33aca39abd6ea124989) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](resources--bigip_http_proxy--reference--group-004.md#canonical-8c2d463f85fe1287fb7cef2f84f85436243cebdd2354f17031db3dc64f1a4ec5) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](resources--bigip_http_proxy--reference--group-004.md#canonical-1d6ec60f7446084c88f9af36b7b69a539d5081909ea28da63067c52423ec5027) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](resources--bigip_http_proxy--reference--group-004.md#canonical-ea14d87d4577b906998a1745b9e316ae3583f5bd8d3178e9b21cbec7c6f0614d) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-51e770596702f303f71b4200bc1a0f56e839914f459cf469a35713cca298c125) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-b78e443dcfce007e3c23dbd3430478e0a631c5a183b8904d617f5e18d35595ab) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-004.md#canonical-68aa3f8a670426ad26814f59720229285d0712d847d3ce356616b56978d92b14) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-004.md#canonical-865ef665233461ffb071e6d5343ba4679dd11b4c8cbf313a799862db7c8a493e) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-004.md#canonical-cd551c5501b6acdecb7222876bdd7718ec38e87769fd3f98a0acab476059d8ce) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-a2e4205dbf11b7d105339ed86edf4bca98609e5ef717dc7baa4935c8f909d0f2) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1d6b16d6a28877d769d76cac703e8f36f4f6692338045b4822fd02c9de1d123d) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-79bc2315de52adf9d54743368d6716b02562fedbcd588937b8c474554ec17d5c) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-3a884abae75ebbe76a2bd80e94b51e697d52d4a62b057fa649f6762947e059a5) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-004.md#canonical-9b703272f479965f87071e224f936a4abb6f6c15d28fa76c26c11d5d8453e36d) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-da41602f3563f511e64796eda211e1bd0f44a74cf086920c5bf3c8deba2941fd) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-004.md#canonical-8880989e586a671c86b9e14a2c3ab7ea1c43dba24a9349314fe84c16a3748876) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-d7ae64d7bbb290997bbc1157e5dfc9a21f1054ad671c913d4e373bf4144cca7e) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-56a124e77db3dbadf4a3c89b3af013d6f3b958f825e23afd87216ba72446e021) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3d3f41cc0a2d56b4788397890278edb02ce74b27b17445dec91d62f470aebd36) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-4a9a8e017b05bd6023348f5dc13786d830ee14c2371c683e446235721fb8fa13) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-004.md#canonical-ecde507f0ca67adcca514d1469af28356075d885dde4d0f9a600e6eec3f20d61) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-5fd62edd63094372582c66b94bc21abadf45e0f80fe8bb9031396ece43607555) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-1b68b2b4af1a9ca9a939340d93faf2154853e1d0efd8bc67e5fce7e91bfcf02e) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-004.md#canonical-04d3e74b85e8b941530ae926514a7ff91a4327ed077a8d97c271d48f52fd31b2) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-3e251cd5b99b7b97b5ba6075a68c514969fd10708782883f6bef4b940f74700b) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-542bdfff731d2a5567fec73c4c7edcb756a306b5434e6810f11e03459a45cd54) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-004.md#canonical-560d757d9ba512ffef2fba84dee817d85506cc2bbbd37a5061e4d182745cc57e) |
| `timeouts` | [timeouts](resources--bigip_http_proxy--reference--group-004.md#canonical-585143a98318f2eadafe1239e1f547db2f40383501e9231fc5776ef9daf6bb62) |
| `timeouts.create` | [timeouts.create](resources--bigip_http_proxy--reference--group-004.md#canonical-c90ab0462223ee50768b2237c5915b4ce774beb9934b36f13104e329ecfe28c5) |
| `timeouts.delete` | [timeouts.delete](resources--bigip_http_proxy--reference--group-004.md#canonical-698e3207550fcc5aea203e60899bd579901393b3af3ea2d72a2f414a3dba6d56) |
| `timeouts.read` | [timeouts.read](resources--bigip_http_proxy--reference--group-004.md#canonical-d58a74198b860233a36e444435b067b08d5eba57b533c2385288effbce2e9147) |
| `timeouts.update` | [timeouts.update](resources--bigip_http_proxy--reference--group-004.md#canonical-52d2b0c3ef27dd72e17522f01941e1c8633d220c0112184a5e67e26ad3ffc0f4) |

<a id="canonical-5ddbc28ce79a0142b16c86fc89c687324593ea06e06a8d567e624f964fdbac7e"></a>

## Next pages — Property reference / a744dbd1f731 / 12

- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac)
- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b)
- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-58e1feceba1cc5dc0e29099842965bf040dcdf9ee383ce30502e44581d236fa2)
- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-04bc9f74502d8ad52eaa52994e94b471c36cc5c7eca34734297c5bf85ef6556d)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-c31e835e2a4d26e1df2850f8f036a9420e9e70004116bdb05f973efc6bb186cf)
- [timeouts](resources--bigip_http_proxy--reference--group-004.md#canonical-6dfd33c696868d9aaa2f87009aa602ef5a1be81177bdfd35ec1953db11750b92)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c142299894672a781a97c27f3ea437582ce6507f51cac04fb8b3e06ff416e0b2"></a>

## advanced_profile — advanced_profile / db12ca4b48ff / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- advanced_profile

<a id="canonical-a245461473faa06a1136f76879ec1ae234563a4a198520f8da9c2f73fbe9dbf8"></a>

Type: `"object"`. single nested block, Optional.

Defines various advanced Profile OPTIONS for a Loadbalancer.

Upstream description:

This defines various advanced Profile OPTIONS for a Loadbalancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable_default_profile")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecf774d577e96b43d0e3d3142f190baa9d8d735eae577afb436d73364e91aca8"></a>

## Direct properties — advanced_profile / db12ca4b48ff / 3

- [disable_spec](resources--bigip_http_proxy--reference--group-001.md#canonical-4d31e42a0296a68ce9c9cdf336020d2f6810fb5dea224ce4d827d4e69e9cc204): complete subsection reference.

- [enable_default_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-df26bedaf81a0b49f13494ff47917a835ebfffd2f7ce59e640e75ff84561444e): complete subsection reference.

<a id="canonical-3aa2c9de030915bbc6fc96c2ce50b665d773202724d1bbf2a1780b25234b601f"></a>

## Next pages — advanced_profile / db12ca4b48ff / 4

- [advanced_profile.disable_spec](resources--bigip_http_proxy--reference--group-001.md#canonical-4d31e42a0296a68ce9c9cdf336020d2f6810fb5dea224ce4d827d4e69e9cc204)
- [advanced_profile.enable_default_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-df26bedaf81a0b49f13494ff47917a835ebfffd2f7ce59e640e75ff84561444e)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4d31e42a0296a68ce9c9cdf336020d2f6810fb5dea224ce4d827d4e69e9cc204"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8157b0f893a744b4bb1f91282936a5732799bde871bfe0b1756fe03d77bc635"></a>

## advanced_profile.disable_spec — advanced_profile.disable_spec / 0631b037a032 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac)
- advanced_profile.disable_spec

<a id="canonical-ca985752719eacf36f287bdd6e54895f0a7d113043d942de92139634fc340fd9"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-0ba060c3fa3531977108b6ae1aea9abe5f7f102121c4c705e30da362dea5e882"></a>

## Direct properties — advanced_profile.disable_spec / 0631b037a032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f414ef4e17efe295534c5a24731b1b7b711862306eea411652fb5016a41d9242"></a>

## Next pages — advanced_profile.disable_spec / 0631b037a032 / 4

- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-df26bedaf81a0b49f13494ff47917a835ebfffd2f7ce59e640e75ff84561444e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80fda63951ca2e514603535869278cf04911a7be261c1dcea88e9dc123b1591a"></a>

## advanced_profile.enable_default_profile — advanced_profile.enable_default_profile / 07d62c6e3283 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac)
- advanced_profile.enable_default_profile

<a id="canonical-21dce0387b5edbaeb5e2c5ed61ae1106bc6b28aa78ca965e5d81816090cd87ad"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable default profile.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_default_profile = {}
```

<a id="canonical-81b1bc93046356518292d138b54a71632e84cb3ae3d59259ee1e10510b332928"></a>

## Direct properties — advanced_profile.enable_default_profile / 07d62c6e3283 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-012ca3480008920103202f2f2c31db2736d848347e5bcc82ede23d6666c2cb82"></a>

## Next pages — advanced_profile.enable_default_profile / 07d62c6e3283 / 4

- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-967ccde362be9a7b2d3e66a16089afb78a33210a69c01630d894fd67bb4f89ac)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d08cc166d5b0912b11173e3ea4fd3f580c07eb56c40ec43c6d766b90046b77db"></a>

## ddos_profile — ddos_profile / b2d3be9f011e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- ddos_profile

<a id="canonical-d9c0552c6bc9f1722a230fd8dbe5c3ec5638749dc0027abdd3a3aa705d057867"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a7cfccafea68d9d09cdabebb28048be6c946dfd8dfa8c0e6a78c75cbfd07e98"></a>

## Direct properties — ddos_profile / b2d3be9f011e / 3

- [disable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-7285f5abd01ecb987ad5abbc05866ca1f78d6c35363106fcb4ded207bacb0750): complete subsection reference.

- [enable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-4fbedccc562e125012fe2cff6d8faa13d9b206c0832aabfe1c7e0c3abcb2222d): complete subsection reference.

<a id="canonical-43d4f6f72b2bc83a812b6b4255d23a75956efb5845bf8852f9d4378bcefc3b05"></a>

## Next pages — ddos_profile / b2d3be9f011e / 4

- [ddos_profile.disable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-7285f5abd01ecb987ad5abbc05866ca1f78d6c35363106fcb4ded207bacb0750)
- [ddos_profile.enable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-4fbedccc562e125012fe2cff6d8faa13d9b206c0832aabfe1c7e0c3abcb2222d)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7285f5abd01ecb987ad5abbc05866ca1f78d6c35363106fcb4ded207bacb0750"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5edba67ef165c49a6a9e066b8371b083869676c57066ad495b6cdbcb7cc1ba70"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / b2fd1733f9a3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-02b0ef95a035fe65d28453b824caf35283ac9c2311bb9354186fe42a31533e92"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_ddos_mitigation = {}
```

<a id="canonical-77af7cd4d306dd5538f8431662c6137363fe4df09a943385a9df9fd2665409f4"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / b2fd1733f9a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20057a02b17ac67d5bd768dd780bdc602d6437841a5745ce6c4e59b751b8e55c"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / b2fd1733f9a3 / 4

- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4fbedccc562e125012fe2cff6d8faa13d9b206c0832aabfe1c7e0c3abcb2222d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02afd080ea96cd0a86bc5e225b3989587d8874252b0a061f0208e3973154180c"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / 6af936405b17 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-9a3dfb04717900fcb908a8a0f8a473c8410e707d740cf4de4cc552170f213380"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_ddos_mitigation = {}
```

<a id="canonical-f806328d2feb50f5eccf898bc3e95da70d12afe015634d5f1b961cada1fed1a7"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / 6af936405b17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1dda3f268e5b18b029c91138653eb77c84fba59277d3549be61d5201c367eb4f"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / 6af936405b17 / 4

- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-ebabda3589267c6fd410a5e557245ee450ae9fbfe0a3e502f12534f9c2a9c17b)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-58e1feceba1cc5dc0e29099842965bf040dcdf9ee383ce30502e44581d236fa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a4f3fda965f59ce9fdc40608e720878a521a0052b56fed61499a2bfb4d3e35e"></a>

## irules — irules / 0570b5c1dc3f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- irules

<a id="canonical-3011b3314c93539cefe83e9ecfe627c8bb8085fb4286f6248a6b2b76bf12ed63"></a>

Type: `"object"`. single nested block, Optional.

IRules Configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-a722e73de7f40b7e62739b5b77399ec944c3b3cc3da72f6ddadb6f3d30e248fc"></a>

## Direct properties — irules / 0570b5c1dc3f / 3

- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-4ca4ecd0717b29c6069e2f18f7ce397d62296b2cabc1d70a0808b710124a9c56): complete subsection reference.

<a id="canonical-c04824591964029217dc6d99ebd7a9360135c1413ea58de4d6f19d0a1fd97a00"></a>

## Next pages — irules / 0570b5c1dc3f / 4

- [irules.irules](resources--bigip_http_proxy--reference--group-001.md#canonical-4ca4ecd0717b29c6069e2f18f7ce397d62296b2cabc1d70a0808b710124a9c56)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4ca4ecd0717b29c6069e2f18f7ce397d62296b2cabc1d70a0808b710124a9c56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bad8ee6260799eb84c435e48a90ed8d38678bf1fb7c5a4b9870cd364ed5cf42"></a>

## irules.irules — irules.irules / a505d8777c32 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-58e1feceba1cc5dc0e29099842965bf040dcdf9ee383ce30502e44581d236fa2)
- irules.irules

<a id="canonical-aa005ef36d45e63cce4634ada13cc58857e9ef184641eddc2bc79cfe7efdbc0e"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to BIG-IP HTTP Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-5af6a376b1753dd693e8ab6600a93367158128872a52073cb4137fe3db827a35"></a>

## Direct properties — irules.irules / a505d8777c32 / 3

<a id="canonical-8d254192484074b5db9678bfdf3f61e29958720f1b2344b7e549bc5c8c20757b"></a>

<a id="canonical-7182902f50dc1545887e6677d0cca4dcd05d23d0d1af8eea145f0342680a898d"></a>

## name property — irules.irules / a505d8777c32 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-b7649b35116c034ac8ecf89f043fc56bbe3e37c310b7aa2ecdaf57766fc72b8f"></a>

<a id="canonical-0dd3f3905b72a4ca281c7744c17355410c5fca3fd3b0bb2219af48a9617c9843"></a>

## namespace property — irules.irules / a505d8777c32 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-769435ceaf302afef8f574b0f897a740242f72ec763fc7501b936b3ccbe007cc"></a>

<a id="canonical-ec0f6ce2294b7872fcfa1c5722844c9e237bdc197fc4e610086b2974d21b56bc"></a>

## tenant property — irules.irules / a505d8777c32 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-704a22bd52cc7c440598fae699c7e710c77d91b482572bc2276b5a003da7670d"></a>

## Next pages — irules.irules / a505d8777c32 / 7

- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-58e1feceba1cc5dc0e29099842965bf040dcdf9ee383ce30502e44581d236fa2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-04bc9f74502d8ad52eaa52994e94b471c36cc5c7eca34734297c5bf85ef6556d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed152ffafcc98bc79945aafd184e5b765def8d76636f5d326c23f0e74f760cc4"></a>

## lb_algorithm — lb_algorithm / 329a3b47c64b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- lb_algorithm

<a id="canonical-639dccf092030b6ea115b031644580ffd2cf87e0e5f57883d719c422d4791da7"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for lb algorithm.

Upstream description:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

<a id="canonical-8db861936877f8001199baad17373b0365c457924cab1f93f8d14bc28b4da55b"></a>

## Direct properties — lb_algorithm / 329a3b47c64b / 3

- [round_robin](resources--bigip_http_proxy--reference--group-001.md#canonical-4eebd314c30dad61a31032080f79bca0a76a57891f2fc5f4d267d5c229c9ab1d): complete subsection reference.

<a id="canonical-fc82e0dbf3af2b99278ff5ea7af264ef0a85ddc8b8c5ee764d102545af69c024"></a>

## Next pages — lb_algorithm / 329a3b47c64b / 4

- [lb_algorithm.round_robin](resources--bigip_http_proxy--reference--group-001.md#canonical-4eebd314c30dad61a31032080f79bca0a76a57891f2fc5f4d267d5c229c9ab1d)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4eebd314c30dad61a31032080f79bca0a76a57891f2fc5f4d267d5c229c9ab1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-361de2325459de49c2ebcbfa540127cd8c16c80a9dd36ac0f01bee9b0178cc54"></a>

## lb_algorithm.round_robin — lb_algorithm.round_robin / 4388a023c471 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-04bc9f74502d8ad52eaa52994e94b471c36cc5c7eca34734297c5bf85ef6556d)
- lb_algorithm.round_robin

<a id="canonical-48e56275efa603b505163573abcdc96ea9355528a4c19ea6603802abd939664d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
round_robin {}
```

<a id="canonical-a1da509f78322ed699aca72ad9b65fc7ebcb777cd96c4a8a05ce3799937fef15"></a>

## Direct properties — lb_algorithm.round_robin / 4388a023c471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-081de3e4bb4451f4318ab1d75bc47994a1e2ef89e8903628e06aa810fb631e6d"></a>

## Next pages — lb_algorithm.round_robin / 4388a023c471 / 4

- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-04bc9f74502d8ad52eaa52994e94b471c36cc5c7eca34734297c5bf85ef6556d)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2047aa6df7d9b40df14b18dc6bdfc3826ca5fc21cbf40fa4fb602d7f8fba848e"></a>

## origin_pools — origin_pools / 2ca23de8c98e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- origin_pools

<a id="canonical-bff420c9dcac9003b1b1cf86d1f48c359ccc36a63f3ffb709f26633f290ed7b5"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
origin_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-186c2701d814ef0c09eb4491b9aa947460006c8fbbfcda0018369d1774117797"></a>

## Direct properties — origin_pools / 2ca23de8c98e / 3

- [pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a): complete subsection reference.

<a id="canonical-424bbb565c13e81b933b95319efa6f2c67973ecd150f5b2aecbb4e9f03d50b0f"></a>

## Next pages — origin_pools / 2ca23de8c98e / 4

- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cb6b66a2cba9f49e5e1a39874d4688b238b19af759f16fddcbdf059a67b17e0"></a>

## origin_pools.pools — origin_pools.pools / 6d0828545740 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- origin_pools.pools

<a id="canonical-fb42e8e2a33adca0af5e42c56a9bf6bbb0410ec6873cc6e1ed31bb4790cb8512"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-95ffad4be072281af437608dbd419cd4e9a10ccdc4b2ad19740b635256f815f3"></a>

## Direct properties — origin_pools.pools / 6d0828545740 / 3

<a id="canonical-b58c37206966132e098899fa2b852838e54668a96401bcfbe21347c28ac6df67"></a>

<a id="canonical-3a429844d28f06e9a946e38f6d5acb08894d76cc23579ff114e71812a2d12253"></a>

## name property — origin_pools.pools / 6d0828545740 / 4

Type: `"string"`. Optional.

Name. Name of the origin pool.

Upstream description:

Name of the origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca): complete subsection reference.

<a id="canonical-8fb27a6dfdd8d9e31450042c56ff5adb75025ed0ff8ce5bafdd59ef470f1a232"></a>

<a id="canonical-b760c566e94aef0be9032e9db159a1e2bdcafc3905ee12f32e847b08b4318fb9"></a>

## priority property — origin_pools.pools / 6d0828545740 / 5

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-5589fb70691586957938c5d1ae461275c095bda1093a3817f40ffc5e17513713"></a>

<a id="canonical-c8dbdfdb3693a1cdd2dede57168cfadc7afc6378fce7eeda94b67e61febceebf"></a>

## weight property — origin_pools.pools / 6d0828545740 / 6

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

Upstream description:

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-189c71d78e9c0dfa75a52b80f380a3036238cbe29478976108b30ed0f61c66ed"></a>

## Next pages — origin_pools.pools / 6d0828545740 / 7

- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4631298cb5d3175a47dfbfc58eebf71b569f6300afbfb8057bdd0e6bb2bd248"></a>

## origin_pools.pools.origin_servers — origin_pools.pools.origin_servers / 3e81649600f2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- origin_pools.pools.origin_servers

<a id="canonical-e5eddd711f418f181d88576d74b8b61708baf248384b3b50eeb1af868c11c168"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the BIG-IP HTTP Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("automatic_port",
    "lb_port"),
  validators.ConflictingObjectAttributes("automatic_port",
    "port"),
  validators.ConflictingObjectAttributes("lb_port",
    "port")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-fee392c0d75e921f1d8819b55a2931ffff1bf4a2f9e9d8a7d6fa3c23707db9ea"></a>

## Direct properties — origin_pools.pools.origin_servers / 3e81649600f2 / 3

- [automatic_port](resources--bigip_http_proxy--reference--group-001.md#canonical-7fd1ed2f7c6eb592ae76626483b9038f671f723143250f0199954a20e776d2d2): complete subsection reference.

- [health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533): complete subsection reference.

- [lb_port](resources--bigip_http_proxy--reference--group-001.md#canonical-02bae522ee166fe4e22ac04806cdfad51eb0c1c70a3745f87889a1535a91716a): complete subsection reference.

- [origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638): complete subsection reference.

<a id="canonical-a265d272f1fb14df5948f28f9b982ee56fe8028cd1c27588bea90d66e7dd0aa4"></a>

<a id="canonical-d3deeb476fcdec27dd3f96943c0ffa05160a1650c28a94d3bee7c12edff9eef7"></a>

## port property — origin_pools.pools.origin_servers / 3e81649600f2 / 4

Type: `"number"`. Optional.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-912718827e22019bc69c210434f459ba04a07a6fd13e79a94f7f25f4fc47ddfe"></a>

## Next pages — origin_pools.pools.origin_servers / 3e81649600f2 / 5

- [origin_pools.pools.origin_servers.automatic_port](resources--bigip_http_proxy--reference--group-001.md#canonical-7fd1ed2f7c6eb592ae76626483b9038f671f723143250f0199954a20e776d2d2)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533)
- [origin_pools.pools.origin_servers.lb_port](resources--bigip_http_proxy--reference--group-001.md#canonical-02bae522ee166fe4e22ac04806cdfad51eb0c1c70a3745f87889a1535a91716a)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7fd1ed2f7c6eb592ae76626483b9038f671f723143250f0199954a20e776d2d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edb220bab3e21bd7ada62b9dbba398744f0331ae20893093f1d24f0dc93007ba"></a>

## origin_pools.pools.origin_servers.automatic_port — origin_pools.pools.origin_servers.automatic_port / fa70b5328cf9 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- origin_pools.pools.origin_servers.automatic_port

<a id="canonical-b8bca6fc5a390d96735ad200ef19553308437b8ddc939bd1f26b8b9ed52e74bf"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
automatic_port = {}
```

<a id="canonical-e3a3688182333c3138509ddc678321a252d5d80aaae5188c8b49757759aa7605"></a>

## Direct properties — origin_pools.pools.origin_servers.automatic_port / fa70b5328cf9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a212fcf1e35008b79f83095b418ccb4d836141338c6c5ce2a24925278a851a1f"></a>

## Next pages — origin_pools.pools.origin_servers.automatic_port / fa70b5328cf9 / 4

- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e77823673f02dc526746abd2b9ddd9bf724066fd7eae5a15a7867a773f10d114"></a>

## origin_pools.pools.origin_servers.health_checks — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- origin_pools.pools.origin_servers.health_checks

<a id="canonical-7d4e468e6a58209e4488ac91128697fffed299f88ed07c12035b3648dc972dad"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for health checks.

Upstream description:

Origin Server Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check",
    "healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-09e108c05c12c2a25c76f66bbad0e8bda309d9079528c723c3ad50be9ba92a82"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 3

- [health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3): complete subsection reference.

<a id="canonical-8cd6a338af752a00b65cdcc2ed2dfd6004022546ddb72e7f770f9328eb1f6d25"></a>

<a id="canonical-f9244710b0963182e3fd6a4f2870c00407866b5b7695d135c94a55396fac491f"></a>

## healthy_threshold property — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 4

Type: `"number"`. Optional.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-48b34c416e0713d385e1f8a366985834a08e520cf721a66c1362702c9c612952"></a>

<a id="canonical-c06f1283342f277c05ee6d96c8a0f7fea81b3cc1dfcf48c9a6ae69ff26db328f"></a>

## interval property — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 5

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-8fb39008ceea2e5e150312389b1cdf44549b56d4ce6a8cd378824dac84a85611"></a>

<a id="canonical-1886a165277dd2867e7ec8b3cbf1a26be8d5769e91f8a347dfddef8d244403d0"></a>

## timeout property — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 6

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-319757f1a535b8cd2f17d9fe83829a558ad8014507c899d313bf25e379c154da"></a>

<a id="canonical-1e851bf9d209080cba0210ab53d97b8fa4608124a9b91ef9c3884e34959f5b0a"></a>

## unhealthy_threshold property — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 7

Type: `"number"`. Optional.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-ab2ada48f6c5221a3793943609555ef6a5bbae21620091e5d8cbb78e923d290f"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks / b97de2c257fd / 8

- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-631da5b0ff84e5a0df9310cbaa9c6d145bd8598e486f4bcdb47f27ac000fd855"></a>

## origin_pools.pools.origin_servers.health_checks.health_check — origin_pools.pools.origin_servers.health_checks.health_check / 911f36c0c281 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533)
- origin_pools.pools.origin_servers.health_checks.health_check

<a id="canonical-9faff68e85044bd7e44dcb09348b77bb07ed49854faffab26101109ed2b8ed24"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-36261f602b047ff1f62a492e454c88c49ffda745a105e89d8e757ab0df7bde7f"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check / 911f36c0c281 / 3

- [icmp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-ecd69bf9d7770606f6a4c4c94344d1ae6ee3915fdb00c229fe6adc2d54f46323): complete subsection reference.

- [tcp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-1605572a0272242822cbdd7ffea415675711305e0da6605b29f0d356eb1c56d1): complete subsection reference.

<a id="canonical-cfb19fc4bf91ebb4ac671a2a1f84c26bce193a686ec2a4d9d1a1170015394c65"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check / 911f36c0c281 / 4

- [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-ecd69bf9d7770606f6a4c4c94344d1ae6ee3915fdb00c229fe6adc2d54f46323)
- [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-1605572a0272242822cbdd7ffea415675711305e0da6605b29f0d356eb1c56d1)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ecd69bf9d7770606f6a4c4c94344d1ae6ee3915fdb00c229fe6adc2d54f46323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58ce257631f416857c8e7b91bd44bfea6696449392c412375114050b78cc2098"></a>

## origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 9d6e29521326 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533)
- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3)
- origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-9198610961bd03ba004dd63f8f9ae5853e65a84e703ffd8ca2c7d6bf2be97312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
icmp_health_check = {}
```

<a id="canonical-0a13229d2e3864c60e5a518733417db0f755685798738820bc86558d5c91c3c3"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 9d6e29521326 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3039e2e3c4353c65c68e4a7fffc1e0fd2bef3ea27ee2d8ce979443aab226b336"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check / 9d6e29521326 / 4

- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-1605572a0272242822cbdd7ffea415675711305e0da6605b29f0d356eb1c56d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72cd3ca6e23f0c6567c62c11e49d0c05d18641fffa3b5f4c8a8daff89314708e"></a>

## origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / d99ff6045edf / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-6dae08e6ab4988641e2c1fd934e48987b3b50c7fed73782308584062a009d533)
- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3)
- origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-d12882321f19a1d130f70a701dfd85de832f4ac5d23a05036a5ba3b5dd93ffdd"></a>

Type: `"object"`. single nested block, Optional.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "send_payload")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-d001999020b6a0f8b48abc7ac6e4afcabcc4eebd2a9e9debeeae2679f15dcbca"></a>

## Direct properties — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / d99ff6045edf / 3

<a id="canonical-5a40146f78fe61d13862ffc4f960f24bd6f3af86caf0ebd825108addac4cf2e1"></a>

<a id="canonical-35d1efa4e07223bb99572a2609a6da601a14be14c86bf64b3b51e6943f892218"></a>

## expected_response property — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / d99ff6045edf / 4

Type: `"string"`. Optional.

Specifies a regular expression pattern which will be matched against response payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-c39bce1b15e838ce934ff946f0e430f395a98f4d00cc42ae7eb28511cd26e589"></a>

<a id="canonical-bcdf2dace77f240bb1639eacf2a611ebb74c469ad0487dcf1211f932c1da3939"></a>

## send_payload property — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / d99ff6045edf / 5

Type: `"string"`. Optional.

Send string. Text string sent in the request.

Upstream description:

Text string sent in the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-9cb32c9602336fb84cdc5599c5a9c49e12ead50eb0a5d284f12a4e1d0ad86543"></a>

## Next pages — origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check / d99ff6045edf / 6

- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-98fed67b83378191b29251ff1aed0e73742acaf3a229b0ad0ffb5df2bb73d4b3)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-02bae522ee166fe4e22ac04806cdfad51eb0c1c70a3745f87889a1535a91716a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec9de7d7a463ad8d57522d4d54970bebff19ff17492d94202d88824f9a18b771"></a>

## origin_pools.pools.origin_servers.lb_port — origin_pools.pools.origin_servers.lb_port / dc212aac1936 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- origin_pools.pools.origin_servers.lb_port

<a id="canonical-e2bebbf461d5873f564360e7923c2176220602f995efcb184a252041127013a8"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
lb_port = {}
```

<a id="canonical-9072effa540f4a9788f29e3519b5db73675eb6f466871087f9f38b7c3e162a01"></a>

## Direct properties — origin_pools.pools.origin_servers.lb_port / dc212aac1936 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65fd841800c73d19fd2e0011685a8aec334dc531c41a497cc53ef5ba1fd9db0e"></a>

## Next pages — origin_pools.pools.origin_servers.lb_port / dc212aac1936 / 4

- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a1df90ef3fb3014a37ea45762f5b5aad39d1e3bd00f6768b8c265ae8644fe07"></a>

## origin_pools.pools.origin_servers.origin_servers — origin_pools.pools.origin_servers.origin_servers / e0e7ebd3d110 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- origin_pools.pools.origin_servers.origin_servers

<a id="canonical-05010ee3dab4c54633a80d65f9fdfe74991299190f45e31127ed79f757a1d6e0"></a>

Type: `"object"`. list nested block, Optional.

List of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-b19c65d313556d8fa751208f49e5d48506ecd89c8e97e13586f95bcafbcf457d"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers / e0e7ebd3d110 / 3

- [k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef): complete subsection reference.

- [private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8): complete subsection reference.

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-a0fb68a40e29f1563916c0752435bd05b633fe257df629a801bd86b305b8aedf): complete subsection reference.

- [public_name](resources--bigip_http_proxy--reference--group-002.md#canonical-82f1b165508c822fa0aa39d03b6ca6cbb3c668bc683b9852c23ec6c73c6ef894): complete subsection reference.

<a id="canonical-812ba47c500da9b9904c9385c5110ba2a5d907f5f966aa53b26de05bf72fb9a7"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers / e0e7ebd3d110 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [origin_pools.pools.origin_servers.origin_servers.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-a0fb68a40e29f1563916c0752435bd05b633fe257df629a801bd86b305b8aedf)
- [origin_pools.pools.origin_servers.origin_servers.public_name](resources--bigip_http_proxy--reference--group-002.md#canonical-82f1b165508c822fa0aa39d03b6ca6cbb3c668bc683b9852c23ec6c73c6ef894)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f087d8d8cf69145a525aa38f2eaf297737829c85a1dc16cda6f90ad671b63eb4"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service — origin_pools.pools.origin_servers.origin_servers.k8s_service / e9fdd062e908 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- origin_pools.pools.origin_servers.origin_servers.k8s_service

<a id="canonical-0300907df319b13dc6850d94b60d71cf2d7ca28baf7de43ce1eb9ca88f0adefc"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-d34e2accda0fb108d58886588b7d596031e9ee4aabdd7bc1916f0dec064d0628"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service / e9fdd062e908 / 3

- [inside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-01845d822c6b9720bb9515bf422cb24c6c40e3b43ccee2c2d2af71546efded00): complete subsection reference.

- [outside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-0e49e62f824149a23def7ba9801f32663ec3f313203efd05283f92f21e932b68): complete subsection reference.

<a id="canonical-755a88bc018d114a6e2ba6f2bc7962ed8990550d284d74240ae9030f221e7029"></a>

<a id="canonical-09523c670e5e61cf03ade8dd9e5eb4ade376a18a086b7456093895dbf0ce7648"></a>

## protocol property — origin_pools.pools.origin_servers.origin_servers.k8s_service / e9fdd062e908 / 4

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9584ba3795b5480ef8ff059b23d75c5cfc0b5a7baede61b17ab168ddc22bab14"></a>

<a id="canonical-606bef90e32414e8d50e39c4e273c0c27976e8ba29e3f7d2f5d30d6bd77f33ff"></a>

## service_name property — origin_pools.pools.origin_servers.origin_servers.k8s_service / e9fdd062e908 / 5

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844): complete subsection reference.

- [vk8s_networks](resources--bigip_http_proxy--reference--group-002.md#canonical-05cbeb1e8c1729eef1f772fe6daa3359311b9d5cbcc129787e11b754c80cf539): complete subsection reference.

<a id="canonical-0b6f4335cea2dd78fa27e53f0fde2a29e2b860df4eb81671fee471d7f4f72856"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service / e9fdd062e908 / 6

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-01845d822c6b9720bb9515bf422cb24c6c40e3b43ccee2c2d2af71546efded00)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-0e49e62f824149a23def7ba9801f32663ec3f313203efd05283f92f21e932b68)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](resources--bigip_http_proxy--reference--group-002.md#canonical-05cbeb1e8c1729eef1f772fe6daa3359311b9d5cbcc129787e11b754c80cf539)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-01845d822c6b9720bb9515bf422cb24c6c40e3b43ccee2c2d2af71546efded00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c27a155b514dbb57f85519f0c3aa6b920eb8e8f19122670a52773f0140ef6c2c"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / bba356275817 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-829d9509735656256b1cb9588c26b8f8d14f2c9ddaa4bd60d5f62883758c1bc0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
inside_network = {}
```

<a id="canonical-7243098829fa53301294421ed18fb324cbe8a7a16724817c49849777e3ffe376"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / bba356275817 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eccddb88982ee6962cc8c72831cb3b66d0eed78e1107914cb4038f7fe643f6f7"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network / bba356275817 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-0e49e62f824149a23def7ba9801f32663ec3f313203efd05283f92f21e932b68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
