---
page_title: "cloudfront.trusted_clients.http_header.headers"
subcategory: ""
description: "List of HTTP header name and value pairs."
xcsh_docs: {"aliases": ["cloudfront trusted clients http header headers"], "body_bytes": 6174, "body_sha256": "sha256:3ed8fc072499df168336336dbd93da620c450344ef372690062f256ee41a8360", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header", "path": "documentation/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/headers/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1320020310120010-1103101210121231-1330321212313201-1013021310221013-0313000020322222-1331002033231033-1320023231221220-3130023230201223", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "trusted_clients", "http_header", "headers"], "schema_version": 1, "sections": [{"aliases": ["exact"], "anchor": "schema-cloudfront--trusted_clients--http_header--headers--exact", "description": "Exclusive with Header value to match exactly.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "trusted_clients", "http_header", "headers", "exact"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-cloudfront--trusted_clients--http_header--headers--name", "description": "Name of the header.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "trusted_clients", "http_header", "headers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex"], "anchor": "schema-cloudfront--trusted_clients--http_header--headers--regex", "description": "Exclusive with Regex match of the header value in re2 format.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "trusted_clients", "http_header", "headers", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of HTTP header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.trusted_clients.http_header.headers

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/)
- [cloudfront.trusted_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/)
- cloudfront.trusted_clients.http_header.headers

<a id="section"></a>

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

## Direct properties

<a id="schema-cloudfront--trusted_clients--http_header--headers--exact"></a>

### exact property

Type: `"string"`. Computed.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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

<a id="schema-cloudfront--trusted_clients--http_header--headers--name"></a>

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

<a id="schema-cloudfront--trusted_clients--http_header--headers--regex"></a>

### regex property

Type: `"string"`. Computed.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

## Next pages

- [cloudfront.trusted_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
