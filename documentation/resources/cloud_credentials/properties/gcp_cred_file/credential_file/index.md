---
page_title: "gcp_cred_file.credential_file"
subcategory: "Infrastructure"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "gcp cred file credential file"], "body_bytes": 2339, "body_sha256": "sha256:13efdf2f002c0128472a3583f599f543f9f024e67ab887c1f1ec7d343b848bb0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "path": "documentation/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_cred_file", "credential_file"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_cred_file--credential_file--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:blindfold_secret_info", "type": "requires"}], "schema_path": ["gcp_cred_file", "credential_file", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp_cred_file--credential_file--clear_secret_info--url", "enforcement": "provider-schema", "group": "gcp_cred_file.credential_file.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file:credential_file:clear_secret_info", "type": "requires"}], "schema_path": ["gcp_cred_file", "credential_file", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/gcp_cred_file/credential_file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
