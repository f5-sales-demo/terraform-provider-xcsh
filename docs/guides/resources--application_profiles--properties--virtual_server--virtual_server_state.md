---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "virtual_server.virtual_server_state for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1871, "body_sha256": "sha256:7df4a9b468c41181f07f893acc66a790f46b674dbfdcd15af001eda144191aca", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "docs/guides/resources--application_profiles--properties--virtual_server--virtual_server_state.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.virtual_server_state for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- virtual_server.virtual_server_state

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Displays the current state on the object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("state_disabled",
    "state_enabled")}
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
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

Terraform syntax:

```terraform
virtual_server_state {
  # Configure direct properties listed below.
}
```

## Direct properties

- [state_disabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md): complete subsection reference.

- [state_enabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md): complete subsection reference.

## Next pages

- [virtual_server.virtual_server_state.state_disabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md)
- [virtual_server.virtual_server_state.state_enabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
