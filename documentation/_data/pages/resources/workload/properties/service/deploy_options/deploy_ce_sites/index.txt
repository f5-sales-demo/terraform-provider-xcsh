---
page_title: "service.deploy_options.deploy_ce_sites"
subcategory: "Container"
description: "This defines a way to deploy a workload on specific Customer sites."
xcsh_docs: {"aliases": ["service deploy options deploy ce sites"], "body_bytes": 1922, "body_sha256": "sha256:b55c1327be0be837f76a7b28208e4115d81ef73b1135b1696123e7e12f1a8770", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites:site"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "documentation/resources/workload/properties/service/deploy_options/deploy_ce_sites/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2030313122020332-3023312023002210-2233321221032220-2320312202220030-3123330003212320-0111110001002000-0220211312020031-3120030222010323", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.deploy_options.deploy_ce_sites:RequiredObjectAttributes:site", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites:site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "deploy_options", "deploy_ce_sites"], "schema_version": 1, "sections": [{"aliases": ["service deploy options deploy ce sites site"], "anchor": "section", "description": "Which customer sites should this workload be deployed.", "document_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites:site", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-service--deploy_options--deploy_ce_sites--site--name", "enforcement": "provider-schema", "group": "service.deploy_options.deploy_ce_sites.site:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites:site", "type": "requires"}], "schema_path": ["service", "deploy_options", "deploy_ce_sites", "site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/deploy_ce_sites/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to deploy a workload on specific Customer sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.deploy_ce_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- service.deploy_options.deploy_ce_sites

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/deploy_ce_sites/site/): complete subsection reference.

## Next pages

- [service.deploy_options.deploy_ce_sites.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/deploy_ce_sites/site/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
