---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41dc2b54e8901f04a10fc31153996b11131f60c973f72a12240a636b25bf3717"></a>

## Property reference — Property reference / 61c3dc06c0b7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- Property reference

<a id="canonical-e36205f11499ba5980f2ea517fbc1c30b12e950e651e1591a14bc01aa427bc3c"></a>

## Direct properties — Property reference / 61c3dc06c0b7 / 3

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-03ed25dd1314b56d5707e18431baf13d1cc9e27ab94078070fef979459e99a09): complete subsection reference.

<a id="canonical-57a0c3f0bcc465eb0e8463e8c2424c88827d4500004942d2813e971b2ad9c636"></a>

<a id="canonical-b0e926a859fe8a32377a39bbb3941d824b183610022a0a0ce7662a32b6fe93de"></a>

## annotations property — Property reference / 61c3dc06c0b7 / 4

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

<a id="canonical-6eccb1d4e757af91e6d288e8b8485ce17b7768465068a9d97578bc31864a3f72"></a>

<a id="canonical-c22c0556e0bf56946a00e8f6861351be6f6c130c6f7f85ced5963ca6e3e655e4"></a>

## connection_timeout property — Property reference / 61c3dc06c0b7 / 5

Type: `"number"`. Optional, Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

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

<a id="canonical-4e78d82f6e055a9718d0dde8299e06d7dc4ea9d039814ed4da9b4bb77787d3a3"></a>

<a id="canonical-e7560944d8c69eede6fd2abbb166e1c905aeed47e7af6d240c5f627743f9ce2d"></a>

## description property — Property reference / 61c3dc06c0b7 / 6

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

<a id="canonical-db5ebf0039af9b0b92a032161fc5b3b75f5accd67d6549035cfd06a7b4f17a91"></a>

<a id="canonical-25ba22d38d7a90601a6d1efceff9a3d67b4120d7af491af7a9c7197f70582d99"></a>

## disable property — Property reference / 61c3dc06c0b7 / 7

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

- [do_not_advertise](resources--proxy--reference--group-001.md#canonical-70c3fa24b943e6d0d0abb21e802eebe41a34b9b523b18f87c83dc0c7fb69b6ee): complete subsection reference.

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc): complete subsection reference.

- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd): complete subsection reference.

<a id="canonical-c9d04c4b3380d5c5ee2289d7b7cc0bb339da89bf6ebe3de245de1fe5bdc02f24"></a>

<a id="canonical-c219235caff30098a210462b1aff1fbadae62e6ee9b0827a87c55f24f5a86d1f"></a>

## id property — Property reference / 61c3dc06c0b7 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6e5cfef879f17c2c3c80b72ef3ea5bf2a137c6d187353acfdcc19d7fb215d5d6"></a>

<a id="canonical-b5325ffe6837521bbffc90f4083166eec0f7cc7ea443cd44443b7000dff5e192"></a>

## labels property — Property reference / 61c3dc06c0b7 / 9

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

<a id="canonical-17569778d9de83f7289e02d766c3f6c09f4e04a39f4bfa3897aff7fe663a0814"></a>

<a id="canonical-7a1f258ebedbebe4fafd7c9ff2d54ccaf0355542b706ff9748774cc47a4f08f4"></a>

## name property — Property reference / 61c3dc06c0b7 / 10

Type: `"string"`. Required.

Name of the Proxy. Must be unique within the namespace.

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

<a id="canonical-054bdc8a0e9e0be6d2140674c5717ad7ce34140212a9e394d510920fdb59a564"></a>

<a id="canonical-b8bac9548052a74684904ed78e05e5a0164b94705b889446fae3610bc8a28980"></a>

## namespace property — Property reference / 61c3dc06c0b7 / 11

Type: `"string"`. Required.

Namespace where the Proxy is created.

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

- [no_forward_proxy_policy](resources--proxy--reference--group-004.md#canonical-ff779654e9d7a3b0554f6fef391a194052fe154808cfcc51b7979ca7298d8d2a): complete subsection reference.

- [no_interception](resources--proxy--reference--group-004.md#canonical-c1d4991763e98c65835bf4a35d131a34666e553de1f223df41092c700befd9fb): complete subsection reference.

- [site_local_inside_network](resources--proxy--reference--group-004.md#canonical-b0a1caa8391902d4f6610c4dbed1b5392a4fd7331ae15afa906a7878ea3a2f40): complete subsection reference.

- [site_local_network](resources--proxy--reference--group-004.md#canonical-73781e83eabf108576e0ba5c0ee0ead9ada60a44a59a56ea627eaf9a0e9570c7): complete subsection reference.

- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2): complete subsection reference.

- [timeouts](resources--proxy--reference--group-005.md#canonical-9751a1a4304e20a986e19abdac2b60f6e30919c2760a358e62b4e168f6bfa05e): complete subsection reference.

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d): complete subsection reference.

<a id="canonical-6292e3fee59e91fd6d7d72792fc3ad89ab3a430d0f86c9b23007e17b72c7c41c"></a>

## All schema paths — Property reference / 61c3dc06c0b7 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-2995d3d9e218487e574fc805123357e795771f05b1e864beb609e91688bf22d8) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-8b8edb217be897359a303cba561e15dd57074092b1cf2a44e8687fbb5dbcaa96) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](resources--proxy--reference--group-001.md#canonical-1aec0ac7b71c5e6c17adb25fba2f140c7fa14b6d80a3b6f5f5feb87a15dcd1eb) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](resources--proxy--reference--group-001.md#canonical-af91e144764b495da43df8082e0a660bd1d4179943a03b021087066302b78fbb) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](resources--proxy--reference--group-001.md#canonical-636f4fefc14c2ffbf6c32be80f8bb69746fb9fad6d2106dccc0293875b3f94b7) |
| `annotations` | [annotations](resources--proxy--reference--group-001.md#canonical-57a0c3f0bcc465eb0e8463e8c2424c88827d4500004942d2813e971b2ad9c636) |
| `connection_timeout` | [connection_timeout](resources--proxy--reference--group-001.md#canonical-6eccb1d4e757af91e6d288e8b8485ce17b7768465068a9d97578bc31864a3f72) |
| `description` | [description](resources--proxy--reference--group-001.md#canonical-4e78d82f6e055a9718d0dde8299e06d7dc4ea9d039814ed4da9b4bb77787d3a3) |
| `disable` | [disable](resources--proxy--reference--group-001.md#canonical-db5ebf0039af9b0b92a032161fc5b3b75f5accd67d6549035cfd06a7b4f17a91) |
| `do_not_advertise` | [do_not_advertise](resources--proxy--reference--group-001.md#canonical-367f44df89655aec2249204d65893391d96f76e60d5f18121b475ccaffd9b744) |
| `dynamic_proxy` | [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-3a1c759354daefc0018b08ac5c61909d8f79022e190b394d77c9912d58f9205e) |
| `dynamic_proxy.disable_dns_masquerade` | [dynamic_proxy.disable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-7ee761d9e86ff17fe98aad73968ba4f636f170f4d9733dc35c330a77013ce29f) |
| `dynamic_proxy.domains` | [dynamic_proxy.domains](resources--proxy--reference--group-001.md#canonical-bc04064377a81302eab5bf8340748f8b7c23850a2e0ccb62527ffdc07392c044) |
| `dynamic_proxy.enable_dns_masquerade` | [dynamic_proxy.enable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-3841990f265b71c08884e49958b237d67360e474abefd9ab959bac573846bf32) |
| `dynamic_proxy.http_proxy` | [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-0ce9745496e96287decb192d617341f375a79913024edc9e018815ee82701bea) |
| `dynamic_proxy.http_proxy.more_option` | [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-86f6c0fa544bab23ecc8c61c30ad0df27e1e8e58a32068578a6554d272303c0c) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy` | [dynamic_proxy.http_proxy.more_option.buffer_policy](resources--proxy--reference--group-001.md#canonical-90d28e56f19585d78a0f31fea579fbc95d05e8aa4399c90d8712018cc3bbfca3) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.http_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-001.md#canonical-7697b8f0b3469dcfd35d23aab043076491869e276e3a800257ba088ede9354e5) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-001.md#canonical-a975dfdcaade03373735b38bc64925e4c284bebdb8947f2e31519517673100ef) |
| `dynamic_proxy.http_proxy.more_option.compression_params` | [dynamic_proxy.http_proxy.more_option.compression_params](resources--proxy--reference--group-002.md#canonical-213e7b68c13f0b23f51adbe941d36c0be190aced7d6965ceb4fe9211b7e2d29b) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_length` | [dynamic_proxy.http_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-002.md#canonical-1ade19f0532e84732c2f7033ae7186b0acc8c69e642b75fb0591c53bed0e398a) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_type` | [dynamic_proxy.http_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-002.md#canonical-d5a48af64dcc846fb7bf88f6f5cd0f43d002614351c39836b6ac02d326b8d2ae) |
| `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-002.md#canonical-ca3bc4ffead0383e7f7326a9fccc72d366b9160a6e526037870d8bd3b49f5c0e) |
| `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-002.md#canonical-70da857101199fd3ff3a66c26937c982eed85bcfbbaae4aee99c5fcbb3b5d301) |
| `dynamic_proxy.http_proxy.more_option.custom_errors` | [dynamic_proxy.http_proxy.more_option.custom_errors](resources--proxy--reference--group-001.md#canonical-033cb53ee0b885ddddfbf577ff345c839203772b1cf073f71f358e3ff8bf4095) |
| `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.http_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-001.md#canonical-31d5780c2e7893c90548be06c3537aa41c2efae8d9c8fc69707b084431178062) |
| `dynamic_proxy.http_proxy.more_option.disable_path_normalize` | [dynamic_proxy.http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-514171c0d87873a53b8b1347eb29a4e06ca64ce549e757979e79b160690e016c) |
| `dynamic_proxy.http_proxy.more_option.enable_path_normalize` | [dynamic_proxy.http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-22767d0b9a6cbfe9bd81c83311122aae19c737be37f93e7594f0bd5a16611f91) |
| `dynamic_proxy.http_proxy.more_option.idle_timeout` | [dynamic_proxy.http_proxy.more_option.idle_timeout](resources--proxy--reference--group-001.md#canonical-a4361a8f620ce684c91f51d713342885a45501b96ccc5c742ccc78dc1fd5fce6) |
| `dynamic_proxy.http_proxy.more_option.max_request_header_size` | [dynamic_proxy.http_proxy.more_option.max_request_header_size](resources--proxy--reference--group-001.md#canonical-42af490ac9a317852867e19cf4f44d1769a0a864374ddab23ba6d0c9aeb45359) |
| `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.http_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-001.md#canonical-30a186f4ea74ef8d4bfa12369133a98140761ba104e29769da647387829fa7d2) |
| `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-6d74ee486240c9b92f6a4044fc64d8cf4754ad39715c4328c92cff3947178213) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-8aab627787a43e6a87aa344b972343838c51e5d1fbf2c144192c3f11e94ecf85) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-24a3f5cc95b0244665752f1f7315b6367d64fc899068e9baf29ddad7b9c439c4) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-b18ef25c6a0cd60cd8fbb47f82056a85f4c15e54566b2ce3fd233a64f1dc8f55) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-01c74856ea4d34518868cb26f7635104fd9057051ec8d73296c570fcf31692d6) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-eb50d4cc3a2d7527e65dda6369bedd22a369db8c7bd5150d157d1f227860dfaa) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-744de0473da1b0382ec08b39e73894c635bcfef56caece6ac673fc2e35bd9b7c) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-d526c7751195e4364d3602722f6c1704e6019684e5dad79faabd0d800875ad45) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-5eebcd0c7e213846afc98db0a339c0529f1b623e03d1c86a2ce767a9814475e4) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-a9c20b2353b22ba658439b51d362dd35c2eb5626ec402fcb037675585ddd6e4e) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-67d20dbcf357e78728dab51dcb495e19e69a2dcffbf11658291b73babf2df95e) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-bc7923339a2797bf0374ec93ad3fb4a589c47c0d9f7d8abc5cc3edca3ff84f99) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-ee5350ec06390e91a58e2e5c95f00bb6ab63b5e63dc5a258969da32e3747c0d8) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-001.md#canonical-eede7123dcd0bbbb5af8cffc0aaa3efb8bf112eaebf2f4a825c4bf3c690a7abd) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-a27640bfabd62f0fc60548a6b293d63a34cad50d2aed5976695dd017416318d6) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-002.md#canonical-a41f98db6bc763e27ace2163d1333c0a9eb8a622fda3fdf7351c325d1e2dc9bc) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-002.md#canonical-33eb189dfaf28c867c6c0ac5b7b240ec5acca8a5ea103655b339e3bfa239ae88) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1297978fcaccffe3e8b3e46c7acfd978245b9897225ccb84c832cff8cb6b974a) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-5dcacd39c1117c5334e9b2761e2e884d03c9c6c1367a47f09aebfc02ecc119d9) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-f7f42434e3aaf79acae9472d6cabfac44dbfefd111a6501e83a1ad14bd692e8e) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-6905e4d9c6c7d1f471567398ff8e62cadc489a2c9c511fc703bfa1a9cb4a9414) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-4cf63e417dbe54475131b2766eeadd7eceec8192d4c68054746233ee29f8ce94) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-eeb8058ae3004af59224908e0331d5a56dad4fae4f2d73377b907f71db37540d) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-48ea85867d4f64879632794f1917c6c829ed1991326146f546c29d3d6a09c7f4) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-e329d30a4bb4af93995301173833541b2c53d5f76fd6c28f88349069f5f5c4f8) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-002.md#canonical-fab2f2e257d2add31da82e8ced5aac3ed54d5d2f923716c3255873a661afb05f) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-001.md#canonical-3781d60b585cae05cfd10c5671102dcb88c9068472debdaec693cd0d9aa029d0) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-c9d5fef1f1dccaac93bbdb2c88b9b873a1b2e445aefe859f91e09853e1b59567) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-002.md#canonical-90dda50f44b36dc30653cc414194ee2bc14ebb3350d1550b9d4ec9d64be7fceb) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-002.md#canonical-0267e9fdb908a2902164573b845e0bede8daaac7d964f4bddf631e8421c0e008) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-002.md#canonical-1e9ea30cac410586e3152382aeb0992181d9d796b1c64dadcb305ba92c40249d) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-002.md#canonical-457cc8e3666c259ea9a59caa709695b60ec62e1028c3420d71d0438cdc4b4d75) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-002.md#canonical-2a2b7025d1d09bd75f702aca6c7d0cfbc88b2b9479b5bb457a909a9b15a3f235) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-002.md#canonical-a05ff8b0e1868e53f63af923df183709231c3e80bd08c8769c92483e90fe80d2) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-002.md#canonical-48e69eb70eef49eb1202a97f1942267ac56a25b75884becc1095114cd74b63c6) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-002.md#canonical-b9b59ed04e5568f76bbacef32b31893c84310a7e561bf2a34e14d0957412a088) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-002.md#canonical-2a777b9a5099769dc6d77e949022bfbc8dd4dede555c148a02c016c79324b982) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-002.md#canonical-2563e709019f874ff61be5a968891f6a71f2b344946bae25df8f7abe1caa787f) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-002.md#canonical-b6e513a5a9f1d0269bd4b962bbf54295423c3ec1d874c688a499e468fcdd84c4) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-002.md#canonical-0be8ce327b355ef34b455ac46b56bd37458d0185c77768e1a06ddff65da825d1) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-002.md#canonical-d37a3ef59c111fd8f308278216a034ff918491e43cb5a28ec818a939a04c8529) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-002.md#canonical-97f8050df33607417bce799b772f2d60568ff9c0db4967b6dc98a29fe2f9dadd) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-002.md#canonical-7a8535c7e42feec9e64840995d770a65e4a14eeff4171de4c346ad5b123ced92) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-002.md#canonical-e8f2aa39c4bfc7b46ddac685037c0456c23e3d1e91bebfd95410d4907a8ace4f) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-5b09c170e04f634fcb5d14ecacef5da066bfd5021daf59be4b6565c9dba4259e) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-03e23172864cc0716eef0bdff6a6f2f77b574b457add46302baedfe843580704) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-002.md#canonical-95cd35b2bab01038c4409635e3ec86135073c1a06ec4597bf78c82a86a8f271d) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-002.md#canonical-326b0ab4b5938c858fe532efe200d1ba6e8f76cd6c0fcbdd5db63ca53f557198) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-002.md#canonical-3342ac460f6f6e0d66003eea0e7963c91aa135d773ec597459a1d7fb158dcb1c) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-b01e1826f76ab13443fc6182520de2732384094c575f43d8dc886d68d3a7c19f) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-8e84acb9e8061441d943276ee9cda520ea4f22d58f29532b7fc839c1829d2917) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-e0f29b3c330c7151e1bd878015843304400cc8b106449e390ea394d8acd8f2a3) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-93a958d9c9accc210d3b69f11c4dd254f2d8a231f5757f1ee7e19fb218b5020f) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-606f3509da6702954ca899cd8af1d29cd1ad7a160a3d55ea06740ec5b6fda24e) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-fc6f42dff80d3194a3bfdb42ccfe102ee4a2e11e5d2fa2fcbe1acb6cebb1580a) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-1b0fb27b140b5ed39a278b56b48ba74f050c41c32e007035a3b94fed7598bb07) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-216fa4893e4d460c33762adac444a33ecae71be89da4186a954614713c3e35ea) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-cff5654e0cd1d2d9bfbb5b692003b3e1b45d9d21381e1cbe915c6aa2774ed06d) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-001.md#canonical-0584a3744e02ad7b62ea0524b6ef902383a7ffca562ac80cea995d09c4b2a32b) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-22419f03471bc9f58e816611e4c854b7c4bdce0654048c00748cf95a03a536e8) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-002.md#canonical-375c5fa5654c8ef753e008a0b587ae5d476900dd1d75ca5fe8b72358022045d9) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-002.md#canonical-41fa3575e1f303eaa8f291aec3788ed92bc829811a3efbaed805ebd09fda7eac) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-7b67e3f04332d65dba95f0c4b3e18b94134d208f1949d7cc2c1179301cb21550) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-934a99e5a31560aeead0ee0eef71e88b3f89a15c73aec62be511860bea3c1477) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-15d8b4062131f5709de4c99484a4126b544d94f116e7fa1e29242b15aa7d6678) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-c72084d45953e553ee02a5e12a1a718fb0cf64c5e32a38aee41a91c939f2062c) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-54b460c3455474051a6454365edb304fab676af063bc2e435c4fb0a708b92216) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-3eddc912eefb51342772ab0a8c39c42d5c22a809c1f7d9c3477f626412c8d037) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-002.md#canonical-01eef4c09d4a692c095d380476bc4a5a0c622ea37457254ba8d6298ea1bab298) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-002.md#canonical-7347bb8fb5c7af03d5f33031d5f82d446c61a947a3e1e5822e0b1187146afe44) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-002.md#canonical-60098772833f15c0b17fdd4163d4b21fbf4a0d1e1b5e159a8aa1eb6fc08f7bc2) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-001.md#canonical-38bc27b46867398c9b3310eb150b7e37f468b62a6d5630df14bd3474e1c87912) |
| `dynamic_proxy.https_proxy` | [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-901bf8f71f3a9b0d1a7c0d6eec9a181156d1b4f066eb5f637974f2549b792a1d) |
| `dynamic_proxy.https_proxy.more_option` | [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-af3207efedb07304c822887e9dd4d8f5dc42ea5596ef958de8dcc8a0a2cb14fd) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy` | [dynamic_proxy.https_proxy.more_option.buffer_policy](resources--proxy--reference--group-002.md#canonical-a164a7cc116b986334ada46dea5e5274d820168e10ea0e96dbaf5a9737e7ad33) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.https_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-002.md#canonical-8225328ae2f824fe4c47841f962edce062c37b9f1c841de8d2e9678491cb7a5f) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-002.md#canonical-d43c9fdb03a8752b15de1b69abd385791c26206c741cc08dddf640ac9797803f) |
| `dynamic_proxy.https_proxy.more_option.compression_params` | [dynamic_proxy.https_proxy.more_option.compression_params](resources--proxy--reference--group-002.md#canonical-e8a3d9b2dac40a251c92b50b8f9c49fce6db87e565f5f2ab7a7757dd36fdfe87) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_length` | [dynamic_proxy.https_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-002.md#canonical-965f2186d92305ebb91e60c7c0eae6bb9429600e74fddb7902c62560bd4bdfc2) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_type` | [dynamic_proxy.https_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-002.md#canonical-263b3c19085c17b141339029b5b606c9353d732fcaf84828d86d52f1bef080b6) |
| `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-002.md#canonical-1296370495ecf496080d4e9162d54b13d9b280aab882b45df968e983425161c9) |
| `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-002.md#canonical-1d99ec991b74b3c75b2c0520f3c4297dea7376dde0e465b373b89a8f7abdd97c) |
| `dynamic_proxy.https_proxy.more_option.custom_errors` | [dynamic_proxy.https_proxy.more_option.custom_errors](resources--proxy--reference--group-002.md#canonical-18c587539eab37a4ac13157411c6eb696cad636d7862b24656a65237238d9635) |
| `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.https_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-002.md#canonical-7731b930a79e803c4efb076ed102764252b164e3523858cd00fa429764bbcc80) |
| `dynamic_proxy.https_proxy.more_option.disable_path_normalize` | [dynamic_proxy.https_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-bfabaf4feafcaedba2eaf63374e4db9d8015f535e4072f5c8f6310d87c0655f4) |
| `dynamic_proxy.https_proxy.more_option.enable_path_normalize` | [dynamic_proxy.https_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-ffcae2af9f51ab79a55332320f6476afff8f0363ba6c04709f451e9b4cc93421) |
| `dynamic_proxy.https_proxy.more_option.idle_timeout` | [dynamic_proxy.https_proxy.more_option.idle_timeout](resources--proxy--reference--group-002.md#canonical-880f7da5f45a05a1a84608ab2d5e5151c58a5b496bf5141c34c01f7635b37301) |
| `dynamic_proxy.https_proxy.more_option.max_request_header_size` | [dynamic_proxy.https_proxy.more_option.max_request_header_size](resources--proxy--reference--group-002.md#canonical-5f34c5c1e56a9d1ad080df6d2630a7968855db34146d8d40b40b960b3972e0e2) |
| `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.https_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-002.md#canonical-95de99c88625409f62b647d01ff7139a8bebdcde264567f13afdb464e2be990c) |
| `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-a2a1ec7e2f95add1eff045afa9d96d463fea1089f70ad48a8444cccc47dcc284) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-8d92f6b4b3d166d1198d8ecdc7cecdd88442c9b083731fb40160260fe624d6e8) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-002.md#canonical-2f6759ecb869e8a8f4542ca70ce1ad7678556500a35b7c1c5142fe8d259b593c) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-002.md#canonical-6285ca628622fa6dace80fe6adab9f202d5945149089d4dc1aeb8db5999e4b59) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-863914f07b423f2a2059cac8875a0fb1ed79f6c1df2d14cf247dc21f7dcd09e3) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0e10d68e221b7d4ed0883ea99aa77c7fdffa7a8cbb98defaaf93445a83e6e450) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-002.md#canonical-da49d0e219b193b891550c18c98ba6827f955dba06e7160a6f6781542dc85384) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-002.md#canonical-6ed108ade63a87d8c558128a8ac9e02f7c47b4680f1033191b0e2b53cfeb28a5) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-002.md#canonical-c94398ae3df12b047785ac0715753c563e8ed136b3425e9ee4cfb36b2533cd8c) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-e8731b958be537a080842874982cc214d1cbbd3431a5d80ae6d2f910974a3dfd) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-44a87d54c7da78a9044f8a908905ff696656930cf454f4bbdda5ae22189d3345) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-adda5788761c2314be2e9676df92eead10fb051d0fb5fe8923eec45fc001296d) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-002.md#canonical-f1d4ea809e0dcd922aa81ca08b87ea6323017104a4673e2b9ba5b79fd1d5ca00) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-002.md#canonical-affa7ab2b4c1a5bf9a0d60ccb68c04db8faf0630dc8de94af550fb40b8c03b59) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-ce072dfd8f22b07a7a0584e657cdafd9d28cdf2b78f8adaf9e342ff8577dc012) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-003.md#canonical-efc7994d878284b47658b493c51441d94b1dcfb5abd77f335affcee835c0d6d2) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-003.md#canonical-d8e65962a54579ec7a124c9aff5c62508eab50bf2f428e04b662121947bb3e49) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-73126fe1365aa11cd858b53715c3f3b37353bf5f968daebbba22a5c1b42d88bb) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-def5db9c9167aec09fd16d0d41983afcf91e6dd8f1f6ce63eb33b87bec0d9079) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-dde2e159539d61784b189d2652da5bf878df7f26b09e8c47796c63d0d7f2fcfe) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-fc5157defd103dde13ce1960fb9e5925d4d9a23e249ea385f5a9775920a326f6) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-b83ce02da65e07f432eaeb8236c58d6b88d1231b72acf58357f7526f70e5c206) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-345ad5df88709767ba92e1c629db27f3591c45ea8c2bd33b5df15be8832f72ca) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-b3236a0e8ce5390de9ed502a55c72adad8f914e67dd633f0af33bd567ac32907) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-9307256dd812fbe704bb57f17b49adad4e7bd86e7eb78776d17f1ed67c1d5c69) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-003.md#canonical-0e268ab06b6c284734e468477205ff53d3851ec725f68001a9f87b4f63aebcc9) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-002.md#canonical-3faa52c82256c455077bb47d3f09382fd7f505423c449a9075b4b1384c7ca17f) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-1d84e689c48818ddff97d653b62c8d868120ffc466292bf07ba246fb0a1f68b5) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-003.md#canonical-f6789ce8c613669bd9b2829337acd0ee081f63d6da84ddbde1e63c3ff227bdd3) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-003.md#canonical-7ea6ab77a932b9e7c75088ee104f1a817160ba9ca3cd6516ab2e689160f5dcb2) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-003.md#canonical-fde5e5547386cc321f7b371871c740feccda33bd6cee8ba5e1857da2f20c5033) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-003.md#canonical-4ac2d6afb03fb4a4e490ee8f13039ac803f882418fcbea2d6fcffc23fc240bcd) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-003.md#canonical-3964b73c11fd6b8bb2c62edcbf9086b344577d305fb15f170f8e3c812097f48d) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-003.md#canonical-3846119cf8019aba0dddf71d12de8326a9e63cf1a97b4be71894e93e3448ad46) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-003.md#canonical-5249f2a8b1f717979a6de522112ce1389114f45c785016a98ffb48d09d58d69d) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-003.md#canonical-75b92ea341db100d88df4ea1c2a78ac81c39194084885cf4d97aa05898b36b29) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-003.md#canonical-3deb605afe342ccd460999e7327c89956952a1ca82389a45a52dd434694a5ee0) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-003.md#canonical-093c9895221585b2389d2577c1cdf18dbd97181335eb988c58dacdbaf98b614c) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-003.md#canonical-bcacd9b3a43d9572bc8e074f870497c5a57bd1dc759f926033f9331f686791f9) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-003.md#canonical-a4aa1aa8fa6eb10ab8a5a7c40aa7e2e469b0ad25f7c72dee45afd7679831e689) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-003.md#canonical-5e4f1ffd90ea6a6cf38fce7fee2c48736cc2ceaf80c0a16b727c86e4e1cb08b8) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-003.md#canonical-caf4a43ab89fe3f5a53a582b611f2240a01a11b9e893fff116d3a45fd9865170) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-003.md#canonical-216babfb31b6427031a2abbef435c7a7d363a01a89ef46d233ff79a7332b7e56) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-003.md#canonical-3c748cbd0e66691206e35dcf6438cea3c91b03e2b5276c850451f24428707a77) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-003.md#canonical-5fc386f84a7d58de9d577685472fa5d2da541182f6e4f1383b16175745d9fec8) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-003.md#canonical-b4e7c27c08e39711eb0b98108dc38bcda293b93ad7ada7990eed58bb5b848c36) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-003.md#canonical-211515a0979f56519f75d8887041cc9ef369958ab5946fb10ec69556b13df2cc) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-003.md#canonical-ea78b164544e10f92ae1b476ba5dbcb22a36a61523c650e4103d387692686c14) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-003.md#canonical-1e7514915131ca4206d128173c150ae5b89b901cc40bc47e234bd9256361b348) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-163822ebf6b7f36adff9bcdb4c85c3a5672defc29b695a5ac48e19f581cab491) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0b3db43cd56782b693caa028ec6c89c89ed3a4015756a6d3b1e8d3051e8e7196) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-50b8af6b02e45ca0e6c179249b0f43eb5fcdf7c60e4100bccefb65d7bd238306) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-ffbe7957824c67e59682f43d924d75cd9a2b4e2431f81d0bff88346fba928202) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-2a1c6cc731fbb6b48decb91280e32d175b4d9141474c645a582e3d633477cfaf) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-3ce960d462604f16a0261ff1f4bdead7aba37d038345e12f9b67ff902442d090) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-a5394f870e19db11a4e3bb91ec029cb220cffc583edd34510e7a84061630dd9f) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-128014b0884c6723fc6964605015f62dc8538327748803cb4229f4a4aa6f7acd) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-003.md#canonical-7b5390589429f465d26dbc3fd56dcdfd8b23d91e2c1248fba4e9daa74a5bc784) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-002.md#canonical-7f668ddd4211b48f0b6f60bdb8d8ca54724ecfde3737441672e6526352972732) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-5f937c3dc2ef3e629d6d71571408731890587b7cd7df22333e695119148f0ef9) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-003.md#canonical-d560f4c8bf2ca77e1578d7e7b36ec679de1b57bbf5f1120a86cbcbbc10c728f4) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-003.md#canonical-c1b4013dfd19df143e5f8368bcaaacc10a10b06712db8fab3b76bc63a445d713) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-90a95e297fd9b7efb3c44a456b243b49fc24e2490459353b37b646b50b16fd1f) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-83bf96cdc6c668f5cff0f6fd40335bc037ddff0caaeb98d2dc6c8785a9202cd7) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-c75d8edfae2841d22066d20ac322fe75661fae1a8516f023f084a0501f0db3a2) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-6c14227b53bd25801135b2cbe2579f2de0e7245afffd5463238cd037d2925408) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-e15beb5ddb5f3560f7bd625fd7e86119c3dfd1a66aa7407e394506f22759d4ee) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-3eaf5caff2dc056b35ad80e1a43faf6aeb735f46df523b67fe50d2789e556cf6) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-bf153412e0676c4ab9701fcdc32ce0768434827a00f158f8e47250726f3be66e) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-f9b627d34bbfc4df0ae4650650627e0ebfe00883456d3b917f5af417db36f82b) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-003.md#canonical-9c46de8e39baff6aa8225697014bb20898c00b57e8205b657f7cc9c5433dcd18) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-002.md#canonical-e0d3993566ef7b4c0ff6cb5d56c1f19da2874a23b6e233b12ec846cc43bd335c) |
| `dynamic_proxy.https_proxy.tls_params` | [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-210a2b6bbc6ab27cc4e75722a9f5117d0d44ffe2b7420767729d96a1e30418c7) |
| `dynamic_proxy.https_proxy.tls_params.no_mtls` | [dynamic_proxy.https_proxy.tls_params.no_mtls](resources--proxy--reference--group-003.md#canonical-d4ee163eb7179b8e39359ce162b32cb47034c9229ce64d7c7f3d3b30a6e128d4) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates` | [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-8d617c6d8034af814ac995ac0c857b181cd0248ca71e1a0a3359df78bbe950f1) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url](resources--proxy--reference--group-003.md#canonical-bbf89f9ace7b6a1a09c43d8e18ec2619d1fe15691781061ded5a47e7dbed29ed) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-df13d7bbd0b2046d83be6150841a67ef7e5d540b828fa9148fe095342293717d) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--proxy--reference--group-003.md#canonical-635a2c821ff8710379e89da9331a1d6088449c1d352d52d2cdbab199ff31b72d) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec](resources--proxy--reference--group-003.md#canonical-6f3a3eec996d9c52528dfae64263200020be520e7a067cccf9746caa005caa24) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-b22f6656a5c5e8c9ddd5ba7633a1e2de2a114f29adcca63d3b3ee7c9f1b059be) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-dcccae23044f069608dde72b6783fd21b0c5fdf8eacf8ed76b858c16bd63befb) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-4a9a58c5d0df485cd935099ee0f31a8756a0e62ae38a021f9ed06a94592136fd) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-003.md#canonical-87389e7338f5dd090043421e82afec5d12636b39b824a7a9605ec159559bd9f4) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location](resources--proxy--reference--group-003.md#canonical-a6e7938d35de63e50d4a3d0538b87748d4ab30d08111ca6bb74e886fda5dbbb6) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--proxy--reference--group-003.md#canonical-86a33e660b6c4ab4f396b4fe8b5ca6b0157b7a291ad9ffda643cfa9fae1649d6) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](resources--proxy--reference--group-003.md#canonical-dee9c7546407893d8989b54fdeaf964c977e5332ca3349b138922bfb6f5f71b5) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--proxy--reference--group-003.md#canonical-6ceb5d6ec3da0d7c825792b12fd823cf9b17458060e4158c67351c99ee44972c) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url](resources--proxy--reference--group-003.md#canonical-432fc51992236b894af410f0a7e1ca49132ad220b6e6c13af9da14fcdeffd5e4) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](resources--proxy--reference--group-003.md#canonical-3e692bf0fd084f1a44772e38b6c08e52b7118770201cc5b01fb3401a3eab0e7a) |
| `dynamic_proxy.https_proxy.tls_params.tls_config` | [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-6d4bebf34276d1a16a9dd763a73dfdfe62fccfa4b12925f3188e875e187b6bcd) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](resources--proxy--reference--group-003.md#canonical-1a8e2799d27d0192cffcc11864eae6f38f7ccab2a9bef1d40c32700fe965f176) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites](resources--proxy--reference--group-003.md#canonical-e252a7f86d6532538972065467acf7513f93010316b503da7533d63208654385) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version](resources--proxy--reference--group-003.md#canonical-81cea1944cd8217f7a9aa079212b5cd0489982fe601f72321cf9f853e2ee72db) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version](resources--proxy--reference--group-003.md#canonical-1d5c797d9cf6f9cdee9f139aef7d9ecd6d73476b0a20e38d083b04ae92e1cbf4) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](resources--proxy--reference--group-003.md#canonical-ca51b8f33ca6b5e920a8fdd3f8b6b34e37d173536e26cfea158515a58ca605f3) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](resources--proxy--reference--group-003.md#canonical-c4264a864d4098eb915947defbf9c623bdfa2c498e6a644ebf9828bc69287251) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](resources--proxy--reference--group-003.md#canonical-ac0239c0589bb69addaa83679afde58acf6d7392534a3c505f396da836a9e925) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls` | [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-67b03affb463ffe6ea8407ba256cd6efb0d61b17b078ed1e3b4dad05e551528c) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` | [dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional](resources--proxy--reference--group-003.md#canonical-a7f6f0ee5b05ace772633f74904970559795c2c4d8556a20cf930a1a10192a4b) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](resources--proxy--reference--group-003.md#canonical-c0becc407463f445eab96c70fc13c3101ab758193b6784874355dec25983921c) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name](resources--proxy--reference--group-003.md#canonical-901a7d42d657db806eb7d4837ed04ce7775ae5b492e9191963532e265bdaa396) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace](resources--proxy--reference--group-003.md#canonical-851cffba861caf64ae73be0f0cdb491dc97c344e031eda4c1a1134cb1cf03567) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant](resources--proxy--reference--group-003.md#canonical-b1829e42a784cc052c5a33e2b4e0e29924fb8d062faceac27d7fddb73cf9af0a) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](resources--proxy--reference--group-003.md#canonical-06f895f01179ae846f56369defe4522aa89d195f8a1fbc5c2340cb06c88afe07) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](resources--proxy--reference--group-003.md#canonical-917fc74064c3486398f3fe2ae8df819e0f44d7b8c8cb4be23f5e1d65c5d01e23) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name](resources--proxy--reference--group-003.md#canonical-57ac69d876804f8f73ca823373e0b8bf7f79cd7f34e5ca07e0cda906d2f5c2e6) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace](resources--proxy--reference--group-003.md#canonical-88e02fbbd954b148188b9aa8ca725c1df5c887a6dad4916af9d6c99bca8c3850) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant](resources--proxy--reference--group-003.md#canonical-9d8de6f07dd5749f0ce6a398049b23f142b0aa8feca65f7ec7add4794ce40566) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url](resources--proxy--reference--group-003.md#canonical-a5a568571cb10bf1be576b577c6bfc835a5418b265546c2f339cb8a97f03c3e1) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](resources--proxy--reference--group-004.md#canonical-d3b70e43bee2f79116ce851707a8bdef4a1332d0aad2d11e5765e6657fd55b06) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](resources--proxy--reference--group-004.md#canonical-e50879c809f6638e947e2d9adcf192224823a7ff4a700ba131c5d83ca30bdc7d) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements](resources--proxy--reference--group-004.md#canonical-337b0ccefb6a87c397016ccf9b8a8df42d7f6ff31bac9f1cb1d24b0f43f731d0) |
| `dynamic_proxy.sni_proxy` | [dynamic_proxy.sni_proxy](resources--proxy--reference--group-004.md#canonical-60310432f74632f6f49cd5f4c5ae0a2e93d92b3f77a6213e158264d2162e76cf) |
| `dynamic_proxy.sni_proxy.idle_timeout` | [dynamic_proxy.sni_proxy.idle_timeout](resources--proxy--reference--group-004.md#canonical-d29cb9b33cdc3b7ff857f92c53c87f9356ff478d992ed75c2ebab679f1486b92) |
| `http_proxy` | [http_proxy](resources--proxy--reference--group-004.md#canonical-0b32a135f0eac087814b5eeb5a2baa7a8860790f0ef07edeea58f4e6eb6bfb9e) |
| `http_proxy.enable_http` | [http_proxy.enable_http](resources--proxy--reference--group-004.md#canonical-a0c518db07a4bc8a6e1d8649c58607ad6c1eb53adb2847382ab8c4d4a9882b2a) |
| `http_proxy.more_option` | [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-a906b91650ee34c26970211cfc973a24f9ce7defdf483e764392b6e79f5daddd) |
| `http_proxy.more_option.buffer_policy` | [http_proxy.more_option.buffer_policy](resources--proxy--reference--group-004.md#canonical-8cf6b5e271a99c0e7f2e0cacab9346cc939a5348c869dd44e61585532a70b8e8) |
| `http_proxy.more_option.buffer_policy.disabled` | [http_proxy.more_option.buffer_policy.disabled](resources--proxy--reference--group-004.md#canonical-cc9e4c73ff321a392c44588b8329db1491c83cd4c1613a86fff4cfe7e0fe3581) |
| `http_proxy.more_option.buffer_policy.max_request_bytes` | [http_proxy.more_option.buffer_policy.max_request_bytes](resources--proxy--reference--group-004.md#canonical-cae7da4dd1f6cff109156d7129d872c79e37fc47bdd30ea80d9e8c905fd28284) |
| `http_proxy.more_option.compression_params` | [http_proxy.more_option.compression_params](resources--proxy--reference--group-004.md#canonical-383e31266524d5de80038a80d4c0a2af4836046f08e784b34e2e0d796d92f0d1) |
| `http_proxy.more_option.compression_params.content_length` | [http_proxy.more_option.compression_params.content_length](resources--proxy--reference--group-004.md#canonical-adf50e60cd45e3caf5ff2d843263680e33c1f77e62b7a2e5383e4b166c454833) |
| `http_proxy.more_option.compression_params.content_type` | [http_proxy.more_option.compression_params.content_type](resources--proxy--reference--group-004.md#canonical-1e338a7438e4d257263a6704514c050c016c9a79c3fbf3d3c1fc01215e5e8d3a) |
| `http_proxy.more_option.compression_params.disable_on_etag_header` | [http_proxy.more_option.compression_params.disable_on_etag_header](resources--proxy--reference--group-004.md#canonical-4595fe55730030125bab9d3260a5d9f4d82cd6e018808036cdba6635e6aa734e) |
| `http_proxy.more_option.compression_params.remove_accept_encoding_header` | [http_proxy.more_option.compression_params.remove_accept_encoding_header](resources--proxy--reference--group-004.md#canonical-9e2c93b3da8c0d5b37d695de58e6471660458df76eeea02a3936912dacfcab5d) |
| `http_proxy.more_option.custom_errors` | [http_proxy.more_option.custom_errors](resources--proxy--reference--group-004.md#canonical-3728d6163c87af12a231ff6334c66ac90542e13399074237d086616f0335c037) |
| `http_proxy.more_option.disable_default_error_pages` | [http_proxy.more_option.disable_default_error_pages](resources--proxy--reference--group-004.md#canonical-ede00cd446bc94f2926bd2f32d27b560d0500fb9ebe92759d423cccce4a0b27c) |
| `http_proxy.more_option.disable_path_normalize` | [http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-004.md#canonical-16a6f43d1719a4a5c5833cc54d3a8cdd50ee26e5af32b8677d6af23280ffa70f) |
| `http_proxy.more_option.enable_path_normalize` | [http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-004.md#canonical-23463db2b6894d74a8c81043382e9770c15ba12f66815b70ea8ed86bb1ca7a15) |
| `http_proxy.more_option.idle_timeout` | [http_proxy.more_option.idle_timeout](resources--proxy--reference--group-004.md#canonical-165c2731825ddf391ef2c730dfd54987e049fba6d39ed8e099734fd1fe7c55c6) |
| `http_proxy.more_option.max_request_header_size` | [http_proxy.more_option.max_request_header_size](resources--proxy--reference--group-004.md#canonical-746169c1a1f2cae9081b4f96861ec304f68eb411cc68cefcdcfd40c979ea1fe1) |
| `http_proxy.more_option.max_requests_per_connection` | [http_proxy.more_option.max_requests_per_connection](resources--proxy--reference--group-004.md#canonical-42725721913a1a8958e594f270a1311cd06a7f059f56064a59d3d55aa2535032) |
| `http_proxy.more_option.no_request_limit_per_connection` | [http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-004.md#canonical-870c0cd6a16e0ac2df872480cf12d204243ea9fb3121cb8811ebd78c717abf03) |
| `http_proxy.more_option.request_cookies_to_add` | [http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-004.md#canonical-6403d5afd60fc9062c8773971be2b52db10f6ab3ce7f450ab29bffedcf0d1653) |
| `http_proxy.more_option.request_cookies_to_add.name` | [http_proxy.more_option.request_cookies_to_add.name](resources--proxy--reference--group-004.md#canonical-d289451206d36ab6ff33d535f276c6b225d8cc21f7b7447e37b479116546f6f4) |
| `http_proxy.more_option.request_cookies_to_add.overwrite` | [http_proxy.more_option.request_cookies_to_add.overwrite](resources--proxy--reference--group-004.md#canonical-23e8e497e2526f48a1c78cd7e3e10fd19975b7be2988151c2a315835d7280474) |
| `http_proxy.more_option.request_cookies_to_add.secret_value` | [http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0ecb558ae890d89890aa26e9d4fdefd65434cbebef436b8d0e7084704008d7c8) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-4412b42b2341068db086c5263e12538a9505eec39a6977245624b9047d824154) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-fdf6faad09357baa950a2006960baef9b8033b58bb5b678fd6921fa241176c93) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-443bac7ee5f9f1c073fa62c6f8f5806f4323204f696018134019f1e61614c982) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-f4804d0d3c5a25c2e33640e82f66380b4779eedd3b06df4460f846dc44024c05) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-c9751314a80a655a95a209f0614bc6cc169cb290ee8bf88de5aabd1a874325d7) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-e85708957d0043bcce26015b337d44a5606ed9c5061620937fc92e3377e43787) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-a5be61a8bf79f9f239a20feaf2c04e209f4f07936a2cefc53cd5c08dfa2a20a2) |
| `http_proxy.more_option.request_cookies_to_add.value` | [http_proxy.more_option.request_cookies_to_add.value](resources--proxy--reference--group-004.md#canonical-958fb0cf0cdf75baa1783e15e13f9a7ccb4547fe42c3874524c601f1825ee886) |
| `http_proxy.more_option.request_cookies_to_remove` | [http_proxy.more_option.request_cookies_to_remove](resources--proxy--reference--group-004.md#canonical-18bac30cb907dff889c8d5e6def87368feec0df166b34f5bc7e39d15f601627f) |
| `http_proxy.more_option.request_headers_to_add` | [http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-004.md#canonical-6c196bb5ba382e4fa32433b6fe72c471d45388eb4b2632512f12726a724ec014) |
| `http_proxy.more_option.request_headers_to_add.append` | [http_proxy.more_option.request_headers_to_add.append](resources--proxy--reference--group-004.md#canonical-c7a39f624ceb748ea3c07b8116996f9a2b442ca2a9dc3d8c72ae136b26ec3c6d) |
| `http_proxy.more_option.request_headers_to_add.name` | [http_proxy.more_option.request_headers_to_add.name](resources--proxy--reference--group-004.md#canonical-631940ef57f7b03fdcb7f34a12e9a97088f458343fc46cdd98ffcdccd2e41b55) |
| `http_proxy.more_option.request_headers_to_add.secret_value` | [http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-b0b5802aa48c22bac8eb685faeac1038628495e389e245ce9c02ffe1eec374f7) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-4de9cf8db78a01aef63995366708a9be56f800f74d410afe21aa9543b8beeb3f) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-d71f1817893768c60aea0c9008d95ac24d60d8256e232c5986cbe355a8323957) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-ed547e4a40cc20e190d97588259f508cfcadcbf2ed5aef67eaf6207098d5a538) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-87cfb1899da2b9bbd8a79577b722a55c290f88654e01a8ac6cd8ad24ca02fe08) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-b5d408b3cbe5f94c165c4d6c7ed9fb2d9c562580db3605cada8378c481020c3a) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-065d036ffe024b38b38b931c6b87a1657c64cd6e6a4fd70eb794796488247451) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-348534e651d7c123fef76b4607e86a2527ba62b13a8274f303f18d9b44788c94) |
| `http_proxy.more_option.request_headers_to_add.value` | [http_proxy.more_option.request_headers_to_add.value](resources--proxy--reference--group-004.md#canonical-83b85d9cf7952cfb34ef6d3b2c59f4278fec4b489d7997cb3d12e07fd5421089) |
| `http_proxy.more_option.request_headers_to_remove` | [http_proxy.more_option.request_headers_to_remove](resources--proxy--reference--group-004.md#canonical-1b4517345ca4208fa81826bceb895c183ffe03fbc339aa122946fcf724fcb1a7) |
| `http_proxy.more_option.response_cookies_to_add` | [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-545ef34733b875f9551ae14b2d9baab2b50a6796d237dffd48333a42cb4e9e09) |
| `http_proxy.more_option.response_cookies_to_add.add_domain` | [http_proxy.more_option.response_cookies_to_add.add_domain](resources--proxy--reference--group-004.md#canonical-c7076a9671d4e51dd81f4fcc993dbe058b861a27f8879a4a5768878c8a0cf517) |
| `http_proxy.more_option.response_cookies_to_add.add_expiry` | [http_proxy.more_option.response_cookies_to_add.add_expiry](resources--proxy--reference--group-004.md#canonical-ee46a5c02108c1869d9762111b8777db1cd77392ac05c468dad1d18fc7900fee) |
| `http_proxy.more_option.response_cookies_to_add.add_httponly` | [http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-004.md#canonical-b0a5ea0e66cc529b1215fcefa6a4a6ea0304a00f807109194a1c4b5ca7a69c7c) |
| `http_proxy.more_option.response_cookies_to_add.add_partitioned` | [http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-004.md#canonical-00a00748160b4aac76f84f81e8b0ad10be7a37de46f37d3135e8bd5c33c7e769) |
| `http_proxy.more_option.response_cookies_to_add.add_path` | [http_proxy.more_option.response_cookies_to_add.add_path](resources--proxy--reference--group-004.md#canonical-0bd5977ae42cd2014585edd03699e1049b769681d8eeb70f7ed30eca80bb1b42) |
| `http_proxy.more_option.response_cookies_to_add.add_secure` | [http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-004.md#canonical-7f7ce3ae882c837c8200be2c7302c9be18c64833590395e0752e5a126d776b2a) |
| `http_proxy.more_option.response_cookies_to_add.ignore_domain` | [http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-004.md#canonical-2986c2f629ee517bfe63ebd4be28a6cd1a25ef5e9524877152b07f903531329c) |
| `http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-004.md#canonical-f1a5a9ed7d3dd7237b088543cdb298ed354b760a8198f7c5806452965f275eec) |
| `http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-004.md#canonical-c98625cec1cde80b40a34e910d2db392edbd9ba2bf2cd60bd89e3b88b37a2fbe) |
| `http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-004.md#canonical-6844a4b0437cade383668be426bd198e6943d9c15d12c971cf791093caef32b1) |
| `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-004.md#canonical-9046fece5c9e5c4c01c186be7433ca4aa81a4d5f7c91f29c2d321dc3c0947cf3) |
| `http_proxy.more_option.response_cookies_to_add.ignore_path` | [http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-004.md#canonical-dcbae0a311485a0c2b8f64799facb1b761f380afeda5f02eadb2f24028a9e837) |
| `http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-004.md#canonical-df6c6072c3b545418789d6dc942451d41d6f02d65db1baa387d86d0aeb5433f4) |
| `http_proxy.more_option.response_cookies_to_add.ignore_secure` | [http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-004.md#canonical-6fff4dc454de26eeca67acf979f13907a0aeed98404df50adb79c05de433e6be) |
| `http_proxy.more_option.response_cookies_to_add.ignore_value` | [http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-004.md#canonical-8014117b70769794f36939ff45149cc9869abba103ba5bfd5e7c7eb408b20a19) |
| `http_proxy.more_option.response_cookies_to_add.max_age_value` | [http_proxy.more_option.response_cookies_to_add.max_age_value](resources--proxy--reference--group-004.md#canonical-b43bee81e88230fe4a110156e8b7ae5d667201aa40719ca3fedd06f842566d3b) |
| `http_proxy.more_option.response_cookies_to_add.name` | [http_proxy.more_option.response_cookies_to_add.name](resources--proxy--reference--group-004.md#canonical-8e6b681d95f879a349f4dc37b7d654fef2b3ab9aca1ac0993df170266f538586) |
| `http_proxy.more_option.response_cookies_to_add.overwrite` | [http_proxy.more_option.response_cookies_to_add.overwrite](resources--proxy--reference--group-004.md#canonical-dbcbed562d99be4ef0311c142458f5395cbffbc7e4b3bc5df8df0dd886f3c90c) |
| `http_proxy.more_option.response_cookies_to_add.samesite_lax` | [http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-004.md#canonical-3b02cdad8394d48d53050bc062d51fe6c8bb3cd583b353c2865d54219afa2925) |
| `http_proxy.more_option.response_cookies_to_add.samesite_none` | [http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-004.md#canonical-3f3285a4892d5b931a61e9ef8f89f44854da746ffd93f795df9e55511cfa5153) |
| `http_proxy.more_option.response_cookies_to_add.samesite_strict` | [http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-004.md#canonical-55a9c9a31daec7b5c4d6ebe8d5b44bbaf1a4461ce50c676b33a306f6324a406a) |
| `http_proxy.more_option.response_cookies_to_add.secret_value` | [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-70978771232ddc64734761c156cc39a218f50403811d217b907d5be3d1172728) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-ff6dc40090338ddf4be9cc24946acb4b4353475cad299ef22189e400371deace) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-7a467bfb090ea249374ef8a3367f5e171df5c48a194d569201bc9f03cc3fd545) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-f45a00ef3cd26d8656bd2278d5a86d4731122e304b276a350e3968ae36a6e6d1) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-19360d722940b59c5ca068921a526103fa7689af6482775757326067b34301a5) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-904e054a6ed6ecddf28200b90c3776b51c50868debd6786e53a83289d6ddfaf3) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-3e609e559185c932697f9d7758d9514932b449cd2556274ea4d1fd767afc7c42) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-6ef38a2804404b4943c54ff293a25b10a32f222908e0e537b15ece3554f550cb) |
| `http_proxy.more_option.response_cookies_to_add.value` | [http_proxy.more_option.response_cookies_to_add.value](resources--proxy--reference--group-004.md#canonical-b7a32287ce5ecf96d5293f3450c02fb1523e8565486ade4a103ab9ed1f2f05f8) |
| `http_proxy.more_option.response_cookies_to_remove` | [http_proxy.more_option.response_cookies_to_remove](resources--proxy--reference--group-004.md#canonical-9070dc7d65283745bd6925624f39c68f02e94ae86937a7e7359d4288cce94be5) |
| `http_proxy.more_option.response_headers_to_add` | [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-004.md#canonical-7ff679e8770a6c67354618df1271d147d18c7c7349c4f600e7907b9369eab9bd) |
| `http_proxy.more_option.response_headers_to_add.append` | [http_proxy.more_option.response_headers_to_add.append](resources--proxy--reference--group-004.md#canonical-f3e18e8971ea7dc96f9b6d473be8ee1766ec117c3224c928a17859a676355333) |
| `http_proxy.more_option.response_headers_to_add.name` | [http_proxy.more_option.response_headers_to_add.name](resources--proxy--reference--group-004.md#canonical-f7a4fe2bf2a5f52eca30d61cf8551572dafb6ee45a9b312118650c5f913d6f79) |
| `http_proxy.more_option.response_headers_to_add.secret_value` | [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-d16dd2fe6edda61da9387cff1a885013e71a7e5578f538b762793190950bab9a) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-004.md#canonical-197c254a4202938d805c9298d2c8aa9a58bbbccb7fdb18736c3596a264e16b77) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-004.md#canonical-9499b04e51c9f08592afab1e3b17b580f61869cdac9ec543e954bfb9347aeba4) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--proxy--reference--group-004.md#canonical-ed693583f7c8f43d7942802f47a000a37d66c00e14b741f502963120f4d835ea) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--proxy--reference--group-004.md#canonical-c74c03120b6c938b79d3f2deda3af57da22ce1dee6bc0db7ac0a0a1b2b190075) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-004.md#canonical-a09d6bb0f7f4ef2962e010690849f16d7077f28128846beb7e7a91a3c3c7423c) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--proxy--reference--group-004.md#canonical-48b2c057aa8af7cdeb4d8f98c2eb90c892d26b12e738f3f8eec0c8801a1d6c08) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](resources--proxy--reference--group-004.md#canonical-3c1363464d9600ce5745864013e910035bb508737c6a6e4287aff270f5dfc59a) |
| `http_proxy.more_option.response_headers_to_add.value` | [http_proxy.more_option.response_headers_to_add.value](resources--proxy--reference--group-004.md#canonical-acaa10f37ce096d6085ab6742d515ddeca971deadd3f1dc80d6712f787e463ca) |
| `http_proxy.more_option.response_headers_to_remove` | [http_proxy.more_option.response_headers_to_remove](resources--proxy--reference--group-004.md#canonical-8259051d13c027d004f95b106fc6e679ae4e5b2fc2dd31ad457742da3b35833b) |
| `id` | [id](resources--proxy--reference--group-001.md#canonical-c9d04c4b3380d5c5ee2289d7b7cc0bb339da89bf6ebe3de245de1fe5bdc02f24) |
| `labels` | [labels](resources--proxy--reference--group-001.md#canonical-6e5cfef879f17c2c3c80b72ef3ea5bf2a137c6d187353acfdcc19d7fb215d5d6) |
| `name` | [name](resources--proxy--reference--group-001.md#canonical-17569778d9de83f7289e02d766c3f6c09f4e04a39f4bfa3897aff7fe663a0814) |
| `namespace` | [namespace](resources--proxy--reference--group-001.md#canonical-054bdc8a0e9e0be6d2140674c5717ad7ce34140212a9e394d510920fdb59a564) |
| `no_forward_proxy_policy` | [no_forward_proxy_policy](resources--proxy--reference--group-004.md#canonical-72d597129d8e04f873ef78a4172e75b919979521b55b6cd8614ee84456067dde) |
| `no_interception` | [no_interception](resources--proxy--reference--group-004.md#canonical-fe2f7d86c6d293b9d1825c60aa5a53f34fd4db2312fcfee1d1e64b58eff6901c) |
| `site_local_inside_network` | [site_local_inside_network](resources--proxy--reference--group-004.md#canonical-d96f7e5101cc20f99673e7fa523e66f77389734b46fd6469a694eb553bc6fe30) |
| `site_local_network` | [site_local_network](resources--proxy--reference--group-004.md#canonical-e8ebfe418e106ba12def83255b7186488931618fa70331fc01027565e775fa8b) |
| `site_virtual_sites` | [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-3f1690011579c5952d206b3af1687ef51b17cc9a189ce572f1a79976d6cc93d4) |
| `site_virtual_sites.advertise_where` | [site_virtual_sites.advertise_where](resources--proxy--reference--group-004.md#canonical-90521e29fc9d52b0d5804df53d52724d9bbca2cbd91a51baea9dde19237f4bc8) |
| `site_virtual_sites.advertise_where.port` | [site_virtual_sites.advertise_where.port](resources--proxy--reference--group-004.md#canonical-b7f570ad2bdce578996dc5d596091f48c794c7143894120436ccc71bd0f42ad7) |
| `site_virtual_sites.advertise_where.site` | [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-004.md#canonical-fae667cfd5b27beb8ab1677ee24b938d147ff538f23a4dfe33f6680ef4114ba8) |
| `site_virtual_sites.advertise_where.site.ip` | [site_virtual_sites.advertise_where.site.ip](resources--proxy--reference--group-004.md#canonical-c86ee1fbe48534a1931be608c257b28e80136b75ee4d0bdb582a35eb146a97f5) |
| `site_virtual_sites.advertise_where.site.network` | [site_virtual_sites.advertise_where.site.network](resources--proxy--reference--group-004.md#canonical-fec5e059bc8261415fd9ed2c9b217b77e44463577ae9c3f2e4a8b76f4b635510) |
| `site_virtual_sites.advertise_where.site.site` | [site_virtual_sites.advertise_where.site.site](resources--proxy--reference--group-004.md#canonical-93e6f0d8a18186a4ddd871408da2e1942152ee62df065c8cdece33315b7191e7) |
| `site_virtual_sites.advertise_where.site.site.name` | [site_virtual_sites.advertise_where.site.site.name](resources--proxy--reference--group-004.md#canonical-d9d38728bdc52588685fa3990dcbe15a083d9fc126c94853ad41b57f2efca797) |
| `site_virtual_sites.advertise_where.site.site.namespace` | [site_virtual_sites.advertise_where.site.site.namespace](resources--proxy--reference--group-004.md#canonical-44a50821a7077ed29926153de2f10702f52e99cbd50ce40dbc31fd345572136f) |
| `site_virtual_sites.advertise_where.site.site.tenant` | [site_virtual_sites.advertise_where.site.site.tenant](resources--proxy--reference--group-004.md#canonical-df75a0c6723909da34cba46c6a9e8f3854b5bef6ac96d18d18de3f50960a4d0e) |
| `site_virtual_sites.advertise_where.use_default_port` | [site_virtual_sites.advertise_where.use_default_port](resources--proxy--reference--group-004.md#canonical-1a6325d2b67e0303382e9e7565890c30eed1e770c388a69406d994f15396c174) |
| `site_virtual_sites.advertise_where.virtual_site` | [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-004.md#canonical-2667ec23946d0986525e114ebffae9ceff274b3b064d723b7a26d63485ecbf4a) |
| `site_virtual_sites.advertise_where.virtual_site.network` | [site_virtual_sites.advertise_where.virtual_site.network](resources--proxy--reference--group-004.md#canonical-4bd8157050ec7704dec50e2e9127025580d4cbf3dab2919ccb0db878bcdf9104) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site` | [site_virtual_sites.advertise_where.virtual_site.virtual_site](resources--proxy--reference--group-005.md#canonical-29a3dfa68cd44bd3d5ef3f0f73688758c68d3fb24942c36ab338dc780692e6c9) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.name](resources--proxy--reference--group-005.md#canonical-3f144b13249edbbbeed35551078f149badf9a482576df616f0b14d24f52d2eb9) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace](resources--proxy--reference--group-005.md#canonical-c4d0dee162da1a1ed2159a1a40d3d09295a7b669e0daab7fec78357a6752e5f3) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant](resources--proxy--reference--group-005.md#canonical-6ea200dc7d3c48d66f310a5a3169581f14c2d9c26320b0c5235686ddeb86fb2c) |
| `timeouts` | [timeouts](resources--proxy--reference--group-005.md#canonical-d0a5e5eb29379d81d6c14549fa8c88f4c612ebb92a69c245ec1b81dd332eb21a) |
| `timeouts.create` | [timeouts.create](resources--proxy--reference--group-005.md#canonical-53629e55c23c4bc2d492abc92dd464d25b3eb05bc151619089e1416130f394b0) |
| `timeouts.delete` | [timeouts.delete](resources--proxy--reference--group-005.md#canonical-69cfb48854127933bda9572f87a34292bba31a0d0a42bb27e52910ed8e3e8c70) |
| `timeouts.read` | [timeouts.read](resources--proxy--reference--group-005.md#canonical-b82a9b406d2613e14b267fb16cd95c013b3f621b4b8413fd5f2452708e50f994) |
| `timeouts.update` | [timeouts.update](resources--proxy--reference--group-005.md#canonical-aedb71e93535ead60c15582b53b2dc5de0d700df665b9537d9e8c123df063f2b) |
| `tls_intercept` | [tls_intercept](resources--proxy--reference--group-005.md#canonical-b09947c5b3a5d27f69f31c0a2da06055276dbe6a003232d163a89b3739b4716c) |
| `tls_intercept.custom_certificate` | [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-1c296d6302f43d9e4f247cef86a712431b0a20f2f8f1d9a2f4e67f5e3ff3088c) |
| `tls_intercept.custom_certificate.certificate_url` | [tls_intercept.custom_certificate.certificate_url](resources--proxy--reference--group-005.md#canonical-a1287c752743836842397822744a2087bca40702fb550b6a3198fbb36a3bafe2) |
| `tls_intercept.custom_certificate.custom_hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-04a425e3786ff2559d5f4c9dab6750912533738667badccc3e1c5fa186a0ddc1) |
| `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](resources--proxy--reference--group-005.md#canonical-5b920e6113704575c5887f60bc917bc5aea10ef0342d75a599188fa2adb6d8d6) |
| `tls_intercept.custom_certificate.description_spec` | [tls_intercept.custom_certificate.description_spec](resources--proxy--reference--group-005.md#canonical-df705fa38c0927ebb693f1fcc097bb6b82558a29516e4096210740136e5f9cd0) |
| `tls_intercept.custom_certificate.disable_ocsp_stapling` | [tls_intercept.custom_certificate.disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-259f6c97fa7cb199208edcdffb53d5fdf28cf2ab47fd6e11806cf0635c80655a) |
| `tls_intercept.custom_certificate.private_key` | [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-a0838f597f776706faefb2de23b7b3b9d5fbf0766c4b13bee2277ea03c898a0c) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-7302469c0f348762abb8a573112f4e741385284ba0449950f3c8ef76b097ce13) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](resources--proxy--reference--group-005.md#canonical-b5cadd47772a6ac5f8c0481a93bbc0342502ed3e28bafb7294d27075348d2507) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](resources--proxy--reference--group-005.md#canonical-b268037a6cb3778a520ea69182e1edb6824c6f1af40b65694dbd22d84867b78a) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](resources--proxy--reference--group-005.md#canonical-f4fd8e4a497a694ebd3047940c5dabced2d9f019559d2c1806cfc4a3bff57461) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info` | [tls_intercept.custom_certificate.private_key.clear_secret_info](resources--proxy--reference--group-005.md#canonical-e581771cbb99e76684adc0a764970febb0a06454b786613284c02e3fe91243a4) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](resources--proxy--reference--group-005.md#canonical-e8bc692703edf76a89430b9de1ef750d3c9f6f7517981bf6ef9e8392c4a8bda3) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [tls_intercept.custom_certificate.private_key.clear_secret_info.url](resources--proxy--reference--group-005.md#canonical-a72935985b94ea510f9a705849f912e48d0fb34468d13e013f02440a2d600140) |
| `tls_intercept.custom_certificate.use_system_defaults` | [tls_intercept.custom_certificate.use_system_defaults](resources--proxy--reference--group-005.md#canonical-548fe943c83be548c70aaa3b4e8599f80a1532f7eda17cdaa52d91b58054e641) |
| `tls_intercept.enable_for_all_domains` | [tls_intercept.enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-f33fd7f2323687d0d1972b7f1f614e7c96f6cbae82067d599aff54c1cda7de22) |
| `tls_intercept.policy` | [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-c540da6563564ab3e14444da469dcd0d5d4a31e560a8f7efdb21002ec1410e30) |
| `tls_intercept.policy.interception_rules` | [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-f660b514493ec2f7c63b152d36fa33187223cb088752e03de2dc5ee62214ecbb) |
| `tls_intercept.policy.interception_rules.disable_interception` | [tls_intercept.policy.interception_rules.disable_interception](resources--proxy--reference--group-005.md#canonical-1da008e286394eb45cd978c1e280dac995581cf8cd9d94e6a6e1f8664a2253c5) |
| `tls_intercept.policy.interception_rules.domain_match` | [tls_intercept.policy.interception_rules.domain_match](resources--proxy--reference--group-005.md#canonical-2ab861b946294eca9c22691deab90ed1143fefd16777a82dbdfd77955902cd82) |
| `tls_intercept.policy.interception_rules.domain_match.exact_value` | [tls_intercept.policy.interception_rules.domain_match.exact_value](resources--proxy--reference--group-005.md#canonical-e1034094d564137d68a4d6717729b30f0802bc08cb3c29efccc56bdd1b3ebaa3) |
| `tls_intercept.policy.interception_rules.domain_match.regex_value` | [tls_intercept.policy.interception_rules.domain_match.regex_value](resources--proxy--reference--group-005.md#canonical-6eeb230b9e9e875c0ecf37c1b510c8a071b25d7baa11e6534bb49ec78df9d34c) |
| `tls_intercept.policy.interception_rules.domain_match.suffix_value` | [tls_intercept.policy.interception_rules.domain_match.suffix_value](resources--proxy--reference--group-005.md#canonical-2cf09a16582aa892704993fa2d0509cc55e19429d8a6fb2405499e69c26733af) |
| `tls_intercept.policy.interception_rules.enable_interception` | [tls_intercept.policy.interception_rules.enable_interception](resources--proxy--reference--group-005.md#canonical-c7023f5cc3272a4896ffb84105d795d302a4130ab69ec9c43684f8b9edcb1602) |
| `tls_intercept.trusted_ca_url` | [tls_intercept.trusted_ca_url](resources--proxy--reference--group-005.md#canonical-b2d18133147a3e800126d410bc2fd31d0b35cc2ee02f8b94571ef7fb74d38d62) |
| `tls_intercept.volterra_certificate` | [tls_intercept.volterra_certificate](resources--proxy--reference--group-005.md#canonical-99682d5181e7da867e7689103b8933908533ea81d6f5f2ea72604302354cc287) |
| `tls_intercept.volterra_trusted_ca` | [tls_intercept.volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-4762ccab0f0ca1c0ab12ae73896fdfda5af2967b72695a47f24a318a0308ce97) |

<a id="canonical-42e5ec0aac71982b1f346d0b24f465f42c789680e3ddf1c368af0c8cf3de23ef"></a>

## Next pages — Property reference / 61c3dc06c0b7 / 13

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-03ed25dd1314b56d5707e18431baf13d1cc9e27ab94078070fef979459e99a09)
- [do_not_advertise](resources--proxy--reference--group-001.md#canonical-70c3fa24b943e6d0d0abb21e802eebe41a34b9b523b18f87c83dc0c7fb69b6ee)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-f27c6a2642f7423263c759642467f19d941ec3cd45f1d6c978b8f6dc4e6633bd)
- [no_forward_proxy_policy](resources--proxy--reference--group-004.md#canonical-ff779654e9d7a3b0554f6fef391a194052fe154808cfcc51b7979ca7298d8d2a)
- [no_interception](resources--proxy--reference--group-004.md#canonical-c1d4991763e98c65835bf4a35d131a34666e553de1f223df41092c700befd9fb)
- [site_local_inside_network](resources--proxy--reference--group-004.md#canonical-b0a1caa8391902d4f6610c4dbed1b5392a4fd7331ae15afa906a7878ea3a2f40)
- [site_local_network](resources--proxy--reference--group-004.md#canonical-73781e83eabf108576e0ba5c0ee0ead9ada60a44a59a56ea627eaf9a0e9570c7)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-ad367c86467f35f5c26ee33c952dcf2fcaf2a7515f5835cae209a925754863b2)
- [timeouts](resources--proxy--reference--group-005.md#canonical-9751a1a4304e20a986e19abdac2b60f6e30919c2760a358e62b4e168f6bfa05e)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-4639ee2404690fbd00e5ce07dfcd45d8207d50ed983459435b1b3df804692b4d)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-03ed25dd1314b56d5707e18431baf13d1cc9e27ab94078070fef979459e99a09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff771e937910a1689052a25c90481f87eeb29aa1cbea0d6e9bca6c64c1ee9b06"></a>

## active_forward_proxy_policies — active_forward_proxy_policies / 00733b6d4b6b / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- active_forward_proxy_policies

<a id="canonical-2995d3d9e218487e574fc805123357e795771f05b1e864beb609e91688bf22d8"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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

OneOf alternatives in this subsection:

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-2995d3d9e218487e574fc805123357e795771f05b1e864beb609e91688bf22d8)
- [no_forward_proxy_policy](resources--proxy--reference--group-004.md#canonical-72d597129d8e04f873ef78a4172e75b919979521b55b6cd8614ee84456067dde)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c868eabafca4065dd1c33ae5656844b44f0c3398c2d9852f4b82f68dce4ed24"></a>

## Direct properties — active_forward_proxy_policies / 00733b6d4b6b / 3

- [forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-6c5eaefc28d1283e7eb0909893d775253005a83a6bf2800ff741c7c37dcc315a): complete subsection reference.

<a id="canonical-f6d36ae4e9bc4511925eb2483d5372ee1823baa7e57b2f0be439af250b22f33a"></a>

## Next pages — active_forward_proxy_policies / 00733b6d4b6b / 4

- [active_forward_proxy_policies.forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-6c5eaefc28d1283e7eb0909893d775253005a83a6bf2800ff741c7c37dcc315a)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-6c5eaefc28d1283e7eb0909893d775253005a83a6bf2800ff741c7c37dcc315a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7793daf8deb61a95a7a16d43aa81799dc1d5594db2532251970f66236cfcdf0"></a>

## active_forward_proxy_policies.forward_proxy_policies — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-03ed25dd1314b56d5707e18431baf13d1cc9e27ab94078070fef979459e99a09)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-8b8edb217be897359a303cba561e15dd57074092b1cf2a44e8687fbb5dbcaa96"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd01ec3864d27b37976b68da8b0a3ba5d61cc05dbe1ab66bb4badcfa12d05cfb"></a>

## Direct properties — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 3

<a id="canonical-1aec0ac7b71c5e6c17adb25fba2f140c7fa14b6d80a3b6f5f5feb87a15dcd1eb"></a>

<a id="canonical-cd778dbbb4b4b11f53d25fb5ec41007db5677b593c4c358591fecad64a26c334"></a>

## name property — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 4

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

<a id="canonical-af91e144764b495da43df8082e0a660bd1d4179943a03b021087066302b78fbb"></a>

<a id="canonical-cd13739929ef15712dce96e993d537daa309015ab7e6c728f5b76f82eb5b1d69"></a>

## namespace property — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 5

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

<a id="canonical-636f4fefc14c2ffbf6c32be80f8bb69746fb9fad6d2106dccc0293875b3f94b7"></a>

<a id="canonical-9645b2e68ddc75b959e6534801d4ab0fcfa6f450abe1f83e0db94183938c372b"></a>

## tenant property — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 6

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

<a id="canonical-6ef13128b37dab771be96e67eece2012793867e9731bf47024be2e365e038828"></a>

## Next pages — active_forward_proxy_policies.forward_proxy_policies / 1020465cc666 / 7

- [active_forward_proxy_policies](resources--proxy--reference--group-001.md#canonical-03ed25dd1314b56d5707e18431baf13d1cc9e27ab94078070fef979459e99a09)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-70c3fa24b943e6d0d0abb21e802eebe41a34b9b523b18f87c83dc0c7fb69b6ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6813cd4cb2faac91b0a5f9875f950e5722cb7794cac7884caacf97731a0544a"></a>

## do_not_advertise — do_not_advertise / aa00e122944f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- do_not_advertise

<a id="canonical-367f44df89655aec2249204d65893391d96f76e60d5f18121b475ccaffd9b744"></a>

Type: `["object", {}]`. Optional.

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

- [do_not_advertise](resources--proxy--reference--group-001.md#canonical-367f44df89655aec2249204d65893391d96f76e60d5f18121b475ccaffd9b744)
- [site_virtual_sites](resources--proxy--reference--group-004.md#canonical-3f1690011579c5952d206b3af1687ef51b17cc9a189ce572f1a79976d6cc93d4)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_advertise = {}
```

<a id="canonical-ca52f6a9c7b6b4a3b2052bfcdf0c8895d2432176f2323eee2ce52e1a5c5a3e68"></a>

## Direct properties — do_not_advertise / aa00e122944f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2436473f7499528645630f89eff4f7e2cb59379e3fc7417d3999b887346bbbc"></a>

## Next pages — do_not_advertise / aa00e122944f / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92678117aa2cd0b8b39effe089176829cd6b6f834a54d29edbc397ed98b36ed1"></a>

## dynamic_proxy — dynamic_proxy / e5457e1eed7c / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- dynamic_proxy

<a id="canonical-3a1c759354daefc0018b08ac5c61909d8f79022e190b394d77c9912d58f9205e"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dynamic\_proxy, http\_proxy\] Configuration parameter for dynamic proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("disable_dns_masquerade",
    "enable_dns_masquerade"),
  validators.ConflictingObjectAttributes("http_proxy",
    "https_proxy"),
  validators.ConflictingObjectAttributes("http_proxy",
    "sni_proxy"),
  validators.ConflictingObjectAttributes("https_proxy",
    "sni_proxy")}
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
  "x-ves-oneof-field-dns_choice": "[\"disable_dns_masquerade\",\"enable_dns_masquerade\"]",
  "x-ves-oneof-field-proxy_choice": "[\"http_proxy\",\"https_proxy\",\"sni_proxy\"]"
}
```

OneOf alternatives in this subsection:

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-3a1c759354daefc0018b08ac5c61909d8f79022e190b394d77c9912d58f9205e)
- [http_proxy](resources--proxy--reference--group-004.md#canonical-0b32a135f0eac087814b5eeb5a2baa7a8860790f0ef07edeea58f4e6eb6bfb9e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dynamic_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3eee18ba00812b782933b23c644fd886d7b3c40672f31bc27a647554168929d"></a>

## Direct properties — dynamic_proxy / e5457e1eed7c / 3

- [disable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-994b5b67b50d6e3a5907a17310db104d483b17a3cb547c4c4826a35062d7c618): complete subsection reference.

<a id="canonical-bc04064377a81302eab5bf8340748f8b7c23850a2e0ccb62527ffdc07392c044"></a>

<a id="canonical-7105477b6c2dfd50c3505e0c9c06b669baa96a17f35957efa360b80ea49a470c"></a>

## domains property — dynamic_proxy / e5457e1eed7c / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

- [enable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-a91ff0d541419c07f0afbdae794ff53b3971551f38ba126e40e60413dd19aa8c): complete subsection reference.

- [http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93): complete subsection reference.

- [https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3): complete subsection reference.

- [sni_proxy](resources--proxy--reference--group-004.md#canonical-d6201277c11d95edb6c35bf0110b8dedb94e654460ba57d9257c01c6cadaca1e): complete subsection reference.

<a id="canonical-4b72b0038a32c76c9f38eb0a0689aba0b0217ed7eeb3181b8380f6410b8f0d23"></a>

## Next pages — dynamic_proxy / e5457e1eed7c / 5

- [dynamic_proxy.disable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-994b5b67b50d6e3a5907a17310db104d483b17a3cb547c4c4826a35062d7c618)
- [dynamic_proxy.enable_dns_masquerade](resources--proxy--reference--group-001.md#canonical-a91ff0d541419c07f0afbdae794ff53b3971551f38ba126e40e60413dd19aa8c)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.sni_proxy](resources--proxy--reference--group-004.md#canonical-d6201277c11d95edb6c35bf0110b8dedb94e654460ba57d9257c01c6cadaca1e)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-994b5b67b50d6e3a5907a17310db104d483b17a3cb547c4c4826a35062d7c618"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a7cdaec8992711e2a0c0f227b08f1f4912f6aa467bfd2daf7fc14b9b8238521"></a>

## dynamic_proxy.disable_dns_masquerade — dynamic_proxy.disable_dns_masquerade / 5695215e3dde / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- dynamic_proxy.disable_dns_masquerade

<a id="canonical-7ee761d9e86ff17fe98aad73968ba4f636f170f4d9733dc35c330a77013ce29f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_dns_masquerade = {}
```

<a id="canonical-326d992554bea307ab5bee6866bd0104b02e130a69893d293d20a137c9cb74ff"></a>

## Direct properties — dynamic_proxy.disable_dns_masquerade / 5695215e3dde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb01fd09438450fe7058980a6d773bd92e1f50d45d44ae51e37ede3dc182bed2"></a>

## Next pages — dynamic_proxy.disable_dns_masquerade / 5695215e3dde / 4

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a91ff0d541419c07f0afbdae794ff53b3971551f38ba126e40e60413dd19aa8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bb4326b82b30f3ff470e368b57269e8df92fb6c6f079a7a64ef9ca8e0e80a0e"></a>

## dynamic_proxy.enable_dns_masquerade — dynamic_proxy.enable_dns_masquerade / 64201319be7b / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- dynamic_proxy.enable_dns_masquerade

<a id="canonical-3841990f265b71c08884e49958b237d67360e474abefd9ab959bac573846bf32"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_dns_masquerade = {}
```

<a id="canonical-9cd81c7b9e4385bd7096dc548308379440bca5835b01475f1c56e2e1fbd38a46"></a>

## Direct properties — dynamic_proxy.enable_dns_masquerade / 64201319be7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-524fb197577902ef5e666ffa63a950d92b4769a7782af4b21469fdb0b7b417d5"></a>

## Next pages — dynamic_proxy.enable_dns_masquerade / 64201319be7b / 4

- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cd28658bd15d49d110cd2cc860fd39433d77de8bac3151a67197188639fd1df"></a>

## dynamic_proxy.http_proxy — dynamic_proxy.http_proxy / b00b9aebb5e2 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- dynamic_proxy.http_proxy

<a id="canonical-0ce9745496e96287decb192d617341f375a79913024edc9e018815ee82701bea"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ffde9de0f58a45bb660c168e1a49aef7eefa1aeeb7283d6bb59b69204f238ba"></a>

## Direct properties — dynamic_proxy.http_proxy / b00b9aebb5e2 / 3

- [more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae): complete subsection reference.

<a id="canonical-f5278afba1b169b4cfe4d1cb3f306cbc53b1057dbd09eee8667aa3b87e1dac17"></a>

## Next pages — dynamic_proxy.http_proxy / b00b9aebb5e2 / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4309bc7786f5b88627680938f6cd9bb411dc522d23ecadcec443fef9299b68d9"></a>

## dynamic_proxy.http_proxy.more_option — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- dynamic_proxy.http_proxy.more_option

<a id="canonical-86f6c0fa544bab23ecc8c61c30ad0df27e1e8e58a32068578a6554d272303c0c"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-a4a394e5c60982c8bff776267e20ff98593ac4e974d353bf82499920d046ca8f"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 3

- [buffer_policy](resources--proxy--reference--group-001.md#canonical-1b83a3550cab13a50f63f4eec743575e67ebc943279d4421284cdf13b8c9a7a5): complete subsection reference.

- [compression_params](resources--proxy--reference--group-001.md#canonical-a03664ddf476be7947a697b484a87b5977879a1c22e704a02565c8d2be2dc84c): complete subsection reference.

<a id="canonical-033cb53ee0b885ddddfbf577ff345c839203772b1cf073f71f358e3ff8bf4095"></a>

<a id="canonical-6a308cbe64d2931e3555e8aab11cf3b687e8a1c8788d5582d3f4e212446b801c"></a>

## custom_errors property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-31d5780c2e7893c90548be06c3537aa41c2efae8d9c8fc69707b084431178062"></a>

<a id="canonical-bc17092f70813185ff86284becffb3ee0a23a48cd6ae1cbc6a4671789e9a35b6"></a>

## disable_default_error_pages property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 5

Type: `"bool"`. Optional.

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

- [disable_path_normalize](resources--proxy--reference--group-002.md#canonical-c4ed3745a27a9c81ec440e1267ace3b87d85ff6c39d101721de6286825f46fb6): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-002.md#canonical-a7f49efacc90aa2781df43b8b4993ca69b62106dab54afc06f452dd61e52af5e): complete subsection reference.

<a id="canonical-a4361a8f620ce684c91f51d713342885a45501b96ccc5c742ccc78dc1fd5fce6"></a>

<a id="canonical-b0379b26c24639bd6f04ba3a730dd7238d77e6c7599aa451f776941364dc8530"></a>

## idle_timeout property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

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

<a id="canonical-42af490ac9a317852867e19cf4f44d1769a0a864374ddab23ba6d0c9aeb45359"></a>

<a id="canonical-5c33ae8763c06d4d92663eb147f9f0bf43f767f296258cc5c2cbe5602eaf7d09"></a>

## max_request_header_size property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

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

<a id="canonical-30a186f4ea74ef8d4bfa12369133a98140761ba104e29769da647387829fa7d2"></a>

<a id="canonical-c72fa9642975011920438be690853410bb650b9152fefa26379b4512c363cc09"></a>

## max_requests_per_connection property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
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

- [no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-7079c85d332c593da75961184f9050dd1376f461677914bad6e33fc70f9ad482): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8): complete subsection reference.

<a id="canonical-eede7123dcd0bbbb5af8cffc0aaa3efb8bf112eaebf2f4a825c4bf3c690a7abd"></a>

<a id="canonical-59d8c4ec489da1bcfb0f4f176c59e4db774a2ab091799bfbb20ceb6b276c9ff7"></a>

## request_cookies_to_remove property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 9

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e): complete subsection reference.

<a id="canonical-3781d60b585cae05cfd10c5671102dcb88c9068472debdaec693cd0d9aa029d0"></a>

<a id="canonical-c2e5327aa29ac90f7a92bc8d62f4e1a229fe53810c16b78b69e57fb2612e028b"></a>

## request_headers_to_remove property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582): complete subsection reference.

<a id="canonical-0584a3744e02ad7b62ea0524b6ef902383a7ffca562ac80cea995d09c4b2a32b"></a>

<a id="canonical-cf994788c104d16663de23bf6b017d1f39209591db65794c616159f16c9217c8"></a>

## response_cookies_to_remove property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 11

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

- [response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0): complete subsection reference.

<a id="canonical-38bc27b46867398c9b3310eb150b7e37f468b62a6d5630df14bd3474e1c87912"></a>

<a id="canonical-fe2d4f68c4e03f03b00c90c031e5d6e1a20b7bc942dcd8dfb1f1e3dee82de69d"></a>

## response_headers_to_remove property — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 12

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-37e19a2930177bf2186696c2fcb6f273b5330cbb2885f40e39a0a7d563b18092"></a>

## Next pages — dynamic_proxy.http_proxy.more_option / ba2519edce50 / 13

- [dynamic_proxy.http_proxy.more_option.buffer_policy](resources--proxy--reference--group-001.md#canonical-1b83a3550cab13a50f63f4eec743575e67ebc943279d4421284cdf13b8c9a7a5)
- [dynamic_proxy.http_proxy.more_option.compression_params](resources--proxy--reference--group-001.md#canonical-a03664ddf476be7947a697b484a87b5977879a1c22e704a02565c8d2be2dc84c)
- [dynamic_proxy.http_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-c4ed3745a27a9c81ec440e1267ace3b87d85ff6c39d101721de6286825f46fb6)
- [dynamic_proxy.http_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-a7f49efacc90aa2781df43b8b4993ca69b62106dab54afc06f452dd61e52af5e)
- [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-7079c85d332c593da75961184f9050dd1376f461677914bad6e33fc70f9ad482)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-a287f58157beee3469042cb53660d86aecb0a9eddfeafec4fe529b182ce1caa8)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-422b9d3942f2fee3a2314415b4e7369985e61622af450adbf5b04e6dfd697b9e)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-5fb8bb7744eaee450b461b600ee40be0351260fb69bf0499ba651c172a799582)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-66b189fe22f6e1bd709169dff08aa89c4f8e277b8ebcecd38987d36df07db0c0)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1b83a3550cab13a50f63f4eec743575e67ebc943279d4421284cdf13b8c9a7a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10b31e4d814f6b0f04964530a04a3a7161a54ec780a7f2699d9c9f5cf132cd7c"></a>

## dynamic_proxy.http_proxy.more_option.buffer_policy — dynamic_proxy.http_proxy.more_option.buffer_policy / f6609fe394f7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-c432cc2f0a18c7c72984dd3233612c0c9ba48d6e45fc6079c32970eb2ef2ec93)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- dynamic_proxy.http_proxy.more_option.buffer_policy

<a id="canonical-90d28e56f19585d78a0f31fea579fbc95d05e8aa4399c90d8712018cc3bbfca3"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-d64aeac428307aa255971a08c5bb1d4b4f4ac1fb27ec617afa9106ee75af411f"></a>

## Direct properties — dynamic_proxy.http_proxy.more_option.buffer_policy / f6609fe394f7 / 3

<a id="canonical-7697b8f0b3469dcfd35d23aab043076491869e276e3a800257ba088ede9354e5"></a>

<a id="canonical-c9aa41171ca588849d9ccf46342b3995a09847185b70dd81025aa9e3706ad122"></a>

## disabled property — dynamic_proxy.http_proxy.more_option.buffer_policy / f6609fe394f7 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-a975dfdcaade03373735b38bc64925e4c284bebdb8947f2e31519517673100ef"></a>

<a id="canonical-f8223fa7209c24d0f66b75bdfa2d4b2a3c9090ce06bfac37d43377ab5697e981"></a>

## max_request_bytes property — dynamic_proxy.http_proxy.more_option.buffer_policy / f6609fe394f7 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

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

<a id="canonical-c4722417e9b2f3ee6a2479fead27b5576ddace3e63a328d7a9b30f5e0de7b4f4"></a>

## Next pages — dynamic_proxy.http_proxy.more_option.buffer_policy / f6609fe394f7 / 6

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-c46b38256a2384f52943a76645f28fee5e4f6a4ade52a3e7d9111e27ec7a65ae)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a03664ddf476be7947a697b484a87b5977879a1c22e704a02565c8d2be2dc84c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
