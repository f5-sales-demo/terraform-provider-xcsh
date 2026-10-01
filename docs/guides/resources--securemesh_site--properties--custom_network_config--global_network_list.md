---
page_title: "custom_network_config.global_network_list"
subcategory: ""
description: "custom_network_config.global_network_list for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1620, "body_sha256": "sha256:85ecbe680ad6d0bc37160e4d0fcffc05a019733cde7e2be109be8f094d252612", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list:global_network_connections"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--global_network_list.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "global_network_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.global_network_list for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- custom_network_config.global_network_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
```

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
global_network_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_network_connections](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections.md): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
