---
page_title: "service.volumes.persistent_volume"
subcategory: "Container"
description: "service.volumes.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1987, "body_sha256": "sha256:9ffcfc1bbe2b197e1df49bdafe466af91c4ad62b62f1e1d41a7e3f6fa72ca285", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes", "path": "documentation/resources/workload/properties/service/volumes/persistent_volume/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["service", "volumes", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.volumes.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- service.volumes.persistent_volume

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

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
persistent_volume {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/storage/): complete subsection reference.

## Next pages

- [service.volumes.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/mount/)
- [service.volumes.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/storage/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
