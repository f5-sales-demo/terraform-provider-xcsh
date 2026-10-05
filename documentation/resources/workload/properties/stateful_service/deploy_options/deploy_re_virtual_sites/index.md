---
page_title: "stateful_service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Regional Edge virtual sites."
xcsh_docs: {"aliases": ["stateful service deploy options deploy re virtual sites"], "body_bytes": 2135, "body_sha256": "sha256:d2e0d5908df3cb419bae83b6ace3854eadff505cd65f6a6e4d2384f88bef8776", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "path": "documentation/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033", "registry_path": "docs/guides/resources--workload--reference--group-029.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.deploy_options.deploy_re_virtual_sites:RequiredObjectAttributes:virtual_site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["stateful service deploy options deploy re virtual sites virtual site"], "anchor": "section", "description": "Which regional edge virtual sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-stateful_service--deploy_options--deploy_re_virtual_sites--virtual_site--name", "enforcement": "provider-schema", "group": "stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site", "type": "requires"}], "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines a way to deploy a workload on specific Regional Edge virtual sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/)
- stateful_service.deploy_options.deploy_re_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

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

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/virtual_site/): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/virtual_site/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
