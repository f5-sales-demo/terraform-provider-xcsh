---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_service_policy."
xcsh_docs: {"aliases": ["service policy"], "body_bytes": 77226, "body_sha256": "sha256:1ae79653fffd581b632dbdd44bbeaa3c8fe96169a19b598263e6904285163137", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:allow_all_requests", "xcsh-docs:resources:service_policy:properties:allow_list", "xcsh-docs:resources:service_policy:properties:any_server", "xcsh-docs:resources:service_policy:properties:deny_all_requests", "xcsh-docs:resources:service_policy:properties:deny_list", "xcsh-docs:resources:service_policy:properties:rule_list", "xcsh-docs:resources:service_policy:properties:server_name_matcher", "xcsh-docs:resources:service_policy:properties:server_selector", "xcsh-docs:resources:service_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:reference", "parent_id": "xcsh-docs:resources:service_policy:fundamentals", "path": "documentation/resources/service_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allow all requests"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:allow_all_requests", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_all_requests"], "syntax": "attribute", "type": "object"}, {"aliases": ["allow list"], "anchor": "section", "description": "List of sources. A request belongs to this list if it satisfies any of the match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:allow_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_next_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "allow_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_next_policy", "type": "conflicts"}], "schema_path": ["allow_list"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["any server"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:any_server", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["any_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["deny all requests"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:deny_all_requests", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deny_all_requests"], "syntax": "attribute", "type": "object"}, {"aliases": ["deny list"], "anchor": "section", "description": "List of sources. A request belongs to this list if it satisfies any of the match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:deny_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_allow,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "deny_list:ConflictingObjectAttributes:default_action_deny,default_action_next_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_next_policy", "type": "conflicts"}], "schema_path": ["deny_list"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule list"], "anchor": "section", "description": "Ordered service-policy rules for non-geographic predicates and actions. Do not use country_list for a geo-only rule here: the platform adds match-all any_ip and any_asn selectors on readback, so the rule can match all traffic. Use deny_list or allow_list with country_list for geographic source matching.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list"], "syntax": "block", "type": "object"}, {"aliases": ["server name"], "anchor": "schema-server_name", "description": "Exclusive with The expected name of the server to which the request API is directed. The actual names for the server are extracted from the HTTP Host header and the name of the virtual_host to which the request is directed. If the request is directed to a virtual K8s service, the actual names also contain the name of", "document_id": "xcsh-docs:resources:service_policy:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["login success", "server name matcher", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:service_policy:properties:server_name_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_name_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["server selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:service_policy:properties:server_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-server_selector--expressions", "enforcement": "provider-schema", "group": "server_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:server_selector", "type": "requires"}], "schema_path": ["server_selector"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:service_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_service_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- Property reference

## Direct properties

- [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_all_requests/): complete subsection reference.

- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/): complete subsection reference.

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

- [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/any_server/): complete subsection reference.

- [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_all_requests/): complete subsection reference.

- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/): complete subsection reference.

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

Name of the Service Policy. Must be unique within the namespace.

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

Namespace where the Service Policy is created.

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

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/): complete subsection reference.

<a id="schema-server_name"></a>

### server_name property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is..

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is directed to a virtual K8s service, the actual names also contain the name of that service. The
predicate evaluates to true if any of the actual names is the same as the expected server name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/): complete subsection reference.

- [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_selector/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_all_requests/#section) |
| `allow_list` | [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#section) |
| `allow_list.asn_list` | [allow_list.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_list/#section) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_list/#schema-allow_list--asn_list--as_numbers) |
| `allow_list.asn_set` | [allow_list.asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_set/#section) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_set/#schema-allow_list--asn_set--name) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_set/#schema-allow_list--asn_set--namespace) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/asn_set/#schema-allow_list--asn_set--tenant) |
| `allow_list.country_list` | [allow_list.country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#schema-allow_list--country_list) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/default_action_allow/#section) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/default_action_deny/#section) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/default_action_next_policy/#section) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/ip_prefix_set/#section) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/ip_prefix_set/#schema-allow_list--ip_prefix_set--name) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/ip_prefix_set/#schema-allow_list--ip_prefix_set--namespace) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/ip_prefix_set/#schema-allow_list--ip_prefix_set--tenant) |
| `allow_list.prefix_list` | [allow_list.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/prefix_list/#section) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/prefix_list/#schema-allow_list--prefix_list--prefixes) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#schema-allow_list--tls_fingerprint_classes) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/#schema-allow_list--tls_fingerprint_values) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-annotations) |
| `any_server` | [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/any_server/#section) |
| `deny_all_requests` | [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_all_requests/#section) |
| `deny_list` | [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#section) |
| `deny_list.asn_list` | [deny_list.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_list/#section) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_list/#schema-deny_list--asn_list--as_numbers) |
| `deny_list.asn_set` | [deny_list.asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_set/#section) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_set/#schema-deny_list--asn_set--name) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_set/#schema-deny_list--asn_set--namespace) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/asn_set/#schema-deny_list--asn_set--tenant) |
| `deny_list.country_list` | [deny_list.country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#schema-deny_list--country_list) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/default_action_allow/#section) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/default_action_deny/#section) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/default_action_next_policy/#section) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/ip_prefix_set/#section) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/ip_prefix_set/#schema-deny_list--ip_prefix_set--name) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/ip_prefix_set/#schema-deny_list--ip_prefix_set--namespace) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/ip_prefix_set/#schema-deny_list--ip_prefix_set--tenant) |
| `deny_list.prefix_list` | [deny_list.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/prefix_list/#section) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/prefix_list/#schema-deny_list--prefix_list--prefixes) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#schema-deny_list--tls_fingerprint_classes) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/#schema-deny_list--tls_fingerprint_values) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-namespace) |
| `rule_list` | [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/#section) |
| `rule_list.rules` | [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/#section) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/metadata/#section) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--description_spec) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/metadata/#schema-rule_list--rules--metadata--name) |
| `rule_list.rules.spec` | [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/#section) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/#schema-rule_list--rules--spec--action) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_asn/#section) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_client/#section) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/any_ip/#section) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/#section) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/#schema-rule_list--rules--spec--api_group_matcher--invert_matcher) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/#schema-rule_list--rules--spec--api_group_matcher--match) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/#section) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_not_present/#section) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/check_present/#section) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/#schema-rule_list--rules--spec--arg_matchers--invert_matcher) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/#section) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/#schema-rule_list--rules--spec--arg_matchers--item--exact_values) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/#schema-rule_list--rules--spec--arg_matchers--item--regex_values) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/item/#schema-rule_list--rules--spec--arg_matchers--item--transformers) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/arg_matchers/#schema-rule_list--rules--spec--arg_matchers--name) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/#section) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_list/#schema-rule_list--rules--spec--asn_list--as_numbers) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/#section) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#section) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#schema-rule_list--rules--spec--asn_matcher--asn_sets--kind) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#schema-rule_list--rules--spec--asn_matcher--asn_sets--name) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#schema-rule_list--rules--spec--asn_matcher--asn_sets--namespace) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#schema-rule_list--rules--spec--asn_matcher--asn_sets--tenant) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/asn_matcher/asn_sets/#schema-rule_list--rules--spec--asn_matcher--asn_sets--uid) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/#section) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/#schema-rule_list--rules--spec--body_matcher--exact_values) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/#schema-rule_list--rules--spec--body_matcher--regex_values) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/body_matcher/#schema-rule_list--rules--spec--body_matcher--transformers) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/#section) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/bot_skip_processing/#section) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/bot_action/none/#section) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/#schema-rule_list--rules--spec--client_name) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/#section) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/#schema-rule_list--rules--spec--client_name_matcher--exact_values) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/#schema-rule_list--rules--spec--client_name_matcher--regex_values) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_name_matcher/#schema-rule_list--rules--spec--client_name_matcher--transformers) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/#section) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/client_selector/#schema-rule_list--rules--spec--client_selector--expressions) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/#section) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/check_not_present/#section) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/check_present/#section) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/#schema-rule_list--rules--spec--cookie_matchers--invert_matcher) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/item/#section) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/item/#schema-rule_list--rules--spec--cookie_matchers--item--exact_values) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/item/#schema-rule_list--rules--spec--cookie_matchers--item--regex_values) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/item/#schema-rule_list--rules--spec--cookie_matchers--item--transformers) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/cookie_matchers/#schema-rule_list--rules--spec--cookie_matchers--name) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/#section) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/#schema-rule_list--rules--spec--domain_matcher--exact_values) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/#schema-rule_list--rules--spec--domain_matcher--regex_values) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/domain_matcher/#schema-rule_list--rules--spec--domain_matcher--transformers) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/#schema-rule_list--rules--spec--expiration_timestamp) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/#section) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/check_not_present/#section) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/check_present/#section) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/#schema-rule_list--rules--spec--headers--invert_matcher) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/item/#section) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/item/#schema-rule_list--rules--spec--headers--item--exact_values) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/item/#schema-rule_list--rules--spec--headers--item--regex_values) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/item/#schema-rule_list--rules--spec--headers--item--transformers) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/headers/#schema-rule_list--rules--spec--headers--name) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/#section) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/#schema-rule_list--rules--spec--http_method--invert_matcher) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/http_method/#schema-rule_list--rules--spec--http_method--methods) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/#section) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/#schema-rule_list--rules--spec--ip_matcher--invert_matcher) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#section) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#schema-rule_list--rules--spec--ip_matcher--prefix_sets--kind) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#schema-rule_list--rules--spec--ip_matcher--prefix_sets--name) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#schema-rule_list--rules--spec--ip_matcher--prefix_sets--namespace) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#schema-rule_list--rules--spec--ip_matcher--prefix_sets--tenant) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/#schema-rule_list--rules--spec--ip_matcher--prefix_sets--uid) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/#section) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/#schema-rule_list--rules--spec--ip_prefix_list--invert_match) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_prefix_list/#schema-rule_list--rules--spec--ip_prefix_list--ip_prefixes) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/#section) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_threat_category_list/#schema-rule_list--rules--spec--ip_threat_category_list--ip_threat_categories) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/#section) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ja4_tls_fingerprint/#schema-rule_list--rules--spec--ja4_tls_fingerprint--exact_values) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/#section) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/check_not_present/#section) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/check_present/#section) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/#schema-rule_list--rules--spec--jwt_claims--invert_matcher) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/item/#section) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/item/#schema-rule_list--rules--spec--jwt_claims--item--exact_values) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/item/#schema-rule_list--rules--spec--jwt_claims--item--regex_values) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/item/#schema-rule_list--rules--spec--jwt_claims--item--transformers) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/jwt_claims/#schema-rule_list--rules--spec--jwt_claims--name) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/#section) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/label_matcher/#schema-rule_list--rules--spec--label_matcher--keys) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/#schema-rule_list--rules--spec--log_rule_evaluation) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/#section) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/default/#section) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/skip_processing/#section) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#section) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--encoded_path_matcher) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--exact_values) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--invert_matcher) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--prefix_values) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--regex_values) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--suffix_values) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/path/#schema-rule_list--rules--spec--path--transformers) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/#section) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/#schema-rule_list--rules--spec--port_matcher--invert_matcher) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/port_matcher/#schema-rule_list--rules--spec--port_matcher--ports) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/#section) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/check_not_present/#section) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/check_present/#section) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/#schema-rule_list--rules--spec--query_params--invert_matcher) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/item/#section) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/item/#schema-rule_list--rules--spec--query_params--item--exact_values) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/item/#schema-rule_list--rules--spec--query_params--item--regex_values) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/item/#schema-rule_list--rules--spec--query_params--item--transformers) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/query_params/#schema-rule_list--rules--spec--query_params--key) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_cookie_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_count_none/#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_cookie_key_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_key_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_cookie_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_value_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_header_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_header_count_none/#section) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_header_key_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_header_key_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_header_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_header_value_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_parameter_count_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_parameter_count_none/#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_parameter_name_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_parameter_name_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_parameter_value_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_parameter_value_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_query_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_query_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_request_line_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_request_line_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_request_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_request_size_none/#section) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/#schema-rule_list--rules--spec--request_constraints--max_url_size_exceeds) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_url_size_none/#section) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/#section) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_any/#section) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/#section) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/segments/#section) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/segments/#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--name) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/segments/#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--namespace) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/segments/#schema-rule_list--rules--spec--segment_policy--dst_segments--segments--tenant) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/intra_segment/#section) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_any/#section) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/#section) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/#section) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/#schema-rule_list--rules--spec--segment_policy--src_segments--segments--name) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/#schema-rule_list--rules--spec--segment_policy--src_segments--segments--namespace) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/#schema-rule_list--rules--spec--segment_policy--src_segments--segments--tenant) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/#section) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/#schema-rule_list--rules--spec--tls_fingerprint_matcher--classes) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/#schema-rule_list--rules--spec--tls_fingerprint_matcher--exact_values) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/tls_fingerprint_matcher/#schema-rule_list--rules--spec--tls_fingerprint_matcher--excluded_values) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/#section) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/#schema-rule_list--rules--spec--user_identity_matcher--exact_values) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/user_identity_matcher/#schema-rule_list--rules--spec--user_identity_matcher--regex_values) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts--exclude_attack_type) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/#section) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--context) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--context_name) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/#schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts--exclude_violation) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/none/#section) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/waf_skip_processing/#section) |
| `server_name` | [server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/#schema-server_name) |
| `server_name_matcher` | [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/#section) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/#schema-server_name_matcher--exact_values) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/#schema-server_name_matcher--regex_values) |
| `server_selector` | [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_selector/#section) |
| `server_selector.expressions` | [server_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_selector/#schema-server_selector--expressions) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [allow_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_all_requests/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/allow_list/)
- [any_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/any_server/)
- [deny_all_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_all_requests/)
- [deny_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/deny_list/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [server_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_name_matcher/)
- [server_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/server_selector/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/timeouts/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
