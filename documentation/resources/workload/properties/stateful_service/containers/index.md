---
page_title: "stateful_service.containers"
subcategory: "Container"
description: "Containers to use for service."
xcsh_docs: {"aliases": ["stateful service containers"], "body_bytes": 8257, "body_sha256": "sha256:e87bb0c90204c6221b4afdbb7bbc9c620ef7277f41efc8b20336154742c024ae", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:custom_flavor", "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service", "path": "documentation/resources/workload/properties/stateful_service/containers/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [{"anchor": "schema-stateful_service--containers--flavor", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--flavor", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--name", "enforcement": "provider-schema", "group": "stateful_service.containers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "containers"], "schema_version": 1, "sections": [{"aliases": ["stateful service containers args"], "anchor": "schema-stateful_service--containers--args", "description": "Arguments to the entrypoint. Overrides the docker image's CMD.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "args"], "syntax": "attribute", "type": "list"}, {"aliases": ["stateful service containers command"], "anchor": "schema-stateful_service--containers--command", "description": "Command to execute. Overrides the docker image's ENTRYPOINT.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "command"], "syntax": "attribute", "type": "list"}, {"aliases": ["stateful service containers custom flavor"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:custom_flavor", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--containers--custom_flavor--name", "enforcement": "provider-schema", "group": "stateful_service.containers.custom_flavor:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:custom_flavor", "type": "requires"}], "schema_path": ["stateful_service", "containers", "custom_flavor"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service containers default flavor"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:default_flavor", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "default_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers flavor"], "anchor": "schema-stateful_service--containers--flavor", "description": "Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "flavor"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers image"], "anchor": "section", "description": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.image:ConflictingObjectAttributes:container_registry,public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image:container_registry", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.image:ConflictingObjectAttributes:container_registry,public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image:public", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--image--name", "enforcement": "provider-schema", "group": "stateful_service.containers.image:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "type": "requires"}], "schema_path": ["stateful_service", "containers", "image"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service containers init container"], "anchor": "schema-stateful_service--containers--init_container", "description": "Specialized container that runs before application container and runs to completion.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "init_container"], "syntax": "attribute", "type": "bool"}, {"aliases": ["stateful service containers liveness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--liveness_check--healthy_threshold", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--liveness_check--interval", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--liveness_check--timeout", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--liveness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "stateful_service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:liveness_check", "type": "requires"}], "schema_path": ["stateful_service", "containers", "liveness_check"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service containers name"], "anchor": "schema-stateful_service--containers--name", "description": "Name of the container.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers readiness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--readiness_check--healthy_threshold", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--readiness_check--interval", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--readiness_check--timeout", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-stateful_service--containers--readiness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "type": "requires"}], "schema_path": ["stateful_service", "containers", "readiness_check"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Containers to use for service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- stateful_service.containers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Containers. Containers to use for service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingListObjectAttributes("default_flavor",
    "flavor")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
containers {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-stateful_service--containers--args"></a>

### args property

Type: `["list", "string"]`. Optional.

Arguments to the entrypoint. Overrides the docker image's CMD.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="schema-stateful_service--containers--command"></a>

### command property

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the docker image's ENTRYPOINT.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/custom_flavor/): complete subsection reference.

- [default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/default_flavor/): complete subsection reference.

<a id="schema-stateful_service--containers--flavor"></a>

### flavor property

Type: `"string"`. Optional.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Additional upstream details:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"),
}
```

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

- [image](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/image/): complete subsection reference.

<a id="schema-stateful_service--containers--init_container"></a>

### init_container property

Type: `"bool"`. Optional.

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

- [liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/liveness_check/): complete subsection reference.

<a id="schema-stateful_service--containers--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the container.

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

- [readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/): complete subsection reference.
