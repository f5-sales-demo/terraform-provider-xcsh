---
page_title: "service.advertise_options.advertise_in_cluster.multi_ports.ports"
subcategory: "Container"
description: "Ports to advertise."
xcsh_docs: {"aliases": ["service advertise options advertise in cluster multi ports ports"], "body_bytes": 3601, "body_sha256": "sha256:85f938e5d445cd958a2b08c9988dbcca446aeee2128a98a0fbed01fb7f820db3", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "path": "documentation/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323", "registry_path": "docs/guides/data-sources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports"], "schema_version": 1, "sections": [{"aliases": ["service advertise options advertise in cluster multi ports ports info"], "anchor": "section", "description": "Port information.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "info"], "syntax": "attribute", "type": "object"}, {"aliases": ["service advertise options advertise in cluster multi ports ports name"], "anchor": "schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--name", "description": "Name of the Port.", "document_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Ports to advertise.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["workloadCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.multi_ports.ports

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/)
- [service.advertise_options.advertise_in_cluster.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="section"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

- [info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/info/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```
