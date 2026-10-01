---
page_title: "virtual_server.connection_rate_limit_mode"
subcategory: ""
description: "virtual_server.connection_rate_limit_mode for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 6523, "body_sha256": "sha256:2d61dc8948e254b3def4695451bd856844d04623b1e4789de44010ed777ec0dd", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.connection_rate_limit_mode for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.connection_rate_limit_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for connection rate limit mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("per_destination_address",
    "per_source_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_source_destination_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_source_address",
    "per_virtual_server_source_destination_address")}
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
  "x-ves-oneof-field-connection_rate_limit_mode_choice": "[\"per_destination_address\",\"per_source_address\",\"per_source_destination_address\",\"per_virtual_server\",\"per_virtual_server_destination_address\",\"per_virtual_server_source_address\",\"per_virtual_server_source_destination_address\"]"
}
```

Terraform syntax:

```terraform
connection_rate_limit_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [per_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md): complete subsection reference.

- [per_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md): complete subsection reference.

- [per_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md): complete subsection reference.

- [per_virtual_server](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md): complete subsection reference.

- [per_virtual_server_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md): complete subsection reference.

- [per_virtual_server_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md): complete subsection reference.

- [per_virtual_server_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md): complete subsection reference.

## Next pages

- [virtual_server.connection_rate_limit_mode.per_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
