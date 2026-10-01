---
page_title: "service.volumes"
subcategory: "Container"
description: "service.volumes for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3245, "body_sha256": "sha256:3a60b87c33fb1326564daa03e7fb02e84cfc0cbbff61cd1f68ab1bd3663c58c9", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir", "xcsh-docs:data-sources:workload:properties:service:volumes:host_path", "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:volumes", "parent_id": "xcsh-docs:data-sources:workload:properties:service", "path": "docs/guides/data-sources--workload--properties--service--volumes.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "volumes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/volumes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.volumes for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- service.volumes

<a id="section"></a>

Type: `"list"`. Computed.

Volumes. Volumes for the service.

Upstream description:

Volumes for the service.

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

## Direct properties

- [empty_dir](data-sources--workload--properties--service--volumes--empty_dir.md): complete subsection reference.

- [host_path](data-sources--workload--properties--service--volumes--host_path.md): complete subsection reference.

<a id="schema-service--volumes--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the volume.

Upstream description:

Name of the volume.

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

- [persistent_volume](data-sources--workload--properties--service--volumes--persistent_volume.md): complete subsection reference.

## Next pages

- [service.volumes.empty_dir](data-sources--workload--properties--service--volumes--empty_dir.md)
- [service.volumes.host_path](data-sources--workload--properties--service--volumes--host_path.md)
- [service.volumes.persistent_volume](data-sources--workload--properties--service--volumes--persistent_volume.md)
- [service](data-sources--workload--properties--service.md)
- [xcsh_workload](../data-sources/workload.md)
