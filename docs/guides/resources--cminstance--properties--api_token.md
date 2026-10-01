---
page_title: "api_token"
subcategory: ""
description: "api_token for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1557, "body_sha256": "sha256:9fa51ebb6c56c26a6fde2b0052bb2449b35eadb8f2a12b8e52ade0f7e17197e5", "canonical_id": "xcsh-docs:resources:cminstance:properties:api_token", "child_ids": ["xcsh-docs:resources:cminstance:properties:api_token:blindfold_secret_info", "xcsh-docs:resources:cminstance:properties:api_token:clear_secret_info"], "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:properties:api_token", "parent_id": "xcsh-docs:resources:cminstance:reference", "path": "docs/guides/resources--cminstance--properties--api_token.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/properties/api_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_token for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_token

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md)
- [Property reference](resources--cminstance--reference.md)
- api_token

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
api_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--cminstance--properties--api_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cminstance--properties--api_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [api_token.blindfold_secret_info](resources--cminstance--properties--api_token--blindfold_secret_info.md)
- [api_token.clear_secret_info](resources--cminstance--properties--api_token--clear_secret_info.md)
- [Property reference](resources--cminstance--reference.md)
- [xcsh_cminstance](../resources/cminstance.md)
