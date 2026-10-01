---
page_title: "enable_forward_proxy.tls_intercept"
subcategory: "Networking"
description: "enable_forward_proxy.tls_intercept for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 3953, "body_sha256": "sha256:03c3b326d15bb8bc5f702f82a2fa30cdb9b8f7cc1862c869bdda8fb8441f7ef3", "canonical_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:enable_for_all_domains", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:policy", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_certificate", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:volterra_trusted_ca"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy", "path": "docs/guides/data-sources--network_connector--properties--enable_forward_proxy--tls_intercept.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_forward_proxy.tls_intercept for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- [Property reference](data-sources--network_connector--reference.md)
- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md)
- enable_forward_proxy.tls_intercept

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

- [custom_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md): complete subsection reference.

- [enable_for_all_domains](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--enable_for_all_domains.md): complete subsection reference.

- [policy](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md): complete subsection reference.

<a id="schema-enable_forward_proxy--tls_intercept--trusted_ca_url"></a>

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

- [volterra_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_certificate.md): complete subsection reference.

- [volterra_trusted_ca](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_trusted_ca.md): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--custom_certificate.md)
- [enable_forward_proxy.tls_intercept.enable_for_all_domains](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--enable_for_all_domains.md)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--policy.md)
- [enable_forward_proxy.tls_intercept.volterra_certificate](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_certificate.md)
- [enable_forward_proxy.tls_intercept.volterra_trusted_ca](data-sources--network_connector--properties--enable_forward_proxy--tls_intercept--volterra_trusted_ca.md)
- [enable_forward_proxy](data-sources--network_connector--properties--enable_forward_proxy.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
