---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cdec6cf944fea4dfbb9e68aac3e7f0cb026918531e5606ce3d378c8150553b5"></a>

## Property reference — Property reference / 51550e4f53ec / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- Property reference

<a id="canonical-2886fd043268303823be5a9e739150f0124b78333a9c4c18ca18574a09dd7083"></a>

## Direct properties — Property reference / 51550e4f53ec / 3

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-a75c7c370855a172adbe82000d2d5c54aaef55974e04c0d0f1fc7e451b8b252d): complete subsection reference.

<a id="canonical-43c8d4401001c42853efd0d13c03fc436d9381c9c5ca4684abc28372b97a744c"></a>

<a id="canonical-0cffcd7faa607d9eda6bdad52316202efb45ce787b74ae6cd10cd912e2835f8e"></a>

## annotations property — Property reference / 51550e4f53ec / 4

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

<a id="canonical-b0f57343862ece8b5ad038f16e89e285707763424855baae41321b190659d07e"></a>

<a id="canonical-949f6ec869f0d845aee98eacb3ed276a40a65ea65534df035e347f3dfa90f4ea"></a>

## connection_timeout property — Property reference / 51550e4f53ec / 5

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-d8bc324b2c781c8510f6bd5d50098da95af83cc871d3bbbe5b6e1a7fb4e2f609"></a>

<a id="canonical-b7e293dcc97f763b9ad5e23da4f19157b260f69a8a5b6db378f894a38019eb8e"></a>

## description property — Property reference / 51550e4f53ec / 6

Type: `"string"`. Computed.

Description of the Proxy.

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

- [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-6ad02a6a79aced4d88d3b8c364f09978072288f3e3db7685c0090af711fc8703): complete subsection reference.

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237): complete subsection reference.

- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2): complete subsection reference.

<a id="canonical-5ca4138f4c316fe75b28cf0175c46c829813170ddf3e67c81c857b0e73894152"></a>

<a id="canonical-bb33eef06702a9199b45ba1532f4cfe19caa4fb3d72a3f8fef9db5954a3a2b8d"></a>

## id property — Property reference / 51550e4f53ec / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6770d01aca12196e3bd7f87f64e86b3cdeed9429237a6b93cf1f54ae1426c5c7"></a>

<a id="canonical-f9a6e2bee2c04549949d58790f82dfa45c7242edea579e4b2cf86eb90fa60c50"></a>

## labels property — Property reference / 51550e4f53ec / 8

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

<a id="canonical-d3bf05cd1e46d7760d249c8f3485484f01c640eac6cec6df5f0f711d64d87c75"></a>

<a id="canonical-30876fb63037873f1cce1446cbae40580b87ce1003b827b9ca5e4271bc36e1f8"></a>

## name property — Property reference / 51550e4f53ec / 9

Type: `"string"`. Required.

Name of the Proxy.

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

<a id="canonical-7b659d3d701f61890141fbaa8face813bd5945157e0373ddb8f31c476bd9135a"></a>

<a id="canonical-2316ecb5e5c5c63480edb7832bd86bcf7c44fd0e1b6e994dc3cebc407ff8b909"></a>

## namespace property — Property reference / 51550e4f53ec / 10

Type: `"string"`. Required.

Namespace where the Proxy exists.

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

- [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-354eed661a275799f673f7dd3b219dd0d9c148389186b9bd151c99738179e6c9): complete subsection reference.

- [no_interception](data-sources--proxy--reference--group-004.md#canonical-4cb287101d7cff732a91ce780d465332c4fc488b430e1c588c72d8d91fff4f3a): complete subsection reference.

- [site_local_inside_network](data-sources--proxy--reference--group-004.md#canonical-511ab4dca177ca5d7d292f2efc016dc5ab5e4f01c63811a03340a2195219d67a): complete subsection reference.

- [site_local_network](data-sources--proxy--reference--group-004.md#canonical-7772efa695c7a5d2abe78ba96890f5435cd65223acd8167714b620ad5f4bda05): complete subsection reference.

- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c): complete subsection reference.

- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6): complete subsection reference.

<a id="canonical-f4a6e69f771928fc85ed128a2379dbbc87bfda288210b2d261a13b89ae82feba"></a>

## All schema paths — Property reference / 51550e4f53ec / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-1759143cc65e63e4b0c8dcca1388f102814349bcf988373b637ba3a2b02967cf) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-eb8b3bd8f3d9b962ec5f4d996202a5f6283787e8627911aa6bf0ea8fc8b97708) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--proxy--reference--group-001.md#canonical-0faa220f2ae05fa331569bcbf78ff59384c2bda93018f397c3496ea8045319d8) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--proxy--reference--group-001.md#canonical-13008d16a77de36490e46d17dd43bf3b843954a06d5dd5c0cfcb3000af498914) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--proxy--reference--group-001.md#canonical-75432b3c0e9cc7b1b6bfcd62ee6411ac49e26f220f80e32b17a476f41c24e16e) |
| `annotations` | [annotations](data-sources--proxy--reference--group-001.md#canonical-43c8d4401001c42853efd0d13c03fc436d9381c9c5ca4684abc28372b97a744c) |
| `connection_timeout` | [connection_timeout](data-sources--proxy--reference--group-001.md#canonical-b0f57343862ece8b5ad038f16e89e285707763424855baae41321b190659d07e) |
| `description` | [description](data-sources--proxy--reference--group-001.md#canonical-d8bc324b2c781c8510f6bd5d50098da95af83cc871d3bbbe5b6e1a7fb4e2f609) |
| `do_not_advertise` | [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-551f1846bcb96d8193dc2dad151583cd0755d4dcfc56f89f987b325ea3ea947f) |
| `dynamic_proxy` | [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-f0194ffb37cd96597d0c0ebcf7db24e8e420fef0265b93dd03471664f5c2b5e8) |
| `dynamic_proxy.disable_dns_masquerade` | [dynamic_proxy.disable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-cc213a8c01504ab52839b0b5ea68503ecdde8ce5493a9e28fe79cc58ecfb3695) |
| `dynamic_proxy.domains` | [dynamic_proxy.domains](data-sources--proxy--reference--group-001.md#canonical-921250ac465e65c42c814a483ad4fcb71509a513543859b724465fe4e131162b) |
| `dynamic_proxy.enable_dns_masquerade` | [dynamic_proxy.enable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-a2b3c12a615021c5665c6905a039d478c51a12a28f4e2eb61743db117f1a70d8) |
| `dynamic_proxy.http_proxy` | [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-704fe152d4940ee3f666eeaa9bf2b16d9807ae15e0b60be728dcbcc8898a5bb1) |
| `dynamic_proxy.http_proxy.more_option` | [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-d1d672e0a511da5170d1d50a91751419738ebc9a1c6c6fd0c54e9c1ffa6871c5) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy` | [dynamic_proxy.http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-001.md#canonical-2aad559a3513fa6d7c58fdbb543a70ba6f19c969e2696cde2f5bca05aea43e0d) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-001.md#canonical-b20cc96af27e3d63953255eb68a6e7386189194bd9b7a1037e14cf045c49ac28) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-001.md#canonical-7a86f4ebe6e4b1c52120484ac4c3856a31f978363814a3677d754e225a709e43) |
| `dynamic_proxy.http_proxy.more_option.compression_params` | [dynamic_proxy.http_proxy.more_option.compression_params](data-sources--proxy--reference--group-001.md#canonical-ddb5a698751ae82d07ae19cfde8d65c799486da8da82025245d665e1e9c9ffa3) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_length` | [dynamic_proxy.http_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-001.md#canonical-18abbee002b0c6fcd299d64f33a4da4d44db9a1266678d89caa9def75034926c) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_type` | [dynamic_proxy.http_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-001.md#canonical-6359224353187f5026696671a8c9cbcc5a623fbbc7d59890b09c5b3944c69ced) |
| `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-001.md#canonical-72367c37475ec266fd3a2374812e6c00b492f55b361375283de84a0cda8924cf) |
| `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-001.md#canonical-9545b3e205fffa4a0cd404d5e3cef7d7bda3bb1259cf591361ad8f6d94086cea) |
| `dynamic_proxy.http_proxy.more_option.custom_errors` | [dynamic_proxy.http_proxy.more_option.custom_errors](data-sources--proxy--reference--group-001.md#canonical-4a7e96a266f7c3e576940053c78a1c3b54c1a5758cc5718c20c4ff74bf1fadbd) |
| `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.http_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-001.md#canonical-aeefd9aa56263e81b2092d95f79dd38e523135d2c7a86a41d21bbdecffeee5e4) |
| `dynamic_proxy.http_proxy.more_option.disable_path_normalize` | [dynamic_proxy.http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-57bddd9adc4aa43bede35dc374a49f0bcca9a5419129381d07eb3d9cdd1e4971) |
| `dynamic_proxy.http_proxy.more_option.enable_path_normalize` | [dynamic_proxy.http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-47829ea8b38dab50a8963d95dd1b5dc5ce13866ace22b5104aaafdbe6603d33c) |
| `dynamic_proxy.http_proxy.more_option.idle_timeout` | [dynamic_proxy.http_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-001.md#canonical-e56f716379069db659b901686890c0aea186e05a93b8232c8c2c7bbd211be0cb) |
| `dynamic_proxy.http_proxy.more_option.max_request_header_size` | [dynamic_proxy.http_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-001.md#canonical-3d39fb99af0c676b0e19d99b60d885a5f3d22e627d792ed31821ad38ccae2a57) |
| `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.http_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-001.md#canonical-b1affec10a5a42c9ea183450f0ad59b60f39306e85826225e557192bb2b42ba9) |
| `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-80ee11b8d04ca234dbcf80614c26090b4c06c2618c43bc9dda69c42f349bf2cd) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-9ce35c4c6e52e041b337beb7d19f70aab8b0dcbe6f740726f6c3a07f9dd7545f) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-6e5aeb7893762817489494e421b01b3e4ee89bd3b0a8e60eeceb7392016c6686) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-f5453a16427af8ff632521d122ab7a402e60f465ccf3448def666deaff2da14e) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3fb33b016fc24bfdc892de5e4dd5a98168bbc715803af0816718c0694ff46cc5) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-65cf6a0a762bb63bc5973ff5de54466e17c3c34bebe5f0c594c6459454de2e05) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-f3d6e57bea5c1198ceace61ce64a4123288e89731d41ebff26409e9076c9309d) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-6b30b46dd62a43bb126965d520a6d8cb38b05b59d0c0c6098dcb7f7ba0702a79) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-e45f43f0ba094fd20a9a13fd127fd8c3442f4bf6b3917eaf1d47e2c88d325020) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-b4ed172e5fc6d15e8e4ca6c88a5ac7c1c3805acde6618bc89e26bc50853d999e) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-36f6b3ec250cc47bdd9dd08aaafec4bcf7b70b21d3fd9856718a4fbc9d4d3d15) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-bbc0e9c4d994d9a9d8f1b62a500ae821c2f4f4f787506b835a2ead83410b8836) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-cb32f26822d62c3e4f1bf19753eb1ffaf800b36c22ee3808fb2d12c096fe6122) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-001.md#canonical-845a8d7ce9d0f8b0ffe43d8aa33cc009aa8764524b46f0ddc7b5532fe4a51fbd) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-cf9bc0b1ab5acbb8f38135bf56b0a1c7379f51d932747f72ea03e6aad104e46a) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-b96b2b32877ef8115ff49626acef799a0d9ea6616de73474c30cc4ad17278d45) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-b87ba95992ff292ffb215832a92d30de2b522309347d7779032ab332ffee3cfd) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-9cb324a65e054ef223dda45bcc6b67dc1ef23074614dcdfc02261642f7b396c0) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-2c525aae3096d26f7f4aa121b7d0082243fd13d2f7b1732911dfd9a9c15b910b) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-c12dd344adf20d9fba46cdbdc5e2b2ab0030ffc81804d11ef8ff36661edac820) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-75d5d310640b787f16fe99c8245a1948f9d8effd867161798f499f3850885f51) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-220e219427ac4374a7c20437f4c84e996db9c133e150469d0264b9ea03909b66) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-7c32c900e9918c52931903998833f4412c8ee1676af35e33d6732546f3e8d6e1) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-762d5cfe37b7864b0da4837d3e05890676697a0007329d6fd38db0aed6817e60) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-fc760adb3f7752471566bb14a087657f57f588d88fcb676a35e738af8fe059b9) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-6bcd4475b71ebcc059325cdc63535eb9c326853fd4c6aa6b1df5dfb2096bd872) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-001.md#canonical-6f1bd4a12305d06b7714256a1766a923a04abb7e4214d67f0596fddf7143c211) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-b3a572c4dd1a865d3c631a18d053dac333bde3c916a1598a6eee1fd627c9f505) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-002.md#canonical-cb1e974fe2874c70be10e4926f360c214b78340550292eae9987d47a9ea4feab) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-002.md#canonical-4f94b6a81d624073ba26caae335631f76de4df9cdddcb1ab63d80a7ad0f7750d) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-002.md#canonical-3469cb9696cc8ec4cd2c98605c51d69ba1786545d71d3ffb018c3846c9b306d6) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-002.md#canonical-14cb5967c1cee5ce033856de1fcadf7d74d11c66009e9f8b4e3da9f4b7b1acab) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-002.md#canonical-83a9f33e5bf8a504d6be725f1bd9ddc6d9998756c287f4e7ae001415080e6cd4) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-002.md#canonical-287657423ed99f877cbdf456e8f3efc0f7c91721b0c185a9d01a30e0bd61e329) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-002.md#canonical-91257ae21ba126981ba4dce676a1226862ac892802200dc2dd4cadbd86bc1752) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-0e24259d035d621dd8c72e11e7dfbc3096ac16a529edb1bc052179c85830d8ce) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-b7d48b36a81a4d99f7e89bb9e05ad187716cf552ee8d8710bec67b5231bb0240) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-7918fc7ef29b6bb31a96c2c21955c66a914c69d5543468d5e78257c0d8247b99) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-ac1dc9080374f544b4b1badf2d9d67e08a2676097c0e7b2d9a631e70063a9ba1) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-002.md#canonical-2152c076e61d1d27b85058a6514179bea567214b49615039e611083e8dc60c12) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-bc533acbf7c9f94a2f328db4535bc8eb69d9db1e6942459d436cb43e365768ec) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-002.md#canonical-f384ed6488240820a8b82c6e255fabc92c85f5a54474a07e9ef05b45444761c3) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-002.md#canonical-b1448c0b159e00097b0003f1f4f521af0c9e67e66e13d847dd1ef0039dfb40da) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-002.md#canonical-833756796b748899149d0ccd7d75f4c0716b52625345054cc12bd20bbdcb1a8a) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-f786b23123ce6e805784a8d85c1892071bdb60b61c1b9b0a2d67506a502dafd1) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-14aa70f388a1b7f4a4d54af68739638142530f1cfb378b15b8897278f0ab58c6) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-002.md#canonical-a27b81d3c41a6b4e91af070a419dd4d6ad8bf84cd90e7d78b8c09ec1331dfb37) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-002.md#canonical-4bd6cdc618de98acfe99da6d16e24d70f00c10a30b2db110c1d4a73571a07981) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-002.md#canonical-ba778213b071323ad77ddf8b0d62c614ebbf84d5ed4ee259e3f8e7109fa660f3) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-c60ba0e12dc0b001324f73f0e75c69cb6bd74aa261c8b326a2482e3f956642d5) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-10a313e1d68580cde62c4758455a6a9184e0b2aeab302c16ddcd99e3cd760176) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-39aa6c2da6a5074b6f8093ae12b4f166176ae20cb82036deca838bed05dad9c3) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-403ed4befd468433ada22e61b6d2f565022ff57e5f479ecaf6ddfd2d61591210) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-eaaee1f0456a2cc457490c3384854fea6ba4717be82e9e609b7743f132bfc015) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-e54aa9b20fcb47384deca75e19750a916a34997c7685d46da5020c946901a509) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-f385cae823cebca5c125e82a6a8ddc276a54012e530e52309eac0ac299faccd9) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-3d0f6af85574c6f520e7a0e9799187a89d281f5cc3017176a506357feb7fbd06) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-6cbbb5115234a90a49db9ee73007b9f3969fc8d9b2e7b33328763152591858fb) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-001.md#canonical-7f6c02099ae9ef57f211dcf0383b7cc38685dc5e608230c453aa1d831fd921ef) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-f69bc7266162a4952117339ca877a862759c9eb77244fc97d573e914d374f367) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-7b1e6eea1fe4541fbf9e2e126bfc4384ab0fa867ce3d6a42b31d9a03b01e4d04) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-464d66ea71999e9db55f03c242cbc19cef718952882a6e6572a22bc4c4da0b82) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-22dac4b42255a01a80a43178db25350ca73295de4755308278fb38e6d4db2c61) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-59a37bcc584ae8fb7e854f148eeb46786c61a36ab1157de02641dfb8faa85fbe) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-a369cafbc3919b0b49a1f3ddf80f657dfc3043c817587b49ee165044df02fbf0) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-116c433a0f4a5c9d808393dd7a76154ace483a1fc59d1b4e82ba1705e6db9a31) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-dfcf56ba74ad33b28712fcfc76e95414eb811ad548c164c9244f42c919a44329) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-b9d8be5054ce58a67c3ffe848ad3672d3a51640f708c6d319dd386bdec403645) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-d287a047e3847362c3fcbc86ddd3a37f93099c97fabb72a43049d50e5f3b03b9) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-62536d204dd2f104d4d4a9149f6045c31e25d004b80cae994766b35718a4df69) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-6cdfb89a88a20e5dfa6e8cda6e8b2177b133cd48efd9839f159ff7b14a65b8a3) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-001.md#canonical-7604440cb653892ecbebef9e68762a15ecf2a9e60534e7dbf8d3da08a9b922ca) |
| `dynamic_proxy.https_proxy` | [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-6ad4339a72076547945ba09bf8fc307a34a29f67ce398104448155d83006fcb8) |
| `dynamic_proxy.https_proxy.more_option` | [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3cf0e4b22e037bb1dc3ef708fda0c0c33056ddc4283dc854f93c09d0fdee94fd) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy` | [dynamic_proxy.https_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-002.md#canonical-69bd3208031b38fd26b2aa91837e6281191a4296139f0b16453cfa33434f791b) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.https_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-002.md#canonical-6367e3a06ba5155d745aa4a2dfbe079fb750cdd3c56e7fac9c6c79d7040afdd8) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-002.md#canonical-497530a83028bc64e0105e32d863ec1c2568e952170ca2ac41e9d1b39db11752) |
| `dynamic_proxy.https_proxy.more_option.compression_params` | [dynamic_proxy.https_proxy.more_option.compression_params](data-sources--proxy--reference--group-002.md#canonical-ddba3a737176500c00c2d6bff89a71bc933d79584bb0a45dc4438bce635bd99e) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_length` | [dynamic_proxy.https_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-002.md#canonical-75cf67314b09f0a6c4a24cb193aca05e0a3adf8fd611e2520dd818916ef6e0e8) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_type` | [dynamic_proxy.https_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-002.md#canonical-22ed6955743f68e540366834b70f45f30ac44fd0c70111d61bf25b4fbd6eac0e) |
| `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-002.md#canonical-231d1f6bcaac5ba011dd6a9145b2d7af36859513fa72e8907107f1ab74bc5a74) |
| `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-002.md#canonical-44a511e78ee28aac524050a3f580d1c6ac46d19f4f73c59762b5622d40572f56) |
| `dynamic_proxy.https_proxy.more_option.custom_errors` | [dynamic_proxy.https_proxy.more_option.custom_errors](data-sources--proxy--reference--group-002.md#canonical-72e122837ffc050846f6c7936a13f93ca2e51a1aee469340127b677295c43f31) |
| `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.https_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-002.md#canonical-2baa9c841948696c1959063a1bb264cc024571f20cf39b86c1c123d35f48f0a6) |
| `dynamic_proxy.https_proxy.more_option.disable_path_normalize` | [dynamic_proxy.https_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-d3f90a915608440c4a7144f830a5d81c2ac2b46e2b4db2349b26523de7ca7c3a) |
| `dynamic_proxy.https_proxy.more_option.enable_path_normalize` | [dynamic_proxy.https_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-b367a7722c034c51b84dcc0c0cb9619a729061d786b55c65190c85f4f8ee5183) |
| `dynamic_proxy.https_proxy.more_option.idle_timeout` | [dynamic_proxy.https_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-002.md#canonical-36f0287ed00e33a311a59e1c6965d688860d44e7f665e40139285ec82bd3ff70) |
| `dynamic_proxy.https_proxy.more_option.max_request_header_size` | [dynamic_proxy.https_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-002.md#canonical-029ce5df60a100fc2d36b0413c5f37a922597fb556618ae567c5dd4a9565dd7e) |
| `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.https_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-002.md#canonical-be53742e5daa434b33196d233b4a2c38f380ca81412ed358c502b8c84e830aca) |
| `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-35e60ef5de69418ae9ba01fcd4290b93c608dfe0766b42f4d6de965fbfef53c5) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-580dcf92d9cdc7d50b7fdc39021ad8ea51029633bb1b1bc35b34a38074113a46) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-002.md#canonical-502cd6877d0ce00585eb2cf1e3dac64caf82b8e030e233a26a05a5c08834b0c7) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-002.md#canonical-5c119413c5bc68d594994e9d04d9812fb9536066ff04904548f5c37870094464) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1f5e3a0ecf9db4804700bcabd3f4183a10d528e7f98f3f559e4efcda64a66991) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-6d06b5af790048c54ef456cd8d42ebbe54d5402e77bed00b4c4d3ae74e32d035) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-002.md#canonical-d3083121c97b2b05f011cafe9156c8a4154f56d015d7520b9311c2fd68a44e32) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-002.md#canonical-de895d44f5138d0f5aeb3d9a63181eb624df449fdc8d95c6077711eb8ad78c74) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-002.md#canonical-c10dec6b6e1f4b831f425a465dc3da7654a9a0b385f81db9ef07d9fc3bdfa9bd) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-172339687e7b655a711aaac6a580307c1af7e1ac6bebda355b14d189e0f68649) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-002.md#canonical-d5c288d2cd887bb9b654e43ec1432b665123057832ea57b263105a06f131dc1d) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-002.md#canonical-096730efe06905075f5c87f5010c6c62525f4d9eb7e388c42f41e2e8bd8cec71) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-002.md#canonical-c94364f5923fc83507e9ac5b1a53f964fb5b8d52f4fab15e40076be92d2129e7) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-002.md#canonical-0726e8d419168462037ee09173dcdac70e0457a08fd9694f9d627fdc45204534) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-f63430d9a25d683a3531ade2dcd32b560d0acf9d586b341b7df7b361a67e8d2b) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-002.md#canonical-a6023ac485758bd9e259bae6625567dcad1dd39a3474bf8bc4ab5a3692988b5d) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-002.md#canonical-a4d908a587095c28084e03aea02b552264088a354613762082b46258c784e397) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-6147240eb6777a7c25b1d006997b86779fcc60972afb7f9f66fb24d8427897c9) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-424acf06daa861352af20f38e4da50040488bf9afc3c8a0b04c483e6b334bb8a) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-3a795ad5a27fcaa569b2d07ceda9ec24c45adb529b744454a2093905ea786c29) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-b8df2512a4e4cb4191425215acefaac4349d0725409087bf3a1fda2dd8de8d76) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-5bedfba2e5fd073c0808e8d57a46101d6823d9aacef4cda73957e55bbd803980) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-4b8ab33a375b6219a61f8c5b80d6d5aa292d5a7e48890800c9b6ad48d599da22) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-db3a361e581b90400396e952f6767ffbced5df338b38a6ed033703a25b80f709) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-9ad87c0b698339a55247edee7ef95e99a35f8cf7864ca3e648fd87a55fee7e99) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-002.md#canonical-d18841128768d88938bd5197a29683b0c9eb4c31278bba8d45d4d67e79bf7a47) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-002.md#canonical-66febcfdcafde796630fd2fbce3c8d64bdc5e27ae5d819cdf3eb770e33d7248d) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-23619f12b49976425833a2a8092238e726fa198abdb9388737e0d2dc77923d03) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-003.md#canonical-8a5823146a74e5b16f78ec8818e900f2b648ea807eda512cfb05fb05ae13fea2) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-003.md#canonical-6638fc37717a7ff2cca8d0e37a8b05b58aa5142954e90348fa3eb7091ab90622) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-003.md#canonical-4c28ed10a0fdb3cc48675c62ee097e6be48a1a6e460e29f8ef1c4c786434d5cf) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-003.md#canonical-85dda0f52b37df43b1d05805d6abcbe65d61e1e8db0c1fffdeb2f53f35f1d8d4) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-003.md#canonical-644c0867e85aee10fc09b534b0163939a46cfd64f39d06d9c0027f3b56e29ed0) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-003.md#canonical-3966aacc75fef5ab937d421021596a1babd4e99798744d1d465b1af7e6082d94) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-003.md#canonical-6d3168df44fa5489e462fcf5f8656bcb9bc0715d824c52216e65dd1678a4015c) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-003.md#canonical-2f5c7933875a5da7d7d95715a23afd0be775fe4f97f88ab17da28baa0f51f62e) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-003.md#canonical-2c6d73b4ba6466b6cd5c06b0a9aad17a754c244c474762c26c1f4f027d26372b) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-003.md#canonical-1ddfd8d690466743888cbfd534408d3d349389cab7563fc7c3482b462dce30ef) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-003.md#canonical-2d50b7c2d4f0d90438bedfc83221a3d13cba9d3d13e573555d76bdc8b411bff8) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-003.md#canonical-8c3aeb67dcdf2044b0d802b05e616ee52d5d64f8322c11b13a1a67b3c7abf404) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-003.md#canonical-a3cea27a2124b3ff9b4f5e327dcba320f12b0a9e1345b28b41e7c5c8f9f9b1e9) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-003.md#canonical-a02b9a539bf1bfb8b5bddb8e8b8e792c91d78553be0cdc0e9019318aa3a85eea) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-003.md#canonical-c87048f7a4faa4667aaa7df09114b202a22b7bababc53c2eae4bdbc148f559ae) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-003.md#canonical-6622981fbb19b45be7792d14afb15b4ad24c9c7482a79ad189209882540a5d13) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-003.md#canonical-148b7101597e7fd711a9efe6e028bda4b4767c697d8006cc3bffb3aeeb91c87a) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-003.md#canonical-c069c83cede53e4bfdaeb8dd58bc0758ff1d1c56245ecb9fe0f11b1b3baf4925) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-003.md#canonical-aa52e29c6aa845c320d3f117689868301b79503341ac7bec64e0cbd2650b6e9f) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-003.md#canonical-3606cac84234f2da507286e77433f8aa4395bf7f3d58942642d9ef8de3a39bc1) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-003.md#canonical-78ae59e16b921337526f185e90ecc4082e2bfe06e03ecb0f0ef00f15e7bb064a) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-f3b5cab8eb23ee10379e96141abc27f1be6038cdbe4e295fb5122e5962313cb7) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-f1fe6da1e64b3fc570bd272a8eff63727beb83bb4e8ffe48d526bb131f130417) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-1560cb8a7b154552cf8be833a649b129657af19523f3a049327a9b7bcc5cd213) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-c85ba01a5a27af2378d6bfb53da2e607a7dd26de808a57a058e5a2f70746cd3c) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-eea9abb573604bdcacb984f98b603e7860950e9b5f66c8b4e58e086684774c89) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-a0cb25254cc11d0dfae190df27ac7cd0a9a74e31ba41b278b8c4ec9f0d824987) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-0fdf06e833d32d59efc1f17db32c3ec5b7dc3fa140ffd6b6733e916039fec238) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-89fb1b8cfe7022cf27f9c9a4fb8b91fc43234ff023749b6e7c711b159e142583) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-003.md#canonical-ce53b2cf825dead8609f0a080bbca8a6062e954b1cdd15d7c1bf2c9cbd3fbcb1) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-002.md#canonical-378bfb2e13ad348ca2130cd877787fd1bee9ad815a42e0a4756ad0d0f7c27966) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-9b45bfb5cd8b1edc14c3692abd6788904e4cdbab69a5a765c19c331c21eefac0) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-003.md#canonical-633955a648f946eddcb5f5891ddcd7e279c46582e6631209ed6a4aa3669af483) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-003.md#canonical-c5a60c5e268b4556334678fca543cabb61e13bd7346725be68363cb38e59766c) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3d8106883166c71283b4641e16832c963bbe966e83330e84f294da2435d360e9) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-cf0267b745215e861db350912996ea2a20d87cb632145a43d5336431d97060cb) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-7fd69dc60c61a9755fd3f03b7c6c42a557cac3275cd15a13e62818a9029637d9) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-f431514af915c59fd778a705ef82d15f68da48540e9ad06f96570a4cd4249838) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-2ac2dc3d1e23337cdd2983c98750e76e7ea2e1c6655e063c0429831e01dcb460) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-0641b13ed20e35b586eb5ae4097d94e769923b0a4b2cce8519ed8db7707b3478) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-7f8214424fd1d14739fca7ea04e9fa57621c07677f8f76e53fe972ea69c817fe) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-8e97f1a028b735e2b500a159526f242040f9d3c06d0678b43e0f66ec87b0ecb5) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-003.md#canonical-461c8c4ba07372ed26ec23edcb06f61e58c39f90c5aeeb43c1c47a834a578453) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-002.md#canonical-ae4b0926ef780a0187ef11b23934bd4c8099ffbf9335a65a7c97cb67ec787fb5) |
| `dynamic_proxy.https_proxy.tls_params` | [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-b067021ddecbfb7bdea8f15c0a3b7a8284e9a34a0ae0290fd5d7c0dd51ca5987) |
| `dynamic_proxy.https_proxy.tls_params.no_mtls` | [dynamic_proxy.https_proxy.tls_params.no_mtls](data-sources--proxy--reference--group-003.md#canonical-42af0f9c6c4760dff3dfc44b896dab5bd158e8bef0a1969fad9b1f21bfde5c09) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates` | [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-42d8386e8bb8cfdbaddd1739bb463a8c97b68920cfcfd2fb56c2bf9c2db24bf6) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url](data-sources--proxy--reference--group-003.md#canonical-707aabad6b99b0c91598a703fbe35984f0078671d0bbb18d442e717157b9620c) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-dfab37429ab384264695434af4eacb7ea3ec602fd7f99ea521bcf0cd15370eed) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-8b5223a44c9f4856da39b7e485bb276d4b4e16af75420f10dd8e7573b4fbd63b) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec](data-sources--proxy--reference--group-003.md#canonical-acff6704caa9c7a62da8418609d11f1799fd78d81eeb9f4068ccddefe969e4ee) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](data-sources--proxy--reference--group-003.md#canonical-d629a47b222684b0b131298c7a8a6d322e61f4d52bc455d547eb32091ab1b387) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-a16adc460bf76dc6d51b3076f70dd303a253ae5147ff072b426a4997fb644415) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2885659e2ccec127a1444c1514c5829a471a72df9e6aed9b4e95665eca93750c) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-003.md#canonical-17fb44d1a11e44613a4219041c341a1e51340d9511ac5cb910da5bac58613a25) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--proxy--reference--group-003.md#canonical-f5a1eb02b8e818ff6f0879019ef659b2801145b3096d12f1538047725605efe3) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-003.md#canonical-657b5db8c01479be8dc59e5e3dabc07479a39baa3fa01a4b7109bd20b5b42421) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-ae8df8a1ffa7d9f5c9be4f5af0ae703029b8e4564ca9ef33f87aa2cf646bf14f) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--proxy--reference--group-003.md#canonical-f8c38042dc37225f773796e3ba2b30a3be6e82e22a44b3d1a7b213a25659f80f) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url](data-sources--proxy--reference--group-003.md#canonical-484083cb58afedb6d0343fc04b1f9dae4b3759c58c2820e0debabe73350c7575) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](data-sources--proxy--reference--group-003.md#canonical-ae860c963855fe3274c28390635a3fb56036a481d2ff01f39b612268ed1d89aa) |
| `dynamic_proxy.https_proxy.tls_params.tls_config` | [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-af5a2684c459e089e47d286f12f72cfd63926328a716038a96807c8ebd3f5ef7) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](data-sources--proxy--reference--group-003.md#canonical-1133624d883253ef3a638b3b95e2b0d175867ef4a97ca00765d95d485c57ea3b) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites](data-sources--proxy--reference--group-003.md#canonical-83f88582e4da8c4ac10b1cee1d1f69d6723c96d140e9de475bc50cb917da6054) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version](data-sources--proxy--reference--group-003.md#canonical-64bf7e3e59708dca8af524c0f133393cea5ad04404bcf5b69920a554c735cf2d) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version](data-sources--proxy--reference--group-003.md#canonical-4a54c3c90ce94163997f466439ed8f2326fc5e12616dd3efd7ceb2cd2dc185af) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](data-sources--proxy--reference--group-003.md#canonical-eb32d43886558355b1961365feee0cc32c405365af7628dca8b4f228879e68ba) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](data-sources--proxy--reference--group-003.md#canonical-68ce8d7c51684821c0bd51e3e12fc35731359211679d4afffa197cfbc92625d9) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](data-sources--proxy--reference--group-003.md#canonical-9df0e21cf558ae5af52e749c2e88eced0e3246130c25a45ddac2fee773c19c86) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls` | [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-8094bcdd06b342700a581ebabe179126907179bc489aa23009630825ff06be1e) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` | [dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional](data-sources--proxy--reference--group-003.md#canonical-535c4236f0098fcda505101f369f70c485ad377a15e5f9c5c9d0739039922df0) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](data-sources--proxy--reference--group-003.md#canonical-736ee12aaea250c78be3efbcdb63704dfec2b4eeb566144c480a308a93287e9b) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name](data-sources--proxy--reference--group-003.md#canonical-1f31c59b664807f9359d9ef1bb3c06aae52e4dd22380fb6b6ddfd31f249efad6) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace](data-sources--proxy--reference--group-003.md#canonical-eb925a76580b2709b8ec8d451e843ceeb77d47c281303c27a7897b0c3c6ba01d) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant](data-sources--proxy--reference--group-003.md#canonical-00ed4f54eeb6784feb43ba30b72f435c24433ebcd07d0bf926680db1db201593) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](data-sources--proxy--reference--group-003.md#canonical-1d5d8f4204341060140c316b98944a4c539d90037f1b55c268daa1467046c8e0) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](data-sources--proxy--reference--group-003.md#canonical-6004d29c78068ef23b4b6776bcc579f738cbe8f02bc30e0060529e928d3a1917) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name](data-sources--proxy--reference--group-003.md#canonical-538bba0f91979556995ffc0b94a93386e342760fc55fa003c8cb426d0993e569) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace](data-sources--proxy--reference--group-003.md#canonical-73fa39338b949005bd55b2380cdb37ebe93116b543949efd83650b92ef8141f7) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant](data-sources--proxy--reference--group-003.md#canonical-87a0a45602650caf16c7dfa58a96d168f36e5e1eae19653e7256b6039655be46) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url](data-sources--proxy--reference--group-003.md#canonical-e42e2b96263f7800c955c967ca54a27f397ad29cbceb881211e6b0c307f3b67e) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](data-sources--proxy--reference--group-003.md#canonical-622b27d83001abc7e82c39365582a3eed90227281a6f40a7416684e18a8dce45) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](data-sources--proxy--reference--group-003.md#canonical-b47097ef7ef8cba77035cc467d9fe4ac4fcb82a90d9f267b79b4de8856bd3f15) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--proxy--reference--group-003.md#canonical-a305e34a808f53687a6d1fbff7c2192bdfcd8c158a420685fcfed6ffe2c5298f) |
| `dynamic_proxy.sni_proxy` | [dynamic_proxy.sni_proxy](data-sources--proxy--reference--group-003.md#canonical-ced420c30de005f32dd0a9e8406225fe3a1efacfea7c72c5f3cec863b11d766f) |
| `dynamic_proxy.sni_proxy.idle_timeout` | [dynamic_proxy.sni_proxy.idle_timeout](data-sources--proxy--reference--group-003.md#canonical-451214f5661a93da1ce4ef7ec159b06ff088ce9d731f8031b63fd089ebdaf053) |
| `http_proxy` | [http_proxy](data-sources--proxy--reference--group-003.md#canonical-9910ac7670bd922a7b6653551d7d805c0038d46e405c6aa3a7c9c42da4befa4d) |
| `http_proxy.enable_http` | [http_proxy.enable_http](data-sources--proxy--reference--group-003.md#canonical-2c56d4f576a38932bc9a158af417ca94de7501ac80d165750c092722b3967127) |
| `http_proxy.more_option` | [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-61f3cd2690b2d520de5b40bd04cf43337ce4ac8694a122013dc205581d7da82a) |
| `http_proxy.more_option.buffer_policy` | [http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-004.md#canonical-797be26098d1bd1213e80419c0319702abe14660d6740170b2a985d317acbe07) |
| `http_proxy.more_option.buffer_policy.disabled` | [http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--reference--group-004.md#canonical-37a644aae40d4b47f0f74387df28cb17dcced8a32c3ad7d5c29801db267d7c68) |
| `http_proxy.more_option.buffer_policy.max_request_bytes` | [http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--reference--group-004.md#canonical-4a7be4dc905ed483af2e6b607c87989eb1c9076511d88b286c1b443b77000dda) |
| `http_proxy.more_option.compression_params` | [http_proxy.more_option.compression_params](data-sources--proxy--reference--group-004.md#canonical-373d60d59d7778b43434e690b39787abfd4ee67b38f6a9908b41a0a51dc7f673) |
| `http_proxy.more_option.compression_params.content_length` | [http_proxy.more_option.compression_params.content_length](data-sources--proxy--reference--group-004.md#canonical-7a702f962cb9b53b2e19df7f105db02feba197cbe9f80cfc2a62e8280494c062) |
| `http_proxy.more_option.compression_params.content_type` | [http_proxy.more_option.compression_params.content_type](data-sources--proxy--reference--group-004.md#canonical-4d234b4bba106a581acf470a4c13dc585f3b44404bd6266015524b5a44718535) |
| `http_proxy.more_option.compression_params.disable_on_etag_header` | [http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--reference--group-004.md#canonical-ce0ce0f749dcf02fca28d7595c6611049806462053fa5841da955f38c912abdd) |
| `http_proxy.more_option.compression_params.remove_accept_encoding_header` | [http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--reference--group-004.md#canonical-ae89c1a8ffb7bf4189a7745afbf7654933347ff445e021975a8b1fe1a25f7e67) |
| `http_proxy.more_option.custom_errors` | [http_proxy.more_option.custom_errors](data-sources--proxy--reference--group-003.md#canonical-bf0ad27d6bf2abf711aeac18a1df2b108f5f19a815632d9e559540c4812e6518) |
| `http_proxy.more_option.disable_default_error_pages` | [http_proxy.more_option.disable_default_error_pages](data-sources--proxy--reference--group-004.md#canonical-3af4f345df0380e1dad3b9e42c99bd2f4af2421a5e839046b1a40a6c2ab831d3) |
| `http_proxy.more_option.disable_path_normalize` | [http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-662018c04261d03c5c0a2bf828570ad1c8aae2ea18d569f4be5dadebd64ae6d7) |
| `http_proxy.more_option.enable_path_normalize` | [http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-c874d73b43fff1c70dd958ec34f9705fa22d506f993cb3d4f0d79e1856be04ae) |
| `http_proxy.more_option.idle_timeout` | [http_proxy.more_option.idle_timeout](data-sources--proxy--reference--group-004.md#canonical-ec43193bbd4ee38eb6271912b3b7a00cc7b571183f62322f3409b5899e24cf03) |
| `http_proxy.more_option.max_request_header_size` | [http_proxy.more_option.max_request_header_size](data-sources--proxy--reference--group-004.md#canonical-66d4b7c92424dd6df0d03fb2a42cebc43f5da119ebacac7dd7af3fef9a763e6f) |
| `http_proxy.more_option.max_requests_per_connection` | [http_proxy.more_option.max_requests_per_connection](data-sources--proxy--reference--group-004.md#canonical-0dbd9fe9f8de2a836e706451a48427f57311234ee1be16c5c40f240e6d312f6d) |
| `http_proxy.more_option.no_request_limit_per_connection` | [http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-004.md#canonical-d4fa92c3041899578bc60733ea6251893d9961f7d7baf110a4b2bb5c24f6a79f) |
| `http_proxy.more_option.request_cookies_to_add` | [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-5714fe10d064e6d9ef3591d24d71eba7e26ffc2ab1871716bb9b63e3efa2d851) |
| `http_proxy.more_option.request_cookies_to_add.name` | [http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--reference--group-004.md#canonical-cf03af0067ba07f8656b10c3865f0b560497181c88d3e52f00a8987d494aa0a5) |
| `http_proxy.more_option.request_cookies_to_add.overwrite` | [http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--reference--group-004.md#canonical-44ba8c83e88a037bbf0a76c4554d0ed20b94e2cfcec087b5793ad33b354d59dc) |
| `http_proxy.more_option.request_cookies_to_add.secret_value` | [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-233cecf31b6368dc87b82d23ea2c093b5d2b7211198027e02d70f78a3c497981) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-203c60fb07e9cc10851474628cab27577ab690ffbd06cc2130c4f7c324396257) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-1c07dcddffd3f8d87a98a9dc33481ff14df02968150ec78a31c550378a4a1c02) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-45f84cc367d99979a78ba877736fc243e60f80bbbbb96a7c909cf9a383b7f05b) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-4ad3c6f8d321bdfc0a526101c729bfa062b0c4b0605b7057f329d6bf2f0e8031) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-0092b1585cbe918a136e069b189f33b14bac367056ff71efc9ea4396b1b4ac9f) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-01330ea03709f123ad56298dee474aa9bfa154a9f4d1b75c1f1a17e141a3e750) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-ccce5965f4466b2f7bf20869142ce84294cdcc4cd9c37109728c2bc26d01f2e4) |
| `http_proxy.more_option.request_cookies_to_add.value` | [http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--reference--group-004.md#canonical-fc1ef5716282e7ffe0ffaaed931b64870f2bc83773a772f675f2f560caaa9110) |
| `http_proxy.more_option.request_cookies_to_remove` | [http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--reference--group-004.md#canonical-7f5a175007c7c16464d2c5fe5e02e1beb84d22ca9f85089e479e3c577183a7b8) |
| `http_proxy.more_option.request_headers_to_add` | [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-064d597b3f5a05dd4e0b1f5c8b99db866ec1d6f375f3293a39f3b72ddd75d14e) |
| `http_proxy.more_option.request_headers_to_add.append` | [http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--reference--group-004.md#canonical-7b0c6d4b245e59af2a8e293f1668164290010e6e725b15a0baedee76e6de2f80) |
| `http_proxy.more_option.request_headers_to_add.name` | [http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--reference--group-004.md#canonical-5de0022c006fa4073802c032ac7e6bcc7ac46568cf7b8ba11e51027f2b48f82d) |
| `http_proxy.more_option.request_headers_to_add.secret_value` | [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-067fd92d4f617afb48c258987785c620a695029061af534250407947ce61480c) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-100c6d6207a86d34c896fbc10f521d8ad81151ce7969178f96e4c1672a1f5763) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-05d4afcc62eda94638a388400566fbcda75a9d8d26d5287ff8b074b13ba48f04) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-3017d7c4e22f40fb3aa439113764ce488fffeac36c13ec626de23641a8379a51) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-f82fa1be55e65a8b038bb2efa5b3f4085524f2ea82f4eca694f42dc04ff2b76b) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-2194f960e0d3a166c57ecf2dacec7788b3a5696a48cb266d434cdc4bbadcecb2) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-0499a2b89e7723a38768c9e7cb7afb55d694a60ca5bcd69ad69b0422c1d1f552) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-6fe0a207e2bfdb7731df37eb1ce350fef6c8bab6ae2c521c3517833f1976b468) |
| `http_proxy.more_option.request_headers_to_add.value` | [http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--reference--group-004.md#canonical-e7da28cfc6a0b4e46d69f6a8842d3402f6cc2008be2c488c277a72a5ca95be66) |
| `http_proxy.more_option.request_headers_to_remove` | [http_proxy.more_option.request_headers_to_remove](data-sources--proxy--reference--group-004.md#canonical-c2f1f2d369e4613e929115e24307a991330dfc8367c7510aa734e136df461220) |
| `http_proxy.more_option.response_cookies_to_add` | [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-606ce31f751ada6c13e19bd0f50712c892514c6d47d97683e3ef0dd646ec4b77) |
| `http_proxy.more_option.response_cookies_to_add.add_domain` | [http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--reference--group-004.md#canonical-97b4604d5cc755ee1a53d9f8ee9388a0abe4a7bf90dcc2e3ffb48085ecfb8fe3) |
| `http_proxy.more_option.response_cookies_to_add.add_expiry` | [http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--reference--group-004.md#canonical-2431521ae261600884774a7152b68e447e4999d2246d4945d555a065993ac3fa) |
| `http_proxy.more_option.response_cookies_to_add.add_httponly` | [http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-004.md#canonical-7d08a8fd4fff4e0f81875904393935656fa3556b704ab6d5483101426125b0fe) |
| `http_proxy.more_option.response_cookies_to_add.add_partitioned` | [http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-004.md#canonical-08209eb1bb5bc53fde378771cc5a3a120d8442c3d4bff80a29f3bd48982e7d4e) |
| `http_proxy.more_option.response_cookies_to_add.add_path` | [http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--reference--group-004.md#canonical-373f164ec3eb4c79ab5acdb75e8d93562fc867b3ecbf0d051571f0ff38f25c7b) |
| `http_proxy.more_option.response_cookies_to_add.add_secure` | [http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-004.md#canonical-c5a604f16a90387d354be6183e10d67ebc34c0ba92caf7435d454269033a532a) |
| `http_proxy.more_option.response_cookies_to_add.ignore_domain` | [http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-004.md#canonical-5574b1e3163d905c719a6a66724eb1afa8fb921c6f8cd73ca1d0a47951e9a63e) |
| `http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-004.md#canonical-968217cd0f38d5f0175478c695cb6a68bc48c3096bf118b83deaa8278517859e) |
| `http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-004.md#canonical-b87a9f030516e794cdc783b7ffd408f200e1585544c48ead55bbeb66b239f609) |
| `http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-004.md#canonical-65e48b18ba1c33277ab226bfa78e4a167cfc954ba239c9dfc24c00a769cb28d3) |
| `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-004.md#canonical-1311df8cd4adf5e3c517a62d056b9ffab03dea838cd48ec285bf9ceea3e3c65b) |
| `http_proxy.more_option.response_cookies_to_add.ignore_path` | [http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-004.md#canonical-7806ae27347760bbaa8fe8b6070d4158f96b94451041dccbbdafb485857ae3ba) |
| `http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-004.md#canonical-fc446e341ea494d1d3da7d78ea75d8eb523e50e84f693c2cc4117f1d0c7e1456) |
| `http_proxy.more_option.response_cookies_to_add.ignore_secure` | [http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-004.md#canonical-3006de3e9e1fec4acfbc33100aa82f4386e8895425fd7c5abd1254d06417eaa2) |
| `http_proxy.more_option.response_cookies_to_add.ignore_value` | [http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-004.md#canonical-c9399f9b4bbea3db90e48be592b52e73762861f138362c0978f4cadb3e17e1d6) |
| `http_proxy.more_option.response_cookies_to_add.max_age_value` | [http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--reference--group-004.md#canonical-ec43c45e4d102d71d3c1bd219bdcd45dec0bd3912b0dc7af7901647f3c0c4a0e) |
| `http_proxy.more_option.response_cookies_to_add.name` | [http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--reference--group-004.md#canonical-a3ad9b64684aed33e7cb929b0f7bc7ba88b247bed4c8e459052a1f6b81bd36d1) |
| `http_proxy.more_option.response_cookies_to_add.overwrite` | [http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--reference--group-004.md#canonical-cefe23f09ddfaf006c3008e79d1f3b6b2b7040be733bc585c2771638ae96d893) |
| `http_proxy.more_option.response_cookies_to_add.samesite_lax` | [http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-004.md#canonical-10f7676365a1ed7f92205b8952c8c861c68c9b31dfe326ee259eafab95d3ee22) |
| `http_proxy.more_option.response_cookies_to_add.samesite_none` | [http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-004.md#canonical-e55b8b25e285e522c9ac718464b87a0c3d5d822f6f81d9379e64f6ac0c118cf1) |
| `http_proxy.more_option.response_cookies_to_add.samesite_strict` | [http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-004.md#canonical-f4c7a0f2084d4f30028be9ae31192bc20d89cdcc5eef83cbe12250d10888fba4) |
| `http_proxy.more_option.response_cookies_to_add.secret_value` | [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-3c6bd7fe5e5aabd1986d3620a8886e3d70c48affd1e2089314bbb022239a182c) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-56ca5041376353dc6292d3fb0eb2d13cf8e4dd29af5ed22213a9b060b3899b88) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-409cc1d0c056e031b8e662e0cbda2b8c5e24662351c456e5f3f7d86d407e4b94) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-0449e63f6dd52d5537800600931d541338a8a44ec46165c2f4d851b0befe470e) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-e6b18a14d4a6d573262ed7d32d10b9f1a6a05297925162e3b232178f663d1874) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-bce1d7c3a37015c5fdc5163c6a7f1d66985309cebf5c3190d44664e8d26fa3dc) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-4e04bfcad35014799d972d6a7f565366b619a99a10132a6830c1a026d885cdc6) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-8f762dc3e39543ac7fa7efe104fbb4da91195c75a0c1641f29ac8d4e8b7482d0) |
| `http_proxy.more_option.response_cookies_to_add.value` | [http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--reference--group-004.md#canonical-2a31a9c9b667f41f8c1d9e9d3b62547e1ad1b5334035c60d6c96f0a476592325) |
| `http_proxy.more_option.response_cookies_to_remove` | [http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--reference--group-004.md#canonical-97687c194baf78dbd355870bf736b97c9fac3001fdc402037ee63342944c7b1e) |
| `http_proxy.more_option.response_headers_to_add` | [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-b04d9684c1c4f12d1f9ffb66d166cc5a7c6ec270742a6d588d7735649e528dc1) |
| `http_proxy.more_option.response_headers_to_add.append` | [http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--reference--group-004.md#canonical-f7197d71e7516cacc49090c1a933cc4e92acdd904746a295b904dad5dff1c064) |
| `http_proxy.more_option.response_headers_to_add.name` | [http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--reference--group-004.md#canonical-1beb3234f82b5a81312aa53bccaf850930e7bced68f50106c75cb364ae8a0121) |
| `http_proxy.more_option.response_headers_to_add.secret_value` | [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-c129bb9958164828e280461cbb6f48105958856a5735157b575c323f31ef83f7) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-d096c1d687241792d52f6e08f1b5c5943911cd26e7a91a3cc983e4e5f12c52e5) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-004.md#canonical-b6bdb3fbd4d7254a16444072e589bd595953e5452e21615db94bd36453defc92) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--reference--group-004.md#canonical-46db97ad85f2dcfe32d4b8b2834d1774256c3ad7cd04d48fff4496861eecbaec) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-004.md#canonical-5a35879f8e4f9c3600281a792c9bed88c9e84146b5732e6dfc5ef15f831e975b) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-1309170c61368ac41bd54fe8e8f452f19e6f8d3346abaad62c9244961b664425) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--reference--group-004.md#canonical-0056204a972b55057afc6b0b90b45f2dae15d41697f22ba6d58f264c47962e0c) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--reference--group-004.md#canonical-2881c4bdbe7c8d28912788467589edc1e151c3a96de5e72dcf42f7432ff5cbad) |
| `http_proxy.more_option.response_headers_to_add.value` | [http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--reference--group-004.md#canonical-33ba97422707d69d098c749818e1b4920354f787482d1ef729830dfc7469845c) |
| `http_proxy.more_option.response_headers_to_remove` | [http_proxy.more_option.response_headers_to_remove](data-sources--proxy--reference--group-004.md#canonical-226e79e4b1ff6553c7343ed9f1f691234a61105003385001794abac631b4c97a) |
| `id` | [id](data-sources--proxy--reference--group-001.md#canonical-5ca4138f4c316fe75b28cf0175c46c829813170ddf3e67c81c857b0e73894152) |
| `labels` | [labels](data-sources--proxy--reference--group-001.md#canonical-6770d01aca12196e3bd7f87f64e86b3cdeed9429237a6b93cf1f54ae1426c5c7) |
| `name` | [name](data-sources--proxy--reference--group-001.md#canonical-d3bf05cd1e46d7760d249c8f3485484f01c640eac6cec6df5f0f711d64d87c75) |
| `namespace` | [namespace](data-sources--proxy--reference--group-001.md#canonical-7b659d3d701f61890141fbaa8face813bd5945157e0373ddb8f31c476bd9135a) |
| `no_forward_proxy_policy` | [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-b4151087faedbfb9198d1954ec55c0ac5c869e873ff01d3dc8c4c316eac84f12) |
| `no_interception` | [no_interception](data-sources--proxy--reference--group-004.md#canonical-472d438ff2425ac3470397b8ccc7afe368e24274f0df9164bff253696d9d3e23) |
| `site_local_inside_network` | [site_local_inside_network](data-sources--proxy--reference--group-004.md#canonical-c5fc7b71f5016f4b6f581708b16103c722f0351bd6c961311178ce76691801e2) |
| `site_local_network` | [site_local_network](data-sources--proxy--reference--group-004.md#canonical-98d9bb15e9aa2e43fb2f0517b15cb21a33e860eb2bd9b73f722d31e9a62f5399) |
| `site_virtual_sites` | [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-16f7c6653a49b2c94edcdda8f045fd6277695fc329d537369a57aa90a3e7193b) |
| `site_virtual_sites.advertise_where` | [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-60ab02db3202b6f59ba79bfa44fef9a6b98b2c3a86e664bdf34d4c10f3e18e16) |
| `site_virtual_sites.advertise_where.port` | [site_virtual_sites.advertise_where.port](data-sources--proxy--reference--group-004.md#canonical-43c271745d0ad1f68128d1f6199146d06460b1b097c25228700d1c8a07a7a43d) |
| `site_virtual_sites.advertise_where.site` | [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-004.md#canonical-86b9b90e9c0653df335ad2284dcd2ee45a51f6baa871260a86abc18f85c3a78c) |
| `site_virtual_sites.advertise_where.site.ip` | [site_virtual_sites.advertise_where.site.ip](data-sources--proxy--reference--group-004.md#canonical-7277efd852e6f788b2dab7eb2c8565ff8b29102bd8849643985887d7ce2ab29f) |
| `site_virtual_sites.advertise_where.site.network` | [site_virtual_sites.advertise_where.site.network](data-sources--proxy--reference--group-004.md#canonical-5afbbcf4f6bcc13fa9401d812a80e1b977b8080e60ffd270427db09109ea17ff) |
| `site_virtual_sites.advertise_where.site.site` | [site_virtual_sites.advertise_where.site.site](data-sources--proxy--reference--group-004.md#canonical-865e01e4b5ab47086ab0d6c006b7412356adfc7c7c2b80335d9cdc8d7059cb2b) |
| `site_virtual_sites.advertise_where.site.site.name` | [site_virtual_sites.advertise_where.site.site.name](data-sources--proxy--reference--group-004.md#canonical-06ad849fe759bf38a7fb5b83bd5cd20be0d38e2b57a70f34e94f54e0bf6468fb) |
| `site_virtual_sites.advertise_where.site.site.namespace` | [site_virtual_sites.advertise_where.site.site.namespace](data-sources--proxy--reference--group-004.md#canonical-8f6a9c2d16fae0e6a960088009cfe057916bbeb5a7d3bdf94bd81a709fe37132) |
| `site_virtual_sites.advertise_where.site.site.tenant` | [site_virtual_sites.advertise_where.site.site.tenant](data-sources--proxy--reference--group-004.md#canonical-7d783636b4a51589d0fd7fd3f783c0fe9fe9bfc076d99f4d9daba22641cdde4c) |
| `site_virtual_sites.advertise_where.use_default_port` | [site_virtual_sites.advertise_where.use_default_port](data-sources--proxy--reference--group-004.md#canonical-d7bfda3ac2b63892d619223bc212a3574847654bc8eb079baebc35ed5f94335b) |
| `site_virtual_sites.advertise_where.virtual_site` | [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-1ac96376af346de2aaffa5294e4160946b7624b783ebe5b2858ef6701d115231) |
| `site_virtual_sites.advertise_where.virtual_site.network` | [site_virtual_sites.advertise_where.virtual_site.network](data-sources--proxy--reference--group-004.md#canonical-2a60f5302a604100d643dd0a1b6cae6d4a603f278113a9f4a6a13aecf2f2daaf) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site` | [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--reference--group-004.md#canonical-e17c031195aa8d766b737be803a5d552c1fdfbfe55dfaeac8ea88f22e2c7863d) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.name](data-sources--proxy--reference--group-004.md#canonical-61ccd381416a92fdf80db8d75e635aa63aff6593106ccc0970a9fc16ccb3ef4c) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace](data-sources--proxy--reference--group-004.md#canonical-a63e4366f453758acb00762d9f0c4e7a356907f9f9e8d36885f09a1005d2b8da) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant](data-sources--proxy--reference--group-004.md#canonical-3e269c202ab856d32feed39266d976c96f39aebc3e0e483a7f1ba0aace24c741) |
| `tls_intercept` | [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-99bb678c4909c834195d7cefc40f5c76cdae7ac5abe2820b0fc8a25451848886) |
| `tls_intercept.custom_certificate` | [tls_intercept.custom_certificate](data-sources--proxy--reference--group-004.md#canonical-ee865bcdf218dd5bc12859c7cc11ed1a461708d498845c81b6fc2a6e87e79fb2) |
| `tls_intercept.custom_certificate.certificate_url` | [tls_intercept.custom_certificate.certificate_url](data-sources--proxy--reference--group-004.md#canonical-ce5c1057088b0ffa89916b33330d2e03323f0266384f98702efc6eecf4f083a4) |
| `tls_intercept.custom_certificate.custom_hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--reference--group-004.md#canonical-1cc1628f0cd934728afb3199ab654f3c46a3beade4baa8572bc873dc1ace6fc9) |
| `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--proxy--reference--group-004.md#canonical-e0ef1598d1cd8415899dea4be66554ef4a1205fb9c87c4f5a212f8468f535c95) |
| `tls_intercept.custom_certificate.description_spec` | [tls_intercept.custom_certificate.description_spec](data-sources--proxy--reference--group-004.md#canonical-88a01b3474ff2c4e16e277479dd5d0636f32168e2936ed51bdfbe1098a3d3f51) |
| `tls_intercept.custom_certificate.disable_ocsp_stapling` | [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--reference--group-004.md#canonical-e1ebbdb2d1d9dda88a86fd700968497de1b3fe1369c448835014e64ce977e661) |
| `tls_intercept.custom_certificate.private_key` | [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-004.md#canonical-7951ef69fb554e3ce16cab4e48e30838205179d360a4659305092e9bb5d7615e) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-9c6019608b52faeb2a90e3946f800ba6b5c42ab618dd32cd402002aa5b31e6dd) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--reference--group-005.md#canonical-df23aea7bffa0d5e1cd368399a94be42625cd91623358525f1ae5749748de698) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--proxy--reference--group-005.md#canonical-f0323c83c9282e20a6f18e8dcb121c749b182f5d9fd0d95aad83869942df13ba) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--proxy--reference--group-005.md#canonical-fbb5e1c961f78125244d0d8fc5073f391e7e116f730c0923b1076b6b9c7c87f7) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info` | [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-68bccfbb040ca55c6936c712c969bcb9a9aab7fb1f04abb189a5191cbf971e02) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--proxy--reference--group-005.md#canonical-8b18ec08dc71a8ec20569200c95edf01af504c7e23f48e4df0c17b9f631ebaea) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--proxy--reference--group-005.md#canonical-c3fabdadf1ee11f0618330660c453f7d721785f3639de7b8ab2e654379de8e08) |
| `tls_intercept.custom_certificate.use_system_defaults` | [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-e44449250351256d37b8fb1d78e992d925bdb943135a3d39f8548a6c17a97621) |
| `tls_intercept.enable_for_all_domains` | [tls_intercept.enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-7c29bb7cba68e6f1366727e73adec2130f869b775fc6dc3b061371e27294af75) |
| `tls_intercept.policy` | [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-91367051670db5eb78bea3f7cf78a1486c6a3990ec69f7eee49d4d4ca1d05264) |
| `tls_intercept.policy.interception_rules` | [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-573829d6fa9cb7a6d326732b36c369af33adbe5cedf9f370e6fb889971ef8949) |
| `tls_intercept.policy.interception_rules.disable_interception` | [tls_intercept.policy.interception_rules.disable_interception](data-sources--proxy--reference--group-005.md#canonical-ecd35760a2153b535ff3a9e6c55b8118b7bb39c052eab1226a1afb742738f03e) |
| `tls_intercept.policy.interception_rules.domain_match` | [tls_intercept.policy.interception_rules.domain_match](data-sources--proxy--reference--group-005.md#canonical-809662283b159958ca30107dbf4d980050faaef8b39c3bf9d084a5351eba0d59) |
| `tls_intercept.policy.interception_rules.domain_match.exact_value` | [tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--proxy--reference--group-005.md#canonical-d97971ff85147bb7c85d4db406cb3b23321cabc5e43c8f73928114ec63c623da) |
| `tls_intercept.policy.interception_rules.domain_match.regex_value` | [tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--proxy--reference--group-005.md#canonical-6452460b530b3855535d152fdecfc14423c1721feabea069acd26d5790d8e1b8) |
| `tls_intercept.policy.interception_rules.domain_match.suffix_value` | [tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--proxy--reference--group-005.md#canonical-44899ebbc98d6ae2972ed59388b62873928e13304739e72d70f46570b3259c25) |
| `tls_intercept.policy.interception_rules.enable_interception` | [tls_intercept.policy.interception_rules.enable_interception](data-sources--proxy--reference--group-005.md#canonical-5022b4b0baab045c96d34bc868b54831c813c7ec712bb86f786979f1b4a9af17) |
| `tls_intercept.trusted_ca_url` | [tls_intercept.trusted_ca_url](data-sources--proxy--reference--group-004.md#canonical-127c8112a718a6ebf430734723e5cc76a264429ac483ab7669aa621e8f60a855) |
| `tls_intercept.volterra_certificate` | [tls_intercept.volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-ee31197a0753d7526ecde5972f9533e04f0a9eb7a80f672bf1a7bbbefd3aade9) |
| `tls_intercept.volterra_trusted_ca` | [tls_intercept.volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-f4c3e1060dc323eb83f9baeea12eba19b7f8710a12497c25cb732b09f41cc7f1) |

<a id="canonical-33f852e2ea7a1d6f2a50665731dd8b879d4a339a70829588805c104d84146595"></a>

## Next pages — Property reference / 51550e4f53ec / 12

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-a75c7c370855a172adbe82000d2d5c54aaef55974e04c0d0f1fc7e451b8b252d)
- [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-6ad02a6a79aced4d88d3b8c364f09978072288f3e3db7685c0090af711fc8703)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-7aac5d7097baac0cba3fea84273eec6641ab91d70f6c0f5bd1aa27f10520b5b2)
- [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-354eed661a275799f673f7dd3b219dd0d9c148389186b9bd151c99738179e6c9)
- [no_interception](data-sources--proxy--reference--group-004.md#canonical-4cb287101d7cff732a91ce780d465332c4fc488b430e1c588c72d8d91fff4f3a)
- [site_local_inside_network](data-sources--proxy--reference--group-004.md#canonical-511ab4dca177ca5d7d292f2efc016dc5ab5e4f01c63811a03340a2195219d67a)
- [site_local_network](data-sources--proxy--reference--group-004.md#canonical-7772efa695c7a5d2abe78ba96890f5435cd65223acd8167714b620ad5f4bda05)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-4687afcca15266478649172c20cd09ff22d8a1f08844927978877710f01e2a7c)
- [tls_intercept](data-sources--proxy--reference--group-004.md#canonical-ecd6c2a1efc573c71ddf03157d70b0e80cfa80d9aeab2eb9e4cf87b4850ef9f6)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-a75c7c370855a172adbe82000d2d5c54aaef55974e04c0d0f1fc7e451b8b252d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9eaab7c232bf966c73122a49d71748e5376b532221aae89341372c3b75041976"></a>

## active_forward_proxy_policies — active_forward_proxy_policies / 8306b8869e8e / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- active_forward_proxy_policies

<a id="canonical-1759143cc65e63e4b0c8dcca1388f102814349bcf988373b637ba3a2b02967cf"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-1759143cc65e63e4b0c8dcca1388f102814349bcf988373b637ba3a2b02967cf)
- [no_forward_proxy_policy](data-sources--proxy--reference--group-004.md#canonical-b4151087faedbfb9198d1954ec55c0ac5c869e873ff01d3dc8c4c316eac84f12)

Select alternatives according to the provider validators above.

<a id="canonical-8f5c4f5c2d6987970c30837b61c9ee2ff86eede68e6b7390c7caa7006095a0ab"></a>

## Direct properties — active_forward_proxy_policies / 8306b8869e8e / 3

- [forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-82e1098578afca0ee9dc7b2de461a0d643be2dcefb8b0355b1526e6fef7d41fe): complete subsection reference.

<a id="canonical-739025eb3caeda7d30a6e5b49aa18ae0fe04857d445a94cf1b712f695cbe3c8c"></a>

## Next pages — active_forward_proxy_policies / 8306b8869e8e / 4

- [active_forward_proxy_policies.forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-82e1098578afca0ee9dc7b2de461a0d643be2dcefb8b0355b1526e6fef7d41fe)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-82e1098578afca0ee9dc7b2de461a0d643be2dcefb8b0355b1526e6fef7d41fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fed5125a7302f726bc5d7d5893535b7e9f7021093ba6922855e256c4290253b"></a>

## active_forward_proxy_policies.forward_proxy_policies — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-a75c7c370855a172adbe82000d2d5c54aaef55974e04c0d0f1fc7e451b8b252d)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-eb8b3bd8f3d9b962ec5f4d996202a5f6283787e8627911aa6bf0ea8fc8b97708"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-10262702c36c0221afdff7daf2c58590319a3e589fe632a3f99a1620631e5659"></a>

## Direct properties — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 3

<a id="canonical-0faa220f2ae05fa331569bcbf78ff59384c2bda93018f397c3496ea8045319d8"></a>

<a id="canonical-07aa91708c3389733aa7c659965f525d0f00bdc441e30f3694347e5dc6b08a4b"></a>

## name property — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 4

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

<a id="canonical-13008d16a77de36490e46d17dd43bf3b843954a06d5dd5c0cfcb3000af498914"></a>

<a id="canonical-92a8fc5e5ee4d41d67430e1277d7a66352b9bccbcc8bb7871c34d947e69d79a2"></a>

## namespace property — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 5

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

<a id="canonical-75432b3c0e9cc7b1b6bfcd62ee6411ac49e26f220f80e32b17a476f41c24e16e"></a>

<a id="canonical-17b0cea253cb374299e455664eb5fbd2bf28f149c11fc2d35cab2df6ec0f5b56"></a>

## tenant property — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 6

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

<a id="canonical-51e39e1dfef1e0c3e032dd7eda31e8eada991fc042468943e94ab8a9fffad47c"></a>

## Next pages — active_forward_proxy_policies.forward_proxy_policies / 460c7bf136d9 / 7

- [active_forward_proxy_policies](data-sources--proxy--reference--group-001.md#canonical-a75c7c370855a172adbe82000d2d5c54aaef55974e04c0d0f1fc7e451b8b252d)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-6ad02a6a79aced4d88d3b8c364f09978072288f3e3db7685c0090af711fc8703"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9699826e7b630ea817b6098deaf246617d659bb11b9c4aca32064992e3b57a17"></a>

## do_not_advertise — do_not_advertise / 1aec714cff0b / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- do_not_advertise

<a id="canonical-551f1846bcb96d8193dc2dad151583cd0755d4dcfc56f89f987b325ea3ea947f"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_advertise, site\_virtual\_sites\] Configuration parameter for do not advertise.

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

- [do_not_advertise](data-sources--proxy--reference--group-001.md#canonical-551f1846bcb96d8193dc2dad151583cd0755d4dcfc56f89f987b325ea3ea947f)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-16f7c6653a49b2c94edcdda8f045fd6277695fc329d537369a57aa90a3e7193b)

Select alternatives according to the provider validators above.

<a id="canonical-c88bd10b330a1edb20353d40bbe9bcc3615e6f77f7e4628a2fa5415d43009c26"></a>

## Direct properties — do_not_advertise / 1aec714cff0b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-980308e5cea8c3fbb19bd8140bdc4aafffa3d14d96fb81967098602a5d1a7885"></a>

## Next pages — do_not_advertise / 1aec714cff0b / 4

- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfeee0212dc02f27b09798336b74b1efda9d9b97b8379ce4f0bdb989979605ec"></a>

## dynamic_proxy — dynamic_proxy / ebf048f2acba / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- dynamic_proxy

<a id="canonical-f0194ffb37cd96597d0c0ebcf7db24e8e420fef0265b93dd03471664f5c2b5e8"></a>

Type: `"single"`. Computed.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"disable_dns_masquerade\",\"enable_dns_masquerade\"]",
  "x-ves-oneof-field-proxy_choice": "[\"http_proxy\",\"https_proxy\",\"sni_proxy\"]"
}
```

OneOf alternatives in this subsection:

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-f0194ffb37cd96597d0c0ebcf7db24e8e420fef0265b93dd03471664f5c2b5e8)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-9910ac7670bd922a7b6653551d7d805c0038d46e405c6aa3a7c9c42da4befa4d)

Select alternatives according to the provider validators above.

<a id="canonical-3a0d57fb7822e02a3d142893093e3644e3b7fc7cd13d9bb0f4cfc51fcef03eaa"></a>

## Direct properties — dynamic_proxy / ebf048f2acba / 3

- [disable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-01453fc80596ccd7bbbc492bcb8f736a914dbb7cb04bd6f458c70924d8a2405c): complete subsection reference.

<a id="canonical-921250ac465e65c42c814a483ad4fcb71509a513543859b724465fe4e131162b"></a>

<a id="canonical-7e19880915026097b4b42b3b1ece1ed061ec620883a65a240a661c4136d8634c"></a>

## domains property — dynamic_proxy / ebf048f2acba / 4

Type: `["list", "string"]`. Computed.

List of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form Supported
Domains and search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

A list of Domains to be proxied. Wildcard hosts are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-9840b58ba1e5a9f257bf7b84015bac735f073e9ff657440c1fcf4f4a65c22012): complete subsection reference.

- [http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540): complete subsection reference.

- [https_proxy](data-sources--proxy--reference--group-002.md#canonical-2125bd393c2cc76b22804679d6c4fd09b3236aa532af4090fa9572561af9e068): complete subsection reference.

- [sni_proxy](data-sources--proxy--reference--group-003.md#canonical-e18c7ee38c8451f8941a75893ffc6cfb577567027ff281ee2641062380837f32): complete subsection reference.

<a id="canonical-d86bd60a9cef3db2fe8378131a715e7074013ccd844645e75d33ec595993e6e5"></a>

## Next pages — dynamic_proxy / ebf048f2acba / 5

- [dynamic_proxy.disable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-01453fc80596ccd7bbbc492bcb8f736a914dbb7cb04bd6f458c70924d8a2405c)
- [dynamic_proxy.enable_dns_masquerade](data-sources--proxy--reference--group-001.md#canonical-9840b58ba1e5a9f257bf7b84015bac735f073e9ff657440c1fcf4f4a65c22012)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-2125bd393c2cc76b22804679d6c4fd09b3236aa532af4090fa9572561af9e068)
- [dynamic_proxy.sni_proxy](data-sources--proxy--reference--group-003.md#canonical-e18c7ee38c8451f8941a75893ffc6cfb577567027ff281ee2641062380837f32)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-01453fc80596ccd7bbbc492bcb8f736a914dbb7cb04bd6f458c70924d8a2405c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3149d8fc557b38e2664a1c5421eae87ba6642d697dbee0cc7ccd986985a9fbba"></a>

## dynamic_proxy.disable_dns_masquerade — dynamic_proxy.disable_dns_masquerade / 12fb5efa417b / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- dynamic_proxy.disable_dns_masquerade

<a id="canonical-cc213a8c01504ab52839b0b5ea68503ecdde8ce5493a9e28fe79cc58ecfb3695"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable dns masquerade.

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

<a id="canonical-5325fa0fa9e373bb3caeba583ed06b74cc27f8740280197714dd3baa4ce42ac3"></a>

## Direct properties — dynamic_proxy.disable_dns_masquerade / 12fb5efa417b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cfe9bbd28ae03cc2e6a618a66d3be8411bfbe63f3455c9b55aadfa841e8b315f"></a>

## Next pages — dynamic_proxy.disable_dns_masquerade / 12fb5efa417b / 4

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-9840b58ba1e5a9f257bf7b84015bac735f073e9ff657440c1fcf4f4a65c22012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c82d0d23098cd7a0675a62a8ff012a49a5d13b9817a676bdca063d903b7acd56"></a>

## dynamic_proxy.enable_dns_masquerade — dynamic_proxy.enable_dns_masquerade / 955dad8bec76 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- dynamic_proxy.enable_dns_masquerade

<a id="canonical-a2b3c12a615021c5665c6905a039d478c51a12a28f4e2eb61743db117f1a70d8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable dns masquerade.

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

<a id="canonical-8f1ca45b8ba705c32becab327bc828881bcf3f5d8e40cc14edc771b73b6716cc"></a>

## Direct properties — dynamic_proxy.enable_dns_masquerade / 955dad8bec76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d58439a337c8fbc88f25ed38925a61817eb182eb8c2ef4cdb57e736cac4679bd"></a>

## Next pages — dynamic_proxy.enable_dns_masquerade / 955dad8bec76 / 4

- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31a35aea35fde1531c2fb9bb8aad390f93a3c8708cc81d8285bcde3b789c4eae"></a>

## dynamic_proxy.http_proxy — dynamic_proxy.http_proxy / 77dab21a17d1 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- dynamic_proxy.http_proxy

<a id="canonical-704fe152d4940ee3f666eeaa9bf2b16d9807ae15e0b60be728dcbcc8898a5bb1"></a>

Type: `"single"`. Computed.

Dynamic HTTP Proxy Type. Parameters for dynamic HTTP proxy.

Upstream description:

Parameters for dynamic HTTP proxy.

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

<a id="canonical-79d61c0a7f91139eef7700d9368fa37f6095b1b50ef7c57442a26851d4dcb63d"></a>

## Direct properties — dynamic_proxy.http_proxy / 77dab21a17d1 / 3

- [more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361): complete subsection reference.

<a id="canonical-7fbb6bd3a7c8c1a25e5af9b26bebe9eb96fbeff19a0a87f9712f38849dde9214"></a>

## Next pages — dynamic_proxy.http_proxy / 77dab21a17d1 / 4

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9e931bb396f225c9e6b36d2df6dd246a59c0c6c4820f1d705c0ae9a5f697eea"></a>

## dynamic_proxy.http_proxy.more_option — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540)
- dynamic_proxy.http_proxy.more_option

<a id="canonical-d1d672e0a511da5170d1d50a91751419738ebc9a1c6c6fd0c54e9c1ffa6871c5"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

<a id="canonical-e5afd8e0a257b5d133a965b4e6b9ab44a1aa9cdc69152acce549d780ff176175"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 3

- [buffer_policy](data-sources--proxy--reference--group-001.md#canonical-965c121e14fa275709805d9b0434299b2b28140b8478b3ee6ef3c6a3c2906950): complete subsection reference.

- [compression_params](data-sources--proxy--reference--group-001.md#canonical-68abb5f7bf9daa19999f5582db955869e957bce3ef06466334e9dc5b56ac097a): complete subsection reference.

<a id="canonical-4a7e96a266f7c3e576940053c78a1c3b54c1a5758cc5718c20c4ff74bf1fadbd"></a>

<a id="canonical-3da7a280b1c969497a49b9968a9c1723dbde9f4bdff05fc0923f50939a09c1fc"></a>

## custom_errors property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 4

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-aeefd9aa56263e81b2092d95f79dd38e523135d2c7a86a41d21bbdecffeee5e4"></a>

<a id="canonical-e9055bc572f34e427d486fba65a9408fbf2770bdf2e9ff2db3421e7a12472956"></a>

## disable_default_error_pages property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 5

Type: `"bool"`. Computed.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-e3ee6da16d394378884e0a37e6bd68b94c7cade6e64580ea06184f06804c39df): complete subsection reference.

- [enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-e0ff37c9edd368c4af8af060179ef18e08b757df5c98833a99108ddf95f6093b): complete subsection reference.

<a id="canonical-e56f716379069db659b901686890c0aea186e05a93b8232c8c2c7bbd211be0cb"></a>

<a id="canonical-8e08274bf9efa4ad9c666d811b8ae33a73d803fa00d190b976a997ae7c7d739a"></a>

## idle_timeout property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 6

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-3d39fb99af0c676b0e19d99b60d885a5f3d22e627d792ed31821ad38ccae2a57"></a>

<a id="canonical-85ca288054e6a0b263d4de1f1755b44422001ede5b6003d034520314acc88e33"></a>

## max_request_header_size property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 7

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-b1affec10a5a42c9ea183450f0ad59b60f39306e85826225e557192bb2b42ba9"></a>

<a id="canonical-879fbb1476bdcf9c14e2c34a6cdc86885823e939b9295f359bc8097760a777e7"></a>

## max_requests_per_connection property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 8

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-4a9867efacb51c00a18d24d1af513862c8ce9d28a477558bf65d2ea9b79c63a1): complete subsection reference.

- [request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-b3c4f122df2800d44df60133acc8c2b410b33d0918a680a5d913900f424c5b98): complete subsection reference.

<a id="canonical-845a8d7ce9d0f8b0ffe43d8aa33cc009aa8764524b46f0ddc7b5532fe4a51fbd"></a>

<a id="canonical-feea537d4572721609c4586f7ce4039cb8733d39e03ae32acfca7bd5fc52c581"></a>

## request_cookies_to_remove property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 9

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-439cd40cd580e0eaadd2a17ad2f020f2b49c233a12a3084075d81c5357c89d15): complete subsection reference.

<a id="canonical-6f1bd4a12305d06b7714256a1766a923a04abb7e4214d67f0596fddf7143c211"></a>

<a id="canonical-aedeee9dea74764e4bec7622a172d48755743cc522688a204f512c3d07da7766"></a>

## request_headers_to_remove property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 10

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-4166d597b91f76efd9158396ac4e58e5a80a58313b28da3daee2294cef26834c): complete subsection reference.

<a id="canonical-7f6c02099ae9ef57f211dcf0383b7cc38685dc5e608230c453aa1d831fd921ef"></a>

<a id="canonical-a94bc2cbfd4ca2ebd442b542cb005b8a46910220d714ac7ca894668b03a60233"></a>

## response_cookies_to_remove property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 11

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-63224269d026632b16ac0145cc6df6785df994b6115bf1aa612d63b8f01e4e1b): complete subsection reference.

<a id="canonical-7604440cb653892ecbebef9e68762a15ecf2a9e60534e7dbf8d3da08a9b922ca"></a>

<a id="canonical-5ccb4ca00d768c1796b104abf028bac04d3f7d7e19efe45df7eb61c34ca5a51f"></a>

## response_headers_to_remove property — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 12

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-13526245f69318c31809a2d8fbe878f3bdbb82a906be28fd5a50d5c316267c0c"></a>

## Next pages — dynamic_proxy.http_proxy.more_option / 9d106d4ee505 / 13

- [dynamic_proxy.http_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-001.md#canonical-965c121e14fa275709805d9b0434299b2b28140b8478b3ee6ef3c6a3c2906950)
- [dynamic_proxy.http_proxy.more_option.compression_params](data-sources--proxy--reference--group-001.md#canonical-68abb5f7bf9daa19999f5582db955869e957bce3ef06466334e9dc5b56ac097a)
- [dynamic_proxy.http_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-001.md#canonical-e3ee6da16d394378884e0a37e6bd68b94c7cade6e64580ea06184f06804c39df)
- [dynamic_proxy.http_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-e0ff37c9edd368c4af8af060179ef18e08b757df5c98833a99108ddf95f6093b)
- [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-4a9867efacb51c00a18d24d1af513862c8ce9d28a477558bf65d2ea9b79c63a1)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-b3c4f122df2800d44df60133acc8c2b410b33d0918a680a5d913900f424c5b98)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-439cd40cd580e0eaadd2a17ad2f020f2b49c233a12a3084075d81c5357c89d15)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-4166d597b91f76efd9158396ac4e58e5a80a58313b28da3daee2294cef26834c)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-63224269d026632b16ac0145cc6df6785df994b6115bf1aa612d63b8f01e4e1b)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-965c121e14fa275709805d9b0434299b2b28140b8478b3ee6ef3c6a3c2906950"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ceb9fafdfb1b6a606951ebee6b5a033630602ec85fb279a62a82b8f9b36209a"></a>

## dynamic_proxy.http_proxy.more_option.buffer_policy — dynamic_proxy.http_proxy.more_option.buffer_policy / 7600d5cb3682 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361)
- dynamic_proxy.http_proxy.more_option.buffer_policy

<a id="canonical-2aad559a3513fa6d7c58fdbb543a70ba6f19c969e2696cde2f5bca05aea43e0d"></a>

Type: `"single"`. Computed.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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

<a id="canonical-2adc04901051749444fd45123360b9a4fa2d719be231a01c457adc11743a1d7b"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.buffer_policy / 7600d5cb3682 / 3

<a id="canonical-b20cc96af27e3d63953255eb68a6e7386189194bd9b7a1037e14cf045c49ac28"></a>

<a id="canonical-d04c6bdd51cb44ee94e47ac73a74e11902eb31da6d2e6105c4172127e8cc445b"></a>

## disabled property — dynamic_proxy.http_proxy.more_option.buffer_policy / 7600d5cb3682 / 4

Type: `"bool"`. Computed.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-7a86f4ebe6e4b1c52120484ac4c3856a31f978363814a3677d754e225a709e43"></a>

<a id="canonical-386366e04553795892fa87ef0e7efcd9c00d3bfa6e746027e23a8fa80bf0f73b"></a>

## max_request_bytes property — dynamic_proxy.http_proxy.more_option.buffer_policy / 7600d5cb3682 / 5

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-285a53554fcf990ad43f9818880fbc673cd8ed6fd63efcdef64e943760226345"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.buffer_policy / 7600d5cb3682 / 6

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-68abb5f7bf9daa19999f5582db955869e957bce3ef06466334e9dc5b56ac097a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fdb41e1596ebbf76cb9e88e86248c9a1ff6d8ddc4fee31ca84bd0bba78ccef9"></a>

## dynamic_proxy.http_proxy.more_option.compression_params — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-9acd4bc4a9e8bb1d7a074a15317383a122f6fd5e80455148f36c321e8dfc97f3)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-d85e83016e8c10aca2da3353807b4605750c7c3f1fa42c3a96f38d39900bd237)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-e4b4bfc28d7f69cc2f254a0187b19feec74bc291300469ae112c76791bbe6540)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361)
- dynamic_proxy.http_proxy.more_option.compression_params

<a id="canonical-ddb5a698751ae82d07ae19cfde8d65c799486da8da82025245d665e1e9c9ffa3"></a>

Type: `"single"`. Computed.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/javascript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

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

<a id="canonical-f22060f9c2b4e3ea13ed1ef23d53526b1ffb9abe4135f9105c5d463ca4ddaed6"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 3

<a id="canonical-18abbee002b0c6fcd299d64f33a4da4d44db9a1266678d89caa9def75034926c"></a>

<a id="canonical-cf37b665d8fb83a0f1f072c7c8edff71ee77bf5a7e993dfd90eb901a206094db"></a>

## content_length property — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 4

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

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
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-6359224353187f5026696671a8c9cbcc5a623fbbc7d59890b09c5b3944c69ced"></a>

<a id="canonical-6c2c7da9908e2dc55f7395a4f4727651fbc5ed588d7d55c9d0de1b764f88c49f"></a>

## content_type property — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 5

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/javascript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/javascript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-72367c37475ec266fd3a2374812e6c00b492f55b361375283de84a0cda8924cf"></a>

<a id="canonical-345602b517b1c40cbab3c658fcc522855a63118c5b0c55f798845a9f30e04de6"></a>

## disable_on_etag_header property — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 6

Type: `"bool"`. Computed.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-9545b3e205fffa4a0cd404d5e3cef7d7bda3bb1259cf591361ad8f6d94086cea"></a>

<a id="canonical-a4aa218456133081fe492152ef3454fd306f71ab0b8e9b12ac49bfa355b4e789"></a>

## remove_accept_encoding_header property — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 7

Type: `"bool"`. Computed.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-2db820cb769c5820ca3bf5ccec4894570c7f78ac809523cbdf436868fb4319a6"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.compression_params / fc457aed46f1 / 8

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-12db5e9e606e82fcd00cdf126b77d2aacab89f13c6a45cf21be215e341e74361)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-e3ee6da16d394378884e0a37e6bd68b94c7cade6e64580ea06184f06804c39df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
