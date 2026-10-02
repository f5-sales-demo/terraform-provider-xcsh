---
page_title: "custom_anonymization.anonymization_config.query_parameter"
subcategory: "Security"
description: "Configure anonymization for HTTP Parameters."
xcsh_docs: {"aliases": ["custom anonymization anonymization config query parameter"], "body_bytes": 3214, "body_sha256": "sha256:eaf03b94a8f450ebbd6d2d540c5799dac1032fe605be413356718014b4738fa5", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1200131222023102-3200231020312002-2100020013033130-2000332231021313-3121212320133221-2232313223013220-0212122011322311-3032333102100223", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--query_parameter--query_param_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.query_parameter:RequiredObjectAttributes:query_param_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "schema_version": 1, "sections": [{"aliases": ["query param name"], "anchor": "schema-custom_anonymization--anonymization_config--query_parameter--query_param_name", "description": "Masks the query parameter value. The setting does not mask the query parameter name. Wildcard matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk (*), or by using only an asterisk to match any query parameter name.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter", "query_param_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure anonymization for HTTP Parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.query_parameter

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.query_parameter

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("query_param_name")}
```

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
query_parameter {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_anonymization--anonymization_config--query_parameter--query_param_name"></a>

### query_param_name property

Type: `"string"`. Optional.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Upstream description:

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
