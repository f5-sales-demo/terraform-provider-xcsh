---
page_title: "xcsh_cdn_purge_command reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command reference."
---

# xcsh_cdn_purge_command reference

<a id="canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf0248ebec19275649a33668ccf7a299c8d97ea8d949d6aef6403024c75eced8"></a>

## Property reference — Property reference / 2e91d610bd33 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- Property reference

<a id="canonical-f6e619959b5427b205e2a8076bdadf68011a73cda3c24fbf32b7b3404a13394a"></a>

## Direct properties — Property reference / 2e91d610bd33 / 3

<a id="canonical-42f44f6fbd620414e666e2ddae7a4676ecf3c108e969d581110ff64e6505a8a8"></a>

<a id="canonical-71fd5d81d3f24feffd65998a815fe222f7ff7242f91b40248e0ed1ef817e4e03"></a>

## annotations property — Property reference / 2e91d610bd33 / 4

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

<a id="canonical-ad255f12fee9ff0383992c7b930923f09f480bff86bc98964e80d796e645e1cc"></a>

<a id="canonical-8aaafe6a5efd943f8f7c722c5cfaf8087003aefa7240f7df3f998132e7c65e03"></a>

## description property — Property reference / 2e91d610bd33 / 5

Type: `"string"`. Computed.

Description of the CDNPurgeCommand.

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

- [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-4374766b3e2c3a0eca8a3bd38340d9e08f09b06567f1d348449c4d8721e07d8b): complete subsection reference.

<a id="canonical-636c0e572ea27ae314bd97cf1fbf76324c8128e9b30d8adf197d7cfb87874824"></a>

<a id="canonical-33009225f2360859f6b19a2f4a2a20f8ea977f60da6dee988fbf116ac87cd4e2"></a>

## hostname property — Property reference / 2e91d610bd33 / 6

Type: `"string"`. Computed.

\[OneOf: hostname, pattern, purge\_all, url\_path\] Exclusive with \[pattern purge\_all url\_path\]
Purge cached content by Hostname.

Upstream description:

Exclusive with \[pattern purge\_all url\_path\] Purge cached content by Hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

OneOf alternatives in this subsection:

- [hostname](data-sources--cdn_purge_command--reference--group-001.md#canonical-636c0e572ea27ae314bd97cf1fbf76324c8128e9b30d8adf197d7cfb87874824)
- [pattern](data-sources--cdn_purge_command--reference--group-001.md#canonical-f8eed7d6579c01a6763b5143898b3aa5b2469bf79a37171b42c5e9d61d6da24a)
- [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-111d8732d65edad7dfad91e6f954e306cb1a54c665e7024089233698c0a8dec2)
- [url_path](data-sources--cdn_purge_command--reference--group-001.md#canonical-ff035d9fd1af76056dd514ae9373c4bdb399cb025807fc9997fa35a18b514420)

Select alternatives according to the provider validators above.

<a id="canonical-d35c99785362f4deedd2684529cbfec2b2c764e7ea79bc9cd550309519c88aad"></a>

<a id="canonical-2a684c763bb8031d1e13e5b5322f942ecf90c64e190d025fa82efc708d770058"></a>

## id property — Property reference / 2e91d610bd33 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f121c05d30d74696084cdb108b9bb01dacccf3a4e3d0760197c4f50c36944b92"></a>

<a id="canonical-884f5005394fb7b6ba3abc9aaa9feaa7b038a6ba33ae370473ce07d47052ee2f"></a>

## labels property — Property reference / 2e91d610bd33 / 8

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

<a id="canonical-6a301837f513638a9814c7165b6700c6fbbcffcbcc74d5e04cf9f28e3ddc87f2"></a>

<a id="canonical-2317042e178d6799c9358925a8db888b0da06efb88badf6848d6a3bfaafabdb2"></a>

## name property — Property reference / 2e91d610bd33 / 9

Type: `"string"`. Required.

Name of the CDNPurgeCommand.

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

<a id="canonical-71e4607f6541312e77eee80e033ca9ea3ae17289c2d95325f09fa0acacfd70b6"></a>

<a id="canonical-f8dcc763e3301af6ce4c117a60855c376fb9565f073ab331472b951724895015"></a>

## namespace property — Property reference / 2e91d610bd33 / 10

Type: `"string"`. Required.

Namespace where the CDNPurgeCommand exists.

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

<a id="canonical-f8eed7d6579c01a6763b5143898b3aa5b2469bf79a37171b42c5e9d61d6da24a"></a>

<a id="canonical-b54aa1e71d949ebe9f8c463a828360378a8465a55fffdb9049f10502449b4cd6"></a>

## pattern property — Property reference / 2e91d610bd33 / 11

Type: `"string"`. Computed.

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

Upstream description:

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-5483f29d0a278357ca1189570c6a39eaccc3d57a00732f95e6456723c83739e4): complete subsection reference.

- [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-fb4fb09f8c68f85f9c5788a8f34feae59a334447f7ff47fe10f5a3af9505b434): complete subsection reference.

<a id="canonical-ff035d9fd1af76056dd514ae9373c4bdb399cb025807fc9997fa35a18b514420"></a>

<a id="canonical-48563d661e99b31d82c8be7537c24b29c1cd429cf308ff68bd7515e8c9a9a3a4"></a>

## url_path property — Property reference / 2e91d610bd33 / 12

Type: `"string"`. Computed.

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

Upstream description:

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

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

- [virtual_host](data-sources--cdn_purge_command--reference--group-001.md#canonical-bc7b51280683fa823f4c73165b6356d246f50cc8395d07deec64f08e073247da): complete subsection reference.

<a id="canonical-28b58368370f87266fb23661468825569eb5c1fcbc1c7176843b274519b63e15"></a>

## All schema paths — Property reference / 2e91d610bd33 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cdn_purge_command--reference--group-001.md#canonical-42f44f6fbd620414e666e2ddae7a4676ecf3c108e969d581110ff64e6505a8a8) |
| `description` | [description](data-sources--cdn_purge_command--reference--group-001.md#canonical-ad255f12fee9ff0383992c7b930923f09f480bff86bc98964e80d796e645e1cc) |
| `hard_purge` | [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-c750db9c17e73e2bc49d5ea3ce0af77aaa071c26dacd70888bd831f8392930fb) |
| `hostname` | [hostname](data-sources--cdn_purge_command--reference--group-001.md#canonical-636c0e572ea27ae314bd97cf1fbf76324c8128e9b30d8adf197d7cfb87874824) |
| `id` | [id](data-sources--cdn_purge_command--reference--group-001.md#canonical-d35c99785362f4deedd2684529cbfec2b2c764e7ea79bc9cd550309519c88aad) |
| `labels` | [labels](data-sources--cdn_purge_command--reference--group-001.md#canonical-f121c05d30d74696084cdb108b9bb01dacccf3a4e3d0760197c4f50c36944b92) |
| `name` | [name](data-sources--cdn_purge_command--reference--group-001.md#canonical-6a301837f513638a9814c7165b6700c6fbbcffcbcc74d5e04cf9f28e3ddc87f2) |
| `namespace` | [namespace](data-sources--cdn_purge_command--reference--group-001.md#canonical-71e4607f6541312e77eee80e033ca9ea3ae17289c2d95325f09fa0acacfd70b6) |
| `pattern` | [pattern](data-sources--cdn_purge_command--reference--group-001.md#canonical-f8eed7d6579c01a6763b5143898b3aa5b2469bf79a37171b42c5e9d61d6da24a) |
| `purge_all` | [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-111d8732d65edad7dfad91e6f954e306cb1a54c665e7024089233698c0a8dec2) |
| `soft_purge` | [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-b5b925da87e8d1a43c93761dd880bc14df3cf6cadd410bb7e76e7ad706cc6006) |
| `url_path` | [url_path](data-sources--cdn_purge_command--reference--group-001.md#canonical-ff035d9fd1af76056dd514ae9373c4bdb399cb025807fc9997fa35a18b514420) |
| `virtual_host` | [virtual_host](data-sources--cdn_purge_command--reference--group-001.md#canonical-22bd7a7244c936102a22b9d4966de3bef68f4d39376a4044ff6ad72a48d6e869) |
| `virtual_host.name` | [virtual_host.name](data-sources--cdn_purge_command--reference--group-001.md#canonical-c5c6abb1679de7bee4010e60f348efcc8eaa427682dd2e42cf7b4eff0b8aa9b3) |
| `virtual_host.namespace` | [virtual_host.namespace](data-sources--cdn_purge_command--reference--group-001.md#canonical-a9b9e1d5b111f57f62f24aa88413c6a50d3f2530dc93c48076e0ce5c8054caa5) |
| `virtual_host.tenant` | [virtual_host.tenant](data-sources--cdn_purge_command--reference--group-001.md#canonical-bea619e1cdf6e610de6ab501a6a1d120467c0fe8079ca77925d9d0e3173f5bed) |

<a id="canonical-5d2594c920dd8ddd310995f48dce03b4b71ab5b88cdb24bd1612b601f99156ee"></a>

## Next pages — Property reference / 2e91d610bd33 / 14

- [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-4374766b3e2c3a0eca8a3bd38340d9e08f09b06567f1d348449c4d8721e07d8b)
- [purge_all](data-sources--cdn_purge_command--reference--group-001.md#canonical-5483f29d0a278357ca1189570c6a39eaccc3d57a00732f95e6456723c83739e4)
- [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-fb4fb09f8c68f85f9c5788a8f34feae59a334447f7ff47fe10f5a3af9505b434)
- [virtual_host](data-sources--cdn_purge_command--reference--group-001.md#canonical-bc7b51280683fa823f4c73165b6356d246f50cc8395d07deec64f08e073247da)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)

<a id="canonical-4374766b3e2c3a0eca8a3bd38340d9e08f09b06567f1d348449c4d8721e07d8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad2e643a22865e2d2af3c1e8493d35249992d7f32307e9ffb62e00a9464c6fcf"></a>

## hard_purge — hard_purge / 51ef266444e9 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- hard_purge

<a id="canonical-c750db9c17e73e2bc49d5ea3ce0af77aaa071c26dacd70888bd831f8392930fb"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hard\_purge, soft\_purge\] Enable this option

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

- [hard_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-c750db9c17e73e2bc49d5ea3ce0af77aaa071c26dacd70888bd831f8392930fb)
- [soft_purge](data-sources--cdn_purge_command--reference--group-001.md#canonical-b5b925da87e8d1a43c93761dd880bc14df3cf6cadd410bb7e76e7ad706cc6006)

Select alternatives according to the provider validators above.

<a id="canonical-f6a1c17fc788256c46b8fba4245f7bd6fd3fdd5f38c05e7a6023db435250b295"></a>

## Direct properties — hard_purge / 51ef266444e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1589db8c971a8e088760b2715bf40e40c3f250da7ecb9738eb310e4f632f36d"></a>

## Next pages — hard_purge / 51ef266444e9 / 4

- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)

<a id="canonical-5483f29d0a278357ca1189570c6a39eaccc3d57a00732f95e6456723c83739e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14d52c47f6d24ed96146367beb877587eedb74d27cb722e5b731b4b175eaeb13"></a>

## purge_all — purge_all / 9ce051017799 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- purge_all

<a id="canonical-111d8732d65edad7dfad91e6f954e306cb1a54c665e7024089233698c0a8dec2"></a>

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

<a id="canonical-9217e536ff052033824601ff0455ba1e0cafd28925cb3957dedf682516e6612c"></a>

## Direct properties — purge_all / 9ce051017799 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89e022c33e7d1190e8259d2ad231a8b11a4abe0de0b365073f9aeb6646bb3e53"></a>

## Next pages — purge_all / 9ce051017799 / 4

- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)

<a id="canonical-fb4fb09f8c68f85f9c5788a8f34feae59a334447f7ff47fe10f5a3af9505b434"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de0c3b89543065e64006019c755a45301f0699ddcb3a8fdd4aefc3f51c7633fb"></a>

## soft_purge — soft_purge / 396e00c40021 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- soft_purge

<a id="canonical-b5b925da87e8d1a43c93761dd880bc14df3cf6cadd410bb7e76e7ad706cc6006"></a>

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

<a id="canonical-6ac4cd765f866f85c9b54077d4252b2467d03c066d8549f4d39092735610d3c5"></a>

## Direct properties — soft_purge / 396e00c40021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09d3ec4d4eccdd4773ca6a36764d47cc71a9eed6c9f8c6e2509189e1d0f9465a"></a>

## Next pages — soft_purge / 396e00c40021 / 4

- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)

<a id="canonical-bc7b51280683fa823f4c73165b6356d246f50cc8395d07deec64f08e073247da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fbd4aff782509cf622981ca515e924f40880c1e0a9fe79313b784e9b60559ef"></a>

## virtual_host — virtual_host / 25f0a532f6b4 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- virtual_host

<a id="canonical-22bd7a7244c936102a22b9d4966de3bef68f4d39376a4044ff6ad72a48d6e869"></a>

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

<a id="canonical-9ee99810dd65bd04157eb4394a4cb3d232befe9517e766dda700e409e3bdda33"></a>

## Direct properties — virtual_host / 25f0a532f6b4 / 3

<a id="canonical-c5c6abb1679de7bee4010e60f348efcc8eaa427682dd2e42cf7b4eff0b8aa9b3"></a>

<a id="canonical-3f89cfcb68ddb5fb32b7c9f677ba9c8e97091b7a5c5bd7d59fc65052a5644d5e"></a>

## name property — virtual_host / 25f0a532f6b4 / 4

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

<a id="canonical-a9b9e1d5b111f57f62f24aa88413c6a50d3f2530dc93c48076e0ce5c8054caa5"></a>

<a id="canonical-f3bf453cdbf8bcb7690c3f830bf6c41bd9ce4e05d678851637ea403ea36d5a1d"></a>

## namespace property — virtual_host / 25f0a532f6b4 / 5

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

<a id="canonical-bea619e1cdf6e610de6ab501a6a1d120467c0fe8079ca77925d9d0e3173f5bed"></a>

<a id="canonical-62f49869a76321a9abd371f21f467f8fced0b0f0d8bcadd73eb49eef28ed48c1"></a>

## tenant property — virtual_host / 25f0a532f6b4 / 6

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

<a id="canonical-71cd2f532b90ed812c37fd1dcaf0f1b629c1e492444ec2380a9a0007dfe705be"></a>

## Next pages — virtual_host / 25f0a532f6b4 / 7

- [Property reference](data-sources--cdn_purge_command--reference--group-001.md#canonical-bfb3a1d5ef969fc07c8a01794300d9a760cd621bff3a35dfc4dfb8ef43f1ec0e)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
