---
page_title: "enable_forward_proxy.tls_intercept"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 5176, "body_sha256": "sha256:d4c0e83d159ed09a285608e9af5178607afc5f4b4a01bd901e852d8beaba5812", "child_ids": ["xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy:tls_intercept", "parent_id": "xcsh-docs:resources:network_connector:properties:enable_forward_proxy", "path": "documentation/resources/network_connector/properties/enable_forward_proxy/tls_intercept/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/enable_forward_proxy/tls_intercept/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_forward_proxy.tls_intercept

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
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

- [custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/): complete subsection reference.

- [enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/enable_for_all_domains/): complete subsection reference.

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/): complete subsection reference.

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

- [volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_certificate/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_trusted_ca/): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- [enable_forward_proxy.tls_intercept.enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/enable_for_all_domains/)
- [enable_forward_proxy.tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/policy/)
- [enable_forward_proxy.tls_intercept.volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_certificate/)
- [enable_forward_proxy.tls_intercept.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/tls_intercept/volterra_trusted_ca/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/enable_forward_proxy/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
