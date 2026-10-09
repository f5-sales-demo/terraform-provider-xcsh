---
page_title: "azure_client_secret.client_secret"
subcategory: "Infrastructure"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["azure client secret client secret"], "body_bytes": 1410, "body_sha256": "sha256:23d71032cbffaf2a5892311b1ae719e45288cf2fb2d384db266b4247ff6acccc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "path": "documentation/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1332123332000132-1021131023201313-3223213033013310-1031303101032022-1033110303231302-2120022113321311-1100101231020320-2003123011131100", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_client_secret", "client_secret"], "schema_version": 1, "sections": [{"aliases": ["azure client secret client secret blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_client_secret", "client_secret", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure client secret client secret clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_client_secret", "client_secret", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_client_secret.client_secret

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/)
- azure_client_secret.client_secret

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/clear_secret_info/): complete subsection reference.
