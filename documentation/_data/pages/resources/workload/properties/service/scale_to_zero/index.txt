---
page_title: "service.scale_to_zero"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service scale to zero"], "body_bytes": 1203, "body_sha256": "sha256:5ca160081bf383bb1f027621e4880ec7474b9606bb36e0d4df7b7e5c28d8f527", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:scale_to_zero", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "documentation/resources/workload/properties/service/scale_to_zero/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "scale_to_zero"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/scale_to_zero/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.scale_to_zero

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- service.scale_to_zero

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for scale to zero.

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
scale_to_zero = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
