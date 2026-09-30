---
page_title: "site_acl.fast_acl_rules.ip_prefix_set"
subcategory: ""
description: "site_acl.fast_acl_rules.ip_prefix_set for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1238, "body_sha256": "sha256:3a8059a5e7e8c8ab1742d9d345136c011cc63d1b51be1161a7310db2a4792edb", "canonical_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "path": "docs/guides/resources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.fast_acl_rules.ip_prefix_set for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_acl.fast_acl_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [site_acl](resources--fast_acl--properties--site_acl.md)
- [site_acl.fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md)
- site_acl.fast_acl_rules.ip_prefix_set

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

- [ref](resources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set--ref.md)
- [site_acl.fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
