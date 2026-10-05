---
page_title: "tls_intercept"
subcategory: ""
description: "Configuration to enable TLS interception."
xcsh_docs: {"aliases": ["tls intercept"], "body_bytes": 4538, "body_sha256": "sha256:1a9e1f34a14b22a9f0628e58ce011290aff64fea2a3749db6cf1f39497353789", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_certificate", "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/tls_intercept/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [{"anchor": "schema-tls_intercept--trusted_ca_url", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:trusted_ca_url,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:custom_certificate,volterra_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:enable_for_all_domains,policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:enable_for_all_domains,policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:custom_certificate,volterra_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept:ConflictingObjectAttributes:trusted_ca_url,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept"], "schema_version": 1, "sections": [{"aliases": ["tls intercept custom certificate"], "anchor": "section", "description": "Handle to fetch certificate and key.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:ConflictingObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-tls_intercept--custom_certificate--certificate_url", "enforcement": "provider-schema", "group": "tls_intercept.custom_certificate:RequiredObjectAttributes:certificate_url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "type": "requires"}], "schema_path": ["tls_intercept", "custom_certificate"], "syntax": "block", "type": "object"}, {"aliases": ["tls intercept enable for all domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:enable_for_all_domains", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "enable_for_all_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept policy"], "anchor": "section", "description": "Policy to enable or disable TLS interception.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_intercept.policy:RequiredObjectAttributes:interception_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:tls_intercept:policy:interception_rules", "type": "requires"}], "schema_path": ["tls_intercept", "policy"], "syntax": "block", "type": "object"}, {"aliases": ["tls intercept trusted ca url"], "anchor": "schema-tls_intercept--trusted_ca_url", "description": "Exclusive with Custom Root CA Certificate for validating upstream server certificate.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls intercept volterra certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "volterra_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:tls_intercept:volterra_trusted_ca", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration to enable TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
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

- [custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/): complete subsection reference.

- [enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/enable_for_all_domains/): complete subsection reference.

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/volterra_certificate/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/volterra_trusted_ca/): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/custom_certificate/)
- [tls_intercept.enable_for_all_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/enable_for_all_domains/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/policy/)
- [tls_intercept.volterra_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/volterra_certificate/)
- [tls_intercept.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/tls_intercept/volterra_trusted_ca/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
