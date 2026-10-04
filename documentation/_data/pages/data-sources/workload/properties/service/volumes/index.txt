---
page_title: "service.volumes"
subcategory: "Container"
description: "Volumes for the service."
xcsh_docs: {"aliases": ["service volumes"], "body_bytes": 3796, "body_sha256": "sha256:4472977f52318f43e59bf947768471a9e1aa37b2669398bc0a669f9a778a8217", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir", "xcsh-docs:data-sources:workload:properties:service:volumes:host_path", "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:volumes", "parent_id": "xcsh-docs:data-sources:workload:properties:service", "path": "documentation/data-sources/workload/properties/service/volumes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes"], "schema_version": 1, "sections": [{"aliases": ["service volumes empty dir"], "anchor": "section", "description": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "volumes", "empty_dir"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes host path"], "anchor": "section", "description": "Volume containing a host mapped path into the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:host_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "volumes", "host_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes name"], "anchor": "schema-service--volumes--name", "description": "Name of the volume.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["service volumes persistent volume"], "anchor": "section", "description": "Volume containing the Persistent Storage for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "volumes", "persistent_volume"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/volumes/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Volumes for the service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/host_path/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/persistent_volume/): complete subsection reference.

## Next pages

- [service.volumes.empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/empty_dir/)
- [service.volumes.host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/host_path/)
- [service.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/persistent_volume/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
