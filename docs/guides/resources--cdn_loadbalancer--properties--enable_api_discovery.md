---
page_title: "enable_api_discovery"
subcategory: "Load Balancing"
description: "enable_api_discovery for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3460, "body_sha256": "sha256:d17163c850758e8d95c55283aea761c2b3936a3fe27301a31111ad88dcbb0026", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--enable_api_discovery.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_api_discovery

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- enable_api_discovery

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_crawler](resources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler.md): complete subsection reference.

- [api_discovery_from_code_scan](resources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md): complete subsection reference.

- [custom_api_auth_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery--custom_api_auth_discovery.md): complete subsection reference.

- [default_api_auth_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery--default_api_auth_discovery.md): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--properties--enable_api_discovery--disable_learn_from_redirect_traffic.md): complete subsection reference.

- [discovered_api_settings](resources--cdn_loadbalancer--properties--enable_api_discovery--discovered_api_settings.md): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--properties--enable_api_discovery--enable_learn_from_redirect_traffic.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery--custom_api_auth_discovery.md)
- [enable_api_discovery.default_api_auth_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery--default_api_auth_discovery.md)
- [enable_api_discovery.disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--properties--enable_api_discovery--disable_learn_from_redirect_traffic.md)
- [enable_api_discovery.discovered_api_settings](resources--cdn_loadbalancer--properties--enable_api_discovery--discovered_api_settings.md)
- [enable_api_discovery.enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--properties--enable_api_discovery--enable_learn_from_redirect_traffic.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
