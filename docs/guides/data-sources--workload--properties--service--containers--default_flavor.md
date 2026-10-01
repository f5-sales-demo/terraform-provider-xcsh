---
page_title: "service.containers.default_flavor"
subcategory: "Container"
description: "service.containers.default_flavor for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1036, "body_sha256": "sha256:704b9a5af31e96ff7f5870517e3ae6064bab33af5de5eeface3da64a4f035d42", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:containers:default_flavor", "child_ids": [], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:containers:default_flavor", "parent_id": "xcsh-docs:data-sources:workload:properties:service:containers", "path": "docs/guides/data-sources--workload--properties--service--containers--default_flavor.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "containers", "default_flavor"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/containers/default_flavor/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.containers.default_flavor for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.default_flavor

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.containers](data-sources--workload--properties--service--containers.md)
- service.containers.default_flavor

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.containers](data-sources--workload--properties--service--containers.md)
- [xcsh_workload](../data-sources/workload.md)
