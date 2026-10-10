---
page_title: "custom_geo_location_selector"
subcategory: ""
description: "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects."
xcsh_docs: {"aliases": ["custom geo location selector"], "body_bytes": 3744, "body_sha256": "sha256:2307197385cd30616c34f236037075bde70966ab37c0a99c5515229d62952d5e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:geo_location_set:properties:custom_geo_location_selector", "parent_id": "xcsh-docs:data-sources:geo_location_set:reference", "path": "documentation/data-sources/geo_location_set/properties/custom_geo_location_selector/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2020332100211112-3033212010031001-3133022212203121-1203000333110321-3032301101133301-1302321210330322-0211210121122321-0302331300332311", "registry_path": "docs/guides/data-sources--geo_location_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_geo_location_selector"], "schema_version": 1, "sections": [{"aliases": ["custom geo location selector expressions"], "anchor": "schema-custom_geo_location_selector--expressions", "description": "Expressions contains the Kubernetes style label expression for selections.", "document_id": "xcsh-docs:data-sources:geo_location_set:properties:custom_geo_location_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_geo_location_selector", "expressions"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/geo_location_set/properties/custom_geo_location_selector/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_geo_location_selector

Breadcrumbs:

- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/properties/)
- custom_geo_location_selector

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_geo\_location\_selector, global\] Type can be used to establish a 'selector
reference' from one object(called selector) to a set of other objects(called selectees) based on the
value of expressions. A label selector is a label query over a set of resources. An empty label
selector matches all objects.

Additional upstream details:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A null label selector matches
no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

- [custom_geo_location_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/properties/custom_geo_location_selector/#section)
- [global](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/properties/global/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-custom_geo_location_selector--expressions"></a>

### expressions property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
