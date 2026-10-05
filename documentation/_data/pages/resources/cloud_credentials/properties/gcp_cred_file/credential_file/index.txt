---
page_title: "gcp_cred_file.credential_file"
subcategory: "Infrastructure"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "gcp cred file credential file"], "body_bytes": 2339, "body_sha256": "sha256:13efdf2f002c0128472a3583f599f543f9f024e67ab887c1f1ec7d343b848bb0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "path": "documentation/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_cred_file", "credential_file"], "schema_version": 1, "sections": [{"aliases": ["gcp cred file credential file blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_cred_file--credential_file--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "type": "requires"}], "schema_path": ["gcp_cred_file", "credential_file", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["gcp cred file credential file clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_cred_file--credential_file--clear_secret_info--url", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "type": "requires"}], "schema_path": ["gcp_cred_file", "credential_file", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file.credential_file

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/)
- gcp_cred_file.credential_file

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
credential_file {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/)
- [gcp_cred_file.credential_file.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
