---
page_title: "detection_settings.stage_new_and_updated_signatures"
subcategory: "Security"
description: "detection_settings.stage_new_and_updated_signatures for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2654, "body_sha256": "sha256:cff7ae57c41954f2849374432c620e63c62373a539fff495294e282bb0c7fd09", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "docs/guides/resources--app_firewall--properties--detection_settings--stage_new_and_updated_signatures.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "stage_new_and_updated_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/stage_new_and_updated_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.stage_new_and_updated_signatures for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.stage_new_and_updated_signatures

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- detection_settings.stage_new_and_updated_signatures

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
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
stage_new_and_updated_signatures {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-detection_settings--stage_new_and_updated_signatures--staging_period"></a>

### staging_period property

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Upstream description:

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
