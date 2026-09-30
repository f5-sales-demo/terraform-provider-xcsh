---
page_title: "https_auto_cert.coalescing_options.strict_coalescing"
subcategory: "Load Balancing"
description: "https_auto_cert.coalescing_options.strict_coalescing for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1172, "body_sha256": "sha256:faebd07af4f6cdc3cd6737d549ae824436a40c8e8645df9680c12337a094744f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options:strict_coalescing", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options:strict_coalescing", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options", "path": "docs/guides/resources--http_loadbalancer--properties--https_auto_cert--coalescing_options--strict_coalescing.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "coalescing_options", "strict_coalescing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/coalescing_options/strict_coalescing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.coalescing_options.strict_coalescing for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_auto_cert.coalescing_options.strict_coalescing

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https_auto_cert](resources--http_loadbalancer--properties--https_auto_cert.md)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--properties--https_auto_cert--coalescing_options.md)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--properties--https_auto_cert--coalescing_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
