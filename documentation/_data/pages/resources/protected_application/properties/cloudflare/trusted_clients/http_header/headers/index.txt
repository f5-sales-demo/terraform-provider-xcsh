---
page_title: "cloudflare.trusted_clients.http_header.headers"
subcategory: ""
description: "List of HTTP header name and value pairs."
xcsh_docs: {"aliases": ["cloudflare trusted clients http header headers"], "body_bytes": 6903, "body_sha256": "sha256:98e6b73849b2d916b43b829649443d2debb6fdd07e2e552bf9770b979e371f36", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "path": "documentation/resources/protected_application/properties/cloudflare/trusted_clients/http_header/headers/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "schema-cloudflare--trusted_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--trusted_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--trusted_clients--http_header--headers--name", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers"], "schema_version": 1, "sections": [{"aliases": ["cloudflare trusted clients http header headers exact"], "anchor": "schema-cloudflare--trusted_clients--http_header--headers--exact", "description": "Exclusive with Header value to match exactly.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers", "exact"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare trusted clients http header headers name"], "anchor": "schema-cloudflare--trusted_clients--http_header--headers--name", "description": "Name of the header.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare trusted clients http header headers regex"], "anchor": "schema-cloudflare--trusted_clients--http_header--headers--regex", "description": "Exclusive with Regex match of the header value in re2 format.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/trusted_clients/http_header/headers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of HTTP header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients.http_header.headers

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/)
- [cloudflare.trusted_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/http_header/)
- cloudflare.trusted_clients.http_header.headers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
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

<a id="schema-cloudflare--trusted_clients--http_header--headers--exact"></a>

### exact property

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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

<a id="schema-cloudflare--trusted_clients--http_header--headers--name"></a>

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

<a id="schema-cloudflare--trusted_clients--http_header--headers--regex"></a>

### regex property

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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

- [cloudflare.trusted_clients.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/http_header/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
