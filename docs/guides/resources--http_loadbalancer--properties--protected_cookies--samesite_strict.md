---
page_title: "protected_cookies.samesite_strict"
subcategory: "Load Balancing"
description: "protected_cookies.samesite_strict for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 952, "body_sha256": "sha256:1b85f526fc4d529946ac23148c885e419b26a5762472bdeaf9a5c605022f91f7", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "path": "docs/guides/resources--http_loadbalancer--properties--protected_cookies--samesite_strict.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protected_cookies", "samesite_strict"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/protected_cookies/samesite_strict/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protected_cookies.samesite_strict for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protected_cookies.samesite_strict

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [protected_cookies](resources--http_loadbalancer--properties--protected_cookies.md)
- protected_cookies.samesite_strict

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
samesite_strict = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protected_cookies](resources--http_loadbalancer--properties--protected_cookies.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
