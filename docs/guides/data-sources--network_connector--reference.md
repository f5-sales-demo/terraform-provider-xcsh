---
page_title: "Property reference"
subcategory: "Networking"
description: "Property reference for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 20099, "body_sha256": "sha256:88aab638c4c12862a9be30b453fd8a030abf4244c4848c44ab521e749ee4b639", "canonical_id": "xcsh-docs:data-sources:network_connector:reference", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:disable_forward_proxy", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr", "xcsh-docs:data-sources:network_connector:properties:sli_to_slo_snat", "xcsh-docs:data-sources:network_connector:properties:slo_to_global_dr"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:reference", "parent_id": "xcsh-docs:data-sources:network_connector:fundamentals", "path": "docs/guides/data-sources--network_connector--reference.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- Property reference

## Direct properties

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

Description of the NetworkConnector.

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

- [disable_forward_proxy](data-sources--network_connector--properties--disable_forward_proxy.md): complete subsection reference.

- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

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

Name of the NetworkConnector.

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

Namespace where the NetworkConnector exists.

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

- [sli_to_global_dr](data-sources--network_connector--properties--sli_to_global_dr.md): complete subsection reference.

- [sli_to_slo_snat](data-sources--network_connector--properties--sli_to_slo_snat.md): complete subsection reference.

- [slo_to_global_dr](data-sources--network_connector--properties--slo_to_global_dr.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_connector--reference.md#schema-annotations) |
| `description` | [description](data-sources--network_connector--reference.md#schema-description) |
| `disable_forward_proxy` | [disable_forward_proxy](data-sources--network_connector--properties--disable_forward_proxy.md#section) |
| `enable_forward_proxy` | [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md#section) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](data-sources--network_connector--properties--enable_forward_proxy.md#schema-enable_forward_proxy--connection_timeout) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](data-sources--network_connector--properties--enable_forward_proxy.md#schema-enable_forward_proxy--max_connect_attempts) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](data-sources--network_connector--properties--enable_forward_proxy--no_interception.md#section) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--certificate_url) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--custom_hash_algorithms.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--custom_hash_algorithms.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--custom_hash_algorithms--hash_algorithms) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--description_spec) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--disable_ocsp_stapling.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--decryption_provider) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--location) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--store_provider) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--provider_ref) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info.md#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--url) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate--use_system_defaults.md#section) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--enable_for_all_domains.md#section) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules.md#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--disable_interception.md#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match.md#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match.md#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match.md#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match.md#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy--interception_rules--enable_interception.md#section) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept.md#schema-enable_forward_proxy--tls_intercept--trusted_ca_url) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_certificate.md#section) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_trusted_ca.md#section) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](data-sources--network_connector--properties--enable_forward_proxy.md#schema-enable_forward_proxy--white_listed_ports) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](data-sources--network_connector--properties--enable_forward_proxy.md#schema-enable_forward_proxy--white_listed_prefixes) |
| `id` | [id](data-sources--network_connector--reference.md#schema-id) |
| `labels` | [labels](data-sources--network_connector--reference.md#schema-labels) |
| `name` | [name](data-sources--network_connector--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--network_connector--reference.md#schema-namespace) |
| `sli_to_global_dr` | [sli_to_global_dr](data-sources--network_connector--properties--sli_to_global_dr.md#section) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md#section) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md#schema-sli_to_global_dr--global_vn--name) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md#schema-sli_to_global_dr--global_vn--namespace) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](data-sources--network_connector--properties--sli_to_global_dr--global_vn.md#schema-sli_to_global_dr--global_vn--tenant) |
| `sli_to_slo_snat` | [sli_to_slo_snat](data-sources--network_connector--properties--sli_to_slo_snat.md#section) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](data-sources--network_connector--properties--sli_to_slo_snat--default_gw_snat.md#section) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](data-sources--network_connector--properties--sli_to_slo_snat--interface_ip.md#section) |
| `slo_to_global_dr` | [slo_to_global_dr](data-sources--network_connector--properties--slo_to_global_dr.md#section) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](data-sources--network_connector--properties--slo_to_global_dr--global_vn.md#section) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](data-sources--network_connector--properties--slo_to_global_dr--global_vn.md#schema-slo_to_global_dr--global_vn--name) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](data-sources--network_connector--properties--slo_to_global_dr--global_vn.md#schema-slo_to_global_dr--global_vn--namespace) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](data-sources--network_connector--properties--slo_to_global_dr--global_vn.md#schema-slo_to_global_dr--global_vn--tenant) |

## Next pages

- [disable_forward_proxy](data-sources--network_connector--properties--disable_forward_proxy.md)
- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md)
- [sli_to_global_dr](data-sources--network_connector--properties--sli_to_global_dr.md)
- [sli_to_slo_snat](data-sources--network_connector--properties--sli_to_slo_snat.md)
- [slo_to_global_dr](data-sources--network_connector--properties--slo_to_global_dr.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
