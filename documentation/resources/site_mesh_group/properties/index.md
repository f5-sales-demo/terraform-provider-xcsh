---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["site mesh group"], "body_bytes": 16196, "body_sha256": "sha256:6f9b613df28086a3d34dfbb7479da5ee1a25092e67a2945b13e7685ff509f0b2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:site_mesh_group:properties:bfd_disabled", "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "xcsh-docs:resources:site_mesh_group:properties:disable_re_fallback", "xcsh-docs:resources:site_mesh_group:properties:enable_re_fallback", "xcsh-docs:resources:site_mesh_group:properties:full_mesh", "xcsh-docs:resources:site_mesh_group:properties:hub_mesh", "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "xcsh-docs:resources:site_mesh_group:properties:timeouts", "xcsh-docs:resources:site_mesh_group:properties:virtual_site"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:reference", "parent_id": "xcsh-docs:resources:site_mesh_group:fundamentals", "path": "documentation/resources/site_mesh_group/properties/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["bfd disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bfd_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["bfd enabled"], "anchor": "section", "description": "BFD parameters.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bfd_enabled--multiplier", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}, {"anchor": "schema-bfd_enabled--receive_interval_milliseconds", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}, {"anchor": "schema-bfd_enabled--transmit_interval_milliseconds", "enforcement": "provider-schema", "group": "bfd_enabled:RequiredObjectAttributes:multiplier,receive_interval_milliseconds,transmit_interval_milliseconds", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_enabled", "type": "requires"}], "schema_path": ["bfd_enabled"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable re fallback"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:disable_re_fallback", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_re_fallback"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable re fallback"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:enable_re_fallback", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_re_fallback"], "syntax": "attribute", "type": "object"}, {"aliases": ["full mesh"], "anchor": "section", "description": "Details of Full Mesh Group Type.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "full_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "full_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh:data_plane_mesh", "type": "conflicts"}], "schema_path": ["full_mesh"], "syntax": "block", "type": "object"}, {"aliases": ["hub mesh"], "anchor": "section", "description": "Details of Hub Full Mesh Group Type.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "hub_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "hub_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "type": "conflicts"}], "schema_path": ["hub_mesh"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:site_mesh_group:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["spoke mesh"], "anchor": "section", "description": "Details of Spoke Mesh Group Type.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "spoke_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "spoke_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "type": "conflicts"}], "schema_path": ["spoke_mesh"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:site_mesh_group:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "Set of sites for which this mesh group config is valid. If 'Type' is Spoke, then it gives set of spoke sites. If 'Type' is Hub, then it gives set of hub sites. If 'Type' is Full Mesh, then it gives set of sites that are connected in full mesh.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_site_mesh_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_disabled/): complete subsection reference.

- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/disable_re_fallback/): complete subsection reference.

- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/enable_re_fallback/): complete subsection reference.

- [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/full_mesh/): complete subsection reference.

- [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the Site Mesh Group. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Namespace where the Site Mesh Group is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-annotations) |
| `bfd_disabled` | [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_disabled/#section) |
| `bfd_enabled` | [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/#section) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--multiplier) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--receive_interval_milliseconds) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/#schema-bfd_enabled--transmit_interval_milliseconds) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-disable) |
| `disable_re_fallback` | [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/disable_re_fallback/#section) |
| `enable_re_fallback` | [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/enable_re_fallback/#section) |
| `full_mesh` | [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/full_mesh/#section) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/full_mesh/control_and_data_plane_mesh/#section) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/full_mesh/data_plane_mesh/#section) |
| `hub_mesh` | [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/#section) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/#section) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/data_plane_mesh/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/#schema-namespace) |
| `spoke_mesh` | [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/#section) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/#section) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/#section) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#section) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--name) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--namespace) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/#schema-spoke_mesh--hub_mesh_group--tenant) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/#schema-timeouts--update) |
| `virtual_site` | [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#section) |
| `virtual_site.kind` | [virtual_site.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#schema-virtual_site--kind) |
| `virtual_site.name` | [virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#schema-virtual_site--name) |
| `virtual_site.namespace` | [virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#schema-virtual_site--namespace) |
| `virtual_site.tenant` | [virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#schema-virtual_site--tenant) |
| `virtual_site.uid` | [virtual_site.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/#schema-virtual_site--uid) |

## Next pages

- [bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_disabled/)
- [bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/bfd_enabled/)
- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/disable_re_fallback/)
- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/enable_re_fallback/)
- [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/full_mesh/)
- [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/)
- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/timeouts/)
- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/virtual_site/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
