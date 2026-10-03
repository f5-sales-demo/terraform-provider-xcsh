---
page_title: "https.disable_path_normalize"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https disable path normalize"], "body_bytes": 1257, "body_sha256": "sha256:0a6a23cc6601bfd6e7685c2728b075b72a4c9aec9193375cf9ee622a8aeaebdc", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:disable_path_normalize", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "documentation/resources/http_loadbalancer/properties/https/disable_path_normalize/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "disable_path_normalize"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/disable_path_normalize/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.disable_path_normalize

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/)
- https.disable_path_normalize

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

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
