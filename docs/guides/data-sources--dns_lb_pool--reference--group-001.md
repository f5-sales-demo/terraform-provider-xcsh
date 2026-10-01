---
page_title: "xcsh_dns_lb_pool reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool reference."
---

# xcsh_dns_lb_pool reference

<a id="canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa471512650c24dec224f7f6233332be29ee47da0eb0bffbba528a90abcef2e5"></a>

## Property reference — Property reference / f3ee0b56ad77 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- Property reference

<a id="canonical-60b8e98bfba9e611d7cb0bc55a3befeb13a7034fdd7e126646ff90ca17021b47"></a>

## Direct properties — Property reference / f3ee0b56ad77 / 3

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670): complete subsection reference.

- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d78390dbfa887147f79c03595f13d2397762cc7d74ddf39ffd3b58896122b2da): complete subsection reference.

<a id="canonical-e06741e67060eaaf8700e066eb8598ba649ff1981720cf33c55e0cccf83ac489"></a>

<a id="canonical-3e86b9f1dfdee36b91ab050ca41721e69f5b19ef748dcb186cd4c7fa77fcb7ac"></a>

## annotations property — Property reference / f3ee0b56ad77 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565): complete subsection reference.

<a id="canonical-6d615d2d9785d4ba29dce71ec2cc4022aeaa69582c088b64b8ee37b632c297cd"></a>

<a id="canonical-9d0cd0a0eb29a5d0d58e98319d6789477bfa62226f48258bdbf8e3360a5d1f65"></a>

## description property — Property reference / f3ee0b56ad77 / 5

Type: `"string"`. Computed.

Description of the DNSLBPool.

Upstream description:

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

<a id="canonical-418f4c0e66dc916eab8ae649b1d9bb8e8809ac6be151f7faeb693fa70dadaba5"></a>

<a id="canonical-3a17b7d9c3829579e48c1715472dabf81f86dfefd7551c1a1e2fb13f6932da79"></a>

## id property — Property reference / f3ee0b56ad77 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5aa0969af1dbe33f60d23456f7f110584c8e55e36e3e4537a3ce7aaf876fffe2"></a>

<a id="canonical-1648e0668df92135f99b14128b1ebd2164dca2b37bca59f7ca5047a6591ad269"></a>

## labels property — Property reference / f3ee0b56ad77 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-978acfe927cfea4ec20e54ed29809f3f6e4de625a8ebf34d09b360431a5460ba"></a>

<a id="canonical-bf36acbccd82f0bf2bc9df3ae0dff33e3f24f3d4a58e29e7c8e981e059dc675d"></a>

## load_balancing_mode property — Property reference / f3ee0b56ad77 / 8

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|RATIO\_MEMBER|STATIC\_PERSIST|PRIORITY\] - ROUND\_ROBIN: Round-Robin Round
Robin will ensure random equal distribution of requests among all pool members in a pool. -
RATIO\_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members
based on the ratio assigned to each pool member - STATIC\_PERSIST.. Possible values are
\`ROUND\_ROBIN\`, \`RATIO\_MEMBER\`, \`STATIC\_PERSIST\`, \`PRIORITY\`. Defaults to
\`ROUND\_ROBIN\`.

Upstream description:

&#8203;- ROUND\_ROBIN: Round-Robin

Round Robin will ensure random equal distribution of requests among all pool members in a pool.
&#8203;- RATIO\_MEMBER: Ratio-Member

Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned
to each pool member &#8203;- STATIC\_PERSIST: Static-Persist

The Static Persist load balancing method uses the persist mask, with the source IP address of the
Local Domain Name Server (LDNS), in a deterministic algorithm to send requests to a specific pool
member. If the DNS resolver passes ECS (EDNS-Client-Subnet) information, then a hash of it will be
used, to send the client to the same pool member &#8203;- PRIORITY: Priority

The Priority load balancing method returns all available endpoints in a pool with the highest
priority. Pool Members have a priority value, starting from zero, where a lower value means a higher
priority.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2a9f11997dde1afdec1b2b9c47a29d79fe7c2ef6781adfb1a8bc976eac447a81): complete subsection reference.

<a id="canonical-6ba18330c8daba91fe453ec3698b5d26c74f1d7bfbc04f28bb769b71a1f63516"></a>

<a id="canonical-91b70bf4c355c10117a48ff1499753a48bc178f087c5ab153afdd04356ab128e"></a>

## name property — Property reference / f3ee0b56ad77 / 9

Type: `"string"`. Required.

Name of the DNSLBPool.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-837935863c6180e011788174d5885aadb413f61a66edcabb6dc3804e0ba482b1"></a>

<a id="canonical-014a1040b013a1c1aa9b29db289be9516fc09e4f3b952a74481cdd2fd63bcd36"></a>

## namespace property — Property reference / f3ee0b56ad77 / 10

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBPool exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-4f476814f8d8e2237a9fa285952a58ae91547f2d40311da123957b32f48c0bb5): complete subsection reference.

<a id="canonical-0f5c9bc31f9d964e8bfab878f54464e7d161fdb93f1a3af5dc4d5db05ed31a21"></a>

<a id="canonical-6dfeb91f296ec3eda178aa24a794f1250d8d462613d1dd4a9e791b295e7dd734"></a>

## ttl property — Property reference / f3ee0b56ad77 / 11

Type: `"number"`. Computed.

\[OneOf: ttl, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Upstream description:

Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default 30) for responses from this pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

OneOf alternatives in this subsection:

- [ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-0f5c9bc31f9d964e8bfab878f54464e7d161fdb93f1a3af5dc4d5db05ed31a21)
- [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-c1961692fcb980d6e5d93906168005087e5b536d25127b156de349beaa30873c)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-d38d32d467bc746f086850c1b52814e108f6277944a51c027513a510ec7003f9): complete subsection reference.

<a id="canonical-c49d414a2615b9c26b09537be82c786fbb429596e48308d924f57463c29899f0"></a>

## All schema paths — Property reference / f3ee0b56ad77 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d47c372a497905f5fdf26c31d00a6296e016eed5be971791671f6d8652d6c55c) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-0d42d1060260750b10643467399e12cf559e0017b07bf0119bba13e146d7eb35) |
| `a_pool.health_check` | [a_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-d8574d7e3310d85d1660600fe676fc02a95603a22a4b08913963e2a6eecf0558) |
| `a_pool.health_check.name` | [a_pool.health_check.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-7ddac6e998b9c602114c71961cea6920bb937627b9e651102b43bb320c22fb6d) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-6ec2d6267e42d83c1090102f15d1b55d6f9f1fe08aaf7730c3797fd418cd84dd) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](data-sources--dns_lb_pool--reference--group-001.md#canonical-99ec7c4ba43ea5ae6750dbf4556ea3b6c01effcb3ce671dd2d63f4a692729fd8) |
| `a_pool.max_answers` | [a_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-a1cdc97e9413ba154a47051291736efab69d82a88b22d899eef72cbc68d29e43) |
| `a_pool.members` | [a_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-a1a4242a3ff38ab465edcca36d0a3e1ebbc7364da3a585ce3e66103a0db9fd6d) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](data-sources--dns_lb_pool--reference--group-001.md#canonical-a63c9470ada9170261760c93160b86d5836fdbf09f834e2fb3015ec3c3a14776) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](data-sources--dns_lb_pool--reference--group-001.md#canonical-dd437ad8f23de33b5cfa9298e42084552e9c76d0b29fbb7c4df1064df0e1911e) |
| `a_pool.members.name` | [a_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-c0d83deca673875be26e5c7e85327e124f3203627a93e3859135711f58afd344) |
| `a_pool.members.priority` | [a_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-cac32380d23d25d3ffd231de1038a4eb38474352bb3517833077d79d1da7e0bc) |
| `a_pool.members.ratio` | [a_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-b4ee41906e8cac4a26f8a65e301db80ad1b75809f007abb2caff4716af4fc8da) |
| `aaaa_pool` | [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-fdffe7f83a9a7e161675bd7a5b68d226f4239527b12f6af89711acbc2b8a91a1) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-0fd3d17785dde6472b62e7344450aa2e915c96d95481b426035b43f5da6abc16) |
| `aaaa_pool.members` | [aaaa_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-8ea6f8c8d3e7d306a3d9ecaefaf1b62f862db2c9b4ab9c7b5e3f44090f98b176) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](data-sources--dns_lb_pool--reference--group-001.md#canonical-57db3eaf63596e984ff896b3b7f3376c79e59d19baa3720c192b20043ee16df7) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](data-sources--dns_lb_pool--reference--group-001.md#canonical-a14e976f3f39ebdb79f968f44a464f4b0f16d3af9b5fa7715cfaf96f4a4b25f9) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-ce1dc1cd19fd8f48e7802ef3c7385196bf6289d43cbd918d855acddba02fb50b) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-19ce419ef368952e3b91767ebe04fb1f752ea8d9dbe1162186e035fc6fa6d171) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-1a5c9d92267a74240076e6bf04d9fd98450a30fbd567a9800d64749f82f8fc67) |
| `annotations` | [annotations](data-sources--dns_lb_pool--reference--group-001.md#canonical-e06741e67060eaaf8700e066eb8598ba649ff1981720cf33c55e0cccf83ac489) |
| `cname_pool` | [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-84eb3b1fc41a65eee05b99b5ab211db907601b58690f37c56740657810d7923f) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-776987590421da2efb3dc9479015f9d4138c8329f3ec6fc9c9e4db40510c4bc2) |
| `cname_pool.health_check` | [cname_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-e2358541a4bdab4a373db59296fbe8251ed848da7654662fe58ba31df279abaa) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-8693e8b8157f709f37d268ed39153e7b88e52743c6a3d5960dd9dc8a7b5b6132) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-0927b858fd47f20843a80b9d659d73809bf7fe8b40b36f55f6a0cf3111aff551) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](data-sources--dns_lb_pool--reference--group-001.md#canonical-f685cff9bd191f399414fbefaf3aa5becd38607c198a092a9200f8aa54b46868) |
| `cname_pool.members` | [cname_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-ca83742d50bef4f79eb553407ee09b6c5d09e413172ab32d869bbecbc39a4ef3) |
| `cname_pool.members.domain` | [cname_pool.members.domain](data-sources--dns_lb_pool--reference--group-001.md#canonical-f2df520cb60bd66caa28f9d7010da3cfa1d765852ec6c6858b13eb3425cfd993) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](data-sources--dns_lb_pool--reference--group-001.md#canonical-aefcab84cab446b188649b0c02e31f81b7bc284e9807d7ed540864fa60888874) |
| `cname_pool.members.name` | [cname_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-a6786f8bdb202c2739ea28fa9b46fb0a7ce09dc4721be605814ee32e606ed71f) |
| `cname_pool.members.priority` | [cname_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-e19ca9d59ac94134397222f3b49920371cb4ded01d95891ef25d2823160d4812) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-e07c0de17c1ffced72cf495b4f17edbb3a7acd1b43d6b6ab31d990db6c6966b7) |
| `description` | [description](data-sources--dns_lb_pool--reference--group-001.md#canonical-6d615d2d9785d4ba29dce71ec2cc4022aeaa69582c088b64b8ee37b632c297cd) |
| `id` | [id](data-sources--dns_lb_pool--reference--group-001.md#canonical-418f4c0e66dc916eab8ae649b1d9bb8e8809ac6be151f7faeb693fa70dadaba5) |
| `labels` | [labels](data-sources--dns_lb_pool--reference--group-001.md#canonical-5aa0969af1dbe33f60d23456f7f110584c8e55e36e3e4537a3ce7aaf876fffe2) |
| `load_balancing_mode` | [load_balancing_mode](data-sources--dns_lb_pool--reference--group-001.md#canonical-978acfe927cfea4ec20e54ed29809f3f6e4de625a8ebf34d09b360431a5460ba) |
| `mx_pool` | [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-531e930254b790101c8a9274746ab4c9706ef8a5fc6a3c06c76b0a0dcd0a87bd) |
| `mx_pool.max_answers` | [mx_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-d7b608e8752a1c8806195e0543909069a22edd3edc0f85e7620cda17fa59399b) |
| `mx_pool.members` | [mx_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-96903a0d4a9c8e83ca51ddf4aa0218be180e883a826e4cf6c1a01426859c2ecf) |
| `mx_pool.members.domain` | [mx_pool.members.domain](data-sources--dns_lb_pool--reference--group-001.md#canonical-949010a8337d68c80f8d0be3dd94ba99eb359c19baeb2c7a9be9296643c1ed68) |
| `mx_pool.members.name` | [mx_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-11b8ee22d305142b102502212b6fe49bf6cd122e59adef6ae4e3765cedc2d9dd) |
| `mx_pool.members.priority` | [mx_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-c09cd470682053303e8ef6e57d1a63821254b56e0922f3bae1442605e75f9826) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-eea8302fb9cd24fd9f33acfaa7cefc13ec205fd364e2d59f07047f457d7a12d6) |
| `name` | [name](data-sources--dns_lb_pool--reference--group-001.md#canonical-6ba18330c8daba91fe453ec3698b5d26c74f1d7bfbc04f28bb769b71a1f63516) |
| `namespace` | [namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-837935863c6180e011788174d5885aadb413f61a66edcabb6dc3804e0ba482b1) |
| `srv_pool` | [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2fb1d1ddd7a907e23873283a6b2f07c31bc580a22f878730f5e08b8aeb54e25f) |
| `srv_pool.max_answers` | [srv_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-653a21f420ccb7632cfea51aaa6aeca47d05680ceef2d331ed19f74568db5109) |
| `srv_pool.members` | [srv_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-e76fe0ce55f461a20a9316a2f1510ad588a2737c6b100662632eedc44e58c916) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](data-sources--dns_lb_pool--reference--group-001.md#canonical-50d107e013d8abdc6aeb865bf1a8759434d03e936d549ecaeec6cc0d9495bd95) |
| `srv_pool.members.name` | [srv_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-b2a7ec36b0a5fcb731203d151a3782dbb12e8362db88c9845019026dc97aecac) |
| `srv_pool.members.port` | [srv_pool.members.port](data-sources--dns_lb_pool--reference--group-001.md#canonical-b5d12418ae0c9c5f37d6e4fbe4d3167d3f57b9a7d64f316b8a3a4397a4a8cc41) |
| `srv_pool.members.priority` | [srv_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-5e5452b54c1d4c7ff262f0a1264269b4005ccd023706caaa27413c4463310be5) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-8f44a4a01f6bbc62c2c6bbf4408594a4f6e2fd45501f6de0f3fc51ea57397ddf) |
| `srv_pool.members.target` | [srv_pool.members.target](data-sources--dns_lb_pool--reference--group-001.md#canonical-89a1e69a7e445b2713ab3d5c207245e5596d0dfaaa0b2f9dfcbca9cdf10dfdb0) |
| `srv_pool.members.weight` | [srv_pool.members.weight](data-sources--dns_lb_pool--reference--group-001.md#canonical-5547e166ec99f13be6a971d46f9d34742c9d8e866fb0592b680a2f6b024ee1a7) |
| `ttl` | [ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-0f5c9bc31f9d964e8bfab878f54464e7d161fdb93f1a3af5dc4d5db05ed31a21) |
| `use_rrset_ttl` | [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-c1961692fcb980d6e5d93906168005087e5b536d25127b156de349beaa30873c) |

<a id="canonical-25330cd1354d5c09700962016ca2bdca9375ec1e1ac558a1b83c5af622e4e611"></a>

## Next pages — Property reference / f3ee0b56ad77 / 13

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d78390dbfa887147f79c03595f13d2397762cc7d74ddf39ffd3b58896122b2da)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2a9f11997dde1afdec1b2b9c47a29d79fe7c2ef6781adfb1a8bc976eac447a81)
- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-4f476814f8d8e2237a9fa285952a58ae91547f2d40311da123957b32f48c0bb5)
- [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-d38d32d467bc746f086850c1b52814e108f6277944a51c027513a510ec7003f9)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bf6df11373289735139c38bacb5d468297d56d980d593ebda6ac6ac1e61751a"></a>

## a_pool — a_pool / d22919da6b10 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- a_pool

<a id="canonical-d47c372a497905f5fdf26c31d00a6296e016eed5be971791671f6d8652d6c55c"></a>

Type: `"single"`. Computed.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

OneOf alternatives in this subsection:

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d47c372a497905f5fdf26c31d00a6296e016eed5be971791671f6d8652d6c55c)
- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-fdffe7f83a9a7e161675bd7a5b68d226f4239527b12f6af89711acbc2b8a91a1)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-84eb3b1fc41a65eee05b99b5ab211db907601b58690f37c56740657810d7923f)
- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-531e930254b790101c8a9274746ab4c9706ef8a5fc6a3c06c76b0a0dcd0a87bd)
- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2fb1d1ddd7a907e23873283a6b2f07c31bc580a22f878730f5e08b8aeb54e25f)

Select alternatives according to the provider validators above.

<a id="canonical-8e7f43d64e4aba3d1d5eacf69892670bbe1126c25c092251dc4a0b55e71ab750"></a>

## Direct properties — a_pool / d22919da6b10 / 3

- [disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-5f5a941cdb59ec6ab3ebacdc0c3257ca6d02a9b6dfd1063c5560af63aa7ff3db): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-e75e2a29b4df1adefbc44ec1f4fde7c7e28d62347e9021a1eedb63195e0efa78): complete subsection reference.

<a id="canonical-a1cdc97e9413ba154a47051291736efab69d82a88b22d899eef72cbc68d29e43"></a>

<a id="canonical-b1578ddbe84c4d397c61bc2653a779589b9fc1e7d16aa73c251491232a42c54b"></a>

## max_answers property — a_pool / d22919da6b10 / 4

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-7b88fb00da79b580d3fba8811785d895af766df441ef4c0deef646abf4d2fe63): complete subsection reference.

<a id="canonical-972ac395c91ee828fffe6ffef749aceeb8ef5c972b45e8c6e040e9027f0de88a"></a>

## Next pages — a_pool / d22919da6b10 / 5

- [a_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-5f5a941cdb59ec6ab3ebacdc0c3257ca6d02a9b6dfd1063c5560af63aa7ff3db)
- [a_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-e75e2a29b4df1adefbc44ec1f4fde7c7e28d62347e9021a1eedb63195e0efa78)
- [a_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-7b88fb00da79b580d3fba8811785d895af766df441ef4c0deef646abf4d2fe63)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-5f5a941cdb59ec6ab3ebacdc0c3257ca6d02a9b6dfd1063c5560af63aa7ff3db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d923f19f099954b5d788b077f9a978380f85b346dcd790cf1a028dcdeba479b6"></a>

## a_pool.disable_health_check — a_pool.disable_health_check / dfa60b6800fa / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- a_pool.disable_health_check

<a id="canonical-0d42d1060260750b10643467399e12cf559e0017b07bf0119bba13e146d7eb35"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable health check.

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

<a id="canonical-c83f6d9729856d4b47bffc74e17c9dcc8f27c602ecc512795b8ed3a42bbb910a"></a>

## Direct properties — a_pool.disable_health_check / dfa60b6800fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f6ffa84a2b412ec997772a1a4ddc245d630ec80fb5f0ca0051d94a8b7513000"></a>

## Next pages — a_pool.disable_health_check / dfa60b6800fa / 4

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-e75e2a29b4df1adefbc44ec1f4fde7c7e28d62347e9021a1eedb63195e0efa78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d890790cec12d17dca2ce1374b023c3e07c6cfb5bd24caee6b40350e0cb3b2e1"></a>

## a_pool.health_check — a_pool.health_check / fedef2a6b9d9 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- a_pool.health_check

<a id="canonical-d8574d7e3310d85d1660600fe676fc02a95603a22a4b08913963e2a6eecf0558"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-d332a1d187eac93ba01ff4d56ae63aa1ec0ebdfa9e6bb3b4cacb956395ed5391"></a>

## Direct properties — a_pool.health_check / fedef2a6b9d9 / 3

<a id="canonical-7ddac6e998b9c602114c71961cea6920bb937627b9e651102b43bb320c22fb6d"></a>

<a id="canonical-1d783b35ac38d4b7f2d434555fc9556205504df95d735653b4c5d1921043e292"></a>

## name property — a_pool.health_check / fedef2a6b9d9 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-6ec2d6267e42d83c1090102f15d1b55d6f9f1fe08aaf7730c3797fd418cd84dd"></a>

<a id="canonical-3564ec5b60380c4724e955cc785294b0c6089d0483121323be5a097673221c60"></a>

## namespace property — a_pool.health_check / fedef2a6b9d9 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-99ec7c4ba43ea5ae6750dbf4556ea3b6c01effcb3ce671dd2d63f4a692729fd8"></a>

<a id="canonical-b8316423cc3594d25ac35141e0f138fe5b55c60e38fba824b382f0a808c80dbe"></a>

## tenant property — a_pool.health_check / fedef2a6b9d9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-54ad5f1cc1bd2b0f73294f2b5651b638c13aaf54e8a43a772ebf042515cbb11b"></a>

## Next pages — a_pool.health_check / fedef2a6b9d9 / 7

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-7b88fb00da79b580d3fba8811785d895af766df441ef4c0deef646abf4d2fe63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ddc5a6150bc821d39406345f225dd4751745a0e589acd343138aa8d274be91f"></a>

## a_pool.members — a_pool.members / 082108404758 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- a_pool.members

<a id="canonical-a1a4242a3ff38ab465edcca36d0a3e1ebbc7364da3a585ce3e66103a0db9fd6d"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5eb0dbec66453e2777c4e5401ba50b5d001ce70beac6eef293dfead92d29a648"></a>

## Direct properties — a_pool.members / 082108404758 / 3

<a id="canonical-a63c9470ada9170261760c93160b86d5836fdbf09f834e2fb3015ec3c3a14776"></a>

<a id="canonical-dd48e55a59e3ffefdb2206cb63283c826e2a3087cb90a18ff06732eeb73831ba"></a>

## disable_spec property — a_pool.members / 082108404758 / 4

Type: `"bool"`. Computed.

Value of true will disable the pool-member.

<a id="canonical-dd437ad8f23de33b5cfa9298e42084552e9c76d0b29fbb7c4df1064df0e1911e"></a>

<a id="canonical-72a0bfaf5e05feab7a4c46bf5962b29e92c7710701d8862d0dca345f6bd57763"></a>

## ip_endpoint property — a_pool.members / 082108404758 / 5

Type: `"string"`. Computed.

Public IP. Public IP address.

Upstream description:

Public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-c0d83deca673875be26e5c7e85327e124f3203627a93e3859135711f58afd344"></a>

<a id="canonical-9920e6ce1c281137c40df0484f4d47746cdcb37e756824ab223d56bfeef1bff4"></a>

## name property — a_pool.members / 082108404758 / 6

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-cac32380d23d25d3ffd231de1038a4eb38474352bb3517833077d79d1da7e0bc"></a>

<a id="canonical-2546214a6ea5df18ca125290b7edef64d08ed562b64cc36c4d4c37edb4a7ae22"></a>

## priority property — a_pool.members / 082108404758 / 7

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-b4ee41906e8cac4a26f8a65e301db80ad1b75809f007abb2caff4716af4fc8da"></a>

<a id="canonical-db3c831995915100b6158c29fe106421c67626b45370181ecfba7e39a3dad442"></a>

## ratio property — a_pool.members / 082108404758 / 8

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-23e9e7d21e592ce566d4cd8bdcc97bc1739660db3a1b54dcccd639df2981c8f8"></a>

## Next pages — a_pool.members / 082108404758 / 9

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d56de2d40fc79e99388ea940843e0962d1462d143440f2c05f671764f4e30670)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-d78390dbfa887147f79c03595f13d2397762cc7d74ddf39ffd3b58896122b2da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4b08de402595690adc8114917cf1c1741f9c18026bf1c16a31bb9b0da18011f"></a>

## aaaa_pool — aaaa_pool / 8d0d170088a7 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- aaaa_pool

<a id="canonical-fdffe7f83a9a7e161675bd7a5b68d226f4239527b12f6af89711acbc2b8a91a1"></a>

Type: `"single"`. Computed.

Pool for AAAA Record.

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

<a id="canonical-fa37a007bb3cfa03a4f1d3672e0d4e3e33fdadbd8331ad99d15e9da8b15ed643"></a>

## Direct properties — aaaa_pool / 8d0d170088a7 / 3

<a id="canonical-0fd3d17785dde6472b62e7344450aa2e915c96d95481b426035b43f5da6abc16"></a>

<a id="canonical-31a4f198ec1a48f8e8a0cf577282d47328a9db796f0683ff86871cb57fff1554"></a>

## max_answers property — aaaa_pool / 8d0d170088a7 / 4

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-fcb77756aceea5f92ef2167c4973babac1a0ad867b5be17383791cdb2fa8fb77): complete subsection reference.

<a id="canonical-5700e9c4a28ae31611d996c6a2420072d2188d0299b219dc5533eda3c752944b"></a>

## Next pages — aaaa_pool / 8d0d170088a7 / 5

- [aaaa_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-fcb77756aceea5f92ef2167c4973babac1a0ad867b5be17383791cdb2fa8fb77)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-fcb77756aceea5f92ef2167c4973babac1a0ad867b5be17383791cdb2fa8fb77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f69608fa2f83bfebfade355c666acefd81af20edfbb799d33a3b6a326bdc3c73"></a>

## aaaa_pool.members — aaaa_pool.members / e8aea2d38a2a / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d78390dbfa887147f79c03595f13d2397762cc7d74ddf39ffd3b58896122b2da)
- aaaa_pool.members

<a id="canonical-8ea6f8c8d3e7d306a3d9ecaefaf1b62f862db2c9b4ab9c7b5e3f44090f98b176"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4d4ccf748229667269b0fd0a715deb8d2583863962b084faafd8370c54e3460a"></a>

## Direct properties — aaaa_pool.members / e8aea2d38a2a / 3

<a id="canonical-57db3eaf63596e984ff896b3b7f3376c79e59d19baa3720c192b20043ee16df7"></a>

<a id="canonical-0c87539be29a6b5afd2e27513a3fa5c44697e5b5ac9662a2a0ce5e622d75ada6"></a>

## disable_spec property — aaaa_pool.members / e8aea2d38a2a / 4

Type: `"bool"`. Computed.

Value of true will disable the pool-member.

<a id="canonical-a14e976f3f39ebdb79f968f44a464f4b0f16d3af9b5fa7715cfaf96f4a4b25f9"></a>

<a id="canonical-f9f3bf4b770d40449a02070dd04e745a0141eba615e27e0a6ccb597175bf87ba"></a>

## ip_endpoint property — aaaa_pool.members / e8aea2d38a2a / 5

Type: `"string"`. Computed.

Public IP. Public IP address.

Upstream description:

Public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-ce1dc1cd19fd8f48e7802ef3c7385196bf6289d43cbd918d855acddba02fb50b"></a>

<a id="canonical-210b34a752b50c886e499121d7b24ffd7c68b7d24bef3f2723cf3d53d278f51a"></a>

## name property — aaaa_pool.members / e8aea2d38a2a / 6

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-19ce419ef368952e3b91767ebe04fb1f752ea8d9dbe1162186e035fc6fa6d171"></a>

<a id="canonical-6e7f4d56fa40db8e8ab29e45cfd5dc41c9a3816263c012f7ce08a6e15bdc63a8"></a>

## priority property — aaaa_pool.members / e8aea2d38a2a / 7

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1a5c9d92267a74240076e6bf04d9fd98450a30fbd567a9800d64749f82f8fc67"></a>

<a id="canonical-d135c1edb3606b7cb982490978a4bb0ce483629755b05567a2e7e3bbca666187"></a>

## ratio property — aaaa_pool.members / e8aea2d38a2a / 8

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-b3f05297d3fc5b9f2264829c7bf0de84b1e53f650368f8ac340a00968471e8b0"></a>

## Next pages — aaaa_pool.members / e8aea2d38a2a / 9

- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-d78390dbfa887147f79c03595f13d2397762cc7d74ddf39ffd3b58896122b2da)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdd0bf27397e00916cd6402be75942d0d9717929893c97cce130cac735fbd40b"></a>

## cname_pool — cname_pool / a7a0841639bc / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- cname_pool

<a id="canonical-84eb3b1fc41a65eee05b99b5ab211db907601b58690f37c56740657810d7923f"></a>

Type: `"single"`. Computed.

Pool for CNAME Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

<a id="canonical-7e503491cb7b5abb4b94d0618ada1ac2e08d03f8c47802feba72161765e3aee8"></a>

## Direct properties — cname_pool / a7a0841639bc / 3

- [disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-616d3490eb1158b8ece754ebab47587c39b1c685ab9c08634de11386e7a061af): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-1fbf4e8c5d5273ba15583af5941e95954a5a932639b06c854c572b77f71ca7ca): complete subsection reference.

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-009f6a32455eebb81fa40e5019f1c9663a89da4087b0a79d40493ceec983074f): complete subsection reference.

<a id="canonical-67fcc9f52be0e2d22a91f46cfadf323528d7cac0cdad104f7aff98a99d3577ca"></a>

## Next pages — cname_pool / a7a0841639bc / 4

- [cname_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-616d3490eb1158b8ece754ebab47587c39b1c685ab9c08634de11386e7a061af)
- [cname_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-1fbf4e8c5d5273ba15583af5941e95954a5a932639b06c854c572b77f71ca7ca)
- [cname_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-009f6a32455eebb81fa40e5019f1c9663a89da4087b0a79d40493ceec983074f)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-616d3490eb1158b8ece754ebab47587c39b1c685ab9c08634de11386e7a061af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1da2ee2fcc32ff08866c68f6f6d36c44d9e5377baab378024e726f9beb86f4e9"></a>

## cname_pool.disable_health_check — cname_pool.disable_health_check / 1498aee3a8f9 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- cname_pool.disable_health_check

<a id="canonical-776987590421da2efb3dc9479015f9d4138c8329f3ec6fc9c9e4db40510c4bc2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable health check.

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

<a id="canonical-d16227e2234eb8b451388178d29bcab9d4b74af95b89db3297a0d49347ba8fe7"></a>

## Direct properties — cname_pool.disable_health_check / 1498aee3a8f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ff4c937b8db09f84ed7e2ec59951530014c22834ee53b4beba61279cdbd5dce"></a>

## Next pages — cname_pool.disable_health_check / 1498aee3a8f9 / 4

- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-1fbf4e8c5d5273ba15583af5941e95954a5a932639b06c854c572b77f71ca7ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43c1ae89ca070d3a846e48d9595b6ea479e2a57ef75088c232fd3050f5077be6"></a>

## cname_pool.health_check — cname_pool.health_check / d9b8eab27a30 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- cname_pool.health_check

<a id="canonical-e2358541a4bdab4a373db59296fbe8251ed848da7654662fe58ba31df279abaa"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2de3214a7e39ef3bf81412050eb6fb73f0606693c24fdca74e75735c0678ef9c"></a>

## Direct properties — cname_pool.health_check / d9b8eab27a30 / 3

<a id="canonical-8693e8b8157f709f37d268ed39153e7b88e52743c6a3d5960dd9dc8a7b5b6132"></a>

<a id="canonical-68bd14fd2f7465142649df1ae7a6c5b7d120f626a73b04e310c6051d06603543"></a>

## name property — cname_pool.health_check / d9b8eab27a30 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0927b858fd47f20843a80b9d659d73809bf7fe8b40b36f55f6a0cf3111aff551"></a>

<a id="canonical-7344a03f3f97aff11c17665c7af2f1d7ad63bb7abf9e2a77e177d94d85939993"></a>

## namespace property — cname_pool.health_check / d9b8eab27a30 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-f685cff9bd191f399414fbefaf3aa5becd38607c198a092a9200f8aa54b46868"></a>

<a id="canonical-073c16601e32ca46922dfea36056f36a9d5e3afdde5ad6304b3648f606fd4ba2"></a>

## tenant property — cname_pool.health_check / d9b8eab27a30 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-7af4b500cc7d0528f22fe540d2ec2b09f42f5ae42fb50973a2d308f2f306e5d6"></a>

## Next pages — cname_pool.health_check / d9b8eab27a30 / 7

- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-009f6a32455eebb81fa40e5019f1c9663a89da4087b0a79d40493ceec983074f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e83438cb81300290b93b93a04fac157d5036abb03c361c0d50a143607f9bcfad"></a>

## cname_pool.members — cname_pool.members / a66ceacee95a / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- cname_pool.members

<a id="canonical-ca83742d50bef4f79eb553407ee09b6c5d09e413172ab32d869bbecbc39a4ef3"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3bf81d30736fa61148e51f277fe1070eaf8f6a869ea6b9fc039baf0df130818b"></a>

## Direct properties — cname_pool.members / a66ceacee95a / 3

<a id="canonical-f2df520cb60bd66caa28f9d7010da3cfa1d765852ec6c6858b13eb3425cfd993"></a>

<a id="canonical-75cb09301039859b73ae55f548e0d3ba66bf5e58bdc87450e7a73fcc0e3472c4"></a>

## domain property — cname_pool.members / a66ceacee95a / 4

Type: `"string"`. Computed.

Specifies the fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-aefcab84cab446b188649b0c02e31f81b7bc284e9807d7ed540864fa60888874"></a>

<a id="canonical-3bd0b02208d7b2d9fa3c28acaf048d3cb3d07ae5786b1089e7c83fe182cac05d"></a>

## final_translation property — cname_pool.members / a66ceacee95a / 5

Type: `"bool"`. Computed.

If this flag is true, the CNAME record will not be translated further.

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

<a id="canonical-a6786f8bdb202c2739ea28fa9b46fb0a7ce09dc4721be605814ee32e606ed71f"></a>

<a id="canonical-1b573a7f0fb85f9cb06c27a1250694e56780eac57258de9056ae17efd609ae65"></a>

## name property — cname_pool.members / a66ceacee95a / 6

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-e19ca9d59ac94134397222f3b49920371cb4ded01d95891ef25d2823160d4812"></a>

<a id="canonical-2c085210f17a5bbf95196ec300f56ab7a5c7c26b6a12a97df25bb88d70fff726"></a>

## priority property — cname_pool.members / a66ceacee95a / 7

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-e07c0de17c1ffced72cf495b4f17edbb3a7acd1b43d6b6ab31d990db6c6966b7"></a>

<a id="canonical-910d9751edd56e46c67ca9735fec21e9eb6c589b949e27a4bdcb98a3905c4d91"></a>

## ratio property — cname_pool.members / a66ceacee95a / 8

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-5aa098f47f39f640db4e86c731ab027fe91c1049b19591d182dc2c7c6097ad7b"></a>

## Next pages — cname_pool.members / a66ceacee95a / 9

- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-694514636eb5ccd7ef5411bd2b825f1f003503cb59043fbb008f4bd830f83565)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-2a9f11997dde1afdec1b2b9c47a29d79fe7c2ef6781adfb1a8bc976eac447a81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-838c781eedc4da7c60153a1a4e8e37bb99078008b8c3bfb3f1d48bb0a274b07c"></a>

## mx_pool — mx_pool / f06a353590b7 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- mx_pool

<a id="canonical-531e930254b790101c8a9274746ab4c9706ef8a5fc6a3c06c76b0a0dcd0a87bd"></a>

Type: `"single"`. Computed.

Pool for MX Record.

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

<a id="canonical-120a88572646c28b5f209b58d344f0fe573b07dda1e2ad699b42966e3fb8ad3c"></a>

## Direct properties — mx_pool / f06a353590b7 / 3

<a id="canonical-d7b608e8752a1c8806195e0543909069a22edd3edc0f85e7620cda17fa59399b"></a>

<a id="canonical-d7821ea02db02732c01f3431dd4074921210172cc188ca6a789d0c55b1f2d491"></a>

## max_answers property — mx_pool / f06a353590b7 / 4

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-cba423ffee41d61a310069ba0305b5a7e69201b90da0283990c7f44b4ebd7068): complete subsection reference.

<a id="canonical-58b60c92a0f8fe81d5b8d689e3a1a9ea738bd735e1b02f15e3ede6d41058b05e"></a>

## Next pages — mx_pool / f06a353590b7 / 5

- [mx_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-cba423ffee41d61a310069ba0305b5a7e69201b90da0283990c7f44b4ebd7068)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-cba423ffee41d61a310069ba0305b5a7e69201b90da0283990c7f44b4ebd7068"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44cb467b7b246a5bca93412b62980e4e93d95a6705c306063990dad27cf04cf5"></a>

## mx_pool.members — mx_pool.members / 34af801387a9 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2a9f11997dde1afdec1b2b9c47a29d79fe7c2ef6781adfb1a8bc976eac447a81)
- mx_pool.members

<a id="canonical-96903a0d4a9c8e83ca51ddf4aa0218be180e883a826e4cf6c1a01426859c2ecf"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9ec36e5591c5154e7860050f2bc57cb1f53dacd22ef10d595d563c3ff39c0037"></a>

## Direct properties — mx_pool.members / 34af801387a9 / 3

<a id="canonical-949010a8337d68c80f8d0be3dd94ba99eb359c19baeb2c7a9be9296643c1ed68"></a>

<a id="canonical-bcbc12e82bd6d13e1382e4e7ef5c966603766bca6ccf142b939af65fc980f66a"></a>

## domain property — mx_pool.members / 34af801387a9 / 4

Type: `"string"`. Computed.

Domain name for routing and identification.

Upstream description:

Domain name for routing and identification

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-11b8ee22d305142b102502212b6fe49bf6cd122e59adef6ae4e3765cedc2d9dd"></a>

<a id="canonical-c92ff4c7500a11acaf3e28948a80e9ffa1c80e13e260b479cba6eb888d733671"></a>

## name property — mx_pool.members / 34af801387a9 / 5

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c09cd470682053303e8ef6e57d1a63821254b56e0922f3bae1442605e75f9826"></a>

<a id="canonical-66785abfce37279a1638098d01bab8cb966c7e01195c1e12260f264dac12c333"></a>

## priority property — mx_pool.members / 34af801387a9 / 6

Type: `"number"`. Computed.

MX Record Priority. MX Record priority.

Upstream description:

MX Record priority.

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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-eea8302fb9cd24fd9f33acfaa7cefc13ec205fd364e2d59f07047f457d7a12d6"></a>

<a id="canonical-9f5799ed858206dff435f72bb473214de43a2e50a25f13da4caef972165dfd1c"></a>

## ratio property — mx_pool.members / 34af801387a9 / 7

Type: `"number"`. Computed.

Load Balancing Ratio. Load Balancing Ratio.

Upstream description:

Load Balancing Ratio.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-ac5ef31345a564ac237703ae812f247582c07aab5a190b92a0d06016cc9a0718"></a>

## Next pages — mx_pool.members / 34af801387a9 / 8

- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2a9f11997dde1afdec1b2b9c47a29d79fe7c2ef6781adfb1a8bc976eac447a81)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-4f476814f8d8e2237a9fa285952a58ae91547f2d40311da123957b32f48c0bb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17e199596b0560df5a0943665ec1b012c1b45b479eb765777f32417e784a3d0a"></a>

## srv_pool — srv_pool / e3d1db01c2b3 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- srv_pool

<a id="canonical-2fb1d1ddd7a907e23873283a6b2f07c31bc580a22f878730f5e08b8aeb54e25f"></a>

Type: `"single"`. Computed.

Pool for SRV Record.

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

<a id="canonical-4067c9e8338097de533616d56fdd1081c421d2a2e7e9da01507ab659ce793ce9"></a>

## Direct properties — srv_pool / e3d1db01c2b3 / 3

<a id="canonical-653a21f420ccb7632cfea51aaa6aeca47d05680ceef2d331ed19f74568db5109"></a>

<a id="canonical-d4a499f554cf5381960f3ec573db4d928d55623a92d0a5dfb77a481c8a08c614"></a>

## max_answers property — srv_pool / e3d1db01c2b3 / 4

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-cd2cca31ca6cfec131c0166dc0dea2904b55abf57ea77c1a2b2f39a94cee2f7e): complete subsection reference.

<a id="canonical-2853ec173c04ed9a486144b64de4c8620e0950f0080c778098557285b601a1ba"></a>

## Next pages — srv_pool / e3d1db01c2b3 / 5

- [srv_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-cd2cca31ca6cfec131c0166dc0dea2904b55abf57ea77c1a2b2f39a94cee2f7e)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-cd2cca31ca6cfec131c0166dc0dea2904b55abf57ea77c1a2b2f39a94cee2f7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b5e39c2a4d632176d9df8bdc2e37a114db02e371d558d8761376ccfbbf6983b"></a>

## srv_pool.members — srv_pool.members / 80a281b67db2 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-4f476814f8d8e2237a9fa285952a58ae91547f2d40311da123957b32f48c0bb5)
- srv_pool.members

<a id="canonical-e76fe0ce55f461a20a9316a2f1510ad588a2737c6b100662632eedc44e58c916"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3257e84430e79ff5f2af7edc871199e7b7efb50d75e17c9dc5e2c9cd99cf9e5e"></a>

## Direct properties — srv_pool.members / 80a281b67db2 / 3

<a id="canonical-50d107e013d8abdc6aeb865bf1a8759434d03e936d549ecaeec6cc0d9495bd95"></a>

<a id="canonical-329f0e77b5941f16d476d8e632b9ec21104ad27bb66c9db261990792cd9a32cf"></a>

## final_translation property — srv_pool.members / 80a281b67db2 / 4

Type: `"bool"`. Computed.

If this flag is true, the SRV record will not be translated further.

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

<a id="canonical-b2a7ec36b0a5fcb731203d151a3782dbb12e8362db88c9845019026dc97aecac"></a>

<a id="canonical-29e0695b6b1a214eecf27a69f3333795fbc9960529a94f6aaa7c6e5b452b27b3"></a>

## name property — srv_pool.members / 80a281b67db2 / 5

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b5d12418ae0c9c5f37d6e4fbe4d3167d3f57b9a7d64f316b8a3a4397a4a8cc41"></a>

<a id="canonical-6d629f3b48597d3192488d777ec834d32e6e9db2b71be9c3c40f4ece8a8229eb"></a>

## port property — srv_pool.members / 80a281b67db2 / 6

Type: `"number"`. Computed.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

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
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5e5452b54c1d4c7ff262f0a1264269b4005ccd023706caaa27413c4463310be5"></a>

<a id="canonical-060baaaebf9415730c58c35c61c1fd46241a889ddf3399d7f8abea201889e4b2"></a>

## priority property — srv_pool.members / 80a281b67db2 / 7

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

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
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-8f44a4a01f6bbc62c2c6bbf4408594a4f6e2fd45501f6de0f3fc51ea57397ddf"></a>

<a id="canonical-b5393aa9ff1762618600ad2aba825f1f244e7a6027e78f8b4119519026148c02"></a>

## ratio property — srv_pool.members / 80a281b67db2 / 8

Type: `"number"`. Computed.

Load Balancing Ratio. Configuration parameter for ratio

Upstream description:

Configuration parameter for ratio

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-89a1e69a7e445b2713ab3d5c207245e5596d0dfaaa0b2f9dfcbca9cdf10dfdb0"></a>

<a id="canonical-cebe227e0a16e4acc0c17fe0f490c5a6d95190e23e991823f199e4a8215661f9"></a>

## target property — srv_pool.members / 80a281b67db2 / 9

Type: `"string"`. Computed.

Domain name of the machine providing the service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-5547e166ec99f13be6a971d46f9d34742c9d8e866fb0592b680a2f6b024ee1a7"></a>

<a id="canonical-9fb4da280bafd3b7d632f3d44c808807c3a4bd58c42b3c2c6236b1ada3344374"></a>

## weight property — srv_pool.members / 80a281b67db2 / 10

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

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
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a072d0a20f9e62613a3b592a3979643fecebf4faadf48c6227aa00c3a092a45f"></a>

## Next pages — srv_pool.members / 80a281b67db2 / 11

- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-4f476814f8d8e2237a9fa285952a58ae91547f2d40311da123957b32f48c0bb5)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-d38d32d467bc746f086850c1b52814e108f6277944a51c027513a510ec7003f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6f8ba4891dc7024c58d11de6031788aea3f71ca538224299681e1ad7d92e5a0"></a>

## use_rrset_ttl — use_rrset_ttl / 9f078f0ae8d3 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- use_rrset_ttl

<a id="canonical-c1961692fcb980d6e5d93906168005087e5b536d25127b156de349beaa30873c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use rrset ttl.

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

<a id="canonical-162aa170f76d99a5d967b2922208d4beeee4ea91f45f27b847afbd413b81464a"></a>

## Direct properties — use_rrset_ttl / 9f078f0ae8d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f811a413c936ad04873b446561f25d16970be918ca03f9d98c85290379b9f654"></a>

## Next pages — use_rrset_ttl / 9f078f0ae8d3 / 4

- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
