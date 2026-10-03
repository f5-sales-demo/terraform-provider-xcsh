---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_view."
xcsh_docs: {"aliases": ["network policy view"], "body_bytes": 23518, "body_sha256": "sha256:b58cdf7c2c7e0c7d2f98b30071d3277c19b0bcd2726fbd15f87da19db2b3d105", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "xcsh-docs:data-sources:network_policy_view:properties:endpoint", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:reference", "parent_id": "xcsh-docs:data-sources:network_policy_view:fundamentals", "path": "documentation/data-sources/network_policy_view/properties/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["egress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections from policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["egress_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint"], "anchor": "section", "description": "Shape of the endpoint choices for a view.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress rules"], "anchor": "section", "description": "Ordered list of rules applied to connections to policy endpoints.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:network_policy_view:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_policy_view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the NetworkPolicyView.

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

- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/): complete subsection reference.

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/): complete subsection reference.

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

Name of the NetworkPolicyView.

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

Type: `"string"`. Optional, Computed.

Namespace where the NetworkPolicyView exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-description) |
| `egress_rules` | [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/#section) |
| `egress_rules.action` | [egress_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/#schema-egress_rules--action) |
| `egress_rules.adv_action` | [egress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/adv_action/#section) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/adv_action/#schema-egress_rules--adv_action--action) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_tcp_traffic/#section) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_traffic/#section) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/all_udp_traffic/#section) |
| `egress_rules.any` | [egress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/any/#section) |
| `egress_rules.applications` | [egress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/applications/#section) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/applications/#schema-egress_rules--applications--applications) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/inside_endpoints/#section) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/#section) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#section) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#schema-egress_rules--ip_prefix_set--ref--kind) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#schema-egress_rules--ip_prefix_set--ref--name) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#schema-egress_rules--ip_prefix_set--ref--namespace) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#schema-egress_rules--ip_prefix_set--ref--tenant) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/#schema-egress_rules--ip_prefix_set--ref--uid) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_matcher/#section) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_matcher/#schema-egress_rules--label_matcher--keys) |
| `egress_rules.label_selector` | [egress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_selector/#section) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/label_selector/#schema-egress_rules--label_selector--expressions) |
| `egress_rules.metadata` | [egress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/metadata/#section) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/metadata/#schema-egress_rules--metadata--description_spec) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/metadata/#schema-egress_rules--metadata--name) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/outside_endpoints/#section) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/prefix_list/#section) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/prefix_list/#schema-egress_rules--prefix_list--prefixes) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/protocol_port_range/#section) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/protocol_port_range/#schema-egress_rules--protocol_port_range--port_ranges) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/protocol_port_range/#schema-egress_rules--protocol_port_range--protocol) |
| `endpoint` | [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/#section) |
| `endpoint.any` | [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/any/#section) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/inside_endpoints/#section) |
| `endpoint.label_selector` | [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/label_selector/#section) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/label_selector/#schema-endpoint--label_selector--expressions) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/outside_endpoints/#section) |
| `endpoint.prefix_list` | [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/prefix_list/#section) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/prefix_list/#schema-endpoint--prefix_list--prefixes) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-id) |
| `ingress_rules` | [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/#section) |
| `ingress_rules.action` | [ingress_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/#schema-ingress_rules--action) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/adv_action/#section) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/adv_action/#schema-ingress_rules--adv_action--action) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_tcp_traffic/#section) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_traffic/#section) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_udp_traffic/#section) |
| `ingress_rules.any` | [ingress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/any/#section) |
| `ingress_rules.applications` | [ingress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/applications/#section) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/applications/#schema-ingress_rules--applications--applications) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/inside_endpoints/#section) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/#section) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#section) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#schema-ingress_rules--ip_prefix_set--ref--kind) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#schema-ingress_rules--ip_prefix_set--ref--name) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#schema-ingress_rules--ip_prefix_set--ref--namespace) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#schema-ingress_rules--ip_prefix_set--ref--tenant) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/ref/#schema-ingress_rules--ip_prefix_set--ref--uid) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_matcher/#section) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_matcher/#schema-ingress_rules--label_matcher--keys) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_selector/#section) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_selector/#schema-ingress_rules--label_selector--expressions) |
| `ingress_rules.metadata` | [ingress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/#section) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/#schema-ingress_rules--metadata--description_spec) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/#schema-ingress_rules--metadata--name) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/outside_endpoints/#section) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/prefix_list/#section) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/prefix_list/#schema-ingress_rules--prefix_list--prefixes) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/#section) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/#schema-ingress_rules--protocol_port_range--port_ranges) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/#schema-ingress_rules--protocol_port_range--protocol) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/#schema-namespace) |

## Next pages

- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/)
- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
