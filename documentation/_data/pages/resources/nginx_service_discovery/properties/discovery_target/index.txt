---
page_title: "discovery_target"
subcategory: ""
description: "discovery_target for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2075, "body_sha256": "sha256:e6e5a92614083e20050f3db1183e30b72f6a3dd6d30ad03fa8115e58b8e80373", "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:config_sync_group", "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance"], "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "parent_id": "xcsh-docs:resources:nginx_service_discovery:reference", "path": "documentation/resources/nginx_service_discovery/properties/discovery_target/index.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["discovery_target"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_target for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
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

- [config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/): complete subsection reference.

- [nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/): complete subsection reference.

## Next pages

- [discovery_target.config_sync_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/config_sync_group/)
- [discovery_target.nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
