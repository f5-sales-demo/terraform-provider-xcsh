---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2fddb19492cc1b85bca0be62943a1aecaad08058d9f75bbef7d90e1a4a07056c"></a>

## api_rate_limit.ip_allowed_list — api_rate_limit.ip_allowed_list / 556d9da21a15 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.ip_allowed_list

<a id="canonical-800223f0d09797bd195d04493eb8ca3ac6b58f002d2cef43b334ce50eb60e640"></a>

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

<a id="canonical-538d44f2a8a732329946b46d818942672a7fdce4650d5b3ebd98111a6c6601f8"></a>

## Direct properties — api_rate_limit.ip_allowed_list / 556d9da21a15 / 3

<a id="canonical-cb31281a92e7a5a35fd3384e0e908f9d8f358583ef7f831c20985ff43027c3eb"></a>

<a id="canonical-e1beeb0db919798603c8a8ec22207c37204e71b3bf33e28b9bcbbd7c8366b15c"></a>

## prefixes property — api_rate_limit.ip_allowed_list / 556d9da21a15 / 4

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

<a id="canonical-82f6a53242cd44871d86435663d7081ae8cee3a2f63d011e418c4217d152746c"></a>

## Next pages — api_rate_limit.ip_allowed_list / 556d9da21a15 / 5

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c6953626a1af97898911862c167c8d2f988e255dd59e4089d14bb78c6831fb36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f86f60d156c6fbd6a3c8bf16d387e3c0c583edba45a7dae37129fce83ea3db91"></a>

## api_rate_limit.no_ip_allowed_list — api_rate_limit.no_ip_allowed_list / 2d451f303882 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-9c1fa3ef36007092adfbd452068ae1fef70f64bafe39de819ecba02c0ae79347"></a>

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

<a id="canonical-6ad5e36e5110220a6c7af1bd0212dbea2a6e708b215bd6866582580c03513352"></a>

## Direct properties — api_rate_limit.no_ip_allowed_list / 2d451f303882 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f06bc85eeb271fc2a2f99a6f3f61ae44d4336c18242b91d0e0a3d25c14e90b9"></a>

## Next pages — api_rate_limit.no_ip_allowed_list / 2d451f303882 / 4

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8196eff4b19224fc5a9899496f98da5bd238d119c5a76482cf49d6b91c6e4717"></a>

## api_rate_limit.server_url_rules — api_rate_limit.server_url_rules / 4119eed9521e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.server_url_rules

<a id="canonical-25c6b1cefdbc7e4ba7c75ca41696e7dab09b4ca01a7521c7bddffacc1a8ed60d"></a>

Type: `"list"`. Computed.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-747d682d97acfe5331dac2533e4920508b21b33c8caca1eae150cc817d24e2a8"></a>

## Direct properties — api_rate_limit.server_url_rules / 4119eed9521e / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-008.md#canonical-3f3a377a00cead66152632520041a0ef1dd1c87660ef03542004fa5ac5f498d1): complete subsection reference.

<a id="canonical-8e35d3fe756e12d7eee461cead3518103a4e5eaf2e2ca921bbe9f037f1b53cf6"></a>

<a id="canonical-aab0c57dfed0345a6803e51b76edf03e0ba6b537f435f9d1ba7bcd736ca2cdd0"></a>

## api_group property — api_rate_limit.server_url_rules / 4119eed9521e / 4

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-b0cbb3790999e88337d060f83b93a20820fe850eca0d89b6184f064b9f0074ba"></a>

<a id="canonical-46bd09c149e3d7e1d2d24364e6567ee479bd9d83d37d9072b57898761934c823"></a>

## base_path property — api_rate_limit.server_url_rules / 4119eed9521e / 5

Type: `"string"`. Computed.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-2c82de0a0a953dc8f74a70479d2c18699b9ea56a95b437fa9fb526def255185a): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377): complete subsection reference.

<a id="canonical-fad88cd3307c560d34f39bc34008efa4180f13fea2255e08c1cf9db8502af104"></a>

<a id="canonical-d8e2c874f8be4987fd533a2cb034aef85e407f3cac7bcd8c211841e3c65b2606"></a>

## specific_domain property — api_rate_limit.server_url_rules / 4119eed9521e / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-6a32700310f47c01710c094d7d3b95102755201cdd9e4065e1431bdf63477ffb"></a>

## Next pages — api_rate_limit.server_url_rules / 4119eed9521e / 7

- [api_rate_limit.server_url_rules.any_domain](data-sources--http_loadbalancer--reference--group-008.md#canonical-3f3a377a00cead66152632520041a0ef1dd1c87660ef03542004fa5ac5f498d1)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c)
- [api_rate_limit.server_url_rules.ref_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-2c82de0a0a953dc8f74a70479d2c18699b9ea56a95b437fa9fb526def255185a)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3f3a377a00cead66152632520041a0ef1dd1c87660ef03542004fa5ac5f498d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccd82f7dd78c80c9f319cb57a2c60e10e054f1ff04ad2d29c66f957d6b9a1586"></a>

## api_rate_limit.server_url_rules.any_domain — api_rate_limit.server_url_rules.any_domain / 2d141690e185 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-1d6d25fab5b998024a96252a633a8e49e6da41cf88a2fe6554638233cb5846d1"></a>

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

<a id="canonical-ddaa3fb54e213112ea7792265de4c2cb893fd570b8e3b6dfe6c7698d1727a989"></a>

## Direct properties — api_rate_limit.server_url_rules.any_domain / 2d141690e185 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39eb123f5d3916d2417bb746801b7bee9844e80db23029569fa64e04147a4ef5"></a>

## Next pages — api_rate_limit.server_url_rules.any_domain / 2d141690e185 / 4

- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77cd671abaa3e61bb1466f8833f24c4aae333b1f0ba11dcf40e0c52f312493c3"></a>

## api_rate_limit.server_url_rules.client_matcher — api_rate_limit.server_url_rules.client_matcher / adf4e7733c18 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-4a3af1e3406b8be1ed9e361d8576b18440f5c2a4bdbfa05d450e1b488dd7687d"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-98df93527f94a14a56dcc2fcab86ca01670848ab5dd11fbc0828af242527e330"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher / adf4e7733c18 / 3

- [any_client](data-sources--http_loadbalancer--reference--group-008.md#canonical-ba6933190be292dfa23ef13b79d129201930c9dad48afd0b4cfe470b63ae4bc9): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-008.md#canonical-1c00971155cf9751cd3584e8a2243b550b9c6eb347e64de027270fd7b105a4b6): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-496f6a6176ab15722f359fff56f8df4bb41db0c9b9942016575363b11a6d1cdc): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-016841196111bd2298d975ff0b9238c66f2d44606ad6db9a013d0d494da75d73): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-008.md#canonical-d592b7bb19182d882850185bc2a1076e104e7f0a7c11c3f0b71256a2757d8eb2): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-dae82020d8e03bbb3212486e6fb939f5d594d59528b81451bd2e5b8a270582c4): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-35a935b3acf4c10937f6fe15dec67923491bd65351d1402ad6718ecdb06eca48): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-d85ed68573f58ddc3d6c283753579d87dcf4c6ddca8c99d10d335eae79fb4346): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-2498b07094c9c645d56bee1748d3f3edaf4f48a62075e03011a6df62969402b7): complete subsection reference.

<a id="canonical-0bf908cb254b60474d022ba8c7657f617f5d4385c8b7d3699e660340c1b83cea"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher / adf4e7733c18 / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](data-sources--http_loadbalancer--reference--group-008.md#canonical-ba6933190be292dfa23ef13b79d129201930c9dad48afd0b4cfe470b63ae4bc9)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](data-sources--http_loadbalancer--reference--group-008.md#canonical-1c00971155cf9751cd3584e8a2243b550b9c6eb347e64de027270fd7b105a4b6)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-496f6a6176ab15722f359fff56f8df4bb41db0c9b9942016575363b11a6d1cdc)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-016841196111bd2298d975ff0b9238c66f2d44606ad6db9a013d0d494da75d73)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](data-sources--http_loadbalancer--reference--group-008.md#canonical-d592b7bb19182d882850185bc2a1076e104e7f0a7c11c3f0b71256a2757d8eb2)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-dae82020d8e03bbb3212486e6fb939f5d594d59528b81451bd2e5b8a270582c4)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-35a935b3acf4c10937f6fe15dec67923491bd65351d1402ad6718ecdb06eca48)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-d85ed68573f58ddc3d6c283753579d87dcf4c6ddca8c99d10d335eae79fb4346)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-2498b07094c9c645d56bee1748d3f3edaf4f48a62075e03011a6df62969402b7)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ba6933190be292dfa23ef13b79d129201930c9dad48afd0b4cfe470b63ae4bc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3180fd0867a29096b6ccb689e0ccaa6ae5e4a34455b6a1ed7ac44de60ea9a9c1"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — api_rate_limit.server_url_rules.client_matcher.any_client / 51328de51ad6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-f80ccef8668ecb66e65b4eaef43d57d7aec385d83c40b105b39de1d0ec98012d"></a>

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

<a id="canonical-358f433981e069208f4ad7fd0e97ee2425e050a2d929e87c7d840474782a7743"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_client / 51328de51ad6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24909fb2da722671dc87eb9b1a4db2d60fb11b16dbfbde3ff0a35421fb4352e4"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_client / 51328de51ad6 / 4

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1c00971155cf9751cd3584e8a2243b550b9c6eb347e64de027270fd7b105a4b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b3f1e6b554738242b9d48b583410f30164779d5d94e526e912879365fb4cba7"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — api_rate_limit.server_url_rules.client_matcher.any_ip / 1a72cafa0d91 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-43a25be7cdbcfb2d2469863a33d732f83b3fac5e6519018332e346021ae23add"></a>

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

<a id="canonical-7ed9492647889d65b93accb602b860b95d0132b32049e42f7408532cd563d9e5"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_ip / 1a72cafa0d91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b12cdd2dd8e28773bf4591a72b521aef74eb2d20f5f6a95a6179205898d78349"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_ip / 1a72cafa0d91 / 4

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-496f6a6176ab15722f359fff56f8df4bb41db0c9b9942016575363b11a6d1cdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fed707edc960303b7662ad2ced4523d8e6cd5a22d1069cde3fbcd4afc0e82da"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — api_rate_limit.server_url_rules.client_matcher.asn_list / 59d404965af4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-ac42f8e758c045800f3aabcb690c209fffeb2df6ef2cf05c3df706df601f52de"></a>

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

<a id="canonical-0f130c7ecc4e3958eafe32e40bab3b54a9b97eb3afa3c5558da2368ee4f85542"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_list / 59d404965af4 / 3

<a id="canonical-91ec9ad2b318694093b5c95c9229411c432a1bd29eb8e0c486bd4ec23c177874"></a>

<a id="canonical-6c951ab9632ec5ed09733476e72333f6a04df11114f4f6497ffc11ca1315e766"></a>

## as_numbers property — api_rate_limit.server_url_rules.client_matcher.asn_list / 59d404965af4 / 4

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

<a id="canonical-20767825b2e7737627518a364ffee6287c91cb44597a78dacfcb8949bd415b06"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_list / 59d404965af4 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-016841196111bd2298d975ff0b9238c66f2d44606ad6db9a013d0d494da75d73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1b94bd0f49aff45803c00d868ad4e67e72696eab040e09e3bb5aad2b86529da"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 2a340a214444 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-911d8556da974d6a53c3c5ff6fee4a05bb90b816de31e3db8b1e3687575d098f"></a>

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

<a id="canonical-13b58f7e116a9a8b3d4247494ce76598b5f56ad53875f8107656d8af2dd73cd2"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 2a340a214444 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-008.md#canonical-9f080107b80a9218d58bd209235f96cbec58f53a00050220ec48c5e34dbc8ef8): complete subsection reference.

<a id="canonical-98e4451aff80f2cbe0ceb6cab4f42c6def54ab7a628237155ae5c261a02f53f8"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 2a340a214444 / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-008.md#canonical-9f080107b80a9218d58bd209235f96cbec58f53a00050220ec48c5e34dbc8ef8)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9f080107b80a9218d58bd209235f96cbec58f53a00050220ec48c5e34dbc8ef8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a52cedd4aea6a77a34350490498ce5c7ee1f587fedc355bfc39258597c63e3f"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-016841196111bd2298d975ff0b9238c66f2d44606ad6db9a013d0d494da75d73)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-23a16b9a20a7f986d9f3466e31a81d6df3bd63335786720103b0e3fda9ddb9a3"></a>

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

<a id="canonical-023838791530a169a87327265073e39da68508a13f1e27c6abe337406d1e8c93"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 3

<a id="canonical-c72c3a370d3712fda8ea9f09f80c373ce68e01a424f58b3a63526f99af6f6c3d"></a>

<a id="canonical-45084ef6fd5c368538f7065bb763107d81069ddc76f3a0d108452c1ef5b89766"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 4

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

<a id="canonical-2e5a4965ac669d86ba7eea932ad5fa7676e864cddaddcd1f80508ab4dc108716"></a>

<a id="canonical-90b2c4f5e439902020fb6eb36ec3374879ac9f3b77875d9692cd429f91bcc98f"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 5

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

<a id="canonical-ac77226090f68af6e73e3aa417d9bf03537512ac59287eec4cffa85ce227525f"></a>

<a id="canonical-fbf75f6da3f2075705471589a0cf938ee98bb4ecbf31caf5598f7c78e3b523ec"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 6

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

<a id="canonical-945efcd0dc7678ec36afc7e64aee1ea46f90e3b27d87dab20f6b9fc6735e05f1"></a>

<a id="canonical-e926ea19b2a1b90f08c5eff3f16a5450eed7282b87317dc5b62bcd4174cc61ae"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 7

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

<a id="canonical-808415290ad7f8f91eb33511a5cd6f47e4c6bc61f0dad67f4d42643e8d0a6606"></a>

<a id="canonical-99146cbd45e4bb88b0c776cef50cfcb33f9a4ad6cfac2fd4c2d17460fb93a8d4"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 8

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

<a id="canonical-7480a5d9ae9cd5f9bd52b2458dc3fc048df3cbb1f72ee8c3ab534010e123277f"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 524aac5cca89 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-016841196111bd2298d975ff0b9238c66f2d44606ad6db9a013d0d494da75d73)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d592b7bb19182d882850185bc2a1076e104e7f0a7c11c3f0b71256a2757d8eb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6519c7aeb16d6566681af0b6cdeffd91845d4239cf349820e1f0442329759f2"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — api_rate_limit.server_url_rules.client_matcher.client_selector / 11ce52ff4032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-a961cc42528d352ed0cb134e65014abe2e0d23cd05e97f7ab0fc346b34bc2259"></a>

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

<a id="canonical-316d7ef0a07cd9c76c5fef643a04244e6e536f4ce4855bbb4d398a1398a1f0df"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.client_selector / 11ce52ff4032 / 3

<a id="canonical-1238537f2d0da949837b6175d3735599b738966700e7317c10b0c5ac38cfc020"></a>

<a id="canonical-a12ecd3d5444288d2299f01fddbd3432a77eb7518742da60b5e85e6ffa351fb6"></a>

## expressions property — api_rate_limit.server_url_rules.client_matcher.client_selector / 11ce52ff4032 / 4

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

<a id="canonical-cf5bed1858e11cf7dd5fa662a8e47c927b3a1a41044e647cf83e3e769930b96f"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.client_selector / 11ce52ff4032 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dae82020d8e03bbb3212486e6fb939f5d594d59528b81451bd2e5b8a270582c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f0168f8aba4357e5bccb0af982b6c749bde460a4859212fa87b9cb46fc507d1"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — api_rate_limit.server_url_rules.client_matcher.ip_matcher / f7b3e8847637 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-8346abab1baa8a8dbef3d1fe24dce79803ba0591184660a641b74a3a0ccc6192"></a>

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

<a id="canonical-83f6046e65eed1a3c4b9fb82ef8d77a8d29a70663328506d46545ac8efbbb5e4"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher / f7b3e8847637 / 3

<a id="canonical-efe17876148047fd3757bd6a9f316068298c8b6f25a954572251471fab0d796a"></a>

<a id="canonical-e71d12a9cc470114b617feef291b28899c8d54ebaa4f7b8ac21be9b9ac08fd00"></a>

## invert_matcher property — api_rate_limit.server_url_rules.client_matcher.ip_matcher / f7b3e8847637 / 4

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-008.md#canonical-37af8ca7a82f805e6ee69eb55bbc74de327b13ff9fb897acb37c0a856f0be12a): complete subsection reference.

<a id="canonical-8c43bef7600879ac4188f51ae2c7b9c8ef2483a10f9e80e76011bdc084ff8f12"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher / f7b3e8847637 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-008.md#canonical-37af8ca7a82f805e6ee69eb55bbc74de327b13ff9fb897acb37c0a856f0be12a)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-37af8ca7a82f805e6ee69eb55bbc74de327b13ff9fb897acb37c0a856f0be12a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7b7197393e4f8f849b72965e605a1bd707ea3761a1b3c99e3522fea2639dc9a"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-dae82020d8e03bbb3212486e6fb939f5d594d59528b81451bd2e5b8a270582c4)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-6f6de7f57e4bb619661ab5a9f44630013573728206f82f0889d94b2c1d5ff795"></a>

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

<a id="canonical-f5fca1ce074ba4d04f50b9f0ab90956fb78bcd1f8c93a3f4ad8f8e6e08a8ebe9"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 3

<a id="canonical-d2d147142db87a3476ba774b581264375935efc4394e1f6fd87d7e5d21dc6207"></a>

<a id="canonical-91e3eb027d79bf00684dc2479db4a5599239cc68bba05432dffde2e9e57f89ae"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 4

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

<a id="canonical-cb5606dd05baf566807d20cf71e08a37f498df85ee82876167efa70fb211ebae"></a>

<a id="canonical-f75cac3ad26559ff45e510b7c0d95ab3810423740d4abe887ae52d56d71af374"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 5

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

<a id="canonical-e408c5279b823ca11fbdaf0dcfbf72b1ed7dff9f0ece4f2abedbc7779391f714"></a>

<a id="canonical-e9d435d75e0fd32b4e2f250465b130a8769cfd507cfc03a5abcd58487cddd64e"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 6

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

<a id="canonical-06269b6b07f9ba5c274532972e5cc39ffa6194c32875b7fd0e417db5f4c870d8"></a>

<a id="canonical-b5945b2f197d048e2a333d0eb2365c8dbcaec9b70189aab562674d15aa0fa32a"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 7

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

<a id="canonical-683502552094bb3e729310051f34af7eeb50727d6a5cc4f6706a2a52e268ffe9"></a>

<a id="canonical-fdc7ab9df8fec473f53a3b6f48fd2c2de1ed1f5ad3c768cd4ef6df025bbb52cc"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 8

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

<a id="canonical-b2dbf3fbb9e3a4050e36b384e480e3d8a2495e95fd75835ec8776666a22ecb1e"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 6cdcccb32d7f / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-dae82020d8e03bbb3212486e6fb939f5d594d59528b81451bd2e5b8a270582c4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-35a935b3acf4c10937f6fe15dec67923491bd65351d1402ad6718ecdb06eca48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31082fa048ee1a746fe0ab0d5279795f901e4b15af0254efea36277ab4db0db9"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / 7037499488e0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-39b6311e55ab95ec56c4c922fcc07280678271dd1cb92b4789ceaf0763bbd951"></a>

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

<a id="canonical-8e85e9d46af86eaf942647c1f9d89f609f41edc0c720054e25188cf3229194f8"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / 7037499488e0 / 3

<a id="canonical-f10fd32846adc0bad8ab5602266e613a0eccbf3595b8c0236d8383affb24be43"></a>

<a id="canonical-dd49b085d6e3fd2b4d08efa1c407e9214a571321c624be91008959c23f8f7759"></a>

## invert_match property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / 7037499488e0 / 4

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

<a id="canonical-eb4261beeb03974ac48fdc25196a67ae68323acd62f7597ff070a2b18998731d"></a>

<a id="canonical-7d169c38294c9c232d06c4fdfc1dd629e65297404d23aad5a064d5a16b308b4e"></a>

## ip_prefixes property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / 7037499488e0 / 5

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

<a id="canonical-5970735ebd29c26d187e51d04c92dfa25f5de859709cc59ad008a6f299c961b3"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / 7037499488e0 / 6

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d85ed68573f58ddc3d6c283753579d87dcf4c6ddca8c99d10d335eae79fb4346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac778a44a3f791b667e53d0bd58eec3e8cf517f3344130c5b7f0515a2b0456c4"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / c71c54155f4d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-555dad6e2de3156ee61aaeb54899c3942cee1eaad3982f207438d0f901592c72"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-cace8936e7fa0aaa8cd219b36236a0389bee724b06014a82f9a1f23398702bf5"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / c71c54155f4d / 3

<a id="canonical-5cea809ab3d6a24cc74681c6e34b1b952e4c5be4d8745f5c4728b755ddeec890"></a>

<a id="canonical-24f3cd12d7e7d3ea7ca2ea1c69d21ca03f7a3d794a7d1291534f4971ce2f0ac9"></a>

## ip_threat_categories property — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / c71c54155f4d / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5a74c4722a6851c63bdcc7ace45bb2dec314f2fd4983ac9bdaf183f3d582adf3"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / c71c54155f4d / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2498b07094c9c645d56bee1748d3f3edaf4f48a62075e03011a6df62969402b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b49e6d49f48c51cfd39e38bd1035730bf7534e31d21a7fc02000c34306dfff09"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-dd38ea58a17c3b1a64802e7490843277500588f7bf6c78c8d4fddadbd9ff67b3"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-e79f7a535e3c832bd59de4e51d8974b11dd7a4630088dadba1b108aa55463c58"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 3

<a id="canonical-83ad5024fd9f00f5b3612c2e4114aba6e95e3d2b096d01721d6a03b8486cf021"></a>

<a id="canonical-4ad904273dc7739c326ec3487c00b98070ac789cc32195e5c9b141029539ef10"></a>

## classes property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-71251b1e7eec7061a2879462501b592e0fff6123a494c1cfe7a1433d8e7b1114"></a>

<a id="canonical-8bfd9ab014e13e27dd72ff6dccdce3362acdbdada7e512df08d9318ae6ca5747"></a>

## exact_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-acd8923ddb51c1023d37968b910cc003e2fc7cc0b020b548e47905c145fef28a"></a>

<a id="canonical-0c1a4445514119e7d4ff2613fbe9cfd0851e78f8a234d7550ac215a9e988017c"></a>

## excluded_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-401b08a156953766590937c06caced8a9a993e99a9f53548788f49b7304f259a"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 8b88974b6d42 / 7

- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-80a1e5bd923e052599779e1f96808497df34db27f474f690a8e503a4840338b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4443f712a8ca570f281a714a4184a8c493b27820d86abb6b223eb74d72698c44"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — api_rate_limit.server_url_rules.inline_rate_limiter / 8942fbe95005 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-e99aa9c26829ea602d428c3be72c49d100872d630cec40bd3d6af9a82b2de79b"></a>

Type: `"single"`. Computed.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-71a32c5b9e97f5734ab1d6f7be55c1fe95ab9be2ba293114c8c5756c42cff60f"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter / 8942fbe95005 / 3

- [ref_user_id](data-sources--http_loadbalancer--reference--group-008.md#canonical-5f0a108075f41492297a42a3d45574eb060d885aa17af12d77ba6e0f35f3a1b1): complete subsection reference.

<a id="canonical-28bbe224bbb1bf9bfde7451c36041b46f40300128f135b322f479c2441e4881f"></a>

<a id="canonical-31a7f9684485ce23410ffff985e37cda333e9204a4cc1c4c3126f972efcdeb49"></a>

## threshold property — api_rate_limit.server_url_rules.inline_rate_limiter / 8942fbe95005 / 4

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-4d4c82680041d660ca5f5556c682710e44f8f6ff5fc1309a3f98ad2440e62b24"></a>

<a id="canonical-79e4669060268955b76ff0156726e23fe68d2bae568acaf7893287f6cf5162b3"></a>

## unit property — api_rate_limit.server_url_rules.inline_rate_limiter / 8942fbe95005 / 5

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-008.md#canonical-90bc06c2761e02b8c056fc83ef15c419232c3b6b4bb3bccfd1424dbe5ff3ee03): complete subsection reference.

<a id="canonical-537ac8c97eee7b5916367e3b6001eeef243a6c13afa1d18534c6613c9e50fea8"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter / 8942fbe95005 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](data-sources--http_loadbalancer--reference--group-008.md#canonical-5f0a108075f41492297a42a3d45574eb060d885aa17af12d77ba6e0f35f3a1b1)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-008.md#canonical-90bc06c2761e02b8c056fc83ef15c419232c3b6b4bb3bccfd1424dbe5ff3ee03)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5f0a108075f41492297a42a3d45574eb060d885aa17af12d77ba6e0f35f3a1b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dfe9941e08937edbfcede4276a6b672789e072ade1fc78cdf5aa6ca77c8b31c"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-814d0112e98f5407aa775d4ac32ae0c92a10b80820e443dd9e06a92b66609b1e"></a>

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

<a id="canonical-7004b671b0af7b32ab032fbceabf6239cd2c46e4eb719e881d3e42ce193cb875"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 3

<a id="canonical-ecd204ccf01e0b4128f0d2e63c0f991f1d26703f4d61548a6bf8e177af5adcd4"></a>

<a id="canonical-d046469784c2c5fc328f6973e9fa250362eaa8861dbd80518c6302506acb5bfe"></a>

## name property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 4

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

<a id="canonical-d5004f5f530d78dfc7b88c88038a9fe62ab6667b76d2e2c897f75db8afde9fc0"></a>

<a id="canonical-ef656e9130a47c67c1c0cd794102f668187601a7f4513b845d27dfc0cd0f67cf"></a>

## namespace property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 5

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

<a id="canonical-6e80405f96f3b4ae32d917f09d139cb02addb530e4b253fd26b3666d89ae0328"></a>

<a id="canonical-d1c33f23cc65ffc43cba3a9d944c832d7a453b16d2c9f4e68f71db46cf005fc7"></a>

## tenant property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 6

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

<a id="canonical-e046f32faa6c734882d49eca59a0f72810e72789bcdb3b0374aac3ab1eec3200"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / d6ecc59c0ec4 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-90bc06c2761e02b8c056fc83ef15c419232c3b6b4bb3bccfd1424dbe5ff3ee03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3abfc264c33a6ea4106fd572fba24a6e9ff8ddf5aa395974a903a7ae44884de6"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / db4b1fbde4e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-925b141a33240d8ed5bf88dc4df42960854693670238c91cd3c98bd75c30522c"></a>

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

<a id="canonical-6e69f9b1a3dcbf9d8b23597679137f84c473b762190b43b58788db9966559385"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / db4b1fbde4e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c455baf9d90d0fb5b7de7aa8665832a13f7bfac42432e37ea97f934211dcad0"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / db4b1fbde4e4 / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-008.md#canonical-3a3f432bca0234ede32148a7a9c478a4d4db48c50bfa8d1075a1c880e126d90c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2c82de0a0a953dc8f74a70479d2c18699b9ea56a95b437fa9fb526def255185a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c40250d4446ab19c8243466190004dda69ad26a5dd10539e61f17e346688b41b"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-4766001a37d6f70be7b9de1006957c9a725785b07e24308c9119ae7e2f35de03"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-aa72219e0014884be8bcf68e837e70bd67dfccde68da1306b9d7a9cf3deee8a2"></a>

## Direct properties — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 3

<a id="canonical-157bedca53dd5b1c98a3ea464b5bcf330b6c602b10d0752a92cfebf551a0e130"></a>

<a id="canonical-19445b2111a94047404523189bf3b46fb27490a9950eb5c7d8ab785e7e2f8d15"></a>

## name property — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 4

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

<a id="canonical-2685daa3a68cfe066369e999fab3b837c67e9accdd86aab09f2e207244d7c1c3"></a>

<a id="canonical-78b922879362b91c4709f41060ed43f32eec8dd6b7021c2220a67fa6f230db9c"></a>

## namespace property — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 5

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

<a id="canonical-a68c00a923376a9e5b46ea5e813760f019fa1eeb2250943e61c0803b26de6ea6"></a>

<a id="canonical-3c5edf6a9dbf632681a9f50d55717b9c66ae1271839c94d36b1cef756e7db223"></a>

## tenant property — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 6

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

<a id="canonical-f1dc51efaf8d9e304123a00c77e14bcc35f7365e957e474ae8824100949cfe37"></a>

## Next pages — api_rate_limit.server_url_rules.ref_rate_limiter / 06b905ab728a / 7

- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b48756052607853e034480f510ae8d30768ebf1fabd1f5d78ae8ab7596d0e0a"></a>

## api_rate_limit.server_url_rules.request_matcher — api_rate_limit.server_url_rules.request_matcher / 4aeca50386aa / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-62bcef76b46990e3cd6be70f52c22ef76f87529d2ae12c3f731a6bcd93bf1eb9"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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

<a id="canonical-620e8e654d2f6afb50ced8e8f12bbdb230c1c437c24809945a24208fb0b93d46"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher / 4aeca50386aa / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182): complete subsection reference.

<a id="canonical-f981d2dbe3d3790582be6f05e2f05d67be5513ded0a1c0eb6c78cdce77322e44"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher / 4aeca50386aa / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a98507b0426837c14bfc3e97c900957110a98cb08757c6b5ed5e9ea766a3466"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 58cfcfee5cb0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-dfd8359f524237921dab858cfbed7cda5dc3c845fee49c51111d6718194f540f"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-6597eb24ff87bf4e57cf05245b7e9632e8b28062049b05d9d58cc35e2b1f2b46"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 58cfcfee5cb0 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-686930621c1eae2b5b17b12a5318294129d9cfc40fab8ecfec132fe2a8bdb474): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-6848207f62c6191b781d2552603caef8fa61da5685099ff86069f0c5bba47ab0): complete subsection reference.

<a id="canonical-fd269f28594531ffe974a5f26d6d69897cdda4ea39080aeacc73c7955246ca12"></a>

<a id="canonical-1bbd2c2a1649f4597512d708f6f9906c78a58da1e4ef671c46c4c35a5fe5a14c"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 58cfcfee5cb0 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-008.md#canonical-1300ead1766e5dcfb966c531305b60073ca4c08b0d813563997e608521b0eefa): complete subsection reference.

<a id="canonical-b01572d9154a383a782f1febf250873468698918eb2f21e676411f2f19d22b87"></a>

<a id="canonical-4cd02884910143e644726e4efc23d789c7571e65e1a7f5e6edf8c81b57283f7f"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 58cfcfee5cb0 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-8fa624ca8f8168fcb8d11b374ca72cf308b0af9777b48be09caed42e14690cdb"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 58cfcfee5cb0 / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-686930621c1eae2b5b17b12a5318294129d9cfc40fab8ecfec132fe2a8bdb474)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-6848207f62c6191b781d2552603caef8fa61da5685099ff86069f0c5bba47ab0)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-008.md#canonical-1300ead1766e5dcfb966c531305b60073ca4c08b0d813563997e608521b0eefa)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-686930621c1eae2b5b17b12a5318294129d9cfc40fab8ecfec132fe2a8bdb474"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d55c8264f17af52d1ea5d65ec250ee399bb4aab5531ac404b15241dda35a52fc"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / fb3e3e8c89fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-9df72186a6e23ff72d4f7e1d52e8b9ece37a0c12083af84c78bab0cb3bb58f07"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-2ce6f9c95af04be17093c8e0ba80794d5d1a823b2c21fc07fc564d253ef58f6e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / fb3e3e8c89fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a464b04304a64e1e20dfea92ac78be2209477f57ad7d6bb67973f79ac33b30d9"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / fb3e3e8c89fd / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6848207f62c6191b781d2552603caef8fa61da5685099ff86069f0c5bba47ab0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fba0dc8215c595685645622e306f71fef04b098d89e40dc74a4a6941a9db25bb"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 039d31f6223d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1279a26e8f3d291a5731db1993d04070f16f31b892a9753730f51478d9761a63"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-977d30190aa971e414dd47990997fd26318ee43f3f05b3540f2b96d78ce23a4f"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 039d31f6223d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f588d88e59d3c48e7a7be1688187e44595b32269dcba79d550613c5a2becd9b7"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / 039d31f6223d / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1300ead1766e5dcfb966c531305b60073ca4c08b0d813563997e608521b0eefa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca4cf8f93bbbe5f708db9587b7746b09663d396051ab5d7ac6313f156a3f2727"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-849179acbbe66f70c60c6dc7dca1b7cc40318695f06177fbfbf2ce7a7c0608f7"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-d1be078b26253b268bfbc9ffb91e85f2b12f4c11523e44db74880a42c57062d6"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 3

<a id="canonical-aa8be8a7b34b1be830c86030d7b9310570a81e275037e691acded42da4de15ab"></a>

<a id="canonical-05bd1ad9adccd0a411a03a6a6f5f6c190dfddea72876807da12d73b32ba95864"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-73c8d840f3cd2ec07f0ffc62d5f6ce1f92d6c02e17127aa591cddb9eae0b86f3"></a>

<a id="canonical-8d0266016a323b5dd98c650a0f833ee36ea79a6648a08d7cb80fddb8b4bfcb7d"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-407a3dd720cd09a80bea9d532b741fda7a77bc452893c832929bda7c735dad5f"></a>

<a id="canonical-6fff0b144cb2cc1f86b410c9c1bcd7b5ca55e108e66d8304fe7af191e49e7e0e"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8fc236f0b964c0925fb531c16ae3f48dda47de48dcaf92092fec1f9e951d991e"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / 810d213fb454 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-008.md#canonical-b2df7960af7d3d55814a3a2b4c4faff4daa0fb689fa24fa372a43e856abbf4e8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed585b2050770e08e7dda3d61fd890544018fa5e3c267daeea59e7e20ac340d5"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — api_rate_limit.server_url_rules.request_matcher.headers / 33458e52aecb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-5490f418f0ba98bdcbc68ca43bf6f7f7ec8c12cb7653b1e5edb9b77b00d9db3f"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-4603ae59688ff108d26fe96adb44fa32400ab0d19a2c18c64c5e42d2079c9988"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers / 33458e52aecb / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-908cc68bc1d892a2d77b3010f556ae733c27606395851d88c183482a54bc5cc1): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-a71e35d42709ea9dccdebffdf1bbe47a7876a558ea4373c025ed5d517460cf71): complete subsection reference.

<a id="canonical-364759d277cfafc111a07b8a2642abe4589a70a27fc3c25f0f845284476a5383"></a>

<a id="canonical-770cb549aee4b5fefa2a9da12182383b2e73801c77f13eb7a625cb680f3f9d17"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.headers / 33458e52aecb / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-008.md#canonical-3dae434630947eb4ec52d202e0fe2c470138f5c399832a114d0eb310e56a6a18): complete subsection reference.

<a id="canonical-d16a04e7c713fd4c1788551331e412874cc62685a0e22bd4513ecb124cec30d4"></a>

<a id="canonical-66da641dcdc6e14367d778d424d5c35dfbbc91e395acaf90c7a42ea2b046dce7"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.headers / 33458e52aecb / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-6605b130bdea4e34bb391d6f7dc130ecb5d1cbcc5ccf5f8ad072a1b01d8d1d66"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers / 33458e52aecb / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-908cc68bc1d892a2d77b3010f556ae733c27606395851d88c183482a54bc5cc1)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-a71e35d42709ea9dccdebffdf1bbe47a7876a558ea4373c025ed5d517460cf71)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-008.md#canonical-3dae434630947eb4ec52d202e0fe2c470138f5c399832a114d0eb310e56a6a18)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-908cc68bc1d892a2d77b3010f556ae733c27606395851d88c183482a54bc5cc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-497206cb7e092a5ed04516252d6a95826c74aafcefb4a36c762c8dbfe10d6669"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 8b5f8f959d94 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-88fdcc1c074316ad75e9e8f4deec6769981ae62aeed7c96f2a68bc51795c6526"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-d659af10ba995ef0c43c9da421f6e9ad05cb6fccbe844f4364e11a85e97a5218"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 8b5f8f959d94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff4ed51c5560bc8a92eb7cc4952e1824af8f3dcffcf23ea7688e6e01b51ea787"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 8b5f8f959d94 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a71e35d42709ea9dccdebffdf1bbe47a7876a558ea4373c025ed5d517460cf71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f186f38917a8c308753a89445f1a7236ce8dff0c4ac0f0f6eb5b35f2236374ba"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 0a1ad5928699 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-b869e16f90885baec5ddbbece8198fcc9bb52797f24f9d9eca8bf99fe464eece"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-5bb6cc76d06e111eed701e27627c90a4786b3d373925dfb8a5c781523bae2ac4"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 0a1ad5928699 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1677c81b3c5697217ac43abd6306a8d171a8e2391dba3dd917cb75672e8f0df6"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 0a1ad5928699 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3dae434630947eb4ec52d202e0fe2c470138f5c399832a114d0eb310e56a6a18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4f9e6f8b61e884b9f26066f1e86c06c6b9346c90a81473ebfaef86eafcc2b6f"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-c3dbfc8b3952eb0f1aa9c03f6b0e58472a628228967e3b7093bf58c9546b3592"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-af5aa71d0640c7b8d8ed1cf436211aab6cd654882374bbe8ac46d14bb460d343"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 3

<a id="canonical-a27c44145d692e3d4ac66a54f09cc36312a7ceca148f04298e77fc0f5ff91445"></a>

<a id="canonical-5d26f2cd21480f0e3581e2f07712b23ec5eaa375cd8f1f75083a13046069de3e"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-073ccf8b4f30ee377984742b0fa8e66147b5eb6e9238fe6cd934efb18cb95e1b"></a>

<a id="canonical-99a567f3a87176e16f18ed991afc63636b5ff1b0cfe114b48845a9116dabac35"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0f7345f360a8d0683a4075440f84facb9fe261a693ab63a64e6beb961069e832"></a>

<a id="canonical-5c9e1aaa259b93cf18f7bffdcd38a51eee2c792d686550add2d1c11a5ac1f255"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-600e72f83a7f56deab9bed31c7ba9de3a595d484ae397177ce2f5a0a01be18c0"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.item / cff8f7c4499b / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-a219aec2a9170ec5c2139f450579c5064241bf197f744a9d948571269d37f599)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-844a5ec57be1a6d1a4496ec887e1bb50e7e5465f57c65170b6aeff8405224310"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 6f4a5117831c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-628104c4285fa20492b3ae5b15ce1cd875befebb65a8b42b9f6141fab08a55b8"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-f5fc74e1772a2f539bb0f942e4f0e9a4532329cc7c6aba8b5dc5393e05b6d057"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 6f4a5117831c / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-b0611f6bdab15ae26c3f2c400ad12d6a5a88b669e72c449cee9c38bf615229b6): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-35e5a1835911db1a45a9ca34ffaab17fdcaf5a55819c72a191595c421e331304): complete subsection reference.

<a id="canonical-cadd03669127770ada7db114b55973a2f4e11757f1a56f9fae0ce1ab406d7cf1"></a>

<a id="canonical-d356fea26cb4dcf37bf82f29cfc4945800f3d5c737eb8c62bf86640b7cdcdbe3"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 6f4a5117831c / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-008.md#canonical-287a73ab0578d03957a5164a53e2bc2ecbf4a4c39bc36cad2e54a11782b81db0): complete subsection reference.

<a id="canonical-241a6f9640d4e82fc6e8c1a26a01bcf7148a65895bac24ccbf513ec82f0cc620"></a>

<a id="canonical-70bbb471a35b7a757f1783cfcd25c714c3a8d9d7398d0a3a821b978929930523"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 6f4a5117831c / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-472ef718c8097903aa5fc4123ab4418266e46f432524f4eddad13edc243fb70a"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims / 6f4a5117831c / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-b0611f6bdab15ae26c3f2c400ad12d6a5a88b669e72c449cee9c38bf615229b6)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-35e5a1835911db1a45a9ca34ffaab17fdcaf5a55819c72a191595c421e331304)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-008.md#canonical-287a73ab0578d03957a5164a53e2bc2ecbf4a4c39bc36cad2e54a11782b81db0)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b0611f6bdab15ae26c3f2c400ad12d6a5a88b669e72c449cee9c38bf615229b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-024966d46c191e5aac1bfff8181d612dfbcd19c02d74d76006437e5ef4682e2a"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / f5d85dbd30f5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-ead674774882d1d674a17a44310eaf96fa34f6b0fc4d00a161e0e7b4cbb643bd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-60b50e625b024749c6752626fcc4cead192f1d4c6a0350a256d58ac671581e34"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / f5d85dbd30f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2699d0a8c7340b0c3be90b8e92b89bdeca5829fa235976b8b22038dae07fde3b"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / f5d85dbd30f5 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-35e5a1835911db1a45a9ca34ffaab17fdcaf5a55819c72a191595c421e331304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23d0c1990d4e129e0ac55124b4ec2034ce4688773ef70236f057f55676162342"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 1841c3a439e5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-98330ab6fb9874fbd57fd9a0641e5f32902b49f9b482b928e4cf3b323b3e98d7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-610819d9c53e1d64ffeb446506112a1f93d21717c39eb327070a5b2d1020d8cf"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 1841c3a439e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-acc9090ae22addb7695549e0ad8f536e889fdb66e83bf11bf420c1d443bee9a8"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 1841c3a439e5 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-287a73ab0578d03957a5164a53e2bc2ecbf4a4c39bc36cad2e54a11782b81db0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d8a37cef7f30d2dfeff2a660a381adcbe79b78c3f5e7d652186cac1943948b6"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-d815bf1b9eaae4571389e4bac70ff444d6ca71e2afa65638af9db7c712bafce5"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-70dc3a51b9fc1004b7b806d0d7e56875ab840987dd4884bf7b8477b4582c0879"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 3

<a id="canonical-746b3496985ea9eea4e3412884abcd01c476ff036bd567cb31991762bb8d3a5c"></a>

<a id="canonical-a7365ed440bfffbf174d87c2eb7d6112b279c714efb78a0bffa33db6957ed98f"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3e3f02ffb9d161f4e7deb5898b6691fcc04ec112b157fe3145267027e96a3823"></a>

<a id="canonical-e565aece76bd7540863900d63a06327e39ca7af980dc93a8d2ec1b875db6cbcd"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4e8eeabf707742f4cd0e2fbb6449f294dd69fe6928818c55c147caf819a159b3"></a>

<a id="canonical-e8392724cc7685256e06bf3c15887e545b43bdf915f58481868c991437619f87"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-23c7796fb2bb2896716a6f90926c94ae7afdab02aa861a844351555c1d8d42a6"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 1265c8188c28 / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-4326190e692d3716da2f0dc5c9f3ab71f4d026fe4f3a121fedce8b14daa442cc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac941c61587b66f9a675437c9ded4c4175333a47ca46ec6b928557acec9c7eb0"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — api_rate_limit.server_url_rules.request_matcher.query_params / 4a0669075bf6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-c451eb1dbfd3454e4d45014a58213f1d2bc6cdc89a29c4cb0475478d13204e27"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-fb9623a300474f0e8c1740a601c2def98f8cfb97e891b9eda8c1f2adc76229dc"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params / 4a0669075bf6 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-d1009b29d16977f80c1aa39dcccb8605d05480aacfb052bbbfce9f140fcd3180): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-79548fd729c338dd8f92e6739f274f33248ded9ba627e75fd772b74d2c0a976a): complete subsection reference.

<a id="canonical-b0d793105a4ba36b44273f590c35324554300dc6c58d7c7a8faf921a41c4b8cb"></a>

<a id="canonical-ee5e7025d25a52ccb60b9d262fa1896ecf683a93d7157318e062c13344547fe0"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.query_params / 4a0669075bf6 / 4

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-008.md#canonical-382a522c5e659b878a470c22bf4cfdc80e07c408d945c0e92751b48c922bd000): complete subsection reference.

<a id="canonical-cf0cd32c7754d1b94798f5063adef426e6af7f10b830f9c4976f26b4f6f79e4c"></a>

<a id="canonical-1dc633de4baeead8dc5dc0a2d174e483f41fad429d83e2a58f1b15019457c171"></a>

## key property — api_rate_limit.server_url_rules.request_matcher.query_params / 4a0669075bf6 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-846e3908f2d9e950df04719e728a96bf445e61ca9560db2fe87ad1e71ead0e34"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params / 4a0669075bf6 / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-d1009b29d16977f80c1aa39dcccb8605d05480aacfb052bbbfce9f140fcd3180)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-79548fd729c338dd8f92e6739f274f33248ded9ba627e75fd772b74d2c0a976a)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-008.md#canonical-382a522c5e659b878a470c22bf4cfdc80e07c408d945c0e92751b48c922bd000)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d1009b29d16977f80c1aa39dcccb8605d05480aacfb052bbbfce9f140fcd3180"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe2c4ffaf2663de52cf9b06dd9af3ddb123a0eb98b9218c11ac7efa78969a6d8"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 0c8df282f180 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-e8d4264ba81de8833b92f52e4abe01c8ac0f0a5976780d968d8d7f11d114bca0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1cc9dee92c4a4ecb357c7c586d0d9a7bd2329696a05f4317831f188083ce1a8a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 0c8df282f180 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a3b5606dc828de58ae3b81766d539fd6cb6a979c9912d4055f9a0804abea6c4"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 0c8df282f180 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-79548fd729c338dd8f92e6739f274f33248ded9ba627e75fd772b74d2c0a976a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60caaff81dd8859733f79a5e0d132ad9e543fade14c568bae30af31d9fc45709"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / be72b57d57dc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-7c7a54697673be62add8fff522af1bb2628bac6e1478ea4bf8e0efac9403e252"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-3756f686d0ce8348931ad0f2fea4916849a8d439b12d3fbea261a9024d9314a9"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / be72b57d57dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a818f08f191b05ec6f2a9725bfaaeeb660012ba09d38a44efd8839e1e9f35fa4"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / be72b57d57dc / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-382a522c5e659b878a470c22bf4cfdc80e07c408d945c0e92751b48c922bd000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2d650dcb4fa25506c6fe06b5067e8a688b6f81eaa7e465d149745d323365918"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-b14a3ed07c058e5a95f4d707dc55d00118d1dcdff2a15cf77f4c974eb9966377)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-5e497f528941089d9f09d9b6ce4a6fbce8636e3487428312d7c29e4f129e4492"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-8615f210cdbd6cb7f446f855a5eb6645e2f7ffb54fdf9bba1f5f1cf5697d343f"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 3

<a id="canonical-699b4fc30044396b327541c52f3bd7fe6cacdd22c45a85264c0725670ac19f10"></a>

<a id="canonical-3a3671485a6523b042ea8bc21df83dfd030a11d8d80623d16fae39d673e15401"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3f1ebbb0c8ad4ca43f0d3fa8d87fe45b8e73468c5c3631b4efc798554db21c26"></a>

<a id="canonical-42238b9927ef7746c9a297f15fed5b5f4214e34eb523952ab1fb56c1b859e9f0"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c96946139cda2857f77c3eb50583e69ef15fccc6b40abee75303d7f22981922e"></a>

<a id="canonical-ed49705d182684e0da32777ba1925632a23c88fc55c88a39de8864591168c97c"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-14cf5797cf7168b5d0b9ecb21560a565b521fbce0759d1a85f6c79e6a64f7c38"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.item / f22b1d7fdbfd / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-a0da55fa16655e1ab8e395bcdb8b4631224ebde6995fbc508dc919f1f5938182)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df6587053c1f0facfe61daffa50d143dbacf1e0f3c39507d0fd80a4c834ba93e"></a>

## api_specification — api_specification / 1dedc2de6b04 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- api_specification

<a id="canonical-2f38a3f5c01bffe8cba7b9c700ba11b1fbfc81c8b8d9e01adda19d3068f7f919"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-2f38a3f5c01bffe8cba7b9c700ba11b1fbfc81c8b8d9e01adda19d3068f7f919)
- [disable_api_definition](data-sources--http_loadbalancer--reference--group-017.md#canonical-343427cfdfef7c824b866ce80898525f5fbf5b84e5bed9e61ebe35e1f2ae9bad)

Select alternatives according to the provider validators above.

<a id="canonical-6f449f15471611bd6795375955e11b45932d1b9de9ea9efe8af39822334b3536"></a>

## Direct properties — api_specification / 1dedc2de6b04 / 3

- [api_definition](data-sources--http_loadbalancer--reference--group-008.md#canonical-53b8b09e1bb34f081c11012c34d978ce9a55cfc87062c88295a3e74c5afa453d): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9): complete subsection reference.

- [validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588): complete subsection reference.

- [validation_disabled](data-sources--http_loadbalancer--reference--group-009.md#canonical-8876fd0662bbfd750e6392c12ef9cfd61476036241f3d647e1ed210220a7cd38): complete subsection reference.

<a id="canonical-fde04ec91dbb66939fe45d8f7c97b159d62dd0606b9507ccb774e8cbcd03e8ed"></a>

## Next pages — api_specification / 1dedc2de6b04 / 4

- [api_specification.api_definition](data-sources--http_loadbalancer--reference--group-008.md#canonical-53b8b09e1bb34f081c11012c34d978ce9a55cfc87062c88295a3e74c5afa453d)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_disabled](data-sources--http_loadbalancer--reference--group-009.md#canonical-8876fd0662bbfd750e6392c12ef9cfd61476036241f3d647e1ed210220a7cd38)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-53b8b09e1bb34f081c11012c34d978ce9a55cfc87062c88295a3e74c5afa453d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c289dc7f953fb0eeab83d3a9a08ee6f97317f9f76c3e7cce5fba8463b020a4b"></a>

## api_specification.api_definition — api_specification.api_definition / 96d7b0bd9610 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- api_specification.api_definition

<a id="canonical-a7e7bce60471ed6d9bc31a32a30be2fb789e7700864240958a53099d1dd88c81"></a>

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

<a id="canonical-93e72bdb5733cd35f8e5f7852a16494d5e15525b5aee1538f3fc57502f06c09b"></a>

## Direct properties — api_specification.api_definition / 96d7b0bd9610 / 3

<a id="canonical-3619cc7882721304f218fe74a9bcac9f514eb222982a79d5adcbe262396d9645"></a>

<a id="canonical-266b08327faa0406fd41403333ef839859298e92ad28242694aaf5d0adb539d6"></a>

## name property — api_specification.api_definition / 96d7b0bd9610 / 4

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

<a id="canonical-6db8824c25d0c3d4a62f91fae1f51230227f9fd90b95d345e761a26d8d103391"></a>

<a id="canonical-dcb202955710a2180a446d3e6d52ff6eb1dc6071355332eae0f408713d804836"></a>

## namespace property — api_specification.api_definition / 96d7b0bd9610 / 5

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

<a id="canonical-13f7026e83f4fdaf3f69d3dd0f2f8df601c89776bddbb67b3db3b9dedf0bd037"></a>

<a id="canonical-6734a934f3714b4105f62f22d782d3dfa2d90c13b5bd561db8fafdd623fceb95"></a>

## tenant property — api_specification.api_definition / 96d7b0bd9610 / 6

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

<a id="canonical-1567146e30c350990bdf58e34ed3e45e86cc42e4d1cf73b21deb2ebc3f69d612"></a>

## Next pages — api_specification.api_definition / 96d7b0bd9610 / 7

- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b283926aba6ab04ffe132c882a6214cf6efe0551f87cf49d5d4a008d421ebe5"></a>

## api_specification.validation_all_spec_endpoints — api_specification.validation_all_spec_endpoints / 75b6b716a1cc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- api_specification.validation_all_spec_endpoints

<a id="canonical-3071646cb070b66a04f5c86972ca2981644120296b9f276c8400ec477f94f19c"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-25b720b7e1b28731038633588ba188ea853dfcd13e864c2d5092bddf7d49ec13"></a>

## Direct properties — api_specification.validation_all_spec_endpoints / 75b6b716a1cc / 3

- [fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f): complete subsection reference.

- [settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf): complete subsection reference.

- [validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c): complete subsection reference.

<a id="canonical-58c8562578fb50cc9fa91fbad6f1cf9f4e478fe92c6f0dc64cadebc4f45d7a23"></a>

## Next pages — api_specification.validation_all_spec_endpoints / 75b6b716a1cc / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73f93a2501a511db6bfdfccc8e9804449ca4625b931ee3d85e81c0429075bf07"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — api_specification.validation_all_spec_endpoints.fall_through_mode / 40bbdbcc12cc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-888c3c416ae75594ea91808c38a43b3fb8be6e579cda446e504a6b841284e13e"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-dfef11bfd1fde9dc4ae5312585ff5be87c66759243ac072364a97bdbbd171e06"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode / 40bbdbcc12cc / 3

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-008.md#canonical-2aee7f45fd18417acbaa7cf46e04f498e95cf5fbfde9543e56f46c3a127b220c): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b): complete subsection reference.

<a id="canonical-1013cd222aec4e12b5aff4b2330e69bc0819ce017ed9def834833fc26372c766"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode / 40bbdbcc12cc / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-008.md#canonical-2aee7f45fd18417acbaa7cf46e04f498e95cf5fbfde9543e56f46c3a127b220c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2aee7f45fd18417acbaa7cf46e04f498e95cf5fbfde9543e56f46c3a127b220c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-546a24ab529d3c676ac5c39ec6c5a556f9521bdafe3b7feb73407bf8531b496a"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 206facfe8735 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-a3ff51c327502d670da3b8f9bbe4d699e43e1a3323db1e13d5ff8e15b518e142"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-0509063ec081d5d2ea36f0f1cc04ec10acc8c6b1eede21f11b9689cdef4855af"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 206facfe8735 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-862ff6c4531df7a6e19ca2accbc1909244b640344cdbadb1e684e45d74a28ad4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 206facfe8735 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bde2e2fe7a59ed0eef32fd664132d927c8444cf73ea9d6de1d1c46eaa64efc3"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a99c434ab5b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-6a1debc92551e7f1a588771846812b3c17840384c6f73963b95df326a0581a46"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

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

<a id="canonical-0cd89ae83a787a497ad2e15422c4b1777d51cfaf943b1aef5791fb0ac4b20cff"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a99c434ab5b9 / 3

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a): complete subsection reference.

<a id="canonical-0ab326d31ffa9a953b1449f58a58d4301b8775a1156a3b4532f6b4b70b83db7b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / a99c434ab5b9 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f8047315629b640ac43addabb2a73a51b0996015d9f79a4b476db6921621a05"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 15e61a4b4ec1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-29e5e5152f76e7c9e31e6db232409b265c07edce0193e92b6500424873cad4e4"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-04a13a226a062d8b9f3a3389af220d216338f1e3f3de63d94609ee965fa59b35"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 15e61a4b4ec1 / 3

- [action_block](data-sources--http_loadbalancer--reference--group-008.md#canonical-55045ebb44a08e5ca399273f65fcb798230c1c3ddb06ae35e44c65f2daa203d4): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-008.md#canonical-a54c362a9a89ef56cdc18100725d603b6285b34c1d5ab6d23468d4c2736a87ce): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-008.md#canonical-a06ad948012fa55d45fe528fec8a5de583be086bad54ad6e62f6a0441bf72862): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-06afe2d77108c161e5b11add33eae536bb73790ddb1e880b3ba2382bbffcb867): complete subsection reference.

<a id="canonical-82c8d7e4254289664f30e43a1beb78f69e87b3c52bd69f8b472c16b44abe6d8b"></a>

<a id="canonical-14d64f423cb338553ab114b8f1c6f8d963db350d1281d43dd7755f0a234e4471"></a>

## api_group property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 15e61a4b4ec1 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-bbcf57abbd079fe3b5786f409c64d6737cb489fd5c51b32dc04ec32df409499a"></a>

<a id="canonical-a2923b7928454881fbd6147d2d5db7db3a346e5089f494375ae9b21b42d6fe00"></a>

## base_path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 15e61a4b4ec1 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-c838a2070be0a009813961ecc76c0294317f622fa57a424a00300c9097782711): complete subsection reference.

<a id="canonical-8c2618a3ca1c228b1e1c73c376302e1897b0e9f9f68c5275ba70e5e2cdc71a46"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 15e61a4b4ec1 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--http_loadbalancer--reference--group-008.md#canonical-55045ebb44a08e5ca399273f65fcb798230c1c3ddb06ae35e44c65f2daa203d4)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--http_loadbalancer--reference--group-008.md#canonical-a54c362a9a89ef56cdc18100725d603b6285b34c1d5ab6d23468d4c2736a87ce)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--http_loadbalancer--reference--group-008.md#canonical-a06ad948012fa55d45fe528fec8a5de583be086bad54ad6e62f6a0441bf72862)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-06afe2d77108c161e5b11add33eae536bb73790ddb1e880b3ba2382bbffcb867)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-c838a2070be0a009813961ecc76c0294317f622fa57a424a00300c9097782711)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-55045ebb44a08e5ca399273f65fcb798230c1c3ddb06ae35e44c65f2daa203d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55609373b33fd8608601a1368f25344183a1e40308a0ef67e6f0637215cecf2c"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 8ebae0b3248c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-22a22275f7dba0005ea57542939361488a4b2b88f0d58a173e14d87e9ec27cd2"></a>

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

<a id="canonical-97929003d7e5281cf27965bcf719b401efd2b38ac8e18ae242a4b3dfb259405e"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 8ebae0b3248c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39445049fe3efff5796bbcad009b278c04ecdfb4f19d03c4792eecf4bfd2a91c"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 8ebae0b3248c / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a54c362a9a89ef56cdc18100725d603b6285b34c1d5ab6d23468d4c2736a87ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab783c9cf41291bd8c7b9782208ace1f56581ce20fdd88bbd182a3a4bc920034"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / f63194f1b8ee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-7459075939a2428046f2d41ba545172f3137ba7ab7dd592a9408695775f00db6"></a>

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

<a id="canonical-8a03db34f594d81dd473f28b855e0568e702b3b2811e2634a115e79fd67fe4b2"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / f63194f1b8ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9420db24553a053db0cbac2564136babb88043e00a4160ac5ab702f9354f5d83"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / f63194f1b8ee / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a06ad948012fa55d45fe528fec8a5de583be086bad54ad6e62f6a0441bf72862"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1db380856794b092f74cf4ef15aec86daee3811832a63654f2339d68d34b573"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / cc578a3f036d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-8da978df291fb4e10f73020a75c44740fb09bdf292754f40463db392cd1c185f"></a>

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

<a id="canonical-28e50e272e0ab4da7b6a984dcf3057cfcf0a632f2a5de0d657eebab88d0a054b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / cc578a3f036d / 3

This is an empty object or choice marker. It has no direct properties.
