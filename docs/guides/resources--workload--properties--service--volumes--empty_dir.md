---
page_title: "service.volumes.empty_dir"
subcategory: "Container"
description: "service.volumes.empty_dir for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2039, "body_sha256": "sha256:9439f84903690ef5b23f8b986683481e919bbf4c21b9858d14cc6394b0a1e8b3", "canonical_id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:empty_dir:mount"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes", "path": "docs/guides/resources--workload--properties--service--volumes--empty_dir.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "volumes", "empty_dir"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.volumes.empty_dir for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.volumes](resources--workload--properties--service--volumes.md)
- service.volumes.empty_dir

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
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
empty_dir {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](resources--workload--properties--service--volumes--empty_dir--mount.md): complete subsection reference.

<a id="schema-service--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [service.volumes.empty_dir.mount](resources--workload--properties--service--volumes--empty_dir--mount.md)
- [service.volumes](resources--workload--properties--service--volumes.md)
- [xcsh_workload](../resources/workload.md)
