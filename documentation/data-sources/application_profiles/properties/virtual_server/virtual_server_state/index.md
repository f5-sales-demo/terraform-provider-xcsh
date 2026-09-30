---
page_title: "virtual_server.virtual_server_state"
subcategory: ""
description: "virtual_server.virtual_server_state for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1946, "body_sha256": "sha256:ca7ddd41011f30195f4ddfdb5f10126d30f1a8b4f7432a35e35487093dc6101a", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state:state_enabled"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/virtual_server_state/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["virtual_server", "virtual_server_state"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/virtual_server_state/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.virtual_server_state for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# virtual_server.virtual_server_state

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
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

- [state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/): complete subsection reference.

- [state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/): complete subsection reference.

## Next pages

- [virtual_server.virtual_server_state.state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/)
- [virtual_server.virtual_server_state.state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
