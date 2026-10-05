---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["site mesh group"], "body_bytes": 14873, "body_sha256": "sha256:2a0a01e0755b1f195737d69dd37fce2475ef0f06873403126743dab04eacb2c8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:bfd_disabled", "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "xcsh-docs:data-sources:site_mesh_group:properties:disable_re_fallback", "xcsh-docs:data-sources:site_mesh_group:properties:enable_re_fallback", "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:reference", "parent_id": "xcsh-docs:data-sources:site_mesh_group:fundamentals", "path": "documentation/data-sources/site_mesh_group/properties/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011", "registry_path": "docs/guides/data-sources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["bfd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["bfd enabled"], "anchor": "section", "description": "BFD parameters.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bfd_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable re fallback"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:disable_re_fallback", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_re_fallback"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable re fallback"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:enable_re_fallback", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_re_fallback"], "syntax": "attribute", "type": "object"}, {"aliases": ["full mesh"], "anchor": "section", "description": "Details of Full Mesh Group Type.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["full_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["hub mesh"], "anchor": "section", "description": "Details of Hub Full Mesh Group Type.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["hub_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:site_mesh_group:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["spoke mesh"], "anchor": "section", "description": "Details of Spoke Mesh Group Type.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["spoke_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "Set of sites for which this mesh group config is valid. If 'Type' is Spoke, then it gives set of spoke sites. If 'Type' is Hub, then it gives set of hub sites. If 'Type' is Full Mesh, then it gives set of sites that are connected in full mesh.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_site_mesh_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_disabled/): complete subsection reference.

- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the SiteMeshGroup.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/disable_re_fallback/): complete subsection reference.

- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/enable_re_fallback/): complete subsection reference.

- [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/): complete subsection reference.

- [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the SiteMeshGroup.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
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

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the SiteMeshGroup exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-annotations) |
| `bfd_disabled` | [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_disabled/#section) |
| `bfd_enabled` | [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/#section) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--multiplier) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--receive_interval_milliseconds) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--transmit_interval_milliseconds) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-description) |
| `disable_re_fallback` | [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/disable_re_fallback/#section) |
| `enable_re_fallback` | [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/enable_re_fallback/#section) |
| `full_mesh` | [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/#section) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/control_and_data_plane_mesh/#section) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/data_plane_mesh/#section) |
| `hub_mesh` | [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/#section) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/#section) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/data_plane_mesh/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/#schema-namespace) |
| `spoke_mesh` | [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/#section) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/#section) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/#section) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#section) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--name) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--namespace) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--tenant) |
| `virtual_site` | [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#section) |
| `virtual_site.kind` | [virtual_site.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#schema-virtual_site--kind) |
| `virtual_site.name` | [virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#schema-virtual_site--name) |
| `virtual_site.namespace` | [virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#schema-virtual_site--namespace) |
| `virtual_site.tenant` | [virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#schema-virtual_site--tenant) |
| `virtual_site.uid` | [virtual_site.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/#schema-virtual_site--uid) |

## Next pages

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_disabled/)
- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/bfd_enabled/)
- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/disable_re_fallback/)
- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/enable_re_fallback/)
- [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/)
- [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/)
- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/)
- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/virtual_site/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
