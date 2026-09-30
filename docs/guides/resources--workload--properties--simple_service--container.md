---
page_title: "simple_service.container"
subcategory: "Container"
description: "simple_service.container for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 8018, "body_sha256": "sha256:b50d715984703c2cfeb7a29afc1126d3363f590448209e81dc1573b396778033", "canonical_id": "xcsh-docs:resources:workload:properties:simple_service:container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:custom_flavor", "xcsh-docs:resources:workload:properties:simple_service:container:default_flavor", "xcsh-docs:resources:workload:properties:simple_service:container:image", "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "docs/guides/resources--workload--properties--simple_service--container.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "container"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.container for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# simple_service.container

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [simple_service](resources--workload--properties--simple_service.md)
- simple_service.container

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ContainerType configures the container information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingObjectAttributes("default_flavor",
    "flavor")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flavor_choice": "[\"custom_flavor\",\"default_flavor\",\"flavor\"]"
}
```

Terraform syntax:

```terraform
container {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-simple_service--container--args"></a>

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

<a id="schema-simple_service--container--command"></a>

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

- [custom_flavor](resources--workload--properties--simple_service--container--custom_flavor.md): complete subsection reference.

- [default_flavor](resources--workload--properties--simple_service--container--default_flavor.md): complete subsection reference.

<a id="schema-simple_service--container--flavor"></a>

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

- [image](resources--workload--properties--simple_service--container--image.md): complete subsection reference.

<a id="schema-simple_service--container--init_container"></a>

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

- [liveness_check](resources--workload--properties--simple_service--container--liveness_check.md): complete subsection reference.

<a id="schema-simple_service--container--name"></a>

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

- [readiness_check](resources--workload--properties--simple_service--container--readiness_check.md): complete subsection reference.

## Next pages

- [simple_service.container.custom_flavor](resources--workload--properties--simple_service--container--custom_flavor.md)
- [simple_service.container.default_flavor](resources--workload--properties--simple_service--container--default_flavor.md)
- [simple_service.container.image](resources--workload--properties--simple_service--container--image.md)
- [simple_service.container.liveness_check](resources--workload--properties--simple_service--container--liveness_check.md)
- [simple_service.container.readiness_check](resources--workload--properties--simple_service--container--readiness_check.md)
- [simple_service](resources--workload--properties--simple_service.md)
- [xcsh_workload](../resources/workload.md)
