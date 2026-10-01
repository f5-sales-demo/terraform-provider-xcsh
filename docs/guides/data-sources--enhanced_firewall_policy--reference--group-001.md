---
page_title: "xcsh_enhanced_firewall_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy reference."
---

# xcsh_enhanced_firewall_policy reference

<a id="canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ac16188a59368bdf5185408fbd9717c5cf732e6691cc95a57d777a31910b2f5"></a>

## Property reference — Property reference / 1b34f62e9ecc / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- Property reference

<a id="canonical-412579144851fa32237a58e114cc5935e59db464443d00d032ba121cb6c70e7e"></a>

## Direct properties — Property reference / 1b34f62e9ecc / 3

- [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-cfcfc1c5f374870b18c521d8bddf34f8136c3c8a886432c73d1b7f24ab6aae80): complete subsection reference.

- [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-59aeaa8f794d56e5bacdc5a578af0339fcceeb74cdfc3828817bf02fd8be9d61): complete subsection reference.

- [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3461c3a8e44ade46badb25f81bc7040807e6c26f425354aae577fd33c070f359): complete subsection reference.

<a id="canonical-2ecc40d36691b8884cdcbfe37f6b11391d2b607f58f73939c923c2d8b1a5deef"></a>

<a id="canonical-c0b69cf14235f7de80c3af29cc74a5459cddf71ffa63d60a7db822858e78f490"></a>

## annotations property — Property reference / 1b34f62e9ecc / 4

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

- [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-39413b52c833c9e4a303b5835c0624caaf50c3752a865bbe85dd0387ca295ebb): complete subsection reference.

- [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e4e79a6e7c4a44dc59bb849aa4c10e25efc820578d9c762664a59173525f0ce2): complete subsection reference.

- [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a9ced0aff7267c68e5e93df248f1a1bdfbff92f8b9089dd5252a1da81a77378d): complete subsection reference.

<a id="canonical-bf98a6b0abd4c6b853173365d181e9b0f18e0c95fc26d3aaba75588cb535038d"></a>

<a id="canonical-6f27b8e62ce77bda0e6a0c1e1540fd5262309daa57b07c7cc3fbbdb8452dd3eb"></a>

## description property — Property reference / 1b34f62e9ecc / 5

Type: `"string"`. Computed.

Description of the EnhancedFirewallPolicy.

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

<a id="canonical-862243fd6abf1ec78a979aadda0b04c26cccbd15f3609cb12f18b562fef2fc2c"></a>

<a id="canonical-25d4b1bdb3786d117450f3213d23251417602ac43bb6ecb3ffa50eb89c204a6c"></a>

## id property — Property reference / 1b34f62e9ecc / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b3f850ae91f2299c889a228806e6da6f647df0462615f1b2365defbc7dd02c6f"></a>

<a id="canonical-c7de049055582e7e1f3c31aedcc4667a0d0812305713654a5c0df6e0d9e48610"></a>

## labels property — Property reference / 1b34f62e9ecc / 7

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

<a id="canonical-3117c7369110f36a43af821df39cd2a0199db500bae56643a4a431376816c6e7"></a>

<a id="canonical-0c7bebbf329064ffa4a22fe96feeb154cf513d78f7e8b6baa64414d54b95e47a"></a>

## name property — Property reference / 1b34f62e9ecc / 8

Type: `"string"`. Required.

Name of the EnhancedFirewallPolicy.

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

<a id="canonical-70269c82db2fd30aab36b3cf3d7a979550466923c9bcbedb11328e18027e5e49"></a>

<a id="canonical-34947ebf04c513b6ca67b6655a365d798de43ab17da00e7be2f63dbc443cb926"></a>

## namespace property — Property reference / 1b34f62e9ecc / 9

Type: `"string"`. Required.

Namespace where the EnhancedFirewallPolicy exists.

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

- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46): complete subsection reference.

<a id="canonical-67ffffbb470ced99d3f7d6db8afd0aead6fd51e25c92b626755938ae6c24e113"></a>

## All schema paths — Property reference / 1b34f62e9ecc / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e3f27025e47507b95d7638f812bdc48c17e9038686c703fb4b941f46627753da) |
| `allowed_destinations` | [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-d4974ab2803adaaee3f70f400e8c1663e83f679eb306ec52a86a8cd9401db4c3) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eef3a2fe13fd3512e1e3380671cbe3297a246518e9bd13c3e91e89b12c616497) |
| `allowed_sources` | [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-f8c955ffb177cb38f81e93ebf2b3fb38fdd5255d559df694060677b7d46799fe) |
| `allowed_sources.prefix` | [allowed_sources.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-41ea8e69c5229124200c0bec80c1c37491d313e7f79ecb65f5bfd5b2ef52a649) |
| `annotations` | [annotations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2ecc40d36691b8884cdcbfe37f6b11391d2b607f58f73939c923c2d8b1a5deef) |
| `denied_destinations` | [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-df137fc16e760f2ecd221e51840999f91b16b2eebc5dde8b2725053880d2dcdd) |
| `denied_destinations.prefix` | [denied_destinations.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-466fdf2d526c2f47888fb01764e057b24256be1e5d9ccec9722ccbd92d864645) |
| `denied_sources` | [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-bca5daa4f7d481bcf0c6d1181976fe2eeaaeba7b8f885e9632e6c3e1857c4a27) |
| `denied_sources.prefix` | [denied_sources.prefix](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-626053e92949e1d08fa9689c84e71115fbd59539dc4225057f846432de9f9498) |
| `deny_all` | [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-198f290455686e4aebb3678a0119af1fee9fb1c6b6b0771a278a5f6d560a16bc) |
| `description` | [description](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-bf98a6b0abd4c6b853173365d181e9b0f18e0c95fc26d3aaba75588cb535038d) |
| `id` | [id](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-862243fd6abf1ec78a979aadda0b04c26cccbd15f3609cb12f18b562fef2fc2c) |
| `labels` | [labels](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-b3f850ae91f2299c889a228806e6da6f647df0462615f1b2365defbc7dd02c6f) |
| `name` | [name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3117c7369110f36a43af821df39cd2a0199db500bae56643a4a431376816c6e7) |
| `namespace` | [namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-70269c82db2fd30aab36b3cf3d7a979550466923c9bcbedb11328e18027e5e49) |
| `rule_list` | [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7300eb53fd1c579ac7b2c592fc830f8de1c70aabc13feb5ea253ae3db577a91d) |
| `rule_list.rules` | [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-419d4ed01e79d80bfbdbad0611fa25adf9773244d45095f052e44376d4801efe) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-816fbeb70224c7ebc85ae41655eacd8f2ae08b320f43072ebddd0d74eaa47169) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-67247361f4e0f44f0d836b70467a160b9f0cf6a0a48811968fc9ce15faacf86f) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-02ded69eb97d2a3ea1aa76065c739fbda9e4ac57434de1fddcd99542900030c0) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-fd17e98498c2d778d159bd34c1b5ff2da96cc971a9759d816354fa4d5dbea150) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-82c1e057683464369fa4472713383288c595bf78bfe9e1464a74a3eb9de841bf) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-d952d45957986b5b281237f3d6693df711f749a5b8ff3fbec4b4a9959c840d56) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eca74b707d4fbcd0cad44c7ff42d9c8fdeb91e975813ab43ffcbe775451c3264) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5275bff585a6157b390a72e2bea94429c2a58152ff9d68911227ca667e75ec99) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-77acd6484f8122672db5479cd1aafd8a6abeee15221439adbfa87a70ee30b42d) |
| `rule_list.rules.allow` | [rule_list.rules.allow](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-88cf1abe8c7aaa44dd82ba4eb6d06c074f18b72b856c012f30b04fa1642b7ce1) |
| `rule_list.rules.applications` | [rule_list.rules.applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-8973384423f63f8e20520389488194037e9655a5c0910cc141b6108f54677637) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5cb3faa8f67210b744f93f23b7d18bd889d524661ccf2836c1466bf08c0f47cb) |
| `rule_list.rules.deny` | [rule_list.rules.deny](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a7bd8bad3dee6aa10c5b893a62da3b0fe291345c123234870d3a1b6861c5e563) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-c8d81e05a189291eada769e6b9488a42d0d4b0bfceea3353a0de72fc5612e6aa) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e10c91b1b7b5c92b656f90f9a180825903644f47a166899ec06d66fcaa9706d1) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-329146244e1d0eb527d1f9b58b1535932dd0fe82b3606585b69a516f09d12d5e) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1d009ed2be3469e1614409147763ac399c686d40098df81a7245ba1b9a71373f) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-954d051366391d85fae50755171facebf07bb29248ff48bbef2969fe1d07564c) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-103c8060ebc0dd7b757f8139c2783eb3546325db6cc2ecb1227b8bc9753e7124) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-caf8b24b3c24ebf838704b8247638ee6f39ab0d55b4eb8a991344a751c96fc26) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-29dc29e6565061b0e02b99d44bc1ae2dfa1de81a3086accdfd90d4d20e9d552f) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-35bbbfabeaa767cd94ea5b2a359d252ba4ee99d7437fc2fb540451fb73df5eb9) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-32e833117a94263cfc57869cfa26ab5ce342fd2dce97fbfc2b2da25e47c38de6) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-8a65194fb6c9a1e99dc2f719757bc60fed46c0d91336f435a6d35c67179e5581) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-8820a44f40c3f68e9be7e87234cc2249fab9ccca9c82cbde36d12d8f74f68ab0) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-deed3bd29f14ef6b2dccb16e2c551523f2668ef57f42e9be067959f79a60a1c4) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-b1a99bd6b803fdde7973fc0b450403436684635ee6bcbb396d5309a59b9ee7a8) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5c7a4ecaa24bb41454752780432b608583b557423088eb225239bb34f57956c0) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e1a87933b7f066a5a0ba4f204139b832ac33cc8cd51836122de7e96e4e194766) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7a51dd1ec04a1abd6715ecd803802ac76878c45ce9115ff199d818758be7b81c) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3eaf193ddbfe86708a9e75bc458f0c960634971d5de9d55a5201abc5017cfcba) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3a52088cd627f5d094b1cb4e125f389a039e10460affa7967b34ca5061b9878b) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0f02621664a169d1e3a5cdf10900c7cd4c11aedeff13f6862dc3471d73ae5749) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-6a6f46db43e88b30faf8a8be51f7cb012945e9652199d244dadedebba2a359eb) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a7adde13fb13242ac1315897c0cc84c7ab1326303724325e05bc48519036afbf) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-15ba39fa6ca98399e8bef0d0879e8c512a7d19c0b687e14f1998b038e985a1f3) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-bb29b528fedec842355d9416a91a858ae71082065f5c6300bc073c8a39cbb355) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-4106ce897a1efaf7d321e4bbfab82ebeccc24c36c5c5fce2ef46807062493c07) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7ca882e1512af0d9460000014edf44542f488a9702a3e0aaeaef5eef2ac4674a) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-47f659fc0202416bf55a4c4a7c17a4fcabf32e9c3aed9dd53cab4e319e1b1fdb) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1be73ada0e10f669d0f752ba0c6e7d7c45b0fc1fc1916637d033be376d79fbd7) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-b71ff82dd17fe0404ef7b4d395827a390b7164b5bd563ed3b912088c73c36eab) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5d70d8745bc115b243b8a36a2ae94f3d966bc1add387cd6381d3aa0b60801cf5) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a6640dd4b2d1614e17ccd2092095446e877176e713fd539be94820aad518e7b7) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-6731a893bd681257846ea8b0fb711bc33981c8d4fce79a052ce24e892d0f3f91) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-f5064eae810b290c5af85aa9830d12dc356444c8488540357d25fec686c5e070) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-74db55bc30473778c3564668f086949304ea30258016cb32d7e9dd94db89d04a) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e9544a0042491d359d2d954db060f5919e250721c6edb51b69c8882b032ea717) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a182ccc0988b735aedef54291767785a44ca3ababffa862ee17d8fcb7082c7b2) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-9bd7a6149e1606ba770b929b190cb2d253fb3abcc07c9b4063fc8c83f31b587d) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e3b1e726923f07e20e252ed8460f767a93fba183a4a2608b0000b64922e356d8) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-706f3021ecc042210fca7740b3d442ba1ae818aed4e44286b9f1bf2679b58daf) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-ef4c95c0f08be56cf7925a0be9442dc09e51f130cfba8b35921bf7d352de8aa3) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a0551567620e54dd4371d62bbdb5423e982808087d4c6d0a7e39991751d94270) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-b0b0dbe39f0a8d1218386d2132f8b902d6df5bbe56d568a753d3569922298a7d) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-888547bdb3fc7048469c45e1d75cbb0ea0c26fe4c040ad145d8def11ca5bf3d5) |

<a id="canonical-db92fb887c09ec124f0b9b26edd8790fdad0d51618b74eade59b067dd35b3b94"></a>

## Next pages — Property reference / 1b34f62e9ecc / 11

- [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-cfcfc1c5f374870b18c521d8bddf34f8136c3c8a886432c73d1b7f24ab6aae80)
- [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-59aeaa8f794d56e5bacdc5a578af0339fcceeb74cdfc3828817bf02fd8be9d61)
- [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3461c3a8e44ade46badb25f81bc7040807e6c26f425354aae577fd33c070f359)
- [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-39413b52c833c9e4a303b5835c0624caaf50c3752a865bbe85dd0387ca295ebb)
- [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e4e79a6e7c4a44dc59bb849aa4c10e25efc820578d9c762664a59173525f0ce2)
- [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a9ced0aff7267c68e5e93df248f1a1bdfbff92f8b9089dd5252a1da81a77378d)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-cfcfc1c5f374870b18c521d8bddf34f8136c3c8a886432c73d1b7f24ab6aae80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b068c862bbbd781a1ba55782c1997e5ffd7e5b913a8eaacf7b3e628567408149"></a>

## allow_all — allow_all / 9951ecacd607 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- allow_all

<a id="canonical-e3f27025e47507b95d7638f812bdc48c17e9038686c703fb4b941f46627753da"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

OneOf alternatives in this subsection:

- [allow_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e3f27025e47507b95d7638f812bdc48c17e9038686c703fb4b941f46627753da)
- [allowed_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-d4974ab2803adaaee3f70f400e8c1663e83f679eb306ec52a86a8cd9401db4c3)
- [allowed_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-f8c955ffb177cb38f81e93ebf2b3fb38fdd5255d559df694060677b7d46799fe)
- [denied_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-df137fc16e760f2ecd221e51840999f91b16b2eebc5dde8b2725053880d2dcdd)
- [denied_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-bca5daa4f7d481bcf0c6d1181976fe2eeaaeba7b8f885e9632e6c3e1857c4a27)
- [deny_all](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-198f290455686e4aebb3678a0119af1fee9fb1c6b6b0771a278a5f6d560a16bc)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7300eb53fd1c579ac7b2c592fc830f8de1c70aabc13feb5ea253ae3db577a91d)

Select alternatives according to the provider validators above.

<a id="canonical-f476d008e26b7d099729aa0d39ba2821cee90647e38fdc81096f9294630ea2ae"></a>

## Direct properties — allow_all / 9951ecacd607 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aaf6eeda3b4ce41820e84e7e3d06212fd4704ab53bf24d1abb23fe792f6e9b94"></a>

## Next pages — allow_all / 9951ecacd607 / 4

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-59aeaa8f794d56e5bacdc5a578af0339fcceeb74cdfc3828817bf02fd8be9d61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-517d82b3d6441ed075d07b91f2e9d8d1d084fb2a548cdcd5cfff76a9aaea6e47"></a>

## allowed_destinations — allowed_destinations / a8d70f2125a6 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- allowed_destinations

<a id="canonical-d4974ab2803adaaee3f70f400e8c1663e83f679eb306ec52a86a8cd9401db4c3"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-34926eb8bd96e453fdc5b90fd0130fef31432ce730c430a3a96e386d3c3c7ca4"></a>

## Direct properties — allowed_destinations / a8d70f2125a6 / 3

<a id="canonical-eef3a2fe13fd3512e1e3380671cbe3297a246518e9bd13c3e91e89b12c616497"></a>

<a id="canonical-80044d1211add7eae6a2bba63d2229674ffa388f1e18a8468ebc853b2b46034c"></a>

## prefix property — allowed_destinations / a8d70f2125a6 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3e5703b4b77238e9bba633119faa093fcba251a1dd1cf35d44f2d7c3f9767d00"></a>

## Next pages — allowed_destinations / a8d70f2125a6 / 5

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-3461c3a8e44ade46badb25f81bc7040807e6c26f425354aae577fd33c070f359"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d75fd9301103684014dc978bfec9ad7f474034fccb01bc7fceff55a704cd80c8"></a>

## allowed_sources — allowed_sources / eca85346f61b / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- allowed_sources

<a id="canonical-f8c955ffb177cb38f81e93ebf2b3fb38fdd5255d559df694060677b7d46799fe"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-d6c8ffc0ade9d95440d3de17b64f647edebd3d895696746d2e604b98d5fa86ec"></a>

## Direct properties — allowed_sources / eca85346f61b / 3

<a id="canonical-41ea8e69c5229124200c0bec80c1c37491d313e7f79ecb65f5bfd5b2ef52a649"></a>

<a id="canonical-ad8064f2042df5abae651cbd6f06e7e33861f8b5c80b4ee4574f87df7542af7d"></a>

## prefix property — allowed_sources / eca85346f61b / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-1ec227f53208c6054e59e66e9f1dafe335d9450c499124d2cc37ee33aded509f"></a>

## Next pages — allowed_sources / eca85346f61b / 5

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-39413b52c833c9e4a303b5835c0624caaf50c3752a865bbe85dd0387ca295ebb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-648966a9a5dfee68541a9b84ec400484d468b51bee913d413d32bb9e6e8906ad"></a>

## denied_destinations — denied_destinations / c74f46f7003c / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- denied_destinations

<a id="canonical-df137fc16e760f2ecd221e51840999f91b16b2eebc5dde8b2725053880d2dcdd"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-8242f6b2c7dda67a63fa890f0e54b8975ac5dcbf9a6b504fe515e267c38e7ed2"></a>

## Direct properties — denied_destinations / c74f46f7003c / 3

<a id="canonical-466fdf2d526c2f47888fb01764e057b24256be1e5d9ccec9722ccbd92d864645"></a>

<a id="canonical-f47080d0cd5a94d3874581390d6080a6b47d35a204a8011b584a4cac57ab2818"></a>

## prefix property — denied_destinations / c74f46f7003c / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-89fd9fa5022765e2538014ba838a2d615b1c3912a4aba5861c8815dd45d49e45"></a>

## Next pages — denied_destinations / c74f46f7003c / 5

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-e4e79a6e7c4a44dc59bb849aa4c10e25efc820578d9c762664a59173525f0ce2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dfc63cfdd18c28bc1c3b2f0bf1f7cf10f4573219ea401822e6c16b0020f510d"></a>

## denied_sources — denied_sources / 0864a07a88c3 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- denied_sources

<a id="canonical-bca5daa4f7d481bcf0c6d1181976fe2eeaaeba7b8f885e9632e6c3e1857c4a27"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-6afe597cdd3f14d270a605a537aa2e4f265277ce6a6fd51c5c9e8c97bdb7135e"></a>

## Direct properties — denied_sources / 0864a07a88c3 / 3

<a id="canonical-626053e92949e1d08fa9689c84e71115fbd59539dc4225057f846432de9f9498"></a>

<a id="canonical-75cb38187a7ea16c945983563ee14ee3f7e5d7cfdf6e6bc280f583c5de5b0d4a"></a>

## prefix property — denied_sources / 0864a07a88c3 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-838e12fbb6c78b39c6577e0950580b150ff4d59cf3fc2b5a1b03f3b70fc92118"></a>

## Next pages — denied_sources / 0864a07a88c3 / 5

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-a9ced0aff7267c68e5e93df248f1a1bdfbff92f8b9089dd5252a1da81a77378d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15cc2300f0bf7e368b2cee13bb56620f0869ede7ef3b746d2127702f9d58438e"></a>

## deny_all — deny_all / 17dd2b16042a / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- deny_all

<a id="canonical-198f290455686e4aebb3678a0119af1fee9fb1c6b6b0771a278a5f6d560a16bc"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fc499ed95b082052bc5a01b63f4c2fce916fe1621223cfb0dbcdbb6b78495391"></a>

## Direct properties — deny_all / 17dd2b16042a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b3bfe9341dfc805f96360fe936d107ca0ed6bdd86a71ed339cc708b36bb03cf"></a>

## Next pages — deny_all / 17dd2b16042a / 4

- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d680f44ccffec9ef97530b338bb9ef2100dd4b7cd1ad7e63164aa1be75ecaca"></a>

## rule_list — rule_list / b9f54bc3c265 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- rule_list

<a id="canonical-7300eb53fd1c579ac7b2c592fc830f8de1c70aabc13feb5ea253ae3db577a91d"></a>

Type: `"single"`. Computed.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

Upstream description:

Custom Enhanced Firewall Policy Rules.

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

<a id="canonical-7096ef07cda50d804fa3759febc0909e1ab49f1d39e11946ade4283b8b76f20a"></a>

## Direct properties — rule_list / b9f54bc3c265 / 3

- [rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd): complete subsection reference.

<a id="canonical-e830dffe383d2f1db77cbd56022665f0cf182907e56e82732f3ad63359ce3e45"></a>

## Next pages — rule_list / b9f54bc3c265 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68ec2af7e62cfb7477ec64b4f2ec7854c68c64aec85ca31ea2ec7099b9d1ca6c"></a>

## rule_list.rules — rule_list.rules / eeb2d03ae687 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- rule_list.rules

<a id="canonical-419d4ed01e79d80bfbdbad0611fa25adf9773244d45095f052e44376d4801efe"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policy Rules.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-7ada09e73705a1ee039392f84e3aa7f3e33dcfc894c6e786dba869dde36b7337"></a>

## Direct properties — rule_list.rules / eeb2d03ae687 / 3

- [advanced_action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e4f3684a769568f87ca5e447e85fc3b91e1340faaf5cff72fb881a76028339dd): complete subsection reference.

- [all_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1cedad637e82da60994edb74a77c357d45e7b2e5354f03fe741a375cc43cdf6d): complete subsection reference.

- [all_sli_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-ecb34d48f9c61cd09d88e1dca3f019c48c92a64dc757e9982a871cd7006be8f3): complete subsection reference.

- [all_slo_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5905025e98ec147d12801f927c5d4c0fc4ce0bec59f527ea7c180343c63d502e): complete subsection reference.

- [all_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3bb13f634f15c334b68f804e553d4057e6d0b71815d60e8a5facea9d9049341a): complete subsection reference.

- [all_tcp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-051a7111078e57336bdb0ca5557d1172f6b9716c8824e020a37f2a3c9b758a5e): complete subsection reference.

- [all_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0c36c26bd5938e7c2f1970ffad94293a1caf6b7e8058dcbebc92cba072d87d8b): complete subsection reference.

- [all_udp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-fe3fecc8404f0673ee1761faa370edca61451e782d5a6139492f240d2e14a82d): complete subsection reference.

- [allow](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-73c74143882d04923b307ad799a116c8756fd12faf5a6b1de4539dc195236ff3): complete subsection reference.

- [applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-96373786a8a4aa4947895e5393c8a847666134f248cf14b4f6bd8702ed70d694): complete subsection reference.

- [deny](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a0a21035f996736d75484b1607c7ddf0bcfae98d58b1d9b44ff91ba34285f37e): complete subsection reference.

- [destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-573a2a0324c83c65c20884f483d4719853821b1c3533e2df2c6a6f963e7f86a6): complete subsection reference.

- [destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-702cdec0f14845c2c40ef42ad90aa716da9839c59ff2c18de36468bb5b60e389): complete subsection reference.

- [destination_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-c36f674bb68f537cbe15cf7ac4e134aa99f07115b686b9b6cfe1aecedbbef1c4): complete subsection reference.

- [destination_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2213f0ef6031f6e976ebf63efcdaa06c5476897e06947435ea365800a416607c): complete subsection reference.

- [insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-741d201d3a2b3a9b759f7bcbc72b6220435bf99158635473deaa5829c684e55b): complete subsection reference.

- [inside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7f4cb07b9d4fcf864aad3e55e56d44042cf3327c81b9fe617ae71ca44bde8da5): complete subsection reference.

- [inside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-725987edd37546fcfa7546e0c604c9c126bcce9e71d5b67e75f2ed6e2f32130c): complete subsection reference.

- [label_matcher](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-52df26bed3cf3d4c76e163ff7a90c450bcd39128201c972d5ab73c876e8e0b80): complete subsection reference.

- [metadata](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3926dfc10ab2393042b86f02ca54db32d33b74bbbb70ee9d01dd507e435a74f3): complete subsection reference.

- [outside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2f3c4cc7e70ef7458b34d41965a84b259837dd5fbf18ce893421c4c61d546f91): complete subsection reference.

- [outside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3d0568a82a224681a0356aab6e1ebf7e31076a3385c8183a369e07519ae12d54): complete subsection reference.

- [protocol_port_range](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-4490cd4db2103499cc04b0e433858740d8007fc86ac0cd6b167290c9369c4123): complete subsection reference.

- [source_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-73f1d7ea85db0752d3ff02c9d7cb127d8a5972c272618e9036c6a365fdca8742): complete subsection reference.

- [source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-202eddb57b6eec894059373000f22e171a2668303e5f41183bea0d7c30ce4aff): complete subsection reference.

- [source_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-052ed70ca5fab404fbdc1c7883895d85223e107b0324b6ae3179b5cbfbde16b0): complete subsection reference.

- [source_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2e1692091408120c306e23d69e92ddcfb142cab3b29f7498f8f2fcab07380bf4): complete subsection reference.

<a id="canonical-23696fa84dabf7d8e5a2e076eb44bb32e81e54158e8d2134ea845240c28657ba"></a>

## Next pages — rule_list.rules / eeb2d03ae687 / 4

- [rule_list.rules.advanced_action](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-e4f3684a769568f87ca5e447e85fc3b91e1340faaf5cff72fb881a76028339dd)
- [rule_list.rules.all_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1cedad637e82da60994edb74a77c357d45e7b2e5354f03fe741a375cc43cdf6d)
- [rule_list.rules.all_sli_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-ecb34d48f9c61cd09d88e1dca3f019c48c92a64dc757e9982a871cd7006be8f3)
- [rule_list.rules.all_slo_vips](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-5905025e98ec147d12801f927c5d4c0fc4ce0bec59f527ea7c180343c63d502e)
- [rule_list.rules.all_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3bb13f634f15c334b68f804e553d4057e6d0b71815d60e8a5facea9d9049341a)
- [rule_list.rules.all_tcp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-051a7111078e57336bdb0ca5557d1172f6b9716c8824e020a37f2a3c9b758a5e)
- [rule_list.rules.all_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-0c36c26bd5938e7c2f1970ffad94293a1caf6b7e8058dcbebc92cba072d87d8b)
- [rule_list.rules.all_udp_traffic](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-fe3fecc8404f0673ee1761faa370edca61451e782d5a6139492f240d2e14a82d)
- [rule_list.rules.allow](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-73c74143882d04923b307ad799a116c8756fd12faf5a6b1de4539dc195236ff3)
- [rule_list.rules.applications](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-96373786a8a4aa4947895e5393c8a847666134f248cf14b4f6bd8702ed70d694)
- [rule_list.rules.deny](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-a0a21035f996736d75484b1607c7ddf0bcfae98d58b1d9b44ff91ba34285f37e)
- [rule_list.rules.destination_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-573a2a0324c83c65c20884f483d4719853821b1c3533e2df2c6a6f963e7f86a6)
- [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-702cdec0f14845c2c40ef42ad90aa716da9839c59ff2c18de36468bb5b60e389)
- [rule_list.rules.destination_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-c36f674bb68f537cbe15cf7ac4e134aa99f07115b686b9b6cfe1aecedbbef1c4)
- [rule_list.rules.destination_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2213f0ef6031f6e976ebf63efcdaa06c5476897e06947435ea365800a416607c)
- [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-741d201d3a2b3a9b759f7bcbc72b6220435bf99158635473deaa5829c684e55b)
- [rule_list.rules.inside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-7f4cb07b9d4fcf864aad3e55e56d44042cf3327c81b9fe617ae71ca44bde8da5)
- [rule_list.rules.inside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-725987edd37546fcfa7546e0c604c9c126bcce9e71d5b67e75f2ed6e2f32130c)
- [rule_list.rules.label_matcher](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-52df26bed3cf3d4c76e163ff7a90c450bcd39128201c972d5ab73c876e8e0b80)
- [rule_list.rules.metadata](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3926dfc10ab2393042b86f02ca54db32d33b74bbbb70ee9d01dd507e435a74f3)
- [rule_list.rules.outside_destinations](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2f3c4cc7e70ef7458b34d41965a84b259837dd5fbf18ce893421c4c61d546f91)
- [rule_list.rules.outside_sources](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-3d0568a82a224681a0356aab6e1ebf7e31076a3385c8183a369e07519ae12d54)
- [rule_list.rules.protocol_port_range](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-4490cd4db2103499cc04b0e433858740d8007fc86ac0cd6b167290c9369c4123)
- [rule_list.rules.source_aws_vpc_ids](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-73f1d7ea85db0752d3ff02c9d7cb127d8a5972c272618e9036c6a365fdca8742)
- [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-202eddb57b6eec894059373000f22e171a2668303e5f41183bea0d7c30ce4aff)
- [rule_list.rules.source_label_selector](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-052ed70ca5fab404fbdc1c7883895d85223e107b0324b6ae3179b5cbfbde16b0)
- [rule_list.rules.source_prefix_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-2e1692091408120c306e23d69e92ddcfb142cab3b29f7498f8f2fcab07380bf4)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-e4f3684a769568f87ca5e447e85fc3b91e1340faaf5cff72fb881a76028339dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7809d7655e1528dcfe1fc609ba1359407d9625b7f4b753af6bcbef4d9bc7203"></a>

## rule_list.rules.advanced_action — rule_list.rules.advanced_action / fc3f98626c0b / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.advanced_action

<a id="canonical-816fbeb70224c7ebc85ae41655eacd8f2ae08b320f43072ebddd0d74eaa47169"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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

<a id="canonical-fbbd43242e4ef27ddd9af16e518b64ca303f5d08eb34a21c928049ef3115c82a"></a>

## Direct properties — rule_list.rules.advanced_action / fc3f98626c0b / 3

<a id="canonical-67247361f4e0f44f0d836b70467a160b9f0cf6a0a48811968fc9ce15faacf86f"></a>

<a id="canonical-bafabb78f0ab843ed09f63a0848e98b22b84ae0c4175bf8ce89c74204c59445c"></a>

## action property — rule_list.rules.advanced_action / fc3f98626c0b / 4

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-87d9318ebc461693cbc73e0268207a6b5a833955715ffe818bfbe8f60dd952bc"></a>

## Next pages — rule_list.rules.advanced_action / fc3f98626c0b / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-1cedad637e82da60994edb74a77c357d45e7b2e5354f03fe741a375cc43cdf6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf252ef2ca1335cdf2ed6be748cb8da450b34ad3da51c84beb7e874b3bbce4f2"></a>

## rule_list.rules.all_destinations — rule_list.rules.all_destinations / 33bbfd3dc978 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_destinations

<a id="canonical-02ded69eb97d2a3ea1aa76065c739fbda9e4ac57434de1fddcd99542900030c0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-007b27ea368592ecf20cb8a757832870a8f31de9d6fd65e5fa8dfde606a1dc33"></a>

## Direct properties — rule_list.rules.all_destinations / 33bbfd3dc978 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-404b7a06f797e2fcedfcd17b5f7e498c07e2e7a69fd45aff63d8dc7ca4ef485e"></a>

## Next pages — rule_list.rules.all_destinations / 33bbfd3dc978 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-ecb34d48f9c61cd09d88e1dca3f019c48c92a64dc757e9982a871cd7006be8f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94e7196b9b550afd5469dcfeaf40db010dc0b95212721369cb710df9c2d1924e"></a>

## rule_list.rules.all_sli_vips — rule_list.rules.all_sli_vips / c4be586646fc / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_sli_vips

<a id="canonical-fd17e98498c2d778d159bd34c1b5ff2da96cc971a9759d816354fa4d5dbea150"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-80d57be26d190d182febf6ef22bf9cb4234c219acd64ccbd431464c4a974c8ec"></a>

## Direct properties — rule_list.rules.all_sli_vips / c4be586646fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5626d3689659a1ad2c63be96a24e54949761b6be524efb29478a17c09f58e396"></a>

## Next pages — rule_list.rules.all_sli_vips / c4be586646fc / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-5905025e98ec147d12801f927c5d4c0fc4ce0bec59f527ea7c180343c63d502e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58ddaf9b4fd1f1fa2c2e4c0de098c8569af53049d8119402fbb4f5825447a5b5"></a>

## rule_list.rules.all_slo_vips — rule_list.rules.all_slo_vips / bbceface4293 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_slo_vips

<a id="canonical-82c1e057683464369fa4472713383288c595bf78bfe9e1464a74a3eb9de841bf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1e387fa284463ee5f6b56b3a94b144f20092c8019b9b6acbb0ffe77cfb26fecf"></a>

## Direct properties — rule_list.rules.all_slo_vips / bbceface4293 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-039adbe1b91f6891a8d456da539400439eaac3d9d7471897d30c19794c791c45"></a>

## Next pages — rule_list.rules.all_slo_vips / bbceface4293 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-3bb13f634f15c334b68f804e553d4057e6d0b71815d60e8a5facea9d9049341a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-681d71a058926d6693dc5fd5811123c1fd71b3c0779eda8abdb91b51f7955298"></a>

## rule_list.rules.all_sources — rule_list.rules.all_sources / 208001700083 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_sources

<a id="canonical-d952d45957986b5b281237f3d6693df711f749a5b8ff3fbec4b4a9959c840d56"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-bf3764fcce1eb1a9145efbfe8bd51e0ab6d1b2eb841ca53aec1394d4dfed5339"></a>

## Direct properties — rule_list.rules.all_sources / 208001700083 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1dbf502c38d0153988e2277fd32ea886a93130c1fa9966ecf6ed3323c3ed1fa"></a>

## Next pages — rule_list.rules.all_sources / 208001700083 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-051a7111078e57336bdb0ca5557d1172f6b9716c8824e020a37f2a3c9b758a5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0db29dc4836fb06ee8009063e72ffba570febe4192cd093b73f58d39738b5a0"></a>

## rule_list.rules.all_tcp_traffic — rule_list.rules.all_tcp_traffic / e2f912922bdd / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_tcp_traffic

<a id="canonical-eca74b707d4fbcd0cad44c7ff42d9c8fdeb91e975813ab43ffcbe775451c3264"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-805b58f35aef8564fd4d57e79e9f8bb94f19c924998d84dc73249e444c1646b5"></a>

## Direct properties — rule_list.rules.all_tcp_traffic / e2f912922bdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-080383ae83945c6a49169039da4a75e374b6feaeee0dbe5b42806371fc57b8e2"></a>

## Next pages — rule_list.rules.all_tcp_traffic / e2f912922bdd / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-0c36c26bd5938e7c2f1970ffad94293a1caf6b7e8058dcbebc92cba072d87d8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f8e05c32159ffc0c841e1901d4d72b28174b644c973e8e57a25400881531ab"></a>

## rule_list.rules.all_traffic — rule_list.rules.all_traffic / f5c230fbdb48 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_traffic

<a id="canonical-5275bff585a6157b390a72e2bea94429c2a58152ff9d68911227ca667e75ec99"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-771d4e244f5e9cdbabb8532188f8bea48ca83eb1fc263a03d2b27449812f1e74"></a>

## Direct properties — rule_list.rules.all_traffic / f5c230fbdb48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-548d798476e52a2a949ba90fe7a1d7dd16294cda8fbb14701d345391e97edf72"></a>

## Next pages — rule_list.rules.all_traffic / f5c230fbdb48 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-fe3fecc8404f0673ee1761faa370edca61451e782d5a6139492f240d2e14a82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dbf007a2a4f00abbf76a4e046c25e4a6a9302ef54ca16bc3970ca8e3af8daac"></a>

## rule_list.rules.all_udp_traffic — rule_list.rules.all_udp_traffic / 56a3f993f0a7 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.all_udp_traffic

<a id="canonical-77acd6484f8122672db5479cd1aafd8a6abeee15221439adbfa87a70ee30b42d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-2a7caae4389b9078460834815c0a453c3ba0f7f665d849143c6f58d4a3795432"></a>

## Direct properties — rule_list.rules.all_udp_traffic / 56a3f993f0a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2588c38193a6cf556d286e1d259395c2d6d9257344898f008305138a77b4253d"></a>

## Next pages — rule_list.rules.all_udp_traffic / 56a3f993f0a7 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-73c74143882d04923b307ad799a116c8756fd12faf5a6b1de4539dc195236ff3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-000ff748f5dfe4d51936963977e7163091e071ddd3270e9abe7e6d0b834d1bc9"></a>

## rule_list.rules.allow — rule_list.rules.allow / 7b7d8444c737 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.allow

<a id="canonical-88cf1abe8c7aaa44dd82ba4eb6d06c074f18b72b856c012f30b04fa1642b7ce1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-a80f92fa681fa47b74f9be16f8e23ff9fa6d9b323a421c0667177990e8f6ed91"></a>

## Direct properties — rule_list.rules.allow / 7b7d8444c737 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c1fcf3cf4f2bf46c661b3c838a9b11bb75bdcbcc54e4ac1256e500aeb295ee1"></a>

## Next pages — rule_list.rules.allow / 7b7d8444c737 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-96373786a8a4aa4947895e5393c8a847666134f248cf14b4f6bd8702ed70d694"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1144b7cc824601d336ab1ba243a28ad38288b588fe454f3163295f6aa72d0bc1"></a>

## rule_list.rules.applications — rule_list.rules.applications / be741b9623b0 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.applications

<a id="canonical-8973384423f63f8e20520389488194037e9655a5c0910cc141b6108f54677637"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-5f1751cfb08d790129ed8eae9799263585ef6ffbd94202d1dfca3a24499d94b1"></a>

## Direct properties — rule_list.rules.applications / be741b9623b0 / 3

<a id="canonical-5cb3faa8f67210b744f93f23b7d18bd889d524661ccf2836c1466bf08c0f47cb"></a>

<a id="canonical-2f4bad31e8c3d54aa922d370539ca723e0beae53b95e27bbb6c847591d2e92c5"></a>

## applications property — rule_list.rules.applications / be741b9623b0 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-279a5ac501a79034917b7022cde4562cf7ff7665a24564b9acbaa8ae4d16ee1e"></a>

## Next pages — rule_list.rules.applications / be741b9623b0 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-a0a21035f996736d75484b1607c7ddf0bcfae98d58b1d9b44ff91ba34285f37e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72c7ea50373f1978b115574640d0d47331b6f73530514099801872dd8d717373"></a>

## rule_list.rules.deny — rule_list.rules.deny / 5a21ca71ab57 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.deny

<a id="canonical-a7bd8bad3dee6aa10c5b893a62da3b0fe291345c123234870d3a1b6861c5e563"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-81cd4ea06bc42e3e4f9b9921979c28c05f690fb3d02da63cca4999ef0d7f772a"></a>

## Direct properties — rule_list.rules.deny / 5a21ca71ab57 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5491e61eae7f3189a9b1fbe63dc6513d540ce8d30f1f7af2b67cb4f2959c3698"></a>

## Next pages — rule_list.rules.deny / 5a21ca71ab57 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-573a2a0324c83c65c20884f483d4719853821b1c3533e2df2c6a6f963e7f86a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b12dd4879069c4b9f3e7f47adf510c74d03058d3497fdc458c8f9846d2b35838"></a>

## rule_list.rules.destination_aws_vpc_ids — rule_list.rules.destination_aws_vpc_ids / 02f2592d71c1 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.destination_aws_vpc_ids

<a id="canonical-c8d81e05a189291eada769e6b9488a42d0d4b0bfceea3353a0de72fc5612e6aa"></a>

Type: `"single"`. Computed.

Configuration parameter for destination aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

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

<a id="canonical-b974f705c0ec0237e62a07b06c27164471184c36c019c5a831c97f532e537fb1"></a>

## Direct properties — rule_list.rules.destination_aws_vpc_ids / 02f2592d71c1 / 3

<a id="canonical-e10c91b1b7b5c92b656f90f9a180825903644f47a166899ec06d66fcaa9706d1"></a>

<a id="canonical-9ffbd1227db2ac5ea2002fc4200d2a43fdbc05723112035af9b5ad99aaa6b1a4"></a>

## vpc_id property — rule_list.rules.destination_aws_vpc_ids / 02f2592d71c1 / 4

Type: `["list", "string"]`. Computed.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bc9a9f4e6bfa083ef2051efd1c07e6960c14cccff0e00a0a6aa36c4aeba7144c"></a>

## Next pages — rule_list.rules.destination_aws_vpc_ids / 02f2592d71c1 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-702cdec0f14845c2c40ef42ad90aa716da9839c59ff2c18de36468bb5b60e389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f918848a39ee2899b0344f8af937f68940217207df02ed5b735c3e8ea940ee6"></a>

## rule_list.rules.destination_ip_prefix_set — rule_list.rules.destination_ip_prefix_set / 3525c050361b / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.destination_ip_prefix_set

<a id="canonical-329146244e1d0eb527d1f9b58b1535932dd0fe82b3606585b69a516f09d12d5e"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-072ac3df28ffad7eb067b2d0223a95094a7a5e0f3b43ebcf1ca6b9feda45aa03"></a>

## Direct properties — rule_list.rules.destination_ip_prefix_set / 3525c050361b / 3

- [ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-530b15481fbb5c2c6bf62cb663feea8886f4d328e2406252ecbac463a57b0b1b): complete subsection reference.

<a id="canonical-de9cbe1421caef8bad3541662f1dd77329d6cff2eea6da9f54bdc8969be5c427"></a>

## Next pages — rule_list.rules.destination_ip_prefix_set / 3525c050361b / 4

- [rule_list.rules.destination_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-530b15481fbb5c2c6bf62cb663feea8886f4d328e2406252ecbac463a57b0b1b)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-530b15481fbb5c2c6bf62cb663feea8886f4d328e2406252ecbac463a57b0b1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32fcd45f96c0826c3128bb3e7fec0cd4191953dfa6128294c6799e18062679da"></a>

## rule_list.rules.destination_ip_prefix_set.ref — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-702cdec0f14845c2c40ef42ad90aa716da9839c59ff2c18de36468bb5b60e389)
- rule_list.rules.destination_ip_prefix_set.ref

<a id="canonical-1d009ed2be3469e1614409147763ac399c686d40098df81a7245ba1b9a71373f"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-fbc99a314dcffde152c5caf3f51e0ba88310e0b4a1634dfcbc358fd7f7613ee0"></a>

## Direct properties — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 3

<a id="canonical-954d051366391d85fae50755171facebf07bb29248ff48bbef2969fe1d07564c"></a>

<a id="canonical-374c6010689c593123d852e37a727f2ebc615d504781cb1b8d034b71ed20e93d"></a>

## kind property — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 4

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

<a id="canonical-103c8060ebc0dd7b757f8139c2783eb3546325db6cc2ecb1227b8bc9753e7124"></a>

<a id="canonical-f369dc6efe27784763ff0cab336cf71eb99dfb3d79c62233c2584c2be3580cef"></a>

## name property — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 5

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

<a id="canonical-caf8b24b3c24ebf838704b8247638ee6f39ab0d55b4eb8a991344a751c96fc26"></a>

<a id="canonical-4b5934c5ccc8d7bebe0f69a94ea5a5e9c499b845eab3d05e1fa46bfaac0425a3"></a>

## namespace property — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 6

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

<a id="canonical-29dc29e6565061b0e02b99d44bc1ae2dfa1de81a3086accdfd90d4d20e9d552f"></a>

<a id="canonical-4f70c607d66f2cd0c0c87e7384709c88f0ee10c638b1d1ce06e2d68243732a3c"></a>

## tenant property — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 7

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

<a id="canonical-35bbbfabeaa767cd94ea5b2a359d252ba4ee99d7437fc2fb540451fb73df5eb9"></a>

<a id="canonical-6272cb71bbbcd271703da382e5aae75f5c325f04b4df6ac7c69f7e3caaece88b"></a>

## uid property — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 8

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

<a id="canonical-8163b24529e2574a19b7dcd3f44cf8aa03c0689a6734c5df1cb90f79b4c0385b"></a>

## Next pages — rule_list.rules.destination_ip_prefix_set.ref / 4db5904d30d6 / 9

- [rule_list.rules.destination_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-702cdec0f14845c2c40ef42ad90aa716da9839c59ff2c18de36468bb5b60e389)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-c36f674bb68f537cbe15cf7ac4e134aa99f07115b686b9b6cfe1aecedbbef1c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a9129f4a73286d15c3ba8a0922986b4b55eb5f20c5be7c1bcc5e1c126fce44d"></a>

## rule_list.rules.destination_label_selector — rule_list.rules.destination_label_selector / 25d81f4294a5 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.destination_label_selector

<a id="canonical-32e833117a94263cfc57869cfa26ab5ce342fd2dce97fbfc2b2da25e47c38de6"></a>

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

<a id="canonical-e2386132f9236c5312977db6ac9ebdc242fd344d194d7c399718cd948959a03b"></a>

## Direct properties — rule_list.rules.destination_label_selector / 25d81f4294a5 / 3

<a id="canonical-8a65194fb6c9a1e99dc2f719757bc60fed46c0d91336f435a6d35c67179e5581"></a>

<a id="canonical-613358ab3e3cf553f4ca764a4bf9af66b4910f4624b8d592a9c8e27923c28e1d"></a>

## expressions property — rule_list.rules.destination_label_selector / 25d81f4294a5 / 4

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

<a id="canonical-5458875361bba33ea26a0845256532c56dd08253c6f8d5c54042e53d3b7a6c73"></a>

## Next pages — rule_list.rules.destination_label_selector / 25d81f4294a5 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-2213f0ef6031f6e976ebf63efcdaa06c5476897e06947435ea365800a416607c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7113c85f5c685a2a6036b6bf87667dd524523ce5c3328590557558610488b215"></a>

## rule_list.rules.destination_prefix_list — rule_list.rules.destination_prefix_list / dd88dbc521ed / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.destination_prefix_list

<a id="canonical-8820a44f40c3f68e9be7e87234cc2249fab9ccca9c82cbde36d12d8f74f68ab0"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-7484375a68e565ba73e0f14886e5373e15da58f338f02ca565b04c2708c76888"></a>

## Direct properties — rule_list.rules.destination_prefix_list / dd88dbc521ed / 3

<a id="canonical-deed3bd29f14ef6b2dccb16e2c551523f2668ef57f42e9be067959f79a60a1c4"></a>

<a id="canonical-17ea3fbb5bd6191cd080dbb9eaab39adafc032e4ca967b7ef9b50b425d410c9f"></a>

## prefixes property — rule_list.rules.destination_prefix_list / dd88dbc521ed / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1b1fc92751d753a65fbd7d14b63f8895ea74a311f833b83481ddfba0283ea5ec"></a>

## Next pages — rule_list.rules.destination_prefix_list / dd88dbc521ed / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-741d201d3a2b3a9b759f7bcbc72b6220435bf99158635473deaa5829c684e55b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa0ec746571edc69b32f6f7ef63fb3c9df97b089dc216dd172f78ef96bd82597"></a>

## rule_list.rules.insert_service — rule_list.rules.insert_service / 6903bbcbc75d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.insert_service

<a id="canonical-b1a99bd6b803fdde7973fc0b450403436684635ee6bcbb396d5309a59b9ee7a8"></a>

Type: `"single"`. Computed.

Action to forward traffic to external service.

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

<a id="canonical-01e0aa4add9db237398312e18ea52ac9890f948c20e653e619560a4618f9c79d"></a>

## Direct properties — rule_list.rules.insert_service / 6903bbcbc75d / 3

- [nfv_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-76f66f55b874602550fd798920a543fb23dc13f2731a4aed0a04f244840b2752): complete subsection reference.

<a id="canonical-d3293860a3ff540e6b9b7be7fd8bf02bb179612d03b99ea58f31aba9ff717dd1"></a>

## Next pages — rule_list.rules.insert_service / 6903bbcbc75d / 4

- [rule_list.rules.insert_service.nfv_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-76f66f55b874602550fd798920a543fb23dc13f2731a4aed0a04f244840b2752)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-76f66f55b874602550fd798920a543fb23dc13f2731a4aed0a04f244840b2752"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10cdad6d0275e390c192a4bd3d7540ec65a11a8e78689b5bc57cb8cc1b7adb11"></a>

## rule_list.rules.insert_service.nfv_service — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-741d201d3a2b3a9b759f7bcbc72b6220435bf99158635473deaa5829c684e55b)
- rule_list.rules.insert_service.nfv_service

<a id="canonical-5c7a4ecaa24bb41454752780432b608583b557423088eb225239bb34f57956c0"></a>

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

<a id="canonical-e03101ae6b8d2a97b6e4bc31758159e40622f11e533a9fb72747370083d31d3f"></a>

## Direct properties — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 3

<a id="canonical-e1a87933b7f066a5a0ba4f204139b832ac33cc8cd51836122de7e96e4e194766"></a>

<a id="canonical-049268be89dd12cd5f1b418a9d25d58fe90f0fd8d6ad32392a8bdf3c1080c682"></a>

## name property — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 4

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

<a id="canonical-7a51dd1ec04a1abd6715ecd803802ac76878c45ce9115ff199d818758be7b81c"></a>

<a id="canonical-da90f9d0edcc0d09e093144530ad35c6424060bb84f7ea035a2a6882777a26de"></a>

## namespace property — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 5

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

<a id="canonical-3eaf193ddbfe86708a9e75bc458f0c960634971d5de9d55a5201abc5017cfcba"></a>

<a id="canonical-70ca8f676bbc5cf88a6908babca04d06d21c8eb9091265b1e7ca77d6beb76148"></a>

## tenant property — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 6

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

<a id="canonical-7ceba1e4f9858ef52d1c63e369d3b52b442f54f1965833cdaf4343ce78f95ca6"></a>

## Next pages — rule_list.rules.insert_service.nfv_service / 4b3aa3baaac7 / 7

- [rule_list.rules.insert_service](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-741d201d3a2b3a9b759f7bcbc72b6220435bf99158635473deaa5829c684e55b)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-7f4cb07b9d4fcf864aad3e55e56d44042cf3327c81b9fe617ae71ca44bde8da5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8700d8b3ad70a798066e18f9396bf5379d57c16f80b3bb92f75d5295ac098b4"></a>

## rule_list.rules.inside_destinations — rule_list.rules.inside_destinations / 44945657395d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.inside_destinations

<a id="canonical-3a52088cd627f5d094b1cb4e125f389a039e10460affa7967b34ca5061b9878b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside destinations.

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

<a id="canonical-b5e3d05da4875a494bb58361c4107bdb017c753962a68bdb177941793c08cc8b"></a>

## Direct properties — rule_list.rules.inside_destinations / 44945657395d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46cff5fb3cbb2ddebb14c4f371e37990e0bb6cb0ac9c9ed157f6a3131af807b6"></a>

## Next pages — rule_list.rules.inside_destinations / 44945657395d / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-725987edd37546fcfa7546e0c604c9c126bcce9e71d5b67e75f2ed6e2f32130c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36fa9f1b88830ee1c1bd000de0164d3995a122141b94a6a4d2db0f970ff01196"></a>

## rule_list.rules.inside_sources — rule_list.rules.inside_sources / ab7a610d5147 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.inside_sources

<a id="canonical-0f02621664a169d1e3a5cdf10900c7cd4c11aedeff13f6862dc3471d73ae5749"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside sources.

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

<a id="canonical-5167bfe52f6fe4597eec5c6be4b99805a268774387100760aee1cbf87188b4d9"></a>

## Direct properties — rule_list.rules.inside_sources / ab7a610d5147 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1a02b11c77e9e730eff16e80e6f721d3ede92fa6ecf8925700ed9ec54f0742c"></a>

## Next pages — rule_list.rules.inside_sources / ab7a610d5147 / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-52df26bed3cf3d4c76e163ff7a90c450bcd39128201c972d5ab73c876e8e0b80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3dd126089376a073ce64939828a2dc38a8c016f27babb0e4d09a505719a5483"></a>

## rule_list.rules.label_matcher — rule_list.rules.label_matcher / 3d8d83ef4c9e / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.label_matcher

<a id="canonical-6a6f46db43e88b30faf8a8be51f7cb012945e9652199d244dadedebba2a359eb"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-030b466129463621615ee712aff13c9421e8cce18a62b09c9f6b9f0662e8dde1"></a>

## Direct properties — rule_list.rules.label_matcher / 3d8d83ef4c9e / 3

<a id="canonical-a7adde13fb13242ac1315897c0cc84c7ab1326303724325e05bc48519036afbf"></a>

<a id="canonical-e1b41efd01223dd8f8c477880a5d375918cb794c536964af784e3861b21f84c0"></a>

## keys property — rule_list.rules.label_matcher / 3d8d83ef4c9e / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ca1f6737edf2b5165b01f78edc58fb4429c0460f94f18500043a09caab7f1b97"></a>

## Next pages — rule_list.rules.label_matcher / 3d8d83ef4c9e / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-3926dfc10ab2393042b86f02ca54db32d33b74bbbb70ee9d01dd507e435a74f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af35a9aebe0fe6bd408177e05ad079c3de564712a7cfbb954dc46f29b8db0903"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / 28afe3faf5a2 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.metadata

<a id="canonical-15ba39fa6ca98399e8bef0d0879e8c512a7d19c0b687e14f1998b038e985a1f3"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3bbe527073d6ef24c9c7a651179496b83a2d2ed34d8c790cc20a0c1e3acaee91"></a>

## Direct properties — rule_list.rules.metadata / 28afe3faf5a2 / 3

<a id="canonical-bb29b528fedec842355d9416a91a858ae71082065f5c6300bc073c8a39cbb355"></a>

<a id="canonical-5ccaf6dbc47b18091c66d718f8c4a83fba98efe07a845eff4b39fc1000148c20"></a>

## description_spec property — rule_list.rules.metadata / 28afe3faf5a2 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-4106ce897a1efaf7d321e4bbfab82ebeccc24c36c5c5fce2ef46807062493c07"></a>

<a id="canonical-e4fa6260d7c0eefed1261b3a14c40a762c9836a3f27361fba1d4a7f90a16bc91"></a>

## name property — rule_list.rules.metadata / 28afe3faf5a2 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-77732427c76eb2db2ea399992919e074fdf9a7a0993cc016d36f51926bdd0247"></a>

## Next pages — rule_list.rules.metadata / 28afe3faf5a2 / 6

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-2f3c4cc7e70ef7458b34d41965a84b259837dd5fbf18ce893421c4c61d546f91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8a0f0b776731b4d27665eb0ec2015b48ecd52a455c28c035c47a89ef65b90ed"></a>

## rule_list.rules.outside_destinations — rule_list.rules.outside_destinations / 26436989095e / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.outside_destinations

<a id="canonical-7ca882e1512af0d9460000014edf44542f488a9702a3e0aaeaef5eef2ac4674a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside destinations.

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

<a id="canonical-b44f8b717940ef0295eed5def3a002cdc0b3de97c304481978252378f0bc43a5"></a>

## Direct properties — rule_list.rules.outside_destinations / 26436989095e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20d5e1ab4a68a218f91ad7c9ad4ea5489c49b240da4c688a4e3ed583cccb68ed"></a>

## Next pages — rule_list.rules.outside_destinations / 26436989095e / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-3d0568a82a224681a0356aab6e1ebf7e31076a3385c8183a369e07519ae12d54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ef94becbfb3bfbe056d9805f86d9e39206113c68c5266a59fad079fc35e0d45"></a>

## rule_list.rules.outside_sources — rule_list.rules.outside_sources / 69c18234403f / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.outside_sources

<a id="canonical-47f659fc0202416bf55a4c4a7c17a4fcabf32e9c3aed9dd53cab4e319e1b1fdb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside sources.

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

<a id="canonical-fe0184e9bf231814da28cdf3c0da5fa3293a6178996fd726831d6316eaf67648"></a>

## Direct properties — rule_list.rules.outside_sources / 69c18234403f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c54576eeaf3ef5e81ffd08a4dd88271fa9f36c86daf768f5e61b24de3cbe0aa2"></a>

## Next pages — rule_list.rules.outside_sources / 69c18234403f / 4

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-4490cd4db2103499cc04b0e433858740d8007fc86ac0cd6b167290c9369c4123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cbf6fb4e72e4c8565e1bfe357875039b17dd0d6d7e734076a036fb2281aec7b"></a>

## rule_list.rules.protocol_port_range — rule_list.rules.protocol_port_range / da4378f1e588 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.protocol_port_range

<a id="canonical-1be73ada0e10f669d0f752ba0c6e7d7c45b0fc1fc1916637d033be376d79fbd7"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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

<a id="canonical-57c551eb310080c4131b975c5dfff7bcdda3858581c3e0ed6ca9a767dc786af3"></a>

## Direct properties — rule_list.rules.protocol_port_range / da4378f1e588 / 3

<a id="canonical-b71ff82dd17fe0404ef7b4d395827a390b7164b5bd563ed3b912088c73c36eab"></a>

<a id="canonical-ba24df6180db227711551896887ce7c6197eafff8d14784c8cfbc8bb72cdb60f"></a>

## port_ranges property — rule_list.rules.protocol_port_range / da4378f1e588 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-5d70d8745bc115b243b8a36a2ae94f3d966bc1add387cd6381d3aa0b60801cf5"></a>

<a id="canonical-e741942dd4ca0e2b65bd5e9d234663c3c3c1f0517e6f7392288d726809d8dc2f"></a>

## protocol property — rule_list.rules.protocol_port_range / da4378f1e588 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-d7b5a2c94b904129607ee216e1fd024b7ca8e057a8e285a15cb340fa53a4f0bc"></a>

## Next pages — rule_list.rules.protocol_port_range / da4378f1e588 / 6

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-73f1d7ea85db0752d3ff02c9d7cb127d8a5972c272618e9036c6a365fdca8742"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ec65e4215e0eab7f57016c5d869a8600e3fba6d1d18ff867b4be8137c5fb6e8"></a>

## rule_list.rules.source_aws_vpc_ids — rule_list.rules.source_aws_vpc_ids / 1daf9688b951 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.source_aws_vpc_ids

<a id="canonical-a6640dd4b2d1614e17ccd2092095446e877176e713fd539be94820aad518e7b7"></a>

Type: `"single"`. Computed.

Configuration parameter for source aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

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

<a id="canonical-765c85783ff5ad95742ad717808da13ae86f2e7e91313a9c7ba40be56f919e01"></a>

## Direct properties — rule_list.rules.source_aws_vpc_ids / 1daf9688b951 / 3

<a id="canonical-6731a893bd681257846ea8b0fb711bc33981c8d4fce79a052ce24e892d0f3f91"></a>

<a id="canonical-21862874f391b72b2713587bb7a589cded4b286954796be9d2cb5ac88213ab7a"></a>

## vpc_id property — rule_list.rules.source_aws_vpc_ids / 1daf9688b951 / 4

Type: `["list", "string"]`. Computed.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-16b831f459376852f5ed294acccbd5d2a754a0565b45c7ddc3ce0233460802a9"></a>

## Next pages — rule_list.rules.source_aws_vpc_ids / 1daf9688b951 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-202eddb57b6eec894059373000f22e171a2668303e5f41183bea0d7c30ce4aff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f95a10825d112ff7b397bf539eee6a39897d6ec09fe17b4128310ced74e2022"></a>

## rule_list.rules.source_ip_prefix_set — rule_list.rules.source_ip_prefix_set / 7b011bdb8fef / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.source_ip_prefix_set

<a id="canonical-f5064eae810b290c5af85aa9830d12dc356444c8488540357d25fec686c5e070"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-a56b3daf5abbda2975c0ed885c7cf920b3cda85833a7a47ba542adf1fcfc5d47"></a>

## Direct properties — rule_list.rules.source_ip_prefix_set / 7b011bdb8fef / 3

- [ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-df24f5f226c299d97b3f57b4f148559b0762442917c1eb8a0e3c7460c00ee37a): complete subsection reference.

<a id="canonical-b4fa790ac320c4071ecd0b57a37d331275813f5dabfcef1abd4b2521ac8dd7a6"></a>

## Next pages — rule_list.rules.source_ip_prefix_set / 7b011bdb8fef / 4

- [rule_list.rules.source_ip_prefix_set.ref](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-df24f5f226c299d97b3f57b4f148559b0762442917c1eb8a0e3c7460c00ee37a)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-df24f5f226c299d97b3f57b4f148559b0762442917c1eb8a0e3c7460c00ee37a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62b592770fa91301d1214d6e1b7e3c19612e0217e2023e34742a3dd330a17c3e"></a>

## rule_list.rules.source_ip_prefix_set.ref — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-202eddb57b6eec894059373000f22e171a2668303e5f41183bea0d7c30ce4aff)
- rule_list.rules.source_ip_prefix_set.ref

<a id="canonical-74db55bc30473778c3564668f086949304ea30258016cb32d7e9dd94db89d04a"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-e395cda5e2a4dc5bbef857a49f4713e90e50918505ef64250141ae14f48a1593"></a>

## Direct properties — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 3

<a id="canonical-e9544a0042491d359d2d954db060f5919e250721c6edb51b69c8882b032ea717"></a>

<a id="canonical-e18f580e40f7f46680cc3b0fb9349ba019519baf5bd1031a23712f2d329a4cab"></a>

## kind property — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 4

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

<a id="canonical-a182ccc0988b735aedef54291767785a44ca3ababffa862ee17d8fcb7082c7b2"></a>

<a id="canonical-a5f5bf7634d0c97713a30b9e908d42566344753f5ec559296f9209d8847ac993"></a>

## name property — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 5

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

<a id="canonical-9bd7a6149e1606ba770b929b190cb2d253fb3abcc07c9b4063fc8c83f31b587d"></a>

<a id="canonical-f42e30a793c9250e4cc625f04be1f79cf72051145ab3a9858c341008444e6961"></a>

## namespace property — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 6

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

<a id="canonical-e3b1e726923f07e20e252ed8460f767a93fba183a4a2608b0000b64922e356d8"></a>

<a id="canonical-2051b98943681e20ccee9e3b39f7e39dcc6f6e09c30979e4b977fd6bd7b52c80"></a>

## tenant property — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 7

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

<a id="canonical-706f3021ecc042210fca7740b3d442ba1ae818aed4e44286b9f1bf2679b58daf"></a>

<a id="canonical-a63e3c17561dc8ce04c98824620c019ee4eb791bcb17b0c60b3d4dcbf280daaa"></a>

## uid property — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 8

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

<a id="canonical-6771a38e21d56cc9d658f28de38f97b0fecfc0113177e1961473175d79645767"></a>

## Next pages — rule_list.rules.source_ip_prefix_set.ref / 6fe6cbe5f5f8 / 9

- [rule_list.rules.source_ip_prefix_set](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-202eddb57b6eec894059373000f22e171a2668303e5f41183bea0d7c30ce4aff)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-052ed70ca5fab404fbdc1c7883895d85223e107b0324b6ae3179b5cbfbde16b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa303a4aea4baf8511577643c4eb3c398baa1458d5efe1b74820aacae2403073"></a>

## rule_list.rules.source_label_selector — rule_list.rules.source_label_selector / ce4221d9d640 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.source_label_selector

<a id="canonical-ef4c95c0f08be56cf7925a0be9442dc09e51f130cfba8b35921bf7d352de8aa3"></a>

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

<a id="canonical-f321ca6e9832ad0844dde2e93beddd8722c5621debc8cbc374fe9fc2259dcc33"></a>

## Direct properties — rule_list.rules.source_label_selector / ce4221d9d640 / 3

<a id="canonical-a0551567620e54dd4371d62bbdb5423e982808087d4c6d0a7e39991751d94270"></a>

<a id="canonical-1fad2b612b0ef8a16358b6e7b586ecd092153a35b6fb80a4bef71c39c88fb825"></a>

## expressions property — rule_list.rules.source_label_selector / ce4221d9d640 / 4

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

<a id="canonical-2bf50db9857a551d1b0b13c215283774385836c5056f8d64db8690c5604dd8ee"></a>

## Next pages — rule_list.rules.source_label_selector / ce4221d9d640 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-2e1692091408120c306e23d69e92ddcfb142cab3b29f7498f8f2fcab07380bf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b392df1c2db1fae0b96697d1eb0f3ee0f330dd9251e7e3e6bb508a8f83385a90"></a>

## rule_list.rules.source_prefix_list — rule_list.rules.source_prefix_list / 8152fbd6a5e4 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Property reference](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [rule_list](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-aaa6e5e1befad8e00258a4b5407ce94145af07c4f00c636d5c66d86047b38f46)
- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- rule_list.rules.source_prefix_list

<a id="canonical-b0b0dbe39f0a8d1218386d2132f8b902d6df5bbe56d568a753d3569922298a7d"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-f220bf10cccffc8607ef600c245dd2ef349ceed9cbac5bfaff52ac3dd00ffe52"></a>

## Direct properties — rule_list.rules.source_prefix_list / 8152fbd6a5e4 / 3

<a id="canonical-888547bdb3fc7048469c45e1d75cbb0ea0c26fe4c040ad145d8def11ca5bf3d5"></a>

<a id="canonical-01e4498deca6029fe2f9333503c870af09ebb5219a0508b85fdb42846ce949c8"></a>

## prefixes property — rule_list.rules.source_prefix_list / 8152fbd6a5e4 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7e2d3b8afb51adb3f6ed4d360c81d63df132b3af5107fde4ba87b89890a977eb"></a>

## Next pages — rule_list.rules.source_prefix_list / 8152fbd6a5e4 / 5

- [rule_list.rules](data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-eb79ecede020d78c0aa7214dbfdcb87e6f2c1310ccf6941cc5b4fd22b4ad42bd)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
