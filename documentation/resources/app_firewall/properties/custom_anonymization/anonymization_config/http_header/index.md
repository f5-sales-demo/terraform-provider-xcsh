---
page_title: "custom_anonymization.anonymization_config.http_header"
subcategory: "Security"
description: "Configure anonymization for HTTP Headers."
xcsh_docs: {"aliases": ["custom anonymization anonymization config http header"], "body_bytes": 2415, "body_sha256": "sha256:a79536f20eee366ad1af6f0de6266cacf06c5b7412205c361b45517817a73568", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3321122222022120-0132202001131102-3002332022300023-3021202332030301-1320003031213312-0031033213302222-3033301111122313-0310322102302033", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "schema-custom_anonymization--anonymization_config--http_header--header_name", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config.http_header:RequiredObjectAttributes:header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config http header header name"], "anchor": "schema-custom_anonymization--anonymization_config--http_header--header_name", "description": "Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (*), or by using only an asterisk to match any HTTP header name.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "http_header", "header_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configure anonymization for HTTP Headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.http_header

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.http_header

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("header_name")}
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
http_header {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_anonymization--anonymization_config--http_header--header_name"></a>

### header_name property

Type: `"string"`. Optional.

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  }
}
```
