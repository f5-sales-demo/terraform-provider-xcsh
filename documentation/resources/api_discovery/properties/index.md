---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_api_discovery."
xcsh_docs: {"aliases": ["api discovery"], "body_bytes": 17423, "body_sha256": "sha256:c80bbf98369815b18d5d8af75abf94c497b6ee0a192537ab3f40b571b7cd4461", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_discovery:properties:custom_auth_types", "xcsh-docs:resources:api_discovery:properties:timeouts", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:reference", "parent_id": "xcsh-docs:resources:api_discovery:fundamentals", "path": "documentation/resources/api_discovery/properties/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["authentication", "credential setup", "credentials", "custom auth types"], "anchor": "section", "description": "Select your custom authentication types to be detected in the API discovery.", "document_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_auth_types--parameter_name", "enforcement": "provider-schema", "group": "custom_auth_types:RequiredListObjectAttributes:parameter_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "type": "requires"}], "schema_path": ["custom_auth_types"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:api_discovery:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:api_discovery:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["user defined api discovery policy"], "anchor": "section", "description": "Rules are evaluated sequentially, top to bottom. If no rules are added, all traffic will be discovered or ignored based on the selection in the 'Default Behaviour of the Rule Set' field.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy:ConflictingObjectAttributes:exclusive,inclusive", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:exclusive", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy:ConflictingObjectAttributes:exclusive,inclusive", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "type": "conflicts"}], "schema_path": ["user_defined_api_discovery_policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_api_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
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

- [custom_auth_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/custom_auth_types/): complete subsection reference.

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

Name of the API Discovery. Must be unique within the namespace.

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

Namespace where the API Discovery is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/): complete subsection reference.

- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-annotations) |
| `custom_auth_types` | [custom_auth_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/custom_auth_types/#section) |
| `custom_auth_types.parameter_name` | [custom_auth_types.parameter_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/custom_auth_types/#schema-custom_auth_types--parameter_name) |
| `custom_auth_types.parameter_type` | [custom_auth_types.parameter_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/custom_auth_types/#schema-custom_auth_types--parameter_type) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/#schema-timeouts--update) |
| `user_defined_api_discovery_policy` | [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/#section) |
| `user_defined_api_discovery_policy.discovery_rules` | [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/#section) |
| `user_defined_api_discovery_policy.discovery_rules.labels` | [user_defined_api_discovery_policy.discovery_rules.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/labels/#section) |
| `user_defined_api_discovery_policy.discovery_rules.metadata` | [user_defined_api_discovery_policy.discovery_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/metadata/#section) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.description_spec` | [user_defined_api_discovery_policy.discovery_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/metadata/#schema-user_defined_api_discovery_policy--discovery_rules--metadata--description_spec) |
| `user_defined_api_discovery_policy.discovery_rules.metadata.name` | [user_defined_api_discovery_policy.discovery_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/metadata/#schema-user_defined_api_discovery_policy--discovery_rules--metadata--name) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties` | [user_defined_api_discovery_policy.discovery_rules.rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/archive/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion.ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/ignore/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.field_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/#schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--field_name) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/#schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--location) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.match_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/#schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--match_type) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/#schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--value) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/inclusion/#section) |
| `user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern` | [user_defined_api_discovery_policy.discovery_rules.rule_properties.pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/#schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--pattern) |
| `user_defined_api_discovery_policy.exclusive` | [user_defined_api_discovery_policy.exclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/#section) |
| `user_defined_api_discovery_policy.exclusive.archive` | [user_defined_api_discovery_policy.exclusive.archive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/archive/#section) |
| `user_defined_api_discovery_policy.exclusive.ignore` | [user_defined_api_discovery_policy.exclusive.ignore](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/exclusive/ignore/#section) |
| `user_defined_api_discovery_policy.inclusive` | [user_defined_api_discovery_policy.inclusive](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/#section) |

## Next pages

- [custom_auth_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/custom_auth_types/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/timeouts/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
