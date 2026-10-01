---
page_title: "job.deploy_options"
subcategory: "Container"
description: "job.deploy_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2458, "body_sha256": "sha256:d21d9bc3f48e8561b9340af72b18639a95adc5f7c8be54d65ebbd0c1712fd6fb", "canonical_id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:deploy_options:all_res", "xcsh-docs:data-sources:workload:properties:job:deploy_options:default_virtual_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_sites", "xcsh-docs:data-sources:workload:properties:job:deploy_options:deploy_re_virtual_sites"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:deploy_options", "parent_id": "xcsh-docs:data-sources:workload:properties:job", "path": "docs/guides/data-sources--workload--properties--job--deploy_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "deploy_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/deploy_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.deploy_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [job](data-sources--workload--properties--job.md)
- job.deploy_options

<a id="section"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

## Direct properties

- [all_res](data-sources--workload--properties--job--deploy_options--all_res.md): complete subsection reference.

- [default_virtual_sites](data-sources--workload--properties--job--deploy_options--default_virtual_sites.md): complete subsection reference.

- [deploy_ce_sites](data-sources--workload--properties--job--deploy_options--deploy_ce_sites.md): complete subsection reference.

- [deploy_ce_virtual_sites](data-sources--workload--properties--job--deploy_options--deploy_ce_virtual_sites.md): complete subsection reference.

- [deploy_re_sites](data-sources--workload--properties--job--deploy_options--deploy_re_sites.md): complete subsection reference.

- [deploy_re_virtual_sites](data-sources--workload--properties--job--deploy_options--deploy_re_virtual_sites.md): complete subsection reference.

## Next pages

- [job.deploy_options.all_res](data-sources--workload--properties--job--deploy_options--all_res.md)
- [job.deploy_options.default_virtual_sites](data-sources--workload--properties--job--deploy_options--default_virtual_sites.md)
- [job.deploy_options.deploy_ce_sites](data-sources--workload--properties--job--deploy_options--deploy_ce_sites.md)
- [job.deploy_options.deploy_ce_virtual_sites](data-sources--workload--properties--job--deploy_options--deploy_ce_virtual_sites.md)
- [job.deploy_options.deploy_re_sites](data-sources--workload--properties--job--deploy_options--deploy_re_sites.md)
- [job.deploy_options.deploy_re_virtual_sites](data-sources--workload--properties--job--deploy_options--deploy_re_virtual_sites.md)
- [job](data-sources--workload--properties--job.md)
- [xcsh_workload](../data-sources/workload.md)
