---
page_title: "ring_hash.hash_policy.cookie.ignore_httponly"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ring hash hash policy cookie ignore httponly"], "body_bytes": 1666, "body_sha256": "sha256:ca542ccd9890ce54eab610091a9ba8b4f7dcaf16f34c6d6e87bf4f053b6a3421", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie:ignore_httponly", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy:cookie", "path": "documentation/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_httponly/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ring_hash", "hash_policy", "cookie", "ignore_httponly"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/ignore_httponly/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash.hash_policy.cookie.ignore_httponly

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [ring_hash](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/)
- [ring_hash.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/)
- [ring_hash.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ring_hash.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/ring_hash/hash_policy/cookie/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
