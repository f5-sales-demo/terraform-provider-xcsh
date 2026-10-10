---
page_title: "custom_anonymization.anonymization_config"
subcategory: "Security"
description: "List of HTTP headers, cookies and query parameters whose values will be masked."
xcsh_docs: {"aliases": ["custom anonymization anonymization config"], "body_bytes": 2640, "body_sha256": "sha256:82076aa0ea669638e2693d15294cde62c09d1bf91b083c8848112d559c27dfdc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config cookie"], "anchor": "section", "description": "Configure anonymization for HTTP Cookies.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--cookie--cookie_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.cookie:RequiredObjectAttributes:cookie_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "cookie"], "syntax": "block", "type": "object"}, {"aliases": ["custom anonymization anonymization config http header"], "anchor": "section", "description": "Configure anonymization for HTTP Headers.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--http_header--header_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.http_header:RequiredObjectAttributes:header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "syntax": "block", "type": "object"}, {"aliases": ["custom anonymization anonymization config query parameter"], "anchor": "section", "description": "Configure anonymization for HTTP Parameters.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--query_parameter--query_param_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.query_parameter:RequiredObjectAttributes:query_param_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "requires"}], "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of HTTP headers, cookies and query parameters whose values will be masked.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
