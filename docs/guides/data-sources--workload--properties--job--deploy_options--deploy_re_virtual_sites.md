---
page_title: "job.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "job.deploy_options.deploy_re_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1331, "body_sha256": "sha256:d7be15000bc0647ed5d6ceb00604bc6e5ac4ab816122f5790f240451c5aa4064", "canonical_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites:virtual_site"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "path": "docs/guides/data-sources--workload--properties--job--deploy_options--deploy_re_virtual_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.deploy_options.deploy_re_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [job](data-sources--workload--properties--job.md)
- [job.deploy_options](data-sources--workload--properties--job--deploy_options.md)
- job.deploy_options.deploy_re_virtual_sites

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

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

- [virtual_site](data-sources--workload--properties--job--deploy_options--deploy_re_virtual_sites--virtual_site.md): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_re_virtual_sites.virtual_site](data-sources--workload--properties--job--deploy_options--deploy_re_virtual_sites--virtual_site.md)
- [job.deploy_options](data-sources--workload--properties--job--deploy_options.md)
- [xcsh_workload](../data-sources/workload.md)
