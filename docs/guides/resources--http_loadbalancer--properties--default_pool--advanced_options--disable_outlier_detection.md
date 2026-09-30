---
page_title: "default_pool.advanced_options.disable_outlier_detection"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.disable_outlier_detection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1240, "body_sha256": "sha256:3e3089f26769d62c1c76a2bd7450d45be6330d7deaf94528cb181adb65d794b6", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--disable_outlier_detection.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "disable_outlier_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/disable_outlier_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.disable_outlier_detection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.advanced_options.disable_outlier_detection

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.disable_outlier_detection

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
disable_outlier_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
