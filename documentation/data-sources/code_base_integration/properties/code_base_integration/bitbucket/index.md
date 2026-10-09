---
page_title: "code_base_integration.bitbucket"
subcategory: ""
description: "BitBucket Cloud Integration."
xcsh_docs: {"aliases": ["code base integration bitbucket"], "body_bytes": 1994, "body_sha256": "sha256:6dbe5fb5b1ff40823cc098b8c1f1f94f1ff931943be4faeebe3bff39a56e8906", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket:passwd"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/bitbucket/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "bitbucket"], "schema_version": 1, "sections": [{"aliases": ["code base integration bitbucket passwd"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket:passwd", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket", "passwd"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket username"], "anchor": "schema-code_base_integration--bitbucket--username", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/bitbucket/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "BitBucket Cloud Integration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.bitbucket

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- code_base_integration.bitbucket

<a id="section"></a>

Type: `"single"`. Computed.

BitBucket Cloud Integration.

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

- [passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket/passwd/): complete subsection reference.

<a id="schema-code_base_integration--bitbucket--username"></a>

### username property

Type: `"string"`. Computed.

BitBucket Username. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
