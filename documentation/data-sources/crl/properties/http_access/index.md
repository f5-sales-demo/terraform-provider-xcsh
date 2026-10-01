---
page_title: "http_access"
subcategory: ""
description: "http_access for xcsh_crl."
xcsh_docs: {"aliases": [], "body_bytes": 1794, "body_sha256": "sha256:1eb9584a0019cd5ffa0b5fc5dc6a515cb1bce818a3be9679de6f1a9f8860bdfa", "child_ids": [], "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:crl:properties:http_access", "parent_id": "xcsh-docs:data-sources:crl:reference", "path": "documentation/data-sources/crl/properties/http_access/index.md", "provider_name": "crl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["http_access"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/properties/http_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_access for xcsh_crl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_access

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/properties/)
- http_access

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-http_access--path"></a>

### path property

Type: `"string"`. Computed.

CRL File path. CRL file location.

Upstream description:

CRL file location.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/properties/)
- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
