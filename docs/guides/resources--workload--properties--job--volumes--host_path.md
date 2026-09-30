---
page_title: "job.volumes.host_path"
subcategory: "Container"
description: "job.volumes.host_path for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2411, "body_sha256": "sha256:e3f33eaaafb8ea87be672ac7acfb0b7fe5bb41e27e9fb70d82dbdfeab7a7e052", "canonical_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "child_ids": ["xcsh-docs:resources:workload:properties:job:volumes:host_path:mount"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "parent_id": "xcsh-docs:resources:workload:properties:job:volumes", "path": "docs/guides/resources--workload--properties--job--volumes--host_path.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "volumes", "host_path"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/host_path/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.volumes.host_path for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job.volumes.host_path

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.volumes](resources--workload--properties--job--volumes.md)
- job.volumes.host_path

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
host_path {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](resources--workload--properties--job--volumes--host_path--mount.md): complete subsection reference.

<a id="schema-job--volumes--host_path--path"></a>

### path property

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

## Next pages

- [job.volumes.host_path.mount](resources--workload--properties--job--volumes--host_path--mount.md)
- [job.volumes](resources--workload--properties--job--volumes.md)
- [xcsh_workload](../resources/workload.md)
