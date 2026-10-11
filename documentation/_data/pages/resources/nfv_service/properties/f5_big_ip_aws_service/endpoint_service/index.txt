---
page_title: "f5_big_ip_aws_service.endpoint_service"
subcategory: ""
description: "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address."
xcsh_docs: {"aliases": ["f5 big ip aws service endpoint service"], "body_bytes": 4662, "body_sha256": "sha256:8deb955ce947acdf9a9f09ff57d83b267f74a7053f80b6236c92803d30114d03", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service endpoint service advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service advertise on slo ip external"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip_external"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service automatic vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "automatic_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service configured vip"], "anchor": "schema-f5_big_ip_aws_service--endpoint_service--configured_vip", "description": "Exclusive with Enter IP address for the default VIP.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "configured_vip"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service endpoint service custom tcp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_tcp_ports"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service custom udp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service default tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "default_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service disable advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "disable_advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service no tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service no udp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_udp_ports"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.endpoint_service

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- f5_big_ip_aws_service.endpoint_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
endpoint_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip/): complete subsection reference.

- [advertise_on_slo_ip_external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip_external/): complete subsection reference.

- [automatic_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/automatic_vip/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--endpoint_service--configured_vip"></a>

### configured_vip property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_tcp_ports/): complete subsection reference.

- [custom_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/): complete subsection reference.

- [default_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/default_tcp_ports/): complete subsection reference.

- [disable_advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/disable_advertise_on_slo_ip/): complete subsection reference.

- [http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/http_port/): complete subsection reference.

- [https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/https_port/): complete subsection reference.

- [no_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_tcp_ports/): complete subsection reference.

- [no_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_udp_ports/): complete subsection reference.
