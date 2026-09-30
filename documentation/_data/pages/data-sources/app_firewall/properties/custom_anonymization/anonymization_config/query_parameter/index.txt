---
page_title: "custom_anonymization.anonymization_config.query_parameter"
subcategory: "Security"
description: "custom_anonymization.anonymization_config.query_parameter for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2716, "body_sha256": "sha256:710fbf812f48439ca07359cda1edf4989bdc23ca957527815d15dadae61df264", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:custom_anonymization:anonymization_config", "path": "documentation/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config", "query_parameter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/query_parameter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization.anonymization_config.query_parameter for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_anonymization.anonymization_config.query_parameter

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/)
- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/)
- custom_anonymization.anonymization_config.query_parameter

<a id="section"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Parameters.

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

<a id="schema-custom_anonymization--anonymization_config--query_parameter--query_param_name"></a>

### query_param_name property

Type: `"string"`. Computed.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Upstream description:

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

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

- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/custom_anonymization/anonymization_config/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
