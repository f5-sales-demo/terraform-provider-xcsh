---
page_title: "f5_big_ip_aws_service.endpoint_service.custom_udp_ports"
subcategory: ""
description: "List of port ranges."
xcsh_docs: {"aliases": ["f5 big ip aws service endpoint service custom udp ports"], "body_bytes": 2282, "body_sha256": "sha256:c4dc201f05eea8fd6a2de90fd0ae1bafa317c30b6e31091e590dd3207eae6630", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3121301132011211-1100200303020202-2312221303312103-3211230311330023-3210031301110131-0023030101213312-2102223221231013-1101322112212331", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service endpoint service custom udp ports ports"], "anchor": "schema-f5_big_ip_aws_service--endpoint_service--custom_udp_ports--ports", "description": "List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports", "ports"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "List of port ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.endpoint_service.custom_udp_ports

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- [f5_big_ip_aws_service.endpoint_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port Range List. List of port ranges.

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
custom_udp_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-f5_big_ip_aws_service--endpoint_service--custom_udp_ports--ports"></a>

### ports property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```
