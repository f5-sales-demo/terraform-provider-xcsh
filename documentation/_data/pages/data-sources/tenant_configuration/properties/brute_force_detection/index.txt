---
page_title: "brute_force_detection"
subcategory: ""
description: "Configuration parameter for brute force detection."
xcsh_docs: {"aliases": ["brute force detection"], "body_bytes": 1573, "body_sha256": "sha256:750eef4b7c0ccedfac0398f622e15d8eddbfbc3640cf732bab7a5a6e6fa53ec6", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "parent_id": "xcsh-docs:data-sources:tenant_configuration:reference", "path": "documentation/data-sources/tenant_configuration/properties/brute_force_detection/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0111231300122321-2102101133133132-2333222030301310-1120221323330101-3321220222021220-0023030302011012-0310120012033313-0121321122311110", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["brute_force_detection"], "schema_version": 1, "sections": [{"aliases": ["brute force detection max login failures", "login", "login result", "sign in"], "anchor": "schema-brute_force_detection--max_login_failures", "description": "How many failures before wait is triggered. When login failure count is hit, user will be temporarily locked for a max duration of 15 minutes.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:brute_force_detection", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["brute_force_detection", "max_login_failures"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/brute_force_detection/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Configuration parameter for brute force detection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# brute_force_detection

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- brute_force_detection

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for brute force detection.

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

<a id="schema-brute_force_detection--max_login_failures"></a>

### max_login_failures property

Type: `"number"`. Computed.

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```
