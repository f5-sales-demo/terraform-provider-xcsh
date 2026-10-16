---
page_title: "service.scale_to_zero"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service scale to zero"], "body_bytes": 969, "body_sha256": "sha256:6f36a096947bd0ccbc8e6e92a1a6305a7091286b2f51801926dae05686bdcde8", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:scale_to_zero", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "documentation/resources/workload/properties/service/scale_to_zero/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "scale_to_zero"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/scale_to_zero/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["workloadCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
scale_to_zero = {}
```

This is an empty object or choice marker. It has no direct properties.
