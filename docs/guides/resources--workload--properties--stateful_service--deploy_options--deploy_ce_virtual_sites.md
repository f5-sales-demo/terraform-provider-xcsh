---
page_title: "stateful_service.deploy_options.deploy_ce_virtual_sites"
subcategory: "Container"
description: "stateful_service.deploy_options.deploy_ce_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1723, "body_sha256": "sha256:3f4b4c5d9f2e94ec3b3d9543a425803b321baf652a92d0ea532c8544f91b5413", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_virtual_sites", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_virtual_sites:virtual_site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_ce_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "path": "docs/guides/resources--workload--properties--stateful_service--deploy_options--deploy_ce_virtual_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_ce_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/deploy_options/deploy_ce_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.deploy_options.deploy_ce_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.deploy_options.deploy_ce_virtual_sites

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.deploy_options](resources--workload--properties--stateful_service--deploy_options.md)
- stateful_service.deploy_options.deploy_ce_virtual_sites

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

- [virtual_site](resources--workload--properties--stateful_service--deploy_options--deploy_ce_virtual_sites--virtual_site.md): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site](resources--workload--properties--stateful_service--deploy_options--deploy_ce_virtual_sites--virtual_site.md)
- [stateful_service.deploy_options](resources--workload--properties--stateful_service--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
