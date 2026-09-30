---
page_title: "single_lb_app.enable_discovery.custom_api_auth_discovery"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.custom_api_auth_discovery for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1532, "body_sha256": "sha256:f8f40a5bc734270b89b6836561cd3925ccdac51d41f2be2e4350b605e79db7a1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery:api_discovery_ref"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--custom_api_auth_discovery.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "custom_api_auth_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.custom_api_auth_discovery for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# single_lb_app.enable_discovery.custom_api_auth_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- single_lb_app.enable_discovery.custom_api_auth_discovery

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

- [api_discovery_ref](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--custom_api_auth_discovery--api_discovery_ref.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--custom_api_auth_discovery--api_discovery_ref.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
