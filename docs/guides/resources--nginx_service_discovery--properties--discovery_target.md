---
page_title: "discovery_target"
subcategory: ""
description: "discovery_target for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1667, "body_sha256": "sha256:bf0816c6cc5277bdb9b724682631f14ebcb135451af34e72e7850fe82cbcd7e4", "canonical_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance"], "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "parent_id": "xcsh-docs:resources:nginx_service_discovery:reference", "path": "docs/guides/resources--nginx_service_discovery--properties--discovery_target.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_target"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_target for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md)
- [Property reference](resources--nginx_service_discovery--reference.md)
- discovery_target

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("config_sync_group",
    "nginx_instance")}
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
  "x-ves-oneof-field-target": "[\"config_sync_group\",\"nginx_instance\"]"
}
```

Terraform syntax:

```terraform
discovery_target {
  # Configure direct properties listed below.
}
```

## Direct properties

- [config_sync_group](resources--nginx_service_discovery--properties--discovery_target--config_sync_group.md): complete subsection reference.

- [nginx_instance](resources--nginx_service_discovery--properties--discovery_target--nginx_instance.md): complete subsection reference.

## Next pages

- [discovery_target.config_sync_group](resources--nginx_service_discovery--properties--discovery_target--config_sync_group.md)
- [discovery_target.nginx_instance](resources--nginx_service_discovery--properties--discovery_target--nginx_instance.md)
- [Property reference](resources--nginx_service_discovery--reference.md)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md)
