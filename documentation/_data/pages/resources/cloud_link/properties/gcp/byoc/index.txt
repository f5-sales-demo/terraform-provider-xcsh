---
page_title: "gcp.byoc"
subcategory: ""
description: "List of GCP Bring You Own Connections."
xcsh_docs: {"aliases": ["gcp byoc"], "body_bytes": 1223, "body_sha256": "sha256:f480584ea4ecd5e339d9cf7cb25e3d60850b5438bb300a5499a053bb5098ac86", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:gcp", "path": "documentation/resources/cloud_link/properties/gcp/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "byoc"], "schema_version": 1, "sections": [{"aliases": ["gcp byoc connections"], "anchor": "section", "description": "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-gcp--byoc--connections--project", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "type": "conflicts"}, {"anchor": "schema-gcp--byoc--connections--interconnect_attachment_name", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}, {"anchor": "schema-gcp--byoc--connections--region", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "schema_path": ["gcp", "byoc", "connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/byoc/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of GCP Bring You Own Connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
