---
page_title: "simple_service.configuration.parameters.file"
subcategory: "Container"
description: "Configuration File for the workload."
xcsh_docs: {"aliases": ["simple service configuration parameters file"], "body_bytes": 5265, "body_sha256": "sha256:c3e90706a6d2da76cab7e6242cb9c2ee7daafade3d2ead10c3eb9e32d95e1a08", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file:mount"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters", "path": "documentation/data-sources/workload/properties/simple_service/configuration/parameters/file/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3002023310331013-0011321210222212-2220122301103122-0200211333123133-1333101122021012-0002032101121130-0023031100212100-1023121300121132", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration", "parameters", "file"], "schema_version": 1, "sections": [{"aliases": ["simple service configuration parameters file data"], "anchor": "schema-simple_service--configuration--parameters--file--data", "description": "File data", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "data"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service configuration parameters file mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file:mount", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service configuration parameters file name"], "anchor": "schema-simple_service--configuration--parameters--file--name", "description": "Name of the file.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service configuration parameters file volume name"], "anchor": "schema-simple_service--configuration--parameters--file--volume_name", "description": "Name of the Volume.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file", "volume_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/configuration/parameters/file/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration File for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration.parameters.file

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/)
- simple_service.configuration.parameters.file

<a id="section"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

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

<a id="schema-simple_service--configuration--parameters--file--data"></a>

### data property

Type: `"string"`. Computed.

Data. File data

Upstream description:

File data

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/file/mount/): complete subsection reference.

<a id="schema-simple_service--configuration--parameters--file--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the file.

Upstream description:

Name of the file.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [simple_service.configuration.parameters.file.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/file/mount/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
