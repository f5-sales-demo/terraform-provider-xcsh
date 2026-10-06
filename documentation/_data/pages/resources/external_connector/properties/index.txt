---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_external_connector."
xcsh_docs: {"aliases": ["external connector"], "body_bytes": 27659, "body_sha256": "sha256:1ad09644d97102b8a643e9d8ec3fb13e2603f39590dbd254b017739c71794d5b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:ce_site_reference", "xcsh-docs:resources:external_connector:properties:gre", "xcsh-docs:resources:external_connector:properties:ipsec", "xcsh-docs:resources:external_connector:properties:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:reference", "parent_id": "xcsh-docs:resources:external_connector:fundamentals", "path": "documentation/resources/external_connector/properties/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["ce site reference"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:external_connector:properties:ce_site_reference", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ce_site_reference--name", "enforcement": "provider-schema", "group": "ce_site_reference:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:ce_site_reference", "type": "requires"}], "schema_path": ["ce_site_reference"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["gre"], "anchor": "section", "description": "External Connector with GRE tunnel.", "document_id": "xcsh-docs:resources:external_connector:properties:gre", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec"], "anchor": "section", "description": "External Connector with IPsec tunnel.", "document_id": "xcsh-docs:resources:external_connector:properties:ipsec", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:external_connector:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_external_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

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

- [ce_site_reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ce_site_reference/): complete subsection reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Additional upstream details:

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

- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the External Connector. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Namespace where the External Connector is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-annotations) |
| `ce_site_reference` | [ce_site_reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ce_site_reference/#section) |
| `ce_site_reference.name` | [ce_site_reference.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--name) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--namespace) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--tenant) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-disable) |
| `gre` | [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/#section) |
| `gre.gre_parameters` | [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/#section) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/peer_ip_address/#section) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/peer_ip_address/#schema-gre--gre_parameters--peer_ip_address--addr) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/#section) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#section) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--kind) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--name) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--namespace) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--tenant) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--uid) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_inside_network/#section) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/site_local_network/#section) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/#section) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--interface) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--node) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/#schema-gre--gre_parameters--tunnel_mtu) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-id) |
| `ipsec` | [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/#section) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/#section) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/#schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#section) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--name) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--namespace) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--tenant) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#section) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--name) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--namespace) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--tenant) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/initiator/#section) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/responder/#section) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/#schema-ipsec--ike_parameters--rm_hostname) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv4/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv4/#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv6/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv6/#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/#schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/#schema-ipsec--ike_parameters--rm_ip_address--ipv6--addr) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_local_ike_id/#section) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/#section) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/#schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#schema-ipsec--ipsec_tunnel_parameters--psk) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--kind) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--name) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--namespace) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--tenant) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--uid) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_inside_network/#section) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_network/#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/timeouts/#schema-timeouts--update) |
