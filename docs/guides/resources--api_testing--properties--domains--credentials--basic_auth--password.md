---
page_title: "domains.credentials.basic_auth.password"
subcategory: ""
description: "domains.credentials.basic_auth.password for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 2118, "body_sha256": "sha256:5c6b9747e4fcb59febe8aa86a1329a63c03e82760345bae054724e0928781b5f", "canonical_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth:password", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth:password:blindfold_secret_info", "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth:password:clear_secret_info"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth:password", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:basic_auth", "path": "docs/guides/resources--api_testing--properties--domains--credentials--basic_auth--password.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "basic_auth", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/basic_auth/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.basic_auth.password for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.basic_auth.password

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
- [Property reference](resources--api_testing--reference.md)
- [domains](resources--api_testing--properties--domains.md)
- [domains.credentials](resources--api_testing--properties--domains--credentials.md)
- [domains.credentials.basic_auth](resources--api_testing--properties--domains--credentials--basic_auth.md)
- domains.credentials.basic_auth.password

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

- [blindfold_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.credentials.basic_auth.password.blindfold_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md)
- [domains.credentials.basic_auth.password.clear_secret_info](resources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md)
- [domains.credentials.basic_auth](resources--api_testing--properties--domains--credentials--basic_auth.md)
- [xcsh_api_testing](../resources/api_testing.md)
