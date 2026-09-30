---
page_title: "enable_forward_proxy.tls_intercept"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 4429, "body_sha256": "sha256:b035e715cef1fddc3321b2eaec159732dd3d1546c2dabe5036b3ab089bd0a440", "canonical_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "path": "docs/guides/resources--network_connector--properties--enable_forward_proxy--tls_intercept.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_forward_proxy.tls_intercept

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- enable_forward_proxy.tls_intercept

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md): complete subsection reference.

- [enable_for_all_domains](resources--network_connector--properties--enable_forward_proxy--tls_intercept--enable_for_all_domains.md): complete subsection reference.

- [policy](resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md): complete subsection reference.

<a id="schema-enable_forward_proxy--tls_intercept--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_certificate.md): complete subsection reference.

- [volterra_trusted_ca](resources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- [enable_forward_proxy.tls_intercept.enable_for_all_domains](resources--network_connector--properties--enable_forward_proxy--tls_intercept--enable_for_all_domains.md)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md)
- [enable_forward_proxy.tls_intercept.volterra_certificate](resources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_certificate.md)
- [enable_forward_proxy.tls_intercept.volterra_trusted_ca](resources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_trusted_ca.md)
- [enable_forward_proxy](resources--network_connector--properties--enable_forward_proxy.md)
- [xcsh_network_connector](../resources/network_connector.md)
