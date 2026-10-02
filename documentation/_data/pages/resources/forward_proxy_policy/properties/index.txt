---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": ["forward proxy policy"], "body_bytes": 34314, "body_sha256": "sha256:792a07212f143406bfe6024f09091500147f18d26e70eb67311ca6ac606e4f07", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:allow_all", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "xcsh-docs:resources:forward_proxy_policy:properties:any_proxy", "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "xcsh-docs:resources:forward_proxy_policy:properties:drp_http_connect", "xcsh-docs:resources:forward_proxy_policy:properties:network_connector", "xcsh-docs:resources:forward_proxy_policy:properties:proxy_label_selector", "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "xcsh-docs:resources:forward_proxy_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:reference", "parent_id": "xcsh-docs:resources:forward_proxy_policy:fundamentals", "path": "documentation/resources/forward_proxy_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["allow list"], "anchor": "section", "description": "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "type": "conflicts"}], "schema_path": ["allow_list"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["any proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:any_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["any_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["deny list"], "anchor": "section", "description": "URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}], "schema_path": ["deny_list"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["drp http connect"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:drp_http_connect", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["drp_http_connect"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:forward_proxy_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["network connector"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:network_connector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-network_connector--name", "enforcement": "provider-schema", "group": "network_connector:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:network_connector", "type": "requires"}], "schema_path": ["network_connector"], "syntax": "block", "type": "object"}, {"aliases": ["proxy label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:proxy_label_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_label_selector--expressions", "enforcement": "provider-schema", "group": "proxy_label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:proxy_label_selector", "type": "requires"}], "schema_path": ["proxy_label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["rule list"], "anchor": "section", "description": "List of custom rules.", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "type": "requires"}], "schema_path": ["rule_list"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:forward_proxy_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_forward_proxy_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- Property reference

## Direct properties

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_all/): complete subsection reference.

- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/): complete subsection reference.

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

- [any_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/any_proxy/): complete subsection reference.

- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/): complete subsection reference.

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

- [drp_http_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/drp_http_connect/): complete subsection reference.

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

Name of the Forward Proxy Policy. Must be unique within the namespace.

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

Namespace where the Forward Proxy Policy is created.

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

- [network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/): complete subsection reference.

- [proxy_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/proxy_label_selector/): complete subsection reference.

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_all/#section) |
| `allow_list` | [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/#section) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/default_action_allow/#section) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/default_action_deny/#section) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/default_action_next_policy/#section) |
| `allow_list.dest_list` | [allow_list.dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/dest_list/#section) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/dest_list/#schema-allow_list--dest_list--ipv6_prefixes) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/dest_list/#schema-allow_list--dest_list--port_ranges) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/dest_list/#schema-allow_list--dest_list--prefixes) |
| `allow_list.http_list` | [allow_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#section) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/any_path/#section) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--exact_value) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--path_exact_value) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--path_prefix_value) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--path_regex_value) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--regex_value) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/http_list/#schema-allow_list--http_list--suffix_value) |
| `allow_list.tls_list` | [allow_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/tls_list/#section) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/tls_list/#schema-allow_list--tls_list--exact_value) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/tls_list/#schema-allow_list--tls_list--regex_value) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/tls_list/#schema-allow_list--tls_list--suffix_value) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-annotations) |
| `any_proxy` | [any_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/any_proxy/#section) |
| `deny_list` | [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/#section) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_allow/#section) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_deny/#section) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/default_action_next_policy/#section) |
| `deny_list.dest_list` | [deny_list.dest_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/#section) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/#schema-deny_list--dest_list--ipv6_prefixes) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/#schema-deny_list--dest_list--port_ranges) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/dest_list/#schema-deny_list--dest_list--prefixes) |
| `deny_list.http_list` | [deny_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#section) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/any_path/#section) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--exact_value) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--path_exact_value) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--path_prefix_value) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--path_regex_value) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--regex_value) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/http_list/#schema-deny_list--http_list--suffix_value) |
| `deny_list.tls_list` | [deny_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/#section) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/#schema-deny_list--tls_list--exact_value) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/#schema-deny_list--tls_list--regex_value) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/tls_list/#schema-deny_list--tls_list--suffix_value) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-disable) |
| `drp_http_connect` | [drp_http_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/drp_http_connect/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/#schema-namespace) |
| `network_connector` | [network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/#section) |
| `network_connector.name` | [network_connector.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/#schema-network_connector--name) |
| `network_connector.namespace` | [network_connector.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/#schema-network_connector--namespace) |
| `network_connector.tenant` | [network_connector.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/#schema-network_connector--tenant) |
| `proxy_label_selector` | [proxy_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/proxy_label_selector/#section) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/proxy_label_selector/#schema-proxy_label_selector--expressions) |
| `rule_list` | [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/#section) |
| `rule_list.rules` | [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/#section) |
| `rule_list.rules.action` | [rule_list.rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/#schema-rule_list--rules--action) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_destinations/#section) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/all_sources/#section) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_list/#section) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_list/#schema-rule_list--rules--dst_asn_list--as_numbers) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/#section) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/#schema-rule_list--rules--dst_asn_set--name) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/#schema-rule_list--rules--dst_asn_set--namespace) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_asn_set/#schema-rule_list--rules--dst_asn_set--tenant) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/#section) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/#schema-rule_list--rules--dst_ip_prefix_set--name) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/#schema-rule_list--rules--dst_ip_prefix_set--namespace) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_ip_prefix_set/#schema-rule_list--rules--dst_ip_prefix_set--tenant) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_label_selector/#section) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_label_selector/#schema-rule_list--rules--dst_label_selector--expressions) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_prefix_list/#section) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/dst_prefix_list/#schema-rule_list--rules--dst_prefix_list--prefixes) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/#section) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#section) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/any_path/#section) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--exact_value) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--path_exact_value) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--path_prefix_value) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--path_regex_value) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--regex_value) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/#schema-rule_list--rules--http_list--http_list--suffix_value) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/#section) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/#schema-rule_list--rules--ip_prefix_set--name) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/#schema-rule_list--rules--ip_prefix_set--namespace) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/ip_prefix_set/#schema-rule_list--rules--ip_prefix_set--tenant) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/label_selector/#section) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/label_selector/#schema-rule_list--rules--label_selector--expressions) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/metadata/#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/no_http_connect_port/#section) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/port_matcher/#section) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/port_matcher/#schema-rule_list--rules--port_matcher--invert_matcher) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/port_matcher/#schema-rule_list--rules--port_matcher--ports) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/prefix_list/#section) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/prefix_list/#schema-rule_list--rules--prefix_list--prefixes) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/#section) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/#section) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/#schema-rule_list--rules--tls_list--tls_list--exact_value) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/#schema-rule_list--rules--tls_list--tls_list--regex_value) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/#schema-rule_list--rules--tls_list--tls_list--suffix_value) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/url_category_list/#section) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/url_category_list/#schema-rule_list--rules--url_category_list--url_categories) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_all/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/allow_list/)
- [any_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/any_proxy/)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/deny_list/)
- [drp_http_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/drp_http_connect/)
- [network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/network_connector/)
- [proxy_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/proxy_label_selector/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/timeouts/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
