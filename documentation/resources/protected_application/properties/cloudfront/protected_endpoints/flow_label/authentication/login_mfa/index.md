---
page_title: "cloudfront.protected_endpoints.flow_label.authentication.login_mfa"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cloudfront protected endpoints flow label authentication login mfa", "login", "login result", "sign in"], "body_bytes": 1664, "body_sha256": "sha256:650a52d20df4d11eb74c332f45fa54b5b72d1cd9f855b95e58be5408599297f0", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication:login_mfa", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:authentication", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1022312121112001-0233130002211103-3313213321332103-0211122111220003-2132120322010312-3020000123013011-3231313033300323-0120210232100033", "registry_path": "docs/guides/resources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "authentication", "login_mfa"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/login_mfa/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.authentication.login_mfa

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [cloudfront.protected_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/authentication/)
- cloudfront.protected_endpoints.flow_label.authentication.login_mfa

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
login_mfa = {}
```

This is an empty object or choice marker. It has no direct properties.
