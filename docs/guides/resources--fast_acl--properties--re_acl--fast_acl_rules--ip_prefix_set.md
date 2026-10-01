---
page_title: "re_acl.fast_acl_rules.ip_prefix_set"
subcategory: ""
description: "re_acl.fast_acl_rules.ip_prefix_set for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1315, "body_sha256": "sha256:dbfeb11c780afc57126db866c065ab51d4e4cbd50e622367dbb8a7a22dc237e0", "canonical_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "path": "docs/guides/resources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl.fast_acl_rules.ip_prefix_set for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [re_acl](resources--fast_acl--properties--re_acl.md)
- [re_acl.fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md)
- re_acl.fast_acl_rules.ip_prefix_set

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

- [ref](resources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [re_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set--ref.md)
- [re_acl.fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
