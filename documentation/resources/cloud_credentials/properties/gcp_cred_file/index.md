---
page_title: "gcp_cred_file"
subcategory: "Infrastructure"
description: "GCP Credentials type."
xcsh_docs: {"aliases": ["gcp cred file"], "body_bytes": 1026, "body_sha256": "sha256:959aa7cdb6bf022397a5336f56602ff27a3f1cf63c1f573e2663f4c5ce922cdf", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/gcp_cred_file/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_cred_file"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "gcp cred file credential file"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "type": "conflicts"}], "schema_path": ["gcp_cred_file", "credential_file"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/gcp_cred_file/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "GCP Credentials type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- gcp_cred_file

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gcp cred file.

Additional upstream details:

GCP Credentials type.

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
gcp_cred_file {
  # Configure direct properties listed below.
}
```

## Direct properties

- [credential_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/): complete subsection reference.
