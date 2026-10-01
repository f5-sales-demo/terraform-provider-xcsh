---
page_title: "rules.ingress_rules.all_traffic"
subcategory: "Security"
description: "rules.ingress_rules.all_traffic for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1113, "body_sha256": "sha256:611d01d6df7d0e7c135e2fb7557ce5d85fa0c383fad071ea10be5a1b1de541c9", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "path": "docs/guides/resources--network_policy--properties--rules--ingress_rules--all_traffic.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ingress_rules", "all_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/all_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ingress_rules.all_traffic for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ingress_rules.all_traffic

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [rules](resources--network_policy--properties--rules.md)
- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- rules.ingress_rules.all_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

Upstream description:

This can be used for messages where no values are needed.

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
all_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- [xcsh_network_policy](../resources/network_policy.md)
