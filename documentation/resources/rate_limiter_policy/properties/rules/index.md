---
page_title: "rules"
subcategory: "Security"
description: "A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 2457, "body_sha256": "sha256:2b6003ac2ed45d976679e5cba78076f88db45edaf6dd2cfadc9e8c568891fe26", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:metadata", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules", "parent_id": "xcsh-docs:resources:rate_limiter_policy:reference", "path": "documentation/resources/rate_limiter_policy/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303", "registry_path": "docs/guides/resources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--metadata--name", "enforcement": "provider-schema", "group": "rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:metadata", "type": "requires"}], "schema_path": ["rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["spec"], "anchor": "section", "description": "Shape of Rate Limiter Rule.", "document_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_country,country_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_country", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:apply_rate_limiter,bypass_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:apply_rate_limiter,custom_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:apply_rate_limiter,bypass_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:bypass_rate_limiter,custom_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_country,country_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:country_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:apply_rate_limiter,custom_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:bypass_rate_limiter,custom_rate_limiter", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "type": "conflicts"}], "schema_path": ["rules", "spec"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of RateLimiterRules that are evaluated sequentially till a matching rule is identified.
Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/): complete subsection reference.

## Next pages

- [rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/metadata/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
