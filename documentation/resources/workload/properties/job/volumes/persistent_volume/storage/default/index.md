---
page_title: "job.volumes.persistent_volume.storage.default"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["job volumes persistent volume storage default"], "body_bytes": 1415, "body_sha256": "sha256:edb70cd6b8e820bdfc0b879012791675d4c5f46d8522cb505b13bc13e1beb785", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage:default", "parent_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage", "path": "documentation/resources/workload/properties/job/volumes/persistent_volume/storage/default/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes", "persistent_volume", "storage", "default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/persistent_volume/storage/default/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.persistent_volume.storage.default

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/)
- [job.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/)
- [job.volumes.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/storage/)
- job.volumes.persistent_volume.storage.default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.
