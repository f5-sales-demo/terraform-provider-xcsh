---
page_title: "waf_exclusion.waf_exclusion_inline_rules.rules"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules.rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10674, "body_sha256": "sha256:11d8ff5007ee495bbe6ff43d55b3fa70d4662fa6f5708351cc99b29e8b420818", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:metadata", "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:waf_skip_processing"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules.rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules.rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [waf_exclusion](data-sources--cdn_loadbalancer--properties--waf_exclusion.md)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [any_domain](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--any_domain.md): complete subsection reference.

- [any_path](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--any_path.md): complete subsection reference.

- [app_firewall_detection_control](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control.md): complete subsection reference.

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--expiration_timestamp"></a>

### expiration_timestamp property

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
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--metadata.md): complete subsection reference.

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--methods"></a>

### methods property

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

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix"></a>

### path_prefix property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex"></a>

### path_regex property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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

- [waf_skip_processing](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--waf_skip_processing.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--any_domain.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--any_path.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--app_firewall_detection_control.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--metadata.md)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules--waf_skip_processing.md)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
