---
page_title: "rules.criteria.tcp"
subcategory: ""
description: "rules.criteria.tcp for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1348, "body_sha256": "sha256:83457c336e8d1ace8ef78442d02d51f5258c5cd26be2d6b5560179389222c1a1", "canonical_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:destination_port", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "path": "docs/guides/data-sources--nat_policy--properties--rules--criteria--tcp.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria", "tcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria.tcp for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
- [Property reference](data-sources--nat_policy--reference.md)
- [rules](data-sources--nat_policy--properties--rules.md)
- [rules.criteria](data-sources--nat_policy--properties--rules--criteria.md)
- rules.criteria.tcp

<a id="section"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

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

- [destination_port](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port.md): complete subsection reference.

- [source_port](data-sources--nat_policy--properties--rules--criteria--tcp--source_port.md): complete subsection reference.

## Next pages

- [rules.criteria.tcp.destination_port](data-sources--nat_policy--properties--rules--criteria--tcp--destination_port.md)
- [rules.criteria.tcp.source_port](data-sources--nat_policy--properties--rules--criteria--tcp--source_port.md)
- [rules.criteria](data-sources--nat_policy--properties--rules--criteria.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
