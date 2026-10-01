---
page_title: "discovery_target.nginx_instance"
subcategory: ""
description: "discovery_target.nginx_instance for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1532, "body_sha256": "sha256:bf27e25554314b05b87326d1f351b7b0d1451643c4f0a1cc3290ec001e32eedc", "canonical_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "child_ids": ["xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance:nginx_instance"], "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target:nginx_instance", "parent_id": "xcsh-docs:resources:nginx_service_discovery:properties:discovery_target", "path": "docs/guides/resources--nginx_service_discovery--properties--discovery_target--nginx_instance.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_target", "nginx_instance"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/properties/discovery_target/nginx_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_target.nginx_instance for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_target.nginx_instance

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md)
- [Property reference](resources--nginx_service_discovery--reference.md)
- [discovery_target](resources--nginx_service_discovery--properties--discovery_target.md)
- discovery_target.nginx_instance

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

NGINXInstance Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nginx_instance")}
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
nginx_instance {
  # Configure direct properties listed below.
}
```

## Direct properties

- [nginx_instance](resources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md): complete subsection reference.

## Next pages

- [discovery_target.nginx_instance.nginx_instance](resources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md)
- [discovery_target](resources--nginx_service_discovery--properties--discovery_target.md)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md)
