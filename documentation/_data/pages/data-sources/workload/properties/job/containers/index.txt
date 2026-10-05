---
page_title: "job.containers"
subcategory: "Container"
description: "Containers to use for the job."
xcsh_docs: {"aliases": ["job containers"], "body_bytes": 8031, "body_sha256": "sha256:60000cd7c3561654d8a5ae22fb7cbbb5a99c333679432c1e2f0ad576e18abe9f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:containers:custom_flavor", "xcsh-docs:data-sources:workload:properties:job:containers:default_flavor", "xcsh-docs:data-sources:workload:properties:job:containers:image", "xcsh-docs:data-sources:workload:properties:job:containers:liveness_check", "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:containers", "parent_id": "xcsh-docs:data-sources:workload:properties:job", "path": "documentation/data-sources/workload/properties/job/containers/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0230022332122200-0100033311212312-3320320103332110-3121012220331330-2123300012201302-0300033022332313-1123221133202010-0322013102131123", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers"], "schema_version": 1, "sections": [{"aliases": ["job containers args"], "anchor": "schema-job--containers--args", "description": "Arguments to the entrypoint. Overrides the docker image's CMD.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "args"], "syntax": "attribute", "type": "list"}, {"aliases": ["job containers command"], "anchor": "schema-job--containers--command", "description": "Command to execute. Overrides the docker image's ENTRYPOINT.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "command"], "syntax": "attribute", "type": "list"}, {"aliases": ["job containers custom flavor"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:custom_flavor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "custom_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers default flavor"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:default_flavor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "default_flavor"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers flavor"], "anchor": "schema-job--containers--flavor", "description": "Container Flavor type - CONTAINER_FLAVOR_TYPE_TINY: Tiny Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_MEDIUM: Medium Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER_FLAVOR_TYPE_LARGE: Large Large containers have limit of 1 vCPU and", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "flavor"], "syntax": "attribute", "type": "string"}, {"aliases": ["job containers image"], "anchor": "section", "description": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:image", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "image"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers init container"], "anchor": "schema-job--containers--init_container", "description": "Specialized container that runs before application container and runs to completion.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "init_container"], "syntax": "attribute", "type": "bool"}, {"aliases": ["job containers liveness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:liveness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "liveness_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers name"], "anchor": "schema-job--containers--name", "description": "Name of the container.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job containers readiness check"], "anchor": "section", "description": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "readiness_check"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/containers/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Containers to use for the job.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- job.containers

<a id="section"></a>

Type: `"list"`. Computed.

Containers. Containers to use for the job.

Upstream description:

Containers to use for the job.

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

## Direct properties

<a id="schema-job--containers--args"></a>

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

<a id="schema-job--containers--command"></a>

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

- [custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/custom_flavor/): complete subsection reference.

- [default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/default_flavor/): complete subsection reference.

<a id="schema-job--containers--flavor"></a>

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

- [image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/): complete subsection reference.

<a id="schema-job--containers--init_container"></a>

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

- [liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/): complete subsection reference.

<a id="schema-job--containers--name"></a>

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

- [readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/): complete subsection reference.

## Next pages

- [job.containers.custom_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/custom_flavor/)
- [job.containers.default_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/default_flavor/)
- [job.containers.image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/)
- [job.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/liveness_check/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
