---
page_title: "stateful_service.deploy_options.deploy_re_sites"
subcategory: "Container"
description: "stateful_service.deploy_options.deploy_re_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1530, "body_sha256": "sha256:dbd68afad23dabd3a06ba0d0ef194194cdd83558cfc54be3f20da61713e877fd", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_sites", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_sites:site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options:deploy_re_sites", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:deploy_options", "path": "docs/guides/resources--workload--properties--stateful_service--deploy_options--deploy_re_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "deploy_options", "deploy_re_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/deploy_options/deploy_re_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.deploy_options.deploy_re_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.deploy_options.deploy_re_sites

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.deploy_options](resources--workload--properties--stateful_service--deploy_options.md)
- stateful_service.deploy_options.deploy_re_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--workload--properties--stateful_service--deploy_options--deploy_re_sites--site.md): complete subsection reference.

## Next pages

- [stateful_service.deploy_options.deploy_re_sites.site](resources--workload--properties--stateful_service--deploy_options--deploy_re_sites--site.md)
- [stateful_service.deploy_options](resources--workload--properties--stateful_service--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
