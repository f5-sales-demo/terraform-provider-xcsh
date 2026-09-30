---
page_title: "rules.criteria.icmp"
subcategory: ""
description: "rules.criteria.icmp for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 909, "body_sha256": "sha256:8b26f6f9a77c0127b1275546f7b92335c7830755ae4315b9087e6e417e94b4be", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "child_ids": [], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "path": "docs/guides/resources--nat_policy--properties--rules--criteria--icmp.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "criteria", "icmp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/icmp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.criteria.icmp for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.criteria.icmp

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- [rules.criteria](resources--nat_policy--properties--rules--criteria.md)
- rules.criteria.icmp

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
icmp = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.criteria](resources--nat_policy--properties--rules--criteria.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
