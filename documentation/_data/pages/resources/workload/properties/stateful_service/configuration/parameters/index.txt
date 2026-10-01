---
page_title: "stateful_service.configuration.parameters"
subcategory: "Container"
description: "stateful_service.configuration.parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2893, "body_sha256": "sha256:43d130e660272c30817299b592e3d36743aa54c558af7ed702682cc2a8401ef4", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:env_var", "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration", "path": "documentation/resources/workload/properties/stateful_service/configuration/parameters/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["stateful_service", "configuration", "parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.configuration.parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/)
- stateful_service.configuration.parameters

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

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/parameters/file/): complete subsection reference.

## Next pages

- [stateful_service.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/parameters/env_var/)
- [stateful_service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/parameters/file/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
