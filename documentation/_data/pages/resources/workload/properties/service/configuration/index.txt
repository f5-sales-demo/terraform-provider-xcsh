---
page_title: "service.configuration"
subcategory: "Container"
description: "Configuration parameters of the workload."
xcsh_docs: {"aliases": ["service configuration"], "body_bytes": 1435, "body_sha256": "sha256:e462840d2498f7d34b192ba23d569e978fc2dd62f1a4c82aaa6f5746c0941511", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:configuration:parameters"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:configuration", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "documentation/resources/workload/properties/service/configuration/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "configuration"], "schema_version": 1, "sections": [{"aliases": ["service configuration parameters"], "anchor": "section", "description": "Parameters for the workload.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:env_var", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file", "type": "conflicts"}], "schema_path": ["service", "configuration", "parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/configuration/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameters of the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.configuration

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- service.configuration

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

- [parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/): complete subsection reference.

## Next pages

- [service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
