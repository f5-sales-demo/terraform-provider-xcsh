---
page_title: "f5_big_ip_aws_service.admin_password"
subcategory: ""
description: "f5_big_ip_aws_service.admin_password for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1907, "body_sha256": "sha256:09ccaaeba2a789886dc40b7a8302d158a7274e5bcdac14e327268195a50ae989", "canonical_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password:blindfold_secret_info", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "docs/guides/resources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "admin_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/admin_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.admin_password for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.admin_password

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- f5_big_ip_aws_service.admin_password

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

- [blindfold_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--clear_secret_info.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.admin_password.blindfold_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md)
- [f5_big_ip_aws_service.admin_password.clear_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--clear_secret_info.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
