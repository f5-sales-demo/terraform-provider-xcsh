---
page_title: "job.configuration.parameters.env_var"
subcategory: "Container"
description: "Environment Variable."
xcsh_docs: {"aliases": ["job configuration parameters env var"], "body_bytes": 3647, "body_sha256": "sha256:9e7d1ee8087e61cb447efdb7ec1792e4ebe26a7aa61d9101dfa55fa4910d046e", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "parent_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters", "path": "documentation/resources/workload/properties/job/configuration/parameters/env_var/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1231200203031232-2130001123233033-0320323132230320-0212230232020300-0111030310202211-2010301223113312-3030222220202212-3120101023010021", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "configuration", "parameters", "env_var"], "schema_version": 1, "sections": [{"aliases": ["job configuration parameters env var name"], "anchor": "schema-job--configuration--parameters--env_var--name", "description": "Name of Environment Variable.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "configuration", "parameters", "env_var", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job configuration parameters env var value"], "anchor": "schema-job--configuration--parameters--env_var--value", "description": "Value of Environment Variable.", "document_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "configuration", "parameters", "env_var", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/configuration/parameters/env_var/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Environment Variable.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration.parameters.env_var

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/)
- [job.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/)
- job.configuration.parameters.env_var

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Environment Variable. Environment Variable.

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
env_var {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-job--configuration--parameters--env_var--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-job--configuration--parameters--env_var--value"></a>

### value property

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
