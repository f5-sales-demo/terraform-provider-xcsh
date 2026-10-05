---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nat_policy."
xcsh_docs: {"aliases": ["nat policy"], "body_bytes": 26574, "body_sha256": "sha256:6fb6e02339ac027e6ea810ec686de1c68a5b9183588071f8c2e05660666e449d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules", "xcsh-docs:resources:nat_policy:properties:site", "xcsh-docs:resources:nat_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:reference", "parent_id": "xcsh-docs:resources:nat_policy:fundamentals", "path": "documentation/resources/nat_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:nat_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules"], "anchor": "section", "description": "List of rules to apply under the NAT Policy. Rule that matches first would be applied.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:cloud_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,segment", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:segment,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:cloud_connect,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:node_interface,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules:ConflictingListObjectAttributes:segment,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:virtual_network", "type": "conflicts"}, {"anchor": "schema-rules--name", "enforcement": "provider-schema", "group": "rules:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules", "type": "requires"}], "schema_path": ["rules"], "syntax": "block", "type": "object"}, {"aliases": ["site"], "anchor": "section", "description": "Reference to Site Object.", "document_id": "xcsh-docs:resources:nat_policy:properties:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site:RequiredObjectAttributes:refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:site:refs", "type": "requires"}], "schema_path": ["site"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:nat_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_nat_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
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

Name of the NAT Policy. Must be unique within the namespace.

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

Namespace where the NAT Policy is created.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/): complete subsection reference.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/#schema-namespace) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/#section) |
| `rules.action` | [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/#section) |
| `rules.action.dynamic` | [rules.action.dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/#section) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/#section) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#section) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--kind) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--name) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--namespace) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--tenant) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/elastic_ips/refs/#schema-rules--action--dynamic--elastic_ips--refs--uid) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/pools/#section) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/dynamic/pools/#schema-rules--action--dynamic--pools--prefixes) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/action/#schema-rules--action--virtual_cidr) |
| `rules.cloud_connect` | [rules.cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/#section) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#section) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--kind) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--name) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--namespace) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--tenant) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/cloud_connect/refs/#schema-rules--cloud_connect--refs--uid) |
| `rules.criteria` | [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#section) |
| `rules.criteria.any` | [rules.criteria.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/any/#section) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#schema-rules--criteria--destination_cidr) |
| `rules.criteria.icmp` | [rules.criteria.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/icmp/#section) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_inside_network/#section) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_network/#section) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/#schema-rules--criteria--source_cidr) |
| `rules.criteria.tcp` | [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/#section) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#section) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/no_port_match/#section) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#schema-rules--criteria--tcp--destination_port--port) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/#schema-rules--criteria--tcp--destination_port--port_ranges) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#section) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/no_port_match/#section) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#schema-rules--criteria--tcp--source_port--port) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/#schema-rules--criteria--tcp--source_port--port_ranges) |
| `rules.criteria.udp` | [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/#section) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#section) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/no_port_match/#section) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#schema-rules--criteria--udp--destination_port--port) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/destination_port/#schema-rules--criteria--udp--destination_port--port_ranges) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#section) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/no_port_match/#section) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#schema-rules--criteria--udp--source_port--port) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/source_port/#schema-rules--criteria--udp--source_port--port_ranges) |
| `rules.disable_spec` | [rules.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/disable_spec/#section) |
| `rules.enable` | [rules.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/enable/#section) |
| `rules.name` | [rules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/#schema-rules--name) |
| `rules.node_interface` | [rules.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/#section) |
| `rules.node_interface.list` | [rules.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/#section) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#section) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--kind) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--name) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--namespace) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--tenant) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/interface/#schema-rules--node_interface--list--interface--uid) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/node_interface/list/#schema-rules--node_interface--list--node) |
| `rules.segment` | [rules.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/#section) |
| `rules.segment.refs` | [rules.segment.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#section) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--kind) |
| `rules.segment.refs.name` | [rules.segment.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--name) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--namespace) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--tenant) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/segment/refs/#schema-rules--segment--refs--uid) |
| `rules.virtual_network` | [rules.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/#section) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#section) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--kind) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--name) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--namespace) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--tenant) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/virtual_network/refs/#schema-rules--virtual_network--refs--uid) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/#section) |
| `site.refs` | [site.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#section) |
| `site.refs.kind` | [site.refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--kind) |
| `site.refs.name` | [site.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--name) |
| `site.refs.namespace` | [site.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--namespace) |
| `site.refs.tenant` | [site.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--tenant) |
| `site.refs.uid` | [site.refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/refs/#schema-site--refs--uid) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/site/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/timeouts/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
