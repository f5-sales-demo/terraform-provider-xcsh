---
page_title: "gcp.byoc"
subcategory: ""
description: "List of GCP Bring You Own Connections."
xcsh_docs: {"aliases": ["gcp byoc"], "body_bytes": 1367, "body_sha256": "sha256:814c56174c8b807731f8a7e569ce9d5e3d6da9f93a4bc3f94456d7727b2bb1cc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:gcp:byoc:connections"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "path": "documentation/data-sources/cloud_link/properties/gcp/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "byoc"], "schema_version": 1, "sections": [{"aliases": ["connections"], "anchor": "section", "description": "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:gcp:byoc:connections", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["gcp", "byoc", "connections"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/gcp/byoc/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of GCP Bring You Own Connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

List of GCP Bring You Own Connections.

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

## Next pages

- [gcp.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/byoc/connections/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/gcp/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
