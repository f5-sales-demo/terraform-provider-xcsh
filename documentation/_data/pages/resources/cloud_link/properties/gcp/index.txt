---
page_title: "gcp"
subcategory: ""
description: "CloudLink for GCP Cloud Provider."
xcsh_docs: {"aliases": ["gcp"], "body_bytes": 1642, "body_sha256": "sha256:ea3124abdab328ab04e08aaa4bdac2053cfc3819c0d61a9f9b6029da7f712e80", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:gcp:byoc", "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp", "parent_id": "xcsh-docs:resources:cloud_link:reference", "path": "documentation/resources/cloud_link/properties/gcp/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp"], "schema_version": 1, "sections": [{"aliases": ["byoc"], "anchor": "section", "description": "List of GCP Bring You Own Connections.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "schema_path": ["gcp", "byoc"], "syntax": "block", "type": "object"}, {"aliases": ["gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp--gcp_cred--name", "enforcement": "provider-schema", "group": "gcp.gcp_cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred", "type": "requires"}], "schema_path": ["gcp", "gcp_cred"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CloudLink for GCP Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- gcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
gcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/): complete subsection reference.

- [gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/): complete subsection reference.

## Next pages

- [gcp.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/)
- [gcp.gcp_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/gcp_cred/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
