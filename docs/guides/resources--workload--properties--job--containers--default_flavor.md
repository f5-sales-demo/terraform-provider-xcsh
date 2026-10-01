---
page_title: "job.containers.default_flavor"
subcategory: "Container"
description: "job.containers.default_flavor for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1043, "body_sha256": "sha256:c919cc1e78e0c3bacf3a3bb5ee7232faa958995bb475ee9965193ac436c890d2", "canonical_id": "xcsh-docs:resources:workload:properties:job:containers:default_flavor", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:default_flavor", "parent_id": "xcsh-docs:resources:workload:properties:job:containers", "path": "docs/guides/resources--workload--properties--job--containers--default_flavor.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "containers", "default_flavor"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/default_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.containers.default_flavor for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.default_flavor

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.containers](resources--workload--properties--job--containers.md)
- job.containers.default_flavor

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

- [job.containers](resources--workload--properties--job--containers.md)
- [xcsh_workload](../resources/workload.md)
