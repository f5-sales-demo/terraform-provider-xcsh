---
page_title: "protected_cookies.ignore_samesite"
subcategory: "Load Balancing"
description: "protected_cookies.ignore_samesite for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1051, "body_sha256": "sha256:d525c8dca98429bba85ec4415aa0e29f565bdb4d77863659eed44c8e0e03b762", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "path": "docs/guides/resources--http_loadbalancer--properties--protected_cookies--ignore_samesite.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protected_cookies", "ignore_samesite"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/protected_cookies/ignore_samesite/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protected_cookies.ignore_samesite for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies.ignore_samesite

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [protected_cookies](resources--http_loadbalancer--properties--protected_cookies.md)
- protected_cookies.ignore_samesite

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
ignore_samesite = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protected_cookies](resources--http_loadbalancer--properties--protected_cookies.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
