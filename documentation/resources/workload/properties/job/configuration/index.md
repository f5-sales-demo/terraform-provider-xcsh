---
page_title: "job.configuration"
subcategory: "Container"
description: "Configuration parameters of the workload."
xcsh_docs: {"aliases": ["job configuration"], "body_bytes": 1399, "body_sha256": "sha256:c76bfcb3f72083f76282654554aeb0bc40712243c805312bd2770c4ff81e558c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:configuration:parameters"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:configuration", "parent_id": "xcsh-docs:resources:workload:properties:job", "path": "documentation/resources/workload/properties/job/configuration/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1320213102010111-2221222102021030-0023213300022333-0003323203322003-2120122032110222-1301331212222103-1023321133120121-1202322320111132", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "configuration"], "schema_version": 1, "sections": [{"aliases": ["job configuration parameters"], "anchor": "section", "description": "Parameters for the workload.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:file", "type": "conflicts"}], "schema_path": ["job", "configuration", "parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/configuration/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameters of the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- job.configuration

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

- [parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/): complete subsection reference.

## Next pages

- [job.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
