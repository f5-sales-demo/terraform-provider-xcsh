---
page_title: "domains"
subcategory: ""
description: "Add and configure testing domains and credentials."
xcsh_docs: {"aliases": ["domains"], "body_bytes": 2973, "body_sha256": "sha256:2b1b5fb741e7f3f1124d087639bdf4c08d8bce02ca756a6119a53f036de109a0", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains", "parent_id": "xcsh-docs:data-sources:api_testing:reference", "path": "documentation/data-sources/api_testing/properties/domains/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains"], "schema_version": 1, "sections": [{"aliases": ["domains allow destructive methods"], "anchor": "schema-domains--allow_destructive_methods", "description": "Enable to allow API Testing to execute against destructive methods. Use with caution as these may modify or DELETE data.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "allow_destructive_methods"], "syntax": "attribute", "type": "bool"}, {"aliases": ["authentication", "credential setup", "credentials", "domains credentials"], "anchor": "section", "description": "Add credentials for API testing to use in the selected environment.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["domains", "credentials"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains domain"], "anchor": "schema-domains--domain", "description": "Add your testing environment domain. Be aware that running tests on a production domain can impact live applications, as API testing cannot distinguish between production and testing environments.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains", "domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Add and configure testing domains and credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["api_testingCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
