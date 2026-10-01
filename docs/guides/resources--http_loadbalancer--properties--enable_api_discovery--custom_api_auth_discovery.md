---
page_title: "enable_api_discovery.custom_api_auth_discovery"
subcategory: "Load Balancing"
description: "enable_api_discovery.custom_api_auth_discovery for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1459, "body_sha256": "sha256:3ad39b416ba3a30140485a1de9810cf73ba69657febce6b117a48aa56bf5d288", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "path": "docs/guides/resources--http_loadbalancer--properties--enable_api_discovery--custom_api_auth_discovery.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.custom_api_auth_discovery for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.custom_api_auth_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md)
- enable_api_discovery.custom_api_auth_discovery

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_discovery_ref](resources--http_loadbalancer--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md): complete subsection reference.

## Next pages

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md)
- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
