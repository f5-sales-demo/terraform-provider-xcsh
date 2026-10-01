---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-c5169e0d5d61fa4e69ae226ae818f3a5e22dee8df56efd690a51e4425f326694"></a>

## name property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / 01fb2ff637e3 / 5

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

<a id="canonical-9240580f393890de6db2b237da74565da93d9a760f442a23815c58ef9c92c8eb"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / 01fb2ff637e3 / 6

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-f2dae3ad02022938cdde6747f41ee3930084d66fd26e741799c1376beffd7377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4ebd2bdc216db4a1646a08b696c2beb3ef9d4393682471d1a5ca3ea299ff22d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375519f6f2ff412e278cb0c548dc6c8b06dbc0c2f98756c34f01cfc9b5973a66"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-f2dae3ad02022938cdde6747f41ee3930084d66fd26e741799c1376beffd7377)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path

<a id="canonical-96d9d456f78279a5b57b7db9acc7fc782f62a679eac90d09ed6cc8f0d8adb800"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-cec188e2c8017413a5535f039f6e9df5ee66974b23c8cf1052819e443c4fbec3"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 3

<a id="canonical-981d2ccea80168a6262e592ce0eaf069436d3e6527b5f7a9dc5aa6550b1f6f6e"></a>

<a id="canonical-e811c582991296203258da7f5559465b5d4fa1d0a720ec84378e264bc0d3e4e0"></a>

## path property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-010e05988801fb332ab59edb16084a9c31667a21e16a3a2cc9e09b66daefe828"></a>

<a id="canonical-225e486240666615d05d5200242e623153f76df5536a52ea22a5734c3f0e68a9"></a>

## prefix property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-25478e83517de2fad0bb6eb8353eec9a06c17d9b7e31b689c5fe9a84e47f810e"></a>

<a id="canonical-cf319a77ce24b0d85033aed9ae0c7fb43dcbc013d435084ca00e05494b65acf1"></a>

## regex property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-37dfd17aebb91a2b61b3d8efe7f06ecd3f9e797ad6860b946b66bc0d19afcd41"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_l / b2ccbf2483f8 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-f2dae3ad02022938cdde6747f41ee3930084d66fd26e741799c1376beffd7377)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da00314637d67e2594beebf79a9c0798386682bcc329e8855bdac1ff01b2ed08"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules / ef6df91799a6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules

<a id="canonical-834e448bfd4f970abdcb062cf0e16b0d57abe13d6c39a05ceac599c5ffb4eff2"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-448e88f3526c9bbc1ddde36ab979054aef2f04b87a94eac0e36fec77dac4ccd3"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules / ef6df91799a6 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-82717e4e9891f1b05143542d751f505e22aad4e10ce2459eadb5d1b7d57537f0): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-0b274c734c3b17a50d2765a65df0e8c709c649acc33380342fa4528c7246a750): complete subsection reference.

<a id="canonical-b06a49b1d1a3d028e3e026347539c2a854b9ad1a06b381c59a226df61b17ee10"></a>

<a id="canonical-116c7ee7f2d6c22747ad89132d63aafc10ad845c37654e5204eef2ee1765d51d"></a>

## javascript_location property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules / ef6df91799a6 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-ec58829b62b0facee7e295f2c18a36f389c10b906af58b993516e2b8873474c1): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-013.md#canonical-51aa5867b4719b68a290dccfbd133b846fb0beb52d0049402feb1694792d2e8a): complete subsection reference.

<a id="canonical-d69089f2d3edf621c1201cbdb86ccf9622830d55b31b910f7ae9e963c5ee27f4"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules / ef6df91799a6 / 5

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-82717e4e9891f1b05143542d751f505e22aad4e10ce2459eadb5d1b7d57537f0)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-0b274c734c3b17a50d2765a65df0e8c709c649acc33380342fa4528c7246a750)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-ec58829b62b0facee7e295f2c18a36f389c10b906af58b993516e2b8873474c1)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path](data-sources--http_loadbalancer--reference--group-013.md#canonical-51aa5867b4719b68a290dccfbd133b846fb0beb52d0049402feb1694792d2e8a)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-82717e4e9891f1b05143542d751f505e22aad4e10ce2459eadb5d1b7d57537f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca6d8303ef91530e9fca8ddc2f203c2c7d5fca13f66c5902841f80bbcc88489e"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any / 23e4f005ee5b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain

<a id="canonical-627e0516c0724e3916dd305dccd2d31d9d36dc2091f3fb20f6ae49d65a595c7b"></a>

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

<a id="canonical-b7df604820f3c383ae12efd79cc8987956cedcdd9f778628e8e16d7fa44f12ad"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any / 23e4f005ee5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25a44e4abf17708e917c0afa75ec811cac3f8b3ca567ec043a3fa11317de7d25"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any / 23e4f005ee5b / 4

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0b274c734c3b17a50d2765a65df0e8c709c649acc33380342fa4528c7246a750"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfd81bc194eeefc6da726752515388ea06e670186fee8e8f71e61bc205e7ce4c"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain

<a id="canonical-352f0fafc92945061879599918428535279c904831492074f4931da05c44d393"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-bfde9a9a31e7f464f79bb08c25b31bfde570278559be286c0f1bc5547f42dd6e"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 3

<a id="canonical-ee57f2386ff8caf04c864c6261dcd29fd24e83300abeba3593c8514956f7aa77"></a>

<a id="canonical-44d6f57c929596d377c9e0960005bcef3586905cea146f48bf6d9aef0df6a3aa"></a>

## exact_value property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-42f40aa09ebf24ebfc25cd5e7aa4e5a53734b3804ada9d2fd5fa41461d267c03"></a>

<a id="canonical-ebb86973b1b866beaa06dda4d8ea76fa5faf0e9cef0ca3b3ec1363f4cdc9c03d"></a>

## regex_value property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-a99ec147296b72b08516c841b54514b2f2190c8388a7e69caabd00bd9c169047"></a>

<a id="canonical-abd7beee202af7622f006a62cc8efd1f3d50f70bff30103652548e9d5403a589"></a>

## suffix_value property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c8c7c63cc41ae4580476c6e2e5c446d41847413f3891c4c83a4a2f21cadfa223"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.dom / df49c19aca5e / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ec58829b62b0facee7e295f2c18a36f389c10b906af58b993516e2b8873474c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c1df9ba88fe17015f3e02cb45349a327068ba097574aa84092957ce6789ac26"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.met / b86e95a096fc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata

<a id="canonical-15c1a0f92d6bfb2f095dc1e9210ef0d53264ee5029f2e4fcf92247ca078bf982"></a>

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

<a id="canonical-566f5219ab0de904b1bde8d1c57e82658ed94c5c0a0dc8783e9dae44d0699b09"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.met / b86e95a096fc / 3

<a id="canonical-02a2d61ff4e235f9a1f17777090c1275a2aa5d6e5402ef4dfb23ebb326ad7116"></a>

<a id="canonical-ec72023e6ed95b612e592def37e9028af2b6cdd842f32024bbc342e7e3b8a8fe"></a>

## description_spec property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.met / b86e95a096fc / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-5755a51ee86b5cb2070704faa0b0375c1662bb3b69250f0aeefbd86d412c6f82"></a>

<a id="canonical-b1ee3932d05fa40c0bff9210383ae8e0ab134d4ee20761a6f583576b71dfff52"></a>

## name property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.met / b86e95a096fc / 5

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

<a id="canonical-9bfb388185ecba6f6eec2224a14d49798d700d9330bb1a6ae2c470eb09655a88"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.met / b86e95a096fc / 6

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-51aa5867b4719b68a290dccfbd133b846fb0beb52d0049402feb1694792d2e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be27c2e40b8706523379b17689d0927b381fe5c8a2373c58dc5b2582eaabe583"></a>

## bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-650a4261a2eb07348b4287e579ffd90a6d7b5be3a66dca13b5323383792f8e58)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path

<a id="canonical-671e18320cb9c2d7477e1c90d69f7f2f8437bc647edfa70424c0922669165f8d"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-1824bd059d82ec674ba5adffa306c648190f6195f8565988d3d0d63be2a6a877"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 3

<a id="canonical-63b949081cfda4cdc5e0dc4a7f5e0821694f270a4ad94e2db170b88b67db30fe"></a>

<a id="canonical-40b94d369e9073484aac7463f45122450e0187a6f1d4d8c0a44a26d2e5b37cbb"></a>

## path property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-fe284c7abe1e7f975109e80b323c338dc2e22eef64aa1bef6d346c925abd8bfe"></a>

<a id="canonical-5144727f6f8a2c4c3584f3ae1013638781acefead871f1cc3c765b7f8edb584b"></a>

## prefix property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-f4111e3b43bfb7d7cfc45a1b1edb27f35e4717d1a95868fc7186e30f95103b1c"></a>

<a id="canonical-d85d301fb39bbc1281d5e5fc90b89935d3b76e59e188e7321cfa1c8adc1b6607"></a>

## regex property — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e8b95ffa56bcfc97a7ad8de85d6ed591c2975e17b14f146ccb41358e9bac6bb6"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.pat / de630568fcc8 / 7

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-a0571087c2a049ebf2eee426cae76d45413f209c6d434181e11e5d7815e74adb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8c014284d0fc758affbbd6b21b645311b1908ccf2b00227def32fc5d4dbf211e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a6fdee3e000edba41f9e6e9a171a6a9d422bdfcc209287e1145b1fea236a72a"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- bot_defense_advanced_protection.both_web_and_mobile.mobile

<a id="canonical-f093e85fe6644f48e8f8f9ea61ce4ebf5851f5efdb2076c17d89d724a492300e"></a>

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

<a id="canonical-2ec411420d4c6f08ce15dd9f7cf0cb12ef1067fe3edf097bc382b0c93b2bb81e"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 3

<a id="canonical-2821fc8754cf4191150140394a11c84559452a55b86b9515eb41a2e75f21a9fd"></a>

<a id="canonical-b25d5cfdb81d9fe2e19ccda4907535175879e4cc9acc84bb6cda23c3bf3489d8"></a>

## name property — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 4

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

<a id="canonical-e7a6885d4a97eb2ca55c95b8f5e03becb2cc627207d076e07de9ed54e480325d"></a>

<a id="canonical-e25895fbb7629449d9e8a37a247da9653af07f5805b8c5ebfcc654b6f953dbbe"></a>

## namespace property — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 5

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

<a id="canonical-68ade6592a6eeca7eea82f4f2afc3882a903f90ecca230de3120fa59da6fa19b"></a>

<a id="canonical-56256a81ddd8ce961d98738c6cce6b7f51c1022eda23c8a93b2e5fc75a655d55"></a>

## tenant property — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 6

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

<a id="canonical-de47f828f6d8b93d2a576dd82a117ad4c67171f8bb4636c0be450e22346445aa"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile / e320eedd5465 / 7

- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29b104af4bf0dc01c2c9e7eca83aae53f29a24141d979c2acfc32ccb6298b56d"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config / 0fb15d016ab1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config

<a id="canonical-d0e0ba64f5444d77db8d458a523a7cdbc12ce974dcc7f370415ca7e809c1099e"></a>

Type: `"single"`. Computed.

Mobile Request Identifier Headers. Mobile Request Identifier Headers.

Upstream description:

Mobile Request Identifier Headers.

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

<a id="canonical-b9056fcdb5fe61487ab219024b983b45a7a25ad51168110c6f4bf081f02acce3"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config / 0fb15d016ab1 / 3

- [mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d): complete subsection reference.

<a id="canonical-21b788f45b9febf3f1b06e109fface233f2f7de6ac20dd3a15a801b3d482468b"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config / 0fb15d016ab1 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60716148b4c87c10c2c92680e0c29f0f87b4a1486e3719d67fb4bc79db69a0fd"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / ea3322863267 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier

<a id="canonical-7a88f211c0334684b47bb24dbe2828f5bdf4012dea6652199c66705e9a3d895c"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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

<a id="canonical-6f571906717cb1f9f45ed3669fbb496bdf4060db10aef75ccd62bb49533775ab"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / ea3322863267 / 3

- [headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1): complete subsection reference.

<a id="canonical-93045153a9ebad9936f08c6591d1ddbd5361f0fee7266c9f854f213317955daf"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / ea3322863267 / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16465bf6819d899c3309e71a37517b20cad381d147440a8f1e397cdc887e0c06"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 6a6b11b8c7ee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-9c8dd0a7ac277dbe624c88941314a1c18e2693b737b2300d0a8904777019e52c"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

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
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-46712aa323540174aec8e92d4f39d14f51958aa153bc7cf71b46643c03b23b9b"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 6a6b11b8c7ee / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-013.md#canonical-0db967517dfe59a18fbf8cc5916ee0a6f422fa2f67d6f6b8402fa02b6afac32f): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-013.md#canonical-0ce6198136bc2ac902d6f8f3020ae1b79f6791f31f4bf57d1776487d141a065a): complete subsection reference.

- [item](data-sources--http_loadbalancer--reference--group-013.md#canonical-b97c24260d14238a10420f504bf54d4a296b362c8104a1a2f132a7903cd72ced): complete subsection reference.

<a id="canonical-eaa1b7499d0ef540249beda0d178e7d9356f1fcf8a53e886aadaff35ed6fd44c"></a>

<a id="canonical-24f434bb77cc1db45096244a86d2243186b0175ed7e8cab0e9c92dc2f06d39eb"></a>

## name property — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 6a6b11b8c7ee / 4

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

<a id="canonical-6a6b8802cb4ea9612645020b033fc1dc853dc07472a940e5970ff8496f823c57"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 6a6b11b8c7ee / 5

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present](data-sources--http_loadbalancer--reference--group-013.md#canonical-0db967517dfe59a18fbf8cc5916ee0a6f422fa2f67d6f6b8402fa02b6afac32f)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present](data-sources--http_loadbalancer--reference--group-013.md#canonical-0ce6198136bc2ac902d6f8f3020ae1b79f6791f31f4bf57d1776487d141a065a)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item](data-sources--http_loadbalancer--reference--group-013.md#canonical-b97c24260d14238a10420f504bf54d4a296b362c8104a1a2f132a7903cd72ced)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0db967517dfe59a18fbf8cc5916ee0a6f422fa2f67d6f6b8402fa02b6afac32f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e39aa83f50faa4ed405ae63f75d72fe81a4e461be80c0e9a46e94894830fa58"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 43ae659685be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-056ef6a730f9ca750fb2e64a38764fa1bb1818979455c5fc5ce93530dc971707"></a>

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

<a id="canonical-37e8ecbc4ff80a9420f85a9d66069e7ac068b4aa58b3b18db5de710f36ce0d25"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 43ae659685be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94fac1e91c58733f6b32f400088a794014f6c8f58453a7774df1f5e4e8fcb622"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 43ae659685be / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0ce6198136bc2ac902d6f8f3020ae1b79f6791f31f4bf57d1776487d141a065a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64cb96306fcc756194baa5708dbc8e8eab573aaae58f77ab05b4d04844992a39"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / f53d675c12fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-e452142154c7ce3e41a948d2965cad143276975cb06cd0f8942b55318734edcc"></a>

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

<a id="canonical-45ca3827c55985e96eb7589fa186a16b404310a7c9a2137a0a79291e7372c3c8"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / f53d675c12fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b596a7b6515e94731d4bb2482c26e96693d46ca74c13091dd3a68bdea5525a6b"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / f53d675c12fb / 4

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b97c24260d14238a10420f504bf54d4a296b362c8104a1a2f132a7903cd72ced"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4e09beb1b01a2dba33f38bcfc328b6f68fa80531a4e6eecc80f375b672573f2"></a>

## bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-013.md#canonical-9196a84c4b6a62f18458b9577ac32708c1ee61c8d748802fcdc651ec0c5aa971)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-013.md#canonical-befaae2a00aee96910d2725f1e9e75b366f161724b4de9138f5a8dd27996a71d)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-d06bc7954f70e99969085dd84c3fe66a840e917e7847abfb3d9cc7e95379dd7d"></a>

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

<a id="canonical-b424b3c5d9638e387e2f5dbd3da2683efab90b5c09bf2dde794e59be125761f4"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 3

<a id="canonical-f228cc04981cdcf2b87fa07db9027875b44a670989fb2769260d3f99b258fdce"></a>

<a id="canonical-79c79ad5a41334c00b1197ec559a14f8f942c0f75b74bf186ca0f522bd39318e"></a>

## exact_values property — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 4

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

<a id="canonical-6859c1d6adcffc6c1974add3357f6968dd90a85d4c412cdd7b3ab70a37546334"></a>

<a id="canonical-155970555aace4f25d104b8ee5f4f1e3193355e9e5f4337b22166343f09a7b48"></a>

## regex_values property — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 5

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

<a id="canonical-3644ce75440c8ced50cbd8427dc18fa1f5429c704d83f89795ca4285b4924a6a"></a>

<a id="canonical-c216f5bf149eb1f7bd30e51df461eaf91e49ae06615a35566cda8145031b2654"></a>

## transformers property — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 6

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

<a id="canonical-72cdbac2c88433673574636097e414cd67097141c588f8755320d06f38ea2a0f"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_ide / 4f3d5af4ce0b / 7

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-74fcc788dee31f42cc92e19cd194b705d22d4c25d8fcb9027cd69baf5b610bd1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-23e613eea4e58145f953449c30e22b147de3622e4e088180e40dff09aa6f48d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-680a31fff14d3672204593f315e0206d17601a319de982e2e324b316bff35340"></a>

## bot_defense_advanced_protection.both_web_and_mobile.web — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- bot_defense_advanced_protection.both_web_and_mobile.web

<a id="canonical-b1c42539abed93f7a7098be1f00eb7c6a082f10a9b657ee4e42040a7428ee995"></a>

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

<a id="canonical-ff695688ebc7a77279d0f9861004f8c1d4cf1931084293f93b1b64e3e601c089"></a>

## Direct properties — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 3

<a id="canonical-61d37b45e90cea1fa06a57b419ca3aa558f8bb224554404da73f73e54c7c3e7a"></a>

<a id="canonical-125c11540d122c45ca0eaf3ff63109cea9ae75154b801dc80b337455b7d28ae1"></a>

## name property — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 4

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

<a id="canonical-95f3a6300ca86b464a3644f9746cdd73b0198bc5d7067c23ec1ef5b61a7c51c1"></a>

<a id="canonical-3df89ff667e89a2fd37e20bd6c4984bb3b74eaafafed89fd438896e1993237c8"></a>

## namespace property — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 5

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

<a id="canonical-85ed5280718016beef10b09263297a9bfb5f4587ddc658dc1f4d5aa7f1a15faf"></a>

<a id="canonical-130b22e2c1a5a929023009684041c240199dae024c5ac838b2a3d22585676f4a"></a>

## tenant property — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 6

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

<a id="canonical-e48427fdeb849f14ab52e5016154eaea984ba8353ac4d1278b1cdc22f0293832"></a>

## Next pages — bot_defense_advanced_protection.both_web_and_mobile.web / 848d9c096494 / 7

- [bot_defense_advanced_protection.both_web_and_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-1fa24d437f01d9e7ffab821aff2e867cc51e336201dc5670a51496596926d543)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3e3bc1b0e0d235c77c8799afc71255c645299d01286b9e95bc84bd8aa55d87be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb1307fd13968dc3337dc9ab4f0091b625d57770fd1055497da1b50fac5d3c71"></a>

## bot_defense_advanced_protection.mobile_only — bot_defense_advanced_protection.mobile_only / 0015b0c22385 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- bot_defense_advanced_protection.mobile_only

<a id="canonical-284987e944502ed26a4c732f74fd9743f3a5c5adbccb29f23510fa37d7e77ea0"></a>

Type: `"single"`. Computed.

Mobile. Mobile only configuration.

Upstream description:

Mobile only configuration.

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

<a id="canonical-1f5261764177f172474e337490453e839c27b78cc768391203b67ed3274efa59"></a>

## Direct properties — bot_defense_advanced_protection.mobile_only / 0015b0c22385 / 3

- [mobile](data-sources--http_loadbalancer--reference--group-013.md#canonical-ac3fcb7f587fe55ee84b06f0258dbc8a3606c53a07159a3e3447254a70c90c42): complete subsection reference.

<a id="canonical-9adfd4f4a5bb9811c58069289e6779010acad17206f487a5ea47ef4975a8abab"></a>

## Next pages — bot_defense_advanced_protection.mobile_only / 0015b0c22385 / 4

- [bot_defense_advanced_protection.mobile_only.mobile](data-sources--http_loadbalancer--reference--group-013.md#canonical-ac3fcb7f587fe55ee84b06f0258dbc8a3606c53a07159a3e3447254a70c90c42)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ac3fcb7f587fe55ee84b06f0258dbc8a3606c53a07159a3e3447254a70c90c42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c84229ef41280a82febac5d49ab2bfa39b5d2366a1f51ed0a4ff628559445cf"></a>

## bot_defense_advanced_protection.mobile_only.mobile — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.mobile_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3e3bc1b0e0d235c77c8799afc71255c645299d01286b9e95bc84bd8aa55d87be)
- bot_defense_advanced_protection.mobile_only.mobile

<a id="canonical-ce411323a7fdd003dda5ea2e6223ae311563468b3ce1935f897767fd364fc3de"></a>

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

<a id="canonical-0dc3f9ce66c663f8d3b5d2dff8577eca15e934d90fe6382db49dddb0bebb354d"></a>

## Direct properties — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 3

<a id="canonical-291763179b7b02cf219632143c8785f7391ae9e3a23b2f78f634b85fc5c4c066"></a>

<a id="canonical-5be916cabe5fc4817eadb3cb325b9fdb8f6f1f4074021c62d71e3c1f59d03e07"></a>

## name property — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 4

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

<a id="canonical-686ea3ce71d5ba9e8a57e95b77cff1e78197000e7cb1ec33d4a03fbf675c6a9a"></a>

<a id="canonical-5c04dd918672bed304fa2359842c003a027c8e4666fd7d5a6e6121e26d0046c4"></a>

## namespace property — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 5

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

<a id="canonical-cde8df7540e69880574fe57ea2e0a97fc1f9003a1aa989b302d635fba4d0d790"></a>

<a id="canonical-a98c8ba28aec3609709cc17d736fc2d3d1081d2964cdb32663444ef09a44e08e"></a>

## tenant property — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 6

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

<a id="canonical-ec83a3421744412780adc62e996fff080a330b69b2d73e1deab900669263f468"></a>

## Next pages — bot_defense_advanced_protection.mobile_only.mobile / 9a22b71e4efb / 7

- [bot_defense_advanced_protection.mobile_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3e3bc1b0e0d235c77c8799afc71255c645299d01286b9e95bc84bd8aa55d87be)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aca2240ac0eba7c9906881fcd6f0d89f48ee54b76a86143a826d01ffa19a24bf"></a>

## bot_defense_advanced_protection.web_only — bot_defense_advanced_protection.web_only / d6a4b8579572 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- bot_defense_advanced_protection.web_only

<a id="canonical-aebc0abfbd9d31d3425eb5f1c428ff6995028e8a9c8acb4c74ad867ac2e0e8a4"></a>

Type: `"single"`. Computed.

Web. Web only configuration.

Upstream description:

Web only configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

<a id="canonical-98719a39167500c81b34dfcaa333659aaba63460fb5dbd685dc7f00f1e0af0b0"></a>

## Direct properties — bot_defense_advanced_protection.web_only / d6a4b8579572 / 3

- [disable_js_insert](data-sources--http_loadbalancer--reference--group-013.md#canonical-a19e7f878d65ebe752500a022aba761302988633b54329cbfff2698d55547d34): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--reference--group-013.md#canonical-baed21816e8ccf268426c1302c537a84f54759aacc5d91c6ce83b2b5a285ee14): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9): complete subsection reference.

- [web](data-sources--http_loadbalancer--reference--group-013.md#canonical-fdbf1bb4c01d654be1d229ae0065c538d72a584824a8926c680fe48bb899031c): complete subsection reference.

<a id="canonical-62d01e38d73cf12ced2a3efe2d6e8a8e932e2b65b329b088891904c064868fc6"></a>

## Next pages — bot_defense_advanced_protection.web_only / d6a4b8579572 / 4

- [bot_defense_advanced_protection.web_only.disable_js_insert](data-sources--http_loadbalancer--reference--group-013.md#canonical-a19e7f878d65ebe752500a022aba761302988633b54329cbfff2698d55547d34)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages](data-sources--http_loadbalancer--reference--group-013.md#canonical-baed21816e8ccf268426c1302c537a84f54759aacc5d91c6ce83b2b5a285ee14)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.web](data-sources--http_loadbalancer--reference--group-013.md#canonical-fdbf1bb4c01d654be1d229ae0065c538d72a584824a8926c680fe48bb899031c)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a19e7f878d65ebe752500a022aba761302988633b54329cbfff2698d55547d34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0a280c9da084ba770d6db805bfc56c76ea2cb134aaf5506fa9e8fccb8915281"></a>

## bot_defense_advanced_protection.web_only.disable_js_insert — bot_defense_advanced_protection.web_only.disable_js_insert / d04c70aec4bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- bot_defense_advanced_protection.web_only.disable_js_insert

<a id="canonical-5f539b2738d606700e383a76e8325868b1fa01930826f1f1bad4c4fa46f7a06e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-a6ea92e25af35aee1988179a45f309de0a0ca3ecce5a87d6ce5fff4541b81561"></a>

## Direct properties — bot_defense_advanced_protection.web_only.disable_js_insert / d04c70aec4bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-534db6a73a16d721606abc612994d0e1fd8a89717c33498cdbe36372b825f164"></a>

## Next pages — bot_defense_advanced_protection.web_only.disable_js_insert / d04c70aec4bf / 4

- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-baed21816e8ccf268426c1302c537a84f54759aacc5d91c6ce83b2b5a285ee14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72c6d6b9544c4aedccf0529a5631084508681b13bde1ec89668799330a404be7"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages — bot_defense_advanced_protection.web_only.js_insert_all_pages / 6f4fa540c61a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- bot_defense_advanced_protection.web_only.js_insert_all_pages

<a id="canonical-60fce55109e0f4790d2e718e7f87871094432b182aa764d6ed92a813ed479144"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages.

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

<a id="canonical-0c8674717115914c98fb8a92177ff16f2b6c2c848b620f043b6ea8e6d43d84b5"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages / 6f4fa540c61a / 3

<a id="canonical-ff96518d34b9c81d6452071b1dae547cd4ce3c982b4754e599d4c22901e0ff1f"></a>

<a id="canonical-ad401fe417c8fd6be623dda4d3bdb14ad3e370e239bfb8a1ad066c92fe3fbd7a"></a>

## javascript_location property — bot_defense_advanced_protection.web_only.js_insert_all_pages / 6f4fa540c61a / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-43c4042011e66dc8c9256e50b1548a39822c207f713565928f170fc21383bf5d"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages / 6f4fa540c61a / 5

- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2017adb2fa684248c00748cad6f55f86ee6909366b3fee97e958f50598f21f5"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except — bot_defense_advanced_protection.web_only.js_insert_all_pages_except / d11ff2dbf427 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except

<a id="canonical-ea8970d506665490feca7b579ee8844eb551b15fe1961f988ce2839d81cb0efa"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages with the exceptions.

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

<a id="canonical-5bf29b91d4ebcb53182f532ec1823cc8841150b08879c8e986066f082dd966fa"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except / d11ff2dbf427 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51): complete subsection reference.

<a id="canonical-c91ce61680031184737d72b015a2399c90a747765b054b5392b26aaa24d84f6a"></a>

<a id="canonical-6ade15bc32830ba9898754a5c03a9846eb1a82dd59548604f4d8ae28e0637f73"></a>

## javascript_location property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except / d11ff2dbf427 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b0d075649ecc85e04f40391efc59c3806b116cee4ddba433b8e1dacedebd6d1a"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except / d11ff2dbf427 / 5

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-427d24a7375da625908ab4ad1b1dda422932c5eb881acdfd020b11b50c98bb3b"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / b3be0a684bb6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list

<a id="canonical-1604a71fb2984feebc99ac58d97f7546a180bfc629491e4bf79c0257096b46e6"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8fba87c7083ca3c40854df42856a7ba9ba32d9fad0e2a267dd24768c509d9e37"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / b3be0a684bb6 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-e043eed10d1af7952070184d18eb79a1d76ce34a0457cecc5e2d6e9b2e45c9f3): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-bc9799d23be3a0fce5db9192b44654c8e537874903fc13a121eff7286305fe0a): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-ef57057ec07cd7f059b74700113b2c92d89a3bc46e0a9664a6232b57f0e81c43): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-013.md#canonical-dfe499c55707f8e8adc6a9e54c86083c9e529f001581f25f1516a911023c42d9): complete subsection reference.

<a id="canonical-3fc57dd0dffaa1a80353ab95c3c14083a3ce3c2e67102a8a04b86097e016c7fc"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / b3be0a684bb6 / 4

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-e043eed10d1af7952070184d18eb79a1d76ce34a0457cecc5e2d6e9b2e45c9f3)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-bc9799d23be3a0fce5db9192b44654c8e537874903fc13a121eff7286305fe0a)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-ef57057ec07cd7f059b74700113b2c92d89a3bc46e0a9664a6232b57f0e81c43)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path](data-sources--http_loadbalancer--reference--group-013.md#canonical-dfe499c55707f8e8adc6a9e54c86083c9e529f001581f25f1516a911023c42d9)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e043eed10d1af7952070184d18eb79a1d76ce34a0457cecc5e2d6e9b2e45c9f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4601054ea8a79086392d7af6d103e9be645c65dfed28878711308da59c92a680"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / bd0e6cb824dd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-aa3a47a3677447da95f4b5bd18de83f43600d3c35e38bbf39a8848832c7cb169"></a>

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

<a id="canonical-db21a5bb58bbddd9d0d208b9a1a3b907a377d804feaad97c2fbdd8babe85743a"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / bd0e6cb824dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16a98b490bf1f91acffc793e46d469e5c46651ccb18332199e5974813848a650"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / bd0e6cb824dd / 4

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bc9799d23be3a0fce5db9192b44654c8e537874903fc13a121eff7286305fe0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6e28d8a47b77bf7c0f946252502a9af3a08f89f682ddfa6250f6f9c613f529a"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-4f684e5db8996ed93d0ca2fd1f5e3c96e5a9d7ffb10d04da91e6c8348b84fb90"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-09cb8399a33e01fad8c281dc6686ac7bec3795bd846b1878543d5557ae533851"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 3

<a id="canonical-54944f30dac12ac0ad5886df8f5ab534a7584eb99694fe60657b5d0e5c65ee54"></a>

<a id="canonical-8680745c944c9610cdbaa7eac95ef67d03733b20f543113ff24018e3c928aef0"></a>

## exact_value property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-fdc8b0998e6fdc94dd04ed3ce3171b97ae231d3a23da6caf68e386d60f10a192"></a>

<a id="canonical-043e4c2bbb6dd36eefb367936b604a57e42b7c35303b1defe34b6ba1c56771bc"></a>

## regex_value property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f36da0367f610c7c3de37775aee89d9d36fed0746a3a5273de7aef62ff2b9450"></a>

<a id="canonical-cded770fddcac21c5e73a6e31e7a78e27ff2b40c985bec2b5c3885381863e766"></a>

## suffix_value property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-a1422178c0ad71ea6b64d4ba7955414dafadadc46f1017082a5867652cf2b893"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 11183caf722a / 7

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ef57057ec07cd7f059b74700113b2c92d89a3bc46e0a9664a6232b57f0e81c43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a61ef25f1201455c08571a97fea3cc3755a51a8deb3fbd21ce0410b0cb437577"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 6091751b52b6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-c3ffeb9a13976b696ccd90a51227da802655f1a77d0465aaf2c3bd88819be9d3"></a>

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

<a id="canonical-ef1ed3c31a1ba74ee2bcadf96006df6c5b2a7e853b8752cdfa26c8c5c43a9bb1"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 6091751b52b6 / 3

<a id="canonical-d8c4b8c4548b6b777447ea220ed9a61fefd96a9bc0bf568d9c79fea4d76b4019"></a>

<a id="canonical-488574c2bafae1e6333f93e6f17f71620e121aada45dd1ed6bed5d6bfdfb3645"></a>

## description_spec property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 6091751b52b6 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-40b689efdabfcdfa7c5e1eeb48a3be649a9470c249e22e01cb9be29442ea170e"></a>

<a id="canonical-d1d3a834e2216d8c29931e8feb188200a077248cbbc9ce0311c1e38a68d61182"></a>

## name property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 6091751b52b6 / 5

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

<a id="canonical-f682e579f91a4ac221e72ff3fb7932a582062ca397dfa4de209376e01faeb333"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 6091751b52b6 / 6

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dfe499c55707f8e8adc6a9e54c86083c9e529f001581f25f1516a911023c42d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701996cea1014ccbd054af4532cb3cb6cea800feb30d0717ac51e40d58df4afc"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-846d62b7811728e0316cd07250790cbaa5fd37ce2a5895a9b8050de096a9190f)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

<a id="canonical-636085768b13ed7198c4c08a24b46fec24c9044de733e75ea59848f3cc3279af"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-2dd702c061095fff4451426ce990c05356b6d1eaf6f980d893488f922e359591"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 3

<a id="canonical-a23cad303cd569ed06a6853d700e59cdf4bb86ab1ed3255d4964c02b1d95c540"></a>

<a id="canonical-350928bc059f1c9b57e624aa0710d977c791d03e5863ec2fc8836f3770f9e996"></a>

## path property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0c63ac0199e8f9ed1f78e5ee03f61f3f1f4556e84260f95cb69a28da3b0b1006"></a>

<a id="canonical-f162ddacf7c03a9498280b62677c1ba64957fbb03e40005d2d5d1e46347098f4"></a>

## prefix property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3084efe96d78a35058b28890567b4d4d9a04bee302dab6256651e515c9e780da"></a>

<a id="canonical-247d676e86734adfa8af6360e90d59a7c04d00989da6f999d101db4b920ff3b1"></a>

## regex property — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-54c72b4a571f738f5d5b59b105ed16b88740499c4ab1a388e2ea5be0419c34e7"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list / 243fb3876435 / 7

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-4b282ac46085d56a90a94ef3051f2d4a4c46242b0db0b9d82f852a3e642a8e51)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c20edbaaf3349ac5b7eacec65bcd9a92ae90ac4ace39dde841f70385f5924d60"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules — bot_defense_advanced_protection.web_only.js_insertion_rules / 8904af91eda3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- bot_defense_advanced_protection.web_only.js_insertion_rules

<a id="canonical-899e08b87c1695aec1de5cfc0b14199f1e040881913e900c997f775241609035"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

<a id="canonical-f14963b1e64eb3961d3d9031909b90cc9a1672e2bb26b874bb773fca96f9eea5"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules / 8904af91eda3 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a): complete subsection reference.

<a id="canonical-ccb1e5d937d50e0dd504dcb545064aff88353845a95cd55a1e5029a3e3c544aa"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules / 8904af91eda3 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3408a804cc728f72064b22ae5d199929c328f18816d911c59b142cc8f9312ab3"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list / 47743fad52cf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list

<a id="canonical-0f00f7012dc68a5f72e3f5209c254712f2a6a1181ebdd79cab406faa461016b9"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f506814be38cd40d8d737fedbf3d52702e0c988e83a8ba0ca2d6cee8d55b09cb"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list / 47743fad52cf / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-f07cff3da4322f06509c9f0a7e6f25c02ba59adb82e3342cb6a5105663003a87): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-a59ff5e758b2e9e9923f7fbbf71be6bd01c2f0b420e81da7892eba27c70b7b2e): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-c57784cbd5f73ac7ba8bfaf9a50e02a73624ac1f74fae76997559839a10a93dc): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-013.md#canonical-29e663fd585a1e65264205d51f5cdca5326de7f2f48710e27058112c482a2620): complete subsection reference.

<a id="canonical-ae337e0487a9214da4314634306348ee582fc5ffb5ca01bec8ae177048e980fb"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list / 47743fad52cf / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-f07cff3da4322f06509c9f0a7e6f25c02ba59adb82e3342cb6a5105663003a87)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-a59ff5e758b2e9e9923f7fbbf71be6bd01c2f0b420e81da7892eba27c70b7b2e)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-c57784cbd5f73ac7ba8bfaf9a50e02a73624ac1f74fae76997559839a10a93dc)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path](data-sources--http_loadbalancer--reference--group-013.md#canonical-29e663fd585a1e65264205d51f5cdca5326de7f2f48710e27058112c482a2620)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f07cff3da4322f06509c9f0a7e6f25c02ba59adb82e3342cb6a5105663003a87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3400d8b76d121a2353c303bd308354a58063b67b6ba419cf71e0c7ded4e2973d"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_dom / 60de6273a8bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain

<a id="canonical-696c36989a35baba592b318fc2071531794fb6cbb4d77e5050cc0f042b3185ba"></a>

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

<a id="canonical-d8f87b9d0d72a8d975e011541d1262205ec44b78c2634d57e2ac8adf6aedd3a2"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_dom / 60de6273a8bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f962c69217efe6ce5be66ff387f9622ee2fd69b5d4e779b067c33b5fe015588a"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_dom / 60de6273a8bb / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a59ff5e758b2e9e9923f7fbbf71be6bd01c2f0b420e81da7892eba27c70b7b2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-742041394206bd1b47d53d2547d06703a51c61dcb678c6a72aacd41d3d6e8fed"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain

<a id="canonical-610251241f27fe12a10e91fd52a9865a701ee2d9a18b737dcba74a61dd2bfcc2"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-992b578ecc1e5f9856e7e08bf501fdefb158d083c12c4a27bf0d86c79eb7a8b4"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 3

<a id="canonical-9057decd01cb4903c1411eb4cce306a9bdef8912c211bce5d915400a2e8bc47c"></a>

<a id="canonical-3f585db2e6d89017503956e672ffe7cace968fcf0ef566360a41bb6102246b04"></a>

## exact_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4bf7ee42b2bc7396fe4fe746bf55d2eebb288ffbea9ae8440435353cecf631c9"></a>

<a id="canonical-9244325531d9a07b5f3011d748eeb89e78f1f603b243bcf8a4692bd5846c092c"></a>

## regex_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-68cd39399384db2910c2b438754ab4c2ca3e9e39f35e81b3d58213c7da499dde"></a>

<a id="canonical-00f152d63b975e3970717b13609703c060a2505e9cbebe83b06908294a7e338a"></a>

## suffix_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-246e6380afe0df288339284eeb341a93bfc63fd3190bbc68c59542c9e3e753e6"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain / c981f88d7acd / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c57784cbd5f73ac7ba8bfaf9a50e02a73624ac1f74fae76997559839a10a93dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5469daf861497523bee4bff445c2910196dd82076ab25ccf6fef0e188292e976"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / b4880c12f544 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

<a id="canonical-41c796192fb7d7781970adfd86a33415c2a25c5fd38f34f52e4b3a93166e6a86"></a>

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

<a id="canonical-4bfa42273baf6e32cb361139698eb40ae7c0bfca6585d17c905a97fd5ca5b66a"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / b4880c12f544 / 3

<a id="canonical-548abd768dd24575d7e396d848ff6cebd0dfe6b3ea910e23b6aae5fa0e99c418"></a>

<a id="canonical-f282262bfb91c5e95a8ca71dba164f766567b4213941d97dc30a93e65956bb80"></a>

## description_spec property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / b4880c12f544 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3dd6210ddc626a281b422a921296fc74930d88e843936539f8c56420097b44a3"></a>

<a id="canonical-bc73309fbc9cfb046c254cb6d8e35c8b8f7de85989bb770693c9967932951824"></a>

## name property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / b4880c12f544 / 5

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

<a id="canonical-8cbe12ee8a52d8ab1d0519a8c44d0126fdfca5f352ca25a6af408a90f0b89558"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / b4880c12f544 / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-29e663fd585a1e65264205d51f5cdca5326de7f2f48710e27058112c482a2620"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2113ccc9997017915ff102e90c79a357e8e05bf686e8e7d22469bd3867bc7b3"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path

<a id="canonical-a07e2fefb504bae1d2ac20ecfc2fd7f2d35a15038993a096f07396a041084814"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-cfe4dc1c3f4a5f13314321e4bf624b48347122939627f374cc298332dfa1ad0a"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 3

<a id="canonical-415f5629614d56746a5182117b64f1e9cbe7743cafa869f4e23dbe0af9baff82"></a>

<a id="canonical-5e64ebadaa7055f8c32271a5c1eb1b348d6c594f26db4571f9e5faf17d9b7210"></a>

## path property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-aae83633c1160a6277777a7b4dca224fb52937bf284a17e8715e83dee043eac7"></a>

<a id="canonical-ee68024043511a2af0b8ba35c4b48b065bc6cbf4321b51bbd41dc16319381e7d"></a>

## prefix property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-f36e12e4444c23c6ff96e82627b1336c46287e6dd613a567327d5528fd0e4578"></a>

<a id="canonical-4a3f2f0ba30e45c48bacd6481b3ef19c48f1692925ff1593a08cab24c579921c"></a>

## regex property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9d8fccaacbe23a2d1dcee71e3c376f75068f1d6963951e13bc0d1db7d90d6d01"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / 244995ec0e80 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-04c82129ace698fa405ab6db052f1e82700420b232dd8dddd9c0334df8859716)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-157702b4b96582ef281e56098130d4911233b6190120a62e335c5635522f4189"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / 96530914e0a1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules

<a id="canonical-dd219e01a87186926003ba1188730a6c7e50045d60e412dd0de08ecaea18f069"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5a220241f9da6223a20a747c34cc6c0f5b962335a3e18b2025bdd159bc70226c"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / 96530914e0a1 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-bcdd7ed27c99006fd6f7cb99c04dea044f5b515c6b3876abca9fa7bb019b6d0b): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-cb120d44cd1fddc3bbfb9ab5a197aef61eb7d6bd92805df8179b6848a62d3fb8): complete subsection reference.

<a id="canonical-88918864b3f4fa334a2b4d708f3969fe8490b9fba412e34581ac272387f46e9a"></a>

<a id="canonical-c579b5437fde56237c4255d2100664446668707d7ae434db11cd21c62bab2133"></a>

## javascript_location property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / 96530914e0a1 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-1e45409bef17bd4230b78f73894f89b287fc441de2f0856cc88f8e463c9c3d1f): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-013.md#canonical-d33a4a62e34f0553068597dabe7ab1bf849b3919d1fd4a4ee7c1e27e2744c4a4): complete subsection reference.

<a id="canonical-5400ade5662c80ae0fbd7badd4684c33cd474856c7dddcd44fda4ec94a8d775f"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / 96530914e0a1 / 5

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-bcdd7ed27c99006fd6f7cb99c04dea044f5b515c6b3876abca9fa7bb019b6d0b)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain](data-sources--http_loadbalancer--reference--group-013.md#canonical-cb120d44cd1fddc3bbfb9ab5a197aef61eb7d6bd92805df8179b6848a62d3fb8)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-1e45409bef17bd4230b78f73894f89b287fc441de2f0856cc88f8e463c9c3d1f)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path](data-sources--http_loadbalancer--reference--group-013.md#canonical-d33a4a62e34f0553068597dabe7ab1bf849b3919d1fd4a4ee7c1e27e2744c4a4)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bcdd7ed27c99006fd6f7cb99c04dea044f5b515c6b3876abca9fa7bb019b6d0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8828ee251f97056af010460436f391ee8c3c31b4da2ce62bfa088ad8f30d6c45"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / a57f63379fbf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain

<a id="canonical-fd818d4acb57472479247c304838cd3852c5a6b5bb95023476fbc6a46c4ac5b5"></a>

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

<a id="canonical-a160edb41ce38a58255cb114f6a51c2d733bcf61ff72d36682430e795fcdaea1"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / a57f63379fbf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-956a5a79745cc161b5045463446e9c9017348984f25fa197211ab60ec953f17c"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / a57f63379fbf / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cb120d44cd1fddc3bbfb9ab5a197aef61eb7d6bd92805df8179b6848a62d3fb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43b56af03a6ec119ef7daea10e5bab6cb59e19cef27242244932e01670809194"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain

<a id="canonical-040058e9688d4c15c43d676757774a832ff0426182e0532ee3b315239441f35f"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-a6f2bbaff704c472ddf0447ea964767c1f154cf4d8c211331af7106b7d627091"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 3

<a id="canonical-46158071be9dcf414ec3779142d6339693dd7144656c432b2e464dc7fb5f1a6e"></a>

<a id="canonical-96bd60b97d62914611c3e1b60af013d162a44ba9646e966d700812e0140fc1d6"></a>

## exact_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-12e5ed597de565ce8fb39c56b2632feeddabe5015fb93a160f95f8772ceacd74"></a>

<a id="canonical-17da85d13d36004d793a58b0c78837a7128214b5571871cf68405d05eabae89b"></a>

## regex_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-09d7bb44bd300353a4a23b707bc9935e142834419a2e9fc553e20f247c159e87"></a>

<a id="canonical-b5e950ef89ab73b3fe3cbdc911607e82cdd93fbf3968f07cecfacd8be6b8c2ba"></a>

## suffix_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-6db196f43be8202e01b0a61faf1078b282b2753f41bac7a832a4b476a0b541ab"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / d554598ee989 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1e45409bef17bd4230b78f73894f89b287fc441de2f0856cc88f8e463c9c3d1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0095e9f660194808afe2cf6774649d6e174007d5430e353e6f283132bc69609c"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / c3a3aae5f9da / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="canonical-965ff788eefa9bea2a11bc6d3b197b58d357b9dc4e8deb115307a8f67d95ae87"></a>

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

<a id="canonical-b533f669fbd0162e962bb0c7022474d2fcbd58a827c7fcaf9fe7906a3b59b5c6"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / c3a3aae5f9da / 3

<a id="canonical-b1f29fb2affedefec275adbf0a6630c9f1942f989fb51f35a2825b9f850ab9a2"></a>

<a id="canonical-f2e60ecf9ed6d31600db1f29e604f4b5254b991d02cfada9022961980cff0267"></a>

## description_spec property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / c3a3aae5f9da / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-772eb606420651553bcf593f1ddcbce299f7c9f82f94573474afea377e99c213"></a>

<a id="canonical-ada0055567b077b46940c2c2f1fe3f7c1d74859f3d5548756e0406976d723044"></a>

## name property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / c3a3aae5f9da / 5

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

<a id="canonical-4715611492b2fa2e818261fb02e339ea43762aeb8bfbc06cb7db6fcc62220953"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / c3a3aae5f9da / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d33a4a62e34f0553068597dabe7ab1bf849b3919d1fd4a4ee7c1e27e2744c4a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f92673417b5f4dceefb97b6779ea7da8bfb730e6d016f923ba016f100aebb04"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-062b801f50315723e95e0cd01929f07c7c1f0090548f97bdd3957f600a318ab9)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path

<a id="canonical-853e1ee2198b447755ef00ed670d380bfda46852053b32100c2407ff556d7d6c"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-a9f9a162d0037e35b9eedb6c442468b37f27f1bed56074f28fb57e48f4da088d"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 3

<a id="canonical-5b6fa1bfad42d7c40bc679d32bbeb8d1f5cfcebdc448e885c62e2c60fe213066"></a>

<a id="canonical-7100d0bcaa8246e4ac4af77177534f4439466d799bc3eea3ea85a5922b72572e"></a>

## path property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3c17032c692afe39e47446f3aefac97b914bdb9c8a1900426293238fab890b54"></a>

<a id="canonical-81b1be556a31718dda381a80c99a96c8d57e0abe52efac392c745cd19ad61496"></a>

## prefix property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-8e0de8d624ca9e46f38b0b5a828bf0d110049ab3c91e4be75c30f058cd0c686e"></a>

<a id="canonical-1ee5d0cb50f1714791baa66b3ed4cc874a407c0cebd3c101497681328571eb72"></a>

## regex property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-857457267c5503393281d770843caab4d79fb9b1c445abf16e0f230bd0f5e9fa"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / a9e15437c11c / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-9921aab47064d1f2d9fcdf854c8195cab8ba0ad74e08bf7c45cca805e0931d9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fdbf1bb4c01d654be1d229ae0065c538d72a584824a8926c680fe48bb899031c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a1d59de79dfbf35ee7944ba935d8f0930f5db7a887e764eed092c35db71541d"></a>

## bot_defense_advanced_protection.web_only.web — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-f654039b97b9e852a5337602bc4f3beb85813ac25f63b4d3da203c820b8e9e9d)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- bot_defense_advanced_protection.web_only.web

<a id="canonical-9c6fb185b48ccd99b0c3e9581f0f1bfad053ceda22d104b2f9f374431984a2a1"></a>

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

<a id="canonical-7fbec419cb7ffeeeadb1a3198210ee9f6d923a2cfd6c189762f49a8ce4e55ac3"></a>

## Direct properties — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 3

<a id="canonical-8b28e532e64fb445eea8c8cb17ac3cdf9e41dc6e574667a8f06c08664a03b0d8"></a>

<a id="canonical-71e234f8d9a5b568d05ffd002e54b8ad6ae1cd5675a0662cf762d7c32040ba78"></a>

## name property — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 4

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

<a id="canonical-c8806e831cbc984204d33536101e60d39150b0aa2b24c7dda3d4dc099a52eecc"></a>

<a id="canonical-f5a823434c13e4749151b2a8d658081433417eca8611a92b9f843378ca554e48"></a>

## namespace property — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 5

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

<a id="canonical-140c1dc2928c7f1f836861f2ace0e43b567019ae6c5b25e0a324995b0b135ae0"></a>

<a id="canonical-3249aaad06e8062f84f28326928dc5df25756a254232af042567ab8550000be4"></a>

## tenant property — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 6

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

<a id="canonical-e0867ed418b87eb16b678fe5e447d742092c1e76b894842b4fac06dcbe71bf59"></a>

## Next pages — bot_defense_advanced_protection.web_only.web / dd4f2ac79f7e / 7

- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-de0cc0637252552be17522ae4281731c9fe06c40c73d8f2646853c19b563a483)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41430ddbdd3181b3f799faaf864ab63eddb7fb58e5e2fe6c621712a28668e42a"></a>

## caching_policy — caching_policy / 319278b6f999 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- caching_policy

<a id="canonical-ed3a44464e740d24cef1699fa02b8d9013c1529c04c3ebae3360133d106ca109"></a>

Type: `"single"`. Computed.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Upstream description:

Caching Policies for the CDN.

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

- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-ed3a44464e740d24cef1699fa02b8d9013c1529c04c3ebae3360133d106ca109)
- [disable_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-766f6acdde7728d639af6252950ccfa5735917daf33c46f87b55bc042df36218)

Select alternatives according to the provider validators above.

<a id="canonical-d655710d496bf0385d0109b3de43fdc52f9fd0ef3f0eb90288976d578c718db4"></a>

## Direct properties — caching_policy / 319278b6f999 / 3

- [custom_cache_rule](data-sources--http_loadbalancer--reference--group-013.md#canonical-93444bad1419495e3ae9adb9f094fb4b74436ce3b5d8bb53ae10e36136792128): complete subsection reference.

- [default_cache_action](data-sources--http_loadbalancer--reference--group-013.md#canonical-8c15fb5ca257bd9470ad8f7f207a29057c5624b4beced8877aa3e1edcd07783d): complete subsection reference.

<a id="canonical-2bbcd5ad9ad8762efd06301254247a4ac4b3d22de47a80e9cc0a7b40130b7aad"></a>

## Next pages — caching_policy / 319278b6f999 / 4

- [caching_policy.custom_cache_rule](data-sources--http_loadbalancer--reference--group-013.md#canonical-93444bad1419495e3ae9adb9f094fb4b74436ce3b5d8bb53ae10e36136792128)
- [caching_policy.default_cache_action](data-sources--http_loadbalancer--reference--group-013.md#canonical-8c15fb5ca257bd9470ad8f7f207a29057c5624b4beced8877aa3e1edcd07783d)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-93444bad1419495e3ae9adb9f094fb4b74436ce3b5d8bb53ae10e36136792128"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31c2521ae9cb4925c695cf3213d9196b9514cdf43698c8b348e66ab7fef8b1e0"></a>

## caching_policy.custom_cache_rule — caching_policy.custom_cache_rule / 6484d9213b40 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- caching_policy.custom_cache_rule

<a id="canonical-c537d5a223f38ada7e8654ff2ba5be1605d659263fa6f273c698647e2c038f7b"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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

<a id="canonical-0a97f438525d775e9d404ca3f51f35d11fec9e75c960bc5e5a38a92567ae90b7"></a>

## Direct properties — caching_policy.custom_cache_rule / 6484d9213b40 / 3

- [cdn_cache_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-972174598834ab0d3b221600a3fcfad13d48935cd814f1b1b9abf064ffb255c5): complete subsection reference.

<a id="canonical-52b8107de6421e5f36bc416d7ad3b7be8810834e7051b64ab23ae69093e14dd5"></a>

## Next pages — caching_policy.custom_cache_rule / 6484d9213b40 / 4

- [caching_policy.custom_cache_rule.cdn_cache_rules](data-sources--http_loadbalancer--reference--group-013.md#canonical-972174598834ab0d3b221600a3fcfad13d48935cd814f1b1b9abf064ffb255c5)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-972174598834ab0d3b221600a3fcfad13d48935cd814f1b1b9abf064ffb255c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a4be95258a1afb6b4c9c44cf1b1fc5156d2efd48bcd74cc25de19a16937571a"></a>

## caching_policy.custom_cache_rule.cdn_cache_rules — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- [caching_policy.custom_cache_rule](data-sources--http_loadbalancer--reference--group-013.md#canonical-93444bad1419495e3ae9adb9f094fb4b74436ce3b5d8bb53ae10e36136792128)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="canonical-c9396a8e498c65b0fa75ba6db4339a9e69530c3b23253ceb93fc52e1ee3c089e"></a>

Type: `"list"`. Computed.

Reference to CDN Cache Rule configuration object.

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

<a id="canonical-40139e12803010e932bfe69d600874ba4411c19c37d624cbb7a66799729475a2"></a>

## Direct properties — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 3

<a id="canonical-94e4354d7aa03ac5b504a06e2c62e83fa6ab2cb126f0f84c2e9ad2ba783254ef"></a>

<a id="canonical-96e071320dbfdf1033bcccb5b5ebb7534af431d66732eada1002ef5704a77b27"></a>

## name property — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 4

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

<a id="canonical-4c398435081a264335cb73f34dcc94c190ed9cc28ed5c4c683865a9895ceb6ee"></a>

<a id="canonical-64543bb9ab773d644e83ec6788a0d9987356351c0d66a88bf98b237d37d97fce"></a>

## namespace property — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 5

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

<a id="canonical-cd341034f0a10bad43ae66469876744293b0a65293a8a16ce58f6bc4b1c2b5a3"></a>

<a id="canonical-c0c82cd6064f0f12a1d56b5606203f999f08e8db241081e350859b24a9871553"></a>

## tenant property — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 6

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

<a id="canonical-5aff6e76bae5a06030cadde224e7dad4cec50a4c9e6a528c6c1684910701d13f"></a>

## Next pages — caching_policy.custom_cache_rule.cdn_cache_rules / 548820d7586f / 7

- [caching_policy.custom_cache_rule](data-sources--http_loadbalancer--reference--group-013.md#canonical-93444bad1419495e3ae9adb9f094fb4b74436ce3b5d8bb53ae10e36136792128)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8c15fb5ca257bd9470ad8f7f207a29057c5624b4beced8877aa3e1edcd07783d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c696f2d9dc0440d9e751e6cc48af3600caca5670d6d4027279f5bff47e8f010"></a>

## caching_policy.default_cache_action — caching_policy.default_cache_action / b81bc91ef895 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- caching_policy.default_cache_action

<a id="canonical-ae5683f7b7d90aba50dbc880319ffd1ea5e3a91d33d96e74cd29f5ba12f60185"></a>

Type: `"single"`. Computed.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

<a id="canonical-dfbc8337b60914e6d03e40eb8a1a8193a2fae17049f9bda437264ab0572f5bba"></a>

## Direct properties — caching_policy.default_cache_action / b81bc91ef895 / 3

- [cache_disabled](data-sources--http_loadbalancer--reference--group-013.md#canonical-aa61fa943f4a5157b7a797cdf618cd5aead06cce98045993d0689dc1af9a76f5): complete subsection reference.

<a id="canonical-7e1b8505e036ae1584a4d13ef6a0193a8d3dc4eac104e00ef6b9c6256c59b48d"></a>

<a id="canonical-7628d56dce71af2b6d5e408a91bde33ff76e125e008c254676fcdd7272fc4361"></a>

## cache_ttl_default property — caching_policy.default_cache_action / b81bc91ef895 / 4

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-d5a5929d858d2daf70402fdd8bcac09cc9e9d15edd8f94edbcdde680d217f059"></a>

<a id="canonical-f438c0293a98e819c7e8611774143f831b9bea60723d9ee91f9bebdf68ed898c"></a>

## cache_ttl_override property — caching_policy.default_cache_action / b81bc91ef895 / 5

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-68b11bdb7eacc8dab0abf501283519994ebf85acc690e36200f3cf56b8113dc1"></a>

## Next pages — caching_policy.default_cache_action / b81bc91ef895 / 6

- [caching_policy.default_cache_action.cache_disabled](data-sources--http_loadbalancer--reference--group-013.md#canonical-aa61fa943f4a5157b7a797cdf618cd5aead06cce98045993d0689dc1af9a76f5)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-aa61fa943f4a5157b7a797cdf618cd5aead06cce98045993d0689dc1af9a76f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b25d305a760a07c0940df1a03be2017517088405ff8567b8651e3d5e05ea968"></a>

## caching_policy.default_cache_action.cache_disabled — caching_policy.default_cache_action.cache_disabled / 3d463256cd11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [caching_policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-a6fb2379e952189b6bcc8f8566767a9972f615481793db4515c819834aa1c952)
- [caching_policy.default_cache_action](data-sources--http_loadbalancer--reference--group-013.md#canonical-8c15fb5ca257bd9470ad8f7f207a29057c5624b4beced8877aa3e1edcd07783d)
- caching_policy.default_cache_action.cache_disabled

<a id="canonical-c3cbc3e2a2f3e7dfa6b75b6f8751e2cf9f0a78dab1af8176de52cbac408f27f3"></a>

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

<a id="canonical-5cb5eef9409c7565664a21ba4e4aa0ad426a839bf15a185e1f54f1dd8d2b9503"></a>

## Direct properties — caching_policy.default_cache_action.cache_disabled / 3d463256cd11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-916980192780c86f09b30d6791b36d2d166d84606ccf3b221f56ad10eb5a4fcf"></a>

## Next pages — caching_policy.default_cache_action.cache_disabled / 3d463256cd11 / 4

- [caching_policy.default_cache_action](data-sources--http_loadbalancer--reference--group-013.md#canonical-8c15fb5ca257bd9470ad8f7f207a29057c5624b4beced8877aa3e1edcd07783d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bf8524ae85110951f9bb08b81c5e0e46ea159fc5e7d34366a39bcf3de81f7c7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f74a14a34a861ddd882590325e7e1d8fe187642219c6babe57a10bda552d31f"></a>

## captcha_challenge — captcha_challenge / f113df44f13f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- captcha_challenge

<a id="canonical-94afe776bd1cbb99d336b9f65fbfcdde67a3292e69b8e5f4e2a2c758edcd1fdb"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

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

- [captcha_challenge](data-sources--http_loadbalancer--reference--group-013.md#canonical-94afe776bd1cbb99d336b9f65fbfcdde67a3292e69b8e5f4e2a2c758edcd1fdb)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-ac78562305fc93dc44bf704750de89f8edddab648fba24d6d00a069000afc00d)
- [js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-7cdc723700c67b34975b8069d2794dc7cfe7f6a6e729f8b679f526683ef39683)
- [no_challenge](data-sources--http_loadbalancer--reference--group-020.md#canonical-19abcca17700aa9300fa125c09f5e481c4b08772c9bf94a63db0db330030b745)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-00114bfff9e8f35fff6fff86c0aeca8550664aeea9d1c4d88e22e58a24abd481)

Select alternatives according to the provider validators above.

<a id="canonical-128e99cdf4a25dd18269196d5b61d3a259a5f3a3cf29e4a48d0fb0e727cb9e42"></a>

## Direct properties — captcha_challenge / f113df44f13f / 3

<a id="canonical-f04ef329437db1985e3e8393787ca16c3e1c66cf02f9811b560413b78f7d58c5"></a>

<a id="canonical-ac4a6b17a2027f588c6868ea5544fd639947af24f3826664c6aeaf4620fb0d13"></a>

## cookie_expiry property — captcha_challenge / f113df44f13f / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-6991630070772dd2fa966dfc54046b2c5f6aa0c5050fa33b09aff8dfeb8fc5a2"></a>

<a id="canonical-758fc1b84e43e144b9a4e05ed58f9af0a4b54de5e2d41065cb29344b70f4b96a"></a>

## custom_page property — captcha_challenge / f113df44f13f / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-792e8df15aff3bc21fee9d0c67cc4ef858437f5b1d5a2d5298216dcbcd6809c6"></a>

## Next pages — captcha_challenge / f113df44f13f / 6

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7064579976ec01f14c56c84404fa017ed778e1777c74032a2fae564a84cffa31"></a>

## client_side_defense — client_side_defense / b3b521b23e38 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- client_side_defense

<a id="canonical-ccf73d4106fd06908dbbbf5f19da013f4da3fd105ff106cc23bf61b3bb160a07"></a>

Type: `"single"`. Computed.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-ccf73d4106fd06908dbbbf5f19da013f4da3fd105ff106cc23bf61b3bb160a07)
- [disable_client_side_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-f7af9d8ade6622f5a01b3a38d83cbba5133e427ad9964c3c53288dc93afc760b)

Select alternatives according to the provider validators above.

<a id="canonical-ad15d6083e35994d042371ee88a12e93eea0792a73504ad7e7e9d02a830121b6"></a>

## Direct properties — client_side_defense / b3b521b23e38 / 3

- [policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b): complete subsection reference.

<a id="canonical-01d70269b39f2314bde25ae3fdc3ff74c535db869333841cf6f26b6cb4a11041"></a>

## Next pages — client_side_defense / b3b521b23e38 / 4

- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17be0f36910d42a6550e6d0ba5444abeca1aa3adb6aad8fd9922377edaf6611c"></a>

## client_side_defense.policy — client_side_defense.policy / f2b860fa9fa3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- client_side_defense.policy

<a id="canonical-f313febdc7878d6806949ddbabecbab4bcc13de138b690d5f38e22a71d7e8965"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

<a id="canonical-6a70282fc27b18820fec1edb6c74620034ef83d7da514529af6f38be56ed9d9c"></a>

## Direct properties — client_side_defense.policy / f2b860fa9fa3 / 3

- [disable_js_insert](data-sources--http_loadbalancer--reference--group-013.md#canonical-5563d9dfd352cf9f3f892971f412e20dd3a89d50c4e0218e16710142b929abd6): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--reference--group-013.md#canonical-f2d06438cc2347ee2266603aae50be7efb0bf2f771867a3bca066947e107fd44): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e): complete subsection reference.

<a id="canonical-59844b435e17a881a5506d7a63e57347b03d4eb9127aaee1b5a27e1335e6efe3"></a>

## Next pages — client_side_defense.policy / f2b860fa9fa3 / 4

- [client_side_defense.policy.disable_js_insert](data-sources--http_loadbalancer--reference--group-013.md#canonical-5563d9dfd352cf9f3f892971f412e20dd3a89d50c4e0218e16710142b929abd6)
- [client_side_defense.policy.js_insert_all_pages](data-sources--http_loadbalancer--reference--group-013.md#canonical-f2d06438cc2347ee2266603aae50be7efb0bf2f771867a3bca066947e107fd44)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-cddcc8b626cac1ca512cd921feaf1053e00e0ff45310b85b8e9292b433ea905e)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5563d9dfd352cf9f3f892971f412e20dd3a89d50c4e0218e16710142b929abd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05cd13c1995931c35417e348942bf2c9d92828ce4399092d8d88ddd6313fd3ad"></a>

## client_side_defense.policy.disable_js_insert — client_side_defense.policy.disable_js_insert / b47d8ce082d0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- client_side_defense.policy.disable_js_insert

<a id="canonical-0c4fac07db224ac40b685c09e0d3d075b972252298b4de598d79e160b6ceaf25"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-5a6314597abc25c429ec935d1108690cb0bc9e07e44bfb0c482cb03588c41272"></a>

## Direct properties — client_side_defense.policy.disable_js_insert / b47d8ce082d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-364472b09d2417cbbf760f3328c7aab7d6c2e9c438e47f4b9c59aecc0a2d95f4"></a>

## Next pages — client_side_defense.policy.disable_js_insert / b47d8ce082d0 / 4

- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f2d06438cc2347ee2266603aae50be7efb0bf2f771867a3bca066947e107fd44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8339d6195927215be4d3cb34f8e71108c3baa84bbd3a84f6639dd5ad91ffefe2"></a>

## client_side_defense.policy.js_insert_all_pages — client_side_defense.policy.js_insert_all_pages / b4b9e2872c92 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-013.md#canonical-367f873504986262d380bdfbd173acad4137578e1de683d131b3d0248b345e6f)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-ec91062075eec6b0f8d484ccd90ab03d75b0fb24c8fcfd39e7978049df8af5f2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for js insert all pages.

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

<a id="canonical-3c155799080bf244c54cc155b5e1fde6998554d5e9dbd34df5f849adb33e7ff4"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages / b4b9e2872c92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2347850d32ac0b2bf164ad3d664f2b4564403804f81c89d1805ab7e628820e18"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages / b4b9e2872c92 / 4

- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-013.md#canonical-187bc8e8a39a49b1be57c0106b3ea4eb53acefd0a6616f1a35708cf01ec7e22b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fec9e6f94146815bcb028735c7065d5e00d72fa1a44bd123c68f5e2e592221b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
