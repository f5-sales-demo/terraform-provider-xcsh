---
page_title: "service.deploy_options.deploy_ce_virtual_sites"
subcategory: "Container"
description: "service.deploy_options.deploy_ce_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2026, "body_sha256": "sha256:133a41202192575d8add0ecc093f91a2e94c7588771a855f5f3c12472389dd4c", "child_ids": ["xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_virtual_sites:virtual_site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "documentation/resources/workload/properties/service/deploy_options/deploy_ce_virtual_sites/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["service", "deploy_options", "deploy_ce_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/deploy_ce_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.deploy_options.deploy_ce_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.deploy_ce_virtual_sites

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- service.deploy_options.deploy_ce_virtual_sites

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

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/deploy_ce_virtual_sites/virtual_site/): complete subsection reference.

## Next pages

- [service.deploy_options.deploy_ce_virtual_sites.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/deploy_ce_virtual_sites/virtual_site/)
- [service.deploy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/deploy_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
