---
page_title: "simple_service.configuration"
subcategory: "Container"
description: "simple_service.configuration for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1143, "body_sha256": "sha256:afee6167792bc45eb5971bf34cc0a9f5fbaf77150caf158b1838ba19eb3c08af", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service:configuration", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration:parameters"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:configuration", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "docs/guides/resources--workload--properties--simple_service--configuration.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "configuration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.configuration for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [simple_service](resources--workload--properties--simple_service.md)
- simple_service.configuration

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

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
configuration {
  # Configure direct properties listed below.
}
```

## Direct properties

- [parameters](resources--workload--properties--simple_service--configuration--parameters.md): complete subsection reference.

## Next pages

- [simple_service.configuration.parameters](resources--workload--properties--simple_service--configuration--parameters.md)
- [simple_service](resources--workload--properties--simple_service.md)
- [xcsh_workload](../resources/workload.md)
