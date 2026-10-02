---
page_title: "gcp"
subcategory: ""
description: "CloudLink for GCP Cloud Provider."
xcsh_docs: {"aliases": ["gcp"], "body_bytes": 1554, "body_sha256": "sha256:981884093058c8fd7688f85cd7e9b7a7e049c1841965112744242ffe78aea4d5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "xcsh-docs:data-sources:cloud_link:properties:gcp:gcp_cred"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "documentation/data-sources/cloud_link/properties/gcp/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp"], "schema_version": 1, "sections": [{"aliases": ["byoc"], "anchor": "section", "description": "List of GCP Bring You Own Connections.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp", "byoc"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:gcp_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp", "gcp_cred"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/gcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "CloudLink for GCP Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

CloudLink for GCP Cloud Provider.

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

## Next pages

- [gcp.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/byoc/)
- [gcp.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/gcp_cred/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
