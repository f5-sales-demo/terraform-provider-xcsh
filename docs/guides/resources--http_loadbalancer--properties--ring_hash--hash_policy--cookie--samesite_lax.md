---
page_title: "ring_hash.hash_policy.cookie.samesite_lax"
subcategory: "Load Balancing"
description: "ring_hash.hash_policy.cookie.samesite_lax for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1277, "body_sha256": "sha256:79745f043719910dccd8718b73c5438ace668f550ae74020c1ec9d611bb1b3eb", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:samesite_lax", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "path": "docs/guides/resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie--samesite_lax.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie", "samesite_lax"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/samesite_lax/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash.hash_policy.cookie.samesite_lax for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie.samesite_lax

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [ring_hash](resources--http_loadbalancer--properties--ring_hash.md)
- [ring_hash.hash_policy](resources--http_loadbalancer--properties--ring_hash--hash_policy.md)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
- ring_hash.hash_policy.cookie.samesite_lax

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
samesite_lax = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--properties--ring_hash--hash_policy--cookie.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
