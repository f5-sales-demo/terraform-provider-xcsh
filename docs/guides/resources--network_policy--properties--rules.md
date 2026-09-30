---
page_title: "rules"
subcategory: "Security"
description: "rules for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1133, "body_sha256": "sha256:bb490ff454626da43c683374202b471fe14f89a30a5671d867e8b611c71f0bc3", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:egress_rules", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "docs/guides/resources--network_policy--properties--rules.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
