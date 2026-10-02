---
page_title: "http_receiver.auth_token"
subcategory: ""
description: "Authentication Token for access."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "http receiver auth token"], "body_bytes": 1612, "body_sha256": "sha256:2252d42c4925350f4900b4ed5dcb62d187c51e127c29bfbd59c1ebb22933f074", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token:token"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver", "path": "documentation/resources/global_log_receiver/properties/http_receiver/auth_token/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver", "auth_token"], "schema_version": 1, "sections": [{"aliases": ["token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token:token", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.auth_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token:token:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_receiver.auth_token.token:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:auth_token:token:clear_secret_info", "type": "conflicts"}], "schema_path": ["http_receiver", "auth_token", "token"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/auth_token/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Authentication Token for access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.auth_token

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/)
- http_receiver.auth_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

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
auth_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/): complete subsection reference.

## Next pages

- [http_receiver.auth_token.token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/auth_token/token/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
