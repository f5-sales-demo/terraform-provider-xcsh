---
page_title: "f5_big_ip_aws_service.endpoint_service"
subcategory: ""
description: "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address."
xcsh_docs: {"aliases": ["f5 big ip aws service endpoint service"], "body_bytes": 4579, "body_sha256": "sha256:69a53008585fe0d2a47cb35340e271c969605fd11dc14dce1d96aea882e255f1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service endpoint service advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service advertise on slo ip external"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip_external"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service automatic vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "automatic_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service configured vip"], "anchor": "schema-f5_big_ip_aws_service--endpoint_service--configured_vip", "description": "Exclusive with Enter IP address for the default VIP.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "configured_vip"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service endpoint service custom tcp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service custom udp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service default tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "default_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service disable advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "disable_advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service no tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service no udp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_udp_ports"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.endpoint_service

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.endpoint_service

<a id="section"></a>

Type: `"single"`. Computed.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-external_vip_choice": "[\"advertise_on_slo_ip\",\"advertise_on_slo_ip_external\",\"disable_advertise_on_slo_ip\"]",
  "x-ves-oneof-field-inside_vip_choice": "[\"automatic_vip\",\"configured_vip\"]",
  "x-ves-oneof-field-tcp_port_choice": "[\"custom_tcp_ports\",\"default_tcp_ports\",\"http_port\",\"https_port\",\"no_tcp_ports\"]",
  "x-ves-oneof-field-udp_port_choice": "[\"custom_udp_ports\",\"no_udp_ports\"]"
}
```

## Direct properties

- [advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip/): complete subsection reference.

- [advertise_on_slo_ip_external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip_external/): complete subsection reference.

- [automatic_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/automatic_vip/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--endpoint_service--configured_vip"></a>

### configured_vip property

Type: `"string"`. Computed.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  }
}
```

- [custom_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_tcp_ports/): complete subsection reference.

- [custom_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/): complete subsection reference.

- [default_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/default_tcp_ports/): complete subsection reference.

- [disable_advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/disable_advertise_on_slo_ip/): complete subsection reference.

- [http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/http_port/): complete subsection reference.

- [https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/https_port/): complete subsection reference.

- [no_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_tcp_ports/): complete subsection reference.

- [no_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_udp_ports/): complete subsection reference.
