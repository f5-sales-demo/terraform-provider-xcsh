---
page_title: "password"
subcategory: ""
description: "password for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1857, "body_sha256": "sha256:8b121fcad3b66443955051bb6cf52e0047d85782f8f67e3aaf0b6d53d79d0f38", "child_ids": ["xcsh-docs:resources:cminstance:properties:password:blindfold_secret_info", "xcsh-docs:resources:cminstance:properties:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:properties:password", "parent_id": "xcsh-docs:resources:cminstance:reference", "path": "documentation/resources/cminstance/properties/password/index.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/properties/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "password for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# password

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/blindfold_secret_info/)
- [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/password/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
