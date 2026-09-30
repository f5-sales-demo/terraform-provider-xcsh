---
page_title: "response_cache.default_response_cache_parameters"
subcategory: "DNS"
description: "response_cache.default_response_cache_parameters for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1032, "body_sha256": "sha256:37eafb776c7dc88819928a48440c1a514c18733a0618889ca12f8762d4d69db6", "canonical_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "path": "docs/guides/resources--dns_load_balancer--properties--response_cache--default_response_cache_parameters.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cache", "default_response_cache_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cache.default_response_cache_parameters for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cache.default_response_cache_parameters

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
- [Property reference](resources--dns_load_balancer--reference.md)
- [response_cache](resources--dns_load_balancer--properties--response_cache.md)
- response_cache.default_response_cache_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default response cache parameters.

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
default_response_cache_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cache](resources--dns_load_balancer--properties--response_cache.md)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
