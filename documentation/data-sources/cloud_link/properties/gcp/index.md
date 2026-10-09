---
page_title: "gcp"
subcategory: ""
description: "CloudLink for GCP Cloud Provider."
xcsh_docs: {"aliases": ["gcp"], "body_bytes": 1024, "body_sha256": "sha256:a8ffd5e45f4e052afb4862a4b479ee4b33713907dae2b58e5ef062fe1decee01", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "xcsh-docs:data-sources:cloud_link:properties:gcp:gcp_cred"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "documentation/data-sources/cloud_link/properties/gcp/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp"], "schema_version": 1, "sections": [{"aliases": ["gcp byoc"], "anchor": "section", "description": "List of GCP Bring You Own Connections.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp", "byoc"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:gcp_cred", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp", "gcp_cred"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/gcp/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "CloudLink for GCP Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- gcp

<a id="section"></a>

Type: `"single"`. Computed.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]"
}
```

## Direct properties

- [byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/byoc/): complete subsection reference.

- [gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/gcp_cred/): complete subsection reference.
