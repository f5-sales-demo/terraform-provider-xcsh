---
page_title: "tls_parameters"
subcategory: ""
description: "tls_parameters for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 3439, "body_sha256": "sha256:8dbf8b436144e23f36bae7617016641c5ca5453e7de3b2d2efb49fc1916f2c2d", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_parameters:client_certificate_optional", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:client_certificate_required", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params", "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:no_client_certificate"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "docs/guides/data-sources--virtual_host--properties--tls_parameters.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

TLS configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

## Direct properties

- [client_certificate_optional](data-sources--virtual_host--properties--tls_parameters--client_certificate_optional.md): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--properties--tls_parameters--client_certificate_required.md): complete subsection reference.

- [common_params](data-sources--virtual_host--properties--tls_parameters--common_params.md): complete subsection reference.

- [no_client_certificate](data-sources--virtual_host--properties--tls_parameters--no_client_certificate.md): complete subsection reference.

<a id="schema-tls_parameters--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [tls_parameters.client_certificate_optional](data-sources--virtual_host--properties--tls_parameters--client_certificate_optional.md)
- [tls_parameters.client_certificate_required](data-sources--virtual_host--properties--tls_parameters--client_certificate_required.md)
- [tls_parameters.common_params](data-sources--virtual_host--properties--tls_parameters--common_params.md)
- [tls_parameters.no_client_certificate](data-sources--virtual_host--properties--tls_parameters--no_client_certificate.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
