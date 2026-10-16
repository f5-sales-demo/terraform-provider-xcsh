---
page_title: "bigip_virtual_server"
subcategory: ""
description: "Set the scope of the API Group to a specific BIG-IP Virtual Server."
xcsh_docs: {"aliases": ["bigip virtual server"], "body_bytes": 1515, "body_sha256": "sha256:5a3db43588ab75561ade0aad9bac35367842a03d2e5afcad0ba48e543edb8abf", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "documentation/data-sources/app_api_group/properties/bigip_virtual_server/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1213011112130221-2112210023332111-0122001122030111-3000113002222221-0222313333311230-0112200212111121-3203220023033102-3020320221222303", "registry_path": "docs/guides/data-sources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bigip_virtual_server"], "schema_version": 1, "sections": [{"aliases": ["bigip virtual server bigip virtual server"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server:bigip_virtual_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bigip_virtual_server", "bigip_virtual_server"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Set the scope of the API Group to a specific BIG-IP Virtual Server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bigip_virtual_server

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- bigip_virtual_server

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

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

OneOf alternatives in this subsection:

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/bigip_virtual_server/#section)
- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/cdn_loadbalancer/#section)
- [http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/http_loadbalancer/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/bigip_virtual_server/bigip_virtual_server/): complete subsection reference.
