---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 31106, "body_sha256": "sha256:c68a6c6f2eb25d35c6c523e48b506038e0c808c80f00aa5aea513f2ad8688f4c", "canonical_id": "xcsh-docs:resources:advertise_policy:reference", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:dualstack", "xcsh-docs:resources:advertise_policy:properties:ipv4", "xcsh-docs:resources:advertise_policy:properties:ipv6", "xcsh-docs:resources:advertise_policy:properties:public_ip", "xcsh-docs:resources:advertise_policy:properties:timeouts", "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "xcsh-docs:resources:advertise_policy:properties:where"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:reference", "parent_id": "xcsh-docs:resources:advertise_policy:fundamentals", "path": "docs/guides/resources--advertise_policy--reference.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Optional, Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Upstream description:

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
"where" is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site object will be used based on the network type. If inside\_vip/outside\_vip is not configured in
the site object, system use interface IP in the respected networks.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

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

- [dualstack](resources--advertise_policy--properties--dualstack.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](resources--advertise_policy--properties--ipv4.md): complete subsection reference.

- [ipv6](resources--advertise_policy--properties--ipv6.md): complete subsection reference.

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

Name of the Advertise Policy. Must be unique within the namespace.

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

Namespace where the Advertise Policy is created.

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

<a id="schema-port"></a>

### port property

Type: `"number"`. Optional, Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Upstream description:

Exclusive with \[port\_ranges\] Port to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [port](resources--advertise_policy--reference.md#schema-port)
- [port_ranges](resources--advertise_policy--reference.md#schema-port_ranges)

Select alternatives according to the provider validators above.

<a id="schema-port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional, Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="schema-protocol"></a>

### protocol property

Type: `"string"`. Optional, Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Protocol to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TCP",
    "UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](resources--advertise_policy--properties--public_ip.md): complete subsection reference.

<a id="schema-skip_xff_append"></a>

### skip_xff_append property

Type: `"bool"`. Optional, Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

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

- [timeouts](resources--advertise_policy--properties--timeouts.md): complete subsection reference.

- [tls_parameters](resources--advertise_policy--properties--tls_parameters.md): complete subsection reference.

- [where](resources--advertise_policy--properties--where.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--advertise_policy--reference.md#schema-address) |
| `annotations` | [annotations](resources--advertise_policy--reference.md#schema-annotations) |
| `description` | [description](resources--advertise_policy--reference.md#schema-description) |
| `disable` | [disable](resources--advertise_policy--reference.md#schema-disable) |
| `dualstack` | [dualstack](resources--advertise_policy--properties--dualstack.md#section) |
| `id` | [id](resources--advertise_policy--reference.md#schema-id) |
| `ipv4` | [ipv4](resources--advertise_policy--properties--ipv4.md#section) |
| `ipv6` | [ipv6](resources--advertise_policy--properties--ipv6.md#section) |
| `labels` | [labels](resources--advertise_policy--reference.md#schema-labels) |
| `name` | [name](resources--advertise_policy--reference.md#schema-name) |
| `namespace` | [namespace](resources--advertise_policy--reference.md#schema-namespace) |
| `port` | [port](resources--advertise_policy--reference.md#schema-port) |
| `port_ranges` | [port_ranges](resources--advertise_policy--reference.md#schema-port_ranges) |
| `protocol` | [protocol](resources--advertise_policy--reference.md#schema-protocol) |
| `public_ip` | [public_ip](resources--advertise_policy--properties--public_ip.md#section) |
| `public_ip.kind` | [public_ip.kind](resources--advertise_policy--properties--public_ip.md#schema-public_ip--kind) |
| `public_ip.name` | [public_ip.name](resources--advertise_policy--properties--public_ip.md#schema-public_ip--name) |
| `public_ip.namespace` | [public_ip.namespace](resources--advertise_policy--properties--public_ip.md#schema-public_ip--namespace) |
| `public_ip.tenant` | [public_ip.tenant](resources--advertise_policy--properties--public_ip.md#schema-public_ip--tenant) |
| `public_ip.uid` | [public_ip.uid](resources--advertise_policy--properties--public_ip.md#schema-public_ip--uid) |
| `skip_xff_append` | [skip_xff_append](resources--advertise_policy--reference.md#schema-skip_xff_append) |
| `timeouts` | [timeouts](resources--advertise_policy--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--advertise_policy--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--advertise_policy--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--advertise_policy--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--advertise_policy--properties--timeouts.md#schema-timeouts--update) |
| `tls_parameters` | [tls_parameters](resources--advertise_policy--properties--tls_parameters.md#section) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](resources--advertise_policy--properties--tls_parameters--client_certificate_optional.md#section) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](resources--advertise_policy--properties--tls_parameters--client_certificate_required.md#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](resources--advertise_policy--properties--tls_parameters--common_params.md#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](resources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](resources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](resources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](resources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](resources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](resources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](resources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](resources--advertise_policy--properties--tls_parameters--no_client_certificate.md#section) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](resources--advertise_policy--properties--tls_parameters.md#schema-tls_parameters--xfcc_header_elements) |
| `where` | [where](resources--advertise_policy--properties--where.md#section) |
| `where.site` | [where.site](resources--advertise_policy--properties--where--site.md#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--advertise_policy--properties--where--site--disable_internet_vip.md#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--advertise_policy--properties--where--site--enable_internet_vip.md#section) |
| `where.site.network_type` | [where.site.network_type](resources--advertise_policy--properties--where--site.md#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](resources--advertise_policy--properties--where--site--ref.md#section) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](resources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](resources--advertise_policy--properties--where--virtual_network.md#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--advertise_policy--properties--where--virtual_network--ref.md#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](resources--advertise_policy--properties--where--virtual_site.md#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--advertise_policy--properties--where--virtual_site--disable_internet_vip.md#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--advertise_policy--properties--where--virtual_site--enable_internet_vip.md#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--advertise_policy--properties--where--virtual_site.md#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--advertise_policy--properties--where--virtual_site--ref.md#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--uid) |

## Next pages

- [dualstack](resources--advertise_policy--properties--dualstack.md)
- [ipv4](resources--advertise_policy--properties--ipv4.md)
- [ipv6](resources--advertise_policy--properties--ipv6.md)
- [public_ip](resources--advertise_policy--properties--public_ip.md)
- [timeouts](resources--advertise_policy--properties--timeouts.md)
- [tls_parameters](resources--advertise_policy--properties--tls_parameters.md)
- [where](resources--advertise_policy--properties--where.md)
- [xcsh_advertise_policy](../resources/advertise_policy.md)
