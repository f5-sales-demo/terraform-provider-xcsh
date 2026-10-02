---
page_title: "stateful_service.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["stateful service configuration parameters"], "body_bytes": 2893, "body_sha256": "sha256:43d130e660272c30817299b592e3d36743aa54c558af7ed702682cc2a8401ef4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:env_var", "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration", "path": "documentation/resources/workload/properties/stateful_service/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:env_var", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:env_var", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "env_var"], "syntax": "block", "type": "object"}, {"aliases": ["file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--configuration--parameters--file--name", "enforcement": "provider-schema", "group": "stateful_service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file", "type": "requires"}, {"anchor": "schema-stateful_service--configuration--parameters--file--volume_name", "enforcement": "provider-schema", "group": "stateful_service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:configuration:parameters:file", "type": "requires"}], "schema_path": ["stateful_service", "configuration", "parameters", "file"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
