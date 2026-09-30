---
page_title: "stateful_service.containers.default_flavor"
subcategory: "Container"
description: "stateful_service.containers.default_flavor for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1048, "body_sha256": "sha256:658cb006400a3225e243247545d9b462000a27dd93d74a861ccea2df40bacb00", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "path": "docs/guides/resources--workload--properties--stateful_service--containers--default_flavor.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "containers", "default_flavor"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/default_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.containers.default_flavor for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.containers.default_flavor

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.containers](resources--workload--properties--stateful_service--containers.md)
- stateful_service.containers.default_flavor

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [stateful_service.containers](resources--workload--properties--stateful_service--containers.md)
- [xcsh_workload](../resources/workload.md)
