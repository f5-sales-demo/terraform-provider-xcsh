---
page_title: "virtual_server.virtual_server_state.state_disabled"
subcategory: ""
description: "virtual_server.virtual_server_state.state_disabled for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1262, "body_sha256": "sha256:8831272bf53dd325deb66bb075dd9ed1f5dda2d353c87a32a27f5ee458fd57e4", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "path": "docs/guides/resources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "virtual_server_state", "state_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.virtual_server_state.state_disabled for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state.state_disabled

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [virtual_server.virtual_server_state](resources--application_profiles--properties--virtual_server--virtual_server_state.md)
- virtual_server.virtual_server_state.state_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
state_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [virtual_server.virtual_server_state](resources--application_profiles--properties--virtual_server--virtual_server_state.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
