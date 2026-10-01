---
page_title: "virtual_server.connection_rate_limit_mode"
subcategory: ""
description: "virtual_server.connection_rate_limit_mode for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 7466, "body_sha256": "sha256:39a8d758dd0fbe65b6b191c39148b601f6e31b5aad0a1750d80448f5d614267e", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.connection_rate_limit_mode for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
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

- [per_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/): complete subsection reference.

- [per_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/): complete subsection reference.

- [per_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/): complete subsection reference.

- [per_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server/): complete subsection reference.

- [per_virtual_server_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/): complete subsection reference.

- [per_virtual_server_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/): complete subsection reference.

- [per_virtual_server_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/): complete subsection reference.

## Next pages

- [virtual_server.connection_rate_limit_mode.per_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
