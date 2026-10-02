---
page_title: "rule_list.rules.spec.ip_matcher"
subcategory: "Security"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["rule list rules spec ip matcher"], "body_bytes": 2381, "body_sha256": "sha256:32cac7e367f23ae8246c486f7c243fc9c70bcb3f79267dc7305556062b9449a3", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3001311313101200-0111313102213231-0330101313233022-0210001223230123-2232331303200003-2213132112112331-2233022233020222-3112013102100001", "registry_path": "docs/guides/data-sources--service_policy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["invert matcher"], "anchor": "schema-rule_list--rules--spec--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:ip_matcher:prefix_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "ip_matcher", "prefix_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.ip_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="schema-rule_list--rules--spec--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/ip_matcher/prefix_sets/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
