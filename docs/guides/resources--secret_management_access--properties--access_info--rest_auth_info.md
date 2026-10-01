---
page_title: "access_info.rest_auth_info"
subcategory: ""
description: "access_info.rest_auth_info for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2290, "body_sha256": "sha256:82a40416eed8d1aeb84b251644ec7bc2951766a7466195f501fd31bf7f47322f", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:basic_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:headers_auth", "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info:query_params_auth"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:rest_auth_info", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info", "path": "docs/guides/resources--secret_management_access--properties--access_info--rest_auth_info.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "rest_auth_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/rest_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.rest_auth_info for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.rest_auth_info

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- access_info.rest_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for REST based hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("basic_auth",
    "headers_auth"),
  validators.ConflictingObjectAttributes("basic_auth",
    "query_params_auth"),
  validators.ConflictingObjectAttributes("headers_auth",
    "query_params_auth")}
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
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

Terraform syntax:

```terraform
rest_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [basic_auth](resources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md): complete subsection reference.

- [headers_auth](resources--secret_management_access--properties--access_info--rest_auth_info--headers_auth.md): complete subsection reference.

- [query_params_auth](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md): complete subsection reference.

## Next pages

- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--properties--access_info--rest_auth_info--basic_auth.md)
- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--properties--access_info--rest_auth_info--headers_auth.md)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--properties--access_info--rest_auth_info--query_params_auth.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
