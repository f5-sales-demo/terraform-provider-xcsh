---
page_title: "data_guard_rules"
subcategory: "Load Balancing"
description: "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature."
xcsh_docs: {"aliases": ["data guard rules"], "body_bytes": 6143, "body_sha256": "sha256:7ab1c54f9eceb488f545959c21574f82d07e03679607a3acca15ae24b1d478c3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:apply_data_guard", "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:path", "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:skip_data_guard"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/data_guard_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data_guard_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["apply data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:apply_data_guard", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "apply_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact value"], "anchor": "schema-data_guard_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data_guard_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data_guard_rules", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules:skip_data_guard", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "skip_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["suffix value"], "anchor": "schema-data_guard_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:data_guard_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/data_guard_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_guard_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- data_guard_rules

<a id="section"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/any_domain/): complete subsection reference.

- [apply_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/apply_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/path/): complete subsection reference.

- [skip_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/skip_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [data_guard_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/any_domain/)
- [data_guard_rules.apply_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/apply_data_guard/)
- [data_guard_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/metadata/)
- [data_guard_rules.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/path/)
- [data_guard_rules.skip_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/data_guard_rules/skip_data_guard/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
