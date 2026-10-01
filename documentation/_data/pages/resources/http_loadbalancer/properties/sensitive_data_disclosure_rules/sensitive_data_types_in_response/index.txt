---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response"
subcategory: "Load Balancing"
description: "sensitive_data_disclosure_rules.sensitive_data_types_in_response for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4015, "body_sha256": "sha256:3ba95f7e7231b6911362625876fcbb89f50724ff32fb78941688ec07ec83249d", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "path": "documentation/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_disclosure_rules.sensitive_data_types_in_response for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mask",
    "report")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
sensitive_data_types_in_response {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/): complete subsection reference.

- [body](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/): complete subsection reference.

- [mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/report/): complete subsection reference.

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/report/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
