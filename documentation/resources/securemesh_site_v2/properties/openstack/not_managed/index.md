---
page_title: "openstack.not_managed"
subcategory: ""
description: "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section."
xcsh_docs: {"aliases": ["openstack not managed"], "body_bytes": 1790, "body_sha256": "sha256:52a37d03365f36b5adacc4001b80162fe4df959ec86306465f01dc0ce7cbc232", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack", "path": "documentation/resources/securemesh_site_v2/properties/openstack/not_managed/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0312033221102032-1100121330120132-2103020332012010-3111210221102130-0320102323211200-3311310000001322-0301130130001001-2213113021333211", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed"], "schema_version": 1, "sections": [{"aliases": ["openstack not managed node list"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["openstack", "not_managed", "node_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openstack/not_managed/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/)
- openstack.not_managed

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

- [node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
