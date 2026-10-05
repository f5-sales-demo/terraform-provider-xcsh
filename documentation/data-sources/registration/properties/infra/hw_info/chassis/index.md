---
page_title: "infra.hw_info.chassis"
subcategory: ""
description: "Chassis information."
xcsh_docs: {"aliases": ["infra hw info chassis"], "body_bytes": 3827, "body_sha256": "sha256:d5dea131d547517ec17aec9a8a300e3b7021445c9190100809f41c20352e1e76", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/chassis/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1001312020030130-3132122321112032-3230333320202221-2132323210011311-1223030323003231-0002333330132110-3102121021030212-0321100020112332", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "chassis"], "schema_version": 1, "sections": [{"aliases": ["infra hw info chassis asset tag"], "anchor": "schema-infra--hw_info--chassis--asset_tag", "description": "Information from /sys/class/dmi/ID/chassis_asset_tag.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "asset_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis serial"], "anchor": "schema-infra--hw_info--chassis--serial", "description": "Information from /sys/class/dmi/ID/chassis_serial.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis type"], "anchor": "schema-infra--hw_info--chassis--type", "description": "Information from /sys/class/dmi/ID/chassis_type.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "type"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info chassis vendor"], "anchor": "schema-infra--hw_info--chassis--vendor", "description": "Information from /sys/class/dmi/ID/chassis_vendor.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info chassis version"], "anchor": "schema-infra--hw_info--chassis--version", "description": "Information from /sys/class/dmi/ID/chassis_version.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:chassis", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "chassis", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/chassis/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Chassis information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.chassis

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.chassis

<a id="section"></a>

Type: `"single"`. Computed.

Chassis Details. Chassis information.

Upstream description:

Chassis information.

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

<a id="schema-infra--hw_info--chassis--asset_tag"></a>

### asset_tag property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Computed.

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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
