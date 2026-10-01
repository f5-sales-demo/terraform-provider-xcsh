---
page_title: "virtual_server.virtual_server_state.state_enabled"
subcategory: ""
description: "virtual_server.virtual_server_state.state_enabled for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1259, "body_sha256": "sha256:0f76312bc2233ecb7110a5770a7be5a9a8c4cc887cb63db4c1572cf1d67bd820", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "path": "docs/guides/resources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "virtual_server_state", "state_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.virtual_server_state.state_enabled for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state.state_enabled

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [virtual_server.virtual_server_state](resources--application_profiles--properties--virtual_server--virtual_server_state.md)
- virtual_server.virtual_server_state.state_enabled

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
state_enabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [virtual_server.virtual_server_state](resources--application_profiles--properties--virtual_server--virtual_server_state.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
