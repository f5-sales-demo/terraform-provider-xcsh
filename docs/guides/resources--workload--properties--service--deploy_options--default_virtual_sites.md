---
page_title: "service.deploy_options.default_virtual_sites"
subcategory: "Container"
description: "service.deploy_options.default_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1095, "body_sha256": "sha256:d9ba8f269b107431441a5df307fae44a7d48e0d75ee03c21112a6202afc6a5d9", "canonical_id": "xcsh-docs:resources:workload:properties:service:deploy_options:default_virtual_sites", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:default_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "docs/guides/resources--workload--properties--service--deploy_options--default_virtual_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "deploy_options", "default_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/default_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.deploy_options.default_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.deploy_options.default_virtual_sites

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- service.deploy_options.default_virtual_sites

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
default_virtual_sites = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
