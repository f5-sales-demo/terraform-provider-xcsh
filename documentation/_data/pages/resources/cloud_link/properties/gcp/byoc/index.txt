---
page_title: "gcp.byoc"
subcategory: ""
description: "List of GCP Bring You Own Connections."
xcsh_docs: {"aliases": ["gcp byoc"], "body_bytes": 1609, "body_sha256": "sha256:3f9f46f9056e1f13d757d2399c4265c8aac5872232e2634704b99c98aedb79b3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:gcp", "path": "documentation/resources/cloud_link/properties/gcp/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "byoc"], "schema_version": 1, "sections": [{"aliases": ["connections"], "anchor": "section", "description": "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-gcp--byoc--connections--project", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "type": "conflicts"}, {"anchor": "schema-gcp--byoc--connections--interconnect_attachment_name", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}, {"anchor": "schema-gcp--byoc--connections--region", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "schema_path": ["gcp", "byoc", "connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/byoc/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of GCP Bring You Own Connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.byoc

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/)
- gcp.byoc

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Upstream description:

List of GCP Bring You Own Connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
```

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

Terraform syntax:

```terraform
byoc {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/): complete subsection reference.

## Next pages

- [gcp.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
