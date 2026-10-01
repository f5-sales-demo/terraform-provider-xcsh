---
page_title: "simple_service.enabled.persistent_volume"
subcategory: "Container"
description: "simple_service.enabled.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1989, "body_sha256": "sha256:5b1329f61ea58a378b03ced2a68015b35be777781dc23b1556781daa71df6a03", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled", "path": "documentation/data-sources/workload/properties/simple_service/enabled/persistent_volume/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["simple_service", "enabled", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/enabled/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.enabled.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.enabled.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/)
- simple_service.enabled.persistent_volume

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/): complete subsection reference.

## Next pages

- [simple_service.enabled.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/mount/)
- [simple_service.enabled.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
