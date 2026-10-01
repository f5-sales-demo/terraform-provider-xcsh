---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-6e9f1afc186c838e91833a59a07fbf55cd27284ce771c9623739d77ecb031daa"></a>

## origin_servers.consul_service.outside_network — origin_servers.consul_service.outside_network / a7fb659f96f5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- origin_servers.consul_service.outside_network

<a id="canonical-81ea009a6dc5bfba306456579588e2349076791a29016f5bf23436335a4a4a97"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-3d7228be45078f88df0851a0ad717c421ad54d91b8a670220602adb9a70f7da1"></a>

## Direct properties — origin_servers.consul_service.outside_network / a7fb659f96f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c9191a5fda8107195529d85ed695944d266fdbbe392353bdf2c99f749e07c5a"></a>

## Next pages — origin_servers.consul_service.outside_network / a7fb659f96f5 / 4

- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1985ff56c48fc30015e0af181abc882b25d246cd44cd685bbfadcca1acf8ac2e"></a>

## origin_servers.consul_service.site_locator — origin_servers.consul_service.site_locator / 44eed66389d9 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- origin_servers.consul_service.site_locator

<a id="canonical-a10a3225aa809327b3dbf2f473b0d5b29becd296c8380b6f08ecc37ba57f5571"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-b0fa428a9b4eb8886fec192449a2379aabaeff26c08c6f88bb11cbd2b18b1b30"></a>

## Direct properties — origin_servers.consul_service.site_locator / 44eed66389d9 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-e1533b74c51e415e63f28c67c849092450007cff3bacc44bd1f6ee1aa01a997b): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-05d40b9aa504b46fbc5bfb39450dccda0bfd042d59e1b6e5bd416fb22e18cd2d): complete subsection reference.

<a id="canonical-8501189f011ccc8039eb467fbb6dbecadd31d58e55dfe626c13274c5def99dc7"></a>

## Next pages — origin_servers.consul_service.site_locator / 44eed66389d9 / 4

- [origin_servers.consul_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-e1533b74c51e415e63f28c67c849092450007cff3bacc44bd1f6ee1aa01a997b)
- [origin_servers.consul_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-05d40b9aa504b46fbc5bfb39450dccda0bfd042d59e1b6e5bd416fb22e18cd2d)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-e1533b74c51e415e63f28c67c849092450007cff3bacc44bd1f6ee1aa01a997b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f88757ed544728696ac87c1b3134fb2e764d28f2010f61e62b1f668db5a1c167"></a>

## origin_servers.consul_service.site_locator.site — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96)
- origin_servers.consul_service.site_locator.site

<a id="canonical-770fbe12b39f6cacf251e79d542924ca63d04ceaecb7ca3efc8c2ae22e175049"></a>

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

<a id="canonical-59cea4e6007c152de971f995a7b03c3531cf3ac113c81884fb478f4985959a48"></a>

## Direct properties — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 3

<a id="canonical-c06af92907fb6d738166468b6cf7c8f815ec6a8ed41c0f9813a751738e698e8c"></a>

<a id="canonical-b411cdb9579295d03f1aca539561fd02251afacc4eb4a30ef6164a78e500a173"></a>

## name property — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 4

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

<a id="canonical-031295adf16a424db5ff2e932b3fc109c9a548a051caa82d0f838fdd85aecb4c"></a>

<a id="canonical-66dc37f5570378983c6e7e916f5c72e998798e0355c741ec02ece0f008a51bb0"></a>

## namespace property — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 5

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

<a id="canonical-e59a2319e0d769a2ee4927ccc39e1d76b1e565d0940b619783560bf6070a61e6"></a>

<a id="canonical-18e2a48fb796e39bb75ffeb2c03f97d20afddd75fee671c5e6dfd76d6159dcee"></a>

## tenant property — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 6

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

<a id="canonical-d3b651b9d225b23dcf2b55bd2190e316d4fe14a67daa999dc01782e446a83599"></a>

## Next pages — origin_servers.consul_service.site_locator.site / 1ec47103ff13 / 7

- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-05d40b9aa504b46fbc5bfb39450dccda0bfd042d59e1b6e5bd416fb22e18cd2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ea04b6a4c1d50d815caaeb0217e28d35835b40e0fdc513f44893964115547c"></a>

## origin_servers.consul_service.site_locator.virtual_site — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96)
- origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-08cbe9c225ba413e24b0be53c6a48fefc9e49dfbbd7466d015c99b01ba662ccb"></a>

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

<a id="canonical-ad632cbdaed609d2408a6cae477c8ce1eb3637f383b1e18db55f153bf4c801e8"></a>

## Direct properties — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 3

<a id="canonical-f5b5c6745711fc4dd42fc696058518c1a667091c4a3be6836b1c12733dbd117c"></a>

<a id="canonical-5ff4b1159c09112d46af64a6286b813e9dd037e7bd3c22cb98234365bf6098c0"></a>

## name property — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 4

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

<a id="canonical-3bd4c7f24836d8b7b07ffe63c647e1951105326af2d25bab328147f2c6258cc4"></a>

<a id="canonical-f3544710fce3d275df7ac1b559e47a173530ab9984f6a7a7187d6829ecc66d4a"></a>

## namespace property — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 5

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

<a id="canonical-8d142dbe166a658889dc26686d5ffa32c24274a78b7dc2725774be568d0375d3"></a>

<a id="canonical-6e5c4a70e95ea8759bf53250cd8040df70a6d5d56aa6a3d076165e7565ea89cb"></a>

## tenant property — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 6

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

<a id="canonical-6a4dc29e49e56b466015411561acd66933bb31959b87a4bea8577e53267e9abb"></a>

## Next pages — origin_servers.consul_service.site_locator.virtual_site / 6e9e0e8d65ac / 7

- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-28bae35109f9a8ba00109ccad831e7e9fe184e62e70e7f54cad54b4a0cb30a96)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e110768cd5fd5e699e6e9b754ece32f982b42b163ce34555cf3cb49447111ee1"></a>

## origin_servers.consul_service.snat_pool — origin_servers.consul_service.snat_pool / 8d4dbe251fc8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- origin_servers.consul_service.snat_pool

<a id="canonical-82d90332c22aed68f17bd76604d3856d579c476ccb2e472aeabf12dee32a52a2"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-e3fc0281d5259f7ce572cbb36d2c133bb857edf72132bb35c6a3fbdf934c77d0"></a>

## Direct properties — origin_servers.consul_service.snat_pool / 8d4dbe251fc8 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-a6c8b119c6d6a38abd7818884620ebc59a6b4d44f60c4d9a4c161979468f1eb0): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-7e579120cb6ec30128eb3c1b8e22873ebea95727e3a670f83e94bee36830764f): complete subsection reference.

<a id="canonical-54aadf70ee5e51a945a49eb18ca98fc16847431308be6a7f59a42df3d93dceac"></a>

## Next pages — origin_servers.consul_service.snat_pool / 8d4dbe251fc8 / 4

- [origin_servers.consul_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-a6c8b119c6d6a38abd7818884620ebc59a6b4d44f60c4d9a4c161979468f1eb0)
- [origin_servers.consul_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-7e579120cb6ec30128eb3c1b8e22873ebea95727e3a670f83e94bee36830764f)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-a6c8b119c6d6a38abd7818884620ebc59a6b4d44f60c4d9a4c161979468f1eb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-216efd4cb71e012b889bcc33feea667f264dcf7ff33947bd7d8bf8e84f4fd029"></a>

## origin_servers.consul_service.snat_pool.no_snat_pool — origin_servers.consul_service.snat_pool.no_snat_pool / d0dbca56f862 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9)
- origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-e284b4b38610bf18f682b4b0f93c66f640e1ee35a32be96eecddc615cd96e794"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-09b7923377b0c0884f1f27ae173ee5cb41146079093a4804d94f04e9f0f4fd04"></a>

## Direct properties — origin_servers.consul_service.snat_pool.no_snat_pool / d0dbca56f862 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b052d48cbbe68df97f0a016e7eec527f57464f1f050df868aec99a0db80bbe26"></a>

## Next pages — origin_servers.consul_service.snat_pool.no_snat_pool / d0dbca56f862 / 4

- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7e579120cb6ec30128eb3c1b8e22873ebea95727e3a670f83e94bee36830764f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74dae145575ef60227a9ebf03cfd5f7ada527930125c2dda2ad609787c36b40e"></a>

## origin_servers.consul_service.snat_pool.snat_pool — origin_servers.consul_service.snat_pool.snat_pool / 2b2617c4bb65 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-ae80c24c1da48cc0070d672299870b8d1715f4294c037096d28e282a16262811)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9)
- origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-caa372535c0c9df868fc35fc3e5bbc07792b4148201a2f76838be632e3d22cfe"></a>

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

<a id="canonical-8a4292aadbd3621c2b81a98213b0973a8485908f51dc5b40b16bb853668234b4"></a>

## Direct properties — origin_servers.consul_service.snat_pool.snat_pool / 2b2617c4bb65 / 3

<a id="canonical-247bb92d9c8b81bc539b01e573827c038572521197379c17977d5b93d786722a"></a>

<a id="canonical-14a836a0ac7b25bdbd3fd06c2c39bdfc56f08053f051002fa6717bb6c8ec668e"></a>

## prefixes property — origin_servers.consul_service.snat_pool.snat_pool / 2b2617c4bb65 / 4

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

<a id="canonical-35829c2364b8b5969f7991f55da9ae7e080df86f4290d52ce676fd81f0cc1793"></a>

## Next pages — origin_servers.consul_service.snat_pool.snat_pool / 2b2617c4bb65 / 5

- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-b87929f642efe394b3364544b3815d73d3dfe072976b77f02f4c8daa2e4b10c9)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-db084b1865c022fee81a4c1d788a9ac1e8c81ee38a27f7ce6dae2ea733e679ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7994ed83350bff41f3305de33e4d17cdd730367769a2b98a8c2934ad76516a2"></a>

## origin_servers.custom_endpoint_object — origin_servers.custom_endpoint_object / 7aa4973ad008 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.custom_endpoint_object

<a id="canonical-ff272def9ed381d6177b5f7d9b90bc088cc5c5f3f480f000ecce5dd8e8d0dd23"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

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

<a id="canonical-8b49dfdb88dbc7ad09313cc1db0961f9182bd3c6bd27bae91a3e906799fe43ee"></a>

## Direct properties — origin_servers.custom_endpoint_object / 7aa4973ad008 / 3

- [endpoint](data-sources--origin_pool--reference--group-002.md#canonical-a0a9c2b2b37645941bd9f2b0cfe76df86db69cff55ab59f36a5c5aca9d8e619d): complete subsection reference.

<a id="canonical-01fb1c1a90b7cf856815188b95cced00dabe6caadc3dc4573e792e941c2113bc"></a>

## Next pages — origin_servers.custom_endpoint_object / 7aa4973ad008 / 4

- [origin_servers.custom_endpoint_object.endpoint](data-sources--origin_pool--reference--group-002.md#canonical-a0a9c2b2b37645941bd9f2b0cfe76df86db69cff55ab59f36a5c5aca9d8e619d)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-a0a9c2b2b37645941bd9f2b0cfe76df86db69cff55ab59f36a5c5aca9d8e619d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba54195805ef3df72a36915e5ccafc5d525fa4933e3546a992459a459bc3e38e"></a>

## origin_servers.custom_endpoint_object.endpoint — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-db084b1865c022fee81a4c1d788a9ac1e8c81ee38a27f7ce6dae2ea733e679ad)
- origin_servers.custom_endpoint_object.endpoint

<a id="canonical-6753b96d3fede410153cbb36986f18bf4e520de52b22f425ba3c88f9ff4257e1"></a>

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

<a id="canonical-78bd8ebdc2610e9d6f16aa237e1abda231399d071a276d11222b709c6645c5b9"></a>

## Direct properties — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 3

<a id="canonical-ddebee53ef7d2931ef60589f026d50e8bab2edbd618712151cb89f87ee3405bb"></a>

<a id="canonical-806ab3e8541a329953eb3cab935ea8166643b972dbb391e2d860200f0ef1ef58"></a>

## name property — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 4

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

<a id="canonical-38680cd0a56b05d06d016b835056fb84303ffbbfd8f2cddbab426fb5f8a6e89a"></a>

<a id="canonical-01b95e8e08220e449386104f926c9bcd9ddf9c28b03e86c9428a6edeab02eb01"></a>

## namespace property — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 5

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

<a id="canonical-df322039f1103e3dd4e8adca81b781d69e088ea2b120ad35c38612513521ae1a"></a>

<a id="canonical-39189fb00c0037880a6a04fbad54477c6f7afbc29146fd0b92e86e31127bdb8f"></a>

## tenant property — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 6

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

<a id="canonical-35b0a6126ea05af9bce7366ce1e1a68e2d364dc62e5a6389a7d6bcebc5134a19"></a>

## Next pages — origin_servers.custom_endpoint_object.endpoint / fe37b911e52f / 7

- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-db084b1865c022fee81a4c1d788a9ac1e8c81ee38a27f7ce6dae2ea733e679ad)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e23805e3f524e8040828887485e2560f9d69d4f8095d0055a84ff6968d39401f"></a>

## origin_servers.k8s_service — origin_servers.k8s_service / aa63cb3e11ce / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.k8s_service

<a id="canonical-6bbc8291e89644837af72bf4e4db993e8ade6f2bce94ccea4b613fe057f9e1f6"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-72ef22419f7cbcf8a4433422e98ef6843cc649dfaa5b8670bcfd38da91a5bf2c"></a>

## Direct properties — origin_servers.k8s_service / aa63cb3e11ce / 3

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-9903c0ea8c97fb675da9cb32d733db91da5eddb7a28959e76a777b5cd559a8a2): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-84bf4df3ae766ac9497ffa320d0fd61d538432e4edbf8247f6c2d6aacee29163): complete subsection reference.

<a id="canonical-8f512e6d85f32ea26f508afca96a30933008f04df84b2f2f3e091bc9c3ece285"></a>

<a id="canonical-9b0a342b5c7ea60f1263d3c30a4dfd545bacc59768f4d1d244827862807c7849"></a>

## protocol property — origin_servers.k8s_service / aa63cb3e11ce / 4

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a33f8a81932d56ea9bbe2fca72c4ccd70ed1535fed8ef592b75bc4c504b57987"></a>

<a id="canonical-30c17c2cf7a7c55b603228331f0d7b195be39291abf45279431da7f64026ca8b"></a>

## service_name property — origin_servers.k8s_service / aa63cb3e11ce / 5

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484): complete subsection reference.

- [vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-c3e2846eaa6820051abe6ad0179f8c65f9eb0aea6f756e50d1522ad6a95c0ac3): complete subsection reference.

<a id="canonical-34c1e29e8b04d3221bcb17afabd57ea77035ce120a58b19eca5154ea2590ec31"></a>

## Next pages — origin_servers.k8s_service / aa63cb3e11ce / 6

- [origin_servers.k8s_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-9903c0ea8c97fb675da9cb32d733db91da5eddb7a28959e76a777b5cd559a8a2)
- [origin_servers.k8s_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-84bf4df3ae766ac9497ffa320d0fd61d538432e4edbf8247f6c2d6aacee29163)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484)
- [origin_servers.k8s_service.vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-c3e2846eaa6820051abe6ad0179f8c65f9eb0aea6f756e50d1522ad6a95c0ac3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-9903c0ea8c97fb675da9cb32d733db91da5eddb7a28959e76a777b5cd559a8a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-452fd7ed2101223648fd7f7dc06dc7537f822c2794f9a1284c2f32c09c1d2ae8"></a>

## origin_servers.k8s_service.inside_network — origin_servers.k8s_service.inside_network / 0ba3798c49f1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- origin_servers.k8s_service.inside_network

<a id="canonical-a3d7d40aeab39cec9818d99d6036c8b3313a717b2f0e843f1545adbd70aacc9e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-d4e20a988f240ab456a7157743cc833c5dca6b7e0c020c3e1239006cc52e6f82"></a>

## Direct properties — origin_servers.k8s_service.inside_network / 0ba3798c49f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d776769b5a132f0cc67c50a65f7a73a6d3450328e8abd236a2bf39c09c5a8ad"></a>

## Next pages — origin_servers.k8s_service.inside_network / 0ba3798c49f1 / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-84bf4df3ae766ac9497ffa320d0fd61d538432e4edbf8247f6c2d6aacee29163"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-939919f080d31c7447b4ba0ba75c0ab5675cc634dc6ff17208c11778b49e152c"></a>

## origin_servers.k8s_service.outside_network — origin_servers.k8s_service.outside_network / 1d3abebabe0e / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- origin_servers.k8s_service.outside_network

<a id="canonical-c626bf8dbd344e7a792c172a5472dd6a4cca2e4804a3abe420080d770c6427ad"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-1a54a2a0ea44ff7ac1aee97aea727d75507a6b8a83da3fd6846b095a5aa2af67"></a>

## Direct properties — origin_servers.k8s_service.outside_network / 1d3abebabe0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5aa5af4ed5ba61ec421e0da332a73dc6a24c68974e47b0b57d9f0e570851bc9"></a>

## Next pages — origin_servers.k8s_service.outside_network / 1d3abebabe0e / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e9b07b040b41b734c740fab33a26c8cc79984885d23bfdc3a3eff79efc7e684"></a>

## origin_servers.k8s_service.site_locator — origin_servers.k8s_service.site_locator / 62e8e0a2bc78 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- origin_servers.k8s_service.site_locator

<a id="canonical-0a5b4f0a144525708ab23af749c2d49078da616928d6619809e499af5a2dcb61"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-2ee206f70abbddd3328c1b7326bca4340141ed79c031e79080ba84fab0662a4a"></a>

## Direct properties — origin_servers.k8s_service.site_locator / 62e8e0a2bc78 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-0b69f9f80e624eab6f0ac35d78e582078eae2cf16865716798d68e63e6d088a4): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-ce6e4093995c9f040f2d308b2cc1c73f8d3498f61f1654fac263c3fcde69974a): complete subsection reference.

<a id="canonical-858ac2ac436aa70424c2c44163158748b972e59e973b7f1104a6404ac617ec46"></a>

## Next pages — origin_servers.k8s_service.site_locator / 62e8e0a2bc78 / 4

- [origin_servers.k8s_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-0b69f9f80e624eab6f0ac35d78e582078eae2cf16865716798d68e63e6d088a4)
- [origin_servers.k8s_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-ce6e4093995c9f040f2d308b2cc1c73f8d3498f61f1654fac263c3fcde69974a)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-0b69f9f80e624eab6f0ac35d78e582078eae2cf16865716798d68e63e6d088a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-704475283ab8d73a42e49a22783c04b121ef24c26d5f0d120333db0bb06f901c"></a>

## origin_servers.k8s_service.site_locator.site — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-5eb033314d7e24a1d5016bdc9504fe0f3b76572fadd0a58328dabfba237502b5"></a>

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

<a id="canonical-ad29195e5d235b48767327b1ce8626e62cec1cce071c50d5ae7ad9121465c377"></a>

## Direct properties — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 3

<a id="canonical-88866a55c5f3712389880e69537e79185b272f1e0c69fc04565dc3ed1e9ad058"></a>

<a id="canonical-7995ddd63b638b409fcad7709dfdfd4952c347687dbba482d3dff90acb304ce3"></a>

## name property — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 4

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

<a id="canonical-22be4be707b1debb1916b3aac6629032180a22891c0755671b6c86fa600cba39"></a>

<a id="canonical-a5cb3cb6ab3af6f1e753c993e759c0d4c4deb12c2080714d27a2a699aa982d85"></a>

## namespace property — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 5

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

<a id="canonical-aa1b7d240a38740465b3c62155cb1eb1d76a1f69c12cb0e44d6602035dafffe5"></a>

<a id="canonical-7cc0134bca06f6177e51d3be5aa95f8a15709428ac2ec899e753219dcc46c150"></a>

## tenant property — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 6

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

<a id="canonical-a6a97b058417e8d3d87ec3f54604162701ad7ccba883c551141114e11faa9f8f"></a>

## Next pages — origin_servers.k8s_service.site_locator.site / 3d8aa9f73867 / 7

- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-ce6e4093995c9f040f2d308b2cc1c73f8d3498f61f1654fac263c3fcde69974a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee43cad039431f204c976dc5db8ce0d0a696cc9362e03771867a7f9f90345fdc"></a>

## origin_servers.k8s_service.site_locator.virtual_site — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-343baa270be35fdaf017735efe0ff309e28fccd478ba56ef97b9ff9ab259e4e0"></a>

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

<a id="canonical-f6179050f86f38a561afd5b78975aaa6884df6e22550dd1a36a9789657af651b"></a>

## Direct properties — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 3

<a id="canonical-d964e1b4fba535b3cf49891676171f534530c157f790bc01f564d247ff36a01d"></a>

<a id="canonical-02fcec24b9e2b08181ad4bbf4b9cf7a8550d6dac4c742aa59c9c04fde6fc69d7"></a>

## name property — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 4

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

<a id="canonical-d001a426353b052d92d11ff0d187b72dbc66b55388c578cd5a36d28867285cee"></a>

<a id="canonical-1fc083b3c327f05081bf54a37ca76429036764a0cd7fcd1ae24b83efd9fd9540"></a>

## namespace property — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 5

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

<a id="canonical-4c76abe91de8873a4725e2159f8bb07a51a6e65f899c998d8f71cb432ecf6c22"></a>

<a id="canonical-fee11c22e1067d8ef32e68a25e0d4ce926f4fbbbc161bc875d803b53b8bf2bf8"></a>

## tenant property — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 6

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

<a id="canonical-ce8e19487beb578bbe9fc99fe33c20ecc710bb41b5d81cf9609c6391bd9a1b88"></a>

## Next pages — origin_servers.k8s_service.site_locator.virtual_site / 6f789cded9ab / 7

- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-e7cbb6e8a0effc5b392dafac20434518b2db7d8dec59d9a3f124542ac00f681b)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20278dabaca1136dd2145de16319f12e4620040342ed29049ad6e3eaf08cc05"></a>

## origin_servers.k8s_service.snat_pool — origin_servers.k8s_service.snat_pool / ed7918f65176 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- origin_servers.k8s_service.snat_pool

<a id="canonical-b079615f41a92b55493c597c479e2fdc4df23a9a4da404542e6d929695c25382"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-6a39046cf173b193e3ae1f6ee20cdc9b98adfab794cb5233147ca2ff38a798d6"></a>

## Direct properties — origin_servers.k8s_service.snat_pool / ed7918f65176 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-89f70c4021589107cc600d57e73e98fe7a8cf24a008a9475a01e6a6175102f39): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-7dc37a05cbf18789c10511e4dc0f65815a654c0ae30ceb97b109543daf38e56c): complete subsection reference.

<a id="canonical-606bfd3bb25e3de6e3a36009c8faa564645c2ff2dd3e9a20774fcbf8793dba0c"></a>

## Next pages — origin_servers.k8s_service.snat_pool / ed7918f65176 / 4

- [origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-89f70c4021589107cc600d57e73e98fe7a8cf24a008a9475a01e6a6175102f39)
- [origin_servers.k8s_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-7dc37a05cbf18789c10511e4dc0f65815a654c0ae30ceb97b109543daf38e56c)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-89f70c4021589107cc600d57e73e98fe7a8cf24a008a9475a01e6a6175102f39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee3be5c3c5aa30e9f4617c190c82a063d82bb3c16e73488e3a073478476651ab"></a>

## origin_servers.k8s_service.snat_pool.no_snat_pool — origin_servers.k8s_service.snat_pool.no_snat_pool / 925f5217096d / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-967f2dda11227aba91a2d6b47ced624d8c83d3c84d1f740187b9826b38988dfa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-66f9077758fdd740636ee223be6d8606c146ddb24dd026a5b1b6b98ccc7e4c81"></a>

## Direct properties — origin_servers.k8s_service.snat_pool.no_snat_pool / 925f5217096d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c22d3a3ba751e4f406c5c0086dd3f1bfee43fa78c6825326fa3cc81ef673996b"></a>

## Next pages — origin_servers.k8s_service.snat_pool.no_snat_pool / 925f5217096d / 4

- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7dc37a05cbf18789c10511e4dc0f65815a654c0ae30ceb97b109543daf38e56c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9223c87a12cfd6c1fd4206fbc8d4393e1a9a779f996edf03afc8902d97b4c6a9"></a>

## origin_servers.k8s_service.snat_pool.snat_pool — origin_servers.k8s_service.snat_pool.snat_pool / 35905d8cd229 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-31ecf11385680a6d231dbec295a7e18cbb1ff6fc33c6d2e2baeac84c5a680f74"></a>

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

<a id="canonical-bb4dabf411cbd4e9c73f50b489e1a41151d29c820a43e1c258278c5f7ff39a1c"></a>

## Direct properties — origin_servers.k8s_service.snat_pool.snat_pool / 35905d8cd229 / 3

<a id="canonical-60c6f910f75a32f6904d9bb5a2106a51bce0d304cda26872dea717824c281909"></a>

<a id="canonical-2133241f35c1a60ee1f26a7ce55e93a5133590cd39f35189f6b06f102a7dcd87"></a>

## prefixes property — origin_servers.k8s_service.snat_pool.snat_pool / 35905d8cd229 / 4

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

<a id="canonical-11a199ef565c8da089216d6c737b58230f5f9ce024fa4cb02ed0195d5a5ee015"></a>

## Next pages — origin_servers.k8s_service.snat_pool.snat_pool / 35905d8cd229 / 5

- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-43844aa15153860664348671b56f4602f7528e86a011fcd1d8a6557634034484)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-c3e2846eaa6820051abe6ad0179f8c65f9eb0aea6f756e50d1522ad6a95c0ac3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3254f107956be976a7ae1ed1ba935da892bd8d5e51a0b784e4c36ef009543002"></a>

## origin_servers.k8s_service.vk8s_networks — origin_servers.k8s_service.vk8s_networks / 565bf3589ea2 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-59e29f02dc7ded37db556319a022b3f6f28a0bf16342300a97fef5760af24395"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

<a id="canonical-8d2c6000a921488aad3a432d0c90ef5f9f77ac6ae1f3ec7522c754b26a7f65f7"></a>

## Direct properties — origin_servers.k8s_service.vk8s_networks / 565bf3589ea2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a4f77bf44d9430d0356cb7c6137c24cf5175fb65d42329ca92c0e3dc479c133"></a>

## Next pages — origin_servers.k8s_service.vk8s_networks / 565bf3589ea2 / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-b0905e5d192363ee195c8619ca622ef62b1f4fcb811edd2202e801a687bbfee1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d03fa4fabd6bd41327fbda47f3248eaadefc2a94c2a89fd6059040eff8caf34"></a>

## origin_servers.private_ip — origin_servers.private_ip / 4e0c28ee0a27 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.private_ip

<a id="canonical-7b78e4b127ca4a304866f4cb48d6a55649dd1d975060949e6faf28ad1e6f74f9"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-e16b8850ae0c37b67267a4e427ecffead96973854872ccb5eb357d403fe930eb"></a>

## Direct properties — origin_servers.private_ip / 4e0c28ee0a27 / 3

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-4be9fc36e6eb8369c02e9108ba05a1898523802f771e91b89846ad9b0c47c90a): complete subsection reference.

<a id="canonical-38655241a32d9f14ca608624a90ec0b6d704a1e6a7699c307a1d6521fea7d0fa"></a>

<a id="canonical-8293983d706ac567bbf1fe95a9c8e910415f06b91f9a029a61afd5da2b664ddf"></a>

## ip property — origin_servers.private_ip / 4e0c28ee0a27 / 4

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-3289db012f6e42e21f862fd0ba901a7de6bc2ee5191bfa9ba16eea94e785b7de): complete subsection reference.

- [segment](data-sources--origin_pool--reference--group-002.md#canonical-decdbd4cbcd34f121182ac4dd4bf15582fb072c42b678d2c96966170532246aa): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1): complete subsection reference.

<a id="canonical-495473d39b8e581ac720679e301466a31cd3bed545b0eb1ec1411c729f36ae54"></a>

## Next pages — origin_servers.private_ip / 4e0c28ee0a27 / 5

- [origin_servers.private_ip.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-4be9fc36e6eb8369c02e9108ba05a1898523802f771e91b89846ad9b0c47c90a)
- [origin_servers.private_ip.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-3289db012f6e42e21f862fd0ba901a7de6bc2ee5191bfa9ba16eea94e785b7de)
- [origin_servers.private_ip.segment](data-sources--origin_pool--reference--group-002.md#canonical-decdbd4cbcd34f121182ac4dd4bf15582fb072c42b678d2c96966170532246aa)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-4be9fc36e6eb8369c02e9108ba05a1898523802f771e91b89846ad9b0c47c90a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ae9bb943dbe267996d5c9b9a8081be3f037d8d321972df83d28655f4ddda6cb"></a>

## origin_servers.private_ip.inside_network — origin_servers.private_ip.inside_network / 1bf2da33c7a0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- origin_servers.private_ip.inside_network

<a id="canonical-7efded0ac379797b3a61bb817de8824bfa40d0c8c7fe589acc5ef510790342cc"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-82474d1cd362deeea89f09e3bb8054c80110a2cd350f8599b83079c85f1c94c9"></a>

## Direct properties — origin_servers.private_ip.inside_network / 1bf2da33c7a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f4da7eb8c5ecd52552667b689b7d3c3e4387902cfa9dfd60c25e7955d38775b"></a>

## Next pages — origin_servers.private_ip.inside_network / 1bf2da33c7a0 / 4

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-3289db012f6e42e21f862fd0ba901a7de6bc2ee5191bfa9ba16eea94e785b7de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a98dd2d69cdff6e5395f689be8366abee12d4235a63b00fb20ca7fac254f293f"></a>

## origin_servers.private_ip.outside_network — origin_servers.private_ip.outside_network / 5cbe4daad42c / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- origin_servers.private_ip.outside_network

<a id="canonical-cd3d4f5ee3e20d12fcff65f7235488d512f998f70115727ede576e2cff4813ec"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-c9b59100167db261a11abd1dc8753162b1672d9236f7144ab150a00c5b695177"></a>

## Direct properties — origin_servers.private_ip.outside_network / 5cbe4daad42c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f423f94b9379556e604a943915523ac88082b1fd49f47e4d9442d2b28ba48e8"></a>

## Next pages — origin_servers.private_ip.outside_network / 5cbe4daad42c / 4

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-decdbd4cbcd34f121182ac4dd4bf15582fb072c42b678d2c96966170532246aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c5c6664a10adda5a8084af248c27b2ce047c556ec26b5d34f7cfb18076e1c8d"></a>

## origin_servers.private_ip.segment — origin_servers.private_ip.segment / bb1fc664f902 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- origin_servers.private_ip.segment

<a id="canonical-e154ac5ae6b19c866aeb2fc4b1b542c0b48ec028e136abd954537f989bf65150"></a>

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

<a id="canonical-6dc6e028ffcfee4d022b191568a4d3403752f954fca87835412ea3bf96495188"></a>

## Direct properties — origin_servers.private_ip.segment / bb1fc664f902 / 3

<a id="canonical-59f84e843d0671e19dd02970aae2ad8492fef7dd5cd2e8b70a1c567934a0b6f7"></a>

<a id="canonical-8bf26892da348703a541246dd5ac6a62f9e5c6efa8d0de94c008b6d655d7cbda"></a>

## name property — origin_servers.private_ip.segment / bb1fc664f902 / 4

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

<a id="canonical-f2d908fafd11e8ed432737cb085ca223c5a07fe0af414f817c884ccc22679f85"></a>

<a id="canonical-278f4a74f95f9e4ed6c3d305e1523162e16a6dc6580a11b2c23144a45bc38e10"></a>

## namespace property — origin_servers.private_ip.segment / bb1fc664f902 / 5

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

<a id="canonical-eff8404654d58067d8a85acb3faae66500a0f418298b896dbf0c62345b2af5e6"></a>

<a id="canonical-1d38b3b1452910b28513012bd0a3134eb753a881c526b8cba0e9ddf013c00ff7"></a>

## tenant property — origin_servers.private_ip.segment / bb1fc664f902 / 6

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

<a id="canonical-2e3e6d3185443c74511ed342d13b978bbeceb545ac7515388a59eaeadf44b280"></a>

## Next pages — origin_servers.private_ip.segment / bb1fc664f902 / 7

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-916c42ce67fc865b8e2aa3215d5ffccb95115145c2dddae9b70c64989e6145a3"></a>

## origin_servers.private_ip.site_locator — origin_servers.private_ip.site_locator / 4c103caaa7e1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- origin_servers.private_ip.site_locator

<a id="canonical-5a6dd14d378cb947f1b9d4aa0e7defc1e79ad4250e833007a9a780131c3935f3"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-0f03a12c34f67c03584e1bc5a44b1ad766e59ec907bd2a1dc7914da8eedb2a00"></a>

## Direct properties — origin_servers.private_ip.site_locator / 4c103caaa7e1 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-6adc7c86b0320ed5c3950c1aa3e9ad7bb0ca5590013e2f7a9f8f42c291a754fa): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-c41b566afbe4a193a711b6f5114e21367e47fab9fe423f3272d9e7de6c3f0c6b): complete subsection reference.

<a id="canonical-66d180715a939c949b18b3d1643c5f6467ec12a171e4b38231e9cb8d88ce54f8"></a>

## Next pages — origin_servers.private_ip.site_locator / 4c103caaa7e1 / 4

- [origin_servers.private_ip.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-6adc7c86b0320ed5c3950c1aa3e9ad7bb0ca5590013e2f7a9f8f42c291a754fa)
- [origin_servers.private_ip.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-c41b566afbe4a193a711b6f5114e21367e47fab9fe423f3272d9e7de6c3f0c6b)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-6adc7c86b0320ed5c3950c1aa3e9ad7bb0ca5590013e2f7a9f8f42c291a754fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-973bc8d3ca3feb28cfe6df9c23b8c77ca3b466e89b79d45d2179a5f89a13c790"></a>

## origin_servers.private_ip.site_locator.site — origin_servers.private_ip.site_locator.site / e50bd1109acc / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c)
- origin_servers.private_ip.site_locator.site

<a id="canonical-245205a82bc0d92de0161f7c07a6edfffd724d1ff8a762fd31a42428566fa030"></a>

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

<a id="canonical-4f90e2c4a1dfb220c9e282c9aaed79ace6755ed4f47cbb691c563df2809a9a95"></a>

## Direct properties — origin_servers.private_ip.site_locator.site / e50bd1109acc / 3

<a id="canonical-85fdfd0fd5587e05f583b01f04ddc8bc8658d3c054576e0a6aa4f29b562cb6d2"></a>

<a id="canonical-4d850277629fd59e935a5952b9f268b8cdb62864457986c9a6277b4a9c40322f"></a>

## name property — origin_servers.private_ip.site_locator.site / e50bd1109acc / 4

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

<a id="canonical-6a0b9421518717fd4eedfa962795fdaf0a0e6eeaf63f0ca9c0929eb432e32aa0"></a>

<a id="canonical-ada53ddb2c4476384fecd440f9986e7f6feaf160ce90f92597bfa034948d8c27"></a>

## namespace property — origin_servers.private_ip.site_locator.site / e50bd1109acc / 5

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

<a id="canonical-066e5585ec4e5d77e75f699c4a8090d2789e82103ef6b8f1a968a4bca326af5c"></a>

<a id="canonical-8588e8b436f4de6391f9919b98df911059b8d215164ca19507ded47322ca4e68"></a>

## tenant property — origin_servers.private_ip.site_locator.site / e50bd1109acc / 6

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

<a id="canonical-df801c445fd964114a574b754c95449d4fc2df911f83feabd723ccc79ac56e88"></a>

## Next pages — origin_servers.private_ip.site_locator.site / e50bd1109acc / 7

- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-c41b566afbe4a193a711b6f5114e21367e47fab9fe423f3272d9e7de6c3f0c6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4d2c4b3d1e1eef73908d9e73a423f63936f7030ac4647690d995974c48d39a9"></a>

## origin_servers.private_ip.site_locator.virtual_site — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-baaf62dce8415ebabda8f37b88487db0493f5acf7d4b64df19655eb935c37852"></a>

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

<a id="canonical-aac45f0580c2bcbc5fa2dc58aa575ade37776686788d2e4f2eeef3ba0a96d1c9"></a>

## Direct properties — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 3

<a id="canonical-9aa1aecfc5dee5d476b533a9c9534a212328eca6374451a4f5cff55695d4d0fc"></a>

<a id="canonical-b4b686840389e07c028e65a61e609d89bec8c3098530f42e879545f8fb08aa3a"></a>

## name property — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 4

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

<a id="canonical-c4f2d0fc7e52e2b65eb4b574071c3e4651704dbfe9065b881dfa99f470a1f960"></a>

<a id="canonical-7cca4c7a01e291228c13d524509327e2abfcae93cf29171344271d792247b9f0"></a>

## namespace property — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 5

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

<a id="canonical-1c020a8223b05ac8133256fbdac84cbfe03822058e6e2e7875e2c2cb584d2486"></a>

<a id="canonical-7d841fa61f62f7d6ffd08190300ef12a1b2621d2329be130c85ac6f5190ba67b"></a>

## tenant property — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 6

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

<a id="canonical-be226a1264d7c42bf6937580012807d966bf5176b5810b633cb218ed1d8b9861"></a>

## Next pages — origin_servers.private_ip.site_locator.virtual_site / 43ff32839314 / 7

- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-6ca3906b042de4181b42be54d927c1073a7ef454592c8583984690ae4a30098c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6203828629632dbac8950515a691bce4a96e9674c8996b483ac55cb5cbf2ea57"></a>

## origin_servers.private_ip.snat_pool — origin_servers.private_ip.snat_pool / d173b83c3687 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- origin_servers.private_ip.snat_pool

<a id="canonical-6c277abbd82cd7606f5f1221aea8416b0e276fa5d07b6c19a2f3ac777394d73a"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-08579706e3c8fda14b7fb208959560b01fdac99b04775c1a01ca71540f07d145"></a>

## Direct properties — origin_servers.private_ip.snat_pool / d173b83c3687 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2e724a4776d95ffb566d481b298b3f991986ef5a0764b74ae16aba83624d7b74): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-c9ed1627ba40f04c8ff8587ce9a846e49278e7752e2f109a896f5847ce521968): complete subsection reference.

<a id="canonical-e22103afe711f8e99ed35a332ee91dfbf2b75945c669350218c203637888d2bd"></a>

## Next pages — origin_servers.private_ip.snat_pool / d173b83c3687 / 4

- [origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2e724a4776d95ffb566d481b298b3f991986ef5a0764b74ae16aba83624d7b74)
- [origin_servers.private_ip.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-c9ed1627ba40f04c8ff8587ce9a846e49278e7752e2f109a896f5847ce521968)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-2e724a4776d95ffb566d481b298b3f991986ef5a0764b74ae16aba83624d7b74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c24e490d9589247faacdf3bd7ebae2b5e3528dd074f1827241209aff2b0d05a"></a>

## origin_servers.private_ip.snat_pool.no_snat_pool — origin_servers.private_ip.snat_pool.no_snat_pool / 240f1d03863a / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-03cb97e9dcba6f935a9010bd047b3d434831b28606f0299eae5dd8f38935391a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-1dbc1240899e976a0509c053bff5e9b267db55c13123af9843bceea6b8fdbfcd"></a>

## Direct properties — origin_servers.private_ip.snat_pool.no_snat_pool / 240f1d03863a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4606c3c7c369c2377c0dfdabe415cb0f3d510f38bd1ee2526376d76ad880b362"></a>

## Next pages — origin_servers.private_ip.snat_pool.no_snat_pool / 240f1d03863a / 4

- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-c9ed1627ba40f04c8ff8587ce9a846e49278e7752e2f109a896f5847ce521968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e24cd033f29ee011655a2e6bc492bf7447b7cc574cc0587b89e20ae8d8abc0b5"></a>

## origin_servers.private_ip.snat_pool.snat_pool — origin_servers.private_ip.snat_pool.snat_pool / 37721ec71eea / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-50cb4555d0f9bf0249b534868d02f0c9b006ba58db7e81b09ea3ac887a853b52)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1)
- origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-aeffe7333c62117e0a957a43be80b3041178db66de2179e59433bbcd1f3131ec"></a>

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

<a id="canonical-e6888720a9590c5a64d815044d68935e0e88bc4f6067212118ce8d15d3cae02f"></a>

## Direct properties — origin_servers.private_ip.snat_pool.snat_pool / 37721ec71eea / 3

<a id="canonical-ab71ad81e970f7eec2dfff64a2122b6730f280d8b47f828162b1b6bab13b2887"></a>

<a id="canonical-2f99695c8e731f6866e82edc62e4c52e76b28cbde1f0895ec77858008e6f2dd4"></a>

## prefixes property — origin_servers.private_ip.snat_pool.snat_pool / 37721ec71eea / 4

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

<a id="canonical-8d64ee58d7483d9dd2d77a9031bcce359a3e253252474fe856fb245d8187b01c"></a>

## Next pages — origin_servers.private_ip.snat_pool.snat_pool / 37721ec71eea / 5

- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3feb6ec9368c2b22eac6fbd43cb10981758d966c66923e03dfb21db3ed34bdc1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301d6de16566e0467f2241357add1fe646ef97a606410497964faca6c111184"></a>

## origin_servers.private_name — origin_servers.private_name / 0d0a3003ebb0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.private_name

<a id="canonical-667792fdfba67427d732773d4a07976b305fdba8299bb49a1eb73821116abd00"></a>

Type: `"single"`. Computed.

Specify origin server with private or public DNS name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

<a id="canonical-5d6dd24c4e5371eea8ae9ae72df1d52a36d888de178ae1cd669fbeef5b4aaacd"></a>

## Direct properties — origin_servers.private_name / 0d0a3003ebb0 / 3

<a id="canonical-0f529d7a2e0a68cc2a322a192f4f56a0b9214c9d189b9ef5c3dc795954baeee7"></a>

<a id="canonical-a8c2d0ff287cd7a4bd7250323e1f245999c28c51166a5372dd30660223ba0521"></a>

## dns_name property — origin_servers.private_name / 0d0a3003ebb0 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-a15cb5719c3db98cd1cd6f9218b055f4bfd5eb64d6af3e836f4b05c546144037): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-bf07ae4c6b72e966ad75a852bc99b57e2f148a0d487f08b36d82ab9659b136c0): complete subsection reference.

<a id="canonical-04c72a2b1a5aa9f776b5733a3f786acbc5ab531b218e45f54a8ad4cc928d1a39"></a>

<a id="canonical-81aa5dac05cd7e73726d8b5fa7ca6ab2abea9264c08fe9ac93f0fad4b2a6c4fb"></a>

## refresh_interval property — origin_servers.private_name / 0d0a3003ebb0 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](data-sources--origin_pool--reference--group-002.md#canonical-929a3c6b7ffbfdc5e3dc32b8e834245d8a35656f1b653c5ecb08d766d0f6e544): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385): complete subsection reference.

<a id="canonical-f84bcdd37fd1b3cfeb54b4e0a08ac9875c6e0253d206e85caed8fe6100c50bab"></a>

## Next pages — origin_servers.private_name / 0d0a3003ebb0 / 6

- [origin_servers.private_name.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-a15cb5719c3db98cd1cd6f9218b055f4bfd5eb64d6af3e836f4b05c546144037)
- [origin_servers.private_name.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-bf07ae4c6b72e966ad75a852bc99b57e2f148a0d487f08b36d82ab9659b136c0)
- [origin_servers.private_name.segment](data-sources--origin_pool--reference--group-002.md#canonical-929a3c6b7ffbfdc5e3dc32b8e834245d8a35656f1b653c5ecb08d766d0f6e544)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-a15cb5719c3db98cd1cd6f9218b055f4bfd5eb64d6af3e836f4b05c546144037"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af150db8d35811a70fc13575b1d6d1b5c64d77439fdd01766798c4194bcbf35c"></a>

## origin_servers.private_name.inside_network — origin_servers.private_name.inside_network / ea7013ec0708 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- origin_servers.private_name.inside_network

<a id="canonical-f61fc861ded55b3aaeaca9bf478fd806cfc1b8ca0038f20f8cff1b617cda6bf0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-3e45413b5653e7e388d1856b3fb5cbd7c646ac79dc1ded1ab2472c63d83fe76d"></a>

## Direct properties — origin_servers.private_name.inside_network / ea7013ec0708 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-426f8511969dcfd499660afa36e053345e808cd5520e462cf2f79e7b77c4d43a"></a>

## Next pages — origin_servers.private_name.inside_network / ea7013ec0708 / 4

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-bf07ae4c6b72e966ad75a852bc99b57e2f148a0d487f08b36d82ab9659b136c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a94d0650aaf00482aa9805cc422d8e0a31ce22d4a35231dffd8eb70d9ff28452"></a>

## origin_servers.private_name.outside_network — origin_servers.private_name.outside_network / 5b6672247a70 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- origin_servers.private_name.outside_network

<a id="canonical-17d12717d08771d18c828494df50dd981a88224e4f07d613c6f9f280b9d7eedb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-32a4819389da1eab583f493547a43b9e9cc7e6f982d9f493427f79ccbfdd0254"></a>

## Direct properties — origin_servers.private_name.outside_network / 5b6672247a70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-375c3da47e8addbbe182634be0d894129df9ed1d3e6291310ff13a537fef58ac"></a>

## Next pages — origin_servers.private_name.outside_network / 5b6672247a70 / 4

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-929a3c6b7ffbfdc5e3dc32b8e834245d8a35656f1b653c5ecb08d766d0f6e544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a24eac70f6fb2f0a3aae1b6a6171b0db0ae860c6c7ecf9674af7c41e7fc72c1a"></a>

## origin_servers.private_name.segment — origin_servers.private_name.segment / 2bee69a3ed8c / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- origin_servers.private_name.segment

<a id="canonical-6904524a0f14570289a82147d89bf02b434923439c9f94eab0c059d42a4d62ba"></a>

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

<a id="canonical-b707f43a8615b1244ba7e1f2535f5cae706ce26a07f9238f779630c5183fd3a6"></a>

## Direct properties — origin_servers.private_name.segment / 2bee69a3ed8c / 3

<a id="canonical-4f89f4147a17b07cd992f4dbc0aeac320a99ee65f10c23f33d71da6090dd7d84"></a>

<a id="canonical-2beb7f9e0bb37a8100afc980ae4a37b2d532686553d907c9c48d60e0a88fc0a5"></a>

## name property — origin_servers.private_name.segment / 2bee69a3ed8c / 4

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

<a id="canonical-b96fa63736b2d770e31bbff123dd65be7b5f843d8a113721a45fb7cf06a7d8c4"></a>

<a id="canonical-316e85a6ed6cdb9291b212334ecf5bc377ce2ec0cb6a0033b082441176740600"></a>

## namespace property — origin_servers.private_name.segment / 2bee69a3ed8c / 5

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

<a id="canonical-648bda1664e6fcc2c9debf64b530eb850e135fc20479cddf9acf764767b1e46b"></a>

<a id="canonical-96514a2a12af8727cfdbe50ded4223736349a6c2f3448c90f9b31fe85dc30e76"></a>

## tenant property — origin_servers.private_name.segment / 2bee69a3ed8c / 6

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

<a id="canonical-03ae8c98ee8f8fcd805b2665714c51f384c4f29ee91f5c78badec01c7295da9d"></a>

## Next pages — origin_servers.private_name.segment / 2bee69a3ed8c / 7

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12e8a9604f964fe671246f1196934515efa5507822cbd7083df978e10aaa06bf"></a>

## origin_servers.private_name.site_locator — origin_servers.private_name.site_locator / cc8cfd3abe97 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- origin_servers.private_name.site_locator

<a id="canonical-7b3ccef1409ee4a76a5e010daf8828653625a1a8c5ef9e936abc9607e6f7621b"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-f22b4fcf1527afe313bfdb23e7fb3f5319baea1f88f46c0916380f7652043ac8"></a>

## Direct properties — origin_servers.private_name.site_locator / cc8cfd3abe97 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-0ede6cd1127643cac7cb505c668736581e30d4197167bd130b4df17347a57ee0): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-1fcdb384b376a5e0ad4a0a67a5812691cbd655cd05d8c343727059c4ccf3283a): complete subsection reference.

<a id="canonical-fa3a5baff657b273a18bd83209afeb78591416dd121d14a32e277db0c9411edb"></a>

## Next pages — origin_servers.private_name.site_locator / cc8cfd3abe97 / 4

- [origin_servers.private_name.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-0ede6cd1127643cac7cb505c668736581e30d4197167bd130b4df17347a57ee0)
- [origin_servers.private_name.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-1fcdb384b376a5e0ad4a0a67a5812691cbd655cd05d8c343727059c4ccf3283a)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-0ede6cd1127643cac7cb505c668736581e30d4197167bd130b4df17347a57ee0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3377c99435bcf40ade08802e1a2e9683d627e039f1f88113234f07f1a21d6575"></a>

## origin_servers.private_name.site_locator.site — origin_servers.private_name.site_locator.site / fc82e2305907 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63)
- origin_servers.private_name.site_locator.site

<a id="canonical-9a4db3929dff81fa270c7131f4866042e502e3175158e4953ded11a8c3044661"></a>

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

<a id="canonical-a08edf648e0b8ae5a12cdddc5cba9f726383084bf0b280b6410e65bc11592e52"></a>

## Direct properties — origin_servers.private_name.site_locator.site / fc82e2305907 / 3

<a id="canonical-65bce5e841889dee3589cff5b5c32d52add7b53801918c3247446512dc530204"></a>

<a id="canonical-1d2208750cc09169336aee1211572ac89d64bbfd2ef29376b7179c1edf7316d5"></a>

## name property — origin_servers.private_name.site_locator.site / fc82e2305907 / 4

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

<a id="canonical-5839b765023d267a3dc1b58077bd510fd57971ecd204668d58b5df83d4db9b58"></a>

<a id="canonical-479a31bfb736aedbc3ceaf009310c85e450018b60875a09b912f791953487869"></a>

## namespace property — origin_servers.private_name.site_locator.site / fc82e2305907 / 5

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

<a id="canonical-6f8fc465f4a258f47d7039fb8427c3acc6b4b84715207bbbedb7850adde7454d"></a>

<a id="canonical-29add6de1c55c8f1c8140947c83a6123e645aee467059ec2b3e361da7b435ddf"></a>

## tenant property — origin_servers.private_name.site_locator.site / fc82e2305907 / 6

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

<a id="canonical-09e754b8a4a2b2de33093f64104d80fccee49b9fde1b9b13fb433e9f736c8fd6"></a>

## Next pages — origin_servers.private_name.site_locator.site / fc82e2305907 / 7

- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-1fcdb384b376a5e0ad4a0a67a5812691cbd655cd05d8c343727059c4ccf3283a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c7d4eec4bdb9b3b8927a14cc818169b915a8c7bf634b4cbfd60cfaa49bf35c2"></a>

## origin_servers.private_name.site_locator.virtual_site — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-4c79d144666db1e529185b266b67efba4eba2299c5f749d2133f917b06c47b1f"></a>

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

<a id="canonical-3ce7da942ddd850e4c94cba25ef71c3396f3ea3a22d58075d69ef3539d9fa113"></a>

## Direct properties — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 3

<a id="canonical-d43cc47befc9042c20c508601dbae18eac429611daef50a92ac133ba2a405cda"></a>

<a id="canonical-4920b6de58e00acf46c5233c891fd4a0e7bbb6a0392fcb36e4315922e775e1d3"></a>

## name property — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 4

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

<a id="canonical-3fa422982493f467854901c9eca3b2b3559ad7c92e4c9e232ddeda183fbd9b30"></a>

<a id="canonical-ea24d90ab54647eecf3045520ff0d16661b74bfa9cbf12c2c7cd1c0dc96c589b"></a>

## namespace property — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 5

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

<a id="canonical-6068d2cb7dd17488712770a55c9c07be7a657b14b849ebe8c7ba805f2eddd00b"></a>

<a id="canonical-e98cfb038607a66e48d7eea386f00b331c20b387896a52927de963dbe1914ce1"></a>

## tenant property — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 6

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

<a id="canonical-d78fc0d6548025dc6ee502f3bf6dfc170ab1df1130aafb135d8db178735ae040"></a>

## Next pages — origin_servers.private_name.site_locator.virtual_site / da2c76729bb3 / 7

- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-943ac7276f7484271337deafdec2a2f869714b36dabcb51a273ad657c0626c63)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41079c1206ed826b697d61dc350eed5635dd12f6349cac378dafe684e9c7ce44"></a>

## origin_servers.private_name.snat_pool — origin_servers.private_name.snat_pool / e3fc12a739a5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- origin_servers.private_name.snat_pool

<a id="canonical-7dbe6978bc6f62e71f4e56661845e2910cfbf09873e5066cda24b7ce71d99c27"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-f5e9fcfecd29f1e9706ec6cf7b98add51283445e304f9eb3671e6d4941767b2f"></a>

## Direct properties — origin_servers.private_name.snat_pool / e3fc12a739a5 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f6619a0bec7635cb6a1de7fd84dcfbecbefe4e959221147c95f4793f5f859eab): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-e6c986207b800c532bb849cfd3a63c1d1dbe56e610523d32493b715bd167f683): complete subsection reference.

<a id="canonical-8003da0fef369d60b1587dbd0072be111ade214c88627af37a1bb28800578758"></a>

## Next pages — origin_servers.private_name.snat_pool / e3fc12a739a5 / 4

- [origin_servers.private_name.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f6619a0bec7635cb6a1de7fd84dcfbecbefe4e959221147c95f4793f5f859eab)
- [origin_servers.private_name.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-e6c986207b800c532bb849cfd3a63c1d1dbe56e610523d32493b715bd167f683)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-f6619a0bec7635cb6a1de7fd84dcfbecbefe4e959221147c95f4793f5f859eab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9ba056930900f80dbae8d537e3cffee908faf27192a1fb58faecca3ff10001b"></a>

## origin_servers.private_name.snat_pool.no_snat_pool — origin_servers.private_name.snat_pool.no_snat_pool / bdf1363bd011 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-6d87f71e424905304d5cce2e2003cbce06a4d6edd847bde0da709e7ee3117f6b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-70a58deba9a2faf4912048797ff8a1fe7c95a78132c7e8d69899d679c23076a1"></a>

## Direct properties — origin_servers.private_name.snat_pool.no_snat_pool / bdf1363bd011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8121cce980d4dee0d37b4f3bdd7713b839fbaaf4e2f30077184f6c8672089f50"></a>

## Next pages — origin_servers.private_name.snat_pool.no_snat_pool / bdf1363bd011 / 4

- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-e6c986207b800c532bb849cfd3a63c1d1dbe56e610523d32493b715bd167f683"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9536b7747af0e239bb5708ba979e978476ba8f2bc15bad57f39dc93f1c59329"></a>

## origin_servers.private_name.snat_pool.snat_pool — origin_servers.private_name.snat_pool.snat_pool / 472566a62f49 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-61a15c78258e70ae043f1c9a7c848a7e88ecf44f8b68583c92181f31c20c63d4)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385)
- origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-8060bd568736191dc8507aba98f9bab9ae42285bf7c1d750613a493346491fb0"></a>

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

<a id="canonical-6a78fffcd044e402f789bf7b540e8ca63078be5d7faefa6106d3b6a37ad41f22"></a>

## Direct properties — origin_servers.private_name.snat_pool.snat_pool / 472566a62f49 / 3

<a id="canonical-2f4e5b5f94a83b91387f1802cc1a79784f1423ddf054beec65d0f9b6d8c32c76"></a>

<a id="canonical-2da342516b02e6afea2440bbb55121128f5df831739cfe0e37bf0a7cfce4656a"></a>

## prefixes property — origin_servers.private_name.snat_pool.snat_pool / 472566a62f49 / 4

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

<a id="canonical-4856116a5710c413cafd9080a8cf42c23db55068f24135dccd8b7ebd1aa959cd"></a>

## Next pages — origin_servers.private_name.snat_pool.snat_pool / 472566a62f49 / 5

- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-f617e76f6c7dddc4c70d9b7b2b9a7606319f346240db83bce468d9978032a385)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-efafe219d02588b88db03465cb41b4d228fdc93ba156e95a4f9774d7c13e45f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb52d1729f602017f4643512b12b7c3ab03d1d6cc9568793ca62b77f0d36c800"></a>

## origin_servers.public_ip — origin_servers.public_ip / cbc470db89ca / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.public_ip

<a id="canonical-5b50a97014f839c7e6097061273397bc043c16d29c282c9c4e363e0c398d2e8f"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-5d3c967c7f7fb989d38949e85cbd2f65f8690046a8144456b77c70850254777e"></a>

## Direct properties — origin_servers.public_ip / cbc470db89ca / 3

<a id="canonical-96df61bbb42bcbf4d6d570d273b1b1027d155ea13f888458e14b768c105a6b9a"></a>

<a id="canonical-f4d914fc6aac4a24dd687c995615faeea166cf6f9cb7568433fb446678504823"></a>

## ip property — origin_servers.public_ip / cbc470db89ca / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-fa3ef13b9ff8037b6e44751bda66dd92e0b414291798fce1102ea2030caaf9c5"></a>

## Next pages — origin_servers.public_ip / cbc470db89ca / 5

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-cb646c3857fcf36e43739596634a67807483b7ec8de38ca2d9dbe38dbb706755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b932c8fb292450e079f29059da1bd78a7b8d6d4817966ba20b64030ad53c6e8"></a>

## origin_servers.public_name — origin_servers.public_name / 932b5ba06723 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.public_name

<a id="canonical-6fb411cd8555e196f633d30e145238f4efd924bbdc6c77917cbb92e6ac41c14b"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-86345d86adc8acc86f26ae25fecd5aa0272d108cc0096cf17090fb2e9ff6530e"></a>

## Direct properties — origin_servers.public_name / 932b5ba06723 / 3

<a id="canonical-60990dda2ca9a53fce1e5e9fab6c1f1553d9cced4c81931f31eeb8faf25cbf3e"></a>

<a id="canonical-decd8c1588f78461901a492601ac52a4a6e1834118082c527bcbeb8e959b00c1"></a>

## dns_name property — origin_servers.public_name / 932b5ba06723 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-eaa734c3a1f00cd790dc29553de4f0e89cd444c1e063f41cfb643b0f0bbdccc9"></a>

<a id="canonical-02f9b42525e3b1cae94e6a2ad0d4fd79fba9280533e5742aaca7b26564f0ffe5"></a>

## refresh_interval property — origin_servers.public_name / 932b5ba06723 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-1ac496ef6fe2ec168d35aceba41218ece46f3c6f07a8e4c58d8fefd268df81e7"></a>

## Next pages — origin_servers.public_name / 932b5ba06723 / 6

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-5edbfdaafe6b2fa0b48a8773b6ef085a8ce8bee684f9d6c33db4a37c37bd7262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67d316de367ca971e725880081ffac105a2f991bbcd4f3741270a783c874fb36"></a>

## origin_servers.vn_private_ip — origin_servers.vn_private_ip / f958271726f8 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.vn_private_ip

<a id="canonical-615972742a9e0fc94f9b8822ece675ed1f58f6ec4a0c83b89a96ad8650f3d09b"></a>

Type: `"single"`. Computed.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-3ca5bd66ab2d5d6e5e35067aa69e9fac2f008ef2a14571b412b1788a4eb25d92"></a>

## Direct properties — origin_servers.vn_private_ip / f958271726f8 / 3

<a id="canonical-83afeac1eed5b3b4760440b08c09eac022157e45e8c376eeb6629c4789738a8c"></a>

<a id="canonical-4ad1e44c8fd3165b12314a7dea6fd1ff4045cd053e275a95870587169b1ef033"></a>

## ip property — origin_servers.vn_private_ip / f958271726f8 / 4

Type: `"string"`. Computed.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-8449650bca14a68fbb2e3547efc0cc1edaf6c6d86003c945292b1dc896b7940b): complete subsection reference.

<a id="canonical-f1539cea58e968532ec5097e87bf933fd255027ae2b83a0ec9d3e7a7f4b99a7e"></a>

## Next pages — origin_servers.vn_private_ip / f958271726f8 / 5

- [origin_servers.vn_private_ip.virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-8449650bca14a68fbb2e3547efc0cc1edaf6c6d86003c945292b1dc896b7940b)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-8449650bca14a68fbb2e3547efc0cc1edaf6c6d86003c945292b1dc896b7940b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84c75fe06343a27285c31f040fcda0e77f36a9041e36f9309044a6b48a06caaa"></a>

## origin_servers.vn_private_ip.virtual_network — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-5edbfdaafe6b2fa0b48a8773b6ef085a8ce8bee684f9d6c33db4a37c37bd7262)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0225eaf983978cc253679db5ab9305bd30e56815dfcda7bb3bf124c9cc946cfe"></a>

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

<a id="canonical-69921390c306b61337ba1d341a4534d12011cf8ac22919107419fd6268ca501e"></a>

## Direct properties — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 3

<a id="canonical-461fcca72d51e6fc9f9dd2f18bf68ec3eefbc90d9c29e2bab0d481c0e34ce521"></a>

<a id="canonical-5f37475680bd8a671f2a009dbb5b52ef61a9c966966746f1c583e518ec406139"></a>

## name property — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 4

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

<a id="canonical-ed4a74dbc0196ab42390a63dba4a755320d697b5f8106d4c5990833788c424db"></a>

<a id="canonical-d3cce6bf14b020d42af182d295d7536e858a98355ba94d91f6638692f8b4786c"></a>

## namespace property — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 5

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

<a id="canonical-f6abe4c886be357bccb89294b6f194ad77afe6036749473e4a8655a3eb852cf0"></a>

<a id="canonical-4957f3d9502c50aa248b8933ea7268c5f78bec45cbf1fbab8a5dce5d2428a3c0"></a>

## tenant property — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 6

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

<a id="canonical-f8209d6f1eb91639ed4a42e709d2c3e9cc175af9ea1c92559e83f80fd39a5d07"></a>

## Next pages — origin_servers.vn_private_ip.virtual_network / ee24ee60e837 / 7

- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-5edbfdaafe6b2fa0b48a8773b6ef085a8ce8bee684f9d6c33db4a37c37bd7262)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-fa5838e570f4c510e55c3914c3f94319a049dfc3fb3ca9a70f4f2a3a098cae3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-015614107be56bbe26992fe0f90d91e8235fd05e08b7184bdf2c7da10524ecde"></a>

## origin_servers.vn_private_name — origin_servers.vn_private_name / d806be358236 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- origin_servers.vn_private_name

<a id="canonical-f6c0b230a9fe56ffbcfe1e0af871b97796562a72ab579c4bd86ec2ca0d1c55ed"></a>

Type: `"single"`. Computed.

Specify origin server with DNS name on Virtual Network.

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

<a id="canonical-4b01e2fe7c6f574f882fe6bdb5cff8dd2b321a9352d11210051fe47959490b48"></a>

## Direct properties — origin_servers.vn_private_name / d806be358236 / 3

<a id="canonical-51d7dccc0e1b4ae14c5c1250eb838b76b0a6de54e530a0863bfe8c3058d929f7"></a>

<a id="canonical-f0d46f8a53940e7986b27e0a84f308517a7ed53b84bdc414dd9861d7bf2442f7"></a>

## dns_name property — origin_servers.vn_private_name / d806be358236 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [private_network](data-sources--origin_pool--reference--group-002.md#canonical-8aa889e2d413bdada0bce02ba76fa298445a234f77994f12c9d19b21534a4bb7): complete subsection reference.

<a id="canonical-18a3f29b051ba4bf6304c4cffccadc094e5defd0af9536d4cdd7b2e13df37597"></a>

## Next pages — origin_servers.vn_private_name / d806be358236 / 5

- [origin_servers.vn_private_name.private_network](data-sources--origin_pool--reference--group-002.md#canonical-8aa889e2d413bdada0bce02ba76fa298445a234f77994f12c9d19b21534a4bb7)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-8aa889e2d413bdada0bce02ba76fa298445a234f77994f12c9d19b21534a4bb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af1bfebfb4c725de73f460669e67bbef29b38689119983adcd5e0d7080dc3bcb"></a>

## origin_servers.vn_private_name.private_network — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-6fc05a53b4cda5572c9fcfe43f673aa0d2695560272c32eb09b75bad5b3047a5)
- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-fa5838e570f4c510e55c3914c3f94319a049dfc3fb3ca9a70f4f2a3a098cae3c)
- origin_servers.vn_private_name.private_network

<a id="canonical-8a04b86b67d07902bfdaf1485f1285772e6062aeba087bcbf9bcbabecc47f304"></a>

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

<a id="canonical-b49a8b4345ee953c4e9b83fa01eb1e86c868b0e058fdab68142bcb8f2a7e8064"></a>

## Direct properties — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 3

<a id="canonical-dd31a30eab24cd74c8153d2c1011ebf0fe38acfa37a38368f548e5c8bfb9d805"></a>

<a id="canonical-060e75a1d204c55e0025d8fda9ed7c56d6034222fdd6694e25a7d1d927b8db07"></a>

## name property — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 4

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

<a id="canonical-50698aa361d95dd078fa56fc4bb4953fba7cba0c88883d1a7794e6681f23c094"></a>

<a id="canonical-7fa7efd4817dcf0453d8c08d151372c20ab1212a1dd9f463e32cf000d71d1558"></a>

## namespace property — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 5

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

<a id="canonical-9a7dfabf7da9cb20c871ef085905da302cf463f5fd7ff34646cae2dd996a8862"></a>

<a id="canonical-b2af6fa94057d88ddbc2512618a6bdf47c86c9fc91ef20eb532848a2ef16e82a"></a>

## tenant property — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 6

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

<a id="canonical-7cfa8fe5b776fb682bd42eab35cd608d3340c8ff6ae7d1beeb07410ff61dea8b"></a>

## Next pages — origin_servers.vn_private_name.private_network / abbcd2ed25fa / 7

- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-fa5838e570f4c510e55c3914c3f94319a049dfc3fb3ca9a70f4f2a3a098cae3c)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-dc185cb410008c05bf9d16e9bbc2d76dec9df69be9a6270d1af105dc52e365cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-845603d3787ac5d23b3914cd022ab85341601ee5eaa5203e44359cee863b1888"></a>

## same_as_endpoint_port — same_as_endpoint_port / 0e7dc84ca843 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- same_as_endpoint_port

<a id="canonical-ec881b452d0dd978f0b3bf07f4f243573b579dc9584c3238af9df0f90788304c"></a>

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

<a id="canonical-01b2b26031972b0471ecd231bb73852142d3795c95b056d3083419ff9323b496"></a>

## Direct properties — same_as_endpoint_port / 0e7dc84ca843 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82ce2b345ae7b029dfc1717a1bad3bf9fc89d3d5894a340cc6ade011c51bc964"></a>

## Next pages — same_as_endpoint_port / 0e7dc84ca843 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a52ef207651b0dc31fd71749e68bece25ebb2f684ba38ed0a2f5c8aace8a45f"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 9d300a11e472 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- upstream_conn_pool_reuse_type

<a id="canonical-6bd239adb6868a830e65cede41e9b2d8cc481cdf1d8d2d182f468c43e4a30c7e"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

<a id="canonical-6138d16c196c06e21fccda561d532805c7baf548f138a0fe11e238a4c5f147bb"></a>

## Direct properties — upstream_conn_pool_reuse_type / 9d300a11e472 / 3

- [disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-4a839e1e5b2fe15fa8846883a07e2e33245345fec20f71a31d0a8954a98c5558): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-7844a385e385cd5c2e66528fe3c1ce43fb61919916fb673f007cc3308605b71b): complete subsection reference.

<a id="canonical-91a58bbb9a709345a7a67a521a3eb35c86980e8ffb17a0af0c58bb8ad3463430"></a>

## Next pages — upstream_conn_pool_reuse_type / 9d300a11e472 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-4a839e1e5b2fe15fa8846883a07e2e33245345fec20f71a31d0a8954a98c5558)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-7844a385e385cd5c2e66528fe3c1ce43fb61919916fb673f007cc3308605b71b)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-4a839e1e5b2fe15fa8846883a07e2e33245345fec20f71a31d0a8954a98c5558"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-032a300c7a51c1c7ce67d8dd484a7447ac3633b8925801a300480eabe0ea9889"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / b5a2040e9c3e / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-695f2671b7219f78298ff6fb9b82bc027c354e9ff781f4367457ae429ba7e127"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable conn pool reuse.

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

<a id="canonical-ea3a90a2512cce61d8a57f54be7ef90a7fd9ad586e499ff9466b6be0668a0e16"></a>

## Direct properties — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / b5a2040e9c3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47d891cf05518a7aac364af0e8a030b7699600f1886bb8a2d31716e65b03c6a4"></a>

## Next pages — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / b5a2040e9c3e / 4

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7844a385e385cd5c2e66528fe3c1ce43fb61919916fb673f007cc3308605b71b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0e6ed61508597869f2a28078d546fe9fd6d3c47bd8aa7bb2d46acdb08bff4bd"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / cd2d367164d1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-73d8e778899a16458998e6722757d7982183e42821eabcd79c62776ff361305e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable conn pool reuse.

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

<a id="canonical-f3a8c40b57ee2ed5c8e17f68c37ece4392d0dddba56c331165c44a425a3df5f9"></a>

## Direct properties — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / cd2d367164d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab44cd0235d1eb5c6d792e2153af3d6073e8c1cd810eba002d8aca9cc5ca94dc"></a>

## Next pages — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / cd2d367164d1 / 4

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-adb23bd86b48fa66a508c03f7e9aaada0d277f11227212aadd78b5f03c16e48d)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68a7e542e9e8b42ed634101f2d66dfb962198f6a077fb0a59d8ea512adb400ea"></a>

## use_tls — use_tls / b77a171671f6 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- use_tls

<a id="canonical-02a9345e3df45300299a1b27753191de2b20bdfeb558d4a0de38bf539a781e09"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-07cb440de564944a107524f5b977708152ba01992fc51bdcfb93672ea7cc5334"></a>

## Direct properties — use_tls / b77a171671f6 / 3

- [default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-d83ebc1a02189e84df943670a1e205874a67869579b7014a6232b11d9c0813c6): complete subsection reference.

- [disable_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-d62ce7b65e102d42c3e2c51e5bce2c7e366d2abe3373407e5b7f787c81308e25): complete subsection reference.

- [disable_sni](data-sources--origin_pool--reference--group-002.md#canonical-f2f411530e937f9cbc450ecbef35efb0efa8ca892daf30c9687f237665dc399e): complete subsection reference.

<a id="canonical-2f4d1a3848d7351750d559796df92c62b839caae0f5575a890646f791c5ac7c4"></a>

<a id="canonical-1b05cad139c1acbba01196675d26dff20877003afd5b66bd90af1cd7f6d1dda0"></a>

## max_session_keys property — use_tls / b77a171671f6 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--origin_pool--reference--group-002.md#canonical-8f31c62c821c62d0b044511b0b519ab78229141e6099214a2e5220c44927b3de): complete subsection reference.

- [skip_server_verification](data-sources--origin_pool--reference--group-002.md#canonical-df4a50b3ce81ae582422e4b8895411c5679d0b509a26f33af955dbf437133e49): complete subsection reference.

<a id="canonical-f718cb2b088a501a8843aae5b022fceee2aac73a6f9d81ef0218a964397a929b"></a>

<a id="canonical-dc9411abc7b91dc731205b35007cd2ef4558cbc728b4873802ceccafb0e7bd12"></a>

## sni property — use_tls / b77a171671f6 / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--origin_pool--reference--group-002.md#canonical-c2a519ac8e9a21aaa54b2ffbcd56547e228bd2cffb1c99145c7a0f215ab109e7): complete subsection reference.

- [use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-ee6ab69ff991ac412a360f9da33ef7822a4a9d6c9331f477e8de86d579bf9d49): complete subsection reference.

- [use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-9bb1add53564965f7909c14ef6c20942943fe963998a8fae10af5890794500f4): complete subsection reference.

- [use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-f472d6715b700c85a732b254a994a203cf5881758a17fbe358df6ac4c344cbb8): complete subsection reference.

- [use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-c7b40a7716cf88d82d712cf8500afb716a8e232257eea09c3dd9275012f1da94): complete subsection reference.

- [volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-7ec79861964ba697e5bc1eb86c5d87001c9a94a4fbb1bbfe59cf314c1d5158bf): complete subsection reference.

<a id="canonical-e62432f8dd1063ad7b7185e335b787d86cd379056f941bbb05064cac91f22275"></a>

## Next pages — use_tls / b77a171671f6 / 6

- [use_tls.default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-d83ebc1a02189e84df943670a1e205874a67869579b7014a6232b11d9c0813c6)
- [use_tls.disable_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-d62ce7b65e102d42c3e2c51e5bce2c7e366d2abe3373407e5b7f787c81308e25)
- [use_tls.disable_sni](data-sources--origin_pool--reference--group-002.md#canonical-f2f411530e937f9cbc450ecbef35efb0efa8ca892daf30c9687f237665dc399e)
- [use_tls.no_mtls](data-sources--origin_pool--reference--group-002.md#canonical-8f31c62c821c62d0b044511b0b519ab78229141e6099214a2e5220c44927b3de)
- [use_tls.skip_server_verification](data-sources--origin_pool--reference--group-002.md#canonical-df4a50b3ce81ae582422e4b8895411c5679d0b509a26f33af955dbf437133e49)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-002.md#canonical-c2a519ac8e9a21aaa54b2ffbcd56547e228bd2cffb1c99145c7a0f215ab109e7)
- [use_tls.use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-ee6ab69ff991ac412a360f9da33ef7822a4a9d6c9331f477e8de86d579bf9d49)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-9bb1add53564965f7909c14ef6c20942943fe963998a8fae10af5890794500f4)
- [use_tls.use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-f472d6715b700c85a732b254a994a203cf5881758a17fbe358df6ac4c344cbb8)
- [use_tls.use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-c7b40a7716cf88d82d712cf8500afb716a8e232257eea09c3dd9275012f1da94)
- [use_tls.volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-7ec79861964ba697e5bc1eb86c5d87001c9a94a4fbb1bbfe59cf314c1d5158bf)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-d83ebc1a02189e84df943670a1e205874a67869579b7014a6232b11d9c0813c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f391f64ff78029cff3cca8eaf0390651199ad4484aea4e92d0acef0f9260992"></a>

## use_tls.default_session_key_caching — use_tls.default_session_key_caching / 0c17a23bb3f3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.default_session_key_caching

<a id="canonical-ba888a9c6e50be9fff457057e4ba0b7d6d61bfb8c4c4f9b6d31a5d2f359e437b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-c83828bc3d6c44e92321ba15fa5d84476dcb2a279c65e494c27a0cc0624df3e0"></a>

## Direct properties — use_tls.default_session_key_caching / 0c17a23bb3f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5fa970c5889cfcc3a4ff4304f409046741213a7188bf09301bae86edc1f1a1c3"></a>

## Next pages — use_tls.default_session_key_caching / 0c17a23bb3f3 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-d62ce7b65e102d42c3e2c51e5bce2c7e366d2abe3373407e5b7f787c81308e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-553b962d2915a2a4b279c4163f9790083444578af17adacd6d1df2d071653c2a"></a>

## use_tls.disable_session_key_caching — use_tls.disable_session_key_caching / 68da5d100845 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.disable_session_key_caching

<a id="canonical-831352753e23156bdbb91496d35c11133f82f0f292149cf556d8aeb77215cb30"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-e9e046525ef91680dbba52703b64dc23c2eaec8ca490ed7537bb7a64cc92f55d"></a>

## Direct properties — use_tls.disable_session_key_caching / 68da5d100845 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65eb800c8424e7d21e98782456de217d9d42a7b3e6e36f1fd8b78804add01343"></a>

## Next pages — use_tls.disable_session_key_caching / 68da5d100845 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-f2f411530e937f9cbc450ecbef35efb0efa8ca892daf30c9687f237665dc399e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ab525bcd26933f8c684c8c7dad1530a58ea24269007930f300d89924792829e"></a>

## use_tls.disable_sni — use_tls.disable_sni / 816d6b3c5a73 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.disable_sni

<a id="canonical-666a0984c79e72a9c9d9e7ed49dc7115cdb6a54f25c3a980bf9ea6cfa8183720"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-830e543d1959ae93740e8d1f7d2ef8d940cf8b67f9f35122df21359b93e2318d"></a>

## Direct properties — use_tls.disable_sni / 816d6b3c5a73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94dc35ad56a5d0391040c0552cc13bbac8c128cb4a03a1bbc851b02a820f86e3"></a>

## Next pages — use_tls.disable_sni / 816d6b3c5a73 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-8f31c62c821c62d0b044511b0b519ab78229141e6099214a2e5220c44927b3de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b10ddb1846941138bce602e1ff6c29536e6cd37c8ebc93ab20295f7473abafa"></a>

## use_tls.no_mtls — use_tls.no_mtls / 358b1e0a6758 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.no_mtls

<a id="canonical-99cbd2d8d8b63b3855a361be4c6a5827bc479ddd4b76aa5230de7d4d154d83db"></a>

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

<a id="canonical-c6103bbc2a0b18954067f81bc439fbf0dd1b029b7b873a7020dac570a05fe624"></a>

## Direct properties — use_tls.no_mtls / 358b1e0a6758 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-586b945be197c0db10dc25701ae3cbf9c371897698d89c855848d71168abf67b"></a>

## Next pages — use_tls.no_mtls / 358b1e0a6758 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-df4a50b3ce81ae582422e4b8895411c5679d0b509a26f33af955dbf437133e49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-153dad675273de195582b7127495494b2fb364df813f972d3a328e9e78933ac2"></a>

## use_tls.skip_server_verification — use_tls.skip_server_verification / 599197ccf3b3 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.skip_server_verification

<a id="canonical-fd23590b59a2ce32728a54efb4eb626edea2dfb7f9f5bd53e0cf6d9781716379"></a>

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

<a id="canonical-cfaf159527d9f820c1b0fbf07f32c2d014681b7972810a97f98e0dba36426701"></a>

## Direct properties — use_tls.skip_server_verification / 599197ccf3b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c0c046f97308dc84d1f223d6a989787103c112e6c02e7d4d9c0a4f1b67818f1"></a>

## Next pages — use_tls.skip_server_verification / 599197ccf3b3 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-c2a519ac8e9a21aaa54b2ffbcd56547e228bd2cffb1c99145c7a0f215ab109e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62fb325a0596e0021a874160c0167d0d3eb390b14792a110baf964cdcaf3df7e"></a>

## use_tls.tls_config — use_tls.tls_config / 56696615c3f0 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- use_tls.tls_config

<a id="canonical-a9649983f595f34dfab45b4b894164595c3bbdf94ade57504f20db1abee483e3"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-1e3747a2410381a2b2b2b6b6720bb8bef5856ebdc4724f86e417cd8ab6936a77"></a>

## Direct properties — use_tls.tls_config / 56696615c3f0 / 3

- [custom_security](data-sources--origin_pool--reference--group-002.md#canonical-7cbd072eee4ad1f00ed521f66a0142e720f190bef250922a0f1baa12276db721): complete subsection reference.

- [default_security](data-sources--origin_pool--reference--group-003.md#canonical-e10ddec2c169d411673281ecd0e3768de11e78dc82c3cbe5df78ae3e01b868f4): complete subsection reference.

- [low_security](data-sources--origin_pool--reference--group-003.md#canonical-9ac2ef49f7189dd822e97a3500009ae9c2fd82baf9309688b2687bd65c942604): complete subsection reference.

- [medium_security](data-sources--origin_pool--reference--group-003.md#canonical-e2d24761b77ce5bf0879ae9b51565441ad80979c03812f8afefea3d1da6609e9): complete subsection reference.

<a id="canonical-585804a8aa068927703261ed09f48c5bd9e41031ffe8a3fcd0d6ae5d17de22fd"></a>

## Next pages — use_tls.tls_config / 56696615c3f0 / 4

- [use_tls.tls_config.custom_security](data-sources--origin_pool--reference--group-002.md#canonical-7cbd072eee4ad1f00ed521f66a0142e720f190bef250922a0f1baa12276db721)
- [use_tls.tls_config.default_security](data-sources--origin_pool--reference--group-003.md#canonical-e10ddec2c169d411673281ecd0e3768de11e78dc82c3cbe5df78ae3e01b868f4)
- [use_tls.tls_config.low_security](data-sources--origin_pool--reference--group-003.md#canonical-9ac2ef49f7189dd822e97a3500009ae9c2fd82baf9309688b2687bd65c942604)
- [use_tls.tls_config.medium_security](data-sources--origin_pool--reference--group-003.md#canonical-e2d24761b77ce5bf0879ae9b51565441ad80979c03812f8afefea3d1da6609e9)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)

<a id="canonical-7cbd072eee4ad1f00ed521f66a0142e720f190bef250922a0f1baa12276db721"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c13fbfded83475f45d403e25af6ca8fd198193eaebde5ad043c155cfb426c3f8"></a>

## use_tls.tls_config.custom_security — use_tls.tls_config.custom_security / 152a8c3d3608 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-d314ac39e03cc8cbfa1def4c1883fc9446d180ce6bff7b20bfe723be7524cfa3)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-13a175ac8642ead74ba9d47f3c000069c0c99a69541facda2dd9ba38ca310af3)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-e8bc8540eeeed3deca5240a7eb530532ba669c77c29f5c9d57749d49d8dc57f1)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-002.md#canonical-c2a519ac8e9a21aaa54b2ffbcd56547e228bd2cffb1c99145c7a0f215ab109e7)
- use_tls.tls_config.custom_security

<a id="canonical-d39a13a0de6ff191f3cc7c0c97a4c2b6d556d3555c6174fc91555c97c3e97753"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-57978e7a39abc7cd485d04fa8009bba4b354d435e1d17051a129c127dbfbc55b"></a>

## Direct properties — use_tls.tls_config.custom_security / 152a8c3d3608 / 3

<a id="canonical-5fd087290530d3962cc09485ff02b6bf92943da2cf7ad2a9c74bc0b0d72c0aec"></a>

<a id="canonical-8582e166d4b4f2dc21bae84604c39faf7a624442cbb70f6eed794006078de97f"></a>

## cipher_suites property — use_tls.tls_config.custom_security / 152a8c3d3608 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fbdd660f964ef55eded06a691021838f69b5e0c8c63a4f0c0b610634eba165e0"></a>
