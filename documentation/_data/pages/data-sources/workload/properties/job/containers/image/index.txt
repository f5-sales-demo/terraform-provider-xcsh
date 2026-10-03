---
page_title: "job.containers.image"
subcategory: "Container"
description: "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any."
xcsh_docs: {"aliases": ["job containers image"], "body_bytes": 5632, "body_sha256": "sha256:4402b9bbbdf95636468908b99116c3a61771ac750548e5bc613be3c8453a1621", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:containers:image:container_registry", "xcsh-docs:data-sources:workload:properties:job:containers:image:public"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:containers:image", "parent_id": "xcsh-docs:data-sources:workload:properties:job:containers", "path": "documentation/data-sources/workload/properties/job/containers/image/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3000331122110313-2022230231322222-3121103022103100-2023330210021132-2231323222300332-3033133010213110-3122112202323201-3121212030130230", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "image"], "schema_version": 1, "sections": [{"aliases": ["job containers image container registry"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:image:container_registry", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "image", "container_registry"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers image name"], "anchor": "schema-job--containers--image--name", "description": "Name is a container image which are usually given a name such as alpine, ubuntu, or quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:image", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "image", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job containers image public"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:image:public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "image", "public"], "syntax": "attribute", "type": "object"}, {"aliases": ["job containers image pull policy"], "anchor": "schema-job--containers--image--pull_policy", "description": "Image pull policy type enumerates the policy choices to use for pulling the image prior to starting the workload - IMAGE_PULL_POLICY_DEFAULT: Default Default will always pull image if :latest tag is specified in image name. If :latest tag is not specified in image name, it will pull image only if it does not already", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:image", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "image", "pull_policy"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/containers/image/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ImageType configures the image to use, how to pull the image, and the associated secrets to use if any.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.image

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- job.containers.image

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/container_registry/): complete subsection reference.

<a id="schema-job--containers--image--name"></a>

### name property

Type: `"string"`. Computed.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

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

- [public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/public/): complete subsection reference.

<a id="schema-job--containers--image--pull_policy"></a>

### pull_policy property

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

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

## Next pages

- [job.containers.image.container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/container_registry/)
- [job.containers.image.public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/image/public/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
