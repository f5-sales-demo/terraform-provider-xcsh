---
page_title: "routes.redirect_route.headers"
subcategory: "Load Balancing"
description: "List of (key, value) headers."
xcsh_docs: {"aliases": ["routes redirect route headers"], "body_bytes": 6627, "body_sha256": "sha256:8cb9efdf6af5fbda80a5c32d31435e19ff09435a90d7a4c15194de2fba3d95cc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "path": "documentation/data-sources/http_loadbalancer/properties/routes/redirect_route/headers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1223101330120300-0132100302130130-2210221302103231-2102211101322213-2323033030232003-3012132221012210-0122030120100112-0033002130131022", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "redirect_route", "headers"], "schema_version": 1, "sections": [{"aliases": ["routes redirect route headers exact"], "anchor": "schema-routes--redirect_route--headers--exact", "description": "Exclusive with Header value to match exactly.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "headers", "exact"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes redirect route headers invert match"], "anchor": "schema-routes--redirect_route--headers--invert_match", "description": "Invert the result of the match to detect missing header or non-matching value.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "headers", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["routes redirect route headers name"], "anchor": "schema-routes--redirect_route--headers--name", "description": "Name of the header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "headers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes redirect route headers presence"], "anchor": "schema-routes--redirect_route--headers--presence", "description": "Exclusive with If true, check for presence of header.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "headers", "presence"], "syntax": "attribute", "type": "bool"}, {"aliases": ["routes redirect route headers regex"], "anchor": "schema-routes--redirect_route--headers--regex", "description": "Exclusive with Regex match of the header value in re2 format.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "headers", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/redirect_route/headers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of (key, value) headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route.headers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/)
- routes.redirect_route.headers

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

<a id="schema-routes--redirect_route--headers--exact"></a>

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

<a id="schema-routes--redirect_route--headers--invert_match"></a>

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

<a id="schema-routes--redirect_route--headers--name"></a>

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

<a id="schema-routes--redirect_route--headers--presence"></a>

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

<a id="schema-routes--redirect_route--headers--regex"></a>

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

- [routes.redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
