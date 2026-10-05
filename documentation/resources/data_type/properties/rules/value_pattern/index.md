---
page_title: "rules.value_pattern"
subcategory: ""
description: "Test"
xcsh_docs: {"aliases": ["rules value pattern"], "body_bytes": 4291, "body_sha256": "sha256:47ae2ec43f84120b6ed61fccd7fdfe9f3529f3c204de3809c216b325694e7046", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "parent_id": "xcsh-docs:resources:data_type:properties:rules", "path": "documentation/resources/data_type/properties/rules/value_pattern/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0210132302033200-2013032022113300-3003000101013321-2012001022212223-1032031301030003-2333213213102000-3013112330122003-0033012210122212", "registry_path": "docs/guides/resources--data_type--reference--group-001.md", "relationships": [{"anchor": "schema-rules--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--regex_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "schema-rules--value_pattern--substring_value", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:regex_value,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.value_pattern:ConflictingObjectAttributes:exact_values,substring_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "value_pattern"], "schema_version": 1, "sections": [{"aliases": ["rules value pattern exact values"], "anchor": "section", "description": "List of exact values to match.", "document_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--value_pattern--exact_values--exact_values", "enforcement": "provider-schema", "group": "rules.value_pattern.exact_values:RequiredObjectAttributes:exact_values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern:exact_values", "type": "requires"}], "schema_path": ["rules", "value_pattern", "exact_values"], "syntax": "block", "type": "object"}, {"aliases": ["rules value pattern regex value"], "anchor": "schema-rules--value_pattern--regex_value", "description": "Exclusive with Search for values matching this regular expression.", "document_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "value_pattern", "regex_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules value pattern substring value"], "anchor": "schema-rules--value_pattern--substring_value", "description": "Exclusive with Search for values that include this substring.", "document_id": "xcsh-docs:resources:data_type:properties:rules:value_pattern", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "value_pattern", "substring_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/properties/rules/value_pattern/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Test", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.value_pattern

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- rules.value_pattern

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

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
value_pattern {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/exact_values/): complete subsection reference.

<a id="schema-rules--value_pattern--regex_value"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-rules--value_pattern--substring_value"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules.value_pattern.exact_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/value_pattern/exact_values/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/rules/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/)
