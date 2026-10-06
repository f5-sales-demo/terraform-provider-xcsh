---
page_title: "rules.egress_rules.ip_prefix_set"
subcategory: "Security"
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["rules egress rules ip prefix set"], "body_bytes": 1246, "body_sha256": "sha256:5f1e703a4e49eab654306e3d584afaff16005b7a98351fc8649dc9dd27fc7a2a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules", "path": "documentation/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0132103010030311-1033233100012302-0121033213310011-2312021232331203-1311112012232203-3311122002212323-1333112021330102-3222210311012011", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rules egress rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "egress_rules", "ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/): complete subsection reference.
