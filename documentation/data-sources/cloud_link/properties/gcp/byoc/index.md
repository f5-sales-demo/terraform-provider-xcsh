---
page_title: "gcp.byoc"
subcategory: ""
description: "List of GCP Bring You Own Connections."
xcsh_docs: {"aliases": ["gcp byoc"], "body_bytes": 942, "body_sha256": "sha256:d4f7986d64d5493d0878ebe739e56613387c7e694be94487d253752c71703962", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:gcp:byoc:connections"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "path": "documentation/data-sources/cloud_link/properties/gcp/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "byoc"], "schema_version": 1, "sections": [{"aliases": ["gcp byoc connections"], "anchor": "section", "description": "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc:connections", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gcp", "byoc", "connections"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/gcp/byoc/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "List of GCP Bring You Own Connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.byoc

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/)
- gcp.byoc

<a id="section"></a>

Type: `"single"`. Computed.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

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

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/byoc/connections/): complete subsection reference.
