---
page_title: "job.volumes"
subcategory: "Container"
description: "job.volumes for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4187, "body_sha256": "sha256:285c171f9bc0c9f797b3dd70b9e5026fda43f8a2636da0c7d22bbb206ce97bcc", "child_ids": ["xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "xcsh-docs:resources:workload:properties:job:volumes:host_path", "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes", "parent_id": "xcsh-docs:resources:workload:properties:job", "path": "documentation/resources/workload/properties/job/volumes/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["job", "volumes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.volumes for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- job.volumes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Volumes. Volumes for the job.

Upstream description:

Volumes for the job.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path"),
  validators.ConflictingListObjectAttributes("empty_dir",
    "persistent_volume"),
  validators.ConflictingListObjectAttributes("host_path",
    "persistent_volume")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
volumes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/host_path/): complete subsection reference.

<a id="schema-job--volumes--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/): complete subsection reference.

## Next pages

- [job.volumes.empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/empty_dir/)
- [job.volumes.host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/host_path/)
- [job.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
