---
page_title: "stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer"
subcategory: "Container"
description: "TCP loadbalancer."
xcsh_docs: {"aliases": ["stateful service advertise options advertise custom ports tcp loadbalancer"], "body_bytes": 3562, "body_sha256": "sha256:50ef60d2af994c8178d0d49e768f1c60a14fa16e4763cc45bf2d366413372642", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:tcp_loadbalancer", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/tcp_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0020101203031230-1223300211330033-2123113022210223-3310333011133100-3330221301313310-3131021233011311-3221200100200012-0300221100330331", "registry_path": "docs/guides/resources--workload--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "tcp_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise custom ports tcp loadbalancer domains"], "anchor": "schema-stateful_service--advertise_options--advertise_custom--ports--tcp_loadbalancer--domains", "description": "A list of additional domains (host/authority header) that will be matched to this loadbalancer. Domains are also used for SNI matching if the `with_sni` is true Domains also indicate the list of names for which DNS resolution will be done by VER.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:tcp_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "tcp_loadbalancer", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["stateful service advertise options advertise custom ports tcp loadbalancer with sni"], "anchor": "schema-stateful_service--advertise_options--advertise_custom--ports--tcp_loadbalancer--with_sni", "description": "Set to true to enable TCP loadbalancer with SNI.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:tcp_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "tcp_loadbalancer", "with_sni"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/tcp_loadbalancer/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "TCP loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["workloadCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/)
- stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

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

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--tcp_loadbalancer--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Additional upstream details:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--tcp_loadbalancer--with_sni"></a>

### with_sni property

Type: `"bool"`. Optional.

Set to true to enable TCP loadbalancer with SNI.

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
