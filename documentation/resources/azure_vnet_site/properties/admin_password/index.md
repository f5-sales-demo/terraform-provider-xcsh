---
page_title: "admin_password"
subcategory: "Infrastructure"
description: "admin_password for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1961, "body_sha256": "sha256:5d6763e3023ccc80df64fcde80700c54004f53d4173a97fbe573159396b6849c", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:admin_password:blindfold_secret_info", "xcsh-docs:resources:azure_vnet_site:properties:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:admin_password", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/admin_password/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_password for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# admin_password

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/admin_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/admin_password/clear_secret_info/): complete subsection reference.

## Next pages

- [admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/admin_password/blindfold_secret_info/)
- [admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/admin_password/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
