---
page_title: "service.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["service configuration parameters"], "body_bytes": 2767, "body_sha256": "sha256:d19d7c64a3c95a2e2c34680c07acd2214e15a3d5b57efb1c6c1cbf8eb8a662f0", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:configuration:parameters:env_var", "xcsh-docs:resources:workload:properties:service:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:configuration:parameters", "parent_id": "xcsh-docs:resources:workload:properties:service:configuration", "path": "documentation/resources/workload/properties/service/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:env_var", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "configuration", "parameters", "env_var"], "syntax": "block", "type": "object"}, {"aliases": ["file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--configuration--parameters--file--name", "enforcement": "provider-schema", "group": "service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file", "type": "requires"}, {"anchor": "schema-service--configuration--parameters--file--volume_name", "enforcement": "provider-schema", "group": "service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:configuration:parameters:file", "type": "requires"}], "schema_path": ["service", "configuration", "parameters", "file"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/)
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

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/file/): complete subsection reference.

## Next pages

- [service.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/env_var/)
- [service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/parameters/file/)
- [service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
