---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 29131, "body_sha256": "sha256:31f6f87d7ce4a5b35658e979805015339befa8115259138663c57baa43bbd685", "canonical_id": "xcsh-docs:data-sources:advertise_policy:reference", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:dualstack", "xcsh-docs:data-sources:advertise_policy:properties:ipv4", "xcsh-docs:data-sources:advertise_policy:properties:ipv6", "xcsh-docs:data-sources:advertise_policy:properties:public_ip", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters", "xcsh-docs:data-sources:advertise_policy:properties:where"], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:reference", "parent_id": "xcsh-docs:data-sources:advertise_policy:fundamentals", "path": "docs/guides/data-sources--advertise_policy--reference.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

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

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

Type: `"string"`. Computed.

Description of the AdvertisePolicy.

Upstream description:

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

- [dualstack](data-sources--advertise_policy--properties--dualstack.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4](data-sources--advertise_policy--properties--ipv4.md): complete subsection reference.

- [ipv6](data-sources--advertise_policy--properties--ipv6.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the AdvertisePolicy.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the AdvertisePolicy exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

Type: `"number"`. Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

Upstream description:

Exclusive with \[port\_ranges\] Port to advertise.

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

- [port](data-sources--advertise_policy--reference.md#schema-port)
- [port_ranges](data-sources--advertise_policy--reference.md#schema-port_ranges)

Select alternatives according to the provider validators above.

<a id="schema-port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Protocol to advertise.

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

- [public_ip](data-sources--advertise_policy--properties--public_ip.md): complete subsection reference.

<a id="schema-skip_xff_append"></a>

### skip_xff_append property

Type: `"bool"`. Computed.

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

- [tls_parameters](data-sources--advertise_policy--properties--tls_parameters.md): complete subsection reference.

- [where](data-sources--advertise_policy--properties--where.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--advertise_policy--reference.md#schema-address) |
| `annotations` | [annotations](data-sources--advertise_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--advertise_policy--reference.md#schema-description) |
| `dualstack` | [dualstack](data-sources--advertise_policy--properties--dualstack.md#section) |
| `id` | [id](data-sources--advertise_policy--reference.md#schema-id) |
| `ipv4` | [ipv4](data-sources--advertise_policy--properties--ipv4.md#section) |
| `ipv6` | [ipv6](data-sources--advertise_policy--properties--ipv6.md#section) |
| `labels` | [labels](data-sources--advertise_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--advertise_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--advertise_policy--reference.md#schema-namespace) |
| `port` | [port](data-sources--advertise_policy--reference.md#schema-port) |
| `port_ranges` | [port_ranges](data-sources--advertise_policy--reference.md#schema-port_ranges) |
| `protocol` | [protocol](data-sources--advertise_policy--reference.md#schema-protocol) |
| `public_ip` | [public_ip](data-sources--advertise_policy--properties--public_ip.md#section) |
| `public_ip.kind` | [public_ip.kind](data-sources--advertise_policy--properties--public_ip.md#schema-public_ip--kind) |
| `public_ip.name` | [public_ip.name](data-sources--advertise_policy--properties--public_ip.md#schema-public_ip--name) |
| `public_ip.namespace` | [public_ip.namespace](data-sources--advertise_policy--properties--public_ip.md#schema-public_ip--namespace) |
| `public_ip.tenant` | [public_ip.tenant](data-sources--advertise_policy--properties--public_ip.md#schema-public_ip--tenant) |
| `public_ip.uid` | [public_ip.uid](data-sources--advertise_policy--properties--public_ip.md#schema-public_ip--uid) |
| `skip_xff_append` | [skip_xff_append](data-sources--advertise_policy--reference.md#schema-skip_xff_append) |
| `tls_parameters` | [tls_parameters](data-sources--advertise_policy--properties--tls_parameters.md#section) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](data-sources--advertise_policy--properties--tls_parameters--client_certificate_optional.md#section) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](data-sources--advertise_policy--properties--tls_parameters--client_certificate_required.md#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--advertise_policy--properties--tls_parameters--common_params.md#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--advertise_policy--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--advertise_policy--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--advertise_policy--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](data-sources--advertise_policy--properties--tls_parameters--no_client_certificate.md#section) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](data-sources--advertise_policy--properties--tls_parameters.md#schema-tls_parameters--xfcc_header_elements) |
| `where` | [where](data-sources--advertise_policy--properties--where.md#section) |
| `where.site` | [where.site](data-sources--advertise_policy--properties--where--site.md#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--advertise_policy--properties--where--site--disable_internet_vip.md#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--advertise_policy--properties--where--site--enable_internet_vip.md#section) |
| `where.site.network_type` | [where.site.network_type](data-sources--advertise_policy--properties--where--site.md#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](data-sources--advertise_policy--properties--where--site--ref.md#section) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--advertise_policy--properties--where--site--ref.md#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](data-sources--advertise_policy--properties--where--virtual_network.md#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--advertise_policy--properties--where--virtual_network--ref.md#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--advertise_policy--properties--where--virtual_network--ref.md#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](data-sources--advertise_policy--properties--where--virtual_site.md#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--advertise_policy--properties--where--virtual_site--disable_internet_vip.md#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--advertise_policy--properties--where--virtual_site--enable_internet_vip.md#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--advertise_policy--properties--where--virtual_site.md#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--advertise_policy--properties--where--virtual_site--ref.md#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--advertise_policy--properties--where--virtual_site--ref.md#schema-where--virtual_site--ref--uid) |

## Next pages

- [dualstack](data-sources--advertise_policy--properties--dualstack.md)
- [ipv4](data-sources--advertise_policy--properties--ipv4.md)
- [ipv6](data-sources--advertise_policy--properties--ipv6.md)
- [public_ip](data-sources--advertise_policy--properties--public_ip.md)
- [tls_parameters](data-sources--advertise_policy--properties--tls_parameters.md)
- [where](data-sources--advertise_policy--properties--where.md)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md)
