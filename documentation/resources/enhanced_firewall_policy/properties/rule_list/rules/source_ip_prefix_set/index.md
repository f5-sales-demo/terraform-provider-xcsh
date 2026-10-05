---
page_title: "rule_list.rules.source_ip_prefix_set"
subcategory: ""
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["rule list rules source ip prefix set"], "body_bytes": 1871, "body_sha256": "sha256:7fb309ca8796fc5e1f4fbb9de0b21af5f0854777ca1208355b47ac94d048e23b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "source_ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rule list rules source ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:source_ip_prefix_set:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "source_ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.source_ip_prefix_set

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.source_ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
source_ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [rule_list.rules.source_ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/source_ip_prefix_set/ref/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
