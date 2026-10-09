---
page_title: "http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation"
subcategory: ""
description: "Use the platform's current default HTTP header transformation behavior."
xcsh_docs: {"aliases": ["http protocol options http protocol enable v1 only header transformation default header transformation"], "body_bytes": 1571, "body_sha256": "sha256:4a57a2b5599ce2ff6d1e30b2bf66ca69932ef76f436eef93efeea2c31e43fd52", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "parent_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "path": "documentation/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3021332131210123-3203200332330321-2021022202222020-3212310220323210-3211011132121111-2130313200201020-0103012312310322-3332312331012220", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_protocol_options", "http_protocol_enable_v1_only", "header_transformation", "default_header_transformation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Use the platform's current default HTTP header transformation behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/)
- [http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/)
- [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/)
- http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

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
