---
page_title: "vmware"
subcategory: ""
description: "vmware for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:cdd0f9ea38510a6dd250f0d681b910b89bda39d1e6b80efc6bb16f231b4d2255", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--vmware.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vmware"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/vmware/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vmware for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- vmware

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

VMware Provider Type. VMware Provider Type.

Upstream description:

VMware Provider Type.

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
vmware {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](resources--securemesh_site_v2--properties--vmware--not_managed.md): complete subsection reference.

## Next pages

- [vmware.not_managed](resources--securemesh_site_v2--properties--vmware--not_managed.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
