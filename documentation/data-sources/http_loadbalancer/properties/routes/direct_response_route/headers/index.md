---
page_title: "routes.direct_response_route.headers"
subcategory: "Load Balancing"
description: "List of (key, value) headers."
xcsh_docs: {"aliases": ["routes direct response route headers"], "body_bytes": 6704, "body_sha256": "sha256:05ae86c6575fa38489e2f122b1a0aaa93f3c143c5ba9d001a8c2fb243f258ee4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "path": "documentation/data-sources/http_loadbalancer/properties/routes/direct_response_route/headers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1312003333310320-0131033011312212-3332203311320023-1310003020010333-3031223213011201-3123033231122132-3222103001020231-0130211211103030", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "direct_response_route", "headers"], "schema_version": 1, "sections": [{"aliases": ["exact"], "anchor": "schema-routes--direct_response_route--headers--exact", "description": "Exclusive with Header value to match exactly.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "headers", "exact"], "syntax": "attribute", "type": "string"}, {"aliases": ["invert match"], "anchor": "schema-routes--direct_response_route--headers--invert_match", "description": "Invert the result of the match to detect missing header or non-matching value.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "headers", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-routes--direct_response_route--headers--name", "description": "Name of the header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "headers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["presence"], "anchor": "schema-routes--direct_response_route--headers--presence", "description": "Exclusive with If true, check for presence of header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "headers", "presence"], "syntax": "attribute", "type": "bool"}, {"aliases": ["regex"], "anchor": "schema-routes--direct_response_route--headers--regex", "description": "Exclusive with Regex match of the header value in re2 format.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "headers", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/direct_response_route/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of (key, value) headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.direct_response_route.headers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/direct_response_route/)
- routes.direct_response_route.headers

<a id="section"></a>

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

## Direct properties

<a id="schema-routes--direct_response_route--headers--exact"></a>

### exact property

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

<a id="schema-routes--direct_response_route--headers--invert_match"></a>

### invert_match property

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

<a id="schema-routes--direct_response_route--headers--name"></a>

### name property

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

<a id="schema-routes--direct_response_route--headers--presence"></a>

### presence property

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

<a id="schema-routes--direct_response_route--headers--regex"></a>

### regex property

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

## Next pages

- [routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/direct_response_route/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
