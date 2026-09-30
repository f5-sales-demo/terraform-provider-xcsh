---
page_title: "stateful_service.volumes.host_path"
subcategory: "Container"
description: "stateful_service.volumes.host_path for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2188, "body_sha256": "sha256:68d8bc479a62883ea0ad554578e69cef19dea0cb5bc047ee33aa0bc770a74639", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes:host_path", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:volumes:host_path:mount"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes:host_path", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes", "path": "docs/guides/data-sources--workload--properties--stateful_service--volumes--host_path.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "volumes", "host_path"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/volumes/host_path/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.volumes.host_path for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.volumes.host_path

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.volumes](data-sources--workload--properties--stateful_service--volumes.md)
- stateful_service.volumes.host_path

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing a host mapped path into the workload.

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

- [mount](data-sources--workload--properties--stateful_service--volumes--host_path--mount.md): complete subsection reference.

<a id="schema-stateful_service--volumes--host_path--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

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

- [stateful_service.volumes.host_path.mount](data-sources--workload--properties--stateful_service--volumes--host_path--mount.md)
- [stateful_service.volumes](data-sources--workload--properties--stateful_service--volumes.md)
- [xcsh_workload](../data-sources/workload.md)
