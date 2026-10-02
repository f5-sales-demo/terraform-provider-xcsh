---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_policy_based_routing."
xcsh_docs: {"aliases": ["policy based routing"], "body_bytes": 29442, "body_sha256": "sha256:ad5958fb5ca19573b23538721b00a9f139d04031b5ec8330a28d5772e9e512f8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "xcsh-docs:data-sources:policy_based_routing:properties:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:reference", "parent_id": "xcsh-docs:data-sources:policy_based_routing:fundamentals", "path": "documentation/data-sources/policy_based_routing/properties/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["forward proxy pbr"], "anchor": "section", "description": "Network(L3/L4) routing policy rule.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr"], "syntax": "attribute", "type": "object"}, {"aliases": ["forwarding class list"], "anchor": "section", "description": "Ordered list of forwarding Class to be used if source application match and no rule match.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forwarding_class_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["forwarding_class_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:policy_based_routing:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["network pbr"], "anchor": "section", "description": "Network(L3/L4) routing policy rule.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["network_pbr"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_policy_based_routing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
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

Description of the PolicyBasedRouting.

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

- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/): complete subsection reference.

- [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/): complete subsection reference.

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

Name of the PolicyBasedRouting.

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

Type: `"string"`. Required.

Namespace where the PolicyBasedRouting exists.

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

- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-description) |
| `forward_proxy_pbr` | [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_sources/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--namespace) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--forwarding_class_list--tenant) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/any_path/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_prefix_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--path_regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--http_list--http_list--suffix_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--namespace) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--ip_prefix_set--tenant) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/label_selector/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/label_selector/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--label_selector--expressions) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--metadata--description_spec) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--metadata--name) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/prefix_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/prefix_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--prefix_list--prefixes) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/#section) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--exact_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--regex_value) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/#schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--suffix_value) |
| `forwarding_class_list` | [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/#section) |
| `forwarding_class_list.name` | [forwarding_class_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/#schema-forwarding_class_list--name) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/#schema-forwarding_class_list--namespace) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/#schema-forwarding_class_list--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/#schema-namespace) |
| `network_pbr` | [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/#section) |
| `network_pbr.any` | [network_pbr.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/any/#section) |
| `network_pbr.label_selector` | [network_pbr.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/#section) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/label_selector/#schema-network_pbr--label_selector--expressions) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/#section) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_tcp_traffic/#section) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_traffic/#section) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/all_udp_traffic/#section) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/any/#section) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/#section) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/applications/#schema-network_pbr--network_pbr_rules--applications--applications) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/#schema-network_pbr--network_pbr_rules--dns_name) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/#section) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/#schema-network_pbr--network_pbr_rules--forwarding_class_list--name) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/#schema-network_pbr--network_pbr_rules--forwarding_class_list--namespace) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/forwarding_class_list/#schema-network_pbr--network_pbr_rules--forwarding_class_list--tenant) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/#section) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#section) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--kind) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--name) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--namespace) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--tenant) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/#schema-network_pbr--network_pbr_rules--ip_prefix_set--ref--uid) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/metadata/#section) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/metadata/#schema-network_pbr--network_pbr_rules--metadata--description_spec) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/metadata/#schema-network_pbr--network_pbr_rules--metadata--name) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/prefix_list/#section) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/prefix_list/#schema-network_pbr--network_pbr_rules--prefix_list--prefixes) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/protocol_port_range/#section) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/protocol_port_range/#schema-network_pbr--network_pbr_rules--protocol_port_range--port_ranges) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/network_pbr_rules/protocol_port_range/#schema-network_pbr--network_pbr_rules--protocol_port_range--protocol) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/#section) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/prefix_list/#schema-network_pbr--prefix_list--prefixes) |

## Next pages

- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/)
- [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forwarding_class_list/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
