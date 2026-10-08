---
page_title: "tls_parameters"
subcategory: ""
description: "TLS configuration for downstream connections."
xcsh_docs: {"aliases": ["tls parameters"], "body_bytes": 3521, "body_sha256": "sha256:94aa1d0c9b6fcf2cd99a5c8898cd64907b6c48b59ccafd017b2bacd214ce08cb", "capabilities": ["load-balancing.tls", "networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "documentation/resources/advertise_policy/properties/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["tls parameters client certificate optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "client_certificate_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters client certificate required"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "client_certificate_required"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters common params"], "anchor": "section", "description": "Information of different aspects for TLS authentication related to ciphers, certificates and trust store.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters no client certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "no_client_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters xfcc header elements"], "anchor": "schema-tls_parameters--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are defined, the header will not be added.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "TLS configuration for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- tls_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for downstream connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("client_certificate_optional",
    "client_certificate_required"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "no_client_certificate"),
  validators.ConflictingObjectAttributes("client_certificate_required",
    "no_client_certificate")}
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
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/client_certificate_optional/): complete subsection reference.

- [client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/client_certificate_required/): complete subsection reference.

- [common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/common_params/): complete subsection reference.

- [no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/no_client_certificate/): complete subsection reference.

<a id="schema-tls_parameters--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
