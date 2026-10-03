---
page_title: "stateful_service.configuration.parameters.file"
subcategory: "Container"
description: "Configuration File for the workload."
xcsh_docs: {"aliases": ["stateful service configuration parameters file"], "body_bytes": 5297, "body_sha256": "sha256:11e80a778b00860c181e6930cf1abd3f4b3d829a1955fbfd7c0b93a4e19f595b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file:mount"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters", "path": "documentation/data-sources/workload/properties/stateful_service/configuration/parameters/file/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1100303330211012-2000203213110113-0211120230032130-3232032133233010-1303222303021102-3312020133132200-2030323321211133-2112032122000213", "registry_path": "docs/guides/data-sources--workload--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "configuration", "parameters", "file"], "schema_version": 1, "sections": [{"aliases": ["stateful service configuration parameters file data"], "anchor": "schema-stateful_service--configuration--parameters--file--data", "description": "File data", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "file", "data"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service configuration parameters file mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file:mount", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "file", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service configuration parameters file name"], "anchor": "schema-stateful_service--configuration--parameters--file--name", "description": "Name of the file.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "file", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service configuration parameters file volume name"], "anchor": "schema-stateful_service--configuration--parameters--file--volume_name", "description": "Name of the Volume.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "file", "volume_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/configuration/parameters/file/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration File for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.configuration.parameters.file

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/)
- [stateful_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/)
- stateful_service.configuration.parameters.file

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

<a id="schema-stateful_service--configuration--parameters--file--data"></a>

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/file/mount/): complete subsection reference.

<a id="schema-stateful_service--configuration--parameters--file--name"></a>

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

<a id="schema-stateful_service--configuration--parameters--file--volume_name"></a>

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

- [stateful_service.configuration.parameters.file.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/file/mount/)
- [stateful_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
