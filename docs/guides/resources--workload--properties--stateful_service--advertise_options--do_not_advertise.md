---
page_title: "stateful_service.advertise_options.do_not_advertise"
subcategory: "Container"
description: "stateful_service.advertise_options.do_not_advertise for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1197, "body_sha256": "sha256:967ec3f6e622147bda04ffc7e1ed6589e06e80e852d6d56de94e94f1f951fc6d", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--do_not_advertise.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "do_not_advertise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/do_not_advertise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.do_not_advertise for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.do_not_advertise

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- stateful_service.advertise_options.do_not_advertise

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [xcsh_workload](../resources/workload.md)
