---
page_title: "service.deploy_options.deploy_re_virtual_sites"
subcategory: "Container"
description: "service.deploy_options.deploy_re_virtual_sites for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1535, "body_sha256": "sha256:a542d8ecc7c91f29d56071f1749020d50c24176eb50950106573926995d67f50", "canonical_id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_virtual_sites", "child_ids": ["xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_virtual_sites:virtual_site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_virtual_sites", "parent_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "path": "docs/guides/resources--workload--properties--service--deploy_options--deploy_re_virtual_sites.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "deploy_options", "deploy_re_virtual_sites"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/deploy_re_virtual_sites/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.deploy_options.deploy_re_virtual_sites for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.deploy_options.deploy_re_virtual_sites

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- service.deploy_options.deploy_re_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [virtual_site](resources--workload--properties--service--deploy_options--deploy_re_virtual_sites--virtual_site.md): complete subsection reference.

## Next pages

- [service.deploy_options.deploy_re_virtual_sites.virtual_site](resources--workload--properties--service--deploy_options--deploy_re_virtual_sites--virtual_site.md)
- [service.deploy_options](resources--workload--properties--service--deploy_options.md)
- [xcsh_workload](../resources/workload.md)
