---
page_title: "rules.key_value_pattern.key_pattern"
subcategory: ""
description: "rules.key_value_pattern.key_pattern for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 4176, "body_sha256": "sha256:4fddd9a4a893904eb3addc2a77efe69070bdeca850693228975afc6ece5669e0", "canonical_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "child_ids": ["xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern:exact_values"], "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern:key_pattern", "parent_id": "xcsh-docs:resources:data_type:properties:rules:key_value_pattern", "path": "docs/guides/resources--data_type--properties--rules--key_value_pattern--key_pattern.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "key_value_pattern", "key_pattern"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/key_value_pattern/key_pattern/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.key_value_pattern.key_pattern for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.key_value_pattern.key_pattern

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md)
- [Property reference](resources--data_type--reference.md)
- [rules](resources--data_type--properties--rules.md)
- [rules.key_value_pattern](resources--data_type--properties--rules--key_value_pattern.md)
- rules.key_value_pattern.key_pattern

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Upstream description:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
key_pattern {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exact_values](resources--data_type--properties--rules--key_value_pattern--key_pattern--exact_values.md): complete subsection reference.

<a id="schema-rules--key_value_pattern--key_pattern--regex_value"></a>

### regex_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-rules--key_value_pattern--key_pattern--substring_value"></a>

### substring_value property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regex\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

## Next pages

- [rules.key_value_pattern.key_pattern.exact_values](resources--data_type--properties--rules--key_value_pattern--key_pattern--exact_values.md)
- [rules.key_value_pattern](resources--data_type--properties--rules--key_value_pattern.md)
- [xcsh_data_type](../resources/data_type.md)
