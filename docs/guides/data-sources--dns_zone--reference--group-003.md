---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-31774f67f56af55071240ed080f1d3dda9d34f86f71f045dfaccf568e24b490b"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha1_digest — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 6468cbc6b1ce / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-7504b1c498fa12303cceeac603e555e8ef7b5c3b3d42a6e3044366b022565d45"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-2d81a35f5585af16586213b5e339bbc11a370146e8983637bfe871a9781f41cb"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 6468cbc6b1ce / 3

<a id="canonical-252b5ef1bf8dd1a84d5310cf5ef457eccc7e1177906ff5805a3520c74e959bd1"></a>

<a id="canonical-21ee92a0fbe151955afbc46e09dd803043ea109e9a7a35fd03e53c2a783d4bc6"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 6468cbc6b1ce / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-bdf8aa8534f6b38d7313b3ed4678fc561e96ee51332f039cd06fa3fcde4d7a6e"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 6468cbc6b1ce / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b58fd918f65b9c174a2110f19d1f177a293cd5825a765eb914707e3999f657cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c41d95e1546e3d42e0bb431e23f5fd11007f68b68a99fe19e8896b165dc5226d"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha256_digest — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 32c5eb7d8ab2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-d8ff2e1b49fcaafcf25674b095dcc6de333733ee24134086260157f195d8926d"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-a072ae8c75ea7fc74688bb3b716ed422df303e6bc7ee56a36e99a51b4f207eab"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 32c5eb7d8ab2 / 3

<a id="canonical-008bf175a56b85d9b5811551e0bdb861aa8b46f0df6566476753f75dce30fa4c"></a>

<a id="canonical-94e5f9d09d2d3e17da2882e8b3062b6addf2eb2a3814d14ba72d25e540fab1b3"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 32c5eb7d8ab2 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-baa75a70a0797225e9b0ec90f3e5713e8d532ec5031bb6c97f06c569d9c33bbe"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 32c5eb7d8ab2 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-c41c505aa680f3668929f4b32ca3e4936276d775404a4916b2132889c0aa9aee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9323508707553adbcbdf8ff722895090f4de57563005de7dd5706c1ece4930f3"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha384_digest — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / d6f2615e7eb5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-c3b79f4b9ccb28a571cd170743737b4855fb7db64f25a19fd13cdad1bf2049c3"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-6496f80679e4ae84bdfafa988115b385bcdd67f2b525e8101d8336d1100282d2"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / d6f2615e7eb5 / 3

<a id="canonical-51a78a6b6df8e956ed5c8cf930e532c05fc017c3de3ad9f901178232b3b919ad"></a>

<a id="canonical-d9a5175e13127f193a6ba37b77eef15c043e74bbec21c6fbfdf04d7e97b8ad9a"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / d6f2615e7eb5 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-4582e6376e380f89320d394e31a7605bc17cca1f098872743a71dc74a62e1464"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / d6f2615e7eb5 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-709880ec57d7e4685ef03f71ff548412c3dde7681ab0dceda336b398bef83fd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f8819f9ecdaf20f370f88c198a10a51287c9750956717daae8faefce27f5c7b"></a>

## primary.rr_set_group.rr_set.eui48_record — primary.rr_set_group.rr_set.eui48_record / fd8b9dbda679 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-810031a11394fca5e4f73d201eef2df69d1438e9d1e5440c666590ae840a695c"></a>

Type: `"single"`. Computed.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

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

<a id="canonical-079d2d97ec0ee87859b55ecf31d822534abedfdbf4d78f424fd0d580e8e31e16"></a>

## Direct properties — primary.rr_set_group.rr_set.eui48_record / fd8b9dbda679 / 3

<a id="canonical-dc9ae63b5aee65c40119880b21730639ac546cc854cc00710030806cc5518992"></a>

<a id="canonical-89ae4f2efd3acde3dd4fe277cd8a6fa0bbe5f5a9e2548cddaf0ab9a39dc91bda"></a>

## name property — primary.rr_set_group.rr_set.eui48_record / fd8b9dbda679 / 4

Type: `"string"`. Computed.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-59a2a901a2075caaeaeeefe5883585a095875bb17c859c7df2d30a095b7f9eff"></a>

<a id="canonical-f5ce20502f84b49bfac93829b4ae421395a11acfd74fa8345baaf5aa2a0bf6b7"></a>

## value property — primary.rr_set_group.rr_set.eui48_record / fd8b9dbda679 / 5

Type: `"string"`. Computed.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-00e684572ecd08a58efd3df9f92bb6a345a13f7b11f8783889e8c8921c23192e"></a>

## Next pages — primary.rr_set_group.rr_set.eui48_record / fd8b9dbda679 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-03ce7d61c3833deb592f05998294842f569a939a7d46517883fb7e7f9482eed1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66eb54ca800d257f22530cbcda8b9db71e603e9741c0c46c290cdcb684db9d39"></a>

## primary.rr_set_group.rr_set.eui64_record — primary.rr_set_group.rr_set.eui64_record / 826919c589fc / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-6686d6e014d02ba9405db6396066587a5b4165c26c4ec0edda637861dfacac5b"></a>

Type: `"single"`. Computed.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

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

<a id="canonical-88bd1ddebe650ae9b0e700c10a61cbf1b5f1899f857e1711be36acdb83e14b10"></a>

## Direct properties — primary.rr_set_group.rr_set.eui64_record / 826919c589fc / 3

<a id="canonical-eb34300a71ecdda472ea92475fd24d1e682d984728c7d08e3c4d430065a51798"></a>

<a id="canonical-959d0ca7b3bef1780775f093008c863456436fbc22a77cb07d8b7c8f34f32cdc"></a>

## name property — primary.rr_set_group.rr_set.eui64_record / 826919c589fc / 4

Type: `"string"`. Computed.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-bd811cdda1c9e4e57406c720f97ace2bdfe9006b450ad361d6963c97754d0cd6"></a>

<a id="canonical-91b1bee3fc9538d1602cd8a3971ee929396a79eea069c49d6338ba8c97e9701b"></a>

## value property — primary.rr_set_group.rr_set.eui64_record / 826919c589fc / 5

Type: `"string"`. Computed.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-3220cfde0fdfb8575ff48d102a7b1106fab1eb31a983d33ec8733fb67d3ea165"></a>

## Next pages — primary.rr_set_group.rr_set.eui64_record / 826919c589fc / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-6c20f8137cf306dce88f0b4265f888b4d3298580871c69a5a06e886ebc338659"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-497af762925d315cfba9056a7c7d2717220ea09654c578bccd04a931e2dc2a52"></a>

## primary.rr_set_group.rr_set.lb_record — primary.rr_set_group.rr_set.lb_record / 02e22267750b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-4408556c0c0318b5870bac11cf58181cf0b13b448ea5649eefd212cb54472ebe"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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

<a id="canonical-48bf62be70676db1e4c22a80065b004e48a57a2176053f13d80cb3d26ab56fc9"></a>

## Direct properties — primary.rr_set_group.rr_set.lb_record / 02e22267750b / 3

<a id="canonical-00e48968331000a826f9874fae909196e7bfbcaa8a18076e25d1a2b62745affb"></a>

<a id="canonical-b94a8f0e18ac4f0a05b25a545eb1df44733fa3a4b804241b8d12bd00a83fe968"></a>

## name property — primary.rr_set_group.rr_set.lb_record / 02e22267750b / 4

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](data-sources--dns_zone--reference--group-003.md#canonical-1911a1a67c9422a6468e1011fef10a8ee1afea9709bc487e56df8e1ab7fb509f): complete subsection reference.

<a id="canonical-b555dc586bef3a12254461202b975ebd5afe0c1a4eab5c87cf2483f722f57acb"></a>

## Next pages — primary.rr_set_group.rr_set.lb_record / 02e22267750b / 5

- [primary.rr_set_group.rr_set.lb_record.value](data-sources--dns_zone--reference--group-003.md#canonical-1911a1a67c9422a6468e1011fef10a8ee1afea9709bc487e56df8e1ab7fb509f)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-1911a1a67c9422a6468e1011fef10a8ee1afea9709bc487e56df8e1ab7fb509f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea56aff2412b7618430a8e1e552c3aab5dba4f8a274f93745e6a76aa5ec60920"></a>

## primary.rr_set_group.rr_set.lb_record.value — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-6c20f8137cf306dce88f0b4265f888b4d3298580871c69a5a06e886ebc338659)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-d9e2b52b7e6d30aab7103a2c86b6dff244ac5947b9a6f0263d3d9dd10bbc4355"></a>

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

<a id="canonical-dde1ff1180d2e9c4f7b3c580050827a597c2d8f610a2629112b4e76ff8709289"></a>

## Direct properties — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 3

<a id="canonical-008a89b7d5e756f29d25a8d1c0b3461156b42cd4732fbe6e0796e2598dad7f4a"></a>

<a id="canonical-2c5ec36ccb5b29d11bddf7f1323e38914f62c8ae7f8785479bf943acaabc7be6"></a>

## name property — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 4

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

<a id="canonical-f7e80e40d79db1a58e1ab44fad3e1e3f742d03c81033766124fdcdb3a92a7b1f"></a>

<a id="canonical-9f65f2f7d5e74e02472f3235bb2a01ca79c092bb5282593346fe70464e2ba2ad"></a>

## namespace property — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 5

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

<a id="canonical-bccc7aac05ad64bc0aa8bd088eb7e2196e49c84b628d380023aa5bf1fe8a1391"></a>

<a id="canonical-d06aaaab828f17e8da95ebc1218be4a7d15598696002734f377df6696116fa7e"></a>

## tenant property — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 6

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

<a id="canonical-4ff391c70cfa5b98e5297bee645dc86c99796f892161cdc499d0f04ad59f4be5"></a>

## Next pages — primary.rr_set_group.rr_set.lb_record.value / 055313512405 / 7

- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-6c20f8137cf306dce88f0b4265f888b4d3298580871c69a5a06e886ebc338659)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-d309f505f031ab7a4ff2608e5832c457699f46ca702fa82f161885667890bdc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58a13e626ddff2db9a8da40535dcf91ea4f2299c6f29ef2d0f3a8ef522472d7f"></a>

## primary.rr_set_group.rr_set.loc_record — primary.rr_set_group.rr_set.loc_record / c47520c74caa / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-b4bcc079b229881885ecb1aefc7aa1add949d4213e77520181f4f9a04b9545a2"></a>

Type: `"single"`. Computed.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

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

<a id="canonical-e3872f2860375e9eec51abfbefc6ec33675bde7be17b34158f495f48788cfd22"></a>

## Direct properties — primary.rr_set_group.rr_set.loc_record / c47520c74caa / 3

<a id="canonical-de598fe5afb8fb6540f6787451a2c9e105397f6c4cf1a8154d207ff4bf6810e6"></a>

<a id="canonical-3b97982939fd9a40c493bcaf341fa336327770987911b936bd84a0762c6a1317"></a>

## name property — primary.rr_set_group.rr_set.loc_record / c47520c74caa / 4

Type: `"string"`. Computed.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-bcecefda9f69a5a7042f512f47e4c0a744163a257f57c17414ffcffab2e200d1): complete subsection reference.

<a id="canonical-3e607d3d1f40f5371a38eacc1cb08d1b6ae5a661ce30f249fea378a3556a2797"></a>

## Next pages — primary.rr_set_group.rr_set.loc_record / c47520c74caa / 5

- [primary.rr_set_group.rr_set.loc_record.values](data-sources--dns_zone--reference--group-003.md#canonical-bcecefda9f69a5a7042f512f47e4c0a744163a257f57c17414ffcffab2e200d1)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-bcecefda9f69a5a7042f512f47e4c0a744163a257f57c17414ffcffab2e200d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1b0a260f5a43ad842e7a15759751f2f5e803ece9f6170520bcab94776a2f134"></a>

## primary.rr_set_group.rr_set.loc_record.values — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-d309f505f031ab7a4ff2608e5832c457699f46ca702fa82f161885667890bdc4)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-6559c3f701fe4a0746cf60de728e90a1452bb8249376e95e56e8e5fa91fdd896"></a>

Type: `"list"`. Computed.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d599a4284d5618701d6ca03e8b2f1a15c45d507643539502486478d9164cf974"></a>

## Direct properties — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 3

<a id="canonical-aa79600c34f1bf1e396d885c156dc549ba5e329ee9463efadaae58896866d7fc"></a>

<a id="canonical-65197364e39bda77201797f107158b71d70dc2cea5ed2de05bd0c289d113251f"></a>

## altitude property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 4

Type: `"number"`. Computed.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-caf16c03914926ee8f90d5e3c2e98628a730d594b51270bcc6d7036ec967d493"></a>

<a id="canonical-680d2257cc9ad62eacd10e81e8ee1960a5e4db0217fc8342374d0dd862ecfc84"></a>

## horizontal_precision property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 5

Type: `"number"`. Computed.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-d4a9b8b6cbc88a12f397731ebf112863e9bb844d2745472299ed4a27701eb31b"></a>

<a id="canonical-b36aeccddaa7b3531f51b3d108007a744811d2db5b03b02bba80057f4da0d9b9"></a>

## latitude_degree property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 6

Type: `"number"`. Computed.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-15f32dd5be966c25d1cba0206fbe2d8524b91742736491eeebd8525ee1c3b258"></a>

<a id="canonical-420587c4cfe73d6a866e289be183fce423e9c98d49b6b0fc19a0b1d4ee21d698"></a>

## latitude_hemisphere property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 7

Type: `"string"`. Computed.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e06d796dfe61401114e1ed2001a8c3ba346bd22cb58d13a1b13cf4dab97dd8ee"></a>

<a id="canonical-ebf3f9219f2c7945c4741bfe1d66e792385babae88d611746a899b901eab1806"></a>

## latitude_minute property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 8

Type: `"number"`. Computed.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-a44bf5120df0917dc83ddfe24f56f0a89496fe6d0e381f49d30a856a22e6a208"></a>

<a id="canonical-49b42f270b73ca2b30713687f7e7fe75fc5d172ef4293edbd8d368b17eecd4ea"></a>

## latitude_second property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 9

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-6633655c7690ad2944868f0e28c98ff6aafbcb100ec65d92459aaa45169a1235"></a>

<a id="canonical-bc69494a66f82cd477f3a81951f40e5cddf06dacc7a516af9159bad597d1fafc"></a>

## location_diameter property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 10

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-950601e71463c53a3da7c8f3d27bfc99d420f1d8b43b1343b61dbeea04ec1e6b"></a>

<a id="canonical-0aed514d7c07942d16bc24d2b9e07e665142935c25e132fffdf56db4b1d31f2f"></a>

## longitude_degree property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 11

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-d586e3896825a8588f12a7000a9098d3d10be67775f46f26c00f25959b11d039"></a>

<a id="canonical-dd2d257be8228fad0979ab854e4f8ba9385f4a3841957751df2cfaf8d60c88f7"></a>

## longitude_hemisphere property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 12

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4c97d03fb13e2b255346f64e4e7664ac7eaea7ae36115076e5f8698ac82c1c03"></a>

<a id="canonical-263a369dd488f69fe3cf35135c28b9734c3633de1500b2e2c8009cefadf687ae"></a>

## longitude_minute property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 13

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-bdaf60beaa069e9b1249e8d4f9b19eb04e4a810c251a97d2459be83f4da7e452"></a>

<a id="canonical-40bc7a2bbe9fadeafb22f760d03b2cacd37c6df0fa875768d14276f8e9a87406"></a>

## longitude_second property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 14

Type: `"number"`. Computed.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-a0f7d5c839c3150c425cab761b3ef5cf9171e606c6295853aca8318a3e936da5"></a>

<a id="canonical-80dd70d5507b5f2ce5d46d81afb36b1933ac78d2503cf8b120c41840fbb238fb"></a>

## vertical_precision property — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 15

Type: `"number"`. Computed.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-a9cfa81c7a2e60ca33f8af48da97d544d7029de60f826fd641197a62d7b8d2dc"></a>

## Next pages — primary.rr_set_group.rr_set.loc_record.values / 92bbd6b9ac73 / 16

- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-d309f505f031ab7a4ff2608e5832c457699f46ca702fa82f161885667890bdc4)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ae20712ac50719e301310b652f0b11851224148a83f26796e8d36e5f7cccfe0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cee6bcb0812f245507510c2a81afbe2fc45bf97e1738f5ccf0e1f03eef1c934"></a>

## primary.rr_set_group.rr_set.mx_record — primary.rr_set_group.rr_set.mx_record / 594eb184edf0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-4ec6c672accef470e07480c5ebb5d615e5007733d71ce6ccb767ff8e0c9bd484"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

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

<a id="canonical-fe6200815fa487f0706d236cc107adb84238f8d58d31306331e02b6c75047740"></a>

## Direct properties — primary.rr_set_group.rr_set.mx_record / 594eb184edf0 / 3

<a id="canonical-f53de048122a83156e354674ad2ae36f5f52a4575edc6d0cb72df4ba7fd7da3c"></a>

<a id="canonical-5eaf7286af58978ddb9ad7321250da6a24f6f80497a197e6eca81153d488183b"></a>

## name property — primary.rr_set_group.rr_set.mx_record / 594eb184edf0 / 4

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-f6261a71ccf33ee648fdd51a5f810a28f111bd404b0812e0e9f28c01fbe8bc7a): complete subsection reference.

<a id="canonical-a23eaa43ae276bf0cd7fbd1f07c099a94d785979bee71347e7c73458e439f343"></a>

## Next pages — primary.rr_set_group.rr_set.mx_record / 594eb184edf0 / 5

- [primary.rr_set_group.rr_set.mx_record.values](data-sources--dns_zone--reference--group-003.md#canonical-f6261a71ccf33ee648fdd51a5f810a28f111bd404b0812e0e9f28c01fbe8bc7a)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-f6261a71ccf33ee648fdd51a5f810a28f111bd404b0812e0e9f28c01fbe8bc7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e41e0659ccbf7f098783f841bd044879fcf1dca92984387bdaa121e491d8507"></a>

## primary.rr_set_group.rr_set.mx_record.values — primary.rr_set_group.rr_set.mx_record.values / adf127559c63 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-ae20712ac50719e301310b652f0b11851224148a83f26796e8d36e5f7cccfe0a)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-845f3db4794d898ba30aff13df27018137858d23bea2f1bbf41c3721b9a3a679"></a>

Type: `"list"`. Computed.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-0bda6492b64defb52f973b1329a697f58b528d60d289033e537bea9db0210e0a"></a>

## Direct properties — primary.rr_set_group.rr_set.mx_record.values / adf127559c63 / 3

<a id="canonical-a65bdd45520d2d51f3c4425e5e83e4af73f09e9692d99ec9c84fe74ce5170c42"></a>

<a id="canonical-d52ddfd0d13b42461e1d464a2252b9912d28c52f8426232b9487321e696ae99e"></a>

## domain property — primary.rr_set_group.rr_set.mx_record.values / adf127559c63 / 4

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-52783135a9dd7048e118602f26abc9f723f1de1765a90bb0e2b77a47f7c30d08"></a>

<a id="canonical-598199a2a6f6434c7be7171246c58a473f23b7ae895dda572d3924749ebdd63f"></a>

## priority property — primary.rr_set_group.rr_set.mx_record.values / adf127559c63 / 5

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ccaf201ea7277aa07df6ec0a0db75886e8a52a8d061a264b5ec97b10064ea45c"></a>

## Next pages — primary.rr_set_group.rr_set.mx_record.values / adf127559c63 / 6

- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-ae20712ac50719e301310b652f0b11851224148a83f26796e8d36e5f7cccfe0a)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-9c832aad21ecb00cad1e4c89f8729406860aecc74fecb638e1b5ec6b98a19504"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e83807ceb6c3631088185b961bb22e487233f29796ff111530e8bb321b8431b5"></a>

## primary.rr_set_group.rr_set.naptr_record — primary.rr_set_group.rr_set.naptr_record / 0cd0d93a5590 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-1fb8e9110182e03b5c8b3763fdfefcfe067428bd38eddbfe00987bfd451a922f"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

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

<a id="canonical-595667372e36b3ad36541984658763ec27a7a35b363bff3f315f08e53b34039f"></a>

## Direct properties — primary.rr_set_group.rr_set.naptr_record / 0cd0d93a5590 / 3

<a id="canonical-66176eba368fdc47fa655911a0968015ce58f83f30e9408de2b841a856a9fce7"></a>

<a id="canonical-567f94120c452d73a012dfffebe0b6e8eee82de54a8c514ae0ea5e3d918606d1"></a>

## name property — primary.rr_set_group.rr_set.naptr_record / 0cd0d93a5590 / 4

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-4bb62a288ce4d5caf3a4c1afde7e3cf8491c006ae51e453a99e82843c4e5c1d7): complete subsection reference.

<a id="canonical-fd759462112e3cfb06b7f5c439e8a3d8c1817e88b675122129548b5532bfe76a"></a>

## Next pages — primary.rr_set_group.rr_set.naptr_record / 0cd0d93a5590 / 5

- [primary.rr_set_group.rr_set.naptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-4bb62a288ce4d5caf3a4c1afde7e3cf8491c006ae51e453a99e82843c4e5c1d7)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4bb62a288ce4d5caf3a4c1afde7e3cf8491c006ae51e453a99e82843c4e5c1d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dfd2184f99341ffd48eaea350e69e725d2130edb0954a33b9b9ca1663757f3d"></a>

## primary.rr_set_group.rr_set.naptr_record.values — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-9c832aad21ecb00cad1e4c89f8729406860aecc74fecb638e1b5ec6b98a19504)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-a0b5123207883a92c073ab214e9e068d0abe695b09ba8123c886b4651420cb8f"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-99425930d22b23483170cfc5b91b1f92dde50e4ead4df76abaad9b539cdd2735"></a>

## Direct properties — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 3

<a id="canonical-4d19819777dbafd38e4ea6c699ff6546cd732359a2b263fd23c570adb6ec75c7"></a>

<a id="canonical-5188b873c6e0053a2d315f886aea9df3e37a9a3c4af39e04750b92c8005c72cf"></a>

## flags property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 4

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-b58adc7bd9fcac51ab30164a0a6365e133493d0ea8d8dab8361fbf95806b1449"></a>

<a id="canonical-b8348c78b4090bce220d418bc1eccf2846cba472b46389942447f2615a9b301a"></a>

## order property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 5

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0ea2ce6e24510c3f1570dd54c719c62425bf49475d3c4f59ab35f0cc8756e68e"></a>

<a id="canonical-5589bf89161cc6e873757c3978d86a31547cfe80e63e08bcf346805e748aa740"></a>

## preference property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 6

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-11b2cfee7e832939dc5aa8881241847c197cbdf9d6141507bb27128b03fb47a2"></a>

<a id="canonical-bac3c3658935c3af671ad0b3a4f1edfa4bde457e9b3a913e6a2f0a7239668441"></a>

## regexp property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 7

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-a082d684611d2966b115a86dab3f65e0b5c891b4cc636474190595402a520115"></a>

<a id="canonical-6ac7fa9cecddcac273861bfd5a12f4ac118a583db816f26835c2b63ef530df85"></a>

## replacement property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 8

Type: `"string"`. Computed.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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

<a id="canonical-2ae9174fbc1f230fba6216d9b7d7e127a442e0ddf100979d410e08a13738e0cf"></a>

<a id="canonical-d2e8dd2e4dd8907ed579ad9173d215f415061b05052a35b85c1b92592c87f96e"></a>

## service property — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 9

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-f18a73c86b1d3b3f72b1354c11e8695d64bd8ce250360edfee780829a6ecc3ef"></a>

## Next pages — primary.rr_set_group.rr_set.naptr_record.values / bcf4fd83ac6a / 10

- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-9c832aad21ecb00cad1e4c89f8729406860aecc74fecb638e1b5ec6b98a19504)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-71210dc397716cff2a83e2fd4597f220ffe3632ca2ffc86e8e930ce2d50568a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ffc9c2b6eaf81f43dcb6bcb3ecefdcdf5a5503d36999eaba274ac21afabf1c"></a>

## primary.rr_set_group.rr_set.ns_record — primary.rr_set_group.rr_set.ns_record / dee51f57d77d / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-3cda141896db18e8ec0c07290bb8033ee8a1250b5e6dd409d3b39a56faaf8753"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

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

<a id="canonical-97717ad4fb3cca1d618ae831604fabd24e3522df6dec2189c98f5b4f77cc214b"></a>

## Direct properties — primary.rr_set_group.rr_set.ns_record / dee51f57d77d / 3

<a id="canonical-2096f321da66047fe88b7d4167da9c51bf668062d05906ba319457e8cdde2171"></a>

<a id="canonical-fedddaaca276d946a40dd783d47cd5e4e4b681c47ee38e7f39cdae9434cd7749"></a>

## name property — primary.rr_set_group.rr_set.ns_record / dee51f57d77d / 4

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-99fc030a555b589c79ab6e502d1865c30ad1456ff856faed6c1532ccdc4bc471"></a>

<a id="canonical-77b240782b21641b69f0e217860f06f83b60b3c20fe434e84876a40c986f2c0f"></a>

## values property — primary.rr_set_group.rr_set.ns_record / dee51f57d77d / 5

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1aab7534b6d8042d00528f98e524a975c29ce1cf44c40685dcef8fe86b50505a"></a>

## Next pages — primary.rr_set_group.rr_set.ns_record / dee51f57d77d / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-a5cc8ab5f4efe1232ec6fc54d96daaf2d33d193a886c5095cede44588163b3fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3e9ca871bfc925b88e7d842b2cf4c9b2ff9744a5355e23751a63e1dfd959e80"></a>

## primary.rr_set_group.rr_set.ptr_record — primary.rr_set_group.rr_set.ptr_record / fc1cf22ce154 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-ccc9c1934d838037e7e1e58ec57e5fd0a548bdd400d2361eaf6a0d6106504316"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

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

<a id="canonical-f9960bca95440bcb44429c928d170d5aca8c2c0861d7c9d263867977cc431da2"></a>

## Direct properties — primary.rr_set_group.rr_set.ptr_record / fc1cf22ce154 / 3

<a id="canonical-b85b91bda539f487865a7bb2a0daff9f8c686d7bc06b941bb1906da52e71e754"></a>

<a id="canonical-6de4d140a0b8acd708ed3324083e7dce99bacfb0162e89e22f0143307166bf80"></a>

## name property — primary.rr_set_group.rr_set.ptr_record / fc1cf22ce154 / 4

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-ca0299aa6b15e3ec4138389ec7a4f31a0c13d959c245232d8eaba8921a51ea72"></a>

<a id="canonical-d6b87367b80fee9b6207cfbfee5ff46233ccf169d982370c5e024d47042088eb"></a>

## values property — primary.rr_set_group.rr_set.ptr_record / fc1cf22ce154 / 5

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2f0838b0492956674d683155c3f6bad1ac4a5c703588d080e8f996c53b08d788"></a>

## Next pages — primary.rr_set_group.rr_set.ptr_record / fc1cf22ce154 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-2b4cb5841ad7eb27081233d9bb8cd9f4791d4ce5744a828ddc7f86377f1446b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69238917d5566a63f69d104163f1871e4bc78c676b5e3f41dab48ac599df8415"></a>

## primary.rr_set_group.rr_set.srv_record — primary.rr_set_group.rr_set.srv_record / b218256b5429 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-40a683168442e781ca09825c53ec073e529ad9fc1ab64bb88871eeaed9a97fb0"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

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

<a id="canonical-18b400702ca5f98a4cee9a098ef89f982f54789bb35386068b4a728834ed23d6"></a>

## Direct properties — primary.rr_set_group.rr_set.srv_record / b218256b5429 / 3

<a id="canonical-b2ea5c34eac052b424f33a9679b967e16b7a4e9567d02e31abc90d1722997343"></a>

<a id="canonical-cc5236cca300fe5a34b3521934d4366b7a09ea58b0b60016ecbf7273d41ac875"></a>

## name property — primary.rr_set_group.rr_set.srv_record / b218256b5429 / 4

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-254f95b230ecdd09b2d748a9d81e991c41970545ba782bb5d7380fd07671138c): complete subsection reference.

<a id="canonical-d1f3bfab9d6be79f55ae981b37bb093f1a45ed314c1d312997a15ec068e9271a"></a>

## Next pages — primary.rr_set_group.rr_set.srv_record / b218256b5429 / 5

- [primary.rr_set_group.rr_set.srv_record.values](data-sources--dns_zone--reference--group-003.md#canonical-254f95b230ecdd09b2d748a9d81e991c41970545ba782bb5d7380fd07671138c)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-254f95b230ecdd09b2d748a9d81e991c41970545ba782bb5d7380fd07671138c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d47606d40c02e62cb040759c3452b6747162dd8359197d6f565160aef2690bec"></a>

## primary.rr_set_group.rr_set.srv_record.values — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-2b4cb5841ad7eb27081233d9bb8cd9f4791d4ce5744a828ddc7f86377f1446b6)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-a6bd35d735280c3a34b61e3450255b35b37036d0abfa24f115017f0669dee80e"></a>

Type: `"list"`. Computed.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-fb9d4d37221ea24af2839de4a8a31dd3894d81c9a3b74d15a4ad2c7b1389f49a"></a>

## Direct properties — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 3

<a id="canonical-c386881c49f2e04db91497ce02d0adae929d7f541e18aecf6b34b426398a4bec"></a>

<a id="canonical-86d3a97ed3cf079bd596e474ede2c616a3dc8d3b099f22d615befe082a341c54"></a>

## port property — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 4

Type: `"number"`. Computed.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-e4a39153296d7a1bd1e616c5dee544f1ae992106f8197278e17bc8d1434c6049"></a>

<a id="canonical-63460ed4a2928fe918495173624e6d0a90ac15567bbce814f668bacf568eb7db"></a>

## priority property — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 5

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-aac320bc3d66ecd97bbf8758914a399ea742b7105da0a8bd907b53512420c031"></a>

<a id="canonical-2fb3c758d21114524ece7597e18e56fa8be76ff899543feb77da8feab32d60cd"></a>

## target property — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 6

Type: `"string"`. Computed.

Hostname of the machine providing the service.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-37a06d4fdd276352840f3780b72320f6f1a6d078304e307a6ecc86d2cc24fda4"></a>

<a id="canonical-202319403ea1df623683e9aae62aa4183cb72a6fbaac4dbe2d8d1e2959877140"></a>

## weight property — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 7

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-31dc00ce2ad84cc4f119a72fd8d700cead2a1f488d7dae4af4288bf3a1f5133b"></a>

## Next pages — primary.rr_set_group.rr_set.srv_record.values / 6d26d3b91c5f / 8

- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-2b4cb5841ad7eb27081233d9bb8cd9f4791d4ce5744a828ddc7f86377f1446b6)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cac3f0d1db2b3537d84ddb9cdd2bf42a09df589e1de50536ed8d5bed464af329"></a>

## primary.rr_set_group.rr_set.sshfp_record — primary.rr_set_group.rr_set.sshfp_record / 409016a5e360 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-991b84ba14f554879e51f8f8d034a0790e4b9a26e957baa6d0d8002c18e9779a"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

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

<a id="canonical-20957a6836e1aff6d2b9f9be4bbf79dfb0f05226803c3b5077a86ffe35be333e"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record / 409016a5e360 / 3

<a id="canonical-6d3405e1f9ef813dfbea16372ab5f98947574df26e381ece6f26d093d7515849"></a>

<a id="canonical-8cbe47b243c93cafbfffd707e46e4a9b1d4fe6d6f3d06292f6850af2dab52305"></a>

## name property — primary.rr_set_group.rr_set.sshfp_record / 409016a5e360 / 4

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394): complete subsection reference.

<a id="canonical-3e7cf0ba21ecfa77b5e26dc9e51c278c1f69d64bf1bf6cce2a03eaed3c3cd3f7"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record / 409016a5e360 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0625438598e71e30405ff311b2ba7e36930fec67d9af2e347918570aa2d71c6e"></a>

## primary.rr_set_group.rr_set.sshfp_record.values — primary.rr_set_group.rr_set.sshfp_record.values / 55560c022a9b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-bdc3cd1f0d8fb860080a18600c3d7135fca67ef979903e95c5b30e633ee89079"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-c5eb15ee8e1a51f7e3ad0add854ec443b04e43d74e4d72f29a6a38c798c30b93"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values / 55560c022a9b / 3

<a id="canonical-702cd975f7b85a731a98fb9a2648971942e3e22fb47bcdf86d3ec2296706eab0"></a>

<a id="canonical-3bd902b48b1161181396cd7014b27ece8bafb305ef5f0349eb8a8d8ba25ad8f0"></a>

## algorithm property — primary.rr_set_group.rr_set.sshfp_record.values / 55560c022a9b / 4

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-663ed1d16591460fa13e555d3ad3d263a0190fed3fb64e44d725a13a6f16c336): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-390e7f3f0e89b1ea789b45502ae12ef068a3b03248e5590ac8051048df8473b1): complete subsection reference.

<a id="canonical-cae655e45acaca00e277eca7a7463b1afb953ca1d067eaf4984598ba1620b873"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values / 55560c022a9b / 5

- [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-663ed1d16591460fa13e555d3ad3d263a0190fed3fb64e44d725a13a6f16c336)
- [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-390e7f3f0e89b1ea789b45502ae12ef068a3b03248e5590ac8051048df8473b1)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-663ed1d16591460fa13e555d3ad3d263a0190fed3fb64e44d725a13a6f16c336"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d820fcf199abd385b4783246c3edcb0ad9cb3aaf477e85f2ea5043c4ecd1e627"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 8a95698808bc / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-741c22aedb434e08f27bb4dec4dad0883a95f4623c7246060b781e8129745fa9"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

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

<a id="canonical-0af23d70008f12804a4e469a33ad7ace2a397fffbad7eae1c18e1b0dd274b7dd"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 8a95698808bc / 3

<a id="canonical-2e01006ab08d3cafd9fd7f0de99b1cd5a4ec57545c4c9ed512dbf72d53b9d03c"></a>

<a id="canonical-90fd855086bf44cdb19b2dd3af6dfec10948dbbaf8f1c2a2b833e64438c029c1"></a>

## fingerprint property — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 8a95698808bc / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2c68b66f96a5e7e43e865305ec92d9a54c39408d592d3632a6aae74909d98363"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 8a95698808bc / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-390e7f3f0e89b1ea789b45502ae12ef068a3b03248e5590ac8051048df8473b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3e5c5a804674355cfba2e5f5267e26abf3e8e45b945cbe67a191659099f792e"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 1be2fac5941e / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-89b6d74582130a3687d7e0babb4edff1ce48cbf839fd07cbab0d77c9ead56516"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

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

<a id="canonical-25d100ba5958fda489696a981ac56170d507b0350b8083004b7b4339c69f068a"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 1be2fac5941e / 3

<a id="canonical-e18c17e61d7ee3a785d620ea8948ec65613eb02cb2edaca787317544096a5217"></a>

<a id="canonical-2bde790e8bdc55a7bd1d70dfd4ce5a0c699fe9073dff70fcdcd010ba493ee875"></a>

## fingerprint property — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 1be2fac5941e / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-3aa353990d1abddbaff7bc144a7b697cbec814db3cf388410d6e3021e6a0500a"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 1be2fac5941e / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-e164fc6ac133ecd0c1f143dbb0c6f7afaca55868e55780f6245a159750810394)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-585687c64fe7863dd11f33eab30674fb6cda389e728007a92787a434b499b119"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f30d259bbeeffa9bcc2ef82d2b5a7904ca4f54e9ead9237784c113f8cc99148b"></a>

## primary.rr_set_group.rr_set.tlsa_record — primary.rr_set_group.rr_set.tlsa_record / 7300f9febf54 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-51ae6a52fa37f05c53e9569f4fef39edfe86bb9079b50a6d9c6a35872a194ac1"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

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

<a id="canonical-0607e479f16c44a2a3bcc7be845d41284225567303e43bf356b3fe81568fba60"></a>

## Direct properties — primary.rr_set_group.rr_set.tlsa_record / 7300f9febf54 / 3

<a id="canonical-c715be307c0a5c7204a80f8a3d53eeeac710528e34dbd35feefad5c40541cbf1"></a>

<a id="canonical-fce9d9ebd5f3163b1b403e9609b12c3972a486beac3f520e725b917c1a3ea4ee"></a>

## name property — primary.rr_set_group.rr_set.tlsa_record / 7300f9febf54 / 4

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-b0e80110fc74531c02b74b402dbc244a5d93a645bc21cab7db8c0dd2a753f72e): complete subsection reference.

<a id="canonical-1efb96316cb7ce86a2db8b856b21ffaf0ef724c60b2992978e1493950008f855"></a>

## Next pages — primary.rr_set_group.rr_set.tlsa_record / 7300f9febf54 / 5

- [primary.rr_set_group.rr_set.tlsa_record.values](data-sources--dns_zone--reference--group-003.md#canonical-b0e80110fc74531c02b74b402dbc244a5d93a645bc21cab7db8c0dd2a753f72e)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b0e80110fc74531c02b74b402dbc244a5d93a645bc21cab7db8c0dd2a753f72e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e43a3d331e404adacdbea51d87fda88abd1a1603157877d29c8ffabc9992e25"></a>

## primary.rr_set_group.rr_set.tlsa_record.values — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-585687c64fe7863dd11f33eab30674fb6cda389e728007a92787a434b499b119)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-b6f8693e9a9d359f5ec6cf202b02097d59134e2ec8a3f723f7dfbd92bce497bf"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-8d1a9cddfeef858963f6da0eebaa3f8f0a708aab0626776a503df42314904ede"></a>

## Direct properties — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 3

<a id="canonical-8af5551b904dff0385968098a0abd1622f7ed6659071d202087ce46158436711"></a>

<a id="canonical-c00fa5ff08dec0ab281545d39db571a6dfcb9139e25060e1993d083fed38264e"></a>

## certificate_association_data property — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 4

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4e3d9ca13702c59461ec445aa3d21cdcb25dcf15e11f285a4b8c01b022732c1a"></a>

<a id="canonical-d6922f29cba89dbc0a3920b3f7cbb89a72b8c10af995cc534c7a167077c103f5"></a>

## certificate_usage property — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 5

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-adc08054a239dd89c74e0b5c5a44dd6cc26e3d904bb21bc4d9f1540ea96bc27a"></a>

<a id="canonical-c49cba1ade8008245dd1634bf0d8cbca2d54664cd7a824d54961597c5dca4391"></a>

## matching_type property — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 6

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3143a02aa4d153101881cecd37c3bcb9882c00cbac8ae74db235fdec2ab4c4dd"></a>

<a id="canonical-c7ba55af6177ea1a700db18e30ce1d0fc6fcc4ba977f6383ebdf2f1652c15893"></a>

## selector property — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 7

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-aa9b5adde0f36bcbeebbd51f951eff5a45fa4e19f6682557081adefeb1d571e9"></a>

## Next pages — primary.rr_set_group.rr_set.tlsa_record.values / f5250ea9b10c / 8

- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-585687c64fe7863dd11f33eab30674fb6cda389e728007a92787a434b499b119)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-960100e50468b23cb2c3dfe4db63f18677dc3da79863cd6d4fe4365a0c756607"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94a5bf80105e2c48effee5b56f6ececce672beb83aed80790ee5db7785846bfd"></a>

## primary.rr_set_group.rr_set.txt_record — primary.rr_set_group.rr_set.txt_record / 13787ced441f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-d81f9f35debece6115d30fca7ff3918acaa4efaf215e4358195592b8baaa6304"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="canonical-a4d70bfe09a26e091652112de110f96ed20555dac6f9fde2bd92001d688db5bb"></a>

## Direct properties — primary.rr_set_group.rr_set.txt_record / 13787ced441f / 3

<a id="canonical-6c6c2cbf3843d3bf3e3a7227b711b5eadc390b920136f98cd51f0024ae2eae2e"></a>

<a id="canonical-febe29c3204e1ecf834122834caa0862a814161b64a53d126061ef9c98423dcf"></a>

## name property — primary.rr_set_group.rr_set.txt_record / 13787ced441f / 4

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-b72114627c9def5e0051fc5d06b94503f3200ce2dc6930fd188fdb9ac1ead36c"></a>

<a id="canonical-282f94cc4b0546c254aa2d6dba9071baf94ea43f139c4fbd40af4eeeab5c9c43"></a>

## values property — primary.rr_set_group.rr_set.txt_record / 13787ced441f / 5

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4e89f17443d2053ef4ddd7d369d672dffb89df53103b7f87a5d3a74176309646"></a>

## Next pages — primary.rr_set_group.rr_set.txt_record / 13787ced441f / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-d5c60d18869554ef8153a2525e58f7e65053f5b72480b618ef3dce844ed7aaa9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80ef68b7f7d6a928f52b49ed801d6110b6dbc5bd2a7f377cb20c2ec6f6879a8f"></a>

## primary.soa_parameters — primary.soa_parameters / 4844449d539a / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- primary.soa_parameters

<a id="canonical-fe80e3d93ec4f76e2a897015f27fef777246fbaa833e06034c70fe3b435e7d18"></a>

Type: `"single"`. Computed.

Configuration parameter for soa parameters.

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

<a id="canonical-50e6dd5dd2799189effd8348f34dc2eb85ced57d76433b761df9300352136359"></a>

## Direct properties — primary.soa_parameters / 4844449d539a / 3

<a id="canonical-613c0cd5abcd80bd093c0b89db6960865bb89714cb9735c9f5d9f52402b55bb2"></a>

<a id="canonical-19206e6243ca1aab9a217a1ddafec56250983e929eb4ebc3cfc7be6760fd1b21"></a>

## expire property — primary.soa_parameters / 4844449d539a / 4

Type: `"number"`. Computed.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-d0af71b88d27ec4abd84e8e2822c845a947cd66de1dc5c8b18a5f410e85c7c9a"></a>

<a id="canonical-ad18629419332e7ae901f8365d28b6f130fef034df76f7b9bdb633d39e5d664a"></a>

## negative_ttl property — primary.soa_parameters / 4844449d539a / 5

Type: `"number"`. Computed.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-d848b58cee658e8309c96991a11a953a280ceaa3b4d37c8862fbd5a71d7dd72c"></a>

<a id="canonical-2b4722da29063def5441c6ee19044eceba6f1ad23424805a3f86954953b45bba"></a>

## refresh property — primary.soa_parameters / 4844449d539a / 6

Type: `"number"`. Computed.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-1e5c9aa98707310ae96be7050f5c07a873e5ba4232e28acf9af45879bb53ed15"></a>

<a id="canonical-da2c5dcf49f6181c7cd7f9a527dd686f21d0349f39e033408a6f75fadb71d1b0"></a>

## retry property — primary.soa_parameters / 4844449d539a / 7

Type: `"number"`. Computed.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-e76276a227d1dc59e27e967b12de72c80bb59a9579e0af2044e73ab51a114991"></a>

<a id="canonical-06e26223d15b5a3051f27f35a6f7eb0391607576ccd2c4c670574145960d8b0b"></a>

## ttl property — primary.soa_parameters / 4844449d539a / 8

Type: `"number"`. Computed.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-d6f3c90d844b60f1c4f6b78cab5660fa393d2326232b5b92cef8a154d885962c"></a>

## Next pages — primary.soa_parameters / 4844449d539a / 9

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c502fe4c64a6814a2b19ff8263c4a9d6d3dc9e819d90bc7dae608ab5f97fd541"></a>

## secondary — secondary / d659ba2ef962 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- secondary

<a id="canonical-f7050a5d3f5c4338703440ab6df7f1041a87d40bcae52230201e1755f8c6f13b"></a>

Type: `"single"`. Computed.

SecondaryDNSCreateSpecType.

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

<a id="canonical-3754e6ff602d42e3feb01fea8bac0934f6386c2e140e24fef1ecfde3da3a13ff"></a>

## Direct properties — secondary / d659ba2ef962 / 3

<a id="canonical-adb4c13cb647a05efcc58c669e59d1dbf1cf12d76838b2c9797f67a814c6d80a"></a>

<a id="canonical-edcbc79e9ecba1264a3b0de888506ee7ccaf0c3b3dfd0d44b2c190c009ae007f"></a>

## primary_servers property — secondary / d659ba2ef962 / 4

Type: `["list", "string"]`. Computed.

Configuration parameter for primary servers.

Upstream description:

Configuration parameter for primary servers

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fcaf440114d560fdc74e768093e7a4e486ec2d00798f0be5a02e14248aab575e"></a>

<a id="canonical-ac0fa8a6701d5fb5943bbd932a19daee14753f9523a552576b7dd41819b5f759"></a>

## tsig_key_algorithm property — secondary / d659ba2ef962 / 5

Type: `"string"`. Computed.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Upstream description:

TSIG key value must be compatible with the specified algorithm

&#8203;- UNDEFINED: UNDEFINED

&#8203;- HMAC\_MD5: HMAC\_MD5

&#8203;- HMAC\_SHA1: HMAC\_SHA1

&#8203;- HMAC\_SHA224: HMAC\_SHA224

&#8203;- HMAC\_SHA256: HMAC\_SHA256

&#8203;- HMAC\_SHA384: HMAC\_SHA384

&#8203;- HMAC\_SHA512: HMAC\_SHA512.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8e26eaa9fdb935cd2a7650cc159f419518ba1e13e8efa1f46b3ea63e50286979"></a>

<a id="canonical-8e34e1d8c8cce059ba4d1dc88f3a21e0f52a5112afb636838b0a4a5e03e7509d"></a>

## tsig_key_name property — secondary / d659ba2ef962 / 6

Type: `"string"`. Computed.

TSIG key name as used in TSIG protocol extension.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd): complete subsection reference.

<a id="canonical-3f07f1b362835c156a484a1ccb6e3a8cf6fef5c691cae7cfb25ec364fcd4e67e"></a>

## Next pages — secondary / d659ba2ef962 / 7

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ac17e12244ed232b1f334662609885b31dc496a20f531146398d17a2f86f233"></a>

## secondary.tsig_key_value — secondary.tsig_key_value / 02ab5cf691ab / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb)
- secondary.tsig_key_value

<a id="canonical-d1bfbfb13633c10816803bd5ccc8f9de6d3de5a269ec5f6fde448d683a3f8dd9"></a>

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

<a id="canonical-0e6f8b4ccfc772b56a33e26fde33d5d616f6978911cfb70751a90f4f1ab0db8e"></a>

## Direct properties — secondary.tsig_key_value / 02ab5cf691ab / 3

- [blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-5efc9c4ffdbea3a7a71521a458c090dac8331fe2b75c1e1351a2520b1381ffc3): complete subsection reference.

- [clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-3b5fddb4fb001bc6fce1fdcc5fb617b0af7ddcb5bfc4b1a74b851717a9e69c42): complete subsection reference.

<a id="canonical-ffe9547ae66a133be7311d4b91cb68c8c0e413ae3da9415ceb9ee59afd8ca535"></a>

## Next pages — secondary.tsig_key_value / 02ab5cf691ab / 4

- [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-5efc9c4ffdbea3a7a71521a458c090dac8331fe2b75c1e1351a2520b1381ffc3)
- [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-3b5fddb4fb001bc6fce1fdcc5fb617b0af7ddcb5bfc4b1a74b851717a9e69c42)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-5efc9c4ffdbea3a7a71521a458c090dac8331fe2b75c1e1351a2520b1381ffc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-819241489f72c3d829f4b9906b35bb8dbd87214cb156df821cc6fa93bc6be6d2"></a>

## secondary.tsig_key_value.blindfold_secret_info — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-fa2710e9450051dae9937dfa4495112c09e86e12f1aa21f53cde58d9f4e3bc37"></a>

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

<a id="canonical-2bcf5396e9067d110d4f22a022b36427919228f820fb59388ba94b8873335072"></a>

## Direct properties — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 3

<a id="canonical-a420e74677c48916ab415030785316c7195fa4b7146ae216b19d6dc2da002aa0"></a>

<a id="canonical-5bc66891aa97880ea5d48628f605d6b7c8fdaf55509a85702264d656db35c680"></a>

## decryption_provider property — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 4

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

<a id="canonical-70dcd13d65ea27b23ba5fff19893a998b8cc4b9af289ad0c31de510950fd5129"></a>

<a id="canonical-967fd4f00dc8f2f0e895a422212782ca3eb1211e42b9dc4f0d6d62761261219e"></a>

## location property — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 5

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

<a id="canonical-7ee6128e259a9c3e8ad1a1eca5bcc5e3281c02e05fa705839a4c6ecb3a34ad86"></a>

<a id="canonical-1bc743e9300eaa531c6941d5482a508aaf358f8aa9af9f9ceb4c9233f71f48cd"></a>

## store_provider property — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 6

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

<a id="canonical-0be458fcd07dafb38f7ad9368c84260494815540beaceec6f97e111c5ca1b2cf"></a>

## Next pages — secondary.tsig_key_value.blindfold_secret_info / 2070eb84bf50 / 7

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-3b5fddb4fb001bc6fce1fdcc5fb617b0af7ddcb5bfc4b1a74b851717a9e69c42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1b58827be2d9969039dab66cf58e9e8cbae8ca319c618f4c2b68cd7b8ce7635"></a>

## secondary.tsig_key_value.clear_secret_info — secondary.tsig_key_value.clear_secret_info / 13fe7e5caefd / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-029c1f7b6041911bb160cbeaa3f1c9be7dca710a7c9edee284e6d8f414ac642a"></a>

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

<a id="canonical-e92712de4fbf750a1c5b67743a68a40facf169b96ea64f9421a07d0687318cc0"></a>

## Direct properties — secondary.tsig_key_value.clear_secret_info / 13fe7e5caefd / 3

<a id="canonical-c6270c5988bc8fb5c0828e1b800f32cbe2f72d1fad7f4f1ff2e65c30f030c355"></a>

<a id="canonical-ad72ecb7ba4a7e9e4da5c4a79529260e8ca2e2786896826ea215bafc1dfbb6db"></a>

## provider_ref property — secondary.tsig_key_value.clear_secret_info / 13fe7e5caefd / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7c6e8732918197fdc42895dafd2e5214ac78dd901130c435e7f5e515d34d7cd7"></a>

<a id="canonical-b1ad1531935bb4af303d59a57636252bb5301899eb2c801a7a9ad999b3147953"></a>

## url property — secondary.tsig_key_value.clear_secret_info / 13fe7e5caefd / 5

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

<a id="canonical-7c380fa765de4f9e2ee843d30eb6ad7a5a83566c7785c78ab44b7bc181c38934"></a>

## Next pages — secondary.tsig_key_value.clear_secret_info / 13fe7e5caefd / 6

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-20c305a687a5168735ab0080db62b7b7534583b9464101e7d57ae1f9105289fd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
