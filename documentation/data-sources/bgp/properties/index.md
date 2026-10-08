---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bgp."
xcsh_docs: {"aliases": ["bgp"], "body_bytes": 26307, "body_sha256": "sha256:8c874ad155a76d4fa1545366501b5d4d57451e1bd98167691eeb082c5e5e2f1a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:bgp_parameters", "xcsh-docs:data-sources:bgp:properties:peers", "xcsh-docs:data-sources:bgp:properties:where"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:reference", "parent_id": "xcsh-docs:data-sources:bgp:fundamentals", "path": "documentation/data-sources/bgp/properties/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["bgp parameters"], "anchor": "section", "description": "BGP parameters for the local site.", "document_id": "xcsh-docs:data-sources:bgp:properties:bgp_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bgp_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:bgp:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["peers"], "anchor": "section", "description": "List of peers.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers"], "syntax": "attribute", "type": "object"}, {"aliases": ["where"], "anchor": "section", "description": "VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual_site It used to refer site or a group of sites indicated by virtual site.", "document_id": "xcsh-docs:data-sources:bgp:properties:where", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_bgp.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bgpCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [bgp_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the BGP.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

Name of the BGP.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Namespace where the BGP exists.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/): complete subsection reference.

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-annotations) |
| `bgp_parameters` | [bgp_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/#section) |
| `bgp_parameters.asn` | [bgp_parameters.asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/#schema-bgp_parameters--asn) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/from_site/#section) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/#schema-bgp_parameters--ip_address) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/bgp_parameters/local_address/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/#schema-namespace) |
| `peers` | [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/#section) |
| `peers.bfd_disabled` | [peers.bfd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_disabled/#section) |
| `peers.bfd_enabled` | [peers.bfd_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/#section) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/#schema-peers--bfd_enabled--multiplier) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/#schema-peers--bfd_enabled--receive_interval_milliseconds) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/bfd_enabled/#schema-peers--bfd_enabled--transmit_interval_milliseconds) |
| `peers.disable_spec` | [peers.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/disable_spec/#section) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_disabled/#section) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/ebgp_multihop_enabled/#section) |
| `peers.external` | [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#section) |
| `peers.external.address` | [peers.external.address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--address) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--address_ipv6) |
| `peers.external.asn` | [peers.external.asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--asn) |
| `peers.external.default_gateway` | [peers.external.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/default_gateway/#section) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/default_gateway_v6/#section) |
| `peers.external.disable_spec` | [peers.external.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/disable_spec/#section) |
| `peers.external.disable_v6` | [peers.external.disable_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/disable_v6/#section) |
| `peers.external.external_connector` | [peers.external.external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/external_connector/#section) |
| `peers.external.family_inet` | [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/#section) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/disable_spec/#section) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/#section) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/#section) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/#schema-peers--external--family_inet--enable--aggregation--ip_prefix) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/options/#section) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/options/summary_only/#section) |
| `peers.external.from_site` | [peers.external.from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/from_site/#section) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/from_site_v6/#section) |
| `peers.external.interface` | [peers.external.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface/#section) |
| `peers.external.interface.name` | [peers.external.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface/#schema-peers--external--interface--name) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface/#schema-peers--external--interface--namespace) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface/#schema-peers--external--interface--tenant) |
| `peers.external.interface_list` | [peers.external.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/#section) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/#section) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/#schema-peers--external--interface_list--interfaces--name) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/#schema-peers--external--interface_list--interfaces--namespace) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/#schema-peers--external--interface_list--interfaces--tenant) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--md5_auth_key) |
| `peers.external.no_authentication` | [peers.external.no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/no_authentication/#section) |
| `peers.external.port` | [peers.external.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--port) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--subnet_begin_offset) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--subnet_begin_offset_v6) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--subnet_end_offset) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/#schema-peers--external--subnet_end_offset_v6) |
| `peers.label` | [peers.label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/#schema-peers--label) |
| `peers.metadata` | [peers.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/metadata/#section) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/metadata/#schema-peers--metadata--description_spec) |
| `peers.metadata.name` | [peers.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/metadata/#schema-peers--metadata--name) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_disabled/#section) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/passive_mode_enabled/#section) |
| `peers.routing_policies` | [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/#section) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/#section) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/all_nodes/#section) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/inbound/#section) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/#section) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/#schema-peers--routing_policies--route_policy--node_name--node) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#section) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#schema-peers--routing_policies--route_policy--object_refs--kind) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#schema-peers--routing_policies--route_policy--object_refs--name) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#schema-peers--routing_policies--route_policy--object_refs--namespace) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#schema-peers--routing_policies--route_policy--object_refs--tenant) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/#schema-peers--routing_policies--route_policy--object_refs--uid) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/outbound/#section) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |
