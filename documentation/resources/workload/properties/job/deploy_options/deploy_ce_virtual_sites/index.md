---
page_title: "job.deploy_options.deploy_ce_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Customer virtual sites."
xcsh_docs: {"aliases": ["job deploy options deploy ce virtual sites"], "body_bytes": 1982, "body_sha256": "sha256:850a746e090769d177779b701a649883cdc94c8db8d1f269bfc849636821c255", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:job:deploy_options", "path": "documentation/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3123031132321222-1211312320321323-2100330023231222-1011202220031232-2332313222331100-1333220330323313-0212113310011030-0000102112201113", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_ce_virtual_sites:RequiredObjectAttributes:virtual_site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "deploy_options", "deploy_ce_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["job deploy options deploy ce virtual sites virtual site"], "anchor": "section", "description": "Which customer virtual sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-job--deploy_options--deploy_ce_virtual_sites--virtual_site--name", "enforcement": "provider-schema", "group": "job.deploy_options.deploy_ce_virtual_sites.virtual_site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites:virtual_site", "type": "requires"}], "schema_path": ["job", "deploy_options", "deploy_ce_virtual_sites", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "This defines a way to deploy a workload on specific Customer virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["workloadCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_ce_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/)
- job.deploy_options.deploy_ce_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/virtual_site/): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_ce_virtual_sites.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/virtual_site/)
- [job.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
