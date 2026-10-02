---
page_title: "blocked_clients.http_header.headers"
subcategory: "Load Balancing"
description: "List of HTTP header name and value pairs."
xcsh_docs: {"aliases": ["blocked clients http header headers"], "body_bytes": 7535, "body_sha256": "sha256:4809df65492e3af328b253455136176dbf565fe0942af7088b87e09ec7649b3d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header", "path": "documentation/resources/http_loadbalancer/properties/blocked_clients/http_header/headers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "schema-blocked_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--presence", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-blocked_clients--http_header--headers--name", "enforcement": "provider-schema", "group": "blocked_clients.http_header.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_clients", "http_header", "headers"], "schema_version": 1, "sections": [{"aliases": ["exact"], "anchor": "schema-blocked_clients--http_header--headers--exact", "description": "Exclusive with Header value to match exactly.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_clients", "http_header", "headers", "exact"], "syntax": "attribute", "type": "string"}, {"aliases": ["invert match"], "anchor": "schema-blocked_clients--http_header--headers--invert_match", "description": "Invert the result of the match to detect missing header or non-matching value.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_clients", "http_header", "headers", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-blocked_clients--http_header--headers--name", "description": "Name of the header.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_clients", "http_header", "headers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["presence"], "anchor": "schema-blocked_clients--http_header--headers--presence", "description": "Exclusive with If true, check for presence of header.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_clients", "http_header", "headers", "presence"], "syntax": "attribute", "type": "bool"}, {"aliases": ["regex"], "anchor": "schema-blocked_clients--http_header--headers--regex", "description": "Exclusive with Regex match of the header value in re2 format.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:blocked_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_clients", "http_header", "headers", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/blocked_clients/http_header/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of HTTP header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_clients.http_header.headers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [blocked_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/blocked_clients/)
- [blocked_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/blocked_clients/http_header/)
- blocked_clients.http_header.headers

<a id="section"></a>

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

## Direct properties

<a id="schema-blocked_clients--http_header--headers--exact"></a>

### exact property

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

<a id="schema-blocked_clients--http_header--headers--invert_match"></a>

### invert_match property

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

<a id="schema-blocked_clients--http_header--headers--name"></a>

### name property

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

<a id="schema-blocked_clients--http_header--headers--presence"></a>

### presence property

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

<a id="schema-blocked_clients--http_header--headers--regex"></a>

### regex property

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

## Next pages

- [blocked_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/blocked_clients/http_header/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
