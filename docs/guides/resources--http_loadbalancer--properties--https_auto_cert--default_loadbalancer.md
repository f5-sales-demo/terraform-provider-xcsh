---
page_title: "https_auto_cert.default_loadbalancer"
subcategory: "Load Balancing"
description: "https_auto_cert.default_loadbalancer for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1085, "body_sha256": "sha256:9a8ce75ccc95806e09e449ef2c23c208ffe064f94f13c3602f9ccf5b940cfd28", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:default_loadbalancer", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:default_loadbalancer", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "path": "docs/guides/resources--http_loadbalancer--properties--https_auto_cert--default_loadbalancer.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "default_loadbalancer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/default_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.default_loadbalancer for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.default_loadbalancer

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- https_auto_cert.default_loadbalancer

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
