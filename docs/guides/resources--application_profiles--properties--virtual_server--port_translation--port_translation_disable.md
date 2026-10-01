---
page_title: "virtual_server.port_translation.port_translation_disable"
subcategory: ""
description: "virtual_server.port_translation.port_translation_disable for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1268, "body_sha256": "sha256:d9e6f85a2bf5546afa840191505778a9b4a79557ac09fe742fd6010f9a4c2e48", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "path": "docs/guides/resources--application_profiles--properties--virtual_server--port_translation--port_translation_disable.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "port_translation", "port_translation_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.port_translation.port_translation_disable for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation.port_translation_disable

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [virtual_server.port_translation](resources--application_profiles--properties--virtual_server--port_translation.md)
- virtual_server.port_translation.port_translation_disable

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
port_translation_disable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [virtual_server.port_translation](resources--application_profiles--properties--virtual_server--port_translation.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
