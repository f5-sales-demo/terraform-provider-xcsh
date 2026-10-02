---
page_title: "rule_list"
subcategory: "Security"
description: "Ordered service-policy rules for non-geographic predicates and actions. Do not use country_list for a geo-only rule here: the platform adds match-all any_ip and any_asn selectors on readback, so the rule can match all traffic. Use deny_list or allow_list with country_list for geographic source matching."
xcsh_docs: {"aliases": ["rule list"], "body_bytes": 1890, "body_sha256": "sha256:575c05bf040cebd4a150132545aaab028dabc2a543747a8b636888f84e597edc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list", "parent_id": "xcsh-docs:resources:service_policy:reference", "path": "documentation/resources/service_policy/properties/rule_list/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0313033202111122-3020302003132312-3312200200230103-2323022300313111-1011031013312303-3010311302110303-3020012330131122-1010302101102100", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "sections": [{"aliases": ["rules"], "anchor": "section", "description": "Define the list of rules (with an order) that should be evaluated by this service policy. Rules are evaluated from top to bottom in the list.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered service-policy rules for non-geographic predicates and actions. Do not use country_list for a geo-only rule here: the platform adds match-all any_ip and any_asn selectors on readback, so the rule can match all traffic. Use deny_list or allow_list with country_list for geographic source matching.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/): complete subsection reference.

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
