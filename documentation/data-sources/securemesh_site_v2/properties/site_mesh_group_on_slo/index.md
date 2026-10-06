---
page_title: "site_mesh_group_on_slo"
subcategory: ""
description: "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used."
xcsh_docs: {"aliases": ["site mesh group on slo"], "body_bytes": 1780, "body_sha256": "sha256:8778cd93ad6afe360e9bcb10430d2c0fc62e0ccac63a34bd1da2208b8103397d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_mesh_group_on_slo"], "schema_version": 1, "sections": [{"aliases": ["site mesh group on slo no site mesh group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "no_site_mesh_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo site mesh group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_mesh_group_on_slo", "site_mesh_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group on slo sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
