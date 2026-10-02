---
page_title: "service.volumes.host_path"
subcategory: "Container"
description: "Volume containing a host mapped path into the workload."
xcsh_docs: {"aliases": ["service volumes host path"], "body_bytes": 2960, "body_sha256": "sha256:39cc94fc7cc4a4e4569e89a6250dde80a88fb60b27a471591151941f4b4b1921", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:host_path:mount"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:host_path", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes", "path": "documentation/resources/workload/properties/service/volumes/host_path/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2311311322220231-2030300230121230-0312211222112320-3012133212233011-3000301230332233-0220013112301302-0322020200213312-1303202001310223", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "schema-service--volumes--host_path--path", "enforcement": "provider-schema", "group": "service.volumes.host_path:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:host_path", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "host_path"], "schema_version": 1, "sections": [{"aliases": ["mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:host_path:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--volumes--host_path--mount--mount_path", "enforcement": "provider-schema", "group": "service.volumes.host_path.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:host_path:mount", "type": "requires"}], "schema_path": ["service", "volumes", "host_path", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "schema-service--volumes--host_path--path", "description": "Path of the directory on the host.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:host_path", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "host_path", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/host_path/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Volume containing a host mapped path into the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.host_path

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- service.volumes.host_path

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/host_path/mount/): complete subsection reference.

<a id="schema-service--volumes--host_path--path"></a>

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

- [service.volumes.host_path.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/host_path/mount/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
