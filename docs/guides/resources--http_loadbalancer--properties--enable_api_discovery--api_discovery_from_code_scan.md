---
page_title: "enable_api_discovery.api_discovery_from_code_scan"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_discovery_from_code_scan for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1570, "body_sha256": "sha256:ce0d0e5a1e9f83765bcce7628a6f294f10376c61c1f3514130fae734dabe606a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "path": "docs/guides/resources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_discovery_from_code_scan for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md)
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

- [code_base_integrations](resources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md)
- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
