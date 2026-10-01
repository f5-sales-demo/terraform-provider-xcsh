---
page_title: "admin_password"
subcategory: "Infrastructure"
description: "admin_password for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1652, "body_sha256": "sha256:525b4cfeac2bc8e556c5812ad8c53f0b1500b3598c16ecdf1a845d2668d2a03b", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:admin_password", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:admin_password:blindfold_secret_info", "xcsh-docs:resources:azure_vnet_site:properties:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:admin_password", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "docs/guides/resources--azure_vnet_site--properties--admin_password.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_password for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_password

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- admin_password

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
admin_password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--azure_vnet_site--properties--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [admin_password.blindfold_secret_info](resources--azure_vnet_site--properties--admin_password--blindfold_secret_info.md)
- [admin_password.clear_secret_info](resources--azure_vnet_site--properties--admin_password--clear_secret_info.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
