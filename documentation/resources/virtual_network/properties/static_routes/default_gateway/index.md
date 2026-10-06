---
page_title: "static_routes.default_gateway"
subcategory: "Networking"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["static routes default gateway"], "body_bytes": 1029, "body_sha256": "sha256:7474e80fa584397a307d5d0c27e6bffa1d8eb63f4aa780db27aa17c2ab87a1ff", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:static_routes:default_gateway", "parent_id": "xcsh-docs:resources:virtual_network:properties:static_routes", "path": "documentation/resources/virtual_network/properties/static_routes/default_gateway/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3221203221010200-0123123213133132-0330102021123113-2001033320003001-2213101113330113-2000003033002131-3011331102000311-2013033232022230", "registry_path": "docs/guides/resources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["static_routes", "default_gateway"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/static_routes/default_gateway/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# static_routes.default_gateway

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/static_routes/)
- static_routes.default_gateway

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

This is an empty object or choice marker. It has no direct properties.
