---
page_title: "custom_anonymization.anonymization_config.http_header"
subcategory: "Security"
description: "Configure anonymization for HTTP Headers."
xcsh_docs: {"aliases": ["custom anonymization anonymization config http header"], "body_bytes": 2157, "body_sha256": "sha256:ba03f5afc2fc39590808be214b23d7f35d7bdc1b4613bea26da1b283b576036e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0100321222011122-2331012001321130-2201200113002333-1232110223010102-3232230233001000-3100023323011100-2232001213202312-1030032031132020", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "http_header"], "schema_version": 1, "sections": [{"aliases": ["custom anonymization anonymization config http header header name"], "anchor": "schema-custom_anonymization--anonymization_config--http_header--header_name", "description": "Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (*), or by using only an asterisk to match any HTTP header name.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_anonymization", "anonymization_config", "http_header", "header_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/http_header/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configure anonymization for HTTP Headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization.anonymization_config.http_header

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Headers.

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

## Direct properties

<a id="schema-custom_anonymization--anonymization_config--http_header--header_name"></a>

### header_name property

Type: `"string"`. Computed.

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
