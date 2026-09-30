---
page_title: "azure.not_managed"
subcategory: ""
description: "azure.not_managed for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:baa16a267a40476a210727882f8840045af5c2d1081406d52febc2e0e0ba5d6c", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure", "path": "docs/guides/resources--securemesh_site_v2--properties--azure--not_managed.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure", "not_managed"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/azure/not_managed/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure.not_managed for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [azure](resources--securemesh_site_v2--properties--azure.md)
- azure.not_managed

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

## Direct properties

- [node_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list.md): complete subsection reference.

## Next pages

- [azure.not_managed.node_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list.md)
- [azure](resources--securemesh_site_v2--properties--azure.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
