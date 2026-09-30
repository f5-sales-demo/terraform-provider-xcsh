---
page_title: "tls_intercept"
subcategory: ""
description: "tls_intercept for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3153, "body_sha256": "sha256:e91e0a7eb242bdb81cdc8c3dd2c0679f9021331ad6cc442285913e98893de598", "canonical_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "xcsh-docs:data-sources:proxy:properties:tls_intercept:enable_for_all_domains", "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy", "xcsh-docs:data-sources:proxy:properties:tls_intercept:volterra_certificate", "xcsh-docs:data-sources:proxy:properties:tls_intercept:volterra_trusted_ca"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "docs/guides/data-sources--proxy--properties--tls_intercept.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_intercept

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- tls_intercept

<a id="section"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

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

## Direct properties

- [custom_certificate](data-sources--proxy--properties--tls_intercept--custom_certificate.md): complete subsection reference.

- [enable_for_all_domains](data-sources--proxy--properties--tls_intercept--enable_for_all_domains.md): complete subsection reference.

- [policy](data-sources--proxy--properties--tls_intercept--policy.md): complete subsection reference.

<a id="schema-tls_intercept--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

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

- [volterra_certificate](data-sources--proxy--properties--tls_intercept--volterra_certificate.md): complete subsection reference.

- [volterra_trusted_ca](data-sources--proxy--properties--tls_intercept--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate](data-sources--proxy--properties--tls_intercept--custom_certificate.md)
- [tls_intercept.enable_for_all_domains](data-sources--proxy--properties--tls_intercept--enable_for_all_domains.md)
- [tls_intercept.policy](data-sources--proxy--properties--tls_intercept--policy.md)
- [tls_intercept.volterra_certificate](data-sources--proxy--properties--tls_intercept--volterra_certificate.md)
- [tls_intercept.volterra_trusted_ca](data-sources--proxy--properties--tls_intercept--volterra_trusted_ca.md)
- [Property reference](data-sources--proxy--reference.md)
- [xcsh_proxy](../data-sources/proxy.md)
