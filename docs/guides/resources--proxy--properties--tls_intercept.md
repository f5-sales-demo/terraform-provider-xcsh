---
page_title: "tls_intercept"
subcategory: ""
description: "tls_intercept for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3830, "body_sha256": "sha256:76b62638068af0723ccafc80c24265662f5e7c8f1ae3707c3435a0ea5fe85636", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_certificate", "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "docs/guides/resources--proxy--properties--tls_intercept.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- tls_intercept

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

- [custom_certificate](resources--proxy--properties--tls_intercept--custom_certificate.md): complete subsection reference.

- [enable_for_all_domains](resources--proxy--properties--tls_intercept--enable_for_all_domains.md): complete subsection reference.

- [policy](resources--proxy--properties--tls_intercept--policy.md): complete subsection reference.

<a id="schema-tls_intercept--trusted_ca_url"></a>

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

- [volterra_certificate](resources--proxy--properties--tls_intercept--volterra_certificate.md): complete subsection reference.

- [volterra_trusted_ca](resources--proxy--properties--tls_intercept--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate](resources--proxy--properties--tls_intercept--custom_certificate.md)
- [tls_intercept.enable_for_all_domains](resources--proxy--properties--tls_intercept--enable_for_all_domains.md)
- [tls_intercept.policy](resources--proxy--properties--tls_intercept--policy.md)
- [tls_intercept.volterra_certificate](resources--proxy--properties--tls_intercept--volterra_certificate.md)
- [tls_intercept.volterra_trusted_ca](resources--proxy--properties--tls_intercept--volterra_trusted_ca.md)
- [Property reference](resources--proxy--reference.md)
- [xcsh_proxy](../resources/proxy.md)
