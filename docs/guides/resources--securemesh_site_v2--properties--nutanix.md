---
page_title: "nutanix"
subcategory: ""
description: "nutanix for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1047, "body_sha256": "sha256:e721b4458f15d23b7cea9c1e18dab57be3399059b3935344b4a6012e60ba123e", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--nutanix.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/nutanix/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# nutanix

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- nutanix

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Nutanix Provider Type. Nutanix Provider Type.

Upstream description:

Nutanix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
nutanix {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](resources--securemesh_site_v2--properties--nutanix--not_managed.md): complete subsection reference.

## Next pages

- [nutanix.not_managed](resources--securemesh_site_v2--properties--nutanix--not_managed.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
