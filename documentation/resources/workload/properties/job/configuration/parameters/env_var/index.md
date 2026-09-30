---
page_title: "job.configuration.parameters.env_var"
subcategory: "Container"
description: "job.configuration.parameters.env_var for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3906, "body_sha256": "sha256:d2faab972c0c544345e403b2e792d610de856499b85d691cec2cbaf425f6ceda", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:configuration:parameters:env_var", "parent_id": "xcsh-docs:resources:workload:properties:job:configuration:parameters", "path": "documentation/resources/workload/properties/job/configuration/parameters/env_var/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["job", "configuration", "parameters", "env_var"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/configuration/parameters/env_var/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.configuration.parameters.env_var for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

Upstream description:

Environment Variable.

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

Upstream description:

Name of Environment Variable.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Value of Environment Variable.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [job.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/configuration/parameters/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
