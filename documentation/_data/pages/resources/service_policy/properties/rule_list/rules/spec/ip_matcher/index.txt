---
page_title: "rule_list.rules.spec.ip_matcher"
subcategory: "Security"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["rule list rules spec ip matcher"], "body_bytes": 2623, "body_sha256": "sha256:2c079373cd8e42a40e0f785cdb15dffaaaa15109ee0ed48a0dee651d70ff966c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1030111002320303-0232312311303320-3110022123122122-1021111312200023-0113312211223231-0101120320103223-1232331122313210-1312313222312210", "registry_path": "docs/guides/resources--service_policy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec ip matcher invert matcher"], "anchor": "schema-rule_list--rules--spec--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules spec ip matcher prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_matcher", "prefix_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
