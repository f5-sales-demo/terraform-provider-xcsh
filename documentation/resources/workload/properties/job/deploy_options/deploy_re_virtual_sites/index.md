---
page_title: "job.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge virtual sites."
xcsh_docs: {"aliases": ["job deploy options deploy re virtual sites"], "body_bytes": 1447, "body_sha256": "sha256:7a05f5cc190620ac784c71f207fb1209b5f729a6945993b7326a9103da600a35", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:job:deploy_options", "path": "documentation/resources/workload/properties/job/deploy_options/deploy_re_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2303323000212010-1331320302002323-2131023001233320-2232113103000220-1031323101300203-3020202312102100-3332323210033031-2112133212011023", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_re_virtual_sites:RequiredObjectAttributes:virtual_site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites:virtual_site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["job deploy options deploy re virtual sites virtual site"], "anchor": "section", "description": "Which regional edge virtual sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-job--deploy_options--deploy_re_virtual_sites--virtual_site--name", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_re_virtual_sites.virtual_site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites:virtual_site", "type": "requires"}], "schema_path": ["job", "deploy_options", "deploy_re_virtual_sites", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/)
- job.deploy_options.deploy_re_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
```

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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_virtual_sites/virtual_site/): complete subsection reference.
