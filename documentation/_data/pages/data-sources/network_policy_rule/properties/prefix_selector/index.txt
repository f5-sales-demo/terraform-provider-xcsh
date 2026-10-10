---
page_title: "prefix_selector"
subcategory: ""
description: "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects."
xcsh_docs: {"aliases": ["prefix selector"], "body_bytes": 3067, "body_sha256": "sha256:5f5348534c7510326ea2baaacc73536fca55f5d3ff498bee24d1f3e16db0464f", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_rule:properties:prefix_selector", "parent_id": "xcsh-docs:data-sources:network_policy_rule:reference", "path": "documentation/data-sources/network_policy_rule/properties/prefix_selector/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3103133210210131-3111230202200002-1023002330233210-1022310220213223-0321323033330030-0203020201030121-1031130113032211-0303302132110012", "registry_path": "docs/guides/data-sources--network_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["prefix_selector"], "schema_version": 1, "sections": [{"aliases": ["prefix selector expressions"], "anchor": "schema-prefix_selector--expressions", "description": "Expressions contains the Kubernetes style label expression for selections.", "document_id": "xcsh-docs:data-sources:network_policy_rule:properties:prefix_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["prefix_selector", "expressions"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/properties/prefix_selector/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# prefix_selector

Breadcrumbs:

- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/)
- prefix_selector

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-prefix_selector--expressions"></a>

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
