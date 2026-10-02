---
page_title: "tls_parameters"
subcategory: ""
description: "TLS configuration for downstream connections."
xcsh_docs: {"aliases": ["tls parameters"], "body_bytes": 4598, "body_sha256": "sha256:185c0e04325cac9cf0fa2c88ba63abc365272ddace595256b9890b5fee939797", "capabilities": ["load-balancing.tls", "networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params", "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "documentation/resources/advertise_policy/properties/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0332333113033300-2023221010021211-3120323123033310-2232030322312121-1101101210023031-2203131103020300-1330132310231221-2020023030033330", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["client certificate optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "client_certificate_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["client certificate required"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_required", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "client_certificate_required"], "syntax": "attribute", "type": "object"}, {"aliases": ["common params"], "anchor": "section", "description": "Information of different aspects for TLS authentication related to ciphers, certificates and trust store.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:common_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params"], "syntax": "block", "type": "object"}, {"aliases": ["no client certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:no_client_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "no_client_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["xfcc header elements"], "anchor": "schema-tls_parameters--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are defined, the header will not be added.", "document_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TLS configuration for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [tls_parameters.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/client_certificate_optional/)
- [tls_parameters.client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/client_certificate_required/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/common_params/)
- [tls_parameters.no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/no_client_certificate/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
