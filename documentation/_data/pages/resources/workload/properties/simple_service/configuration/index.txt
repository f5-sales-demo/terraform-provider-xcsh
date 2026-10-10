---
page_title: "simple_service.configuration"
subcategory: "Container"
description: "Configuration parameters of the workload."
xcsh_docs: {"aliases": ["simple service configuration"], "body_bytes": 1093, "body_sha256": "sha256:30d9d26e218c3ec486c90563b51fead88b1767968ae2f43a3036c5837c352b15", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration:parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:configuration", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "documentation/resources/workload/properties/simple_service/configuration/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0320212022302321-0101303101101210-3333223110220111-3211312311002320-0032113021302113-3220323310123330-0201320221111101-0213000020311313", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration"], "schema_version": 1, "sections": [{"aliases": ["simple service configuration parameters"], "anchor": "section", "description": "Parameters for the workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["simple_service", "configuration", "parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/configuration/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration parameters of the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- simple_service.configuration

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

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
configuration {
  # Configure direct properties listed below.
}
```

## Direct properties

- [parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/): complete subsection reference.
