---
page_title: "simple_service.disabled"
subcategory: "Container"
description: "simple_service.disabled for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 949, "body_sha256": "sha256:ed6183cc0ff2997e7ab26b4d8cd56bb2f5729912d2d635bfece51fe7200f6b62", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:disabled", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "docs/guides/resources--workload--properties--simple_service--disabled.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.disabled for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.disabled

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [simple_service](resources--workload--properties--simple_service.md)
- simple_service.disabled

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
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [simple_service](resources--workload--properties--simple_service.md)
- [xcsh_workload](../resources/workload.md)
