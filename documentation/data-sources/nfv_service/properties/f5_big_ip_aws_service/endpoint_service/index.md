---
page_title: "f5_big_ip_aws_service.endpoint_service"
subcategory: ""
description: "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address."
xcsh_docs: {"aliases": ["f5 big ip aws service endpoint service"], "body_bytes": 7426, "body_sha256": "sha256:d4d8fb3a242719764de04852de248b72b6f46dd0e61496166b48a380e8a40b83", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service"], "schema_version": 1, "sections": [{"aliases": ["advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise on slo ip external"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "advertise_on_slo_ip_external"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "automatic_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["configured vip"], "anchor": "schema-f5_big_ip_aws_service--endpoint_service--configured_vip", "description": "Exclusive with Enter IP address for the default VIP.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "configured_vip"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom tcp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom udp ports"], "anchor": "section", "description": "List of port ranges.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["default tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "default_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable advertise on slo ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "disable_advertise_on_slo_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["no tcp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_tcp_ports"], "syntax": "attribute", "type": "object"}, {"aliases": ["no udp ports"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "no_udp_ports"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

Upstream description:

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

## Next pages

- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip/)
- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/advertise_on_slo_ip_external/)
- [f5_big_ip_aws_service.endpoint_service.automatic_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/automatic_vip/)
- [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_tcp_ports/)
- [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/)
- [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/default_tcp_ports/)
- [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/disable_advertise_on_slo_ip/)
- [f5_big_ip_aws_service.endpoint_service.http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/http_port/)
- [f5_big_ip_aws_service.endpoint_service.https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/https_port/)
- [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_tcp_ports/)
- [f5_big_ip_aws_service.endpoint_service.no_udp_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/no_udp_ports/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
