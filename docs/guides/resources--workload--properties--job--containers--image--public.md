---
page_title: "job.containers.image.public"
subcategory: "Container"
description: "job.containers.image.public for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1104, "body_sha256": "sha256:67cacbf2015084447d6dcaf31b750aa5b186ef3654c102123a6c362ee1c04d7c", "canonical_id": "xcsh-docs:resources:workload:properties:job:containers:image:public", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:image:public", "parent_id": "xcsh-docs:resources:workload:properties:job:containers:image", "path": "docs/guides/resources--workload--properties--job--containers--image--public.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "containers", "image", "public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/image/public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.containers.image.public for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.image.public

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.containers](resources--workload--properties--job--containers.md)
- [job.containers.image](resources--workload--properties--job--containers--image.md)
- job.containers.image.public

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
public = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [job.containers.image](resources--workload--properties--job--containers--image.md)
- [xcsh_workload](../resources/workload.md)
