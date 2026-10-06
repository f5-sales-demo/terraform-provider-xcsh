---
page_title: "https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation"
subcategory: "Load Balancing"
description: "Use the platform's current default HTTP header transformation behavior."
xcsh_docs: {"aliases": ["https http protocol options http protocol enable v1 only header transformation default header transformation"], "body_bytes": 1762, "body_sha256": "sha256:6e677a9d4a04e90d07c01735e2abb1743bb6acd8541bbd6940f915226a3d93c2", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "path": "documentation/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "default_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Use the platform's current default HTTP header transformation behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/)
- [https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/http_protocol_options/)
- [https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.
