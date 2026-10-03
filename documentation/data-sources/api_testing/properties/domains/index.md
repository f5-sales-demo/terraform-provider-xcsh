---
page_title: "domains"
subcategory: ""
description: "Add and configure testing domains and credentials."
xcsh_docs: {"aliases": ["domains"], "body_bytes": 3348, "body_sha256": "sha256:50f4b8f21125cba6cdb4a0283f29d97b8178925fc2d81a4b3644a25132448420", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains", "parent_id": "xcsh-docs:data-sources:api_testing:reference", "path": "documentation/data-sources/api_testing/properties/domains/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains"], "schema_version": 1, "sections": [{"aliases": ["domains allow destructive methods"], "anchor": "schema-domains--allow_destructive_methods", "description": "Enable to allow API Testing to execute against destructive methods. Use with caution as these may modify or DELETE data.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "allow_destructive_methods"], "syntax": "attribute", "type": "bool"}, {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "anchor": "section", "description": "Add credentials for API testing to use in the selected environment.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["domains", "credentials"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains domain"], "anchor": "schema-domains--domain", "description": "Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add and configure testing domains and credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- domains

<a id="section"></a>

Type: `"list"`. Computed.

Add and configure testing domains and credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

## Direct properties

<a id="schema-domains--allow_destructive_methods"></a>

### allow_destructive_methods property

Type: `"bool"`. Computed.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/): complete subsection reference.

<a id="schema-domains--domain"></a>

### domain property

Type: `"string"`. Computed.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

## Next pages

- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
