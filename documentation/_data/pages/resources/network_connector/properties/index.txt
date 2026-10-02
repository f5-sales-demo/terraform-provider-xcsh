---
page_title: "Property reference"
subcategory: "Networking"
description: "Property reference for xcsh_network_connector."
xcsh_docs: {"aliases": ["network connector"], "body_bytes": 25331, "body_sha256": "sha256:7d17145fdce02fdb24b184b92ab5417171cf932541c2f9cf669c22be7126e70f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:properties:disable_forward_proxy", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "xcsh-docs:resources:network_connector:properties:sli_to_global_dr", "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "xcsh-docs:resources:network_connector:properties:slo_to_global_dr", "xcsh-docs:resources:network_connector:properties:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:reference", "parent_id": "xcsh-docs:resources:network_connector:fundamentals", "path": "documentation/resources/network_connector/properties/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331", "registry_path": "docs/guides/resources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_connector:properties:disable_forward_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "enable forward proxy", "operation timeout"], "anchor": "section", "description": "Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes: Forward proxy does application protocol detection and server name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols doesn't have client sending the first data. In such cases,", "document_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy:ConflictingObjectAttributes:no_interception,tls_intercept", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:no_interception", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_forward_proxy:ConflictingObjectAttributes:no_interception,tls_intercept", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "type": "conflicts"}], "schema_path": ["enable_forward_proxy"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:network_connector:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_global_dr", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_global_dr"], "syntax": "block", "type": "object"}, {"aliases": ["sli to slo snat"], "anchor": "section", "description": "X-example: \"\" description.", "document_id": "xcsh-docs:resources:network_connector:properties:sli_to_slo_snat", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_slo_snat"], "syntax": "block", "type": "object"}, {"aliases": ["slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:resources:network_connector:properties:slo_to_global_dr", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slo_to_global_dr"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:network_connector:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
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

- [disable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/disable_forward_proxy/): complete subsection reference.

- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/): complete subsection reference.

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

Name of the Network Connector. Must be unique within the namespace.

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

Namespace where the Network Connector is created.

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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/): complete subsection reference.

- [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-disable) |
| `disable_forward_proxy` | [disable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/disable_forward_proxy/#section) |
| `enable_forward_proxy` | [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/#section) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/#schema-enable_forward_proxy--connection_timeout) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/#schema-enable_forward_proxy--max_connect_attempts) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/no_interception/#section) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/#schema-enable_forward_proxy--tls_intercept--custom_certificate--certificate_url) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/custom_hash_algorithms/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/custom_hash_algorithms/#schema-enable_forward_proxy--tls_intercept--custom_certificate--custom_hash_algorithms--hash_algorithms) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/#schema-enable_forward_proxy--tls_intercept--custom_certificate--description_spec) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/disable_ocsp_stapling/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--decryption_provider) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--location) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--blindfold_secret_info--store_provider) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/#section) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--provider_ref) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/#schema-enable_forward_proxy--tls_intercept--custom_certificate--private_key--clear_secret_info--url) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/use_system_defaults/#section) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/enable_for_all_domains/#section) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/disable_interception/#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/#section) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--exact_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--regex_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/domain_match/#schema-enable_forward_proxy--tls_intercept--policy--interception_rules--domain_match--suffix_value) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/interception_rules/enable_interception/#section) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/#schema-enable_forward_proxy--tls_intercept--trusted_ca_url) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_certificate/#section) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_trusted_ca/#section) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/#schema-enable_forward_proxy--white_listed_ports) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/#schema-enable_forward_proxy--white_listed_prefixes) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/#schema-namespace) |
| `sli_to_global_dr` | [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/#section) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/#section) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/#schema-sli_to_global_dr--global_vn--name) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/#schema-sli_to_global_dr--global_vn--namespace) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/global_vn/#schema-sli_to_global_dr--global_vn--tenant) |
| `sli_to_slo_snat` | [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/#section) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/default_gw_snat/#section) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/interface_ip/#section) |
| `slo_to_global_dr` | [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/#section) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/global_vn/#section) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/global_vn/#schema-slo_to_global_dr--global_vn--name) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/global_vn/#schema-slo_to_global_dr--global_vn--namespace) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/global_vn/#schema-slo_to_global_dr--global_vn--tenant) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [disable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/disable_forward_proxy/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_global_dr/)
- [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/sli_to_slo_snat/)
- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/slo_to_global_dr/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/timeouts/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
