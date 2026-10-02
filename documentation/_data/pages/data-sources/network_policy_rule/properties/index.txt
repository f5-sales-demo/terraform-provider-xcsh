---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_rule."
xcsh_docs: {"aliases": ["network policy rule"], "body_bytes": 14420, "body_sha256": "sha256:28882c3824ac4c7a187f067ee58b1949570b9b7c229ee104747008f88649df30", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy_rule:properties:advanced_action", "xcsh-docs:data-sources:network_policy_rule:properties:ip_prefix_set", "xcsh-docs:data-sources:network_policy_rule:properties:label_matcher", "xcsh-docs:data-sources:network_policy_rule:properties:prefix", "xcsh-docs:data-sources:network_policy_rule:properties:prefix_selector"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_rule:reference", "parent_id": "xcsh-docs:data-sources:network_policy_rule:fundamentals", "path": "documentation/data-sources/network_policy_rule/properties/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300", "registry_path": "docs/guides/data-sources--network_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-action", "description": "Network policy rule action configures the action to be taken on rule match Apply deny action on rule match Apply allow action on rule match.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action"], "syntax": "attribute", "type": "string"}, {"aliases": ["advanced action"], "anchor": "section", "description": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:advanced_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["label matcher"], "anchor": "section", "description": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:label_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["label_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["ports"], "anchor": "schema-ports", "description": "List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ports"], "syntax": "attribute", "type": "list"}, {"aliases": ["prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:prefix", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["prefix"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Labe", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:prefix_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["prefix_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol"], "anchor": "schema-protocol", "description": "Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.", "document_id": "xcsh-docs:data-sources:network_policy_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/)
- Property reference

## Direct properties

<a id="schema-action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/advanced_action/): complete subsection reference.

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

Description of the NetworkPolicyRule.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/label_matcher/): complete subsection reference.

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

Name of the NetworkPolicyRule.

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

Namespace where the NetworkPolicyRule exists.

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

<a id="schema-ports"></a>

### ports property

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix/): complete subsection reference.

- [prefix_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix_selector/): complete subsection reference.

<a id="schema-protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-action) |
| `advanced_action` | [advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/advanced_action/#section) |
| `advanced_action.action` | [advanced_action.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/advanced_action/#schema-advanced_action--action) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-id) |
| `ip_prefix_set` | [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/#section) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#section) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--kind) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--name) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--namespace) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--tenant) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--uid) |
| `label_matcher` | [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/label_matcher/#section) |
| `label_matcher.keys` | [label_matcher.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/label_matcher/#schema-label_matcher--keys) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-namespace) |
| `ports` | [ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-ports) |
| `prefix` | [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix/#section) |
| `prefix.prefix` | [prefix.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix/#schema-prefix--prefix) |
| `prefix_selector` | [prefix_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix_selector/#section) |
| `prefix_selector.expressions` | [prefix_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix_selector/#schema-prefix_selector--expressions) |
| `protocol` | [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/#schema-protocol) |

## Next pages

- [advanced_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/advanced_action/)
- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/ip_prefix_set/)
- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/label_matcher/)
- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix/)
- [prefix_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/prefix_selector/)
- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/)
