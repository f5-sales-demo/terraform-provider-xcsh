---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "virtual_server.virtual_server_state for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1592, "body_sha256": "sha256:8cf4368621146dc65c4a011268e3e4a0d8ed3a6169f2a5f68d9605436866e21a", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server--virtual_server_state.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.virtual_server_state for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- virtual_server.virtual_server_state

<a id="section"></a>

Type: `"single"`. Computed.

Displays the current state on the object.

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

## Direct properties

- [state_disabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md): complete subsection reference.

- [state_enabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md): complete subsection reference.

## Next pages

- [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md)
- [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
