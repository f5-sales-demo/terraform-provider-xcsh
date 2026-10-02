---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_network_policy."
xcsh_docs: {"aliases": ["network policy"], "body_bytes": 25986, "body_sha256": "sha256:7d4ed867349f305b7f5747ad825e99be6b21906d28c13fb1ecd1a315dd6fc60f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:endpoint", "xcsh-docs:resources:network_policy:properties:rules", "xcsh-docs:resources:network_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:reference", "parent_id": "xcsh-docs:resources:network_policy:fundamentals", "path": "documentation/resources/network_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["endpoint"], "anchor": "section", "description": "Shape of the endpoint choices for a view.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}], "schema_path": ["endpoint"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:network_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules"], "anchor": "section", "description": "Shape of Rule Choice.", "document_id": "xcsh-docs:resources:network_policy:properties:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:network_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
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

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/): complete subsection reference.

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

Name of the Network Policy. Must be unique within the namespace.

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

Namespace where the Network Policy is created.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-disable) |
| `endpoint` | [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/#section) |
| `endpoint.any` | [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/any/#section) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/inside_endpoints/#section) |
| `endpoint.label_selector` | [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/#section) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/#schema-endpoint--label_selector--expressions) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/outside_endpoints/#section) |
| `endpoint.prefix_list` | [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/#section) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/#schema-endpoint--prefix_list--prefixes) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/#schema-namespace) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/#section) |
| `rules.egress_rules` | [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/#section) |
| `rules.egress_rules.action` | [rules.egress_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/#schema-rules--egress_rules--action) |
| `rules.egress_rules.adv_action` | [rules.egress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/adv_action/#section) |
| `rules.egress_rules.adv_action.action` | [rules.egress_rules.adv_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/adv_action/#schema-rules--egress_rules--adv_action--action) |
| `rules.egress_rules.all_tcp_traffic` | [rules.egress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/all_tcp_traffic/#section) |
| `rules.egress_rules.all_traffic` | [rules.egress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/all_traffic/#section) |
| `rules.egress_rules.all_udp_traffic` | [rules.egress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/all_udp_traffic/#section) |
| `rules.egress_rules.any` | [rules.egress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/any/#section) |
| `rules.egress_rules.applications` | [rules.egress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/applications/#section) |
| `rules.egress_rules.applications.applications` | [rules.egress_rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/applications/#schema-rules--egress_rules--applications--applications) |
| `rules.egress_rules.inside_endpoints` | [rules.egress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/inside_endpoints/#section) |
| `rules.egress_rules.ip_prefix_set` | [rules.egress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/#section) |
| `rules.egress_rules.ip_prefix_set.ref` | [rules.egress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#section) |
| `rules.egress_rules.ip_prefix_set.ref.kind` | [rules.egress_rules.ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#schema-rules--egress_rules--ip_prefix_set--ref--kind) |
| `rules.egress_rules.ip_prefix_set.ref.name` | [rules.egress_rules.ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#schema-rules--egress_rules--ip_prefix_set--ref--name) |
| `rules.egress_rules.ip_prefix_set.ref.namespace` | [rules.egress_rules.ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#schema-rules--egress_rules--ip_prefix_set--ref--namespace) |
| `rules.egress_rules.ip_prefix_set.ref.tenant` | [rules.egress_rules.ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#schema-rules--egress_rules--ip_prefix_set--ref--tenant) |
| `rules.egress_rules.ip_prefix_set.ref.uid` | [rules.egress_rules.ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/#schema-rules--egress_rules--ip_prefix_set--ref--uid) |
| `rules.egress_rules.label_matcher` | [rules.egress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/label_matcher/#section) |
| `rules.egress_rules.label_matcher.keys` | [rules.egress_rules.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/label_matcher/#schema-rules--egress_rules--label_matcher--keys) |
| `rules.egress_rules.label_selector` | [rules.egress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/label_selector/#section) |
| `rules.egress_rules.label_selector.expressions` | [rules.egress_rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/label_selector/#schema-rules--egress_rules--label_selector--expressions) |
| `rules.egress_rules.metadata` | [rules.egress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/metadata/#section) |
| `rules.egress_rules.metadata.description_spec` | [rules.egress_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/metadata/#schema-rules--egress_rules--metadata--description_spec) |
| `rules.egress_rules.metadata.name` | [rules.egress_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/metadata/#schema-rules--egress_rules--metadata--name) |
| `rules.egress_rules.outside_endpoints` | [rules.egress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/outside_endpoints/#section) |
| `rules.egress_rules.prefix_list` | [rules.egress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/prefix_list/#section) |
| `rules.egress_rules.prefix_list.prefixes` | [rules.egress_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/prefix_list/#schema-rules--egress_rules--prefix_list--prefixes) |
| `rules.egress_rules.protocol_port_range` | [rules.egress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/protocol_port_range/#section) |
| `rules.egress_rules.protocol_port_range.port_ranges` | [rules.egress_rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/protocol_port_range/#schema-rules--egress_rules--protocol_port_range--port_ranges) |
| `rules.egress_rules.protocol_port_range.protocol` | [rules.egress_rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/protocol_port_range/#schema-rules--egress_rules--protocol_port_range--protocol) |
| `rules.ingress_rules` | [rules.ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/#section) |
| `rules.ingress_rules.action` | [rules.ingress_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/#schema-rules--ingress_rules--action) |
| `rules.ingress_rules.adv_action` | [rules.ingress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/adv_action/#section) |
| `rules.ingress_rules.adv_action.action` | [rules.ingress_rules.adv_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/adv_action/#schema-rules--ingress_rules--adv_action--action) |
| `rules.ingress_rules.all_tcp_traffic` | [rules.ingress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_tcp_traffic/#section) |
| `rules.ingress_rules.all_traffic` | [rules.ingress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_traffic/#section) |
| `rules.ingress_rules.all_udp_traffic` | [rules.ingress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/all_udp_traffic/#section) |
| `rules.ingress_rules.any` | [rules.ingress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/any/#section) |
| `rules.ingress_rules.applications` | [rules.ingress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/applications/#section) |
| `rules.ingress_rules.applications.applications` | [rules.ingress_rules.applications.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/applications/#schema-rules--ingress_rules--applications--applications) |
| `rules.ingress_rules.inside_endpoints` | [rules.ingress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/inside_endpoints/#section) |
| `rules.ingress_rules.ip_prefix_set` | [rules.ingress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/#section) |
| `rules.ingress_rules.ip_prefix_set.ref` | [rules.ingress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#section) |
| `rules.ingress_rules.ip_prefix_set.ref.kind` | [rules.ingress_rules.ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#schema-rules--ingress_rules--ip_prefix_set--ref--kind) |
| `rules.ingress_rules.ip_prefix_set.ref.name` | [rules.ingress_rules.ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#schema-rules--ingress_rules--ip_prefix_set--ref--name) |
| `rules.ingress_rules.ip_prefix_set.ref.namespace` | [rules.ingress_rules.ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#schema-rules--ingress_rules--ip_prefix_set--ref--namespace) |
| `rules.ingress_rules.ip_prefix_set.ref.tenant` | [rules.ingress_rules.ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#schema-rules--ingress_rules--ip_prefix_set--ref--tenant) |
| `rules.ingress_rules.ip_prefix_set.ref.uid` | [rules.ingress_rules.ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/#schema-rules--ingress_rules--ip_prefix_set--ref--uid) |
| `rules.ingress_rules.label_matcher` | [rules.ingress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_matcher/#section) |
| `rules.ingress_rules.label_matcher.keys` | [rules.ingress_rules.label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_matcher/#schema-rules--ingress_rules--label_matcher--keys) |
| `rules.ingress_rules.label_selector` | [rules.ingress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_selector/#section) |
| `rules.ingress_rules.label_selector.expressions` | [rules.ingress_rules.label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/label_selector/#schema-rules--ingress_rules--label_selector--expressions) |
| `rules.ingress_rules.metadata` | [rules.ingress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/metadata/#section) |
| `rules.ingress_rules.metadata.description_spec` | [rules.ingress_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/metadata/#schema-rules--ingress_rules--metadata--description_spec) |
| `rules.ingress_rules.metadata.name` | [rules.ingress_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/metadata/#schema-rules--ingress_rules--metadata--name) |
| `rules.ingress_rules.outside_endpoints` | [rules.ingress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/outside_endpoints/#section) |
| `rules.ingress_rules.prefix_list` | [rules.ingress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/prefix_list/#section) |
| `rules.ingress_rules.prefix_list.prefixes` | [rules.ingress_rules.prefix_list.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/prefix_list/#schema-rules--ingress_rules--prefix_list--prefixes) |
| `rules.ingress_rules.protocol_port_range` | [rules.ingress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/protocol_port_range/#section) |
| `rules.ingress_rules.protocol_port_range.port_ranges` | [rules.ingress_rules.protocol_port_range.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/protocol_port_range/#schema-rules--ingress_rules--protocol_port_range--port_ranges) |
| `rules.ingress_rules.protocol_port_range.protocol` | [rules.ingress_rules.protocol_port_range.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/protocol_port_range/#schema-rules--ingress_rules--protocol_port_range--protocol) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/timeouts/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
