---
page_title: "infra.sw_info"
subcategory: ""
description: "SWInfo holds information about sw version."
xcsh_docs: {"aliases": ["infra sw info"], "body_bytes": 1327, "body_sha256": "sha256:c030266864fbaa0592578875f32626fe8c208a738d6c52bed74a649f4dd82d44", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:sw_info", "parent_id": "xcsh-docs:data-sources:registration:properties:infra", "path": "documentation/data-sources/registration/properties/infra/sw_info/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1231311000020020-2022312100213123-2023300303302201-1200001213120331-1211131030203221-1032233033010103-3210023002233002-0031113231233303", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "sw_info"], "schema_version": 1, "sections": [{"aliases": ["infra sw info sw version"], "anchor": "schema-infra--sw_info--sw_version", "description": "SW Version in the site.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:sw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "sw_info", "sw_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/sw_info/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "SWInfo holds information about sw version.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["registrationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.sw_info

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- infra.sw_info

<a id="section"></a>

Type: `"single"`. Computed.

SWInfo holds information about sw version.

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

<a id="schema-infra--sw_info--sw_version"></a>

### sw_version property

Type: `"string"`. Computed.

SW Version. SW Version in the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
