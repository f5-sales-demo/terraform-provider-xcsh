---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector"
subcategory: "Load Balancing"
description: "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects."
xcsh_docs: {"aliases": ["api rate limit bypass rate limiting rules bypass rate limiting rules client matcher client selector"], "body_bytes": 5230, "body_sha256": "sha256:5ed0c6f2f7fc62a628181c73abfd8ccae6b67a3a8de2e8b6cefa586198644518", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:client_selector", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/client_selector/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-007.md", "relationships": [{"anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--client_selector--expressions", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:client_selector", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "client_selector"], "schema_version": 1, "sections": [{"aliases": ["expressions"], "anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--client_selector--expressions", "description": "Expressions contains the Kubernetes style label expression for selections.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:client_selector", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "client_selector", "expressions"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/client_selector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

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

Terraform syntax:

```terraform
client_selector {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--client_selector--expressions"></a>

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

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
