---
page_title: "rule_list"
subcategory: "Security"
description: "Ordered service-policy rules for non-geographic predicates and actions. Do not use country_list for a geo-only rule here: the platform adds match-all any_ip and any_asn selectors on readback, so the rule can match all traffic. Use deny_list or allow_list with country_list for geographic source matching."
xcsh_docs: {"aliases": ["rule list"], "body_bytes": 1790, "body_sha256": "sha256:fe0b47bd27e8846fd11d0e7b1911a71c6c2f1eae53627a3ee184668c22eac462", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "documentation/data-sources/service_policy/properties/rule_list/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033", "registry_path": "docs/guides/data-sources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "sections": [{"aliases": ["rule list rules"], "anchor": "section", "description": "Define the list of rules (with an order) that should be evaluated by this service policy. Rules are evaluated from top to bottom in the list.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Ordered service-policy rules for non-geographic predicates and actions. Do not use country_list for a geo-only rule here: the platform adds match-all any_ip and any_asn selectors on readback, so the rule can match all traffic. Use deny_list or allow_list with country_list for geographic source matching.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source..

Upstream description:

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source
matching.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/): complete subsection reference.

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
