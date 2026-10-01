---
page_title: "k8s_cluster_role_selector"
subcategory: "Container"
description: "k8s_cluster_role_selector for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 3870, "body_sha256": "sha256:c3b6fb3f556c1f7fd7b44dbb945d41e046a4f0716aa035ba398b92394f45aa14", "canonical_id": "xcsh-docs:data-sources:k8s_cluster_role:properties:k8s_cluster_role_selector", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role:properties:k8s_cluster_role_selector", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:reference", "path": "docs/guides/data-sources--k8s_cluster_role--properties--k8s_cluster_role_selector.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["k8s_cluster_role_selector"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/properties/k8s_cluster_role_selector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "k8s_cluster_role_selector for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# k8s_cluster_role_selector

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
- [Property reference](data-sources--k8s_cluster_role--reference.md)
- k8s_cluster_role_selector

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: k8s\_cluster\_role\_selector, policy\_rule\_list, yaml\] Type can be used to establish a
'selector reference' from one object(called selector) to a set of other objects(called selectees)
based on the value of expressions. A label selector is a label query over a set of resources. An
empty label selector matches all objects.

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

- [k8s_cluster_role_selector](data-sources--k8s_cluster_role--properties--k8s_cluster_role_selector.md#section)
- [policy_rule_list](data-sources--k8s_cluster_role--properties--policy_rule_list.md#section)
- [yaml](data-sources--k8s_cluster_role--reference.md#schema-yaml)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-k8s_cluster_role_selector--expressions"></a>

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

- [Property reference](data-sources--k8s_cluster_role--reference.md)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
