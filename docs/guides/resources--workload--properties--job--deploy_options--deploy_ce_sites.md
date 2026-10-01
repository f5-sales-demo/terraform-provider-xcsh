---
page_title: "job.deploy_options.deploy_ce_sites"
subcategory: "Container"
description: "job.deploy_options.deploy_ce_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1476, "body_sha256": "sha256:c313d51dd8365f967bdbb7e1b88b0b423554466d228f56b59eabc31f8fecceea", "canonical_id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "child_ids": ["xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites:site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "parent_id": "xcsh-docs:resources:workload:properties:job:deploy_options", "path": "docs/guides/resources--workload--properties--job--deploy_options--deploy_ce_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "deploy_options", "deploy_ce_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/deploy_options/deploy_ce_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.deploy_options.deploy_ce_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options.deploy_ce_sites

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.deploy_options](resources--workload--properties--job--deploy_options.md)
- job.deploy_options.deploy_ce_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
```

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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--workload--properties--job--deploy_options--deploy_ce_sites--site.md): complete subsection reference.

## Next pages

- [job.deploy_options.deploy_ce_sites.site](resources--workload--properties--job--deploy_options--deploy_ce_sites--site.md)
- [job.deploy_options](resources--workload--properties--job--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
