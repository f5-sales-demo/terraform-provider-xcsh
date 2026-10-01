---
page_title: "simple_service.do_not_advertise"
subcategory: "Container"
description: "simple_service.do_not_advertise for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1000, "body_sha256": "sha256:2af3393ab671b283ee664d7746d23b51b6fe60b661b87e763693c6a9345c7069", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:do_not_advertise", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "docs/guides/resources--workload--properties--simple_service--do_not_advertise.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "do_not_advertise"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/do_not_advertise/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.do_not_advertise for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.do_not_advertise

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [simple_service](resources--workload--properties--simple_service.md)
- simple_service.do_not_advertise

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

- [simple_service](resources--workload--properties--simple_service.md)
- [xcsh_workload](../resources/workload.md)
