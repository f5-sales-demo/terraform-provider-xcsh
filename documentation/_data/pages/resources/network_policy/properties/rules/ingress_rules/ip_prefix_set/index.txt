---
page_title: "rules.ingress_rules.ip_prefix_set"
subcategory: "Security"
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["rules ingress rules ip prefix set"], "body_bytes": 1757, "body_sha256": "sha256:4dd12765200bbb218fd9e28b8d4a006acffa53694b1c5e6cddd0383e687697e8", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "path": "documentation/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0023333031132130-1100013001112130-0030201103122211-0022200330023310-1220102022312303-2310032133113000-1333332000200231-2113220132203232", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "ingress_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rules ingress rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "ingress_rules", "ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ingress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [rules.ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/)
- rules.ingress_rules.ip_prefix_set

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [rules.ingress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/ip_prefix_set/ref/)
- [rules.ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/ingress_rules/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
