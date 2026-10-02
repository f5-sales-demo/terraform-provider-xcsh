---
page_title: "stateful_service.deploy_options.deploy_ce_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Customer sites."
xcsh_docs: {"aliases": ["stateful service deploy options deploy ce sites"], "body_bytes": 2021, "body_sha256": "sha256:186467582df0fd010480d1cb715c1113918273e1a8cef62944060fc78f348b1e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_sites", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "path": "documentation/resources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200", "registry_path": "docs/guides/resources--workload--reference--group-029.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.deploy_options.deploy_ce_sites:RequiredObjectAttributes:site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_ce_sites"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "Which customer sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-stateful_service--deploy_options--deploy_ce_sites--site--name", "enforcement": "provider-schema", "group": "stateful_service.deploy_options.deploy_ce_sites.site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_sites:site", "type": "requires"}], "schema_path": ["stateful_service", "deploy_options", "deploy_ce_sites", "site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to deploy a workload on specific Customer sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options.deploy_ce_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/)
- stateful_service.deploy_options.deploy_ce_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/site/): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.deploy_ce_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/deploy_ce_sites/site/)
- [stateful_service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
