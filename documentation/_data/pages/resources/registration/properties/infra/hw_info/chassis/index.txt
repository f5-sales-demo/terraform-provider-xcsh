---
page_title: "infra.hw_info.chassis"
subcategory: ""
description: "Chassis information."
xcsh_docs: {"aliases": ["infra hw info chassis"], "body_bytes": 3637, "body_sha256": "sha256:7d52a4e27e0f90976e7b172e301ac7d2c557bb0fb6a1480fcde8deea7e552070", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/chassis/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0323323131133211-0011220111233113-3322310102330001-3332123213202301-1222121212331031-1010111023003000-1112303301202203-3232031100203230", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "chassis"], "schema_version": 1, "sections": [{"aliases": ["infra hw info chassis asset tag"], "anchor": "schema-infra--hw_info--chassis--asset_tag", "description": "Information from /sys/class/dmi/ID/chassis_asset_tag.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis serial"], "anchor": "schema-infra--hw_info--chassis--serial", "description": "Information from /sys/class/dmi/ID/chassis_serial.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis type"], "anchor": "schema-infra--hw_info--chassis--type", "description": "Information from /sys/class/dmi/ID/chassis_type.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "type"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info chassis vendor"], "anchor": "schema-infra--hw_info--chassis--vendor", "description": "Information from /sys/class/dmi/ID/chassis_vendor.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis version"], "anchor": "schema-infra--hw_info--chassis--version", "description": "Information from /sys/class/dmi/ID/chassis_version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:chassis", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Chassis information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["registrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.chassis

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.chassis

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Chassis Details. Chassis information.

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
chassis {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--chassis--asset_tag"></a>

### asset_tag property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-infra--hw_info--chassis--serial"></a>

### serial property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_serial.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-infra--hw_info--chassis--type"></a>

### type property

Type: `"number"`. Optional.

Information from /sys/class/dmi/ID/chassis\_type.

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

<a id="schema-infra--hw_info--chassis--vendor"></a>

### vendor property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_vendor.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-infra--hw_info--chassis--version"></a>

### version property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_version.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
