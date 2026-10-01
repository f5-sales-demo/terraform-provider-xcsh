---
page_title: "tls_parameters"
subcategory: ""
description: "tls_parameters for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4103, "body_sha256": "sha256:32dc97ba727822f19be47bb0dfcc185af01885e6aaa129d0b9321e69adfe0edd", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:client_certificate_optional", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:client_certificate_required", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:no_client_certificate"], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "documentation/data-sources/advertise_policy/properties/tls_parameters/index.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
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

- [client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_optional/): complete subsection reference.

- [client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_required/): complete subsection reference.

- [common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/): complete subsection reference.

- [no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/no_client_certificate/): complete subsection reference.

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

- [tls_parameters.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_optional/)
- [tls_parameters.client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/client_certificate_required/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/)
- [tls_parameters.no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/no_client_certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
