---
page_title: "service.deploy_options.all_res"
subcategory: "Container"
description: "service.deploy_options.all_res for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1053, "body_sha256": "sha256:41c0ba92937d3d9f13b710772bbacfb87d343629006d07d24156d691ed1b2879", "canonical_id": "xcsh-docs:resources:workload:properties:service:deploy_options:all_res", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:all_res", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "docs/guides/resources--workload--properties--service--deploy_options--all_res.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "deploy_options", "all_res"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/all_res/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.deploy_options.all_res for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.all_res

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- service.deploy_options.all_res

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
all_res = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
