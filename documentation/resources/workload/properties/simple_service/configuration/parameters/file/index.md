---
page_title: "simple_service.configuration.parameters.file"
subcategory: "Container"
description: "Configuration File for the workload."
xcsh_docs: {"aliases": ["simple service configuration parameters file"], "body_bytes": 5933, "body_sha256": "sha256:f67690cdbe88a049671875a1d3b3204f42234a57da08520016e423476a86e687", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters", "path": "documentation/resources/workload/properties/simple_service/configuration/parameters/file/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3232233330311213-3312202130112202-2202331322032213-3123201210003330-2110033310211003-0103230313133313-3012010123011230-2022023131233102", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "schema-simple_service--configuration--parameters--file--name", "enforcement": "provider-schema", "group": "simple_service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "type": "requires"}, {"anchor": "schema-simple_service--configuration--parameters--file--volume_name", "enforcement": "provider-schema", "group": "simple_service.configuration.parameters.file:RequiredObjectAttributes:name,volume_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration", "parameters", "file"], "schema_version": 1, "sections": [{"aliases": ["data"], "anchor": "schema-simple_service--configuration--parameters--file--data", "description": "File data", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "data"], "syntax": "attribute", "type": "string"}, {"aliases": ["mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--configuration--parameters--file--mount--mount_path", "enforcement": "provider-schema", "group": "simple_service.configuration.parameters.file.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file:mount", "type": "requires"}], "schema_path": ["simple_service", "configuration", "parameters", "file", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-simple_service--configuration--parameters--file--name", "description": "Name of the file.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["volume name"], "anchor": "schema-simple_service--configuration--parameters--file--volume_name", "description": "Name of the Volume.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:configuration:parameters:file", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "volume_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/configuration/parameters/file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration File for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration.parameters.file

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/)
- simple_service.configuration.parameters.file

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-simple_service--configuration--parameters--file--data"></a>

### data property

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/file/mount/): complete subsection reference.

<a id="schema-simple_service--configuration--parameters--file--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-simple_service--configuration--parameters--file--volume_name"></a>

### volume_name property

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [simple_service.configuration.parameters.file.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/file/mount/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/configuration/parameters/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
