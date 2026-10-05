---
page_title: "protected_cookies.samesite_strict"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["protected cookies samesite strict"], "body_bytes": 1308, "body_sha256": "sha256:3bd10405ac0d0b2957b47d8aa615c30f1f1b0e2601a6edbd103b9f659c09f645", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "path": "documentation/resources/http_loadbalancer/properties/protected_cookies/samesite_strict/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protected_cookies", "samesite_strict"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/protected_cookies/samesite_strict/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies.samesite_strict

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [protected_cookies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/)
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

- [protected_cookies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
