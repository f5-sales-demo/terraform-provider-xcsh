---
page_title: "nutanix.not_managed"
subcategory: ""
description: "nutanix.not_managed for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1219, "body_sha256": "sha256:2b8de66575678acf6384c7ae3f7ad5cb5d6d9a72c7e65d926aea0b9b38ab3adb", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix", "path": "docs/guides/data-sources--securemesh_site_v2--properties--nutanix--not_managed.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix", "not_managed"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix.not_managed for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# nutanix.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [nutanix](data-sources--securemesh_site_v2--properties--nutanix.md)
- nutanix.not_managed

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [node_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md): complete subsection reference.

## Next pages

- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md)
- [nutanix](data-sources--securemesh_site_v2--properties--nutanix.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
