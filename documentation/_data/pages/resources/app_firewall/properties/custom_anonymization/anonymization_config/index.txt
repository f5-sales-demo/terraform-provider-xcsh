---
page_title: "custom_anonymization.anonymization_config"
subcategory: "Security"
description: "List of HTTP headers, cookies and query parameters whose values will be masked."
xcsh_docs: {"aliases": ["custom anonymization anonymization config"], "body_bytes": 3477, "body_sha256": "sha256:fe24426fed8f424df1f4e94ea7d227eac704253fa3214a4317143786e55e0992", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config cookie"], "anchor": "section", "description": "Configure anonymization for HTTP Cookies.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--cookie--cookie_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.cookie:RequiredObjectAttributes:cookie_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "syntax": "block", "type": "object"}, {"aliases": ["custom anonymization anonymization config http header"], "anchor": "section", "description": "Configure anonymization for HTTP Headers.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--http_header--header_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.http_header:RequiredObjectAttributes:header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "syntax": "block", "type": "object"}, {"aliases": ["custom anonymization anonymization config query parameter"], "anchor": "section", "description": "Configure anonymization for HTTP Parameters.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--query_parameter--query_param_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.query_parameter:RequiredObjectAttributes:query_param_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of HTTP headers, cookies and query parameters whose values will be masked.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/)
- custom_anonymization.anonymization_config

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP headers, cookies and query parameters whose values will be masked.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "http_header"),
  validators.ConflictingListObjectAttributes("cookie",
    "query_parameter"),
  validators.ConflictingListObjectAttributes("http_header",
    "query_parameter")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
anonymization_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/): complete subsection reference.

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/): complete subsection reference.

- [query_parameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/): complete subsection reference.

## Next pages

- [custom_anonymization.anonymization_config.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/cookie/)
- [custom_anonymization.anonymization_config.http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/)
- [custom_anonymization.anonymization_config.query_parameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
