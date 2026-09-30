---
page_title: "protected_cookies.ignore_secure"
subcategory: "Load Balancing"
description: "protected_cookies.ignore_secure for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 939, "body_sha256": "sha256:d16c9e8c8f8c7a37cd69dfd23127a4b8d55db1780d27f508d9952f2372140ba7", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies:ignore_secure", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies:ignore_secure", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies", "path": "docs/guides/resources--cdn_loadbalancer--properties--protected_cookies--ignore_secure.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protected_cookies", "ignore_secure"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/protected_cookies/ignore_secure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protected_cookies.ignore_secure for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protected_cookies.ignore_secure

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [protected_cookies](resources--cdn_loadbalancer--properties--protected_cookies.md)
- protected_cookies.ignore_secure

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
ignore_secure = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protected_cookies](resources--cdn_loadbalancer--properties--protected_cookies.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
