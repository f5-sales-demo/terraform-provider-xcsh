---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_data_type."
xcsh_docs: {"aliases": ["data type"], "body_bytes": 17774, "body_sha256": "sha256:14c8e11e415ce53ef5192d0885c38ea37e7e2eefe5987d7335ad5aa1523941b2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_type:properties:rules", "xcsh-docs:resources:data_type:properties:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:reference", "parent_id": "xcsh-docs:resources:data_type:fundamentals", "path": "documentation/resources/data_type/properties/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["compliances"], "anchor": "schema-compliances", "description": "Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance breach.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["compliances"], "syntax": "attribute", "type": "list"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["is pii"], "anchor": "schema-is_pii", "description": "Select this option to classify the custom data type as personally identifiable information (PII)", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["is_pii"], "syntax": "attribute", "type": "bool"}, {"aliases": ["is sensitive data"], "anchor": "schema-is_sensitive_data", "description": "Select this option to classify the custom data type as sensitive, enabling detection of API vulnerabilities related to this data type.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["is_sensitive_data"], "syntax": "attribute", "type": "bool"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:data_type:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules"], "anchor": "section", "description": "Configure key/value or regex match rules to enable the platform to detect this custom data type in the API request or response.", "document_id": "xcsh-docs:resources:data_type:properties:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:data_type:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_data_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
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

<a id="schema-compliances"></a>

### compliances property

Type: `["list", "string"]`. Optional.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach. Possible values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`,
\`APPI\`, \`HIPAA\`, \`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`,
\`ISO\_IEC\_27701\`, \`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Upstream description:

Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-is_pii"></a>

### is_pii property

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as personally identifiable information (PII).

Upstream description:

Select this option to classify the custom data type as personally identifiable information (PII)

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

<a id="schema-is_sensitive_data"></a>

### is_sensitive_data property

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as sensitive, enabling detection of API
vulnerabilities related to this data type.

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

Name of the Data Type. Must be unique within the namespace.

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

Namespace where the Data Type is created.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-annotations) |
| `compliances` | [compliances](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-compliances) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-id) |
| `is_pii` | [is_pii](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-is_pii) |
| `is_sensitive_data` | [is_sensitive_data](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-is_sensitive_data) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/#schema-namespace) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/#section) |
| `rules.key_pattern` | [rules.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/#section) |
| `rules.key_pattern.exact_values` | [rules.key_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/exact_values/#section) |
| `rules.key_pattern.exact_values.exact_values` | [rules.key_pattern.exact_values.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/exact_values/#schema-rules--key_pattern--exact_values--exact_values) |
| `rules.key_pattern.regex_value` | [rules.key_pattern.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/#schema-rules--key_pattern--regex_value) |
| `rules.key_pattern.substring_value` | [rules.key_pattern.substring_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_pattern/#schema-rules--key_pattern--substring_value) |
| `rules.key_value_pattern` | [rules.key_value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/#section) |
| `rules.key_value_pattern.key_pattern` | [rules.key_value_pattern.key_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/#section) |
| `rules.key_value_pattern.key_pattern.exact_values` | [rules.key_value_pattern.key_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/#section) |
| `rules.key_value_pattern.key_pattern.exact_values.exact_values` | [rules.key_value_pattern.key_pattern.exact_values.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/exact_values/#schema-rules--key_value_pattern--key_pattern--exact_values--exact_values) |
| `rules.key_value_pattern.key_pattern.regex_value` | [rules.key_value_pattern.key_pattern.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/#schema-rules--key_value_pattern--key_pattern--regex_value) |
| `rules.key_value_pattern.key_pattern.substring_value` | [rules.key_value_pattern.key_pattern.substring_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/key_pattern/#schema-rules--key_value_pattern--key_pattern--substring_value) |
| `rules.key_value_pattern.value_pattern` | [rules.key_value_pattern.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/#section) |
| `rules.key_value_pattern.value_pattern.exact_values` | [rules.key_value_pattern.value_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/#section) |
| `rules.key_value_pattern.value_pattern.exact_values.exact_values` | [rules.key_value_pattern.value_pattern.exact_values.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/exact_values/#schema-rules--key_value_pattern--value_pattern--exact_values--exact_values) |
| `rules.key_value_pattern.value_pattern.regex_value` | [rules.key_value_pattern.value_pattern.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/#schema-rules--key_value_pattern--value_pattern--regex_value) |
| `rules.key_value_pattern.value_pattern.substring_value` | [rules.key_value_pattern.value_pattern.substring_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/key_value_pattern/value_pattern/#schema-rules--key_value_pattern--value_pattern--substring_value) |
| `rules.value_pattern` | [rules.value_pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/#section) |
| `rules.value_pattern.exact_values` | [rules.value_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/exact_values/#section) |
| `rules.value_pattern.exact_values.exact_values` | [rules.value_pattern.exact_values.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/exact_values/#schema-rules--value_pattern--exact_values--exact_values) |
| `rules.value_pattern.regex_value` | [rules.value_pattern.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/#schema-rules--value_pattern--regex_value) |
| `rules.value_pattern.substring_value` | [rules.value_pattern.substring_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/#schema-rules--value_pattern--substring_value) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/timeouts/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
