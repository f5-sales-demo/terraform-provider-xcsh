---
page_title: "domains.credentials.login_endpoint.json_payload"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["domains credentials login endpoint json payload", "login", "login result", "sign in"], "body_bytes": 2744, "body_sha256": "sha256:c6c9e18522c5cec722ff97b5437f0c8e21254d5b85e240864e9ec7cdb9f8f06f", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload", "parent_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint", "path": "documentation/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2022021213220201-2233331101332133-1312233301111020-3102231111200321-2303231331113321-0120220311133230-2201100332012212-2011001001012123", "registry_path": "docs/guides/resources--api_testing--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint.json_payload:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint.json_payload:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "login_endpoint", "json_payload"], "schema_version": 1, "sections": [{"aliases": ["domains credentials login endpoint json payload blindfold secret info", "login", "login result", "sign in"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--login_endpoint--json_payload--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint.json_payload.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:blindfold_secret_info", "type": "requires"}], "schema_path": ["domains", "credentials", "login_endpoint", "json_payload", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["domains credentials login endpoint json payload clear secret info", "login", "login result", "sign in"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-domains--credentials--login_endpoint--json_payload--clear_secret_info--url", "enforcement": "provider-schema", "group": "domains.credentials.login_endpoint.json_payload.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_testing:properties:domains:credentials:login_endpoint:json_payload:clear_secret_info", "type": "requires"}], "schema_path": ["domains", "credentials", "login_endpoint", "json_payload", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.login_endpoint.json_payload

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/)
- [domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/)
- domains.credentials.login_endpoint.json_payload

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
json_payload {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/blindfold_secret_info/)
- [domains.credentials.login_endpoint.json_payload.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/json_payload/clear_secret_info/)
- [domains.credentials.login_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/properties/domains/credentials/login_endpoint/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
