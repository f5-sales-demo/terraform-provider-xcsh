---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-de255b4567a2f226a97c30866b898d06f983c5dc1b19f276a1c1a56219d65744"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10467a555190fb7caf1cfba84f546b59d1d08460104c65f2a5053aec448a89b9"></a>

## routes.simple_route.headers — routes.simple_route.headers / 8ebdbc6d10f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.headers

<a id="canonical-d315331cfdc59ae7dd8fc140b108c36e3427137a9b40959d54304863885b169e"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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

<a id="canonical-3a79946b52b2762e0e8ca695b7174c661682be4110412460562bd3a5ee93fc6f"></a>

## Direct properties — routes.simple_route.headers / 8ebdbc6d10f7 / 3

<a id="canonical-a56df5c737ef9445783defe41d59106eb044b46bc822b73ae3202ffe24248545"></a>

<a id="canonical-78182936d48328bdba344e35972688902b50f3f5ae692c3701f1a5fecaba74ee"></a>

## exact property — routes.simple_route.headers / 8ebdbc6d10f7 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-fe8037a88f06fd9bf33d6044c72ac6b3a0ed9911d9649bf9d2ac3abf2e1124f4"></a>

<a id="canonical-4b156fa0d15819e51f0cb9bfd624b2022c9f491082e331ec8bca15f32215b124"></a>

## invert_match property — routes.simple_route.headers / 8ebdbc6d10f7 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-9986971e57e61d612df01f3e5b918defadfcb4693cc37522b1ab53de131249e6"></a>

<a id="canonical-94b6ff5b9a4f6df73fbda7f62f8e7584b3624442d9e6033b1a8353ee36f7d289"></a>

## name property — routes.simple_route.headers / 8ebdbc6d10f7 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-e62429d68f1fd780605dd3ea7b3a1bdff91fcd94141bc70a76af30c2c0d5ce4b"></a>

<a id="canonical-3c890d26aae9e230f0d1e1778dc2929c0e18317b71e71a480a3203994722d112"></a>

## presence property — routes.simple_route.headers / 8ebdbc6d10f7 / 7

Type: `"bool"`. Computed.

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

<a id="canonical-8f72775d2daa57872b99451e644e74b0560dac48c1b7d4454abd1a1d7ee1cf96"></a>

<a id="canonical-8ef51d2c270dd283bd53809649353542339d5157c62dbe12e9161f64d411cc7c"></a>

## regex property — routes.simple_route.headers / 8ebdbc6d10f7 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-d6d26fffcb6e77e7b31b2121bd9e9f3ba50982a25edca0c40815e83bf196ef52"></a>

## Next pages — routes.simple_route.headers / 8ebdbc6d10f7 / 9

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-29cba187b019955171440e9188dcedbe57464be112234f015041aea135f6c35f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3161b5fbf1785912b649013b10fce7b0e29c2614a9a167b2592b215fbdbcc244"></a>

## routes.simple_route.incoming_port — routes.simple_route.incoming_port / dd73fa62388e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.incoming_port

<a id="canonical-4eef22c2becf3e11688abf0961040fd10b5c04cef32dbda9a1fbcb4d137817fa"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-75c202cb89c66c60dc892575123752530c85e82c3c898f2ae16018b2acc01188"></a>

## Direct properties — routes.simple_route.incoming_port / dd73fa62388e / 3

- [no_port_match](data-sources--http_loadbalancer--reference--group-025.md#canonical-4c260984063466891c51238e0009c316b3a4ab555ccfe783785f475169dd3814): complete subsection reference.

<a id="canonical-8dc97e39347a7bc9f81eb80183e44606c46b6ddeec678a7580b216ee6539a385"></a>

<a id="canonical-6fe01d76c668587f3ad11255a8df304de1e5c24f884c3cd7e90e494eed7e6eb3"></a>

## port property — routes.simple_route.incoming_port / dd73fa62388e / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-d8a340795b6046c985428a794077262462822a66050972d0cc359f2447a4d745"></a>

<a id="canonical-aea741c19fac9a16caae4bff88878f838a9ee315b6b6afb6c9555ec448ba9575"></a>

## port_ranges property — routes.simple_route.incoming_port / dd73fa62388e / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-f787e9e49ed9b5bb274ef00396719fb90350b47ed46e7edef46f9af7aa29ae43"></a>

## Next pages — routes.simple_route.incoming_port / dd73fa62388e / 6

- [routes.simple_route.incoming_port.no_port_match](data-sources--http_loadbalancer--reference--group-025.md#canonical-4c260984063466891c51238e0009c316b3a4ab555ccfe783785f475169dd3814)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4c260984063466891c51238e0009c316b3a4ab555ccfe783785f475169dd3814"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a38b51fcec06271ac4c0d191a5f74ab4f6c48bbba03860cd3e5238dcecaefc5"></a>

## routes.simple_route.incoming_port.no_port_match — routes.simple_route.incoming_port.no_port_match / 1cb78b507d7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.incoming_port](data-sources--http_loadbalancer--reference--group-025.md#canonical-29cba187b019955171440e9188dcedbe57464be112234f015041aea135f6c35f)
- routes.simple_route.incoming_port.no_port_match

<a id="canonical-ff3912223ef33747e173df204c7a2f663bde5028d3a41b00b9fd86d2382134c2"></a>

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

<a id="canonical-f05fbb1acdcedd4a37343a88e96ed92e15895b91501314cc288fac01c0b6d387"></a>

## Direct properties — routes.simple_route.incoming_port.no_port_match / 1cb78b507d7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5cdbe6d5a9b7ef2ac1464bf4e4ebef30d2f93f6d532749f31495cc3d2fd07e5"></a>

## Next pages — routes.simple_route.incoming_port.no_port_match / 1cb78b507d7d / 4

- [routes.simple_route.incoming_port](data-sources--http_loadbalancer--reference--group-025.md#canonical-29cba187b019955171440e9188dcedbe57464be112234f015041aea135f6c35f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c4f42ec629e666093af8dfed5ace558f729c49003cae9dbf5f3e5e339306d83"></a>

## routes.simple_route.origin_pools — routes.simple_route.origin_pools / e7b4caf7c05c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.origin_pools

<a id="canonical-f75fd18134535e3877248a688b97ea0fc83a780490d166ee62f6d71efdc7e2e3"></a>

Type: `"list"`. Computed.

Origin Pools. Origin Pools for this route.

Upstream description:

Origin Pools for this route.

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

<a id="canonical-a6784ba4696e7eec0551645cff1276956eff36ff152339c43b05971c4c5cfb8b"></a>

## Direct properties — routes.simple_route.origin_pools / e7b4caf7c05c / 3

- [cluster](data-sources--http_loadbalancer--reference--group-025.md#canonical-eca26043981394278df1909224b60a40cdcb2779b40346839e786844b5241e5a): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-025.md#canonical-9a12004c1de886cc358951e3a8d5cb8b0b3bd3fe2d2f4141e65b1a8e2e8a8f63): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-025.md#canonical-44f7f2c84c07b3be66d4541f91690aa425a1b3f4ddc567410d165821f591ec33): complete subsection reference.

<a id="canonical-ca150c63dfece531027e267035395db92e2179256e46c015ee886e5319549f54"></a>

<a id="canonical-35ea45e1c19d7781dac23066caa4120ad91415309b21edd2a07099b61e0bee8e"></a>

## priority property — routes.simple_route.origin_pools / e7b4caf7c05c / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2ab1a137aad071fd20a47f8aa5094681864a1e0367e96e7f423290724beb773f"></a>

<a id="canonical-582a7a996e4740ac99c66172e9e5b40d66ba51914b987d2393afaffb38f91778"></a>

## weight property — routes.simple_route.origin_pools / e7b4caf7c05c / 5

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-f878e52ba6dd5ea7d5e8c90ed7186f5a47d8ea1c413168ce1b9eecacacc95742"></a>

## Next pages — routes.simple_route.origin_pools / e7b4caf7c05c / 6

- [routes.simple_route.origin_pools.cluster](data-sources--http_loadbalancer--reference--group-025.md#canonical-eca26043981394278df1909224b60a40cdcb2779b40346839e786844b5241e5a)
- [routes.simple_route.origin_pools.endpoint_subsets](data-sources--http_loadbalancer--reference--group-025.md#canonical-9a12004c1de886cc358951e3a8d5cb8b0b3bd3fe2d2f4141e65b1a8e2e8a8f63)
- [routes.simple_route.origin_pools.pool](data-sources--http_loadbalancer--reference--group-025.md#canonical-44f7f2c84c07b3be66d4541f91690aa425a1b3f4ddc567410d165821f591ec33)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-eca26043981394278df1909224b60a40cdcb2779b40346839e786844b5241e5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63be196c1b30b74b374910c7ccef95a30b55fbe756c49e00593593d44c5fd800"></a>

## routes.simple_route.origin_pools.cluster — routes.simple_route.origin_pools.cluster / b113ad3eecff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- routes.simple_route.origin_pools.cluster

<a id="canonical-3dec77ef5d99225c31f9fc2fb4830dda9d30f3857c97d0b51e8476f7e087c7fc"></a>

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

<a id="canonical-6c26ecaac32d8aa092cfb9dbc99b9945df63b356ca0c44e59e7c3cb31edc7fb7"></a>

## Direct properties — routes.simple_route.origin_pools.cluster / b113ad3eecff / 3

<a id="canonical-2556895fc09be5a6311a293fa7e1f052ab3efabba545c494f764913160b541b9"></a>

<a id="canonical-a962d691cf75e33c5750e97f0815a3534813ad3d377579a702a566b90ecd5612"></a>

## name property — routes.simple_route.origin_pools.cluster / b113ad3eecff / 4

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

<a id="canonical-a8a3dbe5f6ff57b9f63e77955ed72423f1d2c55c81312aa145e953b0b7d90154"></a>

<a id="canonical-2da5d14bbb805b6f22039939c86cec5dd5a61b89e1763dd60043a30a089ea110"></a>

## namespace property — routes.simple_route.origin_pools.cluster / b113ad3eecff / 5

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

<a id="canonical-cc92ec3cc4edf5a5b47e3862aa0b26d9866bd32aeb4d3ac4e2ca3a62f059ae20"></a>

<a id="canonical-22f2cc1448e5cfaa9de2764c70202b2742223d1b29c0bd062cb529b560a30bbd"></a>

## tenant property — routes.simple_route.origin_pools.cluster / b113ad3eecff / 6

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

<a id="canonical-d6fa713f30268a0d4f474baaa87c53bf38a7b6436adaf7d30b79ab7bcd9b5ca6"></a>

## Next pages — routes.simple_route.origin_pools.cluster / b113ad3eecff / 7

- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9a12004c1de886cc358951e3a8d5cb8b0b3bd3fe2d2f4141e65b1a8e2e8a8f63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73b2e5eec26b84ba300c8426b68e99d35323de89f83e46bd3128723ee07255f0"></a>

## routes.simple_route.origin_pools.endpoint_subsets — routes.simple_route.origin_pools.endpoint_subsets / ed4183d8cae4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- routes.simple_route.origin_pools.endpoint_subsets

<a id="canonical-c51dc6ee93832f0c8159716537fcae7ae88fc55be38729a0c026d7f317b441e8"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-b94e187ea497546a2160828b2d6eff12bf365b4cb3c72ecd953894c85925c5c7"></a>

## Direct properties — routes.simple_route.origin_pools.endpoint_subsets / ed4183d8cae4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-094b75b976c21e2e292864d1c1b36272534536fcb6b96c0d72396eaf73a5aeb1"></a>

## Next pages — routes.simple_route.origin_pools.endpoint_subsets / ed4183d8cae4 / 4

- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-44f7f2c84c07b3be66d4541f91690aa425a1b3f4ddc567410d165821f591ec33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a063975a775daa380b0651fbada8be9b57ae0182936add9edd2b233c04af5901"></a>

## routes.simple_route.origin_pools.pool — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- routes.simple_route.origin_pools.pool

<a id="canonical-f6c2015454f5a8943cbe63b36efa52c3df49e00b0cbf00ba4d5119ddb8e2a503"></a>

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

<a id="canonical-214664c3a0281640614e5c4192c88417432a896b1efc28e6f578315f0cd613d3"></a>

## Direct properties — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 3

<a id="canonical-e5a438bc9ab00f772734b91e05dbf7810d3d521f8be3be2e5c832f28db5e94db"></a>

<a id="canonical-e495f658c736d666eedc1a59ef92f976a8e49df440ba70105dfcf8536b6815c5"></a>

## name property — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 4

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

<a id="canonical-d9c8e29c872ee4172f048faa549ec6ea712d07f7b1857b7382ba600f05c14974"></a>

<a id="canonical-90e155f0a573ccf08281dd8e7cba8d92681ba2349488feacb03fe6f7b95a7d7e"></a>

## namespace property — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 5

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

<a id="canonical-92a16697b023dc3319a600737a9845b9b5dfcf074c59765989c7cdf9d48896c8"></a>

<a id="canonical-5ed7dfc38b3f73200f7172248a7eacb8117bbe5209f15da3fba7ef5eafb6f1fa"></a>

## tenant property — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 6

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

<a id="canonical-3c3788e934c80cf681f3e3d281251c33437e1a755ead11449cdaddcd1c36cd59"></a>

## Next pages — routes.simple_route.origin_pools.pool / c9df5ed2b571 / 7

- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-740eb0b78dceeed12b23d9b6594c906b796d3474248f1238748fe40172dadfc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa191408032ef227c97b11e9b1ceb4dd00a0c1dee9d1cbc009064547636efd8a"></a>

## routes.simple_route.path — routes.simple_route.path / 8087f4fa1527 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.path

<a id="canonical-7f31603edf0353a29fbadcfc45f637aac4157aab71b3666928defeb647a03320"></a>

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

<a id="canonical-a8e76ec5df08331ce1961b830068684cbfeabbe3bcaee40aafb3783babedd483"></a>

## Direct properties — routes.simple_route.path / 8087f4fa1527 / 3

<a id="canonical-8c6f890115d441c50ad3fe689332637201be6185cf7d719f5d5203a8a39f16d4"></a>

<a id="canonical-ec4a1362b778fb1f41d9807110019db4de24c487c1748dccb23982ca1359d0d0"></a>

## path property — routes.simple_route.path / 8087f4fa1527 / 4

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

<a id="canonical-b2dc1167eaa93b0744a61f42ea444268944b4970a985346f39a87ca58c00a9f8"></a>

<a id="canonical-58cf83cd73474228949b6f2395caab3a1ac2e71609862b2f579379c5b16077ca"></a>

## prefix property — routes.simple_route.path / 8087f4fa1527 / 5

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

<a id="canonical-032d3122a25aafb628b66248f85c38dda397dba3c0466173b840488c4601660c"></a>

<a id="canonical-cfad3a3da50c20047d2385cceb549f0882358fa1ffa3206601745f509c59cbb6"></a>

## regex property — routes.simple_route.path / 8087f4fa1527 / 6

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

<a id="canonical-dd8f7d1bb1920720600fc91388525a73c7a9e7a906631d96600002318018d28f"></a>

## Next pages — routes.simple_route.path / 8087f4fa1527 / 7

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94bcdbc1b88c1eecde199b54c2086206f96a5f17d045f2c40680b99526caac10"></a>

## routes.simple_route.query_params — routes.simple_route.query_params / b15f2a1443cd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.query_params

<a id="canonical-f3370d3acd6242986a60a307e1b6673a6248ec42cd90bcfa7e1427ed6defd2bb"></a>

Type: `"single"`. Computed.

Handling of incoming query parameters in simple route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

<a id="canonical-c9b1880ac4d39448d107014bd1cd2bc89917d63f1a028f16656cdeb0af0d725a"></a>

## Direct properties — routes.simple_route.query_params / b15f2a1443cd / 3

- [remove_all_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-8cdb33478294418767dd816dc830d23c99f402196e132a64d8f43cbb4b870a5c): complete subsection reference.

<a id="canonical-67e5013d520c3bb53026e5234526535dccb44c49afcfdf867eb0c1a73e2350b8"></a>

<a id="canonical-b5727c28052bdbb46370003e9dca6befb25e02aecb57a94cb3d9a48da03837bc"></a>

## replace_params property — routes.simple_route.query_params / b15f2a1443cd / 4

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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

- [retain_all_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-9ddccf72fbaffb32b71046a94edf5a44831fedf8ef8c1ec96ceab9b6f25cef4a): complete subsection reference.

<a id="canonical-acb686b41f93940d3b1fc04566c1fc9ea4972881795d193c022111ac090b0089"></a>

## Next pages — routes.simple_route.query_params / b15f2a1443cd / 5

- [routes.simple_route.query_params.remove_all_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-8cdb33478294418767dd816dc830d23c99f402196e132a64d8f43cbb4b870a5c)
- [routes.simple_route.query_params.retain_all_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-9ddccf72fbaffb32b71046a94edf5a44831fedf8ef8c1ec96ceab9b6f25cef4a)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8cdb33478294418767dd816dc830d23c99f402196e132a64d8f43cbb4b870a5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9cb68cc2246fb5986af10f1c551cdf2bc4a65d136ea2fb8e5165f35ff1144fd"></a>

## routes.simple_route.query_params.remove_all_params — routes.simple_route.query_params.remove_all_params / 03d57d3c72f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38)
- routes.simple_route.query_params.remove_all_params

<a id="canonical-b2f0cc2e60e1745f73d62521c7472cb426b121f2655b8adc4edba334635b2078"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-5f9b3d1698fbd9757a5ca9ee2220812a90ad27c82da420fb29a440320ebde4f3"></a>

## Direct properties — routes.simple_route.query_params.remove_all_params / 03d57d3c72f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa3bf8b3c383ceb8052b64e5d92223f81afac17120ddafc0450284b8f67c9ac4"></a>

## Next pages — routes.simple_route.query_params.remove_all_params / 03d57d3c72f4 / 4

- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9ddccf72fbaffb32b71046a94edf5a44831fedf8ef8c1ec96ceab9b6f25cef4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5ec99ba397a5deb8477bb6988fab2dea195640ca2957c2330b44fea17a65459"></a>

## routes.simple_route.query_params.retain_all_params — routes.simple_route.query_params.retain_all_params / 07496bc5402c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38)
- routes.simple_route.query_params.retain_all_params

<a id="canonical-81be6fc96b85f2dcdafb21bc45749652475ddcd89bf56337ef9d93131a63a717"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-d4fb9f1883c7417d128d1d2f024f4c82a25e94ce86d5c9d81c074ce20063b22f"></a>

## Direct properties — routes.simple_route.query_params.retain_all_params / 07496bc5402c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c15bfdedb7cc4d2af148a45045d457c1aa233072456a74227cd06e4c8d0e223f"></a>

## Next pages — routes.simple_route.query_params.retain_all_params / 07496bc5402c / 4

- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88609ed97ec6ebb6a5d05013c69e51110c589ad9b0ef4cb5d753e4003a2fb8de"></a>

## sensitive_data_disclosure_rules — sensitive_data_disclosure_rules / b0a751b47245 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- sensitive_data_disclosure_rules

<a id="canonical-87928b419dcb5fee9943310ff64bb470b658024521ba5e670523e82bf3c2aa8a"></a>

Type: `"single"`. Computed.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

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

<a id="canonical-ca51abd686b4078b9bf8660a3a63d318f88cd7e74a953ec298aac1ceb6ffaeb3"></a>

## Direct properties — sensitive_data_disclosure_rules / b0a751b47245 / 3

- [sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039): complete subsection reference.

<a id="canonical-35dcd18773554129fa767b2ea78a1d520c695c780d7ec09d04ee8a4349bfa3e6"></a>

## Next pages — sensitive_data_disclosure_rules / b0a751b47245 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f877a156a6c351f40b40ae27618396030e34f5e83b762d558f90499dfce9270"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response — sensitive_data_disclosure_rules.sensitive_data_types_in_response / c2a7d6f7cd0e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="canonical-3e7a463165766bf3066f00b4d51a687e163bf0c08960dca7063e51cc4fa21b4c"></a>

Type: `"list"`. Computed.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-efe1c2f044d7e5ee3beca7724eacec92f58c3d8a3f1af11a3ab3dce5fa49a35e"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response / c2a7d6f7cd0e / 3

- [api_endpoint](data-sources--http_loadbalancer--reference--group-025.md#canonical-2b9568dee744376a6cfb0e507f149ed813bf7f99bfe6c482380eb289f70a3102): complete subsection reference.

- [body](data-sources--http_loadbalancer--reference--group-025.md#canonical-c41031d9a347f27ae95c4a87a11f18d172861ac7beb438b54f86e9091154d010): complete subsection reference.

- [mask](data-sources--http_loadbalancer--reference--group-025.md#canonical-94900bc9e2941f7fe2e6247386b952003a31fd0900af7aaf140643b6a921a4cc): complete subsection reference.

- [report](data-sources--http_loadbalancer--reference--group-025.md#canonical-1998a17548ed66ce01c1b3c7306bd9bd8370d23fe231bb9e0c8ae7ecd0e80911): complete subsection reference.

<a id="canonical-0728c4830ad20dfc7c46b70ee51272179f413238d08938f47aef15f39133f510"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response / c2a7d6f7cd0e / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](data-sources--http_loadbalancer--reference--group-025.md#canonical-2b9568dee744376a6cfb0e507f149ed813bf7f99bfe6c482380eb289f70a3102)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](data-sources--http_loadbalancer--reference--group-025.md#canonical-c41031d9a347f27ae95c4a87a11f18d172861ac7beb438b54f86e9091154d010)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](data-sources--http_loadbalancer--reference--group-025.md#canonical-94900bc9e2941f7fe2e6247386b952003a31fd0900af7aaf140643b6a921a4cc)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](data-sources--http_loadbalancer--reference--group-025.md#canonical-1998a17548ed66ce01c1b3c7306bd9bd8370d23fe231bb9e0c8ae7ecd0e80911)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2b9568dee744376a6cfb0e507f149ed813bf7f99bfe6c482380eb289f70a3102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-825658372a9260c81ad05a571ba6d0af0c971e1016d009f37af5933e37670378"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / cc92fb9a9a2c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

<a id="canonical-1cfd7c851f5c4eec47846f4d833f17dab73717d1d4b270db4010036b7b27e66d"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

<a id="canonical-5bce440db63d21a6c6b9f619d66e57dd93407d6b3e20388ee037fc7258b40d90"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / cc92fb9a9a2c / 3

<a id="canonical-252c167e02cdff71a171a5c1d10915bfe649ab30c5282718612bd7a07fe3d66b"></a>

<a id="canonical-bd13d150a7611139254792a4c9b4b399e7e20d9b9e6b08deaa9f51bebf4a6dd4"></a>

## methods property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / cc92fb9a9a2c / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a0aceaf7de1f8fb4d18d51220b6154a0d18a56bf71eb73f3539d1dcc230f2453"></a>

<a id="canonical-0aed9e7060cde99268785d5488a22413b9300e65daf102f100de460b77fdb737"></a>

## path property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / cc92fb9a9a2c / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1b273d43aa40cde3950d5e455998443ba914bec36f7ba0874729a57075468ae6"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / cc92fb9a9a2c / 6

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c41031d9a347f27ae95c4a87a11f18d172861ac7beb438b54f86e9091154d010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701b073238ebb30907da81641f6ee8b096d13a93e1bee20819bd4622c239de98"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.body — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / 43394d5babf2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="canonical-92b07ce056595cc2fb1722d6cd428e479b2068dac33f5a52e7b5e65bacd26d5e"></a>

Type: `"single"`. Computed.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Upstream description:

OPTIONS for HTTP Body Masking.

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

<a id="canonical-4c723729eb1137cd0754eb2f2125579e77d73cedc961a3ff7e1331c489c12d71"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / 43394d5babf2 / 3

<a id="canonical-5a7199e87155cfc7a3a201712c834fb719ffa98f2b860067ff09ff334871a279"></a>

<a id="canonical-920aeb1295a987b6437add0a7a3d5d1dce150387928c52928cd7af35ea3bcb23"></a>

## fields property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / 43394d5babf2 / 4

Type: `["list", "string"]`. Computed.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes.

Upstream description:

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

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
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cfd2e0bfc69453722240ab564f936dbb7038232dc7e4ebe010f79c520ec2e189"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / 43394d5babf2 / 5

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-94900bc9e2941f7fe2e6247386b952003a31fd0900af7aaf140643b6a921a4cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c660c64db475e39ae1e3b8f5d20cc19977e252b4b34e8fce97c99fe96c4539ff"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / fce46c86a8a7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

<a id="canonical-19c1058b668fe3667251a952140ea1941484e38ac9ab8ed5bfed6cb215193b26"></a>

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

<a id="canonical-0444cfd70d8e26cb472a532791be03f8785d83043c09b352f3a91ddf8d79f344"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / fce46c86a8a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff9437a403590c0df6d3b1b11e5c8e7cdcad513af9e19c333193dbdd136e7cd2"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / fce46c86a8a7 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1998a17548ed66ce01c1b3c7306bd9bd8370d23fe231bb9e0c8ae7ecd0e80911"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7dc873673f08868ce54e30feb1c4f53c3157a86d84703c814c7c057f5dccbfd"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.report — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / afc3ae0429a3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-025.md#canonical-b6c3a1417be308a40152934c38c279e30755355f2674475995819812f65a9671)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.report

<a id="canonical-1ccf971c1d9e9494d250d1f01bb75533a0b8ee115ab806eb13108ef2d891e82b"></a>

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

<a id="canonical-edb41b94c0e3a50bfab263560ad45f1da34d7dbc5d07ee5a028c16d4ef876266"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / afc3ae0429a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a376f0bed896dd2ab979630fccea51f322b5f5cc4ffbe39794ef0481b6e5579"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / afc3ae0429a3 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-025.md#canonical-a2d4dca1b0322d070a21281a8f69b6827775e598e96f4dcb97406bc25c939039)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-555e1fd9db4016e1051d09b85d6284a1b66657e93495db6fab8e93b1b137257d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b41c7b043ac5bf10a07f888ec1f6652d24d637f31e573c9b963e5ee3562bfa2"></a>

## sensitive_data_policy — sensitive_data_policy / 37bad8bbefb2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- sensitive_data_policy

<a id="canonical-6f1aff19cc28fb21bf8c1d5a1a8c4b73143011f8d76a167baae77b17eb1dbba2"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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

<a id="canonical-fcf96c53f98aa33d66826e30bf5c6936931e803ab66fa3b49f0d199d4255769e"></a>

## Direct properties — sensitive_data_policy / 37bad8bbefb2 / 3

- [sensitive_data_policy_ref](data-sources--http_loadbalancer--reference--group-025.md#canonical-c3846f9e660bfe6df97941b246f4535a09de655d4a77fdd3e5e28af0e32561c0): complete subsection reference.

<a id="canonical-f30be024dd3f6a509158182765b63223256b015f50b2551765b1f5dc1a295163"></a>

## Next pages — sensitive_data_policy / 37bad8bbefb2 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](data-sources--http_loadbalancer--reference--group-025.md#canonical-c3846f9e660bfe6df97941b246f4535a09de655d4a77fdd3e5e28af0e32561c0)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c3846f9e660bfe6df97941b246f4535a09de655d4a77fdd3e5e28af0e32561c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fdc0eb15f3dd7befaf6ca57a6794ccccb045e6b9f5f73f58510bcebad413389"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-025.md#canonical-555e1fd9db4016e1051d09b85d6284a1b66657e93495db6fab8e93b1b137257d)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-3e8f5a6d4e2059ceed39810427a90f90bde4068c8614b68e2125fbfea6058390"></a>

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

<a id="canonical-8331c04cab9e2937bbf6067f5a93e4dbe171f0cd74ef9f4b67ed341f92a870b4"></a>

## Direct properties — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 3

<a id="canonical-bca7b4256becca34aecc16cf9bfeaa7aef38f9c7c51906c3124f085ac9627565"></a>

<a id="canonical-ee054b5055095400e2b637114251e7e1101d2b5fa42bf1a76ef48c3e3316b73e"></a>

## name property — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 4

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

<a id="canonical-94e87abc64cdca32d92eb74dd808c0ad4943d52f66913b536ed0b1397b3115b2"></a>

<a id="canonical-c81ebc3bdab2186c7aba44b773fe95c7122f5105cd84cadf91bfea4c7acda5c6"></a>

## namespace property — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 5

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

<a id="canonical-d727d014154f309be877d61ac366e94102462508f8e687c7bd02edbbfc1b61f4"></a>

<a id="canonical-1ba3dc2c2ac1c5e9826ae1a5ec64405bf3353840eb0a2be62b1ee3dd6bbc8ac2"></a>

## tenant property — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 6

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

<a id="canonical-6136b7feb26f686b5bd77926e79f06fe8180f4ae1ce2bf8705ce0b354db8f15e"></a>

## Next pages — sensitive_data_policy.sensitive_data_policy_ref / 2eb74fc3a65d / 7

- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-025.md#canonical-555e1fd9db4016e1051d09b85d6284a1b66657e93495db6fab8e93b1b137257d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ce1121f145bd5fca71bb8b705f4d8310b91f390aa749824440eb6b34acfb0fde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78308c83b6786b99c5ab0f47ad7ad5777383245d90f841ee84f1d9f19ba3303f"></a>

## service_policies_from_namespace — service_policies_from_namespace / 071891c3af5a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- service_policies_from_namespace

<a id="canonical-a4e857ce5c82c665d299b139f646572ffb4441c242172fe9f2c2c5d658a3f7c7"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-a76917e048446e44cb6143a2eb57b2a6ecb694c595a74be3b9bddc25d2ddcd37"></a>

## Direct properties — service_policies_from_namespace / 071891c3af5a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f57d81244198d4afb8ab97a9b4dc79a648ac0330e5d01a41a22711e6791551c7"></a>

## Next pages — service_policies_from_namespace / 071891c3af5a / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43b1e67872c793a0ba37e90404327f66f7e9a8b0fdd5e5e94fc416aab2763e67"></a>

## single_lb_app — single_lb_app / f6997384a6a6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- single_lb_app

<a id="canonical-0d805775af5bb40f150457d8a9b2167a896dc325aa306e1b3bb62586f01860ee"></a>

Type: `"single"`. Computed.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

<a id="canonical-eef6826286ef9dae32e6702ae930c6f30ab355a2f183347d7b91cb5e4001e001"></a>

## Direct properties — single_lb_app / f6997384a6a6 / 3

- [disable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-6fb5c03dec7dad834ae0347cd74b07f993e738f167d80c43da88d86a4e197b01): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-025.md#canonical-b73e41554d5cbeaf48edd241ac42851ce5cf2d5ef850e922225884fb262d205c): complete subsection reference.

- [enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-025.md#canonical-507953ce2096f48d5567042afa8acd9f7068cb765a3c4a4797c23c382b99b3ea): complete subsection reference.

<a id="canonical-75c7e29262a4e65fc5c56bf28b2ae3324b44e9c86abda5be19a6228197159c41"></a>

## Next pages — single_lb_app / f6997384a6a6 / 4

- [single_lb_app.disable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-6fb5c03dec7dad834ae0347cd74b07f993e738f167d80c43da88d86a4e197b01)
- [single_lb_app.disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-025.md#canonical-b73e41554d5cbeaf48edd241ac42851ce5cf2d5ef850e922225884fb262d205c)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-025.md#canonical-507953ce2096f48d5567042afa8acd9f7068cb765a3c4a4797c23c382b99b3ea)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6fb5c03dec7dad834ae0347cd74b07f993e738f167d80c43da88d86a4e197b01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44879fa54f845999cee6a99fc0ec7f132f479eebcd753a73130e549042707215"></a>

## single_lb_app.disable_discovery — single_lb_app.disable_discovery / 523f4a47f032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- single_lb_app.disable_discovery

<a id="canonical-18ed0bb77ea70db3c332e4511afa0ebc7c731742e6ac70b81755a44b23045fdf"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable discovery.

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

<a id="canonical-f3b8933c1ed5d16644b6f8e1db0efc9464186a467af1d17bd88a83b8983ffd78"></a>

## Direct properties — single_lb_app.disable_discovery / 523f4a47f032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b2b4241d853d49cfdcaa610c345dbad3ab53d6350084fdb044da26637156c1e"></a>

## Next pages — single_lb_app.disable_discovery / 523f4a47f032 / 4

- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b73e41554d5cbeaf48edd241ac42851ce5cf2d5ef850e922225884fb262d205c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ded2812ae1997323ed884b73ea7c1f9918d7d77db161e787dd5f61f6be07a0bd"></a>

## single_lb_app.disable_malicious_user_detection — single_lb_app.disable_malicious_user_detection / f80ef1a3dda2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- single_lb_app.disable_malicious_user_detection

<a id="canonical-bdd8cabc7cf2c44d75beac6430b21fb9e5ef333557dae9f61a5fd3a39cc22782"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable malicious user detection.

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

<a id="canonical-f8abda4b193f6d061398aebd1d4c0a3b14851288a5c56100d2a303a358b95e4f"></a>

## Direct properties — single_lb_app.disable_malicious_user_detection / f80ef1a3dda2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa9ada7fb679a837672ac1bbedb3de637ce9f43fec2651b980b6d54062218285"></a>

## Next pages — single_lb_app.disable_malicious_user_detection / f80ef1a3dda2 / 4

- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ad08bcb3bca844b692b0f3217615a033dfb64b4c59d7bee0d194459996c4651"></a>

## single_lb_app.enable_discovery — single_lb_app.enable_discovery / bde312260d30 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- single_lb_app.enable_discovery

<a id="canonical-4725c0192f7b7a0d5feb0b414023a38b3eab034cdf0f73fab0639dfeb6fe46d2"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-e31e173e1768a04db3d8552d721887e99079992229f53bd3fefea8f4185425ed"></a>

## Direct properties — single_lb_app.enable_discovery / bde312260d30 / 3

- [api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43): complete subsection reference.

- [custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-c99d1494f501f3ff7604ea8bd163174ab1788b38a448a7d8159e626d2f2e7648): complete subsection reference.

- [default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-b53f9b55449615aaf681f61c9ef59638e281815f53791972627009e6d764e346): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-025.md#canonical-8896bb5f18dfa45f13a197bc96243c77f4c7aec854b13f934ea1a72424ddf8bb): complete subsection reference.

- [discovered_api_settings](data-sources--http_loadbalancer--reference--group-025.md#canonical-6777310fbcf668d673401c38f3540b70ccf09fae31c1583989bed04fc1f9e0c5): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-025.md#canonical-74b5d9db3adb30e3c2c4f4bd453643eeb0185c6885bad2ea15d00516f937edc9): complete subsection reference.

<a id="canonical-694f18c3e97f8f49db0d86ad372cee991bbe7dc7024c7645cee96a4598b2cc56"></a>

## Next pages — single_lb_app.enable_discovery / bde312260d30 / 4

- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-c99d1494f501f3ff7604ea8bd163174ab1788b38a448a7d8159e626d2f2e7648)
- [single_lb_app.enable_discovery.default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-b53f9b55449615aaf681f61c9ef59638e281815f53791972627009e6d764e346)
- [single_lb_app.enable_discovery.disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-025.md#canonical-8896bb5f18dfa45f13a197bc96243c77f4c7aec854b13f934ea1a72424ddf8bb)
- [single_lb_app.enable_discovery.discovered_api_settings](data-sources--http_loadbalancer--reference--group-025.md#canonical-6777310fbcf668d673401c38f3540b70ccf09fae31c1583989bed04fc1f9e0c5)
- [single_lb_app.enable_discovery.enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-025.md#canonical-74b5d9db3adb30e3c2c4f4bd453643eeb0185c6885bad2ea15d00516f937edc9)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f920a5fa7b9f1a9c1a6974a987f99973e21d98f66f4838526e37d95ee7d85565"></a>

## single_lb_app.enable_discovery.api_crawler — single_lb_app.enable_discovery.api_crawler / da5702053dcc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.api_crawler

<a id="canonical-1145f9d01134667e3fbb7933b5ca6812dd0840bc6dbcb53019b1434780d1ba29"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-e98c0827a27d3f998c7db19f896e5fa2c184ff4a7780fd222cdb6cb8b0bb272e"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler / da5702053dcc / 3

- [api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65): complete subsection reference.

- [disable_api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-8be09332ec35fd60c40a7866fbf3016b10d6feba121b82c529bfbe72db2ca774): complete subsection reference.

<a id="canonical-e8a732aa7ee5ec3817f4031548bb24f0cd119972819b8e32ca068362d2ae5cb9"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler / da5702053dcc / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [single_lb_app.enable_discovery.api_crawler.disable_api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-8be09332ec35fd60c40a7866fbf3016b10d6feba121b82c529bfbe72db2ca774)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53cc26a92529f567588946d6b32dda7bd8f2339854e0544e6b6a3db46559dd5d"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config — single_lb_app.enable_discovery.api_crawler.api_crawler_config / d096d4045222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

<a id="canonical-7edb6700639c9940fef143acb6d12afa7ccc33fc08887ef03e8efdab73d2ccc0"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

<a id="canonical-0d755a54d7b80db5be638476fc46123f58ed90e762d061d8c0fe0b15f8fb6f24"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config / d096d4045222 / 3

- [domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537): complete subsection reference.

<a id="canonical-78198b795d8f054dbf34714bd00340d4715ae8f3d9c9f4884b9b82b80e50dea7"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config / d096d4045222 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1bc5c4f609e221bc49bfb6c3809aeee81ed907bfa98a97c519de33ff34857a4"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / cd56ad73a6ff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-dd91530676fb89c1a7a5ded9279f7142f82c279c4671227172bb0ae23842c77c"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-168b9183f198aaafc029f2c8a53ffcd86d1be08d7389ff36fe4fe39bc8c14520"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / cd56ad73a6ff / 3

<a id="canonical-acd3b6079b7a9a6a9d7cbb4a81b41a75ab5f7cd1ec15ede411fb99da0e0221c9"></a>

<a id="canonical-065a45a53e4ecb561d7b0b717ae7280aad18f0994c7f657c8d5fd93ce17cb610"></a>

## domain property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / cd56ad73a6ff / 4

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c): complete subsection reference.

<a id="canonical-66202c83b8b0055d004898c0c139f4eeaca63b02f66f58c6901ae2640df28046"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / cd56ad73a6ff / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19e1f5e8c7acf79b73d4c2ca44bb98e6f010f8aaab1cb429d95ac65e3057f489"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 29bfe64c33ad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-ada5dccef8eb22d0a40e09ad4b12c0f0219642bb2bc9913de6581e5752eb760b"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-155ebf56cea15d13320c56fbe21e7f58fdad03926fa996c0d8b847e008cfacaa"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 29bfe64c33ad / 3

- [password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4): complete subsection reference.

<a id="canonical-0699cf55288f3ceca29739eaae727422b40abeac27d1c4e9590f2c489b6d230a"></a>

<a id="canonical-4c5e070129fa34f628f56dc2a71b3b5093cfea6787bc2587d4c166e1d33c8ce4"></a>

## user property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 29bfe64c33ad / 4

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f8da2255b7cd320153ab6c4590de49961b859e524107dceda7562d2379cc437e"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 29bfe64c33ad / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68cef28585f9865f43cfd993fdf816f72964dd9551ffe670acb10f612f94aeb7"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 90f632f45417 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-4da9b48dc31a5a97c74fe26acf9ead40223178a01427b85e66f352092143cee6"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-53e42078993f477f2024285851b258c30ee191f191fde0ac5f5438d01f18138c"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 90f632f45417 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-025.md#canonical-10e53ab1c5b775be216907dbd755fbde21b75d5204afeb006285df1b52182690): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-025.md#canonical-c6f06b5af5bf78ed30c4d7e5c1f0108d934a18864639c2ee99eaf069fa0766bd): complete subsection reference.

<a id="canonical-ec018137c658279de869156f6479ad3ad07b5a794ce85bb02eb86ace40d843f5"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 90f632f45417 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-025.md#canonical-10e53ab1c5b775be216907dbd755fbde21b75d5204afeb006285df1b52182690)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](data-sources--http_loadbalancer--reference--group-025.md#canonical-c6f06b5af5bf78ed30c4d7e5c1f0108d934a18864639c2ee99eaf069fa0766bd)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-10e53ab1c5b775be216907dbd755fbde21b75d5204afeb006285df1b52182690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e53581cef79a97f6e6842f7130478a4c79589a5dc10cd141e42bfbbf4d63d2a"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-14e514247dd465f52a3837acf494a7de70677e93795071666d57bf7cf3e5c4b1"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-68efad826289961e9e352a585228b5c148c1c42f27b5e25da83ef00783248684"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 3

<a id="canonical-b20b322bd88e8aafcc6a5c25265c4f59270b56d05f84fb60776790912abeea9d"></a>

<a id="canonical-19c8f629f9c143991364193a538606e334e0a669fef6145b8d629dfd8f041440"></a>

## decryption_provider property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0f60782822ab292f35ba442bd0e26f99c25b207d7dbc643bed16d971ea2c971c"></a>

<a id="canonical-7ca94a3496420b84e7f481ad5a06dab80e689e9d2afafdda11cadb02f3dd20f8"></a>

## location property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-f4e61c21d4b4e70a82344ee0dc0996716df65b4c38b30fe84f30b3d15f9d4485"></a>

<a id="canonical-2f3e92756f3e8add1adf22ea5b7bdcb133f68317673ecf36ef02d149ae2cc635"></a>

## store_provider property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-b44a2d5d1a4d38ac80a06f20f1e3cc2c6dc1ea2fc1658922cd793440005d6f41"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 3936db41eed2 / 7

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c6f06b5af5bf78ed30c4d7e5c1f0108d934a18864639c2ee99eaf069fa0766bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c56a89defbcf89ea3f1502f0251246aec7ef9f91dbd986cf2646ebb99bf4df5b"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 557fa0fe4985 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f63adaab6aff3664e39e327f1c16eeeb21b6c4ad4cb42199deefe84a6de6b65)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-025.md#canonical-861d170fc74d67f289258df426603e12d89afe406c3558e68f6eab6372b5b537)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-025.md#canonical-1f51b0866ced5132d89f8029302067cbb6eb427c54f93b50a548abc26de85f8c)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-e9639e1c4862914230ffcb8af4291d8342f370940a89ad738136f9264167f093"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-fa941870a14d164a0091b11c114768213495ad95cf0cdcb747d9c6d828a2ea10"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 557fa0fe4985 / 3

<a id="canonical-8977cfa422cd67c73c58a9f343b4e2f6136158ae18c5b20c66205ff37f9386cc"></a>

<a id="canonical-fb788967f7b3976e9d9041118aa6f098af53a2c135efb682e5e2d82e08a4b56c"></a>

## provider_ref property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 557fa0fe4985 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7fb468b9475f9bccf3ae7ee22f33a50facd154d1eb982d6d84ab8d7efb428d5f"></a>

<a id="canonical-06238752318ff6daeeecb2c744d13f01df55dac13e420e303b3d69aec79a322c"></a>

## url property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 557fa0fe4985 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3c50beb5be0be8ccf8ac15140c61f7360b1d183b2b7bfb55f57db5b45ac04063"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 557fa0fe4985 / 6

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-025.md#canonical-dc05a8a6f1de9eaabe4181bcbc410f5caeafb09bcc1a3ab5dbdf06dbc4a085d4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8be09332ec35fd60c40a7866fbf3016b10d6feba121b82c529bfbe72db2ca774"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c22e5be809c32fde71ca851baaf8270080bafa7a32714bfdcaa57b1dd0f6c71e"></a>

## single_lb_app.enable_discovery.api_crawler.disable_api_crawler — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 98e66ccd35cc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- single_lb_app.enable_discovery.api_crawler.disable_api_crawler

<a id="canonical-8395bf9af75f56ceb0b27b615ecfee385b0c9fff30ab944dccf2bc3f73f58860"></a>

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

<a id="canonical-af9195af4234fa888399bb3e7f016cbeaa9632908509c27bc84716e690f57c30"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 98e66ccd35cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c346546df367d55d31e19ca34a65e39d24e3320002a60504f8ecd7da69ac474"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 98e66ccd35cc / 4

- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-025.md#canonical-5ac1e6a357b5d4600983189e8c95c060d044419158cc12dbeb99906145ea135f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8cb677dab113482b60fb078db397d25ca760e3546282e218a8068d3f15b0344"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan — single_lb_app.enable_discovery.api_discovery_from_code_scan / 907fd6a858c4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="canonical-8e23907c51aa2c318698259dc4f24d7b92e8ca807aa9ec86d3e1040f85ebc4af"></a>

Type: `"single"`. Computed.

Select Code Base and Repositories.

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

<a id="canonical-e3fea9e7f064fd5228878ac7a670fdf736fd4517375944d8015d54a558480664"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan / 907fd6a858c4 / 3

- [code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb): complete subsection reference.

<a id="canonical-e6d261e6cbbe10021407defa766eb9f9eefcd48c8e08fb2701d164720381fe9e"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan / 907fd6a858c4 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8356ac878b4b9e28cf1373d66cd1233f9ab60a7577fd4bb29832fc4157c80f2"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 54f9b385ba68 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-dc9a97a5729b69f352b966f00606c46fc9594e242246b92ec1b81d6f1aba2a20"></a>

Type: `"list"`. Computed.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fdaf1aae8fa3a52b3b720923c7e15128c400a65bc9043504f8e29cba4912e8d2"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 54f9b385ba68 / 3

- [all_repos](data-sources--http_loadbalancer--reference--group-025.md#canonical-3461f4428724625dfcc813ab1b5b71a45b6add09fd177cecfd50142d81da9708): complete subsection reference.

- [code_base_integration](data-sources--http_loadbalancer--reference--group-025.md#canonical-04a8702a19b8e34cf6f4c331fbc257b9232738750ab8a38f0fa17805ce692774): complete subsection reference.

- [selected_repos](data-sources--http_loadbalancer--reference--group-025.md#canonical-bd075d8474343bded439961cdbd3952d8a95ef38d67800b6f3226855f54781e6): complete subsection reference.

<a id="canonical-8183a35e62a63eedc4755c6c34aa8839a333a9d9760db86b01142400e6b3dfee"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 54f9b385ba68 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](data-sources--http_loadbalancer--reference--group-025.md#canonical-3461f4428724625dfcc813ab1b5b71a45b6add09fd177cecfd50142d81da9708)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](data-sources--http_loadbalancer--reference--group-025.md#canonical-04a8702a19b8e34cf6f4c331fbc257b9232738750ab8a38f0fa17805ce692774)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](data-sources--http_loadbalancer--reference--group-025.md#canonical-bd075d8474343bded439961cdbd3952d8a95ef38d67800b6f3226855f54781e6)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3461f4428724625dfcc813ab1b5b71a45b6add09fd177cecfd50142d81da9708"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06c8584a78c44cf40f8fa3ddfdfe4554f336322a787837ba77077d1314c68bb1"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 77dcbedcb6d8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-e352b9d149fef8b8b02432d3c449767e76eada5af5f04c4cb9b8ef1dbae2eabe"></a>

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

<a id="canonical-c0942000ed13360b4a1bbb67510dce5bbd8d02a347aa57063ccbda04f3141d3c"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 77dcbedcb6d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19ed01a8b32b54f2b76de165f24564da00e1078492d24427cc028acea66c3e67"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 77dcbedcb6d8 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-04a8702a19b8e34cf6f4c331fbc257b9232738750ab8a38f0fa17805ce692774"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fef0406f355d0123a22e856a6c18b9137069b1c86c2bcaf8c68975213d26883d"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-fc11b2440c290ec885dee48474c6c1ba38beb213697f983f10a7b88d8759d624"></a>

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

<a id="canonical-2a4a6b89fa0b33f748ef407387ee078e6cc9c6e5c550627666227a697853acce"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 3

<a id="canonical-db08296b1cb7ba9591f0c7908cf91d34b540bb7ae526e41dfd32d737fbf7c37a"></a>

<a id="canonical-f97d40910ffad20dca1d152ea18487ea5ba161452e72b2d1435146b580cff183"></a>

## name property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 4

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

<a id="canonical-b64915d12da3f0c6edf802b48aff9668aca5feedc621beeb2b957b277af284fb"></a>

<a id="canonical-73f3d3d744f1ba1575eabe2391745056145fa8b12ab005e93e0f5c1215476772"></a>

## namespace property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 5

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

<a id="canonical-143ff192ca70bb13975d6d1381244638abf999b6255a7b9013fc987403c07e41"></a>

<a id="canonical-6354a93e1b8eb12f67d76f3469e1d4213c08bc1eb667813b12a72393f1c0bbad"></a>

## tenant property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 6

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

<a id="canonical-ced436050079049b8fc0f1b42a084338abc8fd152660a6905a8990c6daabc42d"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 5fb902f95289 / 7

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bd075d8474343bded439961cdbd3952d8a95ef38d67800b6f3226855f54781e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09d77ae3509e49150fb6c4de0e0021cac4efffcd164438347b1e2e79da2064d3"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / a6d5ca627f36 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-025.md#canonical-3a8df5a1b605d299aa819f5418b2e1575168302a5d62e591f50039c666c14a43)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-04c4af93c7445d12c942665ffc56eb2400c9621a59e24d13f62e096a0d2e64db"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-32b8abd6770ba556c39e2a10ef1c19dfd90d4e790701323916dfd38ec47d32d7"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / a6d5ca627f36 / 3

<a id="canonical-206ca0efd34dd14a3dd2dc0768b69f49bae3be6e038731e635979da936a2146e"></a>

<a id="canonical-cefea1594b83ffc79b8a2dc7ce91e5b32b50f38ed67512a96ff3fee4ead41554"></a>

## api_code_repo property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / a6d5ca627f36 / 4

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-79f5ffdc71ba222b991904768987621e6af396a618468ebbd0335219d8815297"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / a6d5ca627f36 / 5

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-025.md#canonical-1d12acd98373ca58273f720ad704e1e88d59052aa839c5ce8d28413163fb5feb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c99d1494f501f3ff7604ea8bd163174ab1788b38a448a7d8159e626d2f2e7648"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-330634e50335ddedcfa990eaaf1f1c27116288aff04f8a866a7cc6422a9f9879"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery — single_lb_app.enable_discovery.custom_api_auth_discovery / 849f29199e20 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="canonical-c88fdc6b73371fec7017e1ed6b35cfb671de9229d54c17e7204a1f04e62db700"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

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

<a id="canonical-af3ee05c02c46976e2cd5ba56f7e6dd1a0f2bbfc33cd78dca210f9ab4bdc86dd"></a>

## Direct properties — single_lb_app.enable_discovery.custom_api_auth_discovery / 849f29199e20 / 3

- [api_discovery_ref](data-sources--http_loadbalancer--reference--group-025.md#canonical-aa5a67097b2d656bf7b250c30fcd037583e31dd337f37ed5e46af853f0bd0a33): complete subsection reference.

<a id="canonical-c5b48da0f88ee387a6e2706a122fae3adc0375f7ea3706fb95d15570335daa59"></a>

## Next pages — single_lb_app.enable_discovery.custom_api_auth_discovery / 849f29199e20 / 4

- [single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref](data-sources--http_loadbalancer--reference--group-025.md#canonical-aa5a67097b2d656bf7b250c30fcd037583e31dd337f37ed5e46af853f0bd0a33)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-aa5a67097b2d656bf7b250c30fcd037583e31dd337f37ed5e46af853f0bd0a33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8a8114ab99c52022efdfa194b1d8a11ea4ce8dafe7735d41e649c4c4728177f"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-c99d1494f501f3ff7604ea8bd163174ab1788b38a448a7d8159e626d2f2e7648)
- single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-d288608dc1a3fd9c842a1d45b2a2576bfc1c6e019ceda04b0a699d843fd3ce78"></a>

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

<a id="canonical-dde90d996c9ab38a9ddee0da7875692f455d1096dddf2e64e90b9e64d9e2b406"></a>

## Direct properties — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 3

<a id="canonical-0c4da6859442a15a00226029b362af2bfc1ed2b0ff934a6e22536153300fff09"></a>

<a id="canonical-33db9add44a0424fbe1c8725e410a640c24c5e08be249d70487666bcd3b4498d"></a>

## name property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 4

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

<a id="canonical-7e03a47ef200e313b9b68bf53478fbd947af943d29686394cb18de5720935a6c"></a>

<a id="canonical-d694997c5b08d370ce526d73a27c4966b30762cd9ee65e104cf66c4c2544f48c"></a>

## namespace property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 5

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

<a id="canonical-0d423da501c99f97e0903e0e709b0b23e7fd18034c30f4d08cf8c5b13ff4a775"></a>

<a id="canonical-2f734154ad60cc9fb359d75d8d9320d2d4c04cfc2341843ef80dc5b98f164d7c"></a>

## tenant property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 6

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

<a id="canonical-10c2041ae826bcbb992eb5721c209d831f032d9fa240eddfbf2060c3e0df58f1"></a>

## Next pages — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / cb18c7bd59cb / 7

- [single_lb_app.enable_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-c99d1494f501f3ff7604ea8bd163174ab1788b38a448a7d8159e626d2f2e7648)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b53f9b55449615aaf681f61c9ef59638e281815f53791972627009e6d764e346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6309695613af06b6e77cb874a9803a5a4b9bb0eaffa56e4211a64f559462eb77"></a>

## single_lb_app.enable_discovery.default_api_auth_discovery — single_lb_app.enable_discovery.default_api_auth_discovery / 0090832144f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.default_api_auth_discovery

<a id="canonical-8ac54d99e2634b7911e8c0aec2d2f6d51d82894c54820a060c24740aa407dff1"></a>

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

<a id="canonical-e780c01d3423a6cffac3c506144c9d3b42c312d95f14b04f1d9b3b50f5289165"></a>

## Direct properties — single_lb_app.enable_discovery.default_api_auth_discovery / 0090832144f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfa59ae7bf551752033e345669aaf8e8cf3d8627fa40850297effdb45c65a2df"></a>

## Next pages — single_lb_app.enable_discovery.default_api_auth_discovery / 0090832144f7 / 4

- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8896bb5f18dfa45f13a197bc96243c77f4c7aec854b13f934ea1a72424ddf8bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7abde5b6b1d78076e7753de17ec8584316802d232a12f16773e543032133220"></a>

## single_lb_app.enable_discovery.disable_learn_from_redirect_traffic — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 5caa2b60c060 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.disable_learn_from_redirect_traffic

<a id="canonical-f163d4f38b8d82185822e72a4efddba1a31ccacfaeddd0a458e0580e2c348b00"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-628f27f798b239c69c23bc94b4dee723226597a432c491c151147dbc482ffdbf"></a>

## Direct properties — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 5caa2b60c060 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e5bebdf57bbe357f3bdaf72dbcf6241d27dc53ccf9f3b02636421ff41b10c7c"></a>

## Next pages — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 5caa2b60c060 / 4

- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6777310fbcf668d673401c38f3540b70ccf09fae31c1583989bed04fc1f9e0c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48312cc09e4fb1fcd0cf7d718c4167df2464aa251c78eb5ba553d2022bc09491"></a>

## single_lb_app.enable_discovery.discovered_api_settings — single_lb_app.enable_discovery.discovered_api_settings / ee39d3a11b0b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="canonical-681a702a86f432970617e6fda8c5b83d53a7fd7fd134fbc147821d5b9aa41000"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="canonical-2959cb87fc336f0b4fc6bf2b94224b5078712b060cb1f8b91b6852c0f5c792bc"></a>

## Direct properties — single_lb_app.enable_discovery.discovered_api_settings / ee39d3a11b0b / 3

<a id="canonical-b8ba5999cfbcfbc7ba017aeba977e2e737d0e49c483ff6bc280ebacd4a616dda"></a>

<a id="canonical-223d2c4b6625c077f9a629671a8ed8e7f40adc1b22fb665f6080567390e67193"></a>

## purge_duration_for_inactive_discovered_apis property — single_lb_app.enable_discovery.discovered_api_settings / ee39d3a11b0b / 4

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-55e8b369a81fbdd0f7e888ed16f982daf238b3eb9d991769bb25d218778ecdce"></a>

## Next pages — single_lb_app.enable_discovery.discovered_api_settings / ee39d3a11b0b / 5

- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-74b5d9db3adb30e3c2c4f4bd453643eeb0185c6885bad2ea15d00516f937edc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95503c3d91999e6c004b89043b39d257b9bb3740d464b30971f302a8088a9891"></a>

## single_lb_app.enable_discovery.enable_learn_from_redirect_traffic — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 140de5e4eb18 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- single_lb_app.enable_discovery.enable_learn_from_redirect_traffic

<a id="canonical-76fd267357d40a466c0845d5f18b3553d98bc0285cd9c3a1b599828177ab023d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-49595041d410dbd7c6b2f45a0c07a45b9dcab76ac3f9700a391942faca7dfd99"></a>

## Direct properties — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 140de5e4eb18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50fe494dde038a686bf8071ddd7f3d264a17d7c038ae9a9a28ee417fe8499abc"></a>

## Next pages — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 140de5e4eb18 / 4

- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-025.md#canonical-84a112a13d341a559a391e70465cc059dae641a202912c47b772dba42895388b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-507953ce2096f48d5567042afa8acd9f7068cb765a3c4a4797c23c382b99b3ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39c3805420cf8bb31e7c91ff78a48a71ddda2ff926b6f03c45c2a5ecbd39aac0"></a>

## single_lb_app.enable_malicious_user_detection — single_lb_app.enable_malicious_user_detection / cd7ec6bc6236 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- single_lb_app.enable_malicious_user_detection

<a id="canonical-aaf062eaaa48a8f868bb75b08ce2aa1bb4f80feb7551824b42aa986b90e8d232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-682f6d0f5f057705ef675f52c7b7263b842f6e967ece61e043345440bdd552c8"></a>

## Direct properties — single_lb_app.enable_malicious_user_detection / cd7ec6bc6236 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-304bda999d4d3a1bcfeebec2e845ea52ccff38c31a296e3eaf8ce8294eba20b2"></a>

## Next pages — single_lb_app.enable_malicious_user_detection / cd7ec6bc6236 / 4

- [single_lb_app](data-sources--http_loadbalancer--reference--group-025.md#canonical-af9239eb79a8754b4400973a73525f0486e4e23c7cece96a2ae3f76907b919f4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b08a6f056656710a9d8870bdb14818e3364d4d3b2f17386ae851550fede5001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a890e7ca9f73f7d4bc5db52c6d69775b9d466dd47daae0173dbab41cf93c57cf"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / b0d595a1b0c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- slow_ddos_mitigation

<a id="canonical-d9359e50869c3b0bc38d1da7ede4848fd8ff604aebe791418cb329907e872570"></a>

Type: `"single"`. Computed.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-025.md#canonical-d9359e50869c3b0bc38d1da7ede4848fd8ff604aebe791418cb329907e872570)
- [system_default_timeouts](data-sources--http_loadbalancer--reference--group-025.md#canonical-b5fa9aedd5d36745268bbea106e39f397aced72deb78777db5b7f86a719d96ca)

Select alternatives according to the provider validators above.

<a id="canonical-a9285ced37ab7a5eec1437499d5da7dd6938cbf9fbaebf6d93c26b736d873014"></a>

## Direct properties — slow_ddos_mitigation / b0d595a1b0c9 / 3

- [disable_request_timeout](data-sources--http_loadbalancer--reference--group-025.md#canonical-6f0101d9c8624ed503c5ff215052556d98b1510b187003cc20b9cd6b1048f569): complete subsection reference.

<a id="canonical-49d799d8acc8955d97d1d98570177bba8f6642230b6a636e381b14af5a7076b8"></a>

<a id="canonical-d0ae7789f60f1daf2fdb9e64e97309cd60d324ed0c37567f7094381f2dd50589"></a>

## request_headers_timeout property — slow_ddos_mitigation / b0d595a1b0c9 / 4

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-26322ec0f9267ee09f62971543ae2af7d8695c60b81a379b13a425c92e1789a3"></a>

<a id="canonical-8567a9817128a5cbd320732965c970074ed39bcc503ac95ac2c7008b8feca9ab"></a>

## request_timeout property — slow_ddos_mitigation / b0d595a1b0c9 / 5

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-5188de670e2234a5b64453bb11961c2602a5ceb480845288a80ce7e0c7b36c61"></a>

## Next pages — slow_ddos_mitigation / b0d595a1b0c9 / 6

- [slow_ddos_mitigation.disable_request_timeout](data-sources--http_loadbalancer--reference--group-025.md#canonical-6f0101d9c8624ed503c5ff215052556d98b1510b187003cc20b9cd6b1048f569)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6f0101d9c8624ed503c5ff215052556d98b1510b187003cc20b9cd6b1048f569"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d85183052fb982732e97ec2a51ff083b15cc1311c5a387a0b6bddb029a0663"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / 827bff2dc8a9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-025.md#canonical-5b08a6f056656710a9d8870bdb14818e3364d4d3b2f17386ae851550fede5001)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-a5fa6a5367cfb648a76389f0e8663c4d1d430dddf67b2eaa7e1bddd9535de4fc"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-f93ad441f63b175968c9ff8da2ee318ba37e08ec491b340023c34b906f17b761"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / 827bff2dc8a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a71d63a4ee5acb91cdbf2bd4849263e17168fe84f070e3fbef040e859225571"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / 827bff2dc8a9 / 4

- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-025.md#canonical-5b08a6f056656710a9d8870bdb14818e3364d4d3b2f17386ae851550fede5001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5e72e062dee812d3a84a0f29d22dbb11853853af05d0b9684665c211ae9eeda8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82f9e15bd3d3f3465d289b78534b73fad8d54f27a0c44a7609bc57e955188261"></a>

## source_ip_stickiness — source_ip_stickiness / 5397b77a2f92 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- source_ip_stickiness

<a id="canonical-aec3f86d7a9328509fd797ba929ddf4c65b50b051dc9cfc35600921fd2f6b5e0"></a>

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

<a id="canonical-f9b2793542e8188694e323a0f1b57519ed25627b159a18b122f65a158dd66446"></a>

## Direct properties — source_ip_stickiness / 5397b77a2f92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9ad4f2aefac44b19898fa693ba4926126e7298b265874fde55c6c941f87daad"></a>

## Next pages — source_ip_stickiness / 5397b77a2f92 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-425fe42824262babbc98f50fc32dde2d8970ab573dfe5fa4c70f101bd9509ebd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9afc49bf9f405caf05fb7329c4ab0dcfd8f846a1e611100381498e1408c13134"></a>

## system_default_timeouts — system_default_timeouts / 5246542ae1ea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- system_default_timeouts

<a id="canonical-b5fa9aedd5d36745268bbea106e39f397aced72deb78777db5b7f86a719d96ca"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system default timeouts.

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

<a id="canonical-8779632528147ea489c72d5bb43fe3da5a604acffe4a97214c4486a91de03dcc"></a>

## Direct properties — system_default_timeouts / 5246542ae1ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2766ffa5ce51808d05abce66a1fed0bc7e9051470feb005e7e85a66496241056"></a>

## Next pages — system_default_timeouts / 5246542ae1ea / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d401abf957a4b8d5c29619747b61befeefe5c96c2df65dc67e2555f280c71601"></a>

## trusted_clients — trusted_clients / 625cd3688bd1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- trusted_clients

<a id="canonical-8ef5e28eadd254ecf2f0b0e65199942dabfc76dc57517fd9e0a9386f7f9d7f25"></a>

Type: `"list"`. Computed.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-322c085f78b9cf620389c062980e3085ad451455da25da86ade31abbe1dee6f7"></a>

## Direct properties — trusted_clients / 625cd3688bd1 / 3

<a id="canonical-74cff7bd2057c0212e918b5552fad95e70e311fc5b2b4ecc78a7963571f1adab"></a>

<a id="canonical-263ae4e2a5c190213abcaed4ad49cedf7b6ea66b6e6f8f7aa62716248d0161be"></a>

## actions property — trusted_clients / 625cd3688bd1 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1aee77f8ca136f4a388d55749d05aa47a8e8addbac96c20f1344b5f22601190d"></a>

<a id="canonical-25c6104eaae64189bdfdf45622cf41f418912dfa6f51616b9c4d4660328dff61"></a>

## as_number property — trusted_clients / 625cd3688bd1 / 5

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-69d51247217700c6fe3abcb00b6e28a51006b8220c7864c0e69157f86d681144): complete subsection reference.

<a id="canonical-b0d73e1efd680d7a89a29ba17efed72a4dfb110c1e1c6e307ff67faa87cfc57f"></a>

<a id="canonical-52cd81731a28eea37c4baef52a1c5dd46ec7109b0aa151b66c804d97ced1d8ab"></a>

## expiration_timestamp property — trusted_clients / 625cd3688bd1 / 6

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--http_loadbalancer--reference--group-025.md#canonical-ec30876b1da32b43623fdde6de1625e854e932cf0153702e8bc041bbd2b5e590): complete subsection reference.

<a id="canonical-2bd5ec364419faa43a570ad2c43ffb5c5e29145c031d5aae54ca584b031eb60e"></a>

<a id="canonical-3c7e90ef36455f0f26e13cb32f3f93fda9528f20a69745c0fb625101c1268997"></a>

## ip_prefix property — trusted_clients / 625cd3688bd1 / 7

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2561bbfb8ff54d8dceb8f63288a9bf32821ebd5cd3b28edc2624b09eb6c414b4"></a>

<a id="canonical-d457fe66d06f8294c3ac214a28f24b3dba2ffa4f2f4f0e9ab85e0254fe9ba9f1"></a>

## ipv6_prefix property — trusted_clients / 625cd3688bd1 / 8

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-025.md#canonical-e2209fd12e6b0077749a6e54f5b667b5b759e7aee2d7e687de7716b9eed754df): complete subsection reference.

- [skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-d6f8f751c565dcccbabca80b4f1693796920ee0d9d39a547a9cdd7a897124846): complete subsection reference.

<a id="canonical-ac973ca855276ff266f92d27309643663f0a9a81dac52864e575068ef1392c6c"></a>

<a id="canonical-a0a7f43cec9e6f95f697806b3bf110e1fc84dce603701a0ce347d717986c1a98"></a>

## user_identifier property — trusted_clients / 625cd3688bd1 / 9

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-0288f8008dd20c138a2eeddf71096e1538c75a5c6da5d2fa216be3b0f41917f4): complete subsection reference.

<a id="canonical-8ec1589635e3acfe20712a59c98eb1ed2a5b98640eaae3ba382960b95c6ecd6e"></a>

## Next pages — trusted_clients / 625cd3688bd1 / 10

- [trusted_clients.bot_skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-69d51247217700c6fe3abcb00b6e28a51006b8220c7864c0e69157f86d681144)
- [trusted_clients.http_header](data-sources--http_loadbalancer--reference--group-025.md#canonical-ec30876b1da32b43623fdde6de1625e854e932cf0153702e8bc041bbd2b5e590)
- [trusted_clients.metadata](data-sources--http_loadbalancer--reference--group-025.md#canonical-e2209fd12e6b0077749a6e54f5b667b5b759e7aee2d7e687de7716b9eed754df)
- [trusted_clients.skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-d6f8f751c565dcccbabca80b4f1693796920ee0d9d39a547a9cdd7a897124846)
- [trusted_clients.waf_skip_processing](data-sources--http_loadbalancer--reference--group-025.md#canonical-0288f8008dd20c138a2eeddf71096e1538c75a5c6da5d2fa216be3b0f41917f4)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-69d51247217700c6fe3abcb00b6e28a51006b8220c7864c0e69157f86d681144"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee53208689a921fc0ad4f5fb0285dda422e4ebf2f9c8af029853bd3c5daa9f9"></a>

## trusted_clients.bot_skip_processing — trusted_clients.bot_skip_processing / 39aff14bbd48 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- trusted_clients.bot_skip_processing

<a id="canonical-414cc8699b87c773f2b1a4ebd8a3563143d7014502027b46a1b3e1d9533e72b0"></a>

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

<a id="canonical-80b326045881e4fa14f01428b6866b04596e22feecbcfb808790d8789693ce96"></a>

## Direct properties — trusted_clients.bot_skip_processing / 39aff14bbd48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-204af5cdf498ef3d3ef0cd28aefb773c16318dcdc3332598945ea71b7b2bc585"></a>

## Next pages — trusted_clients.bot_skip_processing / 39aff14bbd48 / 4

- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ec30876b1da32b43623fdde6de1625e854e932cf0153702e8bc041bbd2b5e590"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8dd04e1592a45f0f2ee5534203fe25e2da7a7d4019e6ed7f6f1d668124f9d2f"></a>

## trusted_clients.http_header — trusted_clients.http_header / 730d6add5cd5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- trusted_clients.http_header

<a id="canonical-8bf4601f25d02ce5b49dccdef7184e8b4f5621be97dad83d70d67b0fbe825270"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

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

<a id="canonical-ac97e969a3070004f192a59fa9c95de4426e4fd2d06df3d97abc23bda6d53893"></a>

## Direct properties — trusted_clients.http_header / 730d6add5cd5 / 3

- [headers](data-sources--http_loadbalancer--reference--group-025.md#canonical-9b89b47a2aa4e5cd317776153fca03345b3dcc0fd74ca28e80af98c143a3ef13): complete subsection reference.

<a id="canonical-e65e35ce2b88245eefe7de86fd0b438a58d61d552c55dd893c12e86c7d58b35d"></a>

## Next pages — trusted_clients.http_header / 730d6add5cd5 / 4

- [trusted_clients.http_header.headers](data-sources--http_loadbalancer--reference--group-025.md#canonical-9b89b47a2aa4e5cd317776153fca03345b3dcc0fd74ca28e80af98c143a3ef13)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9b89b47a2aa4e5cd317776153fca03345b3dcc0fd74ca28e80af98c143a3ef13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab1b6ac9ba95543d878d4aeabe3d814c276110be27a93d86eda9a5aa60c92cae"></a>

## trusted_clients.http_header.headers — trusted_clients.http_header.headers / d3121a502495 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [trusted_clients.http_header](data-sources--http_loadbalancer--reference--group-025.md#canonical-ec30876b1da32b43623fdde6de1625e854e932cf0153702e8bc041bbd2b5e590)
- trusted_clients.http_header.headers

<a id="canonical-ee3bcbad9ef2e813a8bd8b3f38403c48dc28077cb1b3397d9e23b0d66082af30"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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

<a id="canonical-d30d1f3168ae743b76d7c9926abd3d617303d0bb0c6e675fae9042b618e38dad"></a>

## Direct properties — trusted_clients.http_header.headers / d3121a502495 / 3

<a id="canonical-44dcaf954cab9f4af3529e4d4d5c140cabae81c9af4c87d3d5d0c3b6c147bd02"></a>

<a id="canonical-ca9d1a88b8d52daa14fb50bf9639330c935f9b0f269d35454cd18ea79abda16d"></a>

## exact property — trusted_clients.http_header.headers / d3121a502495 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-fae9f508882b391052d788ce9e9804b9448b14bcd4e19d7438a35937914f34b9"></a>

<a id="canonical-cb0eaf899065ddab457644b2994db11b92d5f8397478dcd4fb2f2d858d8cd911"></a>

## invert_match property — trusted_clients.http_header.headers / d3121a502495 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-ff0c37b980ec5981311a0b5386c122eb4a64c2a35a39d40c39a9d17d44ce9f2a"></a>

<a id="canonical-fc9c3007df232532da4dbf4d98b721b53e9c3adb53ca496c995c1599bd333ad7"></a>

## name property — trusted_clients.http_header.headers / d3121a502495 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-a2f97edf3e9914dd43be5651c23fab74dcf48f15addc5b126574509bf51c7ff5"></a>

<a id="canonical-8520b83e03ca769a9d213dd17a33a23b2fabea3e537eb31346a97cee7eaee66b"></a>

## presence property — trusted_clients.http_header.headers / d3121a502495 / 7

Type: `"bool"`. Computed.

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

<a id="canonical-f1432d2af3a34ea023c93acd351715b9412061724351346eeb5123d6233e8191"></a>

<a id="canonical-dc783e04da22ee9a912d04af69223bbb86514ad7d02506800d780684926a6c55"></a>

## regex property — trusted_clients.http_header.headers / d3121a502495 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-543cd6b7a4f3a2752d3b477bd19d8e49ca479b29c5dd4170285e15b60050a3f8"></a>

## Next pages — trusted_clients.http_header.headers / d3121a502495 / 9

- [trusted_clients.http_header](data-sources--http_loadbalancer--reference--group-025.md#canonical-ec30876b1da32b43623fdde6de1625e854e932cf0153702e8bc041bbd2b5e590)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e2209fd12e6b0077749a6e54f5b667b5b759e7aee2d7e687de7716b9eed754df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b07be57d77733b913ce5e4d9467e206f5c885d948d4b44ba99dc012371072845"></a>

## trusted_clients.metadata — trusted_clients.metadata / e84ec3015e37 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- trusted_clients.metadata

<a id="canonical-bbb82cef2af80671c586999ee035739225435678859f06323d4d4046e7e2529a"></a>

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

<a id="canonical-45044b48a62c3a027d93d51555d5cc326996624af3ab4065471cc53a84ae5ff9"></a>

## Direct properties — trusted_clients.metadata / e84ec3015e37 / 3

<a id="canonical-489694b9443ed2d34dbca1978a9621f4cb605d4b0e449ae717996150440c3ecb"></a>

<a id="canonical-32e5ed9929a6054b990d7c7f2b3ab83f4ff3cb4072cfc05d4409e8428a9b3cf1"></a>

## description_spec property — trusted_clients.metadata / e84ec3015e37 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-b58bcb045a7de359c4583dcf2252a4d81bbfb560b3fac4cfd2444ec04c330655"></a>

<a id="canonical-e6fe491891065a7fb1268a5aac8256fc8e3dc069620b90aa09cdb1ca9066ea73"></a>

## name property — trusted_clients.metadata / e84ec3015e37 / 5

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

<a id="canonical-055eba6db67c97a92886ed746546feb6d165427d12363549fb11ab7c1f2ac6a7"></a>

## Next pages — trusted_clients.metadata / e84ec3015e37 / 6

- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d6f8f751c565dcccbabca80b4f1693796920ee0d9d39a547a9cdd7a897124846"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b0b1eaf3bb4a0b86bddeb7bfc2a5ea5efb5f75860b183a1a12e2dc8119edd9b"></a>

## trusted_clients.skip_processing — trusted_clients.skip_processing / 2c0190716c1f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- trusted_clients.skip_processing

<a id="canonical-d62eea2cade4c4bbb89647151eaa941d05485303b02599373ac770412e54e376"></a>

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

<a id="canonical-7d9f6f8fb01d647f937578f2b4d66dc3bd87a9afbd4ce5db25ff0dce88145e8a"></a>

## Direct properties — trusted_clients.skip_processing / 2c0190716c1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8ecd559456af852f516427c39ec2376e97c60de4bf054cd6864e4c7ad4dc580"></a>

## Next pages — trusted_clients.skip_processing / 2c0190716c1f / 4

- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0288f8008dd20c138a2eeddf71096e1538c75a5c6da5d2fa216be3b0f41917f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ddb54a9f21787aeb99d7af4aace17f3dc4cdab13dc9eb030551db0f37345cfb"></a>

## trusted_clients.waf_skip_processing — trusted_clients.waf_skip_processing / c5175ebd039e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- trusted_clients.waf_skip_processing

<a id="canonical-a0cf0fd5730d4efe12e1ad8f00364486ef6d6ed2933b45b43f64ca3e332948f0"></a>

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

<a id="canonical-bc13c610d3bdb9de5bc4f70c757ed4da1d463e25de3afa5a9726d7384ca1979b"></a>

## Direct properties — trusted_clients.waf_skip_processing / c5175ebd039e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2d11ed4d8f736018368fe622f637bc08a1995247faef3942ece58978f45bdee"></a>

## Next pages — trusted_clients.waf_skip_processing / c5175ebd039e / 4

- [trusted_clients](data-sources--http_loadbalancer--reference--group-025.md#canonical-dcb51cdd2d49f8b7dbd34a96c17f533bb9ec118232e063bb7e15a5f11b0fb9c8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ab8c8d0d78dee7b8447509f4d2cb563d77206d05b230ebcefad4e0a53423b0f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-253134e37f2b71cf36b63383e2f496946c4d17d3afc0fe11e6c5ed00c24e330d"></a>

## user_id_client_ip — user_id_client_ip / 94f7126a240d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- user_id_client_ip

<a id="canonical-622e7778a2ffdff8a24d11b93b5602720ec36a55ca37ae0940fe86972ce8f45d"></a>

Type: `["object", {}]`. Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

- [user_id_client_ip](data-sources--http_loadbalancer--reference--group-025.md#canonical-622e7778a2ffdff8a24d11b93b5602720ec36a55ca37ae0940fe86972ce8f45d)
- [user_identification](data-sources--http_loadbalancer--reference--group-025.md#canonical-4e045fe79e8c661c4a83e23fbc7454d38420f79fa4c0829c8cd9a112b278e320)

Select alternatives according to the provider validators above.

<a id="canonical-1064da608a517762c474b0033238b363435ef08809d20c9d0f4011cd9a13ad64"></a>

## Direct properties — user_id_client_ip / 94f7126a240d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec6dbbf0f3f4da21f1e586002ceece80fda76ef0f7ba089b58c96843b2c63d98"></a>

## Next pages — user_id_client_ip / 94f7126a240d / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f09f430b73b0f125c1db603c889dfb4e71bc501f71ee4876284156502d4dc20a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1275a1440d592c8f35b80b888cf1687de1361890640efb0f5a89351c45267a1"></a>

## user_identification — user_identification / dfdcbc05625a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- user_identification

<a id="canonical-4e045fe79e8c661c4a83e23fbc7454d38420f79fa4c0829c8cd9a112b278e320"></a>

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

<a id="canonical-89662c67948e138d93fe66a708796912b72a4c215a5b34f8d4567014c8b705ab"></a>

## Direct properties — user_identification / dfdcbc05625a / 3

<a id="canonical-a5c3e1278164535e909a445565a6a2794d589b1345f94140ed6deba90cf1c352"></a>

<a id="canonical-513341639e5f4b27020d2c8b3514ad0582c11c49622a361dd4924dc32e7c4616"></a>

## name property — user_identification / dfdcbc05625a / 4

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

<a id="canonical-75e106a865502dc00735ef0bbe4fdf90435f20e557851ac5c58a4a4377196886"></a>
