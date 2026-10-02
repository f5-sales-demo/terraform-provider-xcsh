---
page_title: "aws.byoc"
subcategory: ""
description: "List of Bring You Own Connection."
xcsh_docs: {"aliases": ["aws byoc"], "body_bytes": 1353, "body_sha256": "sha256:312e33d03b7158045ae4b2017f5cd8c9a5be1bc70c030e7c66a1113f70ce6e6c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc", "parent_id": "xcsh-docs:data-sources:cloud_link:properties:aws", "path": "documentation/data-sources/cloud_link/properties/aws/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc"], "schema_version": 1, "sections": [{"aliases": ["connections"], "anchor": "section", "description": "List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but will be used for connecting sites and REs.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:aws:byoc:connections", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["aws", "byoc", "connections"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/aws/byoc/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Bring You Own Connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/)
- aws.byoc

<a id="section"></a>

Type: `"single"`. Computed.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

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

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/): complete subsection reference.

## Next pages

- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/byoc/connections/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/aws/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
