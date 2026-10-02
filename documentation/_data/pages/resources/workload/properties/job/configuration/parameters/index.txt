---
page_title: "job.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["job configuration parameters"], "body_bytes": 2711, "body_sha256": "sha256:8856fbaff3b72dcab184dfd817465ca98d8da2fc795310e87fc377b6fc03df4c", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "xcsh-docs:resources:workload:properties:job:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:configuration:parameters", "parent_id": "xcsh-docs:resources:workload:properties:job:configuration", "path": "documentation/resources/workload/properties/job/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0133220330020323-3011333113101313-2130312310111112-0311033201321131-0021223021201022-0000333212001300-1011312123033212-1311222111312130", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.configuration.parameters:ConflictingListObjectAttributes:env_var,file", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:file", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "configuration", "parameters", "env_var"], "syntax": "block", "type": "object"}, {"aliases": ["file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:file", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--configuration--parameters--file--name", "enforcement": "provider-schema", "group": "job.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:file", "type": "requires"}, {"anchor": "schema-job--configuration--parameters--file--volume_name", "enforcement": "provider-schema", "group": "job.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:file", "type": "requires"}], "schema_path": ["job", "configuration", "parameters", "file"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/)
- job.configuration.parameters

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

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/file/): complete subsection reference.

## Next pages

- [job.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/env_var/)
- [job.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/file/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
