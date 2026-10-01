---
page_title: "service.scale_to_zero"
subcategory: "Container"
description: "service.scale_to_zero for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 946, "body_sha256": "sha256:e00728780513a1a4b737b36b32c12d5e8bbff3cbea161f074ea91419023ff23d", "canonical_id": "xcsh-docs:resources:workload:properties:service:scale_to_zero", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:scale_to_zero", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "docs/guides/resources--workload--properties--service--scale_to_zero.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "scale_to_zero"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/scale_to_zero/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.scale_to_zero for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.scale_to_zero

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- service.scale_to_zero

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for scale to zero.

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
scale_to_zero = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service](resources--workload--properties--service.md)
- [xcsh_workload](../resources/workload.md)
