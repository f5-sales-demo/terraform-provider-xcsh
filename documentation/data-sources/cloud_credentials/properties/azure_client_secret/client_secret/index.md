---
page_title: "azure_client_secret.client_secret"
subcategory: "Infrastructure"
description: "azure_client_secret.client_secret for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 2011, "body_sha256": "sha256:896e94e85e9b8d3ab5bb25b152d3ca3f19cbfefc8f25695926cac5a51d54e288", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:blindfold_secret_info", "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret:client_secret", "parent_id": "xcsh-docs:data-sources:cloud_credentials:properties:azure_client_secret", "path": "documentation/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/index.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["azure_client_secret", "client_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_client_secret.client_secret for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [azure_client_secret.client_secret.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/)
- [azure_client_secret.client_secret.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/client_secret/clear_secret_info/)
- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
