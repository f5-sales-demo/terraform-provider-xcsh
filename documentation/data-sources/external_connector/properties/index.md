---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_external_connector."
xcsh_docs: {"aliases": ["external connector"], "body_bytes": 26077, "body_sha256": "sha256:cdf3ad65f2d3d0a73efea05fa3fff50507a838a1d1d25a79519f53f17039f026", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:ce_site_reference", "xcsh-docs:data-sources:external_connector:properties:gre", "xcsh-docs:data-sources:external_connector:properties:ipsec"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:reference", "parent_id": "xcsh-docs:data-sources:external_connector:fundamentals", "path": "documentation/data-sources/external_connector/properties/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["ce site reference"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ce_site_reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ce_site_reference"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["gre"], "anchor": "section", "description": "External Connector with GRE tunnel.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec"], "anchor": "section", "description": "External Connector with IPsec tunnel.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ipsec"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:external_connector:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_external_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["external_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
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

- [ce_site_reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ce_site_reference/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ExternalConnector.

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

- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/): complete subsection reference.

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

Name of the ExternalConnector.

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

Namespace where the ExternalConnector exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-annotations) |
| `ce_site_reference` | [ce_site_reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ce_site_reference/#section) |
| `ce_site_reference.name` | [ce_site_reference.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--name) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--namespace) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ce_site_reference/#schema-ce_site_reference--tenant) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-description) |
| `gre` | [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/#section) |
| `gre.gre_parameters` | [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/#section) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/#section) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/peer_ip_address/#schema-gre--gre_parameters--peer_ip_address--addr) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/#section) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#section) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--kind) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--name) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--namespace) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--tenant) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/segment/refs/#schema-gre--gre_parameters--segment--refs--uid) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_inside_network/#section) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/site_local_network/#section) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/#section) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--interface) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--node) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/tunnel_eps/#schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/#schema-gre--gre_parameters--tunnel_mtu) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-id) |
| `ipsec` | [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/#section) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/#section) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/#section) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/dpd_keep_alive_timer/#schema-ipsec--ike_parameters--dpd_keep_alive_timer--timeout) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#section) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--name) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--namespace) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase1_profile/#schema-ipsec--ike_parameters--ike_phase1_profile--tenant) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#section) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--name) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--namespace) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/ike_phase2_profile/#schema-ipsec--ike_parameters--ike_phase2_profile--tenant) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/initiator/#section) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/responder/#section) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/#schema-ipsec--ike_parameters--rm_hostname) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv4/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv4/#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv6/#section) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/dual_stack/ipv6/#schema-ipsec--ike_parameters--rm_ip_address--dual_stack--ipv6--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv4/#schema-ipsec--ike_parameters--rm_ip_address--ipv4--addr) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/#section) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/rm_ip_address/ipv6/#schema-ipsec--ike_parameters--rm_ip_address--ipv6--addr) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/use_default_local_ike_id/#section) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/#section) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/#section) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/peer_ip_address/#schema-ipsec--ipsec_tunnel_parameters--peer_ip_address--addr) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#schema-ipsec--ipsec_tunnel_parameters--psk) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#section) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--kind) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--name) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--namespace) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--tenant) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/segment/refs/#schema-ipsec--ipsec_tunnel_parameters--segment--refs--uid) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_inside_network/#section) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/site_local_network/#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#section) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/#schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/#schema-ipsec--ipsec_tunnel_parameters--tunnel_mtu) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/#schema-namespace) |
