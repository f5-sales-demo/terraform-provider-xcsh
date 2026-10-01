---
page_title: "virtual_server.connection_rate_limit_mode"
subcategory: ""
description: "virtual_server.connection_rate_limit_mode for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3985, "body_sha256": "sha256:1b214474585c1f01f5c7bb8d0b8e90d1e2099c586245ae83d4b164b178dd3e4d", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.connection_rate_limit_mode for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- virtual_server.connection_rate_limit_mode

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for connection rate limit mode.

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

## Direct properties

- [per_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md): complete subsection reference.

- [per_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md): complete subsection reference.

- [per_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md): complete subsection reference.

- [per_virtual_server](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md): complete subsection reference.

- [per_virtual_server_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md): complete subsection reference.

- [per_virtual_server_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md): complete subsection reference.

- [per_virtual_server_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md): complete subsection reference.

## Next pages

- [virtual_server.connection_rate_limit_mode.per_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
