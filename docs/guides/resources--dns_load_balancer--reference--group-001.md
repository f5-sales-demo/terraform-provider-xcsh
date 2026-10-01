---
page_title: "xcsh_dns_load_balancer reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer reference."
---

# xcsh_dns_load_balancer reference

<a id="canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df2baca58bb9f3d5a54c5f06c0a22c59c79266a7c00542fbc99a5c91067f6f23"></a>

## Property reference — Property reference / d3f57d8f26c4 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- Property reference

<a id="canonical-054f266ace725fdc8c34547e20dc1778c905eaef8079e71068232121dfce44f2"></a>

## Direct properties — Property reference / d3f57d8f26c4 / 3

<a id="canonical-1b5949113572e6365ce2f1ed657b0a3aba2f6233967d3db7d68124943573c685"></a>

<a id="canonical-5e7ab2f78003158bdca9464d654512a7d807a84aa7795a061e0191d677b320c6"></a>

## annotations property — Property reference / d3f57d8f26c4 / 4

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

<a id="canonical-ffe2493580cecbc1c9c2d58d14e4e93b744cd49e1d156c5cc8d7e5495899a604"></a>

<a id="canonical-8b3a0fcbea3b102c9ca130199a2fb112d28ccd9c3bfd50d0ab9548dd9924bb7e"></a>

## description property — Property reference / d3f57d8f26c4 / 5

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

<a id="canonical-089dc3abfbb2a87ed8bb8d27230db20b0750733d940ea1cc04c37fbb84901979"></a>

<a id="canonical-ff2ee6de43a82a09a15beb3949f2a8efa55cc12ea5ee8ed29bbdffbb563a5696"></a>

## disable property — Property reference / d3f57d8f26c4 / 6

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

- [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-66903525c1859b667ea2957b50840ae8b11821e2bd1d1929a6d161fa6b678948): complete subsection reference.

<a id="canonical-ec3e36b5d2cb4f33a56b3358baf436dca4f445da70f57ba7fffeb19837e31638"></a>

<a id="canonical-3d0ca48585cfd19ff390ceaa73f49a5c0964f559765c1e7b40cd6770fa8d77a0"></a>

## id property — Property reference / d3f57d8f26c4 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-27b92831365542162ed4d43fcab0fb857b82f800148172f41fcb7ddacb79ea08"></a>

<a id="canonical-67f5be4ffdc934d22c3b3d05cd9f3373a61067804bb17895748571c8be4a4aeb"></a>

## labels property — Property reference / d3f57d8f26c4 / 8

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

<a id="canonical-fcb2128d307839183eb83acb324520819d20f073b769563ea4ebf11276800ae7"></a>

<a id="canonical-3cbfeeeaa1531b3b4b25d8b9ea6d0b95b1bf28b85533ceff7c34300bdca687f8"></a>

## name property — Property reference / d3f57d8f26c4 / 9

Type: `"string"`. Required.

Name of the DNS Load Balancer. Must be unique within the namespace.

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

<a id="canonical-1d3056565ea77cd6470cf56b248f5f586200d343c3db0742db6b30a185ef5302"></a>

<a id="canonical-6fa3d9c3c660adf1f9c5982e413dc2603a9daf7492d06c0dbd2bae62f521d39a"></a>

## namespace property — Property reference / d3f57d8f26c4 / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS Load Balancer. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

<a id="canonical-b67624f8489f4d5844893b87099eeb867012293e355668968a00e378ab30fce3"></a>

<a id="canonical-1a3eace4b00197230b2737029f6fce83901a7ec76212d67534fa7a9f5a6a6ab8"></a>

## record_type property — Property reference / d3f57d8f26c4 / 11

Type: `"string"`. Optional, Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"),
}
```

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

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f): complete subsection reference.

- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b): complete subsection reference.

- [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-c8451a107d6f68291ff2f63e446d22ac295d5a644c522b0fb13c71b284a4dbc5): complete subsection reference.

<a id="canonical-89996f85f042add2a4d9833f91d6f017344c9aebe3508b22dea217925d79d95c"></a>

## All schema paths — Property reference / d3f57d8f26c4 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_load_balancer--reference--group-001.md#canonical-1b5949113572e6365ce2f1ed657b0a3aba2f6233967d3db7d68124943573c685) |
| `description` | [description](resources--dns_load_balancer--reference--group-001.md#canonical-ffe2493580cecbc1c9c2d58d14e4e93b744cd49e1d156c5cc8d7e5495899a604) |
| `disable` | [disable](resources--dns_load_balancer--reference--group-001.md#canonical-089dc3abfbb2a87ed8bb8d27230db20b0750733d940ea1cc04c37fbb84901979) |
| `fallback_pool` | [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-99572917375b49cabd73f86b55c555ff526031d7829ac39b5a18464a7426c473) |
| `fallback_pool.name` | [fallback_pool.name](resources--dns_load_balancer--reference--group-001.md#canonical-bc7898104725bcf03e8cacf83b41cb5627bf0ca32bedd00f69bf627742eb90d1) |
| `fallback_pool.namespace` | [fallback_pool.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-b3cd9c822c0f6fbb757a80812a0859e09a48d7ea70889ce8c8412d61465a9130) |
| `fallback_pool.tenant` | [fallback_pool.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-6f1e20a8138a7974aa8d43be253edb1ab0e0bb01654b25a0ef88dd9cb0523f3a) |
| `id` | [id](resources--dns_load_balancer--reference--group-001.md#canonical-ec3e36b5d2cb4f33a56b3358baf436dca4f445da70f57ba7fffeb19837e31638) |
| `labels` | [labels](resources--dns_load_balancer--reference--group-001.md#canonical-27b92831365542162ed4d43fcab0fb857b82f800148172f41fcb7ddacb79ea08) |
| `name` | [name](resources--dns_load_balancer--reference--group-001.md#canonical-fcb2128d307839183eb83acb324520819d20f073b769563ea4ebf11276800ae7) |
| `namespace` | [namespace](resources--dns_load_balancer--reference--group-001.md#canonical-1d3056565ea77cd6470cf56b248f5f586200d343c3db0742db6b30a185ef5302) |
| `record_type` | [record_type](resources--dns_load_balancer--reference--group-001.md#canonical-b67624f8489f4d5844893b87099eeb867012293e355668968a00e378ab30fce3) |
| `response_cache` | [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-45d76d4f9bc7b80a2ece9fd77eb148ac9c6a9c735d82e1b0cd5733e84d053eb6) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-9662f3e8a9e8b2d1d6574f51104aa418c8f04df97d0ff0b21fcc0f833d5f4ae0) |
| `response_cache.disable_spec` | [response_cache.disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-a5c6129bb619dc038c1ee28aa8f6a42c72b2f7ea60c3b883bd39cfff05f09381) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-afdb6a4937e413ca5b73cdd7444b6a0c14fb3cf2ac14b200fba0b800e63c132e) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](resources--dns_load_balancer--reference--group-001.md#canonical-229222a7930ceced5a005778f5987674a09a262b029990af46ecf6d872852951) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](resources--dns_load_balancer--reference--group-001.md#canonical-3b53289f51fa9944a8458ce30b401ed311487170f7c76b893c0e147e5c7d15f3) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](resources--dns_load_balancer--reference--group-001.md#canonical-eb9c3bcad33aba50c18abe9a72bdb17c3252eb98696bf787ae9084fc73d308fe) |
| `rule_list` | [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-453307fc6ebfd02301a311068afea86dd882d674c7b20badbd671143b232e326) |
| `rule_list.rules` | [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-6e8867a89cae12b0f0385f2e1deab07aed1e12e5d4c73e5efb0f8aa008205418) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-3355076f9a428b313455e2b96b84bea8fcd226b702634bd2dedc70227f02697c) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](resources--dns_load_balancer--reference--group-001.md#canonical-0e8999c58518d4d16d3474825e044146d67512ed7fd02642da4414f544c02401) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-912990b2755d16f308be788bde1855af5c1e8dcfb0340bb6108bd56c23ff51f2) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-f1414d23b2696d9900dfd0ab572b9d2768721e13f8e325ddc62cb7ee9338e08c) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](resources--dns_load_balancer--reference--group-001.md#canonical-509aab550b0c2d9ea350d0f660a52d39c7328ebc98d5f1170b3ea05351d7f2e5) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](resources--dns_load_balancer--reference--group-001.md#canonical-1e56d9ba38157c4137e53fc9f73872543e7e204b0bab3738a0e098ea4acbd228) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-9f8bb2906147e4ae12debfd31f71528c4ee38e7f1b84f075e07b795cc444c82d) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-f3e531e3eaabf86fcdf6bb5c7cfdbd5b700eaaf4a935a3d077fd3bdfcbec31fb) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](resources--dns_load_balancer--reference--group-001.md#canonical-dadc8df59fcaf3664846f4e144720b529aae298b9ac82d4912c0dfd262cffe83) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-db55335b29e3cb971d42f76b55400c88c5f4c4b7b2aaa5063d64a6e047067e61) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](resources--dns_load_balancer--reference--group-001.md#canonical-7e011c2d7d65c804f7bd8e31b1c8db35fca9c9a0f6fef7b174e82f808ce2c51b) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-dbf13acab9ba22936a1856977ce734b73c37e7197e8c61ad7ed225eb19384d12) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](resources--dns_load_balancer--reference--group-001.md#canonical-b75ce5ac848c03175d6fc61e490514e9b84af9ddaf1a18708825667852515f01) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-aa6677edda6ffd7e56e9ed55648b29d412002afc9c97c95c53835951a58475cb) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-b2be701b773d0ed3c47c53580435ee9243ddd0d0717e69c1fd0849c6080b091e) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-41341d04982ae49f2dd77d2a52f6174aab37dc97200685de5cfd9b6cbe3a05cc) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](resources--dns_load_balancer--reference--group-001.md#canonical-2a2dc22cf9befc076765bc7be2bebbddee5238c71bdaccbe460c1713e1782bba) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](resources--dns_load_balancer--reference--group-001.md#canonical-e863496bc7abe1111b9d228737b92136da36ceba3cefde7db959a1807a5c3fac) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-228731d8ef66958a051c95c3a38de3baa45a7a659a3e7eb32d1ad6ebc1faca93) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-de95805c4a8b121b375b68cd668c3ed6ac82b8599b3237775a8cdadfad02c89a) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-6693edc51b380eab83e71414e40b26f1cd09de3d402d97b5b5cd4709213336d0) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](resources--dns_load_balancer--reference--group-001.md#canonical-c0fd46901b0c336d708d4aa6924f374e0ae0fb4f0c610ab7dc3a1dd50e3bd5ca) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](resources--dns_load_balancer--reference--group-001.md#canonical-afa5cdeda5c8c1c635bdf338eba00ee6703d9ffc8b001294eb206d245d71ce55) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-d6ed60d8fec4a731e4c9987957165cc9efa2b0dc4213f8e0935ad12e915ed7e2) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-129a71e2fa38357fd51c2c0769eac96afa80a9e0d4dbd98891565a38ec5d02c7) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](resources--dns_load_balancer--reference--group-001.md#canonical-7d73206f1617e6d33f1c79869441f347d617d311e9d5d8d278f0d48b9375982c) |
| `rule_list.rules.pool` | [rule_list.rules.pool](resources--dns_load_balancer--reference--group-001.md#canonical-35d606fa8d7100d1246630f56b3dfc6e95ddd8eb0e7c62749d33b2ffb01a304a) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](resources--dns_load_balancer--reference--group-001.md#canonical-d47de1c496baaf1023ba658ec796baffd0858058305f5e1622f39ab045e479f2) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-4aaf4904d0ce07863d906518f62770a2cfe13026abcdbf8b07e80749421a07d8) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-385e6f7fa963843bb764c34a2782fb70b76415e5ac7a1134a272db18d05aede7) |
| `rule_list.rules.score` | [rule_list.rules.score](resources--dns_load_balancer--reference--group-001.md#canonical-c76dd0fa11206f7eaf75d1cfc4530602788da6baca96e84c0ae23b9f003cc230) |
| `timeouts` | [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-de31dce4f63da5424332ab0ef92d64f7c784a2e5b26d38a99070145d116c0f15) |
| `timeouts.create` | [timeouts.create](resources--dns_load_balancer--reference--group-001.md#canonical-05b55872e4476bcc2c6d5f21836575bf451507690b2112a3824e1feefe575590) |
| `timeouts.delete` | [timeouts.delete](resources--dns_load_balancer--reference--group-001.md#canonical-2838e77b0a0e115b786553838097a11f43c24f0547569bf412e0da6c8b552c71) |
| `timeouts.read` | [timeouts.read](resources--dns_load_balancer--reference--group-001.md#canonical-34d3ecc0ff4fb08f7570a922a133856f28aa240f3c40857b528c2477f127c569) |
| `timeouts.update` | [timeouts.update](resources--dns_load_balancer--reference--group-001.md#canonical-22fb306612ba42b47b3b8f2e11daadc3e87381d0c8b9decdd256d31ffd35ed69) |

<a id="canonical-7dfd5fa76b51aaa92f38120ed55b9b278d27e9c4135a65b5369f454fef57d8b1"></a>

## Next pages — Property reference / d3f57d8f26c4 / 13

- [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-66903525c1859b667ea2957b50840ae8b11821e2bd1d1929a6d161fa6b678948)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-c8451a107d6f68291ff2f63e446d22ac295d5a644c522b0fb13c71b284a4dbc5)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-66903525c1859b667ea2957b50840ae8b11821e2bd1d1929a6d161fa6b678948"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd4e3b2974b2d0687d397b3ae9be34c3fbc1bff7483e59a731bc5ff42702f9df"></a>

## fallback_pool — fallback_pool / e04560c54585 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- fallback_pool

<a id="canonical-99572917375b49cabd73f86b55c555ff526031d7829ac39b5a18464a7426c473"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
fallback_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1e8c3a7e04cb74a04f1c774515455f6f292a440195bb757c040a2af714faf91"></a>

## Direct properties — fallback_pool / e04560c54585 / 3

<a id="canonical-bc7898104725bcf03e8cacf83b41cb5627bf0ca32bedd00f69bf627742eb90d1"></a>

<a id="canonical-196f729de5e83c6df74770c62979f3320cd10dbfe24bca54d8d0a83990981a36"></a>

## name property — fallback_pool / e04560c54585 / 4

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

<a id="canonical-b3cd9c822c0f6fbb757a80812a0859e09a48d7ea70889ce8c8412d61465a9130"></a>

<a id="canonical-512706a89c858cc87bfeca42f7dab9856be5d1c4a3af0244e9c780bf78fe2410"></a>

## namespace property — fallback_pool / e04560c54585 / 5

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

<a id="canonical-6f1e20a8138a7974aa8d43be253edb1ab0e0bb01654b25a0ef88dd9cb0523f3a"></a>

<a id="canonical-86e84e98167479c17e40245ad1db6ff4a8f0404d01d230c066b65bfa9c9edb20"></a>

## tenant property — fallback_pool / e04560c54585 / 6

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

<a id="canonical-59215513ae695ee14a53d33ccd001b53fdd7092421f2149370f989cc848b1ff5"></a>

## Next pages — fallback_pool / e04560c54585 / 7

- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6318c84bb0e9a385992a1b8feb6227b71b6e6bbd27885fcbb78ccc9c0885208e"></a>

## response_cache — response_cache / 4745e5663357 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- response_cache

<a id="canonical-45d76d4f9bc7b80a2ece9fd77eb148ac9c6a9c735d82e1b0cd5733e84d053eb6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "response_cache_parameters"),
  validators.ConflictingObjectAttributes("disable_spec",
    "response_cache_parameters")}
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
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

Terraform syntax:

```terraform
response_cache {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0d7b08d1e6a1c26fccd4cfd565b974cecc90fc82b9b9f4c3785b784fd6cf1e1"></a>

## Direct properties — response_cache / 4745e5663357 / 3

- [default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-8e8eb77278de0e606f8760cffda6317c255d49580400d9217362ea884c1a32ae): complete subsection reference.

- [disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-6bcd535c73a17941014c6afe7a24308e3cf68b5699a88c2ed9996bb0d16df03b): complete subsection reference.

- [response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-fd583ad9838fb4bce20e42007ccb9a45845ccb99681646797041e216c47a039c): complete subsection reference.

<a id="canonical-c20cf3eb7def16945f8d0ae2fed7588fed8827ca4096811f6419da9fae685055"></a>

## Next pages — response_cache / 4745e5663357 / 4

- [response_cache.default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-8e8eb77278de0e606f8760cffda6317c255d49580400d9217362ea884c1a32ae)
- [response_cache.disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-6bcd535c73a17941014c6afe7a24308e3cf68b5699a88c2ed9996bb0d16df03b)
- [response_cache.response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-fd583ad9838fb4bce20e42007ccb9a45845ccb99681646797041e216c47a039c)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-8e8eb77278de0e606f8760cffda6317c255d49580400d9217362ea884c1a32ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b00a3397d34f5540fa38e333865fb432d5f2bffa588e42e4279253bf67776109"></a>

## response_cache.default_response_cache_parameters — response_cache.default_response_cache_parameters / d53e1a444fb1 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- response_cache.default_response_cache_parameters

<a id="canonical-9662f3e8a9e8b2d1d6574f51104aa418c8f04df97d0ff0b21fcc0f833d5f4ae0"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_response_cache_parameters = {}
```

<a id="canonical-7aef614d5fd908d8f759578aeda30bad271415637b4c9146106737f5fb34a7ea"></a>

## Direct properties — response_cache.default_response_cache_parameters / d53e1a444fb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-668a0b42dfc77c65ec107dfc0e5ec976f38ab853bf004686c90b23e757fffb7c"></a>

## Next pages — response_cache.default_response_cache_parameters / d53e1a444fb1 / 4

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-6bcd535c73a17941014c6afe7a24308e3cf68b5699a88c2ed9996bb0d16df03b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0397523ec3504bc2bebdc44b3f20df1a6f4c897716f07e11cb4a043a4ec7e6aa"></a>

## response_cache.disable_spec — response_cache.disable_spec / bf9b68ce6476 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- response_cache.disable_spec

<a id="canonical-a5c6129bb619dc038c1ee28aa8f6a42c72b2f7ea60c3b883bd39cfff05f09381"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-cdfc010945890ff8d2a9bf8c243e48377bf372641af300f94f328527db1a4e22"></a>

## Direct properties — response_cache.disable_spec / bf9b68ce6476 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47a9b9f1c3e0b0f4b454212b65b304b048c9c3f7261685fe93a7e666c62c947e"></a>

## Next pages — response_cache.disable_spec / bf9b68ce6476 / 4

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-fd583ad9838fb4bce20e42007ccb9a45845ccb99681646797041e216c47a039c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b7978baa7eda8a97efa04f3ab109a33d6ca719881b024bf5fb2cbfeda1864d4"></a>

## response_cache.response_cache_parameters — response_cache.response_cache_parameters / 91af17109901 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- response_cache.response_cache_parameters

<a id="canonical-afdb6a4937e413ca5b73cdd7444b6a0c14fb3cf2ac14b200fba0b800e63c132e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_cidr_ipv6")}
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
response_cache_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2ac9c352dd7f6f8707b247c9828cdee405578fc205cf48098be34bf932de162"></a>

## Direct properties — response_cache.response_cache_parameters / 91af17109901 / 3

<a id="canonical-229222a7930ceced5a005778f5987674a09a262b029990af46ecf6d872852951"></a>

<a id="canonical-c80d425b6cc6af702b5ab55dcacb199e155bde5a6c3535ea98b7b8274782205f"></a>

## cache_cidr_ipv4 property — response_cache.response_cache_parameters / 91af17109901 / 4

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv4 clients.

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

<a id="canonical-3b53289f51fa9944a8458ce30b401ed311487170f7c76b893c0e147e5c7d15f3"></a>

<a id="canonical-1da1c63194ff86ba2d65dc8ea7d862a9d0b2282e166ea35847dc4f9b5c8162c7"></a>

## cache_cidr_ipv6 property — response_cache.response_cache_parameters / 91af17109901 / 5

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv6 clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 128),
}
```

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

<a id="canonical-eb9c3bcad33aba50c18abe9a72bdb17c3252eb98696bf787ae9084fc73d308fe"></a>

<a id="canonical-0d846495ff4cab309d378024bc523d2f41b919bb2833abaa64b5ddde0e392e9f"></a>

## cache_ttl property — response_cache.response_cache_parameters / 91af17109901 / 6

Type: `"number"`. Optional.

TTL. TTL for response cache.

Upstream description:

TTL for response cache.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

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

<a id="canonical-41c617bf2fa6d42028acc586ba19215186a6f2ec5dabf9f66b01d76f458c8de6"></a>

## Next pages — response_cache.response_cache_parameters / 91af17109901 / 7

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-aa60266cd4394ce959a7090db8cb4bc1acc46a79e575d6446390ae4f894e258f)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67182c0ac0a77d69129b0a302ae0ecabfbb6612f716dd4104b2d99029e43397a"></a>

## rule_list — rule_list / ed3a8da825de / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- rule_list

<a id="canonical-453307fc6ebfd02301a311068afea86dd882d674c7b20badbd671143b232e326"></a>

Type: `"object"`. single nested block, Optional.

Load Balancing Rule List. List of the Load Balancing Rules.

Upstream description:

List of the Load Balancing Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-aa7f15bfcb3637cfb71029ae52e8db46ff80ff2fe7650c27aacd78f17a0c1777"></a>

## Direct properties — rule_list / ed3a8da825de / 3

- [rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419): complete subsection reference.

<a id="canonical-c26a4b5b1f53597b293e85de996a92ea32404f2fa3668ee51608d259ef6f690d"></a>

## Next pages — rule_list / ed3a8da825de / 4

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90f4c763076d899a183a8d557521dbae2447147cd3519b401584c022dbdd7997"></a>

## rule_list.rules — rule_list.rules / 970573a6a01c / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- rule_list.rules

<a id="canonical-6e8867a89cae12b0f0385f2e1deab07aed1e12e5d4c73e5efb0f8aa008205418"></a>

Type: `"object"`. list nested block, Optional.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("score"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("ip_prefix_list",
    "ip_prefix_set")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-8d08ef2d62bf102d3cf1a258bc418d98a5486c5cad8a90c5d2b118593ce402bf"></a>

## Direct properties — rule_list.rules / 970573a6a01c / 3

- [asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-f9aeaf0086c209612524a0ab7b42a4179cf12698a48893112d990991474d6f17): complete subsection reference.

- [asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-9b0d9fe5b21e6df25b92fb28e8120d54daeaadfe89861528804fbf9956bb2021): complete subsection reference.

- [geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-3ebffd5e0fb62cc862c46acbd943c541caf87fec78c19b2bdf15f333cc0c2824): complete subsection reference.

- [geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-feccbd9e566d993c7883d7964810bb448b71dd4177b045564885e67302d74b94): complete subsection reference.

- [ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-afd66fcdb230a81581a465d7aadf0341288526ec465c09aad54f941c4550abf1): complete subsection reference.

- [ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-494d9c08d7d1c46099f27c5c5b012b09b4d65f820e78387576adab90ba62e3fd): complete subsection reference.

- [pool](resources--dns_load_balancer--reference--group-001.md#canonical-44f4ea735f2b8602f2c096c96af203ceb25b8858714a8dabc8c0ad7ca8129209): complete subsection reference.

<a id="canonical-c76dd0fa11206f7eaf75d1cfc4530602788da6baca96e84c0ae23b9f003cc230"></a>

<a id="canonical-cb399daeba7d39b77f34c7a22dbabb9b1aad4e5ec97b058c222dc68241df0be1"></a>

## score property — rule_list.rules / 970573a6a01c / 4

Type: `"number"`. Optional.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32767),
}
```

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

<a id="canonical-a658f7835dd3f2687b8f065850af2b8b7367641c7ea3cd58ae6adc9048311e4e"></a>

## Next pages — rule_list.rules / 970573a6a01c / 5

- [rule_list.rules.asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-f9aeaf0086c209612524a0ab7b42a4179cf12698a48893112d990991474d6f17)
- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-9b0d9fe5b21e6df25b92fb28e8120d54daeaadfe89861528804fbf9956bb2021)
- [rule_list.rules.geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-3ebffd5e0fb62cc862c46acbd943c541caf87fec78c19b2bdf15f333cc0c2824)
- [rule_list.rules.geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-feccbd9e566d993c7883d7964810bb448b71dd4177b045564885e67302d74b94)
- [rule_list.rules.ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-afd66fcdb230a81581a465d7aadf0341288526ec465c09aad54f941c4550abf1)
- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-494d9c08d7d1c46099f27c5c5b012b09b4d65f820e78387576adab90ba62e3fd)
- [rule_list.rules.pool](resources--dns_load_balancer--reference--group-001.md#canonical-44f4ea735f2b8602f2c096c96af203ceb25b8858714a8dabc8c0ad7ca8129209)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-f9aeaf0086c209612524a0ab7b42a4179cf12698a48893112d990991474d6f17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9c246694cf6381a7e30862ad98463ca4ab6fb2602c91d43f91ccc9798e1cbd4"></a>

## rule_list.rules.asn_list — rule_list.rules.asn_list / 81026bc57ad0 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.asn_list

<a id="canonical-3355076f9a428b313455e2b96b84bea8fcd226b702634bd2dedc70227f02697c"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea38f783d7c070d80f1c9aa9773f81f1595c5550ad3842a9c7f4a4c0343ddfcd"></a>

## Direct properties — rule_list.rules.asn_list / 81026bc57ad0 / 3

<a id="canonical-0e8999c58518d4d16d3474825e044146d67512ed7fd02642da4414f544c02401"></a>

<a id="canonical-c54ecf7ba5be1ffac4b8a350810b7b60a872336dc27633739b7523f85a756042"></a>

## as_numbers property — rule_list.rules.asn_list / 81026bc57ad0 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-8c4d91e332623674dc156592d213fbd9a83b5f899c96b0ac6bc755fb34be0eef"></a>

## Next pages — rule_list.rules.asn_list / 81026bc57ad0 / 5

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-9b0d9fe5b21e6df25b92fb28e8120d54daeaadfe89861528804fbf9956bb2021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12e3e64b0c030be89d3c6ff803ecfda1208b332cca8683ea85afad2ff3917633"></a>

## rule_list.rules.asn_matcher — rule_list.rules.asn_matcher / 79c511cf8b39 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.asn_matcher

<a id="canonical-912990b2755d16f308be788bde1855af5c1e8dcfb0340bb6108bd56c23ff51f2"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-26c7d2951d13982e94d6172ed834249b0b2715c77069145515aa0b13ea993f0d"></a>

## Direct properties — rule_list.rules.asn_matcher / 79c511cf8b39 / 3

- [asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-3b943ed3246aed981004f8af5d4f517ed45f0781587caa223e33ebcd891c3e1a): complete subsection reference.

<a id="canonical-25a3eaca06ec0950d64a012496a20a529a90554cce986785e4191b3cf3e94c55"></a>

## Next pages — rule_list.rules.asn_matcher / 79c511cf8b39 / 4

- [rule_list.rules.asn_matcher.asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-3b943ed3246aed981004f8af5d4f517ed45f0781587caa223e33ebcd891c3e1a)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-3b943ed3246aed981004f8af5d4f517ed45f0781587caa223e33ebcd891c3e1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-919e22f2cc97a61a4fc3d45f97884d72e2e8d209547db5c846bb31ecec69b84f"></a>

## rule_list.rules.asn_matcher.asn_sets — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-9b0d9fe5b21e6df25b92fb28e8120d54daeaadfe89861528804fbf9956bb2021)
- rule_list.rules.asn_matcher.asn_sets

<a id="canonical-f1414d23b2696d9900dfd0ab572b9d2768721e13f8e325ddc62cb7ee9338e08c"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-891f775486ba4baed92a28164474ec5e46b038a61be0d0128261e74310fd198e"></a>

## Direct properties — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 3

<a id="canonical-509aab550b0c2d9ea350d0f660a52d39c7328ebc98d5f1170b3ea05351d7f2e5"></a>

<a id="canonical-f42eb52b1d1d84797ddfbf5f376e99d49f3bd8ac2a906d2b6b92abc7687789c6"></a>

## kind property — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 4

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

<a id="canonical-1e56d9ba38157c4137e53fc9f73872543e7e204b0bab3738a0e098ea4acbd228"></a>

<a id="canonical-4fb79ad0d31afba88f2cdc0db67f5a9d50d2766b4c4c548ddbc1880513d66e85"></a>

## name property — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 5

Type: `"string"`. Optional.

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

<a id="canonical-9f8bb2906147e4ae12debfd31f71528c4ee38e7f1b84f075e07b795cc444c82d"></a>

<a id="canonical-c2963709b03083112c6ace71755188acbe9da5611ef20d702faec03b8b066eaa"></a>

## namespace property — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 6

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

<a id="canonical-f3e531e3eaabf86fcdf6bb5c7cfdbd5b700eaaf4a935a3d077fd3bdfcbec31fb"></a>

<a id="canonical-68ea609391019fe4525582edf88550db2861f8503a45312c2e6abffac69bdb6d"></a>

## tenant property — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 7

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

<a id="canonical-dadc8df59fcaf3664846f4e144720b529aae298b9ac82d4912c0dfd262cffe83"></a>

<a id="canonical-823443a695258ae2183b1f74c2fee9024e50160319672ae9d563a0f411882327"></a>

## uid property — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 8

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

<a id="canonical-9ae1578b262815abde5aa72045cf255037e9d2175286de3ca809c31c6318c09b"></a>

## Next pages — rule_list.rules.asn_matcher.asn_sets / 69f4c1793f19 / 9

- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-9b0d9fe5b21e6df25b92fb28e8120d54daeaadfe89861528804fbf9956bb2021)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-3ebffd5e0fb62cc862c46acbd943c541caf87fec78c19b2bdf15f333cc0c2824"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a82ff11a4ab9bfa9f2d97a8076105616b94b922ade20f22ac10e6ef6fe772ec"></a>

## rule_list.rules.geo_location_label_selector — rule_list.rules.geo_location_label_selector / ef67752e25ba / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.geo_location_label_selector

<a id="canonical-db55335b29e3cb971d42f76b55400c88c5f4c4b7b2aaa5063d64a6e047067e61"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
geo_location_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-12ed46b305e3ad5a21c0f7047549c41a1566ae241f7e36bfd7d7353c5d73f2eb"></a>

## Direct properties — rule_list.rules.geo_location_label_selector / ef67752e25ba / 3

<a id="canonical-7e011c2d7d65c804f7bd8e31b1c8db35fca9c9a0f6fef7b174e82f808ce2c51b"></a>

<a id="canonical-fd0e4c1c58a475f729498f0e68f69ec20dc42e1373071a23b52b522e4aeef5cd"></a>

## expressions property — rule_list.rules.geo_location_label_selector / ef67752e25ba / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-24eb0194ccd00bcc7137735782e216dc88ccda00e1de7ebcc4b24278546fa005"></a>

## Next pages — rule_list.rules.geo_location_label_selector / ef67752e25ba / 5

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-feccbd9e566d993c7883d7964810bb448b71dd4177b045564885e67302d74b94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a735687a43db84481f11e66f89a372ea3640e180d69cd71a8b6617e461aff10"></a>

## rule_list.rules.geo_location_set — rule_list.rules.geo_location_set / 3830cd29ee4f / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.geo_location_set

<a id="canonical-dbf13acab9ba22936a1856977ce734b73c37e7197e8c61ad7ed225eb19384d12"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
geo_location_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-8912ff6b7371733fec33e0d982ca9d96a7203946104425085f39b32adc4390c8"></a>

## Direct properties — rule_list.rules.geo_location_set / 3830cd29ee4f / 3

<a id="canonical-b75ce5ac848c03175d6fc61e490514e9b84af9ddaf1a18708825667852515f01"></a>

<a id="canonical-e1ada0aa6a8d6dfb01682e5847124fc2823f75900028c9c085481f0ae728b5d7"></a>

## name property — rule_list.rules.geo_location_set / 3830cd29ee4f / 4

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

<a id="canonical-aa6677edda6ffd7e56e9ed55648b29d412002afc9c97c95c53835951a58475cb"></a>

<a id="canonical-a56fa5db5f7b326f2149d0d3e39078a7bf4a6c68be348d36025902e8a17fa960"></a>

## namespace property — rule_list.rules.geo_location_set / 3830cd29ee4f / 5

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

<a id="canonical-b2be701b773d0ed3c47c53580435ee9243ddd0d0717e69c1fd0849c6080b091e"></a>

<a id="canonical-0365dceb574fbd38db6e671cb693a994bfed54c55d7ec996c4a5d6a023637d88"></a>

## tenant property — rule_list.rules.geo_location_set / 3830cd29ee4f / 6

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

<a id="canonical-653e7607df3f91cd956bee8b0bee0b6eac186815713c0edeceb1e6c636923282"></a>

## Next pages — rule_list.rules.geo_location_set / 3830cd29ee4f / 7

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-afd66fcdb230a81581a465d7aadf0341288526ec465c09aad54f941c4550abf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c69762d4b405aedeb90d4f9abe350f37bb007f482cdbfcbd324164c038d5c367"></a>

## rule_list.rules.ip_prefix_list — rule_list.rules.ip_prefix_list / e65b1b7445f2 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.ip_prefix_list

<a id="canonical-41341d04982ae49f2dd77d2a52f6174aab37dc97200685de5cfd9b6cbe3a05cc"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-359a625554c26e7b02986757b04ac77debd65084ca0f65808608b7344919361d"></a>

## Direct properties — rule_list.rules.ip_prefix_list / e65b1b7445f2 / 3

<a id="canonical-2a2dc22cf9befc076765bc7be2bebbddee5238c71bdaccbe460c1713e1782bba"></a>

<a id="canonical-fd527ef174f4562c7e6aecbf4faa9d151111e99f8770b183b9f1cb4ec5cca45f"></a>

## invert_match property — rule_list.rules.ip_prefix_list / e65b1b7445f2 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-e863496bc7abe1111b9d228737b92136da36ceba3cefde7db959a1807a5c3fac"></a>

<a id="canonical-de9011b001f01fd154ea0793c5d31274f41b809311a83b6fbefa5bf46993d1f8"></a>

## ip_prefixes property — rule_list.rules.ip_prefix_list / e65b1b7445f2 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-619ad731514e84949ace6c12c80dc45b55341131a902616837d841305ce33294"></a>

## Next pages — rule_list.rules.ip_prefix_list / e65b1b7445f2 / 6

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-494d9c08d7d1c46099f27c5c5b012b09b4d65f820e78387576adab90ba62e3fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e14a621754645267854a42beda0660388abfa8a4da433dedd71acdbf38c2144a"></a>

## rule_list.rules.ip_prefix_set — rule_list.rules.ip_prefix_set / 949ff0acd8ec / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.ip_prefix_set

<a id="canonical-228731d8ef66958a051c95c3a38de3baa45a7a659a3e7eb32d1ad6ebc1faca93"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c1bb84877b6e0e083b962573fb6c5304bfe452578737a94fe670dd8b87bd4b3"></a>

## Direct properties — rule_list.rules.ip_prefix_set / 949ff0acd8ec / 3

<a id="canonical-de95805c4a8b121b375b68cd668c3ed6ac82b8599b3237775a8cdadfad02c89a"></a>

<a id="canonical-cd4e82b034b0fc6af6dc86db1eab3cd40b14e424d325e257f5c0a257393fbfa9"></a>

## invert_matcher property — rule_list.rules.ip_prefix_set / 949ff0acd8ec / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-25f62f24d6bb636174bb3a64de2a6a84f8834b578507bfe23d561fb40919726e): complete subsection reference.

<a id="canonical-3c8377cbe7ed745c4b7914c33d66cb24c1ac1b609a3cabd24d6de4e0f4328b29"></a>

## Next pages — rule_list.rules.ip_prefix_set / 949ff0acd8ec / 5

- [rule_list.rules.ip_prefix_set.prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-25f62f24d6bb636174bb3a64de2a6a84f8834b578507bfe23d561fb40919726e)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-25f62f24d6bb636174bb3a64de2a6a84f8834b578507bfe23d561fb40919726e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3716bdab69352846f9dcc4cca7e78cc1653c80eba067087f1a6221093605364"></a>

## rule_list.rules.ip_prefix_set.prefix_sets — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-494d9c08d7d1c46099f27c5c5b012b09b4d65f820e78387576adab90ba62e3fd)
- rule_list.rules.ip_prefix_set.prefix_sets

<a id="canonical-6693edc51b380eab83e71414e40b26f1cd09de3d402d97b5b5cd4709213336d0"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-8757a1654a47571fcef5b26a946b5f9cf84b8eefdae1140a854d9bcbccfa8ebd"></a>

## Direct properties — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 3

<a id="canonical-c0fd46901b0c336d708d4aa6924f374e0ae0fb4f0c610ab7dc3a1dd50e3bd5ca"></a>

<a id="canonical-329669e291f82e326a475d11b808577d994ff4aab9cdfc6a166afc8b4db6b2ea"></a>

## kind property — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 4

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

<a id="canonical-afa5cdeda5c8c1c635bdf338eba00ee6703d9ffc8b001294eb206d245d71ce55"></a>

<a id="canonical-f9ee1edb56ffc631031ed8b7203006fba79c46cc04b1b3893e9c3e052ff2d9f6"></a>

## name property — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 5

Type: `"string"`. Optional.

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

<a id="canonical-d6ed60d8fec4a731e4c9987957165cc9efa2b0dc4213f8e0935ad12e915ed7e2"></a>

<a id="canonical-c68d96fc129af4c05b6e950de2a00a8dcc3e8c78527eac0fcc9c9c8e3588245d"></a>

## namespace property — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 6

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

<a id="canonical-129a71e2fa38357fd51c2c0769eac96afa80a9e0d4dbd98891565a38ec5d02c7"></a>

<a id="canonical-1f7849e702d3eb016bbc83f74c04f613faeda3a7c3c1a5caa1de13e40eae91a1"></a>

## tenant property — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 7

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

<a id="canonical-7d73206f1617e6d33f1c79869441f347d617d311e9d5d8d278f0d48b9375982c"></a>

<a id="canonical-83f6067201d1e28b843ae83aa2546b585757b9350fed57bfbe5ca08ec70d86b7"></a>

## uid property — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 8

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

<a id="canonical-78599054fe657f8d6b879bf35c872c3671a642a034f88b136959b05d2a65f2af"></a>

## Next pages — rule_list.rules.ip_prefix_set.prefix_sets / 4cb3cd8e04e1 / 9

- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-494d9c08d7d1c46099f27c5c5b012b09b4d65f820e78387576adab90ba62e3fd)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-44f4ea735f2b8602f2c096c96af203ceb25b8858714a8dabc8c0ad7ca8129209"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c70a14ff2de79be37ba661098002711a50f6848b918d7df824fb89e63e703db"></a>

## rule_list.rules.pool — rule_list.rules.pool / 8edb7df91d96 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-db045313cc27480c97f5211e3373570534cff29f994a6c8d84c49215d690e79b)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- rule_list.rules.pool

<a id="canonical-35d606fa8d7100d1246630f56b3dfc6e95ddd8eb0e7c62749d33b2ffb01a304a"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a3f848dca470dd37ca92536c7fac9022e827ff3661a20b804840109ac3b373e"></a>

## Direct properties — rule_list.rules.pool / 8edb7df91d96 / 3

<a id="canonical-d47de1c496baaf1023ba658ec796baffd0858058305f5e1622f39ab045e479f2"></a>

<a id="canonical-0c50d49e6ef2aff2ed089dec58d0925b56d0984a04aedf79dfb8600325bf9b56"></a>

## name property — rule_list.rules.pool / 8edb7df91d96 / 4

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

<a id="canonical-4aaf4904d0ce07863d906518f62770a2cfe13026abcdbf8b07e80749421a07d8"></a>

<a id="canonical-c46d5eea706ef790f45e844257fa3e675b332e1ad3e54171e46cc00c0c7d2c9e"></a>

## namespace property — rule_list.rules.pool / 8edb7df91d96 / 5

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

<a id="canonical-385e6f7fa963843bb764c34a2782fb70b76415e5ac7a1134a272db18d05aede7"></a>

<a id="canonical-084aeaf108b3318a42ce4dbabd0806366f1335c9193e34f8fd8324de07d3289a"></a>

## tenant property — rule_list.rules.pool / 8edb7df91d96 / 6

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

<a id="canonical-22d382cbc4208e6a0b16523a84b9309fe2c3c528f090bb98bc47dfd00d81feea"></a>

## Next pages — rule_list.rules.pool / 8edb7df91d96 / 7

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-b978a7cc7e76986d9741a716943224f099e851f2255926ad5a4719fcab084419)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)

<a id="canonical-c8451a107d6f68291ff2f63e446d22ac295d5a644c522b0fb13c71b284a4dbc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fb68ca6160f1e19dfd2bdab31e1a97cae440acee289207f70a1ab24be12ef20"></a>

## timeouts — timeouts / adf28e3d11f0 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- timeouts

<a id="canonical-de31dce4f63da5424332ab0ef92d64f7c784a2e5b26d38a99070145d116c0f15"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf58c4bc20c6f8f28a07849caba9cc47792538cd99628fb40039ef350d144fa4"></a>

## Direct properties — timeouts / adf28e3d11f0 / 3

<a id="canonical-05b55872e4476bcc2c6d5f21836575bf451507690b2112a3824e1feefe575590"></a>

<a id="canonical-628375519a01df9e36bfaf8d64f842121a79bd865d20a02242101458927ec803"></a>

## create property — timeouts / adf28e3d11f0 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2838e77b0a0e115b786553838097a11f43c24f0547569bf412e0da6c8b552c71"></a>

<a id="canonical-d05e8288545b8761574ab29b1c91c2f9bd3698272ea92073af6cabd868208fb0"></a>

## delete property — timeouts / adf28e3d11f0 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-34d3ecc0ff4fb08f7570a922a133856f28aa240f3c40857b528c2477f127c569"></a>

<a id="canonical-8a3512d6c2f9541b625e5a8ad3559bdbac880f565f172c4f970008871e767576"></a>

## read property — timeouts / adf28e3d11f0 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-22fb306612ba42b47b3b8f2e11daadc3e87381d0c8b9decdd256d31ffd35ed69"></a>

<a id="canonical-3991c8819a77d44437c5d4b8447fd4c342f02580faf9284fd05af6f4be30811c"></a>

## update property — timeouts / adf28e3d11f0 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ce87e713ddb33e2f32cd93acc504b315fec50feb49ec5b57e96887d4667f4875"></a>

## Next pages — timeouts / adf28e3d11f0 / 8

- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-f932e202054f58f09900a6d1e9d671ba4fcf260d4501fc9af69356d4abde8d0d)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0fcdaf880e64bfc375d92833e3901f826bf9b25df73e37402ea88ec899069baf)
