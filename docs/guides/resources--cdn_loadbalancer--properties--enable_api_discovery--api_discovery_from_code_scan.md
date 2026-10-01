---
page_title: "enable_api_discovery.api_discovery_from_code_scan"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_discovery_from_code_scan for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1561, "body_sha256": "sha256:1825b28cc68b4300149a7025bc0c1c67bb1324b8ac3b0896aa1cb9a7028eab2d", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "path": "docs/guides/resources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_discovery_from_code_scan for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery.md)
- enable_api_discovery.api_discovery_from_code_scan

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

## Direct properties

- [code_base_integrations](resources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- [enable_api_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
