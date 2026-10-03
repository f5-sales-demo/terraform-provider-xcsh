---
page_title: "policy_based_challenge.rule_list.rules.spec.headers"
subcategory: "Load Balancing"
description: "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header"
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec headers"], "body_bytes": 6917, "body_sha256": "sha256:9b2630fbf179c8bcab7ac32a123f5ba3574f01cd64ef92d7813bce123721b426", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_not_present", "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_present", "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:item"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:item", "type": "conflicts"}, {"anchor": "schema-policy_based_challenge--rule_list--rules--spec--headers--name", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec headers check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_not_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec headers check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:check_present", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec headers invert matcher"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--headers--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["policy based challenge rule list rules spec headers item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers:item", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers", "item"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec headers name"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--headers--name", "description": "A case-insensitive HTTP header name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "headers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.headers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/check_present/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--headers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/item/): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--headers--name"></a>

### name property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/check_not_present/)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/check_present/)
- [policy_based_challenge.rule_list.rules.spec.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/headers/item/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
