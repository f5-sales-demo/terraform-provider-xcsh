---
page_title: "virtual_server.auto_last_hop.auto_last_hop_default"
subcategory: ""
description: "virtual_server.auto_last_hop.auto_last_hop_default for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1273, "body_sha256": "sha256:a371c32561f043561c8e6698d27b87fe23b7c62b752038de5d051934f20eaac5", "canonical_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop:auto_last_hop_default", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop", "path": "docs/guides/resources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_default.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server", "auto_last_hop", "auto_last_hop_default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.auto_last_hop.auto_last_hop_default for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.auto_last_hop.auto_last_hop_default

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [virtual_server.auto_last_hop](resources--application_profiles--properties--virtual_server--auto_last_hop.md)
- virtual_server.auto_last_hop.auto_last_hop_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for auto last hop default.

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
auto_last_hop_default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [virtual_server.auto_last_hop](resources--application_profiles--properties--virtual_server--auto_last_hop.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
