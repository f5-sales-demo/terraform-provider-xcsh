---
page_title: "baremetal"
subcategory: ""
description: "Baremetal Provider Type."
xcsh_docs: {"aliases": ["baremetal"], "body_bytes": 1472, "body_sha256": "sha256:922520a1e71d5564f69b03ba85822427326f9bb1b741649ce81047ec7f493233", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/baremetal/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["baremetal"], "schema_version": 1, "sections": [{"aliases": ["not managed"], "anchor": "section", "description": "This section will show nodes associated with this site. Note: For sites that are not orchestrated by F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it will be shown in this section.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["baremetal", "not_managed"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/baremetal/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Baremetal Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- baremetal

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Baremetal Provider Type. Baremetal Provider Type.

Upstream description:

Baremetal Provider Type.

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
baremetal {
  # Configure direct properties listed below.
}
```

## Direct properties

- [not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/baremetal/not_managed/): complete subsection reference.

## Next pages

- [baremetal.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/baremetal/not_managed/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
