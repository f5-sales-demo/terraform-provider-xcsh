---
page_title: "re_select.specific_re"
subcategory: ""
description: "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs."
xcsh_docs: {"aliases": ["re select specific re"], "body_bytes": 2430, "body_sha256": "sha256:aa50fb08546450b60fc9b5406f0c76fa0b24e823e1c8ff7cb965da9c4ab9aaf4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "path": "documentation/data-sources/securemesh_site_v2/properties/re_select/specific_re/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2230220011013020-1013300303333213-0023102001132331-1320101332011200-1033032011123222-1321012111200110-1122113323120212-1131130233033003", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select", "specific_re"], "schema_version": 1, "sections": [{"aliases": ["re select specific re backup re"], "anchor": "schema-re_select--specific_re--backup_re", "description": "Select backup RE for this site, cannot be the same as Primary RE.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_re", "backup_re"], "syntax": "attribute", "type": "string"}, {"aliases": ["re select specific re primary re"], "anchor": "schema-re_select--specific_re--primary_re", "description": "Select primary RE for this site.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_re", "primary_re"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/re_select/specific_re/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select.specific_re

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [re_select](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/re_select/)
- re_select.specific_re

<a id="section"></a>

Type: `"single"`. Computed.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

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

<a id="schema-re_select--specific_re--backup_re"></a>

### backup_re property

Type: `"string"`. Computed.

Select backup RE for this site, cannot be the same as Primary RE.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="schema-re_select--specific_re--primary_re"></a>

### primary_re property

Type: `"string"`. Computed.

Primary RE Geography. Select primary RE for this site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
