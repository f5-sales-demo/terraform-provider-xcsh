---
page_title: "stateful_service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "stateful_service.deploy_options.deploy_re_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2135, "body_sha256": "sha256:d2e0d5908df3cb419bae83b6ace3854eadff505cd65f6a6e4d2384f88bef8776", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites:virtual_site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "path": "documentation/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.deploy_options.deploy_re_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
