---
page_title: "custom_geo_location_selector"
subcategory: ""
description: "custom_geo_location_selector for xcsh_geo_location_set."
xcsh_docs: {"aliases": [], "body_bytes": 4492, "body_sha256": "sha256:1bba4074aa1354024e3f23c6ea3f44dc42c7372337ef6a16d1079dcfd2163610", "child_ids": [], "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:geo_location_set:properties:custom_geo_location_selector", "parent_id": "xcsh-docs:resources:geo_location_set:reference", "path": "documentation/resources/geo_location_set/properties/custom_geo_location_selector/index.md", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["custom_geo_location_selector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/properties/custom_geo_location_selector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_geo_location_selector for xcsh_geo_location_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_geo_location_selector

Breadcrumbs:

- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/)
- custom_geo_location_selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_geo\_location\_selector, global\] Type can be used to establish a 'selector
reference' from one object(called selector) to a set of other objects(called selectees) based on the
value of expressions. A label selector is a label query over a set of resources. An empty label
selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

OneOf alternatives in this subsection:

- [custom_geo_location_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/custom_geo_location_selector/#section)
- [global](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/global/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_geo_location_selector {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_geo_location_selector--expressions"></a>

### expressions property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/)
- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/)
