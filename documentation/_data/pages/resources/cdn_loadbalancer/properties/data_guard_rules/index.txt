---
page_title: "data_guard_rules"
subcategory: "Load Balancing"
description: "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature."
xcsh_docs: {"aliases": ["data guard rules"], "body_bytes": 6917, "body_sha256": "sha256:8e3bb3b70498c2302fe8a0916ccf9d2edc43c6bc7c98ae0831ed3caa8e4ef463", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:apply_data_guard", "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:skip_data_guard"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/data_guard_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "schema-data_guard_rules--exact_value", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--exact_value", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--suffix_value", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--suffix_value", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:apply_data_guard,skip_data_guard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:apply_data_guard", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_guard_rules:ConflictingListObjectAttributes:apply_data_guard,skip_data_guard", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:skip_data_guard", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["data_guard_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["apply data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:apply_data_guard", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "apply_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact value"], "anchor": "schema-data_guard_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-data_guard_rules--metadata--name", "enforcement": "provider-schema", "group": "data_guard_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:metadata", "type": "requires"}], "schema_path": ["data_guard_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-data_guard_rules--path--path", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--path--path", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--path--prefix", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--path--prefix", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--path--regex", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}, {"anchor": "schema-data_guard_rules--path--regex", "enforcement": "provider-schema", "group": "data_guard_rules.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:path", "type": "conflicts"}], "schema_path": ["data_guard_rules", "path"], "syntax": "block", "type": "object"}, {"aliases": ["skip data guard"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules:skip_data_guard", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "skip_data_guard"], "syntax": "attribute", "type": "object"}, {"aliases": ["suffix value"], "anchor": "schema-data_guard_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:data_guard_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_guard_rules", "suffix_value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/data_guard_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data Guard prevents responses from exposing sensitive information by masking the data. The system masks credit card numbers and social security numbers leaked from the application from within the HTTP response with a string of asterisks (*). Note: App Firewall should be enabled, to use Data Guard feature.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_guard_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- data_guard_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("apply_data_guard",
    "skip_data_guard"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
data_guard_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/any_domain/): complete subsection reference.

- [apply_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/apply_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--exact_value"></a>

### exact_value property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/path/): complete subsection reference.

- [skip_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/skip_data_guard/): complete subsection reference.

<a id="schema-data_guard_rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [data_guard_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/any_domain/)
- [data_guard_rules.apply_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/apply_data_guard/)
- [data_guard_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/metadata/)
- [data_guard_rules.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/path/)
- [data_guard_rules.skip_data_guard](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/data_guard_rules/skip_data_guard/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
