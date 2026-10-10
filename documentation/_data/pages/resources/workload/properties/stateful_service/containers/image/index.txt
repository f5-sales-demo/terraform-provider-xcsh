---
page_title: "stateful_service.containers.image"
subcategory: "Container"
description: "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any."
xcsh_docs: {"aliases": ["stateful service containers image"], "body_bytes": 5038, "body_sha256": "sha256:b3a5bfaca145c4801653671f8ea23cd6f8e5b75b6b151dc678e0fab9963a437b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:image:container_registry", "xcsh-docs:resources:workload:properties:stateful_service:containers:image:public"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers", "path": "documentation/resources/workload/properties/stateful_service/containers/image/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "containers", "image"], "schema_version": 1, "sections": [{"aliases": ["stateful service containers image container registry"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image:container_registry", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "containers", "image", "container_registry"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service containers image name"], "anchor": "schema-stateful_service--containers--image--name", "description": "Name is a container image which are usually given a name such as alpine, ubuntu, or quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "image", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service containers image public"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image:public", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "image", "public"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service containers image pull policy"], "anchor": "schema-stateful_service--containers--image--pull_policy", "description": "Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:image", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "image", "pull_policy"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/image/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers.image

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/)
- stateful_service.containers.image

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

## Direct properties

- [container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/image/container_registry/): complete subsection reference.

<a id="schema-stateful_service--containers--image--name"></a>

### name property

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/image/public/): complete subsection reference.

<a id="schema-stateful_service--containers--image--pull_policy"></a>

### pull_policy property

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Additional upstream details:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
