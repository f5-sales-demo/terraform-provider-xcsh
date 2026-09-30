---
page_title: "service.configuration.parameters"
subcategory: "Container"
description: "service.configuration.parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2170, "body_sha256": "sha256:246da2253fbf7962f8de0ba14e22e4a94d98fc1f0a10c02b1ff39b928212bbc9", "canonical_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters", "child_ids": ["xcsh-docs:resources:workload:properties:service:configuration:parameters:env_var", "xcsh-docs:resources:workload:properties:service:configuration:parameters:file"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:configuration:parameters", "parent_id": "xcsh-docs:resources:workload:properties:service:configuration", "path": "docs/guides/resources--workload--properties--service--configuration--parameters.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "configuration", "parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.configuration.parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.configuration](resources--workload--properties--service--configuration.md)
- service.configuration.parameters

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [env_var](resources--workload--properties--service--configuration--parameters--env_var.md): complete subsection reference.

- [file](resources--workload--properties--service--configuration--parameters--file.md): complete subsection reference.

## Next pages

- [service.configuration.parameters.env_var](resources--workload--properties--service--configuration--parameters--env_var.md)
- [service.configuration.parameters.file](resources--workload--properties--service--configuration--parameters--file.md)
- [service.configuration](resources--workload--properties--service--configuration.md)
- [xcsh_workload](../resources/workload.md)
