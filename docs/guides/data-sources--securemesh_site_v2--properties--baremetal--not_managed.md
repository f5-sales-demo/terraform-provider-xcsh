---
page_title: "baremetal.not_managed"
subcategory: ""
description: "baremetal.not_managed for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1336, "body_sha256": "sha256:f323538cba47c0f1a56eb2b2f7d3f4d8ece8511eb24901f9380e3ae257ec3578", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal", "path": "docs/guides/data-sources--securemesh_site_v2--properties--baremetal--not_managed.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["baremetal", "not_managed"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/baremetal/not_managed/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "baremetal.not_managed for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [baremetal](data-sources--securemesh_site_v2--properties--baremetal.md)
- baremetal.not_managed

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

- [node_list](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list.md): complete subsection reference.

## Next pages

- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list.md)
- [baremetal](data-sources--securemesh_site_v2--properties--baremetal.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
