---
page_title: "site_mesh_group_on_slo"
subcategory: ""
description: "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used."
xcsh_docs: {"aliases": ["site mesh group on slo"], "body_bytes": 2802, "body_sha256": "sha256:9dd6171aa52dbc0f6e13b58398042cb53c76f5e27ed321d5a5594cb8a08bd8da", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_mesh_group_on_slo"], "schema_version": 1, "sections": [{"aliases": ["site mesh group on slo no site mesh group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "no_site_mesh_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo site mesh group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_mesh_group_on_slo", "site_mesh_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_mesh_group_on_slo

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- site_mesh_group_on_slo

<a id="section"></a>

Type: `"single"`. Computed.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

## Direct properties

- [no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/): complete subsection reference.

- [site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [site_mesh_group_on_slo.no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/)
- [site_mesh_group_on_slo.site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/)
- [site_mesh_group_on_slo.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
