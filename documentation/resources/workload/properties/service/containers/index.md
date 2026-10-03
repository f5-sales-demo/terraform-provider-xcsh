---
page_title: "service.containers"
subcategory: "Container"
description: "Containers to use for service."
xcsh_docs: {"aliases": ["service containers"], "body_bytes": 9227, "body_sha256": "sha256:2365b019c376cbdeaa01efbef442b19f473eb387a2d2db34bbdec45bf68c8198", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:containers:custom_flavor", "xcsh-docs:resources:workload:properties:service:containers:default_flavor", "xcsh-docs:resources:workload:properties:service:containers:image", "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "xcsh-docs:resources:workload:properties:service:containers:readiness_check"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "documentation/resources/workload/properties/service/containers/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [{"anchor": "schema-service--containers--flavor", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers", "type": "conflicts"}, {"anchor": "schema-service--containers--flavor", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:custom_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:custom_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:custom_flavor,default_flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:default_flavor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers:ConflictingListObjectAttributes:default_flavor,flavor", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:default_flavor", "type": "conflicts"}, {"anchor": "schema-service--containers--name", "enforcement": "provider-schema", "group": "service.containers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers"], "schema_version": 1, "sections": [{"aliases": ["service containers args"], "anchor": "schema-service--containers--args", "description": "Arguments to the entrypoint. Overrides the docker image's CMD.", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "args"], "syntax": "attribute", "type": "list"}, {"aliases": ["service containers command"], "anchor": "schema-service--containers--command", "description": "Command to execute. Overrides the docker image's ENTRYPOINT.", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "command"], "syntax": "attribute", "type": "list"}, {"aliases": ["service containers custom flavor"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:custom_flavor", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--containers--custom_flavor--name", "enforcement": "provider-schema", "group": "service.containers.custom_flavor:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:custom_flavor", "type": "requires"}], "schema_path": ["service", "containers", "custom_flavor"], "syntax": "block", "type": "object"}, {"aliases": ["service containers default flavor"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:default_flavor", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "default_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["service containers flavor"], "anchor": "schema-service--containers--flavor", "description": "Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "flavor"], "syntax": "attribute", "type": "string"}, {"aliases": ["service containers image"], "anchor": "section", "description": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:image", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.image:ConflictingObjectAttributes:container_registry,public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:image:container_registry", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.image:ConflictingObjectAttributes:container_registry,public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:image:public", "type": "conflicts"}, {"anchor": "schema-service--containers--image--name", "enforcement": "provider-schema", "group": "service.containers.image:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:image", "type": "requires"}], "schema_path": ["service", "containers", "image"], "syntax": "block", "type": "object"}, {"aliases": ["service containers init container"], "anchor": "schema-service--containers--init_container", "description": "Specialized container that runs before application container and runs to completion.", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "init_container"], "syntax": "attribute", "type": "bool"}, {"aliases": ["service containers liveness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-service--containers--liveness_check--healthy_threshold", "enforcement": "provider-schema", "group": "service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-service--containers--liveness_check--interval", "enforcement": "provider-schema", "group": "service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-service--containers--liveness_check--timeout", "enforcement": "provider-schema", "group": "service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "type": "requires"}, {"anchor": "schema-service--containers--liveness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "service.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "type": "requires"}], "schema_path": ["service", "containers", "liveness_check"], "syntax": "block", "type": "object"}, {"aliases": ["service containers name"], "anchor": "schema-service--containers--name", "description": "Name of the container.", "document_id": "xcsh-docs:resources:workload:properties:service:containers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["service containers readiness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.containers.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-service--containers--readiness_check--healthy_threshold", "enforcement": "provider-schema", "group": "service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-service--containers--readiness_check--interval", "enforcement": "provider-schema", "group": "service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-service--containers--readiness_check--timeout", "enforcement": "provider-schema", "group": "service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "type": "requires"}, {"anchor": "schema-service--containers--readiness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "service.containers.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:readiness_check", "type": "requires"}], "schema_path": ["service", "containers", "readiness_check"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Containers to use for service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["workloadCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- service.containers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

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

<a id="schema-service--containers--args"></a>

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

<a id="schema-service--containers--command"></a>

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

- [custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/custom_flavor/): complete subsection reference.

- [default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/default_flavor/): complete subsection reference.

<a id="schema-service--containers--flavor"></a>

### flavor property

Type: `"string"`. Optional.

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

- [image](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/image/): complete subsection reference.

<a id="schema-service--containers--init_container"></a>

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

- [liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/): complete subsection reference.

<a id="schema-service--containers--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the container.

Upstream description:

Name of the container.

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

- [readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/readiness_check/): complete subsection reference.

## Next pages

- [service.containers.custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/custom_flavor/)
- [service.containers.default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/default_flavor/)
- [service.containers.image](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/image/)
- [service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/)
- [service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/readiness_check/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
