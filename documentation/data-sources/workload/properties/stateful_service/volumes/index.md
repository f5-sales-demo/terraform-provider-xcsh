---
page_title: "stateful_service.volumes"
subcategory: "Container"
description: "Ephemeral volumes for the service."
xcsh_docs: {"aliases": ["stateful service volumes"], "body_bytes": 3600, "body_sha256": "sha256:cba7006462b033cb3c12f5b4f584c438aea640f307ee8b52a1b6996922b1a7ac", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:volumes:empty_dir", "xcsh-docs:data-sources:workload:properties:stateful_service:volumes:host_path"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "path": "documentation/data-sources/workload/properties/stateful_service/volumes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212", "registry_path": "docs/guides/data-sources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "volumes"], "schema_version": 1, "sections": [{"aliases": ["empty dir"], "anchor": "section", "description": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes:empty_dir", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "volumes", "empty_dir"], "syntax": "attribute", "type": "object"}, {"aliases": ["host path"], "anchor": "section", "description": "Volume containing a host mapped path into the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes:host_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "volumes", "host_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-stateful_service--volumes--name", "description": "Name of the volume.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:volumes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "volumes", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/volumes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Ephemeral volumes for the service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- stateful_service.volumes

<a id="section"></a>

Type: `"list"`. Computed.

Ephemeral Volumes. Ephemeral volumes for the service.

Upstream description:

Ephemeral volumes for the service.

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

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/host_path/): complete subsection reference.

<a id="schema-stateful_service--volumes--name"></a>

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

## Next pages

- [stateful_service.volumes.empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/empty_dir/)
- [stateful_service.volumes.host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/volumes/host_path/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
