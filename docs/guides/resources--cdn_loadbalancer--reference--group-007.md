---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-9dd1370f6ad42110a3557c536ed5368150c25a37a783d9d01c3a9ffbae479bd4"></a>

## Next pages — blocked_clients / d242ae1a4cdf / 10

- [blocked_clients.bot_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-a50a91f48f5009c57fe80afca029ee6a6f73abf255da1b37f018a311496e98e5)
- [blocked_clients.http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-017e08a116e7b1d2ab78e4d3f0bd5dc5fd9de77a221c8c74f81ddf180a3e9b8f)
- [blocked_clients.metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-a0800c1e7ec633a290825977d56239e8fc6884a8764a77c7e044aceb6015db37)
- [blocked_clients.skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-a2464a9d2e7a6d789619971e9f65e9165e3cd233adb79fdd2537c5a0a9726037)
- [blocked_clients.waf_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-7835e37f3afa0217718dfbbeeeeeb60e407d642aafeff7fe43d0ef55c61d8fdf)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a50a91f48f5009c57fe80afca029ee6a6f73abf255da1b37f018a311496e98e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0625c1cb9d36cd6e98d997916b55d5053c76c58400920a20d85b145551b60cb3"></a>

## blocked_clients.bot_skip_processing — blocked_clients.bot_skip_processing / 2d17ca707c9c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- blocked_clients.bot_skip_processing

<a id="canonical-f71055ec6a2d32cf9fdd6db27d2cbb82c9ff7d6d314917f4af2381bfcef690d1"></a>

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
bot_skip_processing = {}
```

<a id="canonical-990ffec36bd9992e42683c57a216ca9476789f4df1593efae4935207766c17ee"></a>

## Direct properties — blocked_clients.bot_skip_processing / 2d17ca707c9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86dec98f51d257bb17fc119912391010a19b6908b1aab00369f1eebc1fc172fa"></a>

## Next pages — blocked_clients.bot_skip_processing / 2d17ca707c9c / 4

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-017e08a116e7b1d2ab78e4d3f0bd5dc5fd9de77a221c8c74f81ddf180a3e9b8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2495acf64d10d6ff789c1e682eaabb568640c47460dfd5f08dccf5dc4364b2ed"></a>

## blocked_clients.http_header — blocked_clients.http_header / c3c53619efe2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- blocked_clients.http_header

<a id="canonical-d8b42ffcdc62de91006a486cd5a799abab100728d549df8d3e29f63608a8a271"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-568a663c112f30653dc89f781008f13892b99faf9018fd8ef4103da0144cf2a6"></a>

## Direct properties — blocked_clients.http_header / c3c53619efe2 / 3

- [headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-07d1d3e42cfa3040b7d5e4f9c3cbf9f2c40b94c933a9e107b5d05fb1426b7303): complete subsection reference.

<a id="canonical-0bb8b7488c6d22bd340b6993093edebbbc55cae2cd3d6817729371e924f6eb77"></a>

## Next pages — blocked_clients.http_header / c3c53619efe2 / 4

- [blocked_clients.http_header.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-07d1d3e42cfa3040b7d5e4f9c3cbf9f2c40b94c933a9e107b5d05fb1426b7303)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-07d1d3e42cfa3040b7d5e4f9c3cbf9f2c40b94c933a9e107b5d05fb1426b7303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c5d9d1d0d2be32fd4f539a3aaa8954428ce924cd4126946831cd6a93869ea1c"></a>

## blocked_clients.http_header.headers — blocked_clients.http_header.headers / 6e4832723974 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [blocked_clients.http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-017e08a116e7b1d2ab78e4d3f0bd5dc5fd9de77a221c8c74f81ddf180a3e9b8f)
- blocked_clients.http_header.headers

<a id="canonical-68b8d93cd929052092d7136dfc022606d1b172a6793ce763c86dce814a909cd4"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-dad94a96a497d81097d74cb597193321d2e9865827a912c43b45829b8aed8f56"></a>

## Direct properties — blocked_clients.http_header.headers / 6e4832723974 / 3

<a id="canonical-a99d4e1b9aeb34f80f470d2426f55b79263571771f2c5aa0181bc00679cde7cc"></a>

<a id="canonical-fb4df8c8252a3340d79c51f9c1d5131b9444b1b01f11a2d2fa820e01f4764eba"></a>

## exact property — blocked_clients.http_header.headers / 6e4832723974 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-cfef0ebd3ac0507f9cb14be41cca4eedc46fb3f9239101a8b34a0e2f4b21d677"></a>

<a id="canonical-4b0b850d6a083bce703a5b72c95378446255555a41956e15ed1b74cf65dc4a7b"></a>

## invert_match property — blocked_clients.http_header.headers / 6e4832723974 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-afa09e29e9e0d3f562f803f4d04c8a3c53a3c3efd19e439e433bceb231afd118"></a>

<a id="canonical-b426559ad5f313da577b12c2a543b3eb4c6b980047f6d1a15c970f4e92ca1960"></a>

## name property — blocked_clients.http_header.headers / 6e4832723974 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-e98536e17d698035548fd7b6885acceb5ae3a84737be25807f8e9cd755c867c0"></a>

<a id="canonical-7171204e41df7afd55052995766d3cf86b2acb72b1eb13e499bd62bf54258d7d"></a>

## presence property — blocked_clients.http_header.headers / 6e4832723974 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-054edc5ccf2bb2bc8c9acfc1ae99a48fef4a56fe72cabf8035b3233977e139be"></a>

<a id="canonical-e03eddfd76a146a2e6a3665c2f3c1200964fe57fd3bf72241c53691bd1c9c725"></a>

## regex property — blocked_clients.http_header.headers / 6e4832723974 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-035ad18d71ca4c63cc0b2a9402169ebf89f0855ba6e4cb977be22ee1c6f01400"></a>

## Next pages — blocked_clients.http_header.headers / 6e4832723974 / 9

- [blocked_clients.http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-017e08a116e7b1d2ab78e4d3f0bd5dc5fd9de77a221c8c74f81ddf180a3e9b8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a0800c1e7ec633a290825977d56239e8fc6884a8764a77c7e044aceb6015db37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b2fec969b597c592a7861dbcb15aed038f37baea13d4ecf0e4386ebea8d63c9"></a>

## blocked_clients.metadata — blocked_clients.metadata / 0a2af2f4fe59 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- blocked_clients.metadata

<a id="canonical-ef6dc7180029461e75fac3c0d5a2d7db9ae2672bb157150e807d02905653c7e3"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-6be204a1c9cdf0e3a0f04063d5448f83daac2902227797d7af949723529e65e0"></a>

## Direct properties — blocked_clients.metadata / 0a2af2f4fe59 / 3

<a id="canonical-278c4ddf4f10accddbcaa8ce5c1f4b7a07959cf3b4eeeb3a8bbac3c7bfe4bd8e"></a>

<a id="canonical-270bf365abed420323b53d5fafae49199e278e58100c334aeedfe9216e6de2ff"></a>

## description_spec property — blocked_clients.metadata / 0a2af2f4fe59 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-9ee174cadf7b2d95cdd0a3ea8b536de65ce5bd732039ed58eeda50fe6f3a1c9e"></a>

<a id="canonical-fff04beea1e5e4dfd678c24066c2cc8ccd72aa036cee4eb41f7e119cc08a39dc"></a>

## name property — blocked_clients.metadata / 0a2af2f4fe59 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-fcf07a4842d21822dee0124abca978b68721f2369194dce2651e26265b33fdcc"></a>

## Next pages — blocked_clients.metadata / 0a2af2f4fe59 / 6

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a2464a9d2e7a6d789619971e9f65e9165e3cd233adb79fdd2537c5a0a9726037"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28b6c7ce1ee3f3d51a7555a2779e5da8172eeb4c0c249f8a7503fd935229ae95"></a>

## blocked_clients.skip_processing — blocked_clients.skip_processing / 5d14b0d337ad / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- blocked_clients.skip_processing

<a id="canonical-f2801f9d50116ad29abf306aa335e9f8970da3363453e486f947b8155e8f6d33"></a>

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
skip_processing = {}
```

<a id="canonical-72885c8ec9bb9e855e29e0fdf0cf8b8c6b9adde1d3fb624d73616df43ae35bd9"></a>

## Direct properties — blocked_clients.skip_processing / 5d14b0d337ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ef6930fe47477c19765f75ff68f2b02bd40db156f0698334c523e85d022d8f0"></a>

## Next pages — blocked_clients.skip_processing / 5d14b0d337ad / 4

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7835e37f3afa0217718dfbbeeeeeb60e407d642aafeff7fe43d0ef55c61d8fdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-059a03e0ee4e237fd519847190c5a5e69ee25a97a8d865c96418fc37d44ed334"></a>

## blocked_clients.waf_skip_processing — blocked_clients.waf_skip_processing / 4130deb1f0de / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- blocked_clients.waf_skip_processing

<a id="canonical-9b9a7d383f75746606b882b98d14be3202a20f99eb3f4aaf7818b384f0eca5c7"></a>

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
waf_skip_processing = {}
```

<a id="canonical-cc90003d3f6b1ba7e0c95b1b0921364722d13cff82eb82af9052d1c27f85ec83"></a>

## Direct properties — blocked_clients.waf_skip_processing / 4130deb1f0de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2bef791a94fa648c5c98427ca1b8550d2931cda6517e9b9a77259c16089f1e3"></a>

## Next pages — blocked_clients.waf_skip_processing / 4130deb1f0de / 4

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83806e1c291c5054ae8d2220bfcd963d401f8f7da032c154f1227f49aa11095c"></a>

## bot_defense — bot_defense / c36a1d658ec7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- bot_defense

<a id="canonical-80d5da23471d6d5a4df84470ea5f938d9637cb248fe75723fa650844e200d023"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_cors_support",
    "enable_cors_support")}
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
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf208711d9c83dac87c5093e8049b8df3784981d9cb5d7022d616a56978d9cd9"></a>

## Direct properties — bot_defense / c36a1d658ec7 / 3

- [disable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-94ee8a8a60207dec862437a362f5085256a940faa6fb59648c3f71db70083905): complete subsection reference.

- [enable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-5c2ead318161388ec3e65edb5b5f74314ddcaf13b1bd2f56742119b59afa8c53): complete subsection reference.

- [policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5): complete subsection reference.

<a id="canonical-82b8821e3b3291eb221152fa98f751c1970e62f19afd445958805783d948e090"></a>

<a id="canonical-dcb52454ac9132c57c7989a3ad877fa9f675e5390fd92c1f70417dfb65414961"></a>

## regional_endpoint property — bot_defense / c36a1d658ec7 / 4

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AUTO",
    "US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-80459b0d7e0cfdacfb4d1e917041c606a93a82cf33d8abebf2f89c65bf967844"></a>

<a id="canonical-6c1b7a9645ea1827030383a77037551d740a8ca3556c9d77f1fcab2a36be5771"></a>

## timeout property — bot_defense / c36a1d658ec7 / 5

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-9f24a18861516e8cb3f48d75495e064b7725b4cee08007cc607015edeb20bcd7"></a>

## Next pages — bot_defense / c36a1d658ec7 / 6

- [bot_defense.disable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-94ee8a8a60207dec862437a362f5085256a940faa6fb59648c3f71db70083905)
- [bot_defense.enable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-5c2ead318161388ec3e65edb5b5f74314ddcaf13b1bd2f56742119b59afa8c53)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-94ee8a8a60207dec862437a362f5085256a940faa6fb59648c3f71db70083905"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e89240a576a544c4914c6cd61712ae8c53a66d96a095e0ade46bd9b485c62bb9"></a>

## bot_defense.disable_cors_support — bot_defense.disable_cors_support / 3aed5c136502 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- bot_defense.disable_cors_support

<a id="canonical-405fe6ac29c84b004c068c8066be95a284453573d93340b2395b42fe2ceee36d"></a>

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
disable_cors_support = {}
```

<a id="canonical-a000e2fd250b9a71dde21b86d8d46200b6442045de55f13c12df78205b4ac904"></a>

## Direct properties — bot_defense.disable_cors_support / 3aed5c136502 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d832bb54c97ddd7558a732b8ea09e3b53b3db97fb18815f3fa209e11fdd471ad"></a>

## Next pages — bot_defense.disable_cors_support / 3aed5c136502 / 4

- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5c2ead318161388ec3e65edb5b5f74314ddcaf13b1bd2f56742119b59afa8c53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0dc2c4d6f889931bc6ff8f08e3e4919fbcd37c2bb5b0ccbbbcbde803b5806a5"></a>

## bot_defense.enable_cors_support — bot_defense.enable_cors_support / 0878f01a9139 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- bot_defense.enable_cors_support

<a id="canonical-2b06aae70683cbb6a53f10eabfccb2bbcef1b8dc9d7a6a7cc793b76d0f2f13ee"></a>

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
enable_cors_support = {}
```

<a id="canonical-c6492943fcc160af57ad256d6e657bc15619776cd4525205e6dd4287236edb9d"></a>

## Direct properties — bot_defense.enable_cors_support / 0878f01a9139 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38cb588fe395bc1c5827da1ed931e04ce6a9334b9707cd1222d656883bf4208b"></a>

## Next pages — bot_defense.enable_cors_support / 0878f01a9139 / 4

- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e827e4cb78dbb566394cf28331388f6fe92d3bf74cf0a605c2ca2ad7a27c9de"></a>

## bot_defense.policy — bot_defense.policy / 4c2b1dd23709 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- bot_defense.policy

<a id="canonical-9161384f78d2d1b845d28fc67c80754a56d78dbe88ad2f4cb7b0f5a310038064"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_app_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a83da7808fab318648dd4b413ef491ac359acc9e5403d2b2e226ff0eac54358"></a>

## Direct properties — bot_defense.policy / 4c2b1dd23709 / 3

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-007.md#canonical-d739f66381ba43d2e36920608e1ae8a66cae979f938ab47a0742745bd985c4e0): complete subsection reference.

- [disable_mobile_sdk](resources--cdn_loadbalancer--reference--group-007.md#canonical-18e8ac26342b4983cedc899e497d4db7cd492459cd116d971f6d31a32dd00fa7): complete subsection reference.

<a id="canonical-065ef172c2ab6433f7a24ceed98ca128ccd461465b2dcc714848591f3cb0b636"></a>

<a id="canonical-b162809c1f8654b63500701fc4867da995597cecc8b426a434fa6c2f498140bc"></a>

## javascript_mode property — bot_defense.policy / 4c2b1dd23709 / 4

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-97b851b00e37cefa7e34ede7fe2a9a9b2cbaffa27193f72834c7e2facbe29236"></a>

<a id="canonical-7edf43db6d3bf575d7a291328faf16b0f9375e7a3d1d5c9f6278a83487920c7c"></a>

## js_download_path property — bot_defense.policy / 4c2b1dd23709 / 5

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-007.md#canonical-256e64cdbef93f0087c399187cf9bb2fccebaa1c9b673a6b6588bce417e461b2): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39): complete subsection reference.

- [mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a): complete subsection reference.

- [protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447): complete subsection reference.

<a id="canonical-bf27ce49ea3251864c9f9ecda366c455af8ab988573d2ad1cc60f83ac7528a0a"></a>

## Next pages — bot_defense.policy / 4c2b1dd23709 / 6

- [bot_defense.policy.disable_js_insert](resources--cdn_loadbalancer--reference--group-007.md#canonical-d739f66381ba43d2e36920608e1ae8a66cae979f938ab47a0742745bd985c4e0)
- [bot_defense.policy.disable_mobile_sdk](resources--cdn_loadbalancer--reference--group-007.md#canonical-18e8ac26342b4983cedc899e497d4db7cd492459cd116d971f6d31a32dd00fa7)
- [bot_defense.policy.js_insert_all_pages](resources--cdn_loadbalancer--reference--group-007.md#canonical-256e64cdbef93f0087c399187cf9bb2fccebaa1c9b673a6b6588bce417e461b2)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d739f66381ba43d2e36920608e1ae8a66cae979f938ab47a0742745bd985c4e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f250548c6ff15e7386ce276232716158eb979fb407135114b489aae65726395a"></a>

## bot_defense.policy.disable_js_insert — bot_defense.policy.disable_js_insert / 33ca9031e542 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.disable_js_insert

<a id="canonical-8d4480f5b7cce6c35a12669a851a2978740f0809ddffe857cc42f0e5aaf9d901"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_js_insert = {}
```

<a id="canonical-70366a70c21d2e92ae99db287e348d3e49bb145d5cfd8866820d1c0ccd61630f"></a>

## Direct properties — bot_defense.policy.disable_js_insert / 33ca9031e542 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9aee3a5d7aa9c3ccc7eba110c67a10bef4cca0cd78b9c6f98fbcdbf62ae25ac0"></a>

## Next pages — bot_defense.policy.disable_js_insert / 33ca9031e542 / 4

- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-18e8ac26342b4983cedc899e497d4db7cd492459cd116d971f6d31a32dd00fa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4ff7e4d8a23b7a91723c9eb3d38824184b5d6bb219975b1ed292065b082927"></a>

## bot_defense.policy.disable_mobile_sdk — bot_defense.policy.disable_mobile_sdk / 3a12a3445d2b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-ab19b85ba50f623e3933611c4c2d2b8011f7d3f33af0849d2bece52328647e79"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-36b9c34908ba345a409ca5ac6805b319ad2e515370cac4506bc58f078e60b104"></a>

## Direct properties — bot_defense.policy.disable_mobile_sdk / 3a12a3445d2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-596157e1ad2676c29a84f368dd4f4b911beaeba58894291aed798b6b20526697"></a>

## Next pages — bot_defense.policy.disable_mobile_sdk / 3a12a3445d2b / 4

- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-256e64cdbef93f0087c399187cf9bb2fccebaa1c9b673a6b6588bce417e461b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff99ec1a7c286f22225b7db09d490dee82b054dfdfb15e7ec73b217f377e88fa"></a>

## bot_defense.policy.js_insert_all_pages — bot_defense.policy.js_insert_all_pages / 048759813f71 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-9250befc839799031e29d19e0542296dc82511090543bd380d81b6fbe2c4a0f6"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ce5ad108055eb19799e22b41d02be5b88f627a996f270a6c3f741ab7b3d639c"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages / 048759813f71 / 3

<a id="canonical-e6260c9290fe89164f549d6e07a1b531e81895d9b4cbde06641816dd68f3c197"></a>

<a id="canonical-e3dc8cceb81e8d606c29051e15bad45e6a0571015028fcc5f19062cf97dbca95"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages / 048759813f71 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

<a id="canonical-6ca9e97f607c80ba72f2e0a88400bc47bc43cf4864386c0ee5f88b43cf45a940"></a>

## Next pages — bot_defense.policy.js_insert_all_pages / 048759813f71 / 5

- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84e6c25f862d2e9ed4512aa87024e42c7ac96556070b28f3e422e7295ae706d5"></a>

## bot_defense.policy.js_insert_all_pages_except — bot_defense.policy.js_insert_all_pages_except / 17c76afe8d1d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-625dec3d8bcc3d1dcf96e434b53d76bc5fecc5a0d1c1198ff2c4b1953554d837"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-7d827a9f46675bf1aba0d53f6985882d4deb84282f06a6a0a20ef53143ae32a6"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except / 17c76afe8d1d / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d): complete subsection reference.

<a id="canonical-4cc5a4408df22cca5845407c6258d382381ebc422901c629be4b367915a431fe"></a>

<a id="canonical-72985942d903a7895b41e6b0cb75c4237b3b7e0ea4851ff9a6cd91331407b6af"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages_except / 17c76afe8d1d / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

<a id="canonical-8753a964787114094b8ff16c77b86369b7dad8a1b07a53bf647c039a5bd6aac1"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except / 17c76afe8d1d / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea06b58768e72153ea43717fd44bbed424c571e1f2a266226854184ea06cd377"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — bot_defense.policy.js_insert_all_pages_except.exclude_list / dfdba6ca35a7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-5986094cb2b1580f94d72d98f7dc6d71dcf78d8a11cd7abf3a396617b6eb9a36"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a1712d4aeca94d74c94f69ef14d66f7698d2838dc6a602905d27a83d43781098"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list / dfdba6ca35a7 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-d3c376c0c116dbe78a4b911c8b8cad42ecdc4deda567e0b2d0f291825d4ea833): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-446d04910e216f4ff721a1ed06133198cd3e5aed3b69461669fd482baead96aa): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-346280031e6d39b8ff124b2e71bbd56f7cf4f7a41e5832d67d8aff5ecbdd7b1c): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-935912c660343d0a77903d5426d8729816213f4e733214e311e192cb6c9812bf): complete subsection reference.

<a id="canonical-587f43a2d2cf8ed12729eb39ffaee25613350d3d598ca1672ae7d55d1236ab47"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list / dfdba6ca35a7 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-d3c376c0c116dbe78a4b911c8b8cad42ecdc4deda567e0b2d0f291825d4ea833)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-446d04910e216f4ff721a1ed06133198cd3e5aed3b69461669fd482baead96aa)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-346280031e6d39b8ff124b2e71bbd56f7cf4f7a41e5832d67d8aff5ecbdd7b1c)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--cdn_loadbalancer--reference--group-007.md#canonical-935912c660343d0a77903d5426d8729816213f4e733214e311e192cb6c9812bf)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d3c376c0c116dbe78a4b911c8b8cad42ecdc4deda567e0b2d0f291825d4ea833"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ce597a38f9aaa9c7f44c74f6c27aa056b575f80119c1bf5aac3a15ad18e63b"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7911211e30f9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-fc072758ca7ceeccd99ffee3473036ee02c6a0350a6d067b7752f8705d4f9154"></a>

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
any_domain = {}
```

<a id="canonical-c0b90232cc5f78552c5c53a04bb94552a8cc45f3a063441416e589ba2cc432b5"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7911211e30f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1727ab45dc81a7ed535480b8551b5234e75fb527d56da0ee9e81faa4def6e3dd"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7911211e30f9 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-446d04910e216f4ff721a1ed06133198cd3e5aed3b69461669fd482baead96aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-182cc89e408f04e6fa906e48fd17e6a18d76e5d4284c03732cb9bb9e881a794b"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-98843e465b928a4b3d2b8ee430596a5b45360764f0d67418f72ee266237d90ee"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-fdbc8529f2191d8831d414e6d4ca198910365ff3211e3ee33f3dcab96fe92866"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 3

<a id="canonical-f1905198a6048c34af2b941911934e4e797aeb2d859259f011125935157cfb57"></a>

<a id="canonical-ead5d5e659c1812c7d0b756d7bc6a2f2076c0e8c31e3e094de894b8dc6758944"></a>

## exact_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-f2476b0ddaa0c96238f6106d5771199fe60c1733e4bac056be12fb6c3c805fb4"></a>

<a id="canonical-462ffb6282054830e8b4b5ebf4222611845cd3f61008fa0fff4cbe3d07cb6458"></a>

## regex_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-af25ee1aea9724999fcea742c6c0332107d4f4eeac6b3882718361ec6268df6b"></a>

<a id="canonical-9b8ec0f7add01c0b45fe0a0b6307a891ed2a6c0dd820627623722dad92e4b339"></a>

## suffix_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-e5cb0a4247160167caee7892390d91e0d4d904b8fe324f7b94b70ea5a6471c1f"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 4231b4dd637e / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-346280031e6d39b8ff124b2e71bbd56f7cf4f7a41e5832d67d8aff5ecbdd7b1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d8ebb5b0aef8a7eb003cde87637f3a96c42c27c8c6fce19b3980df7ba8556a"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 0737a9bfb218 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-111ea0455631630ddea6b683a8a5c75805264102cb238b5491052f825bcf3258"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-347eab205b9e210ce766c9fe296202d2c55ee8321c8ce097d555fb8ae3bb2f0b"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 0737a9bfb218 / 3

<a id="canonical-b6f79fa45b91d4c06d54a0ee17b6ee9eac1e82b909b1cc0f9fccbcc4b3000692"></a>

<a id="canonical-4813d2daebda182138a015599e6856de10427df078c88708ab355e5f16700ec5"></a>

## description_spec property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 0737a9bfb218 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-d710cf73c9915590623f9150c3e9909dd5a79df3b7a0ef7810a2b44e975ce55c"></a>

<a id="canonical-3bafb7a30a0b9bfcbcd9dcc5cbaf33b110f5a6a838b410e05555f654dbcca543"></a>

## name property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 0737a9bfb218 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2383bbb1c6cde2d30dbf51122097c863cac2fb3aae5ea79395c142697569e9ef"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 0737a9bfb218 / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-935912c660343d0a77903d5426d8729816213f4e733214e311e192cb6c9812bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375891db9389c031b8734187d054bf7d4142c7d6b4a6499314868ba8a9d8f218"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-007.md#canonical-8f6c68b96547529761095395e4284989f9d09877e224f20c5126ac0a547569d5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-10be119035013d098ff1d61a83d9c7078e722daf52f275db733cc7a9f4317e47"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-d76144327f75fbfbaaf67a3ac2b9ee20fb5699b917257f6e8fa6389321553cf1"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 3

<a id="canonical-3e2b2379fd905fe64866fefff673932801f4fc9ed261772328a3cdb41613aadb"></a>

<a id="canonical-0cd97e07ac23dcd0fa8de66f73956cb4bcea29f9d32a744493f1828100b1a0c9"></a>

## path property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-67c686699af45e8b596654fbadfe3a805d9c3ca2f9e72e96e17e467aaaee32fe"></a>

<a id="canonical-d608a340b4580d4ef36e1e37e49d73c0933c0676f3b083b5e58607f5e4df7073"></a>

## prefix property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-37ab5ce6068dc8f882cfa621c8d6e92682fb40c2b027c094c663d11501ae57a5"></a>

<a id="canonical-49f085f2877959bfbc15853669f014fd81cb97ab4996eb8b90b36955ec7f5f62"></a>

## regex property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-31a5ee4fb0e8a518114b12f4d6038d2d89fd7a1bba969eb3e18b445575056b11"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 02db0649e156 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-f9b3f2e02b7ea491eb7367777c8b2a8cdf5870110ba9fb15f91eae4dd529139d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b372e8f0e57795abd295190ac5932451a979c649ab88e1f39e40d417989d5610"></a>

## bot_defense.policy.js_insertion_rules — bot_defense.policy.js_insertion_rules / 7385c6707444 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.js_insertion_rules

<a id="canonical-38fdfef74b1e98f8374c0d2b99ddee9a04234f4c36ed6606dc9d5fec24e61c72"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-af3f3308570318f4acdc5377df8b45009cca84aaa4afad5cf5bd691896982654"></a>

## Direct properties — bot_defense.policy.js_insertion_rules / 7385c6707444 / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144): complete subsection reference.

- [rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463): complete subsection reference.

<a id="canonical-9ec104d61e26641945d9f377008fd0d21110b0e700dd20368ba233eeadec787c"></a>

## Next pages — bot_defense.policy.js_insertion_rules / 7385c6707444 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b946e7c140807a8a0db137d6878ba599f0393c6d39668adb1cf3523f2e9f49c0"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — bot_defense.policy.js_insertion_rules.exclude_list / 1f9c025aa089 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-2d9a2ce8afe9afe1c7c6c572fac3a94e50b2598ea23b84beeb72e79dd2517f0a"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-f908606e98dfe52a7a3954ea8217d34adbb988349686576783a580e6ee6feea5"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list / 1f9c025aa089 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0fdd944bf35108d2996c63930562c7d09c334b078930c48d5751619cd96ec9a8): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-1b3f9e6e429a62d420ae6a64cbef1b59ac41f2231dae4d01ebfbfba871547350): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-7843bcacb2af1c0c8ccc84916b46e3d850088485d391118063d347475a6c768f): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-f63b9a4fd53bf4e3df7788c5fe32b7de1a985adea703300f036c9684e7cc080c): complete subsection reference.

<a id="canonical-e8d37d777c6ff962e17086f1891c1be2a29ac282b5361547c0434cb6f95fecaf"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list / 1f9c025aa089 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0fdd944bf35108d2996c63930562c7d09c334b078930c48d5751619cd96ec9a8)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-1b3f9e6e429a62d420ae6a64cbef1b59ac41f2231dae4d01ebfbfba871547350)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-7843bcacb2af1c0c8ccc84916b46e3d850088485d391118063d347475a6c768f)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](resources--cdn_loadbalancer--reference--group-007.md#canonical-f63b9a4fd53bf4e3df7788c5fe32b7de1a985adea703300f036c9684e7cc080c)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0fdd944bf35108d2996c63930562c7d09c334b078930c48d5751619cd96ec9a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68f0833992798acb543e008617bb382bfdbc422574b470219d32b2b18d5df3c0"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 82607cf10ad9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-32ff84183f2108b2713b3c37fc9f2823090cc6df3557eb7d8b3ae44fda74682e"></a>

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
any_domain = {}
```

<a id="canonical-7d1f5360a0a853219e8ca2d8a3a261714bda8309afc0ec59fb69ada7fcaa7cf2"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 82607cf10ad9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3512dfd7111a323a8e38500413a8670dd39f0f46cfe0bc8c596f675f1a469770"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 82607cf10ad9 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1b3f9e6e429a62d420ae6a64cbef1b59ac41f2231dae4d01ebfbfba871547350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba36756d88a3257279db954494c863fe802e0e85225a51071206c0f8c74c5efc"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-4a9550b44f26a94c61cd57e1b5a97276cb00ee16a0fa35a44e643a19b82c263a"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2558b8e92904c9bda5c7c95d2f3a071baaaa12f8ee4f07ae719cc6d7614c3c9e"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 3

<a id="canonical-e9b873ce28d619bb45a40f7e01406f0b1298abc056b1000e40e3663c8077e397"></a>

<a id="canonical-18c2888af8d14495989c56aa813be42093aa72dda4c95649c6f395bf54736cd6"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-60d9bab04cfb329dd2c7482f98f3bd44044654412137bf02c7ba06ac7bb73cb6"></a>

<a id="canonical-a86a2d6f374642f1d264e7630c85d28bb87342cd4a2938cc127ee7aeb82e0af0"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-67d16337e0a3577420bac1812859d6bf7ef8ba5476f4ce1fcf2146b59b7fa6ef"></a>

<a id="canonical-0086fb9e8638dab3b5ba8cbd2fe0baca0eb3e91e5e622d840435e3d14571d675"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-e1eb5aa04fcaaff3961098c1b347a5b5c55b0438a5d8020dece38df3249aabfa"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.domain / b1c2a4fefa91 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7843bcacb2af1c0c8ccc84916b46e3d850088485d391118063d347475a6c768f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0225dfe0df58d9d384cbde0ae70f4853df9a978c98ee34710bb8ad09c823f1a"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 701b511a74fe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-3d7018c446cbb2a495ee71cffe850cd079c0e1dc71c8fb049f01cc68bf99c11c"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-f65b7969546b3d4cc75f07e8bb6bc28ac6825fbf384bbe203db5d1842c8f5651"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 701b511a74fe / 3

<a id="canonical-116326de4bb0308a174fb383b506c69b0282e6c64e351073e3ef7843cf80f051"></a>

<a id="canonical-acd13acbb11e4cc59914ccda71bfde518e41f8eb5ebfd6887bc051281013ff15"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 701b511a74fe / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-6923f1e73902dfd6ab520ff11f6e9b8302881fbc95e34a146837181d6f7627ca"></a>

<a id="canonical-b58e44bbb64b62574b0c381bf817adc792598da65a03a9440ef50efd51c0b1bf"></a>

## name property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 701b511a74fe / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-89b8fd53dc74fdafbb55aba8c0845bdf2024ba3f0572f5a66c4d5b3f5780d24e"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 701b511a74fe / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f63b9a4fd53bf4e3df7788c5fe32b7de1a985adea703300f036c9684e7cc080c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13bc9de627c92386eb06b0823984ccbc141b3112387750fa2c14ab93fa986448"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-90397bdfaaca4f761a22c68e1aea38b78bbfe8ac842993743b743951dc915311"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-76668228efc82b55ab4e24fd5b73c01f4ab8e8758d9789146fb52554ef272919"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 3

<a id="canonical-5a30f8e76d7f51dbcb9c491f1fa5f50d7e88d4d46e455720f0f43a55ceac4731"></a>

<a id="canonical-e00de8a6c28f8292cab97eea3a5ae3073616c4db9a676566effbdbd4471630b7"></a>

## path property — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-5e2dcdbc5ce39350896e995f780c8e733782244a276d9836da90db2a2e71d895"></a>

<a id="canonical-aabdcdd5c7cba53017cd327d70203562d8ef6343297fa62ad33ca0bd1d5feab8"></a>

## prefix property — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-1f6c9eb8d218027b72f9dea73ad34829baca0db132661518d113f3f969baa6e7"></a>

<a id="canonical-ed9426d1168c1dd56faeb5e3d3396991be875e9b9d78e2d4c1c2f72a2d93f77d"></a>

## regex property — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-86f33e789598ed9e1b3fa00364a3e319d1d303aad99563bf2af1f5e7120bda14"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.path / cb6eaacdfc39 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-2ed9f0f2ed6e93c6a8eae08414837b5091400eb7e11dc7cdbf4b8cc498c42144)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8fd512949a2e6e6790829a2174b803d3f3df1682c6e7f43d206746c81864d7e"></a>

## bot_defense.policy.js_insertion_rules.rules — bot_defense.policy.js_insertion_rules.rules / 72fd2138eef7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-2a1c42f6ced782d2059002af425137c30509ac3adcc9306b62d33f63eb7e574e"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c042f1de29da7b124ae0d9f74da6e9c243635a6b3784d4ce734a5347120658d"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules / 72fd2138eef7 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-92a67a61b013b1c6d73e62e3a7db3592edea6ef14730714ec2c9cc9726b9422a): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-00e2f5a19f4ebb77ede3096a873e948ff0ab32c874f5af6d2026e2982ae14a70): complete subsection reference.

<a id="canonical-c6eb634e25e797e1b5a5135e728f13a2e18ca4de7302badbd2ce79715d12bab5"></a>

<a id="canonical-d1fa36c36e4469cf41b2f8d0800b89d122b5f61d4cae9154bf4e8efad71b4b66"></a>

## javascript_location property — bot_defense.policy.js_insertion_rules.rules / 72fd2138eef7 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-2df5ce039ee61d3853ff29f465b368ec2b4298d6acf90bd5738dbe7230399959): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-007.md#canonical-7c39d581dce7bfdf400fe0f7a1bab283efc2f3d864a8e45c54897d43a22ffc82): complete subsection reference.

<a id="canonical-6da23c3ed3aa2ea6dfd959d5af6f4c62418ead5813e01239760d54419a6c4e7b"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules / 72fd2138eef7 / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-92a67a61b013b1c6d73e62e3a7db3592edea6ef14730714ec2c9cc9726b9422a)
- [bot_defense.policy.js_insertion_rules.rules.domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-00e2f5a19f4ebb77ede3096a873e948ff0ab32c874f5af6d2026e2982ae14a70)
- [bot_defense.policy.js_insertion_rules.rules.metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-2df5ce039ee61d3853ff29f465b368ec2b4298d6acf90bd5738dbe7230399959)
- [bot_defense.policy.js_insertion_rules.rules.path](resources--cdn_loadbalancer--reference--group-007.md#canonical-7c39d581dce7bfdf400fe0f7a1bab283efc2f3d864a8e45c54897d43a22ffc82)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-92a67a61b013b1c6d73e62e3a7db3592edea6ef14730714ec2c9cc9726b9422a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6804faedd37ec6ac0864bbccc5800bfa67fe1c3fd7a8f81468e7d5a84d2a6309"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — bot_defense.policy.js_insertion_rules.rules.any_domain / e1733c4c0c69 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-77128001eda70f72f4a238fc4e496e013ab722463a3616b48aebc42a7eee39f6"></a>

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
any_domain = {}
```

<a id="canonical-546875830656811348f7ded68d0c6f0d6f79f054fb9001c9568bf5736e1d413e"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.any_domain / e1733c4c0c69 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb9c2ddba446564d273296bd7a3aa60bf2275e453dfdc01f3a470cca75cbbf56"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.any_domain / e1733c4c0c69 / 4

- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-00e2f5a19f4ebb77ede3096a873e948ff0ab32c874f5af6d2026e2982ae14a70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86b4309c7d3469bc438957721777ff55de74f4bd99a69f1630a4f96eaac078cb"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-de674dd762b4230cc54a1067df1387a8e74e0096c384d84c0deeef6f04635887"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-da8422a5943a0a45fb9e695d5367ebcd4f727c5a34c4f3f4ec4ab8c4e4ad12ce"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 3

<a id="canonical-8058214700c76972d89a27a2eae70b8f7ce3fec5ead70175a413c90d8964b9de"></a>

<a id="canonical-fa0b720a107b520be162ddf1599442e348bed64f4529a91c2d846bb9a3938df2"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-a881bde309e54592c15c22be7f64a7102e6fceed6ad88df55aaaa7c1aa8d6723"></a>

<a id="canonical-64d1438fe998ddfd2fad17224ff8db4acb7cbca36276369d236548ff8019bed1"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-b41cbc9fcde45772a4819057c57dd2093650acd0bc0eec19dae5caba7d04116c"></a>

<a id="canonical-33ffe62573e761142ee6e6c1cd7113829151fb2a2b983bdb3d1815a52cb11c6e"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-5fd3339b8b0f87bf5643f4af80e9765b39c3b3fac0e3847a41c04fdcbdb12c7f"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.domain / 03b97a9c01ca / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2df5ce039ee61d3853ff29f465b368ec2b4298d6acf90bd5738dbe7230399959"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bff5a5b789952237ebb62cd1bc4c8f874006334feeb0e4423d9f85a36ea8acdf"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — bot_defense.policy.js_insertion_rules.rules.metadata / bdf0509d0f3e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-a3d9d40a6f543aaffb730cec4ecdea868b6b17cf8cc8fbc206d6a0a8bea0acd0"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-01126ee1db47c01cd5d45ec97b484a9e975cf180dd68355915ab906051fc14f5"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.metadata / bdf0509d0f3e / 3

<a id="canonical-ff2b688270c8e15c33f049b7aeaa5ea592c9f2cc4b54748fffa3b9c7fca02dfd"></a>

<a id="canonical-fed211c10032ee67b3d482eca464e3902d423b635462f0b865a316259c36553b"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.rules.metadata / bdf0509d0f3e / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a05dfc2b6a156b9b4743c21b79d7c2743713d17f6ea4ee6a8a447eb1099c28d8"></a>

<a id="canonical-00069f8649e6a10bdb0a7b98463fb7bc6d038913669fd84e1cc509432364ff4f"></a>

## name property — bot_defense.policy.js_insertion_rules.rules.metadata / bdf0509d0f3e / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0146296ec890d7c96f23d7b1435547a827c8847f58e59bdb429a871cfe02fd33"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.metadata / bdf0509d0f3e / 6

- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7c39d581dce7bfdf400fe0f7a1bab283efc2f3d864a8e45c54897d43a22ffc82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4daa04a120ac58614e8b648b52a757eb209b88a9eeb50173ef0834dccecab537"></a>

## bot_defense.policy.js_insertion_rules.rules.path — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-d1f84956cd3136f798d4c7175e06974f53d1278f274d5dbbc389976cfb98ef39)
- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-9b8b44f421c84553d8324f9a1cd2080243be3c5fae918537e6d82a26b668f922"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-993f403a7aa82666c25103826bcf0b0df21807c3890234db292cd73aa9d9b0f8"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 3

<a id="canonical-28b5bacab3cf4294a70fd47ae4ffd65f0cd69a0380c5089a97d28f0c551ec683"></a>

<a id="canonical-2e94362c28de3d484576e98851a1fb7e60af36f611e113b3039fe361ac8e7e8b"></a>

## path property — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-9e31c68bc1c9f1e4c7921c58eb34bf1dcb02a4f951dc33f32ad74d1d2d52bf89"></a>

<a id="canonical-4f6e12914f75205deeba66396774d3efbd4006dd352067feb3e7f5e57986669e"></a>

## prefix property — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-77520ba15fd69339608bd6b81831a7961e37471b6537f82ff8998868c3d0e39b"></a>

<a id="canonical-1bdad8941d3dd47e456bf938d85c9809a8a31be871e1a09526042e78ee5b20cf"></a>

## regex property — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-113ebb1493d32457cd920f6b4b67300615460707d1d0b906f72b2e40a7879d7b"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.path / b4587d821dcf / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-8a9686742bf2d8e48d16f1a4ffe3a53bc0a71880a46cece6757f9cf3cb553463)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c77846a5f2363a42de4088087d010b691e7c98555b923a745537ce755e9989f2"></a>

## bot_defense.policy.mobile_sdk_config — bot_defense.policy.mobile_sdk_config / 9a9417abd23a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-ab0be0c2e67394b4dc1adc51000df63199d79c0165d82d7be0ae7db30554f528"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b369075b13e3239f6fc02c757969626b14adc16d372250a35853932768fe10c"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config / 9a9417abd23a / 3

- [mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98): complete subsection reference.

<a id="canonical-be2a3afae93116ad1219ca8a05ee679e3ff3c0c293aff423f36736cc751fbc4d"></a>

## Next pages — bot_defense.policy.mobile_sdk_config / 9a9417abd23a / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5543df71e0127c6fc7c932d9104cc64f6c7b524fcbe4115a4af9885fcf4d7415"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — bot_defense.policy.mobile_sdk_config.mobile_identifier / 66802ab4a91d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-2525cba804b18b9f8bff26885f6d7b54d1eae19f89ca8d446a3479fbb81ed989"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-2da63f19e5264b4e877780ccfbaaed4a73036006aa8043c9087c31be256e1562"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier / 66802ab4a91d / 3

- [headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb): complete subsection reference.

<a id="canonical-bca52b0b449c22e448b68f49f65646f06694ab25c67f9b769d7fdd17f911ef4d"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier / 66802ab4a91d / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-839f29936f0179131a7a21d06900a509f455e63d96818b10248d2b460b50cc09"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 7ed1040338dd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-5a6fd02f2351707956577f28e4c2f8333bd060edbae6d8ee80c8e31ba838b482"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-d8cfb6bdfd85bbefaeb8434ca9d7c5e831b79200bbdca8f84b6226836d99ecd5"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 7ed1040338dd / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-8c97565e05335967e8cbbb93d9319bc92ae61f09f868886cfc8a5c9309bdddc5): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-922dbbf651b6402fc2ee71353eccc900e31dd47c81056b166933e2ee6d327ee2): complete subsection reference.

- [item](resources--cdn_loadbalancer--reference--group-007.md#canonical-a0ace724fe533c4b69df83b3626a133cc9f7fcb1ab8a197636a6151174399ea4): complete subsection reference.

<a id="canonical-48b4026ea52fead03a084d7b1ea59844b787001e04d3465c927c2db17e8ae872"></a>

<a id="canonical-a628372a28bc9b5cbc5ef2a034e128277de0d4cf0866002ce992de4be4ca1aac"></a>

## name property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 7ed1040338dd / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-b423c6b6983f0c74298ef27942b4dd5a83934b63a756b6cb2cd3615e35fd9823"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 7ed1040338dd / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-8c97565e05335967e8cbbb93d9319bc92ae61f09f868886cfc8a5c9309bdddc5)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](resources--cdn_loadbalancer--reference--group-007.md#canonical-922dbbf651b6402fc2ee71353eccc900e31dd47c81056b166933e2ee6d327ee2)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](resources--cdn_loadbalancer--reference--group-007.md#canonical-a0ace724fe533c4b69df83b3626a133cc9f7fcb1ab8a197636a6151174399ea4)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8c97565e05335967e8cbbb93d9319bc92ae61f09f868886cfc8a5c9309bdddc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf9dfbf380441eb26c26c28bac67696a66c1d310419188c81e31a0b9e5eb46e0"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 654415fb0dab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-0945681754b1914060c93afa757711e03acd7502c672faa73905db37f90c4fbc"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-73b7e3a27f83308df4ea2e9b47752bb8933dd93a59e23ec27b9f659da093682b"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 654415fb0dab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-582780a72f7c48053565fc987b21ef686d498d0c5cc0da350170f4402a235658"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 654415fb0dab / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-922dbbf651b6402fc2ee71353eccc900e31dd47c81056b166933e2ee6d327ee2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d52604491c63aaeb9af96557f3394f91e49b50a794624180623b622bc253738"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 148822a4a405 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-7a7aec443746546cbee79de49a8a487db2a7488d5c12afd93ac66132635d45f0"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-c7de687927624f23b048a186054a34172e5d23190300db2d97644e5135ea8ebf"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 148822a4a405 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4b49432e5d3aa5256aab2323b934c9a8f8a592426d49e9901c65a453f60e64d"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 148822a4a405 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a0ace724fe533c4b69df83b3626a133cc9f7fcb1ab8a197636a6151174399ea4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-751b8c3877cb4d96e1a5c9b2153e9644cc23835e4419ce215f64928751fdb61d"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.mobile_sdk_config](resources--cdn_loadbalancer--reference--group-007.md#canonical-2553e2626be82bec4464f495bf27aa0533ca244a88ab61527723e79eec52287a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--cdn_loadbalancer--reference--group-007.md#canonical-e20c441928e8754de9d1687c61dfef0d44ed5cdad7833d5211cbedd10117ba98)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-348360dc24e688cfdb1408f5e0adecdc24590c021c623a2308c3f27a38f92aba"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-c249a8baea3cefe1260c744a42197a5bff35abf82f209ec50bad93e0c8c3d9ff"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 3

<a id="canonical-d58ffbb676bc20a6717255fb96ca557826d1d9acda6e605086a8025f28bd4892"></a>

<a id="canonical-6631a122079d89f5cc6df1d61433a662d5dbfdf56b298021688a2bd973392c36"></a>

## exact_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-f91448742376a2b2b0a6f53bd568d0627d9493e388bf29d3296464ed773afa4b"></a>

<a id="canonical-497bfe3330557fc5d56f4ea44909e05e5f6d6177a4a0555ab98f45549e9a489a"></a>

## regex_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-6565deeb730a4ae7afc4c71e155af4b00c1455fb19900831ebe65e5f1f2b21cd"></a>

<a id="canonical-ad3c04fe0970a7542d55531eae0cec2fb34b3047f4f829bdc64b9ee3f978edce"></a>

## transformers property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-05086e8bffb02f307f94ebb7d01592a04aff98d2789b0d31b3d67e90628c9e0e"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 68d2843f21f0 / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-8380315f28b43d6fcb6aff6d109dc548e6a9f88157d411133f071c264b3822bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-560127813ea36c7b279f8973d0d690ef7664ac6e21fdfa93c3718e659ca72099"></a>

## bot_defense.policy.protected_app_endpoints — bot_defense.policy.protected_app_endpoints / 0ef12a6642fa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-c1665aa641aafe01bcbb4d7d2804a3793ab81257e567c7807d6e96a2a5267e88"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("allow_good_bots",
    "mitigate_good_bots"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile",
    "web"),
  validators.ConflictingListObjectAttributes("mobile",
    "web_mobile"),
  validators.ConflictingListObjectAttributes("web",
    "web_mobile")}
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

Terraform syntax:

```terraform
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-8c6db76c79de204e6ffff431ec59602e2eeb200566d473eea994ac3bf040f2eb"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints / 0ef12a6642fa / 3

- [allow_good_bots](resources--cdn_loadbalancer--reference--group-007.md#canonical-87f7d154e834b655bb6041832df94d87878c78a941af67715a169928798fe8ed): complete subsection reference.

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-ebc58ad7cc26596f71391343e598ee91d73355a5e6a86cc51a0b6a9784968faf): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-c4ff27344b553683decdca6c18be063e421ba6d2392abf71b66957d988c2faa9): complete subsection reference.

- [flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-5050d23a49f2710337c38d68110ede2e89f95d017c27256ec8ac176a3456ceb2): complete subsection reference.

<a id="canonical-c1d73382655cb6a4beae2c37e3023d892ccd38dd8a83c1ab46369e8f310a36dd"></a>

<a id="canonical-a36f52c5fc4689dd8c5204dd30d1d7501da69a29d9301e33565c0f4a648e5984"></a>

## http_methods property — bot_defense.policy.protected_app_endpoints / 0ef12a6642fa / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-f471de63fd69f94c5921649127ee5d6712ee6ee4cab4f27c40db8f4ebc288d90): complete subsection reference.

- [mitigate_good_bots](resources--cdn_loadbalancer--reference--group-008.md#canonical-352c9c6b2777c8bfcd2d546c2c818a300e1a1337e2a59cbdcd5b026ea8da08ca): complete subsection reference.

- [mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-f16163a87e781bdfadca2e4dbb0ad838c5a685bd6f14b94283036c08c9cca611): complete subsection reference.

- [mobile](resources--cdn_loadbalancer--reference--group-009.md#canonical-b8b5785988c2538a52493ad8157015d8f1ba98bac382a915fd589b9380999a6d): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-abdeb6e2af99c135f6189ecbc2cbb3bddda5ff11a216c11dc75c9a3a6bc39bac): complete subsection reference.

<a id="canonical-c67f7889fb5f59d866b572a19264d46d6bb8ae93712bc4d853c701fcc9d622a8"></a>

<a id="canonical-d007a7d02e8c825c6312cbd15ba6316a9e294dcca8f0c607c711eb246331a3b7"></a>

## protocol property — bot_defense.policy.protected_app_endpoints / 0ef12a6642fa / 5

Type: `"string"`. Optional.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOTH",
    "HTTP",
    "HTTPS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b): complete subsection reference.

- [undefined_flow_label](resources--cdn_loadbalancer--reference--group-009.md#canonical-536994a1b4b82c44b0a6cdce0bb07153922fd8b35f3f716b9800ddc3784246cc): complete subsection reference.

- [web](resources--cdn_loadbalancer--reference--group-009.md#canonical-1783c8921d2e090218141c9bc3dc8b752ed8a54e2470ed1838c571db1c056084): complete subsection reference.

- [web_mobile](resources--cdn_loadbalancer--reference--group-009.md#canonical-4fcf4fe1ba8a4f2d73ee2ae019775f61bddbe047d0f65fd6e7f1652cece45590): complete subsection reference.

<a id="canonical-87856429e4317d647d1a506c4e08f681b33ee9711aeebf76f900f45789aa6bd1"></a>

## Next pages — bot_defense.policy.protected_app_endpoints / 0ef12a6642fa / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](resources--cdn_loadbalancer--reference--group-007.md#canonical-87f7d154e834b655bb6041832df94d87878c78a941af67715a169928798fe8ed)
- [bot_defense.policy.protected_app_endpoints.any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-ebc58ad7cc26596f71391343e598ee91d73355a5e6a86cc51a0b6a9784968faf)
- [bot_defense.policy.protected_app_endpoints.domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-c4ff27344b553683decdca6c18be063e421ba6d2392abf71b66957d988c2faa9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-5050d23a49f2710337c38d68110ede2e89f95d017c27256ec8ac176a3456ceb2)
- [bot_defense.policy.protected_app_endpoints.metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-f471de63fd69f94c5921649127ee5d6712ee6ee4cab4f27c40db8f4ebc288d90)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](resources--cdn_loadbalancer--reference--group-008.md#canonical-352c9c6b2777c8bfcd2d546c2c818a300e1a1337e2a59cbdcd5b026ea8da08ca)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-f16163a87e781bdfadca2e4dbb0ad838c5a685bd6f14b94283036c08c9cca611)
- [bot_defense.policy.protected_app_endpoints.mobile](resources--cdn_loadbalancer--reference--group-009.md#canonical-b8b5785988c2538a52493ad8157015d8f1ba98bac382a915fd589b9380999a6d)
- [bot_defense.policy.protected_app_endpoints.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-abdeb6e2af99c135f6189ecbc2cbb3bddda5ff11a216c11dc75c9a3a6bc39bac)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](resources--cdn_loadbalancer--reference--group-009.md#canonical-536994a1b4b82c44b0a6cdce0bb07153922fd8b35f3f716b9800ddc3784246cc)
- [bot_defense.policy.protected_app_endpoints.web](resources--cdn_loadbalancer--reference--group-009.md#canonical-1783c8921d2e090218141c9bc3dc8b752ed8a54e2470ed1838c571db1c056084)
- [bot_defense.policy.protected_app_endpoints.web_mobile](resources--cdn_loadbalancer--reference--group-009.md#canonical-4fcf4fe1ba8a4f2d73ee2ae019775f61bddbe047d0f65fd6e7f1652cece45590)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-87f7d154e834b655bb6041832df94d87878c78a941af67715a169928798fe8ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae59045472a178a716ac45ae4b93ca477acb0369f5ec67b1b2c440026914e0ad"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — bot_defense.policy.protected_app_endpoints.allow_good_bots / cf4f1c639e78 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-a678a0d641a78a2b18cbc61a9cb9f592aeafbfcd4fd92640bda5e0b7e94987a0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow good bots.

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
allow_good_bots = {}
```

<a id="canonical-21e7f83277fd27e37697292b08a152e6d42e1a7967edfbeaf0c86f4d370b5223"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.allow_good_bots / cf4f1c639e78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-282860072a448e92c941adb06dd97a90b9cd5660315bdf2a017125f78d6d3636"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.allow_good_bots / cf4f1c639e78 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ebc58ad7cc26596f71391343e598ee91d73355a5e6a86cc51a0b6a9784968faf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c088f28dbfc3c9e04ad0c6e6958629b495f6253fec44cd985a01c6892d9e3049"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — bot_defense.policy.protected_app_endpoints.any_domain / b01400d87fd4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-8c6cfe912dbd83deeebfa90339b22425ec05800b68372650b47e91b965bdeef9"></a>

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
any_domain = {}
```

<a id="canonical-c4105b80bb9884f66b5e19d7ee5371d87a8752bbb73266691f38c57769243712"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.any_domain / b01400d87fd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d741b19fa0989721545c0f87248858fca8fc0f2bf6ca30cbe0c07489850dcef4"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.any_domain / b01400d87fd4 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c4ff27344b553683decdca6c18be063e421ba6d2392abf71b66957d988c2faa9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8b98dd30afb7079c671f9957b277fb2a4aa7c9f2dd930a983671d6ca1ef571b"></a>

## bot_defense.policy.protected_app_endpoints.domain — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-e7f91ea6f6779b3927187756474dc192c4138fc9116e1945ea8747683b519edf"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-c919e4c947a23bc704835a2b18cdec5f23a6251a024c72e8917f3ee9c4a2348f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 3

<a id="canonical-ef6466eb6562fb2641dd66c7ef7e7065f81dce74e3664f495b36f72d7466d6a4"></a>

<a id="canonical-81f206ffd1a8c3fe82e04228f0efee563db4cd19addbf39e92e680f820d1431e"></a>

## exact_value property — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-d643e06eae0faec56f4d14259d9eccba5420825fd5918a5fe5f5d779a0bafc11"></a>

<a id="canonical-a11619f49bb21d59c01e53926ec52fe200eecfb548cbb64b5dbb77198282573f"></a>

## regex_value property — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-5792757f104d0fed5464b43b691679bf95793a1f84bfb70222288cd0d2d4a6af"></a>

<a id="canonical-f627bbcc6b9ad63f49854f80cf8e1e4ae8eaf4f14b8ba1dfe762e1a2ff6d01e0"></a>

## suffix_value property — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-88058be1997665f6e00c2f5590b30e50db6af6670e469d13a1773dbc12b4dbc6"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.domain / 474aac7e26fa / 7

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50421900bd93e01601be5cc5a0a52c4fc78e25851147a97e8df70c8052833042"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — bot_defense.policy.protected_app_endpoints.flow_label / c97e729c12d0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-45da7bb6ceaeda23d30629741916ab31b19edcf713925abac9c0d8de3e12df25"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad937372382117030f3959ad33a46d7afacaf03c63da82fe2e9d98238bc15c07"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label / c97e729c12d0 / 3

- [account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8): complete subsection reference.

- [authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-e3da21fa35aff647f5498fd4002e343e3ae15690009f0429982c82f5883f4182): complete subsection reference.

- [financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-799efabbb807d8ad48978f3be98efd52a8c4df2894e550288764a9c98e632cbe): complete subsection reference.

- [flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-3e30c85bd6294be6b35f63449d98dd89698e10d102a67a5d454518f839b75c6c): complete subsection reference.

- [profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-961cbba0ad5002592dff6e707fc03c92593d83085fc52e45c7d90a15e0d15b0e): complete subsection reference.

- [search](resources--cdn_loadbalancer--reference--group-008.md#canonical-6a7c5814bbdf7279f8b54f02fc5cfd8e6746206b514d85841cb4a073b16a85c7): complete subsection reference.

- [shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-9ab4070f1626a3cb981ec45cc65874aa24a2f8975b049ad54ea8efd8ca4be1d2): complete subsection reference.

<a id="canonical-dab7ca515b83f35c6b2d1082ed63e04cf835bc97a05494f904566961a911894e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label / c97e729c12d0 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-e3da21fa35aff647f5498fd4002e343e3ae15690009f0429982c82f5883f4182)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-799efabbb807d8ad48978f3be98efd52a8c4df2894e550288764a9c98e632cbe)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-3e30c85bd6294be6b35f63449d98dd89698e10d102a67a5d454518f839b75c6c)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-961cbba0ad5002592dff6e707fc03c92593d83085fc52e45c7d90a15e0d15b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-6a7c5814bbdf7279f8b54f02fc5cfd8e6746206b514d85841cb4a073b16a85c7)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-9ab4070f1626a3cb981ec45cc65874aa24a2f8975b049ad54ea8efd8ca4be1d2)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f05703523da7c29c5c74babcfc56541e592e22c29849fcdd06711775d0150347"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 715d66da0f28 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-824140303fe34665756b18a891d52a12978edfb1ab67a674e0340f741652b784"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-27ff9df460da89b1cebd6446c4723a5c7ac7dfd625f8372c842dd09f17359802"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 715d66da0f28 / 3

- [create](resources--cdn_loadbalancer--reference--group-007.md#canonical-7eb2ddcf12db27858331d7b9595b0faaeba0df5b6bed4cfe30a8ed3eebd9859d): complete subsection reference.

- [password_reset](resources--cdn_loadbalancer--reference--group-007.md#canonical-92e24f0e1bd197bc65fc381183f004416f174ea76b355fbb71e7e8b6dca62b26): complete subsection reference.

<a id="canonical-a1667e54cb973f5cb26449c2a668056c51ad0b31e8e868f0b5473fd33de4fcbd"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 715d66da0f28 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](resources--cdn_loadbalancer--reference--group-007.md#canonical-7eb2ddcf12db27858331d7b9595b0faaeba0df5b6bed4cfe30a8ed3eebd9859d)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](resources--cdn_loadbalancer--reference--group-007.md#canonical-92e24f0e1bd197bc65fc381183f004416f174ea76b355fbb71e7e8b6dca62b26)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7eb2ddcf12db27858331d7b9595b0faaeba0df5b6bed4cfe30a8ed3eebd9859d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa8f30662b63959277476b2b42ba35012b40b22828efb0e10365cc1a104438e"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 1c2f182f61e8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-db0e42666a8961e312d994f632653c490b009d7fc0c9ce9d8baef5ce7a953a83"></a>

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
create = {}
```

<a id="canonical-127b550f4e9ed1639f5a1b0ba5d7592886b98a68f24711e1914e7194514da5f8"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 1c2f182f61e8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dbee411be1a2685ee37051f6fcffc61b646a925b12805c6ebc6372169bd70634"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 1c2f182f61e8 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-92e24f0e1bd197bc65fc381183f004416f174ea76b355fbb71e7e8b6dca62b26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6d162294dd19914c2c542ffd98de1444a47643ba530f305443fa9648499e83d"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / da3023bb0fe1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-fbb0ec93b562fb8c9bc4a798e7337797c16b5086a37a9982ebdc59706164fd2a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

<a id="canonical-35ce6e08d9867cb9de5800661916a72db8bddeab57cf37b2bf4fc168123fbdf0"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / da3023bb0fe1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6dd26132ca9723f114cc14cc5e2e727ee9db3eb262d05aa651df29d30d07b5e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / da3023bb0fe1 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-7172ca4b0f34588561644a1f559784a745c3a837771447da41a9260cbda75bc8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e3da21fa35aff647f5498fd4002e343e3ae15690009f0429982c82f5883f4182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63bb427dc7692c096972e4140d50d1dd2eae0d971228c1a86cbe45d398a6dec3"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 87f66448856c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-483c4e96b25910bf456f807c35815029e101f291433e38f825d3bed752320b69"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-399444fbdd5ac1be992f3b6ecfbc4f4286b795926e235a83d36235781e295eed"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 87f66448856c / 3

- [login](resources--cdn_loadbalancer--reference--group-007.md#canonical-929922c0087873eb5da1d1c302701a9a71de653f87438ead65d89a96bb67f686): complete subsection reference.

- [login_mfa](resources--cdn_loadbalancer--reference--group-008.md#canonical-0eed9c065aefcd47e754db763092784e1f86591b121dcaad759cb19ad1b4179f): complete subsection reference.

- [login_partner](resources--cdn_loadbalancer--reference--group-008.md#canonical-db3f09f768ebb2392c30b3adc3fadfe6cf1c71503140181a5b73e882c4124b89): complete subsection reference.

- [logout](resources--cdn_loadbalancer--reference--group-008.md#canonical-558c2b0505e4d02d4671d336a49d1e173195c02f7eb17ff78a08e7b4cb7885ef): complete subsection reference.

- [token_refresh](resources--cdn_loadbalancer--reference--group-008.md#canonical-1f9a37f461385556f8f16ce053c744ae033714972f7ac04c1db02052c2f58c56): complete subsection reference.

<a id="canonical-b83b73dfc7190b7d82ff13d761d746cbcd614ce594f695f5bec9c7e99a0766a1"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 87f66448856c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-007.md#canonical-929922c0087873eb5da1d1c302701a9a71de653f87438ead65d89a96bb67f686)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](resources--cdn_loadbalancer--reference--group-008.md#canonical-0eed9c065aefcd47e754db763092784e1f86591b121dcaad759cb19ad1b4179f)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](resources--cdn_loadbalancer--reference--group-008.md#canonical-db3f09f768ebb2392c30b3adc3fadfe6cf1c71503140181a5b73e882c4124b89)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](resources--cdn_loadbalancer--reference--group-008.md#canonical-558c2b0505e4d02d4671d336a49d1e173195c02f7eb17ff78a08e7b4cb7885ef)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](resources--cdn_loadbalancer--reference--group-008.md#canonical-1f9a37f461385556f8f16ce053c744ae033714972f7ac04c1db02052c2f58c56)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-929922c0087873eb5da1d1c302701a9a71de653f87438ead65d89a96bb67f686"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0527caa19b1c292265cf417f50682cac0184c6585e11ff18bff56ae1970b8987"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 2344c075d639 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-37b19a61e2569a361e4f37614e2d5fe9947b119048d43ab4261367ba574aff1c)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-007.md#canonical-e3da21fa35aff647f5498fd4002e343e3ae15690009f0429982c82f5883f4182)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-69af84df30cec7d109a5ff4967e2fbf5c89511023ec5d5c82f5700790cb62869"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ad42ed9575a7a56c41fbd77d98821d45d8f967543121d666a0521e12cad0985"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 2344c075d639 / 3

- [disable_transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-3e0cbe009a776d00ba05596e4067cc07db1c937cc19858463ee141febb6f6460): complete subsection reference.

- [transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-8a7edf417e6ccace43eb9019134705415e90db9babd7e069ea45eb0b027aea2a): complete subsection reference.
