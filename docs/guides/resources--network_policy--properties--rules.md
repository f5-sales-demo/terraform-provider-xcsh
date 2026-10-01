---
page_title: "rules"
subcategory: "Security"
description: "rules for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1232, "body_sha256": "sha256:7e796cbd1e535b984bbfbabd24526d45d0612fb1841417515055009878848f5d", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:egress_rules", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "docs/guides/resources--network_policy--properties--rules.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Rule Choice. Shape of Rule Choice.

Upstream description:

Shape of Rule Choice.

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
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [egress_rules](resources--network_policy--properties--rules--egress_rules.md): complete subsection reference.

- [ingress_rules](resources--network_policy--properties--rules--ingress_rules.md): complete subsection reference.

## Next pages

- [rules.egress_rules](resources--network_policy--properties--rules--egress_rules.md)
- [rules.ingress_rules](resources--network_policy--properties--rules--ingress_rules.md)
- [Property reference](resources--network_policy--reference.md)
- [xcsh_network_policy](../resources/network_policy.md)
