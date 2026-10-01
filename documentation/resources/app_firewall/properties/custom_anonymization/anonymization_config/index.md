---
page_title: "custom_anonymization.anonymization_config"
subcategory: "Security"
description: "custom_anonymization.anonymization_config for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 3477, "body_sha256": "sha256:846329edb66deebf85cc5328abec7f593ccc3087d2e55a96659fc2c0512c28ad", "child_ids": ["xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter"], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "parent_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization", "path": "documentation/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_anonymization", "anonymization_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/anonymization_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_anonymization.anonymization_config for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
