---
page_title: "f5_big_ip_aws_service.endpoint_service"
subcategory: ""
description: "f5_big_ip_aws_service.endpoint_service for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 7667, "body_sha256": "sha256:6c583f3f016318a4a1586563441732fb7bc84857e4ad7066e29a9296210d8a60", "canonical_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "path": "docs/guides/resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.endpoint_service for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.endpoint_service

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- f5_big_ip_aws_service.endpoint_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Upstream description:

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_on_slo_ip",
    "advertise_on_slo_ip_external"),
  validators.ConflictingObjectAttributes("advertise_on_slo_ip",
    "disable_advertise_on_slo_ip"),
  validators.ConflictingObjectAttributes("advertise_on_slo_ip_external",
    "disable_advertise_on_slo_ip"),
  validators.ConflictingObjectAttributes("automatic_vip",
    "configured_vip"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "default_tcp_ports"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "http_port"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "https_port"),
  validators.ConflictingObjectAttributes("custom_tcp_ports",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("custom_udp_ports",
    "no_udp_ports"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "http_port"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "https_port"),
  validators.ConflictingObjectAttributes("default_tcp_ports",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("http_port",
    "https_port"),
  validators.ConflictingObjectAttributes("http_port",
    "no_tcp_ports"),
  validators.ConflictingObjectAttributes("https_port",
    "no_tcp_ports")}
```

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

- [advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip.md): complete subsection reference.

- [advertise_on_slo_ip_external](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip_external.md): complete subsection reference.

- [automatic_vip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--automatic_vip.md): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--endpoint_service--configured_vip"></a>

### configured_vip property

Type: `"string"`. Optional.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Upstream description:

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

- [custom_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_tcp_ports.md): complete subsection reference.

- [custom_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_udp_ports.md): complete subsection reference.

- [default_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--default_tcp_ports.md): complete subsection reference.

- [disable_advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--disable_advertise_on_slo_ip.md): complete subsection reference.

- [http_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--http_port.md): complete subsection reference.

- [https_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--https_port.md): complete subsection reference.

- [no_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_tcp_ports.md): complete subsection reference.

- [no_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_udp_ports.md): complete subsection reference.

## Next pages

- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip.md)
- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip_external.md)
- [f5_big_ip_aws_service.endpoint_service.automatic_vip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--automatic_vip.md)
- [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_tcp_ports.md)
- [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_udp_ports.md)
- [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--default_tcp_ports.md)
- [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--disable_advertise_on_slo_ip.md)
- [f5_big_ip_aws_service.endpoint_service.http_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--http_port.md)
- [f5_big_ip_aws_service.endpoint_service.https_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--https_port.md)
- [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_tcp_ports.md)
- [f5_big_ip_aws_service.endpoint_service.no_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_udp_ports.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
