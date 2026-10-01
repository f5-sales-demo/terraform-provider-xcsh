---
page_title: "xcsh_dns_load_balancer reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer reference."
---

# xcsh_dns_load_balancer reference

<a id="canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b934be2bc00c0ee40f2acf1cabc5e3934de6725685859afcd9e19f97079a65e4"></a>

## Property reference — Property reference / d3e0c6e95775 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- Property reference

<a id="canonical-0a9fe853ef9a427f11162ce673f82d43fc3aaf1d8acb2911337a3e4a72b8c307"></a>

## Direct properties — Property reference / d3e0c6e95775 / 3

<a id="canonical-17ed0e8b33d11dd6a627dd6b072852ee99037034d7fdbdd5f6393ed60cfe6935"></a>

<a id="canonical-5cb11d0895fb6009b481ee48eef7605c1c84a842d78a5d4a005f4a2fed757b46"></a>

## annotations property — Property reference / d3e0c6e95775 / 4

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

<a id="canonical-994ccd1b8b0c3bce27626cedcdf3939c74c0c7b0e80d05bf8ed81024e371f24a"></a>

<a id="canonical-e9eb4add345083cbd802aa36e9af1e2e5b76bdc91b890b6eb609fe3585ec1fcd"></a>

## description property — Property reference / d3e0c6e95775 / 5

Type: `"string"`. Computed.

Description of the DNSLoadBalancer.

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

- [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-91689721d111edcdc2f9bd3b9d4cb2e9d91386c7b2f8c7ace5638f31a61676a4): complete subsection reference.

<a id="canonical-fdf4494078a8c219686c4a3ec55b9f1f9fa516003ab1f48991833510660cdf6a"></a>

<a id="canonical-e0a41e7454c0d0e523e3791215e03a0ac6fdeb761a722b820de23227859f6de2"></a>

## id property — Property reference / d3e0c6e95775 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2bb9c05547ca772f0acd3266884af58ffd5fe7f28b54fe95fc3ce356970eab49"></a>

<a id="canonical-772e348bc8d122cfe33dce4ad19d41ac7b8827858507d2604d607a9b83334666"></a>

## labels property — Property reference / d3e0c6e95775 / 7

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

<a id="canonical-7d4224993bd7fbc48942849e25cce9ef23a955ad18c481d6592528070ad770d4"></a>

<a id="canonical-73ca5c99e90a0ee31ce90126c4978861a840e6c9d728cdefd32e02ff536a8fab"></a>

## name property — Property reference / d3e0c6e95775 / 8

Type: `"string"`. Required.

Name of the DNSLoadBalancer.

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

<a id="canonical-74e255980a8f3407e50f30bdec94c77ffa0ebb718cad5f66ae788f516d5418ca"></a>

<a id="canonical-3355406d3e0077d674813420daa83da5ac2f165dc56fac013551f7b8dc3d5d6b"></a>

## namespace property — Property reference / d3e0c6e95775 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSLoadBalancer exists.

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

<a id="canonical-7cbcdf00de1a5ae2f27f04650581e0bddcb41dc2cfdfc0707089e4b10a6d5f8e"></a>

<a id="canonical-d756b67dea28609bcbef422ce848cc6fbaf4f07a4c01b8d046077bfbdbeae1aa"></a>

## record_type property — Property reference / d3e0c6e95775 / 10

Type: `"string"`. Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Receipt-pinned upstream constraints:

```json
{
  "default": "A",
  "enum": [
    "A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98): complete subsection reference.

- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb): complete subsection reference.

<a id="canonical-4f5606a27ea0b1e2fda37244a620a4644cf9219df4ab6b0256e877bd59095b3f"></a>

## All schema paths — Property reference / d3e0c6e95775 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_load_balancer--reference--group-001.md#canonical-17ed0e8b33d11dd6a627dd6b072852ee99037034d7fdbdd5f6393ed60cfe6935) |
| `description` | [description](data-sources--dns_load_balancer--reference--group-001.md#canonical-994ccd1b8b0c3bce27626cedcdf3939c74c0c7b0e80d05bf8ed81024e371f24a) |
| `fallback_pool` | [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-e8c5129ec6fe62d06fdbe65fd6f60aceec4b7087e5e59fce529171c01942eca0) |
| `fallback_pool.name` | [fallback_pool.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-23c1b82cdb10c19f48aedaa391289627b1ad25ec71d7b567a790516f395e4656) |
| `fallback_pool.namespace` | [fallback_pool.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-2992d0db5003aa4bd8df105b354904e4276c7a61e0a04b0b5712b7d1661591e4) |
| `fallback_pool.tenant` | [fallback_pool.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-71abaa4d730ee8438a7257a00d85530352aa77f3f3fba9af56c516fba7ecfe7b) |
| `id` | [id](data-sources--dns_load_balancer--reference--group-001.md#canonical-fdf4494078a8c219686c4a3ec55b9f1f9fa516003ab1f48991833510660cdf6a) |
| `labels` | [labels](data-sources--dns_load_balancer--reference--group-001.md#canonical-2bb9c05547ca772f0acd3266884af58ffd5fe7f28b54fe95fc3ce356970eab49) |
| `name` | [name](data-sources--dns_load_balancer--reference--group-001.md#canonical-7d4224993bd7fbc48942849e25cce9ef23a955ad18c481d6592528070ad770d4) |
| `namespace` | [namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-74e255980a8f3407e50f30bdec94c77ffa0ebb718cad5f66ae788f516d5418ca) |
| `record_type` | [record_type](data-sources--dns_load_balancer--reference--group-001.md#canonical-7cbcdf00de1a5ae2f27f04650581e0bddcb41dc2cfdfc0707089e4b10a6d5f8e) |
| `response_cache` | [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-7cfe9a95ee328f7eeabab7f9c8fb715c7b07b9edc99527a52d8559b86630b089) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-673da709697284f01a4e70783795f3d92cd04304f0f803169b9e146769056456) |
| `response_cache.disable_spec` | [response_cache.disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-9751e454dac57df5f303486541a540656adcf17c88aa0348dbcd0a8345b7519c) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-d8767fca9b3a67422b0a6bcb45a347d239d4c78318878513044f219888315895) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](data-sources--dns_load_balancer--reference--group-001.md#canonical-f41c8e4a086c566852cd3704c78074efb9e7386c0d73ea6030cd9cf01ec02963) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](data-sources--dns_load_balancer--reference--group-001.md#canonical-cd6ae1051ab8884672239bf2f002db23f1cf1366077aed7c7af34532395b66bc) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](data-sources--dns_load_balancer--reference--group-001.md#canonical-f66abbe64c45a620c086d6b5e48b8f34e7f91cc2044d2a2ff733642572c5fda4) |
| `rule_list` | [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-0d171646501b635a47af94705300296f6729ca3b114dad351c11aef5523ff836) |
| `rule_list.rules` | [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-63ce5123c65f129dc9ca0b72d50398bcd1d4cc27b3d7a7d3a59c21aa73c4d106) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-ecdd3dd2ceaddbc8cd8ae7ee4cb39bc21c4b413a0922789f43448ebd380a15c1) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](data-sources--dns_load_balancer--reference--group-001.md#canonical-03986889da197ede11cb7a8c600df3df454895a271d817f65ea6f5a723b5ec18) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-b688745454476e1d17fd945e5c39c6ecea8aad0300f02687600499b27cffd074) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-8429025c0faec37bb9960ba1a4e45673da6bd54bf00bd1b0b6dbd9a01408289f) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](data-sources--dns_load_balancer--reference--group-001.md#canonical-c45f50851fee88991d187f75c026f0499f6d2a8bee96bc59633ac37aa7aeba32) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-7dec2b42c54635f2fe40013d710cc76014234a273fa40bff06a1b32ee2645422) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-3cc6285c98895289f67ccd7269ede69ac32ec4c12453afb356b5f34d29c4c1c1) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-30cb40d4bc8835949c88c082b15fbf7043f482ba33141039b926672c2e365882) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](data-sources--dns_load_balancer--reference--group-001.md#canonical-bd1a8d110dbeaa6c3cbd971305f4518bb54a5ee7fc1297542a6e0ea5d609985d) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-fc9451f15d642b74a61208be4ac0c04f3c274d21bf738817e4712cf5f4d587c4) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](data-sources--dns_load_balancer--reference--group-001.md#canonical-27d51df4a7571e232ecb59c33df8f4ade8e83f594379663989f1b0e64b998916) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-61f33c75d032d45d69dc8259a2d0a16bc2011ca7703af9951b479cc90687e849) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-2714ebcde88ad02e85c2a18cd17216b0d2d553945efb06028afd6fad43e2a2e2) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-c4c1eb984c0b3c51000272335d6f7249090e48658e4db8d3182457f75dcdfbc6) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-af3eff618c9ebd2c08e0917846dea2e144c40bb0846226e0fc3599aa636303eb) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-97e80f8caef86d7f8c426e54032259842f14d65a500d5a4bd83e5860dd404928) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](data-sources--dns_load_balancer--reference--group-001.md#canonical-05424b6a4fdf697bbdd4c67a14878cd56eb5d4434c6f15bb9cfee8e61e9bad2c) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](data-sources--dns_load_balancer--reference--group-001.md#canonical-0d29b02f34a42f06862546532bb3ef59457722461d409a4e5bf730891bc32514) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-fe3dd5cfeba252d9bb0c8779dd274eb5bd57eddb6d737b87ffc189a7038afc71) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-4a8d9e4f3b1cc5d11b899a11a430babfa97a01ffeae8f8a5b6bf00f7fe7589e6) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-15f28ebeaa6b793ee7bd989662ad2a2e8bacf8c78b31804017314dbdd7d751ee) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](data-sources--dns_load_balancer--reference--group-001.md#canonical-dae4cd923592dc2cfd0ad7dff9d96eca94e23e9496412853d9a70506176b7f92) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-8d2b354bed8099245e06ac27bb1c77ebe6bb3cb79440408601d3bb5ef49855b9) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-0c7cd24027917f60a91f13895dc315aff7dd55357c329957e4bb0f263679ada7) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-4ea4fa4013cc9457a3d7611deb0e7da8fb1bbaee0557bd3ad588e26f18585d4f) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](data-sources--dns_load_balancer--reference--group-001.md#canonical-8978222a9d09130d529f541b5df5a849ed6f346c5afe1b3c2c3637ff7a6bcd06) |
| `rule_list.rules.pool` | [rule_list.rules.pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-2753f32a16f272891a5c07bde76f2469c1de47edda68119af793e9434185b308) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-249e46d2687ff70ddfa314e98e3a2e7de0b399ec2a2615ab8711a0dc9a98feeb) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-c501d6333d1092a9e087d12d46567480ca482631edb0e9789523ee2fc8e48164) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-6713a588ddc8e38b6b0f065add59c48d03121b7ae2abb3dd7b58cf66a2a29005) |
| `rule_list.rules.score` | [rule_list.rules.score](data-sources--dns_load_balancer--reference--group-001.md#canonical-8a214d41a279f96938e2629addb1f8e4383852e3a73ba2a3b0d417c7eb081b85) |

<a id="canonical-59d65c3b1138f2129edbc4a39b99594c03a648d42b3ed793065eda17e8bbd2a4"></a>

## Next pages — Property reference / d3e0c6e95775 / 12

- [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-91689721d111edcdc2f9bd3b9d4cb2e9d91386c7b2f8c7ace5638f31a61676a4)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-91689721d111edcdc2f9bd3b9d4cb2e9d91386c7b2f8c7ace5638f31a61676a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-090e84b7df74105f5b3d3b08611ec0255dc76f65992b91ef4d340a4f2a8dfb57"></a>

## fallback_pool — fallback_pool / 7b54a2afaf5e / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- fallback_pool

<a id="canonical-e8c5129ec6fe62d06fdbe65fd6f60aceec4b7087e5e59fce529171c01942eca0"></a>

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

<a id="canonical-7aa6c43c839f8b1f4e1175183d900f3b3adc7849d4233aaf4e48e9fab537ab25"></a>

## Direct properties — fallback_pool / 7b54a2afaf5e / 3

<a id="canonical-23c1b82cdb10c19f48aedaa391289627b1ad25ec71d7b567a790516f395e4656"></a>

<a id="canonical-3b1543ef3215df36d89620f61d6dd584173d40b3ea9658407045a51ab940c901"></a>

## name property — fallback_pool / 7b54a2afaf5e / 4

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

<a id="canonical-2992d0db5003aa4bd8df105b354904e4276c7a61e0a04b0b5712b7d1661591e4"></a>

<a id="canonical-12145881415eb46feecd8405cb7053cf8a2fc5d3ae2a2e94bc716b6d012118ff"></a>

## namespace property — fallback_pool / 7b54a2afaf5e / 5

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

<a id="canonical-71abaa4d730ee8438a7257a00d85530352aa77f3f3fba9af56c516fba7ecfe7b"></a>

<a id="canonical-e8f875e2542bdfd407fc0661bdd97a73c674810199c4495e568a5b32f8c8ced6"></a>

## tenant property — fallback_pool / 7b54a2afaf5e / 6

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

<a id="canonical-13dc580b4088ec1150ffa954b13c658f0d80e39ca4e10cd083049dc125290d88"></a>

## Next pages — fallback_pool / 7b54a2afaf5e / 7

- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7ea0feeba35037e7b8e40daeda329ac93cd130daeb78b8e027c641e23c608e4"></a>

## response_cache — response_cache / f084afb90bfa / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- response_cache

<a id="canonical-7cfe9a95ee328f7eeabab7f9c8fb715c7b07b9edc99527a52d8559b86630b089"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

<a id="canonical-7f6ce2c8c2b6a4207be96858580d9de5aa05d276226db20bea1a7b8f6e28f471"></a>

## Direct properties — response_cache / f084afb90bfa / 3

- [default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-0f175a1aee5f3a73460d007bab34aa29ed4dc4a37f290aa6e7670dd7fda739a8): complete subsection reference.

- [disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-755453700c8cf3b0bb7b4f736cb499e4b28dbbf46ee53734a22a9d7c5273ea5f): complete subsection reference.

- [response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-fabd9aa96e2129400396344b1a8168a6d0dd72f29c65845bcba7325b91cd2c0f): complete subsection reference.

<a id="canonical-28ba219f417b2e8d53673c0323d8c6f7a5a246917e9a129761f75fb27d7a29cd"></a>

## Next pages — response_cache / f084afb90bfa / 4

- [response_cache.default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-0f175a1aee5f3a73460d007bab34aa29ed4dc4a37f290aa6e7670dd7fda739a8)
- [response_cache.disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-755453700c8cf3b0bb7b4f736cb499e4b28dbbf46ee53734a22a9d7c5273ea5f)
- [response_cache.response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-fabd9aa96e2129400396344b1a8168a6d0dd72f29c65845bcba7325b91cd2c0f)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-0f175a1aee5f3a73460d007bab34aa29ed4dc4a37f290aa6e7670dd7fda739a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3f616973f973c4d6d24aae0e6e44fd6f0468085ad52653382669b226fa1b1b5"></a>

## response_cache.default_response_cache_parameters — response_cache.default_response_cache_parameters / b8ff8dc1a52a / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- response_cache.default_response_cache_parameters

<a id="canonical-673da709697284f01a4e70783795f3d92cd04304f0f803169b9e146769056456"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default response cache parameters.

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

<a id="canonical-5a770249583f95ae6e1a2760ed24e2090119e1ba9036ed6169e035ac007e2b15"></a>

## Direct properties — response_cache.default_response_cache_parameters / b8ff8dc1a52a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ee9b2efd79fd971ef9e6d835cf0ff0480129a40d6d2b5a90f4e5f2411de1cdb"></a>

## Next pages — response_cache.default_response_cache_parameters / b8ff8dc1a52a / 4

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-755453700c8cf3b0bb7b4f736cb499e4b28dbbf46ee53734a22a9d7c5273ea5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02a954fccf737cb3c71fceb69f7ca263eca23d7b381b4d385fc07168e88a2396"></a>

## response_cache.disable_spec — response_cache.disable_spec / 0c2f8535778c / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- response_cache.disable_spec

<a id="canonical-9751e454dac57df5f303486541a540656adcf17c88aa0348dbcd0a8345b7519c"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-46492b7131716050597f27dbac5c10f345565b0fd480992aa048d5d8570ee581"></a>

## Direct properties — response_cache.disable_spec / 0c2f8535778c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cbf59997f825cbf3ceaf8308785bc11cf11f220ded6ed42c9a13fc036b1dbc97"></a>

## Next pages — response_cache.disable_spec / 0c2f8535778c / 4

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-fabd9aa96e2129400396344b1a8168a6d0dd72f29c65845bcba7325b91cd2c0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75867aea1a1674e7a984570af609728daafa6a0a47cd57afc21c91a0f0a4b44b"></a>

## response_cache.response_cache_parameters — response_cache.response_cache_parameters / 7045571eabef / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- response_cache.response_cache_parameters

<a id="canonical-d8767fca9b3a67422b0a6bcb45a347d239d4c78318878513044f219888315895"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache parameters.

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

<a id="canonical-49ab2a384c48bf07b1ac3bc5b35d0c62de21daf430376c53822c96791f85656c"></a>

## Direct properties — response_cache.response_cache_parameters / 7045571eabef / 3

<a id="canonical-f41c8e4a086c566852cd3704c78074efb9e7386c0d73ea6030cd9cf01ec02963"></a>

<a id="canonical-8f09ee8719a5413e1476eaa9f94be9ebfa171017f7687b3d6b6662fb7f410f58"></a>

## cache_cidr_ipv4 property — response_cache.response_cache_parameters / 7045571eabef / 4

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv4 clients.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-cd6ae1051ab8884672239bf2f002db23f1cf1366077aed7c7af34532395b66bc"></a>

<a id="canonical-b9225da4535da63aa54a5bd5baf41ae79c064cce39d4bdd41e607459b574b8e1"></a>

## cache_cidr_ipv6 property — response_cache.response_cache_parameters / 7045571eabef / 5

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv6 clients.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-f66abbe64c45a620c086d6b5e48b8f34e7f91cc2044d2a2ff733642572c5fda4"></a>

<a id="canonical-b812ce67b820c651f47c269a2c57dd0805ee6fd5d38c0629d8ca5e883b5f5201"></a>

## cache_ttl property — response_cache.response_cache_parameters / 7045571eabef / 6

Type: `"number"`. Computed.

TTL. TTL for response cache.

Upstream description:

TTL for response cache.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

<a id="canonical-d3ed899232fe4d73a754106009e98890fd9a5962b71fb693bc34abf9b7d9f3c1"></a>

## Next pages — response_cache.response_cache_parameters / 7045571eabef / 7

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-4aa10dc173dfc56e307ab26d5be6f96d536cd4a18e0a3cbe4c579f49c05bba98)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61ecff34a7a393a3f45c5fe263cce06c364e7f7f8c4b4ec945d35d72b5d4046b"></a>

## rule_list — rule_list / 1fafee4287ad / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- rule_list

<a id="canonical-0d171646501b635a47af94705300296f6729ca3b114dad351c11aef5523ff836"></a>

Type: `"single"`. Computed.

Load Balancing Rule List. List of the Load Balancing Rules.

Upstream description:

List of the Load Balancing Rules.

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

<a id="canonical-a6fe44904633383a773e873bc769ad1d413cc6cd879797219ff1911b8db689e9"></a>

## Direct properties — rule_list / 1fafee4287ad / 3

- [rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53): complete subsection reference.

<a id="canonical-10f3cebe241c6f01880d6c2a68265944d0f78db8f5e6d7165f60ec590437a607"></a>

## Next pages — rule_list / 1fafee4287ad / 4

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f26a96860cc2de23609687871fdf0f1bfc2792d29883ce321f914a1ddebb7bad"></a>

## rule_list.rules — rule_list.rules / 8287a953e097 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- rule_list.rules

<a id="canonical-63ce5123c65f129dc9ca0b72d50398bcd1d4cc27b3d7a7d3a59c21aa73c4d106"></a>

Type: `"list"`. Computed.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4ee3d482f9bda1450c67597154d8796808ec5dfc61ec43fb50484b72875f6981"></a>

## Direct properties — rule_list.rules / 8287a953e097 / 3

- [asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-36cef2bb5c3699c9f1ea9f573595512316d11649e23eeb6e59bf06d72b1143bb): complete subsection reference.

- [asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-b654e638e8a9c760817a8ade7179f78218ea676d58fdff692fa4f8af5d6b0b89): complete subsection reference.

- [geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-6ba93adef43510d21004dc9452a2dc60c0ae5580287681797bd1044f997e9bec): complete subsection reference.

- [geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-910651eaf4f230193a89b38690bbd37af41d6fde0262e03ecc8ad4af957c2f82): complete subsection reference.

- [ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-4eed935f58ce99166233a6300425b1bc432240071250442f56023e0743f61b0a): complete subsection reference.

- [ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-d364457067a34a40f2b4c16d21959d96b7cb8c5e6f2575fbf68c8167e39f964d): complete subsection reference.

- [pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-16577a3d7cdfd1251e300f550473bfbdb69d56de0fd60814ba4745307269a9bb): complete subsection reference.

<a id="canonical-8a214d41a279f96938e2629addb1f8e4383852e3a73ba2a3b0d417c7eb081b85"></a>

<a id="canonical-ae579e56558d731e4f13f291b292a2e3a2867fbfd45bf81f26ebadaa18136ba2"></a>

## score property — rule_list.rules / 8287a953e097 / 4

Type: `"number"`. Computed.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32767,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  }
}
```

<a id="canonical-41fd373a0c46f73ab1e8c24f817983a4e643cdabc402b51f8c88862ecc9499da"></a>

## Next pages — rule_list.rules / 8287a953e097 / 5

- [rule_list.rules.asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-36cef2bb5c3699c9f1ea9f573595512316d11649e23eeb6e59bf06d72b1143bb)
- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-b654e638e8a9c760817a8ade7179f78218ea676d58fdff692fa4f8af5d6b0b89)
- [rule_list.rules.geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-6ba93adef43510d21004dc9452a2dc60c0ae5580287681797bd1044f997e9bec)
- [rule_list.rules.geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-910651eaf4f230193a89b38690bbd37af41d6fde0262e03ecc8ad4af957c2f82)
- [rule_list.rules.ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-4eed935f58ce99166233a6300425b1bc432240071250442f56023e0743f61b0a)
- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-d364457067a34a40f2b4c16d21959d96b7cb8c5e6f2575fbf68c8167e39f964d)
- [rule_list.rules.pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-16577a3d7cdfd1251e300f550473bfbdb69d56de0fd60814ba4745307269a9bb)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-36cef2bb5c3699c9f1ea9f573595512316d11649e23eeb6e59bf06d72b1143bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acf6037252943645581ddb5433759f123e3ad040af6136fc222f328e8e1d2555"></a>

## rule_list.rules.asn_list — rule_list.rules.asn_list / 092f0607adbd / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.asn_list

<a id="canonical-ecdd3dd2ceaddbc8cd8ae7ee4cb39bc21c4b413a0922789f43448ebd380a15c1"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-d1ecef70945d942f10781549c0fdc6de93f1410bf3c4fd0c47781d5af006c224"></a>

## Direct properties — rule_list.rules.asn_list / 092f0607adbd / 3

<a id="canonical-03986889da197ede11cb7a8c600df3df454895a271d817f65ea6f5a723b5ec18"></a>

<a id="canonical-7a2f77814b0adddc0deff7656ab2409c22cb2c00917a210ca06b503925c6dae9"></a>

## as_numbers property — rule_list.rules.asn_list / 092f0607adbd / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6a050547bf15ec8f0cdc430edc199e241a0a2e59d9227e967b13b239a8a0558f"></a>

## Next pages — rule_list.rules.asn_list / 092f0607adbd / 5

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-b654e638e8a9c760817a8ade7179f78218ea676d58fdff692fa4f8af5d6b0b89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1de24bf4d03d4c634c938679e2aad2e76ba2c5372934d774411121077e990d6b"></a>

## rule_list.rules.asn_matcher — rule_list.rules.asn_matcher / adb6b3b88f9f / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.asn_matcher

<a id="canonical-b688745454476e1d17fd945e5c39c6ecea8aad0300f02687600499b27cffd074"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-e5b77ff700218c86847e882bccdc99b1ec17d1989e3c3ff31017531de66afcb8"></a>

## Direct properties — rule_list.rules.asn_matcher / adb6b3b88f9f / 3

- [asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-f432c788b10112e671a5157ec12ef3c6c714da3493db277ef70eca748c6b0068): complete subsection reference.

<a id="canonical-8b1cc685d87b29fb5bc6520ee698789106cec6cf6e45a660b4f4039d1547a773"></a>

## Next pages — rule_list.rules.asn_matcher / adb6b3b88f9f / 4

- [rule_list.rules.asn_matcher.asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-f432c788b10112e671a5157ec12ef3c6c714da3493db277ef70eca748c6b0068)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-f432c788b10112e671a5157ec12ef3c6c714da3493db277ef70eca748c6b0068"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1616797b489896924e0bcca75366f9dfcc45e4d097a1f7b8d8eee63ec48cd28"></a>

## rule_list.rules.asn_matcher.asn_sets — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-b654e638e8a9c760817a8ade7179f78218ea676d58fdff692fa4f8af5d6b0b89)
- rule_list.rules.asn_matcher.asn_sets

<a id="canonical-8429025c0faec37bb9960ba1a4e45673da6bd54bf00bd1b0b6dbd9a01408289f"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-7bc2d1fbe4243cced0b1bb834ac698346404fc45692fa5ecf1a29c2321442c77"></a>

## Direct properties — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 3

<a id="canonical-c45f50851fee88991d187f75c026f0499f6d2a8bee96bc59633ac37aa7aeba32"></a>

<a id="canonical-8edb105c241cc237e573f0b3ae2ef28942b39dd35925e2ad2c104ef747478086"></a>

## kind property — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
  }
}
```

<a id="canonical-7dec2b42c54635f2fe40013d710cc76014234a273fa40bff06a1b32ee2645422"></a>

<a id="canonical-39df92b95193239fa1d51b84338cd0ad607d8352069bdccaa198897076a9860b"></a>

## name property — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
  }
}
```

<a id="canonical-3cc6285c98895289f67ccd7269ede69ac32ec4c12453afb356b5f34d29c4c1c1"></a>

<a id="canonical-bf13c44c4855dedea2fa78c4cd493eb5c7bde83aeb60dc1aa31a78beb3becf1a"></a>

## namespace property — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-30cb40d4bc8835949c88c082b15fbf7043f482ba33141039b926672c2e365882"></a>

<a id="canonical-0f8b25dd4c375ea36a63ed3b72c647c2dcd02b3cbb50a54e6627c88f23a04d74"></a>

## tenant property — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
  }
}
```

<a id="canonical-bd1a8d110dbeaa6c3cbd971305f4518bb54a5ee7fc1297542a6e0ea5d609985d"></a>

<a id="canonical-e67bfa103045d036b8f451599557f2a8f8c6588545ebf7a0472a6a9e11a083ab"></a>

## uid property — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
  }
}
```

<a id="canonical-2c44656cca84d099ca187164aa7becf0c3046e73f1c32551b0193c7b268d9ee9"></a>

## Next pages — rule_list.rules.asn_matcher.asn_sets / f07efa2ae9dd / 9

- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-b654e638e8a9c760817a8ade7179f78218ea676d58fdff692fa4f8af5d6b0b89)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-6ba93adef43510d21004dc9452a2dc60c0ae5580287681797bd1044f997e9bec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a02f6f919ddae0732b46ab8a130467fdd3468747dfb8f41aa6dd7ecc28bfa584"></a>

## rule_list.rules.geo_location_label_selector — rule_list.rules.geo_location_label_selector / a5fb22e1b0ac / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.geo_location_label_selector

<a id="canonical-fc9451f15d642b74a61208be4ac0c04f3c274d21bf738817e4712cf5f4d587c4"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-5b04a6252a8aea78c1950955077f6f54c561c06d7cf02cb3ccbf372cfa036018"></a>

## Direct properties — rule_list.rules.geo_location_label_selector / a5fb22e1b0ac / 3

<a id="canonical-27d51df4a7571e232ecb59c33df8f4ade8e83f594379663989f1b0e64b998916"></a>

<a id="canonical-5c629fa409188b46bf6ea370a7feaf205da21711171db0f763ffb15ad773a22a"></a>

## expressions property — rule_list.rules.geo_location_label_selector / a5fb22e1b0ac / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-7535cdfbaa9fd8c1d8a03d8ce037dfeecedf7f77602f3190e8876e64357441d9"></a>

## Next pages — rule_list.rules.geo_location_label_selector / a5fb22e1b0ac / 5

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-910651eaf4f230193a89b38690bbd37af41d6fde0262e03ecc8ad4af957c2f82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5365472afdec3b61bcb91c233c747a8dcbaa653de80bf0ca756cc24b43138800"></a>

## rule_list.rules.geo_location_set — rule_list.rules.geo_location_set / ed3e1303896f / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.geo_location_set

<a id="canonical-61f33c75d032d45d69dc8259a2d0a16bc2011ca7703af9951b479cc90687e849"></a>

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

<a id="canonical-131ea501293085ff1be61cb06b212741142c6b4d741a32bbe292a41e68dbc75d"></a>

## Direct properties — rule_list.rules.geo_location_set / ed3e1303896f / 3

<a id="canonical-2714ebcde88ad02e85c2a18cd17216b0d2d553945efb06028afd6fad43e2a2e2"></a>

<a id="canonical-fb977aa7dceecbca7f56639c1080fcfa639664eafc4742b3072984489f541fbc"></a>

## name property — rule_list.rules.geo_location_set / ed3e1303896f / 4

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

<a id="canonical-c4c1eb984c0b3c51000272335d6f7249090e48658e4db8d3182457f75dcdfbc6"></a>

<a id="canonical-43e1de370257eb95143c01300fcd9df6b37737ea8fd2b9a7e3d2d513676dad57"></a>

## namespace property — rule_list.rules.geo_location_set / ed3e1303896f / 5

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

<a id="canonical-af3eff618c9ebd2c08e0917846dea2e144c40bb0846226e0fc3599aa636303eb"></a>

<a id="canonical-c36f4f3c363019143fc96eaf52b4e0a8a6d6f852adf4f9dde7003d5aa21b0207"></a>

## tenant property — rule_list.rules.geo_location_set / ed3e1303896f / 6

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

<a id="canonical-189d7aeda445ab188ab892e7c9972a0c47597e61daef9f8e5f9c885ed6008aab"></a>

## Next pages — rule_list.rules.geo_location_set / ed3e1303896f / 7

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-4eed935f58ce99166233a6300425b1bc432240071250442f56023e0743f61b0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d102c8918844dc5967b1e9ebab5a329ef6cd1c8dd33c1bb37c8c8fd1846d49b0"></a>

## rule_list.rules.ip_prefix_list — rule_list.rules.ip_prefix_list / 56e61cfa659d / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.ip_prefix_list

<a id="canonical-97e80f8caef86d7f8c426e54032259842f14d65a500d5a4bd83e5860dd404928"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-3c3f73243d9ccca3a9bfbcc811548e7ba91758287446ed9ccb99a4fb1c6a5932"></a>

## Direct properties — rule_list.rules.ip_prefix_list / 56e61cfa659d / 3

<a id="canonical-05424b6a4fdf697bbdd4c67a14878cd56eb5d4434c6f15bb9cfee8e61e9bad2c"></a>

<a id="canonical-f866af56f8ce3aa8bd6eed3cddfa6f0b42487441f4dad46ffe85eadd5eaaba5b"></a>

## invert_match property — rule_list.rules.ip_prefix_list / 56e61cfa659d / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-0d29b02f34a42f06862546532bb3ef59457722461d409a4e5bf730891bc32514"></a>

<a id="canonical-6a5bd5403f1fe55bdf6b0fe60d6d2198cb51c281d75a3d26b72ef27360cdfe09"></a>

## ip_prefixes property — rule_list.rules.ip_prefix_list / 56e61cfa659d / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a4e3c62f5ec06802ec1ef6f13038c70d0f5a61b68778b410c29c62ec29d44499"></a>

## Next pages — rule_list.rules.ip_prefix_list / 56e61cfa659d / 6

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-d364457067a34a40f2b4c16d21959d96b7cb8c5e6f2575fbf68c8167e39f964d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6feb0d5159b3ddb434f235dcfa00273cd4256cda3cbe977ecaf1ca642b9c42c0"></a>

## rule_list.rules.ip_prefix_set — rule_list.rules.ip_prefix_set / 2209669853a6 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.ip_prefix_set

<a id="canonical-fe3dd5cfeba252d9bb0c8779dd274eb5bd57eddb6d737b87ffc189a7038afc71"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-403d4cfc06b2e6e5afdc5456e797a502031f4f0a50441e9250609356430bf5f6"></a>

## Direct properties — rule_list.rules.ip_prefix_set / 2209669853a6 / 3

<a id="canonical-4a8d9e4f3b1cc5d11b899a11a430babfa97a01ffeae8f8a5b6bf00f7fe7589e6"></a>

<a id="canonical-d8b09201f617693b5dc3fb3d729662ee1df592f19920397c53bf3534e40993b1"></a>

## invert_matcher property — rule_list.rules.ip_prefix_set / 2209669853a6 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-e8d94104f331029157b14fb2a71facd917042b252d9ad9008be269d0dcdc46b2): complete subsection reference.

<a id="canonical-7952a0996f8eaf66beee509a5257f465b5cd417023d539624762b37ccabda81f"></a>

## Next pages — rule_list.rules.ip_prefix_set / 2209669853a6 / 5

- [rule_list.rules.ip_prefix_set.prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-e8d94104f331029157b14fb2a71facd917042b252d9ad9008be269d0dcdc46b2)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-e8d94104f331029157b14fb2a71facd917042b252d9ad9008be269d0dcdc46b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d580ffd4dff0c08274d3f565e98622867a8a37c84db580db75006594d1c459d"></a>

## rule_list.rules.ip_prefix_set.prefix_sets — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-d364457067a34a40f2b4c16d21959d96b7cb8c5e6f2575fbf68c8167e39f964d)
- rule_list.rules.ip_prefix_set.prefix_sets

<a id="canonical-15f28ebeaa6b793ee7bd989662ad2a2e8bacf8c78b31804017314dbdd7d751ee"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-70bcaab4404f6afce321b39cb9c57f8342ca4ec2b6adc947c246b0bd03f29739"></a>

## Direct properties — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 3

<a id="canonical-dae4cd923592dc2cfd0ad7dff9d96eca94e23e9496412853d9a70506176b7f92"></a>

<a id="canonical-18fdac1eae87dbef0076193ebc541e3f707e4091638fe7e4c4d7b335c8d2a89a"></a>

## kind property — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
  }
}
```

<a id="canonical-8d2b354bed8099245e06ac27bb1c77ebe6bb3cb79440408601d3bb5ef49855b9"></a>

<a id="canonical-47b866dff0898f87f72832cdb91e89db7d5cbb6dcbaeda4585f86dec73d35976"></a>

## name property — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
  }
}
```

<a id="canonical-0c7cd24027917f60a91f13895dc315aff7dd55357c329957e4bb0f263679ada7"></a>

<a id="canonical-1f7fb37a31618226702a8766343284ee93cb6e046de07e1a896e21dd21602935"></a>

## namespace property — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4ea4fa4013cc9457a3d7611deb0e7da8fb1bbaee0557bd3ad588e26f18585d4f"></a>

<a id="canonical-5055563513a928cfc5cbda3b85a761d79f51aa3eeb8b8e850a05ceb7f8d52943"></a>

## tenant property — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
  }
}
```

<a id="canonical-8978222a9d09130d529f541b5df5a849ed6f346c5afe1b3c2c3637ff7a6bcd06"></a>

<a id="canonical-544ad98812242efbef212fb5f8aad5490dace39c769ee6d58adbf1b91ee81f95"></a>

## uid property — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
  }
}
```

<a id="canonical-29c05e89d52bf2c09f3de7f0390e9d21a36b9513d8a52580b7c296a949bdcb00"></a>

## Next pages — rule_list.rules.ip_prefix_set.prefix_sets / c8025d31fc28 / 9

- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-d364457067a34a40f2b4c16d21959d96b7cb8c5e6f2575fbf68c8167e39f964d)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-16577a3d7cdfd1251e300f550473bfbdb69d56de0fd60814ba4745307269a9bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a90d6ac2607adf0f000fdce10d66bad0433a711c9324d67a6ab52f28a493916"></a>

## rule_list.rules.pool — rule_list.rules.pool / 10165f41cacf / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-de83d2932248cb9be96fc334defc36e2686068cd9ef5f35213e8dc567cd904aa)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-60cbc8179586ebaa383956e20a2e19533dddb1e726eaf9dcdcadf636fb3a4deb)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- rule_list.rules.pool

<a id="canonical-2753f32a16f272891a5c07bde76f2469c1de47edda68119af793e9434185b308"></a>

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

<a id="canonical-14855b892ec1dd2df986ab9bc96ae1c1aa065bdc0102131b5b883f5d25a40f5c"></a>

## Direct properties — rule_list.rules.pool / 10165f41cacf / 3

<a id="canonical-249e46d2687ff70ddfa314e98e3a2e7de0b399ec2a2615ab8711a0dc9a98feeb"></a>

<a id="canonical-66aa8fbab47481dc97f565696a0f7df7c9320af4f82c896a38f3881473dd2037"></a>

## name property — rule_list.rules.pool / 10165f41cacf / 4

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

<a id="canonical-c501d6333d1092a9e087d12d46567480ca482631edb0e9789523ee2fc8e48164"></a>

<a id="canonical-5f393c7cd0b37aaf7a7b3a319bf33cc00025349f146362f4911b29562e37915b"></a>

## namespace property — rule_list.rules.pool / 10165f41cacf / 5

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

<a id="canonical-6713a588ddc8e38b6b0f065add59c48d03121b7ae2abb3dd7b58cf66a2a29005"></a>

<a id="canonical-7c0c052ecc4b466c097f40e57bbb6379accf1008f1cb0f7981f41a758afed091"></a>

## tenant property — rule_list.rules.pool / 10165f41cacf / 6

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

<a id="canonical-6f65ef3aa8653c2dc3d47257307de7e1bb33df69147bf88c8d2986cc30b56ec1"></a>

## Next pages — rule_list.rules.pool / 10165f41cacf / 7

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1e206e8b1dfbb7bb7209903af431388775262d4ff7606a917e1db6c03fc58f53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
