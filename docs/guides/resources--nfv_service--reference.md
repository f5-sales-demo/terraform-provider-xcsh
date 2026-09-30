---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 88080, "body_sha256": "sha256:1ab13d14ca6b3ffcaea7995b4af1443ee31c7a512f0772841ae11b64f36a50ee", "canonical_id": "xcsh-docs:resources:nfv_service:reference", "child_ids": ["xcsh-docs:resources:nfv_service:properties:disable_https_management", "xcsh-docs:resources:nfv_service:properties:disable_ssh_access", "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "xcsh-docs:resources:nfv_service:properties:https_management", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "xcsh-docs:resources:nfv_service:properties:timeouts"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:reference", "parent_id": "xcsh-docs:resources:nfv_service:fundamentals", "path": "docs/guides/resources--nfv_service--reference.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [disable_https_management](resources--nfv_service--properties--disable_https_management.md): complete subsection reference.

- [disable_ssh_access](resources--nfv_service--properties--disable_ssh_access.md): complete subsection reference.

- [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md): complete subsection reference.

- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md): complete subsection reference.

- [https_management](resources--nfv_service--properties--https_management.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Nfv Service. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Nfv Service is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
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
  }
}
```

- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md): complete subsection reference.

- [timeouts](resources--nfv_service--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nfv_service--reference.md#schema-annotations) |
| `description` | [description](resources--nfv_service--reference.md#schema-description) |
| `disable` | [disable](resources--nfv_service--reference.md#schema-disable) |
| `disable_https_management` | [disable_https_management](resources--nfv_service--properties--disable_https_management.md#section) |
| `disable_ssh_access` | [disable_ssh_access](resources--nfv_service--properties--disable_ssh_access.md#section) |
| `enabled_ssh_access` | [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md#section) |
| `enabled_ssh_access.advertise_on_sli` | [enabled_ssh_access.advertise_on_sli](resources--nfv_service--properties--enabled_ssh_access--advertise_on_sli.md#section) |
| `enabled_ssh_access.advertise_on_slo` | [enabled_ssh_access.advertise_on_slo](resources--nfv_service--properties--enabled_ssh_access--advertise_on_slo.md#section) |
| `enabled_ssh_access.advertise_on_slo_sli` | [enabled_ssh_access.advertise_on_slo_sli](resources--nfv_service--properties--enabled_ssh_access--advertise_on_slo_sli.md#section) |
| `enabled_ssh_access.domain_suffix` | [enabled_ssh_access.domain_suffix](resources--nfv_service--properties--enabled_ssh_access.md#schema-enabled_ssh_access--domain_suffix) |
| `enabled_ssh_access.node_ssh_ports` | [enabled_ssh_access.node_ssh_ports](resources--nfv_service--properties--enabled_ssh_access--node_ssh_ports.md#section) |
| `enabled_ssh_access.node_ssh_ports.node_name` | [enabled_ssh_access.node_ssh_ports.node_name](resources--nfv_service--properties--enabled_ssh_access--node_ssh_ports.md#schema-enabled_ssh_access--node_ssh_ports--node_name) |
| `enabled_ssh_access.node_ssh_ports.ssh_port` | [enabled_ssh_access.node_ssh_ports.ssh_port](resources--nfv_service--properties--enabled_ssh_access--node_ssh_ports.md#schema-enabled_ssh_access--node_ssh_ports--ssh_port) |
| `f5_big_ip_aws_service` | [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md#section) |
| `f5_big_ip_aws_service.admin_password` | [f5_big_ip_aws_service.admin_password](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password.md#section) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md#section) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md#schema-f5_big_ip_aws_service--admin_password--blindfold_secret_info--decryption_provider) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.location](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md#schema-f5_big_ip_aws_service--admin_password--blindfold_secret_info--location) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--blindfold_secret_info.md#schema-f5_big_ip_aws_service--admin_password--blindfold_secret_info--store_provider) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info` | [f5_big_ip_aws_service.admin_password.clear_secret_info](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--clear_secret_info.md#section) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` | [f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--clear_secret_info.md#schema-f5_big_ip_aws_service--admin_password--clear_secret_info--provider_ref) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.url` | [f5_big_ip_aws_service.admin_password.clear_secret_info.url](resources--nfv_service--properties--f5_big_ip_aws_service--admin_password--clear_secret_info.md#schema-f5_big_ip_aws_service--admin_password--clear_secret_info--url) |
| `f5_big_ip_aws_service.admin_username` | [f5_big_ip_aws_service.admin_username](resources--nfv_service--properties--f5_big_ip_aws_service.md#schema-f5_big_ip_aws_service--admin_username) |
| `f5_big_ip_aws_service.aws_tgw_site_params` | [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params.md#section) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site.md#section) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site.md#schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--name) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site.md#schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--namespace) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant](resources--nfv_service--properties--f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site.md#schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--tenant) |
| `f5_big_ip_aws_service.endpoint_service` | [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md#section) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip.md#section) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--advertise_on_slo_ip_external.md#section) |
| `f5_big_ip_aws_service.endpoint_service.automatic_vip` | [f5_big_ip_aws_service.endpoint_service.automatic_vip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--automatic_vip.md#section) |
| `f5_big_ip_aws_service.endpoint_service.configured_vip` | [f5_big_ip_aws_service.endpoint_service.configured_vip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md#schema-f5_big_ip_aws_service--endpoint_service--configured_vip) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_tcp_ports.md#section) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_tcp_ports.md#schema-f5_big_ip_aws_service--endpoint_service--custom_tcp_ports--ports) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_udp_ports.md#section) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_udp_ports.md#schema-f5_big_ip_aws_service--endpoint_service--custom_udp_ports--ports) |
| `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--default_tcp_ports.md#section) |
| `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--disable_advertise_on_slo_ip.md#section) |
| `f5_big_ip_aws_service.endpoint_service.http_port` | [f5_big_ip_aws_service.endpoint_service.http_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--http_port.md#section) |
| `f5_big_ip_aws_service.endpoint_service.https_port` | [f5_big_ip_aws_service.endpoint_service.https_port](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--https_port.md#section) |
| `f5_big_ip_aws_service.endpoint_service.no_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_tcp_ports.md#section) |
| `f5_big_ip_aws_service.endpoint_service.no_udp_ports` | [f5_big_ip_aws_service.endpoint_service.no_udp_ports](resources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--no_udp_ports.md#section) |
| `f5_big_ip_aws_service.market_place_image` | [f5_big_ip_aws_service.market_place_image](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image.md#section) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g200_mbps.md#section) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](resources--nfv_service--properties--f5_big_ip_aws_service--market_place_image--awafpay_g3_gbps.md#section) |
| `f5_big_ip_aws_service.nodes` | [f5_big_ip_aws_service.nodes](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md#section) |
| `f5_big_ip_aws_service.nodes.automatic_prefix` | [f5_big_ip_aws_service.nodes.automatic_prefix](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--automatic_prefix.md#section) |
| `f5_big_ip_aws_service.nodes.aws_az_name` | [f5_big_ip_aws_service.nodes.aws_az_name](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md#schema-f5_big_ip_aws_service--nodes--aws_az_name) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet` | [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md#section) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id` | [f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet.md#schema-f5_big_ip_aws_service--nodes--mgmt_subnet--existing_subnet_id) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md#section) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param.md#schema-f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param--ipv4) |
| `f5_big_ip_aws_service.nodes.node_name` | [f5_big_ip_aws_service.nodes.node_name](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md#schema-f5_big_ip_aws_service--nodes--node_name) |
| `f5_big_ip_aws_service.nodes.reserved_mgmt_subnet` | [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](resources--nfv_service--properties--f5_big_ip_aws_service--nodes--reserved_mgmt_subnet.md#section) |
| `f5_big_ip_aws_service.nodes.tunnel_prefix` | [f5_big_ip_aws_service.nodes.tunnel_prefix](resources--nfv_service--properties--f5_big_ip_aws_service--nodes.md#schema-f5_big_ip_aws_service--nodes--tunnel_prefix) |
| `f5_big_ip_aws_service.ssh_key` | [f5_big_ip_aws_service.ssh_key](resources--nfv_service--properties--f5_big_ip_aws_service.md#schema-f5_big_ip_aws_service--ssh_key) |
| `f5_big_ip_aws_service.tags` | [f5_big_ip_aws_service.tags](resources--nfv_service--properties--f5_big_ip_aws_service.md#schema-f5_big_ip_aws_service--tags) |
| `https_management` | [https_management](resources--nfv_service--properties--https_management.md#section) |
| `https_management.advertise_on_internet` | [https_management.advertise_on_internet](resources--nfv_service--properties--https_management--advertise_on_internet.md#section) |
| `https_management.advertise_on_internet.public_ip` | [https_management.advertise_on_internet.public_ip](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md#section) |
| `https_management.advertise_on_internet.public_ip.name` | [https_management.advertise_on_internet.public_ip.name](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md#schema-https_management--advertise_on_internet--public_ip--name) |
| `https_management.advertise_on_internet.public_ip.namespace` | [https_management.advertise_on_internet.public_ip.namespace](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md#schema-https_management--advertise_on_internet--public_ip--namespace) |
| `https_management.advertise_on_internet.public_ip.tenant` | [https_management.advertise_on_internet.public_ip.tenant](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md#schema-https_management--advertise_on_internet--public_ip--tenant) |
| `https_management.advertise_on_internet_default_vip` | [https_management.advertise_on_internet_default_vip](resources--nfv_service--properties--https_management--advertise_on_internet_default_vip.md#section) |
| `https_management.advertise_on_sli_vip` | [https_management.advertise_on_sli_vip](resources--nfv_service--properties--https_management--advertise_on_sli_vip.md#section) |
| `https_management.advertise_on_sli_vip.no_mtls` | [https_management.advertise_on_sli_vip.no_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--no_mtls.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates` | [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` | [https_management.advertise_on_sli_vip.tls_certificates.certificate_url](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md#schema-https_management--advertise_on_sli_vip--tls_certificates--certificate_url) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--custom_hash_algorithms.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--custom_hash_algorithms.md#schema-https_management--advertise_on_sli_vip--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `https_management.advertise_on_sli_vip.tls_certificates.description_spec` | [https_management.advertise_on_sli_vip.tls_certificates.description_spec](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md#schema-https_management--advertise_on_sli_vip--tls_certificates--description_spec) |
| `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--disable_ocsp_stapling.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key` | [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info--location) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_sli_vip--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info.md#section) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_sli_vip--tls_certificates--private_key--clear_secret_info--url) |
| `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates--use_system_defaults.md#section) |
| `https_management.advertise_on_sli_vip.tls_config` | [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config.md#section) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security` | [https_management.advertise_on_sli_vip.tls_config.custom_security](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md#section) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md#schema-https_management--advertise_on_sli_vip--tls_config--custom_security--cipher_suites) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.max_version](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md#schema-https_management--advertise_on_sli_vip--tls_config--custom_security--max_version) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.min_version](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md#schema-https_management--advertise_on_sli_vip--tls_config--custom_security--min_version) |
| `https_management.advertise_on_sli_vip.tls_config.default_security` | [https_management.advertise_on_sli_vip.tls_config.default_security](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--default_security.md#section) |
| `https_management.advertise_on_sli_vip.tls_config.low_security` | [https_management.advertise_on_sli_vip.tls_config.low_security](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--low_security.md#section) |
| `https_management.advertise_on_sli_vip.tls_config.medium_security` | [https_management.advertise_on_sli_vip.tls_config.medium_security](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--medium_security.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls` | [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls.md#schema-https_management--advertise_on_sli_vip--use_mtls--client_certificate_optional) |
| `https_management.advertise_on_sli_vip.use_mtls.crl` | [https_management.advertise_on_sli_vip.use_mtls.crl](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--crl.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.name` | [https_management.advertise_on_sli_vip.use_mtls.crl.name](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--crl.md#schema-https_management--advertise_on_sli_vip--use_mtls--crl--name) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` | [https_management.advertise_on_sli_vip.use_mtls.crl.namespace](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--crl.md#schema-https_management--advertise_on_sli_vip--use_mtls--crl--namespace) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` | [https_management.advertise_on_sli_vip.use_mtls.crl.tenant](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--crl.md#schema-https_management--advertise_on_sli_vip--use_mtls--crl--tenant) |
| `https_management.advertise_on_sli_vip.use_mtls.no_crl` | [https_management.advertise_on_sli_vip.use_mtls.no_crl](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--no_crl.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--trusted_ca.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_sli_vip--use_mtls--trusted_ca--name) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_sli_vip--use_mtls--trusted_ca--namespace) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_sli_vip--use_mtls--trusted_ca--tenant) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls.md#schema-https_management--advertise_on_sli_vip--use_mtls--trusted_ca_url) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--xfcc_disabled.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--xfcc_options.md#section) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls--xfcc_options.md#schema-https_management--advertise_on_sli_vip--use_mtls--xfcc_options--xfcc_header_elements) |
| `https_management.advertise_on_slo_internet_vip` | [https_management.advertise_on_slo_internet_vip](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip.md#section) |
| `https_management.advertise_on_slo_internet_vip.no_mtls` | [https_management.advertise_on_slo_internet_vip.no_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--no_mtls.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates` | [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--certificate_url) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--description_spec) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--disable_ocsp_stapling.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info--location) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_internet_vip--tls_certificates--private_key--clear_secret_info--url) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates--use_system_defaults.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_config` | [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--custom_security.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_internet_vip--tls_config--custom_security--cipher_suites) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_internet_vip--tls_config--custom_security--max_version) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_internet_vip--tls_config--custom_security--min_version) |
| `https_management.advertise_on_slo_internet_vip.tls_config.default_security` | [https_management.advertise_on_slo_internet_vip.tls_config.default_security](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--default_security.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_config.low_security` | [https_management.advertise_on_slo_internet_vip.tls_config.low_security](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--low_security.md#section) |
| `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` | [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config--medium_security.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls` | [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--client_certificate_optional) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--crl.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.name](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--crl--name) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--crl--namespace) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--crl--tenant) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--no_crl.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca--name) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca--namespace) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca--tenant) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--trusted_ca_url) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--xfcc_disabled.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--xfcc_options.md#section) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls--xfcc_options.md#schema-https_management--advertise_on_slo_internet_vip--use_mtls--xfcc_options--xfcc_header_elements) |
| `https_management.advertise_on_slo_sli` | [https_management.advertise_on_slo_sli](resources--nfv_service--properties--https_management--advertise_on_slo_sli.md#section) |
| `https_management.advertise_on_slo_sli.no_mtls` | [https_management.advertise_on_slo_sli.no_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_sli--no_mtls.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates` | [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` | [https_management.advertise_on_slo_sli.tls_certificates.certificate_url](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates.md#schema-https_management--advertise_on_slo_sli--tls_certificates--certificate_url) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--custom_hash_algorithms.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--custom_hash_algorithms.md#schema-https_management--advertise_on_slo_sli--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `https_management.advertise_on_slo_sli.tls_certificates.description_spec` | [https_management.advertise_on_slo_sli.tls_certificates.description_spec](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates.md#schema-https_management--advertise_on_slo_sli--tls_certificates--description_spec) |
| `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--disable_ocsp_stapling.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key` | [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--location) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info.md#section) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--url) |
| `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--use_system_defaults.md#section) |
| `https_management.advertise_on_slo_sli.tls_config` | [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config.md#section) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security` | [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md#section) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md#schema-https_management--advertise_on_slo_sli--tls_config--custom_security--cipher_suites) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.max_version](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md#schema-https_management--advertise_on_slo_sli--tls_config--custom_security--max_version) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.min_version](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md#schema-https_management--advertise_on_slo_sli--tls_config--custom_security--min_version) |
| `https_management.advertise_on_slo_sli.tls_config.default_security` | [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--default_security.md#section) |
| `https_management.advertise_on_slo_sli.tls_config.low_security` | [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--low_security.md#section) |
| `https_management.advertise_on_slo_sli.tls_config.medium_security` | [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--medium_security.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls` | [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls.md#schema-https_management--advertise_on_slo_sli--use_mtls--client_certificate_optional) |
| `https_management.advertise_on_slo_sli.use_mtls.crl` | [https_management.advertise_on_slo_sli.use_mtls.crl](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--crl.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.name` | [https_management.advertise_on_slo_sli.use_mtls.crl.name](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--crl.md#schema-https_management--advertise_on_slo_sli--use_mtls--crl--name) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` | [https_management.advertise_on_slo_sli.use_mtls.crl.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--crl.md#schema-https_management--advertise_on_slo_sli--use_mtls--crl--namespace) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` | [https_management.advertise_on_slo_sli.use_mtls.crl.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--crl.md#schema-https_management--advertise_on_slo_sli--use_mtls--crl--tenant) |
| `https_management.advertise_on_slo_sli.use_mtls.no_crl` | [https_management.advertise_on_slo_sli.use_mtls.no_crl](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--no_crl.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--trusted_ca.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_sli--use_mtls--trusted_ca--name) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_sli--use_mtls--trusted_ca--namespace) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_sli--use_mtls--trusted_ca--tenant) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls.md#schema-https_management--advertise_on_slo_sli--use_mtls--trusted_ca_url) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--xfcc_disabled.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--xfcc_options.md#section) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--properties--https_management--advertise_on_slo_sli--use_mtls--xfcc_options.md#schema-https_management--advertise_on_slo_sli--use_mtls--xfcc_options--xfcc_header_elements) |
| `https_management.advertise_on_slo_vip` | [https_management.advertise_on_slo_vip](resources--nfv_service--properties--https_management--advertise_on_slo_vip.md#section) |
| `https_management.advertise_on_slo_vip.no_mtls` | [https_management.advertise_on_slo_vip.no_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_vip--no_mtls.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates` | [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_vip.tls_certificates.certificate_url](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates.md#schema-https_management--advertise_on_slo_vip--tls_certificates--certificate_url) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--custom_hash_algorithms.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--custom_hash_algorithms.md#schema-https_management--advertise_on_slo_vip--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `https_management.advertise_on_slo_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_vip.tls_certificates.description_spec](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates.md#schema-https_management--advertise_on_slo_vip--tls_certificates--description_spec) |
| `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--disable_ocsp_stapling.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info--location) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info.md#schema-https_management--advertise_on_slo_vip--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--clear_secret_info.md#section) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_vip--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key--clear_secret_info.md#schema-https_management--advertise_on_slo_vip--tls_certificates--private_key--clear_secret_info--url) |
| `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--use_system_defaults.md#section) |
| `https_management.advertise_on_slo_vip.tls_config` | [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config.md#section) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security` | [https_management.advertise_on_slo_vip.tls_config.custom_security](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--custom_security.md#section) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_vip--tls_config--custom_security--cipher_suites) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.max_version](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_vip--tls_config--custom_security--max_version) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.min_version](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--custom_security.md#schema-https_management--advertise_on_slo_vip--tls_config--custom_security--min_version) |
| `https_management.advertise_on_slo_vip.tls_config.default_security` | [https_management.advertise_on_slo_vip.tls_config.default_security](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--default_security.md#section) |
| `https_management.advertise_on_slo_vip.tls_config.low_security` | [https_management.advertise_on_slo_vip.tls_config.low_security](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--low_security.md#section) |
| `https_management.advertise_on_slo_vip.tls_config.medium_security` | [https_management.advertise_on_slo_vip.tls_config.medium_security](resources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_config--medium_security.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls` | [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls.md#schema-https_management--advertise_on_slo_vip--use_mtls--client_certificate_optional) |
| `https_management.advertise_on_slo_vip.use_mtls.crl` | [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_vip.use_mtls.crl.name](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_vip--use_mtls--crl--name) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_vip.use_mtls.crl.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_vip--use_mtls--crl--namespace) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_vip.use_mtls.crl.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--crl.md#schema-https_management--advertise_on_slo_vip--use_mtls--crl--tenant) |
| `https_management.advertise_on_slo_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--no_crl.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_vip--use_mtls--trusted_ca--name) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_vip--use_mtls--trusted_ca--namespace) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--trusted_ca.md#schema-https_management--advertise_on_slo_vip--use_mtls--trusted_ca--tenant) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls.md#schema-https_management--advertise_on_slo_vip--use_mtls--trusted_ca_url) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_disabled.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_options.md#section) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--properties--https_management--advertise_on_slo_vip--use_mtls--xfcc_options.md#schema-https_management--advertise_on_slo_vip--use_mtls--xfcc_options--xfcc_header_elements) |
| `https_management.default_https_port` | [https_management.default_https_port](resources--nfv_service--properties--https_management--default_https_port.md#section) |
| `https_management.domain_suffix` | [https_management.domain_suffix](resources--nfv_service--properties--https_management.md#schema-https_management--domain_suffix) |
| `https_management.https_port` | [https_management.https_port](resources--nfv_service--properties--https_management.md#schema-https_management--https_port) |
| `id` | [id](resources--nfv_service--reference.md#schema-id) |
| `labels` | [labels](resources--nfv_service--reference.md#schema-labels) |
| `name` | [name](resources--nfv_service--reference.md#schema-name) |
| `namespace` | [namespace](resources--nfv_service--reference.md#schema-namespace) |
| `palo_alto_fw_service` | [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md#section) |
| `palo_alto_fw_service.auto_setup` | [palo_alto_fw_service.auto_setup](resources--nfv_service--properties--palo_alto_fw_service--auto_setup.md#section) |
| `palo_alto_fw_service.auto_setup.admin_password` | [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password.md#section) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md#section) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info--decryption_provider) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info--location) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--admin_password--blindfold_secret_info--store_provider) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md#section) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md#schema-palo_alto_fw_service--auto_setup--admin_password--clear_secret_info--provider_ref) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--admin_password--clear_secret_info.md#schema-palo_alto_fw_service--auto_setup--admin_password--clear_secret_info--url) |
| `palo_alto_fw_service.auto_setup.admin_username` | [palo_alto_fw_service.auto_setup.admin_username](resources--nfv_service--properties--palo_alto_fw_service--auto_setup.md#schema-palo_alto_fw_service--auto_setup--admin_username) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys` | [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys.md#section) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key.md#section) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info.md#section) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info--decryption_provider) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info--location) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--blindfold_secret_info--store_provider) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--clear_secret_info.md#section) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--clear_secret_info.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--clear_secret_info--provider_ref) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--clear_secret_info.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--private_key--clear_secret_info--url) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key](resources--nfv_service--properties--palo_alto_fw_service--auto_setup--manual_ssh_keys.md#schema-palo_alto_fw_service--auto_setup--manual_ssh_keys--public_key) |
| `palo_alto_fw_service.aws_tgw_site` | [palo_alto_fw_service.aws_tgw_site](resources--nfv_service--properties--palo_alto_fw_service--aws_tgw_site.md#section) |
| `palo_alto_fw_service.aws_tgw_site.name` | [palo_alto_fw_service.aws_tgw_site.name](resources--nfv_service--properties--palo_alto_fw_service--aws_tgw_site.md#schema-palo_alto_fw_service--aws_tgw_site--name) |
| `palo_alto_fw_service.aws_tgw_site.namespace` | [palo_alto_fw_service.aws_tgw_site.namespace](resources--nfv_service--properties--palo_alto_fw_service--aws_tgw_site.md#schema-palo_alto_fw_service--aws_tgw_site--namespace) |
| `palo_alto_fw_service.aws_tgw_site.tenant` | [palo_alto_fw_service.aws_tgw_site.tenant](resources--nfv_service--properties--palo_alto_fw_service--aws_tgw_site.md#schema-palo_alto_fw_service--aws_tgw_site--tenant) |
| `palo_alto_fw_service.disable_panaroma` | [palo_alto_fw_service.disable_panaroma](resources--nfv_service--properties--palo_alto_fw_service--disable_panaroma.md#section) |
| `palo_alto_fw_service.instance_type` | [palo_alto_fw_service.instance_type](resources--nfv_service--properties--palo_alto_fw_service.md#schema-palo_alto_fw_service--instance_type) |
| `palo_alto_fw_service.pan_ami_bundle1` | [palo_alto_fw_service.pan_ami_bundle1](resources--nfv_service--properties--palo_alto_fw_service--pan_ami_bundle1.md#section) |
| `palo_alto_fw_service.pan_ami_bundle2` | [palo_alto_fw_service.pan_ami_bundle2](resources--nfv_service--properties--palo_alto_fw_service--pan_ami_bundle2.md#section) |
| `palo_alto_fw_service.panorama_server` | [palo_alto_fw_service.panorama_server](resources--nfv_service--properties--palo_alto_fw_service--panorama_server.md#section) |
| `palo_alto_fw_service.panorama_server.authorization_key` | [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key.md#section) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info.md#section) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info.md#schema-palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info--decryption_provider) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info.md#schema-palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info--location) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info.md#schema-palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info--store_provider) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info.md#section) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info.md#schema-palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info--provider_ref) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info.md#schema-palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info--url) |
| `palo_alto_fw_service.panorama_server.device_group_name` | [palo_alto_fw_service.panorama_server.device_group_name](resources--nfv_service--properties--palo_alto_fw_service--panorama_server.md#schema-palo_alto_fw_service--panorama_server--device_group_name) |
| `palo_alto_fw_service.panorama_server.server` | [palo_alto_fw_service.panorama_server.server](resources--nfv_service--properties--palo_alto_fw_service--panorama_server.md#schema-palo_alto_fw_service--panorama_server--server) |
| `palo_alto_fw_service.panorama_server.template_stack_name` | [palo_alto_fw_service.panorama_server.template_stack_name](resources--nfv_service--properties--palo_alto_fw_service--panorama_server.md#schema-palo_alto_fw_service--panorama_server--template_stack_name) |
| `palo_alto_fw_service.service_nodes` | [palo_alto_fw_service.service_nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes.md#section) |
| `palo_alto_fw_service.service_nodes.nodes` | [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md#section) |
| `palo_alto_fw_service.service_nodes.nodes.aws_az_name` | [palo_alto_fw_service.service_nodes.nodes.aws_az_name](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md#schema-palo_alto_fw_service--service_nodes--nodes--aws_az_name) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet.md#section) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet.md#schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--existing_subnet_id) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--subnet_param.md#section) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--subnet_param.md#schema-palo_alto_fw_service--service_nodes--nodes--mgmt_subnet--subnet_param--ipv4) |
| `palo_alto_fw_service.service_nodes.nodes.node_name` | [palo_alto_fw_service.service_nodes.nodes.node_name](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md#schema-palo_alto_fw_service--service_nodes--nodes--node_name) |
| `palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes--reserved_mgmt_subnet.md#section) |
| `palo_alto_fw_service.ssh_key` | [palo_alto_fw_service.ssh_key](resources--nfv_service--properties--palo_alto_fw_service.md#schema-palo_alto_fw_service--ssh_key) |
| `palo_alto_fw_service.tags` | [palo_alto_fw_service.tags](resources--nfv_service--properties--palo_alto_fw_service.md#schema-palo_alto_fw_service--tags) |
| `palo_alto_fw_service.version` | [palo_alto_fw_service.version](resources--nfv_service--properties--palo_alto_fw_service.md#schema-palo_alto_fw_service--version) |
| `timeouts` | [timeouts](resources--nfv_service--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--nfv_service--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--nfv_service--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--nfv_service--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--nfv_service--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [disable_https_management](resources--nfv_service--properties--disable_https_management.md)
- [disable_ssh_access](resources--nfv_service--properties--disable_ssh_access.md)
- [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md)
- [f5_big_ip_aws_service](resources--nfv_service--properties--f5_big_ip_aws_service.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [timeouts](resources--nfv_service--properties--timeouts.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
