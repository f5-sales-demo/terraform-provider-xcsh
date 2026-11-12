---
page_title: "detection_settings.stage_new_signatures"
subcategory: "Security"
description: "Attack Signatures staging configuration."
xcsh_docs: {"aliases": ["detection settings stage new signatures"], "body_bytes": 2529, "body_sha256": "sha256:3dbd9a632eb0ec745f8b193c4295ba04a8c72492ac53bbd339c2ea35e7fcbd48", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_signatures", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "documentation/resources/app_firewall/properties/detection_settings/stage_new_signatures/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1202022030301321-3210231222102302-0033311203113112-3202001300100302-3301322321002130-1113331033222120-2001233200331100-0210022113121301", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "schema-detection_settings--stage_new_signatures--staging_period", "enforcement": "provider-schema", "group": "detection_settings.stage_new_signatures:RequiredObjectAttributes:staging_period", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_signatures", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "stage_new_signatures"], "schema_version": 1, "sections": [{"aliases": ["detection settings stage new signatures staging period"], "anchor": "schema-detection_settings--stage_new_signatures--staging_period", "description": "Define staging period in days. The default staging period is 7 days and the max supported staging period is 20 days.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "stage_new_signatures", "staging_period"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/stage_new_signatures/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Attack Signatures staging configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["app_firewallCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.stage_new_signatures

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
- detection_settings.stage_new_signatures

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("staging_period")}
```

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
stage_new_signatures {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-detection_settings--stage_new_signatures--staging_period"></a>

### staging_period property

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Staging period in days. Default 7, max 20. Applies to both stage_new_and_updated_signatures and stage_new_signatures.",
      "source": "api-probed",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```
