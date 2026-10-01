---
page_title: "user_session_expiration.absolute_timeout"
subcategory: ""
description: "user_session_expiration.absolute_timeout for xcsh_tenant_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1855, "body_sha256": "sha256:85f90dc5c5ce94187a1e7fa6289e967a85bfe07e58f426a20389da98be4295f8", "canonical_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes"], "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "path": "docs/guides/resources--tenant_configuration--properties--user_session_expiration--absolute_timeout.md", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_session_expiration.absolute_timeout for xcsh_tenant_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
- [Property reference](resources--tenant_configuration--reference.md)
- [user_session_expiration](resources--tenant_configuration--properties--user_session_expiration.md)
- user_session_expiration.absolute_timeout

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents the session expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
absolute_timeout {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md): complete subsection reference.

- [minutes](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md): complete subsection reference.

## Next pages

- [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--hours.md)
- [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--properties--user_session_expiration--absolute_timeout--minutes.md)
- [user_session_expiration](resources--tenant_configuration--properties--user_session_expiration.md)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md)
