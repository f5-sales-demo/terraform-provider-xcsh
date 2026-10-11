---
page_title: "gcp"
subcategory: ""
description: "CloudLink for GCP Cloud Provider."
xcsh_docs: {"aliases": ["gcp"], "body_bytes": 1124, "body_sha256": "sha256:69f67c0c2c4192314d541cd0fe53fc4d3dc1bc9363bdadce60be54cd7207e298", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:gcp:byoc", "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp", "parent_id": "xcsh-docs:resources:cloud_link:reference", "path": "documentation/resources/cloud_link/properties/gcp/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp"], "schema_version": 1, "sections": [{"aliases": ["gcp byoc"], "anchor": "section", "description": "List of GCP Bring You Own Connections.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "schema_path": ["gcp", "byoc"], "syntax": "block", "type": "object"}, {"aliases": ["gcp gcp cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp--gcp_cred--name", "enforcement": "provider-schema", "group": "gcp.gcp_cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:gcp_cred", "type": "requires"}], "schema_path": ["gcp", "gcp_cred"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "CloudLink for GCP Cloud Provider.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
