---
page_title: "azure_pfx_certificate.password"
subcategory: "Infrastructure"
description: "azure_pfx_certificate.password for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1919, "body_sha256": "sha256:dd62828f13da01f1f84f63f5a83d2f192aa308b120e8ef6f2a47616b89152edc", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate:password", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "path": "docs/guides/resources--cloud_credentials--properties--azure_pfx_certificate--password.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_pfx_certificate", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/azure_pfx_certificate/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_pfx_certificate.password for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_pfx_certificate.password

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Property reference](resources--cloud_credentials--reference.md)
- [azure_pfx_certificate](resources--cloud_credentials--properties--azure_pfx_certificate.md)
- azure_pfx_certificate.password

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
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [azure_pfx_certificate.password.blindfold_secret_info](resources--cloud_credentials--properties--azure_pfx_certificate--password--blindfold_secret_info.md)
- [azure_pfx_certificate.password.clear_secret_info](resources--cloud_credentials--properties--azure_pfx_certificate--password--clear_secret_info.md)
- [azure_pfx_certificate](resources--cloud_credentials--properties--azure_pfx_certificate.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
