---
page_title: "nutanix.not_managed"
subcategory: ""
description: "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section."
xcsh_docs: {"aliases": ["nutanix not managed"], "body_bytes": 1673, "body_sha256": "sha256:fabf26099ad293b6fee89b626c7f969facf7abc65c17e1fd0d9050372813a996", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix", "path": "documentation/data-sources/securemesh_site_v2/properties/nutanix/not_managed/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nutanix", "not_managed"], "schema_version": 1, "sections": [{"aliases": ["node list"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/)
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

- [node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/): complete subsection reference.

## Next pages

- [nutanix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
