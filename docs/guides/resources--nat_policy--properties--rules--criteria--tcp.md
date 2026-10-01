---
page_title: "rules.criteria.tcp"
subcategory: ""
description: "rules.criteria.tcp for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1430, "body_sha256": "sha256:67d22873ed582c8cf50292571d4c418390ad268cbde10e1ced77731e59ebdc31", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "path": "docs/guides/resources--nat_policy--properties--rules--criteria--tcp.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria", "tcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria.tcp for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- [rules.criteria](resources--nat_policy--properties--rules--criteria.md)
- rules.criteria.tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [destination_port](resources--nat_policy--properties--rules--criteria--tcp--destination_port.md): complete subsection reference.

- [source_port](resources--nat_policy--properties--rules--criteria--tcp--source_port.md): complete subsection reference.

## Next pages

- [rules.criteria.tcp.destination_port](resources--nat_policy--properties--rules--criteria--tcp--destination_port.md)
- [rules.criteria.tcp.source_port](resources--nat_policy--properties--rules--criteria--tcp--source_port.md)
- [rules.criteria](resources--nat_policy--properties--rules--criteria.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
