---
page_title: "http_access"
subcategory: ""
description: "Configuration parameter for http access."
xcsh_docs: {"aliases": ["http access"], "body_bytes": 2044, "body_sha256": "sha256:a20da35fb4817955cc7aba2e7663341506db6879fe565c1c0ed3d425f171ad3a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "id": "xcsh-docs:resources:crl:properties:http_access", "parent_id": "xcsh-docs:resources:crl:reference", "path": "documentation/resources/crl/properties/http_access/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3321030000232202-1110003002322111-3110232113133210-0120020233031003-2100102002203222-0101012223111031-0331330123323332-3220011030031300", "registry_path": "docs/guides/resources--crl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_access"], "schema_version": 1, "sections": [{"aliases": ["http access path"], "anchor": "schema-http_access--path", "description": "CRL file location.", "document_id": "xcsh-docs:resources:crl:properties:http_access", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_access", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/properties/http_access/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for http access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_access

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/properties/)
- http_access

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http access.

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

Terraform syntax:

```terraform
http_access {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-http_access--path"></a>

### path property

Type: `"string"`. Optional.

CRL File path. CRL file location.

Upstream description:

CRL file location.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/properties/)
- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
