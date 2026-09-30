---
page_title: "rules.action.dynamic"
subcategory: ""
description: "rules.action.dynamic for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1567, "body_sha256": "sha256:9162971cdc9a9fecbd541b5f6c2d7ebcccdceee42f4c0abb7284b50c0c8df534", "canonical_id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:elastic_ips", "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic:pools"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:action:dynamic", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:action", "path": "docs/guides/resources--nat_policy--properties--rules--action--dynamic.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "action", "dynamic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/action/dynamic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action.dynamic for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.action.dynamic

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md)
- [Property reference](resources--nat_policy--reference.md)
- [rules](resources--nat_policy--properties--rules.md)
- [rules.action](resources--nat_policy--properties--rules--action.md)
- rules.action.dynamic

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Upstream description:

Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("elastic_ips",
    "pools")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

Terraform syntax:

```terraform
dynamic {
  # Configure direct properties listed below.
}
```

## Direct properties

- [elastic_ips](resources--nat_policy--properties--rules--action--dynamic--elastic_ips.md): complete subsection reference.

- [pools](resources--nat_policy--properties--rules--action--dynamic--pools.md): complete subsection reference.

## Next pages

- [rules.action.dynamic.elastic_ips](resources--nat_policy--properties--rules--action--dynamic--elastic_ips.md)
- [rules.action.dynamic.pools](resources--nat_policy--properties--rules--action--dynamic--pools.md)
- [rules.action](resources--nat_policy--properties--rules--action.md)
- [xcsh_nat_policy](../resources/nat_policy.md)
