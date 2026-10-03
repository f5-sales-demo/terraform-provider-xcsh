---
page_title: "stateful_service.containers"
subcategory: "Container"
description: "Containers to use for service."
xcsh_docs: {"aliases": ["stateful service containers"], "body_bytes": 8369, "body_sha256": "sha256:4e92b789e6cabc51dbb620984f9161e441f687befe97f3ccb37ef6ac0e562f90", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:containers:custom_flavor", "xcsh-docs:data-sources:workload:properties:stateful_service:containers:default_flavor", "xcsh-docs:data-sources:workload:properties:stateful_service:containers:image", "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check", "xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "path": "documentation/data-sources/workload/properties/stateful_service/containers/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003", "registry_path": "docs/guides/data-sources--workload--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "containers"], "schema_version": 1, "sections": [{"aliases": ["stateful service containers args"], "anchor": "schema-stateful_service--containers--args", "description": "Arguments to the entrypoint. Overrides the docker image's CMD.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "args"], "syntax": "attribute", "type": "list"}, {"aliases": ["stateful service containers command"], "anchor": "schema-stateful_service--containers--command", "description": "Command to execute. Overrides the docker image's ENTRYPOINT.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "command"], "syntax": "attribute", "type": "list"}, {"aliases": ["stateful service containers custom flavor"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:custom_flavor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "containers", "custom_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers default flavor"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:default_flavor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "default_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers flavor"], "anchor": "schema-stateful_service--containers--flavor", "description": "Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "flavor"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers image"], "anchor": "section", "description": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:image", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "containers", "image"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers init container"], "anchor": "schema-stateful_service--containers--init_container", "description": "Specialized container that runs before application container and runs to completion.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "init_container"], "syntax": "attribute", "type": "bool"}, {"aliases": ["stateful service containers liveness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:liveness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "containers", "liveness_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers name"], "anchor": "schema-stateful_service--containers--name", "description": "Name of the container.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers readiness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "containers", "readiness_check"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/containers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Containers to use for service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- stateful_service.containers

<a id="section"></a>

Type: `"list"`. Computed.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

<a id="schema-stateful_service--containers--args"></a>

### args property

Type: `["list", "string"]`. Computed.

Arguments to the entrypoint. Overrides the docker image's CMD.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="schema-stateful_service--containers--command"></a>

### command property

Type: `["list", "string"]`. Computed.

Command to execute. Overrides the docker image's ENTRYPOINT.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/custom_flavor/): complete subsection reference.

- [default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/default_flavor/): complete subsection reference.

<a id="schema-stateful_service--containers--flavor"></a>

### flavor property

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/image/): complete subsection reference.

<a id="schema-stateful_service--containers--init_container"></a>

### init_container property

Type: `"bool"`. Computed.

Specialized container that runs before application container and runs to completion.

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

- [liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/liveness_check/): complete subsection reference.

<a id="schema-stateful_service--containers--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the container.

Upstream description:

Name of the container.

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

- [readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/readiness_check/): complete subsection reference.

## Next pages

- [stateful_service.containers.custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/custom_flavor/)
- [stateful_service.containers.default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/default_flavor/)
- [stateful_service.containers.image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/image/)
- [stateful_service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/liveness_check/)
- [stateful_service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/containers/readiness_check/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
