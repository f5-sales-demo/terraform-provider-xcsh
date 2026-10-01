---
page_title: "default_pool.advanced_options.http1_config"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.http1_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1487, "body_sha256": "sha256:1b28c5872f8388d66dd68479b71ce5b1bc488e1cd2f261e887cf5cdfea627c43", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.http1_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- default_pool.advanced_options.http1_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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
http1_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config--header_transformation.md): complete subsection reference.

## Next pages

- [default_pool.advanced_options.http1_config.header_transformation](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config--header_transformation.md)
- [default_pool.advanced_options](resources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
