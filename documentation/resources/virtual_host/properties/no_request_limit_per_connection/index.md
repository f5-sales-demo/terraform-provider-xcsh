---
page_title: "no_request_limit_per_connection"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no request limit per connection"], "body_bytes": 929, "body_sha256": "sha256:d7bd42e8dd21e2b819d2d63143b337ca4345392b13a4401f31edc4d4257ecdfe", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:no_request_limit_per_connection", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/no_request_limit_per_connection/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1201021232211320-1113031321300211-0230223223322233-2013130232110130-1300111221111201-1102111321313123-0013222103000031-1210301000110230", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_request_limit_per_connection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/no_request_limit_per_connection/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_request_limit_per_connection

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- no_request_limit_per_connection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

Additional upstream details:

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.
