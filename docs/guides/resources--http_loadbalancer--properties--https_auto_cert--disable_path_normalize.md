---
page_title: "https_auto_cert.disable_path_normalize"
subcategory: "Load Balancing"
description: "https_auto_cert.disable_path_normalize for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1060, "body_sha256": "sha256:739c6cd2c84ef5d41c3cd3b7c2e27f131e39d2b3c5ee6a5d9a519689472ff5ed", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:disable_path_normalize", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:disable_path_normalize", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "path": "docs/guides/resources--http_loadbalancer--properties--https_auto_cert--disable_path_normalize.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "disable_path_normalize"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/disable_path_normalize/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.disable_path_normalize for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.disable_path_normalize

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- https_auto_cert.disable_path_normalize

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_path_normalize = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
